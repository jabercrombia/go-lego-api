package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type LegoSet struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Theme        string  `json:"theme"`
	ThumbnailUrl *string `json:"thumbnailurl"`
}

var db *sql.DB

func init() {
	var err error
	dbURL := os.Getenv("POSTGRES_URL")
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
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
