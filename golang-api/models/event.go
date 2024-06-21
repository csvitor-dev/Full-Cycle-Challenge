package models

import (
	"errors"
	"time"
)

type Rating string

const (
	RatingLivre Rating = "L"
	Rating10    Rating = "10"
	Rating12    Rating = "12"
	Rating14    Rating = "14"
	Rating16    Rating = "16"
	Rating18    Rating = "18"
)

type Event struct {
	ID           string
	Name         string
	Location     string
	Organization string
	Rating       Rating
	Date         time.Time
	ImageURL     string
	Capacity     int
	Price        float64
	PartnerID    string
	Spots        []Spot
	Tickets      []Ticket
}

// methods
func (e *Event) Validate() error {
	if e.Name == "" {
		return errors.New("event name is required")
	}
	if e.Date.Before(time.Now()) {
		return errors.New("event date must be in the future")
	}
	if e.Capacity <= 0 {
		return errors.New("event capacity must be greater than zero")
	}
	if e.Price <= 0 {
		return errors.New("event price must be greater than zero")
	}

	return nil
}
