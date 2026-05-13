package main

import (
	"log"
	"net/http"
	"os"

	"github.com/garvityadav/water-tracker/internal/database"
	"github.com/garvityadav/water-tracker/internal/handlers"
	"github.com/joho/godotenv"
)

func main() {
	// loading dotenv file
	err := godotenv.Load()
	// checking for errors
	if err != nil {
		log.Fatal("main: Error loading .env file")
	}
	// setting up port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// setup db
	db, dbErr := database.Connect(os.Getenv("DB"))
	if dbErr != nil {
		log.Fatalf("database.Connect:%v", dbErr)
	}

	// closing database
	db.Close()
	// creating http server
	mux := http.NewServeMux()

	// Endpoints starts here --------

	// /health
	mux.HandleFunc("GET /health", handlers.HealthHandler)

	// GET history /water/history
	mux.HandleFunc("GET /water/history", handlers.HistoryHandler)

	//
	// Endpoints ends here ----------
	// server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	// dev printing port
	log.Printf("main: starting server on port: %s ", port)

	log.Fatal(server.ListenAndServe())
}
