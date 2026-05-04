package db

import (
	"database/sql"
	"time"

	"github.com/leventeblanar/GO_learn/weather-logger/api"
	_ "github.com/lib/pq"
)

func InitDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertReading(db *sql.DB, w api.CurrentWeather) error {

	_, err := db.Exec(`
		INSERT INTO readings (timestamp, temp, windspeed, weathercode)
		VALUES ($1, $2, $3, $4)`,
		time.Now().UTC(), w.Temperature, w.Windspeed, w.Weathercode,
	)

	return err
}

func GetReadings(db *sql.DB, limit int) ([]api.CurrentWeather, error) {

	rows, err := db.Query("SELECT temp, windspeed, weathercode FROM readings ORDER BY timestamp DESC LIMIT $1", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []api.CurrentWeather

	for rows.Next() {
		var w api.CurrentWeather
		err := rows.Scan(&w.Temperature, &w.Windspeed, &w.Weathercode)
		if err != nil {
			return nil, err
		}
		results = append(results, w)
	}

	return results, nil
}