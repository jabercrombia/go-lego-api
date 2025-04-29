package handler

import (
	"database/sql"
	"encoding/json"
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

// HandleRequest is the exported function that Vercel needs
func HandleRequest(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sets)
}

func init() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to Postgres
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
