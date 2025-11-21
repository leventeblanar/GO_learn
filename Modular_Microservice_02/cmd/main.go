package main

import (
	"fmt"
	"log"

	"modularmicserv/internal/db"
	"modularmicserv/internal/repository"
)

func main() {
	conn, err := db.ConnectDb()
	if err != nil {
		log.Fatalf("cannot connect to database:", err)
	}
	defer conn.Close()

	fmt.Println("Database connected")
	
	// trackRepo := repository.NewAlbumRepository(conn)

	trackRepo := repository.NewTrackRepository(conn)

	// tracks, err := trackRepo.GetAllTracks()
	// if err != nil {
	// 	log.Fatalf("failed to get all tracks: %w", err)
	// }

	// fmt.Printf("\n === Found %d cikk === \n", len(tracks))
	// for i, t := range tracks {
	// 	fmt.Printf("%d. ID: %d, Név: %s, Hossz (ms): %d, AlbumID: %d\n", i+1, t.TrackId, t.Name, t.Milliseconds, t.AlbumId)
	// }

	// albumTracks, err := trackRepo.GetAlbumWithTracks(4)
	// if err != nil {
	// 	log.Fatalf("failed fetch all tracks: %w", err)
	// }

	// if albumTracks == nil {
	// 	fmt.Println("\n Album not found")
	// } else {
	// 	fmt.Printf("\n Album címe: %s, Albumon található dalok száma: %d\n", albumTracks.Title,len(albumTracks.Tracks))
	// }

	// for i, t := range albumTracks.Tracks {
	// 	fmt.Printf("%d. ID: %d, Számcím: %s, Hossz: %d, AlbumId: %d\n", i+1, t.TrackId, t.Name, t.Milliseconds, t.AlbumId)
	// }

	// stats, err := trackRepo.GetAllAlbumStat()
	// if err != nil {
	// 	log.Fatalf("failed to get album stats", err)
	// }

	// fmt.Printf("\n Number of stat rows: %d\n", len(stats))
	// for i, s := range stats {
	// 	fmt.Printf("%d. Album_ID: %d, Album title: %s, Track count: %d, Average duration: %2f, Total Duration: %2f\n", i+1, s.AlbumId, s.Title, s.TrackCount, s.AverageDuration, s.TotalDuration)
	// }

	// newTrack := model.Track {
	// 	TrackId: 3504,
	// 	Name:		"My test song",
	// 	AlbumId:	1,
	// 	Milliseconds: 240000,
	// 	MediaTypeId: 2,
	// 	UnitPrice: 0.99,
	// }

	// err = trackRepo.CreateTrack(newTrack)
	// if err != nil {
	// 	log.Fatalf("failed to create track: %v", err)
	// }

	// fmt.Println("Track successfully created!")

	// albumTracks, err := trackRepo.GetAlbumWithTracks(1)
	// if err != nil {
	// 	log.Fatalf("failed to fetch album: %v", err)
	// }

	// fmt.Printf("\nAlbum: %s, Tracks: %d\n", albumTracks.Title, len(albumTracks.Tracks))
	// for i, t := range albumTracks.Tracks {
	// 	fmt.Printf("%d. %s (%d ms)\n", i+1, t.Name, t.Milliseconds)
	// }

	err = trackRepo.UpdateTrackDuration(3504, 300000)
	if err != nil {
		log.Fatalf("failed to update track duration: %v", err)
	}

	fmt.Println("Track updated successfully!")

	albumTracks, err := trackRepo.GetAlbumWithTracks(1)
	if err != nil {
		log.Fatalf("failed to fetch album: %v", err)
	}

	fmt.Printf("\nAlbum: %s, Tracks: %d\n", albumTracks.Title, len(albumTracks.Tracks))
	for i, t := range albumTracks.Tracks {
		fmt.Printf("%d. %s (%d ms)\n", i+1, t.Name, t.Milliseconds)
	}
}	

