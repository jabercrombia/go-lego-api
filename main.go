package main

import (
	"fmt"
	"go-lego-api/api"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	_ "go-lego-api/docs" // Swag will generate this

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	frontendURL := os.Getenv("FRONT_END_URL")
	if frontendURL == "" {
		fmt.Println("Please set the FRONT_END_URL environment variable")
		return
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
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{frontendURL},    // Allow requests from your frontend
		AllowedMethods: []string{"GET"},          // Allowed methods
		AllowedHeaders: []string{"Content-Type"}, // Allowed headers
	}).Handler(r)

	// Start the server with CORS applied
	port := "8080" // Or use an environment variable for the port
	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
