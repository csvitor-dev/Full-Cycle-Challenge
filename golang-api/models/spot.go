package models

import (
	"errors"
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       int `json:"id"`
	EventID  int `json:"event_id"`
	Name     string `json:"name"`// all name uses the rule: Letter+Number. Ex: A1, B2, C3, etc.
	Status   SpotStatus `json:"status"`
}

// Validate checks if the spot data is valid.
func (s *Spot) Validate() error {
	if len(s.Name) == 0 {
		return errors.New("spot name is required")
	}
	if len(s.Name) < 2 {
		return errors.New("spot name must be at least 2 characters long")
	}
	// Validate if the spot name is in the correct format
	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return errors.New("spot name must start with a letter")
	}
	if s.Name[1] < '0' || s.Name[1] > '9' {
		return errors.New("spot name must end with a number")
	}
	return nil
}
