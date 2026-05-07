package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"

	"modularmicserv/internal/db"
	"modularmicserv/internal/model"
	"modularmicserv/internal/repository"
)

func main() {
	conn, err := db.ConnectDb()
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	defer conn.Close()

	fmt.Println("Database connected")
	
	// trackRepo := repository.NewAlbumRepository(conn)

	trackRepo := repository.NewTrackRepository(conn)

	track, err := getTrackFromUser()
	if err != nil {
		log.Fatalf("failed to create track: %v", err)
	}

	err = trackRepo.CreateTrackWithValidation(track)
	if err != nil {
		log.Fatalf("failed to create track: %v", err)
	}

	fmt.Printf("\nTrack '%s' successfully created with ID %d!\n", track.Name, track.TrackId)

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

	// err = trackRepo.UpdateTrackDuration(3504, 300000)
	// if err != nil {
	// 	log.Fatalf("failed to update track duration: %v", err)
	// }

	// fmt.Println("Track updated successfully!")

	// albumTracks, err := trackRepo.GetAlbumWithTracks(1)
	// if err != nil {
	// 	log.Fatalf("failed to fetch album: %v", err)
	// }

	// fmt.Printf("\nAlbum: %s, Tracks: %d\n", albumTracks.Title, len(albumTracks.Tracks))
	// for i, t := range albumTracks.Tracks {
	// 	fmt.Printf("%d. %s (%d ms)\n", i+1, t.Name, t.Milliseconds)
	// }
}	

func getTrackFromUser() (model.Track, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== New Track Creation ===")

	trackId, err := readInt(reader, "Track ID: ")
	if err != nil {
		return model.Track{}, err
	}

	name, err := readString(reader, "Track name: ")
	if err != nil {
		return model.Track{}, err
	}

	albumId, err := readInt(reader, "Album ID: ")
	if err != nil {
		return model.Track{}, err
	}

	milliseconds, err := readInt(reader, "Duration (ms): ")
	if err != nil {
		return model.Track{}, err
	}

	mediaTypeId, err := readInt(reader, "Media Type ID: ")
	if err != nil {
		return model.Track{}, err
	}

	unitPrice, err := readFloat(reader, "Unit Pricer: ")
	if err != nil {
		return model.Track{}, err
	}

	return model.Track {
		TrackId: trackId,
		Name:	name,
		AlbumId: albumId,
		Milliseconds: milliseconds,
		MediaTypeId: mediaTypeId,
		UnitPrice: unitPrice,
	}, nil
	
}


func readString(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}


func readInt(reader *bufio.Reader, prompt string) (int, error) {
	str, err := readString(reader, prompt)
	if err != nil {
		return 0, nil
	}
	return strconv.Atoi(str)
}


func readFloat(reader *bufio.Reader, prompt string) (float64, error) {
	str, err := readString(reader, prompt)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(str, 64)
}