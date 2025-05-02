package main

import (
	"go-lego-api/api"
	"log"
	"net/http"
	"os"

	_ "go-lego-api/docs" // Swag will generate this

	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Load .env from root directory
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	frontendURL := os.Getenv("FRONT_END_URL")
	if frontendURL == "" {
		log.Fatal("Please set the FRONT_END_URL environment variable")
	}

	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/api/allsets", api.GetAllLegoSets).Methods("GET")
	r.HandleFunc("/api/sets/{id:[0-9]+}", api.GetLegoSetByID).Methods("GET")
	r.HandleFunc("/api/health", api.HealthCheck).Methods("GET")

	// Swagger docs
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	// CORS configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{frontendURL},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(r)

	port := "8080"
	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
