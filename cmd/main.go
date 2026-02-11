package main

import (
	"log"
	"net/http"
)

func main() {
	Mux := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: Mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server is broken with error: %v", err)
	}
	log.Println("Server is starting on port 8080")
}
