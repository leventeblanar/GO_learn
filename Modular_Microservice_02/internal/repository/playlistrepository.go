package repository

import (
	"database/sql"
	"fmt"

	"modularmicserv/internal/model"
)

type PlaylistRepository struct {
	db 			*sql.DB
}

func NewPlaylistRepository(db *sql.DB) *PlaylistRepository {
	return &PlaylistRepository{db: db}
}

func (pr *PlaylistRepository) GetAllPlaylists() ([]model.Playlist, error) {

	query := `SELECT playlist_id, name FROM playlist`

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying playlists: %w", err)
	}
	defer rows.Close()

	var playlists []model.Playlist

	for rows.Next() {
		var playlist model.Playlist
		err := rows.Scan(&playlist.PlaylistId, &playlist.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning playlist: %w", err)
		}
		playlists = append(playlists, playlist)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return playlists, nil
}


func (pr *PlaylistRepository) GetPlaylistWithTracks(playlistId int) (*model.Playlist, error) {

	query := `SELECT playlist_id, name, 
	
	FROM playlist WHERE playlist_id = $1`

	rows, err := pr.db.Query(query, playlistId)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var playlist model.Playlist
	var tracks []model.Track
	playlistSet := false

	for rows.Next(){
		var track model.Track
		var playlistId int
		var playlistName string

		err := rows.Scan(&playlistId, &playlistName, &track.TrackId, &track.Milliseconds, &track.AlbumId)
		if err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		if !playlistSet {
			playlist.PlaylistId = playlistId
			playlist.Name = playlistName
			playlistSet = true
		}

		tracks = append(tracks, track)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	playlist.Tracks = tracks
	return &playlist, nil
}