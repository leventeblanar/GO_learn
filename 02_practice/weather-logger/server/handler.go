package server

import (
	"database/sql"
	"encoding/json"
	"net/http"

	dbpackage "github.com/leventeblanar/GO_learn/weather-logger/db"
)

func ReadingsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		readings, err := dbpackage.GetReadings(db, 20)
		if err != nil {
			http.Error(w, "Failed to get readings", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(readings)
	}
}