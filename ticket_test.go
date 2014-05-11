package main

import (
	"bytes"
	"testing"
)

func TestTickets(t *testing.T) {
	t1 := NewTicket()
	t2 := NewTicket()

	if t1.Value > t2.Value {
		t.Error("Should increment value of new tickets")
	}

	if bytes.Equal([]byte(t1.Slug), []byte(t2.Slug)) {
		t.Error("Ticket should have unique slugs")
	}

	if FindBySlug(t1.Slug) != t1 {
		t.Error("Should be able to find tickets by slug")
	}

	ResetCounter()

	if NewTicket().Value != 1 {
		t.Error("ResetCounter() should reset the ticket counter")
	}
}
