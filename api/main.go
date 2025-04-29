package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// LegoSet represents a sample row
type LegoSet struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Theme        string  `json:"theme"`
	ThumbnailUrl *string `json:"thumbnailurl"`
}

var db *sql.DB

// HandleRequest is the function that serves the data from the database
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Check if the path is correct
	if r.URL.Path != "/api/legosets" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Query to get Lego sets from the table
	rows, err := db.Query("SELECT id, name, theme, thumbnailurl FROM lego_table LIMIT 10")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sets []LegoSet
	for rows.Next() {
		var set LegoSet
		if err := rows.Scan(&set.ID, &set.Name, &set.Theme, &set.ThumbnailUrl); err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		sets = append(sets, set)
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sets)
}

func init() {
	// Load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	dbURL := os.Getenv("POSTGRES_URL")
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}
}

func main() {
	// Set up the route and start the server
	http.HandleFunc("/api/legosets", HandleRequest)
	port := "8080"
	fmt.Println("Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
