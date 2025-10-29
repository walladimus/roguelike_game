package main

import (
	"log"
	"net/http"
	"os"

	"roguelike_game/backend/internal/httpserver"
)

func main()  {
	// read port or default to 8081
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	//Build the router
	handler := httpserver.NewRouter()

	//Log startup
	log.Printf("server starting on: %s", port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("server error: %v", err)
	}
}