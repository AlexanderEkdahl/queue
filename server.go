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
	x := &struct {
		QR string
		R string
	}{fmt.Sprintf("http://%v/n", *host), fmt.Sprintf("%.0f", EstimatedTotalQueueLength())}
	templates.ExecuteTemplate(w, "qr.html", x)
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
		R string
		C int
	}{ticket, fmt.Sprintf("%.0f", EstimatedTotalQueueLength()), current}

	templates.ExecuteTemplate(w, "ticket.html", x)
}

func serveCounter(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "counter.html", NewCustomer())
}

func serveManual(w http.ResponseWriter, r *http.Request) {
	ticket := NewTicket()
	templates.ExecuteTemplate(w, "manual.html", ticket)
}

func main() {
	flag.Parse()

	http.HandleFunc("/n", serveNew)
	http.HandleFunc("/t/", serveTicket)
	http.HandleFunc("/c/", serveCounter)
	http.HandleFunc("/m/", serveManual)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", serveQR)
	log.Fatal(http.ListenAndServe(*host, nil))
}
