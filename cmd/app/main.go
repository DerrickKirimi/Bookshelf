package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/DerrickKirimi/Bookshelf/config"
	"github.com/DerrickKirimi/Bookshelf/server/router"	
)

func main() {
	appConf := config.AppConfig()

	appRouter := router.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Greet)

	address := fmt.Sprintf(":%d", appConf.Server.Port) //Note colon

	log.Printf("Starting Server %s\n", address)
	s := &http.Server{
		Addr:         	address, //port declaration
		Handler:      	appRouter,
		ReadTimeout:	appConf.Server.TimeoutRead,
		WriteTimeout: 	appConf.Server.TimeoutWrite,
		IdleTimeout:	appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed { //defensive probe. Note closed server exception
		log.Fatal("Server startup failed")
	}
}

func Greet(w http.ResponseWriter, r *http.Request) { //request pointer significance
	fmt.Fprintf(w, "Welcome on the voyage. May you be intimately familiar with its beautiful depths. Surf the wave. Don't forget your scuba! ") //Takes responsewriter
}
