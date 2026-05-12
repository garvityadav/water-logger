package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// loading dotenv file
	err := godotenv.Load()
	// checking for errors
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// setting up port
	port := os.Getenv("PORT")

	// creating http server
	mux := http.NewServeMux()
	// /health
	mux.HandleFunc("/health", healthHandler)

	// server
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	server.ListenAndServe()
	// dev printing port
	fmt.Println("starting server on port: ", port)
}
