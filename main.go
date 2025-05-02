package main

import (
	"go-lego-api/api"
	"log"
	"net/http"
	"os"

	_ "go-lego-api/docs" // Swag will generate this

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set up the mux router
	r := mux.NewRouter()

	// Handle the Swagger UI
	// r.Handle("/swagger/", httpSwagger.WrapHandler)

	// Serve the static files
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Route to get all sets
	r.HandleFunc("/api/allsets", api.GetAllLegoSets).Methods("GET")

	// Route to get a specific set by ID
	// redo this it might be the issue
	r.HandleFunc("/api/sets/{id:[0-9]+}", api.GetLegoSetByID).Methods("GET") // Ensure that ID matches only numbers
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)                  // Set the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
