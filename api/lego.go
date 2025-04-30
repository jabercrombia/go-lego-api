package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// LegoHandler responds to /api/lego requests
func LegoHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve DB connection URL from environment variable
	dbURL := os.Getenv("POSTGRES_URL")
	if dbURL == "" {
		http.Error(w, "POSTGRES_URL environment variable not set", http.StatusInternalServerError)
		return
	}

	// Connect to PostgreSQL DB
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to connect to database: %v", err), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Query database to fetch Lego sets
	rows, err := db.Query("SELECT set_id, name, year, id FROM lego_table LIMIT 10")
	if err != nil {
		http.Error(w, fmt.Sprintf("Query error: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Prepare response data
	var legoSets []map[string]interface{}
	for rows.Next() {
		var setID, name string
		var year, id int
		if err := rows.Scan(&setID, &name, &year, &id); err != nil {
			http.Error(w, fmt.Sprintf("Scan error: %v", err), http.StatusInternalServerError)
			return
		}

		legoSets = append(legoSets, map[string]interface{}{
			"set_id": setID,
			"name":   name,
			"year":   year,
			"id":     id,
		})
	}

	// Check for row iteration error
	if err := rows.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Row iteration error: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the data in JSON format
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(legoSets); err != nil {
		http.Error(w, fmt.Sprintf("JSON encoding error: %v", err), http.StatusInternalServerError)
	}
}
