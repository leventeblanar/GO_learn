package repository

import (
	"database/sql"
	"fmt"

	"modularmicserv/internal/model"
)

type TrackRepository struct {
	db 			*sql.DB
}

func NewTrackRepository(db *sql.DB) *TrackRepository {
	return &TrackRepository{db: db}
}

func (t *TrackRepository) GetAllTracks() ([]model.Track, error) {

	query := `
	SELECT track_id, name, album_id, milliseconds
	FROM track
	`

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var tracks []model.Track

	for rows.Next() {
		var item model.Track
		err := rows.Scan(&item.TrackId, &item.Name, &item.Milliseconds, &item.AlbumId)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}
		tracks = append(tracks, item)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rws error: %w", err)
	}

	return tracks, nil
}

func (t *TrackRepository) GetAlbumWithTracks(albumId int) (*model.Album, error) {
	query := `
		SELECT 
			a.album_id, a.title,
			t.track_id, t.name, t.milliseconds, t.album_id
		FROM album a
		LEFT JOIN track t ON a.album_id = t.album_id
		WHERE a.album_id = $1
	`

	rows, err := t.db.Query(query, albumId)
	if err != nil {
		return nil, fmt.Errorf("error querying rows: %w", err)
	}
	defer rows.Close()

	var album model.Album
	var tracks []model.Track
	albumSet := false

	for rows.Next() {
		var track model.Track
		var albumId int
		var albumTitle string

		err := rows.Scan(
			&albumId, &albumTitle,
			&track.TrackId, &track.Name, &track.Milliseconds, &track.AlbumId,
		)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		if !albumSet {
			album.AlbumId = albumId
			album.Title = albumTitle
			albumSet = true
		}

		tracks = append(tracks, track)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	album.Tracks = tracks
	return &album, nil
}

func (t *TrackRepository) CreateTrack(track model.Track) error {

	query := `
	INSERT INTO track(track_id, name, album_id, milliseconds, media_type_id, unit_price)
	VALUES($1, $2, $3, $4, $5, $6)
	`

	result, err := t.db.Exec(
		query,
		track.TrackId,
		track.Name,
		track.AlbumId,
		track.Milliseconds,
		track.MediaTypeId,
		track.UnitPrice,
	)
	if err != nil {
		return fmt.Errorf("failed to insert track: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were inserted")
	}

	return nil
}

func (t *TrackRepository) UpdateTrackDuration(trackId int, newDuration int) error {

	query := `
	update track t
	set milliseconds = $1
	where t.track_id = $2
	`

	updateResult, err := t.db.Exec(
		query, newDuration, trackId,
	)
	if err != nil {
		return fmt.Errorf("failed to update track: %d", trackId)
	}

	rowsAffected, err := updateResult.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return nil
}


func (t *TrackRepository)DeleteTrack(trackId int) error {

	var trackExists bool

	checkTrackquery := `
	SELECT EXISTS (
    SELECT 1
    FROM track t
    WHERE t.track_id = $1
	);
	`
	

	checkPlaylisttrack := `
	SELECT EXISTS (
    SELECT 1
    FROM playlist_track t
    WHERE t.track_id = $1
	);
	`

	deleteQuery := `
	delete from track t
	where t.track_id = $1;
	`

	err := t.db.QueryRow(checkTrackquery, trackId).Scan(&trackExists)
	if err != nil {
		return fmt.Errorf("error checking track existence: %w", err)
	}

	if !trackExists {
		return fmt.Errorf("track %d not found", trackId)
	}

	var inPlaylist bool

	err = t.db.QueryRow(checkPlaylisttrack, trackId).Scan(&inPlaylist)
	if err != nil {
		return fmt.Errorf("error checking track existance: %w", err)
	}

	if inPlaylist {
		return fmt.Errorf("cannot delete: track %d is in playlists", trackId)
	}


	result, err := t.db.Exec(deleteQuery, trackId)
	if err != nil {
		return fmt.Errorf("error deleting track: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows deleted")
	}

	return nil
}


func (t *TrackRepository) DeleteTrackForce(trackId int) error {

	var trackExists bool
	checkQuery := `SELECT EXISTS (SELECT 1 FROM track WHERE track_id = $1)`

	err := t.db.QueryRow(checkQuery, trackId).Scan(&trackExists)
	if err != nil {
		return fmt.Errorf("error checking track: %w", err)
	}

	if !trackExists {
		return fmt.Errorf("track %d not found", trackId)
	}

	//tranzakció

	tx, err := t.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to beging transaction: %w", err)
	}

	// bármi hiba - rollback
	defer tx.Rollback()

	// törlés a playlist táblából - FONTOS: t.Exec() nem t.db.Exec()
	_, err = tx.Exec(`DELETE FROM playlist_track WHERE track_id = $1`, trackId)
	if err != nil {
		return fmt.Errorf("failed to remove track from playlist: %w", err)
	}

	// törlés a track táblából
	result, err := tx.Exec("DELETE FROM track WHERE track_id = $1", trackId)
	if err != nil {
		return fmt.Errorf("failed to delete track: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no track was deleted")
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}