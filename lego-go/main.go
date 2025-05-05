package main

import (
	"fmt"
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
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	if os.Getenv("FRONT_END_URL") == "" {
		log.Fatal("Please set the FRONT_END_URL environment variable")
	}

	log.Println("Error loading .env file")
	fmt.Print("Hello, ")

	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/api/allsets", api.GetAllLegoSets).Methods("GET")
	r.HandleFunc("/api/sets/{id:[0-9]+}", api.GetLegoSetByID).Methods("GET")

	// Swagger docs
	r.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler)

	// CORS configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("FRONT_END_URL")},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(r)

	port := "8080"
	log.Printf("Server is running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
