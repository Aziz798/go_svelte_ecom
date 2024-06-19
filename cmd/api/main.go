package main

import (
	"go_ecom/internal/server"
	"log"
)

func main() {
	// Initialize the server
	srv := server.NewServer()

	// Start the server and handle any errors
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
