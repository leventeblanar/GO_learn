package repository

import (
	"fmt"
	"database/sql"
	"strings"

	"modularmicserv/internal/model"
)

type StatisticsRepository struct {
	db			*sql.DB
}

func NewStatisticsRepository(db *sql.DB) *StatisticsRepository {
	return &StatisticsRepository{db: db}
}

func (sr *StatisticsRepository)GetAlbumStatistics() ([]model.AlbumStatistics, error) {

	query := `
	select 
	a.artist_id, 
	a.name as artist_name, 
	al.title as album_title, 
	count(t.track_id) as track_count,
	round(sum(t.milliseconds)/60000, 2) as total_minutes,
	round(avg(t.milliseconds)/60000, 2) as average_minutes 
	from artist a 
	left join album al on a.artist_id = al.artist_id 
	left join track t on al.album_id = t.album_id 
	group by a.artist_id, a.name, al.title
	order by a.artist_id asc;
	`

	rows, err := sr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var albumStatistics []model.AlbumStatistics

	for rows.Next() {
		var AlbumStatistic model.AlbumStatistics
		err := rows.Scan(&AlbumStatistic.ID, &AlbumStatistic.ArtistName, &AlbumStatistic.AlbumTitle, &AlbumStatistic.TrackCount, &AlbumStatistic.TotalMinutes, &AlbumStatistic.AverageMinutes)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		albumStatistics = append(albumStatistics, AlbumStatistic)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return albumStatistics, nil

}



func (sr *StatisticsRepository) InsertStatistics(stats []model.AlbumStatistics) error {
	
	if len(stats) == 0 {
		return nil
	}

	query := `
	INSERT INTO public.album_statistics (artist_name, album_title, track_count, total_minutes, average_minutes) VALUES
	`

	values := []interface{}{}
	placeholders := []string{}

	for i, stat := range stats {
		base := i * 5
		// Minden stat-hoz 5 érték tartozik (artist_name, album_title, track_count, total_minutes, average_minutes).
		// stat[0] → $1,  $2,  $3,  $4,  $5   (base = 0*5 = 0,  tehát 0+1, 0+2, 0+3...)
		// stat[1] → $6,  $7,  $8,  $9,  $10  (base = 1*5 = 5,  tehát 5+1, 5+2, 5+3...)
		// stat[2] → $11, $12, $13, $14, $15  (base = 2*5 = 10, tehát 10+1, 10+2...)
		placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", base+1, base+2, base+3, base+4, base+5),)
		values = append(values, stat.ArtistName, stat.AlbumTitle, stat.TrackCount, stat.TotalMinutes, stat.AverageMinutes)
	}

	query += strings.Join(placeholders, ", ")

	_, err := sr.db.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("batch insert failed: %w", err)
	}

	return nil
}