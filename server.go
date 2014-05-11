package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	templates = template.Must(template.ParseGlob("templates/*"))

	host = flag.String("host", "localhost:8080", "Host")
)

func serveQR(w http.ResponseWriter, r *http.Request) {
	target := fmt.Sprintf("http://%v/n", *host)
	templates.ExecuteTemplate(w, "qr.html", target)
}

func serveNew(w http.ResponseWriter, r *http.Request) {
	ticket := NewTicket()
	http.Redirect(w, r, "/t/"+ticket.Slug, http.StatusTemporaryRedirect)
}

func serveTicket(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Path[len("/t/"):]
	ticket := FindBySlug(slug)
	if ticket == nil {
		fmt.Fprintf(w, "Ticket not found")
		return
	}

	x := &struct {
		T *Ticket
		R float64
		C int
	}{ticket, EstimatedQueueLength(ticket.Value), current}

	templates.ExecuteTemplate(w, "ticket.html", x)
}

func serveCounter(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "counter.html", NewCustomer())
}

func serveManual(w http.ResponseWriter, r *http.Request) {
	ticket := NewTicket()
	templates.ExecuteTemplate(w, "manual.html", ticket)
}

func serveRoot(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "root.html", EstimatedTotalQueueLength())
}

func main() {
	flag.Parse()

	http.HandleFunc("/qr", serveQR)
	http.HandleFunc("/n", serveNew)
	http.HandleFunc("/t/", serveTicket)
	http.HandleFunc("/c/", serveCounter)
	http.HandleFunc("/m/", serveManual)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", serveRoot)
	log.Fatal(http.ListenAndServe(*host, nil))
}
