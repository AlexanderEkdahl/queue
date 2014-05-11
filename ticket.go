package main

import (
	"crypto/rand"
	"time"
)

const (
	// SmothingFactor determines the smoothing used to estimate the current speed
	SmothingFactor = 0.1
	// SlugLength determines the length of the unique ticket slug
	SlugLength = 5
)

var (
	counter      int
	current      int
	tickets      map[string]*Ticket
	averageSpeed float64
	lastTime     time.Time
)

// Ticket represents a single ticket hold by a user
type Ticket struct {
	Value int
	Slug  string
}

// ResetCounter resets the ticket value counter
func ResetCounter() {
	// TODO(AlexanderEkdahl): should lock a mutex for extra safety
	counter = 0
	current = 0
	lastTime = time.Now()
	// This makes old ticket inaccessible
	tickets = make(map[string]*Ticket)
}

// returns a unique slug
func randomSlug() string {
	alphanum := "23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijklmnpqrstuvwxyz"
	var bytes = make([]byte, SlugLength)
	rand.Read(bytes)
	for {
		for i, b := range bytes {
			bytes[i] = alphanum[b%byte(len(alphanum))]
		}

		// Make sure the slug is unique
		slug := string(bytes)
		if _, ok := tickets[slug]; !ok {
			return slug
		}
	}
}

// NewTicket creates and returns a new ticket
func NewTicket() *Ticket {
	// TODO(AlexanderEkdahl): should lock a mutex for extra safety
	counter++
	ticket := &Ticket{
		Value: counter,
		Slug:  randomSlug(),
	}
	tickets[ticket.Slug] = ticket
	return ticket
}

// FindBySlug returns the ticket matching the slug, returns nil if the ticket
// does not exist
func FindBySlug(slug string) *Ticket {
	return tickets[slug]
}

// NewCustomer increases the current counter and returns a new number
func NewCustomer() int {
	// TODO(AlexanderEkdahl): should lock a mutex for extra safety
	restimateSpeed()
	current++
	return current
}

// EstimatedTotalQueueLength returns the estimated total queue length
func EstimatedTotalQueueLength() float64 {
	// these methods fail to take into consideration the current time.
	return float64(counter-current) / averageSpeed
}

// EstimatedQueueLength returns the estimated total queue length
func EstimatedQueueLength(n int) float64 {
	// these methods fail to take into consideration the current time.
	return float64(n-current) / averageSpeed
}

// Estimate average speed using an exponential moving average
// http://en.wikipedia.org/wiki/Moving_average#Exponential_moving_average
func restimateSpeed() {
	lastSpeed := 1 / time.Since(lastTime).Seconds()
	lastTime = time.Now()
	averageSpeed = SmothingFactor*lastSpeed + (1-SmothingFactor)*averageSpeed
}

func init() {
	ResetCounter()
}
