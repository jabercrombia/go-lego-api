package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// GetLegoSetByID godoc
// @Summary Get a LEGO set by ID
// @Description Retrieves a LEGO set by its ID
// @Tags lego
// @Produce json
// @Param id path int true "LEGO Set ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/sets/{id} [get]
func GetLegoSetByID(w http.ResponseWriter, r *http.Request) {
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

	// Get the set ID from the URL parameters
	id := r.URL.Path[len("/api/sets/"):]

	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Query database to fetch Lego sets
	rows, err := db.Query("SELECT set_id, name, year, theme, subtheme, themegroup, category, pieces, minifigs, agerange_min, us_retailprice, brickseturl, thumbnailurl, imageurl, id FROM lego_table WHERE id = $1", id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Query error: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Prepare response data
	var legoSets []map[string]interface{}
	for rows.Next() {
		var setID, name, theme, category, brickseturl string
		var year, id int

		// Use sql.NullString for nullable fields
		var nullableSubtheme sql.NullString
		var nullableThemegroup sql.NullString
		var nullableMinifigs sql.NullInt32
		var nullablePieces sql.NullInt32
		var nullableAgerangeMin sql.NullInt32
		var nullableUsRetailPrice sql.NullFloat64
		var nullableThumbnailURL sql.NullString
		var nullableImageURL sql.NullString

		// Scan the values into variables
		if err := rows.Scan(&setID, &name, &year, &theme, &nullableSubtheme, &nullableThemegroup, &category, &nullablePieces, &nullableMinifigs, &nullableAgerangeMin, &nullableUsRetailPrice, &brickseturl, &nullableThumbnailURL, &nullableImageURL, &id); err != nil {
			http.Error(w, fmt.Sprintf("Scan error: %v", err), http.StatusInternalServerError)
			return
		}

		// Append each lego set to the response slice
		legoSets = append(legoSets, map[string]interface{}{
			"set_id":         setID,
			"name":           name,
			"year":           year,
			"theme":          theme,
			"subtheme":       getNullableString(nullableSubtheme),
			"themegroup":     getNullableString(nullableThemegroup),
			"category":       category,
			"pieces":         getNullableInt(nullablePieces),
			"minifigs":       getNullableInt(nullableMinifigs),
			"agerange_min":   getNullableInt(nullableAgerangeMin),
			"us_retailprice": getNullableFloat(nullableUsRetailPrice),
			"brickseturl":    brickseturl,
			"thumbnailurl":   getNullableString(nullableThumbnailURL),
			"imageurl":       getNullableString(nullableImageURL),
			"id":             id,
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
