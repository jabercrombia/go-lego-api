package main

import (
	"go-lego-api/api"
	"log"
	"net/http"
	"os"

	_ "go-lego-api/docs" // Swag will generate this

	"github.com/gorilla/mux"
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
	r.HandleFunc("/api/sets", api.GetAllLegoSets).Methods("GET")

	// Route to get a specific set by ID
	r.HandleFunc("/api/sets/{id:[0-9]+}", api.GetLegoSetByID).Methods("GET") // Ensure that ID matches only numbers

	// Serve Swagger UI
	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swagger-ui"))))

	// Set the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
