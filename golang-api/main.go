package main

import (
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	// add handlers
	server := &http.Server{
		Addr: ":3000",
		Handler: r,
	}

	log.Println("HTTP server listen and serve in port 3000")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Error: %v\n", err)
	}

	log.Println("Shutdown...")
}