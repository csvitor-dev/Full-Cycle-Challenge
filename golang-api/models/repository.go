package models

type EventRepository interface {
	ListEvents() ([]Event, error)
	FindEventByID(eventID string) (*Event, error)
	FindSpotsByEventID(eventID string) ([]*Spot, error)
	CreateTicket(ticket *Ticket) error
	ReserveSpot(spotID, ticketID string) error
}
