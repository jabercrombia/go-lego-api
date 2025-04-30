package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type LegoSet struct {
	SetID  string `json:"set_id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	ID     int    `json:"id"`
	Pieces int    `json:"pieces"`
}

// This is the function Vercel looks for
func Handler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		http.Error(w, "Failed to connect to DB", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT set_id, name, year, id, pieces FROM lego_table LIMIT 10")
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sets []LegoSet
	for rows.Next() {
		var set LegoSet
		if err := rows.Scan(&set.SetID, &set.Name, &set.Year, &set.ID, &set.Pieces); err != nil {
			http.Error(w, "Scan error", http.StatusInternalServerError)
			return
		}
		sets = append(sets, set)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sets)
}
