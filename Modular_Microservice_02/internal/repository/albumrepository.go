package repository

import (
	"fmt"
	"database/sql"

	"modularmicserv/internal/model"
)

type AlbumRepository struct {
	db			*sql.DB 
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (ar *AlbumRepository)GetAllAlbumStat() ([]model.AlbumStats, error) {

	query := `
	select
	a.album_id,
	a.title,
	Count(a.title),
	SUM(t.milliseconds),
	AVG(t.milliseconds)
	from album a
	left join track t on a.album_id  = t.album_id
	group by a.album_id
	order by count(a.title) ASC
	`

	rows, err := ar.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error pulling rows: %w", err)
	}
	defer rows.Close()

	var albumStats []model.AlbumStats

	for rows.Next() {
	var albumStat model.AlbumStats
		err := rows.Scan(&albumStat.AlbumId, &albumStat.Title, &albumStat.TrackCount, &albumStat.TotalDuration, &albumStat.AverageDuration)
		if err != nil {
		return nil, fmt.Errorf("error scanning rows: %w", err)
		}
		albumStats = append(albumStats, albumStat)
	}

	if err := rows.Err();err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return albumStats, nil
}