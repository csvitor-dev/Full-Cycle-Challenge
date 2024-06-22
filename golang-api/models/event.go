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
	ID           int `json:"id"`
	Name         string `json:"name"`
	Location     string `json:"location"`
	Organization string `json:"organization"`
	Rating       Rating `json:"rating"`
	Date         time.Time `json:"date"`
	ImageURL     string `json:"image_url"`
	Price        float64 `json:"price"`
	Spots        []Spot `json:"spots"`
	CreateAt time.Time `json:"created_at"`
}

// methods
func (e *Event) Validate() error {
	if e.Name == "" {
		return errors.New("event name is required")
	}
	if e.Date.Before(time.Now()) {
		return errors.New("event date must be in the future")
	}
	if e.Price <= 0 {
		return errors.New("event price must be greater than zero")
	}

	return nil
}
