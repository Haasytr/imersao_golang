package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TicketType string

var TicketPriceZero = errors.New("ticket price must be greater than zero")

const (
	TicketTypeHalf TicketType = "half"
	TicketTypeFull TicketType = "full"
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func isValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return TicketPriceZero
	}
	return nil
}

func NewTicket(event *Event, spot *Spot, ticketType TicketType) (*Ticket, error) {
	if !isValidTicketType(ticketType) {
		return nil, errors.New(("Invalid ticket type"))
	}

	ticket := &Ticket{
		ID:         uuid.New().String(),
		EventID:    event.ID,
		Spot:       spot,
		TicketType: ticketType,
		Price:      event.Price,
	}
	ticket.CalculatePrice()
	if err := ticket.Validate(); err != nil {
		return nil, err
	}
	return ticket, nil
}
