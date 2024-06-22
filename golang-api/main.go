package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/csvitor-dev/Full-Cycle-Challenge/golang-api/models"
)

type DataOutPutDTO struct {
	Events []models.Event `json:"events"`
	Spots []models.Spot `json:"spots"`
}

var globalData DataOutPutDTO

func loadData() {
	data, err := os.Open("./data.json")
	if err != nil {
		log.Fatalf("Error [Open fail]: %s", err)
	}
	defer data.Close()

	decoder := json.NewDecoder(data)
	err = decoder.Decode(&globalData)
	if err != nil {
		log.Fatalf("Error [Decode fail]: %s", err)
	}
}

// HTTP handlers
// GET /events
func listEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(globalData.Events)
}
// GET /events/:eventId
func findEvent(w http.ResponseWriter, r *http.Request) {

}
// GET /events/:eventId/spots
func listSpots(w http.ResponseWriter, r *http.Request) {

}
// POST /event/:eventId/reserve
func reserveSpots(w http.ResponseWriter, r *http.Request) {

}

func main() {
	loadData()

	r := http.NewServeMux()

	r.HandleFunc("/events", listEvents)
	r.HandleFunc("/events/{eventId}", findEvent)
	r.HandleFunc("/events/{eventId}/spots", listSpots)
	r.HandleFunc("POST /event/{eventId}/reserve", reserveSpots)
	
	server := &http.Server{
		Addr: ":3000",
		Handler: r,
	}

	log.Println("HTTP server listen and serve in port 3000")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Error [Server error]: %v\n", err)
	}

	log.Println("Shutdown...")
}