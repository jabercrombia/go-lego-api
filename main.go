package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"go-lego-api/api"
)

func main() {
	// Load env vars
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to PostgreSQL
	dbURL := os.Getenv("POSTGRES_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	// Initialize routes
	http.HandleFunc("/api/legosets", api.GetLegoSetsHandler(db))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
