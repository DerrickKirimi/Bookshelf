package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)

	log.Println("Starting Server :8080")
	s := &http.Server{
		Addr:         ":8080", //Note port declaration
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed { //defensive probe. Note closed server exception
		log.Fatal("Server startup failed")
	}
}

func greet(w http.ResponseWriter, r *http.Request) { //request pointer significance
	fmt.Fprintf(w, "Welcome on the voyage. May you be intimately familiar with its beautiful depths. Surf the wave. Don't forget your scuba! ") //Takes responsewriter 
}
