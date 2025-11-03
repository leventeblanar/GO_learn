package main

import (
	"fmt"
	"errors"
)

type Song struct {
	title		string
	artist		string
	duration	int
}

type MusicPlayer struct {
	playlist		[]Song
	currentIndex	int
	isPlaying		bool
}

// AddSong(song Song) - dal hozzáadása
func (mp *MusicPlayer) AddSong(song Song) {
	mp.playlist = append(mp.playlist, song)
	fmt.Printf("%s hozzáadva a lejátszási listádhoz\n", song.title)
}

// Play() error - lejátszás indítása (error ha üres a playlist)
func (mp *MusicPlayer) Play() error {
	if len(mp.playlist) == 0 {
		return errors.New("a lejátszási lista üres")
	} else {
		mp.isPlaying = true
		mp.currentIndex = 0
		return nil
	}
}

// Pause() - megállítás
func (mp *MusicPlayer) Pause() error {
	if !mp.isPlaying {
		return errors.New("a lejátszó nincs bekapcsolva")
	} else {
		mp.isPlaying = false
		return nil
	}
}

// Next() error - következő dal (error ha nincs több)
func (mp *MusicPlayer) Next() error {
	if mp.currentIndex >= len(mp.playlist)-1 {
		return errors.New("nincs több dal")
	} else {
		mp.currentIndex++
		return nil
	}
}

// Previous() error - előző dal (error ha az elsőnél van)
func (mp *MusicPlayer) Previous() error {
	if mp.currentIndex == 0 {
		return errors.New("ez az első dal")
	} else {
		mp.currentIndex--
		return nil
	}
}

// GetCurrentSong() (Song, error) - visszaadja az aktuális dalt (error ha üres)
func (mp MusicPlayer) GetCurrentSong() (Song, error) {
	if len(mp.playlist) == 0 {
		return Song{} , errors.New("a lejátszási lista üres")
	}
	song := mp.playlist[mp.currentIndex]
	return song, nil
}

// GetPlaylistInfo() string - playlist összefoglaló (pl. "5 dal, 15:32 perc")
func (mp MusicPlayer) GetPlaylistInfo() string {
	sumDalok := len(mp.playlist)
	var lenDalok int
	for _, song := range mp.playlist {
		lenDalok = lenDalok + song.duration
	}
	return fmt.Sprintf("Jelenleg %d db dal van a listában, ami %d mp hosszú összesen\n", sumDalok, lenDalok)
}

func main() {

	var MusicPlayer1 MusicPlayer

	MusicPlayer1.AddSong(Song{title: "Kolbász", artist: "Mackó", duration: 220})
	MusicPlayer1.AddSong(Song{title: "Nincsen tej", artist: "Prognózis", duration: 300})
	MusicPlayer1.AddSong(Song{title: "Highway to hell", artist: "AC/DC", duration: 420})

	if err := MusicPlayer1.Play(); err != nil {
		fmt.Println("Hiba: ", err)
	} else {
		fmt.Println("Lejátszás megkezdve!")
	}

	song, err := MusicPlayer1.GetCurrentSong()
	if err != nil {
		fmt.Println("Hiba:", err)
	} else {
		fmt.Printf("Most szól: %s - %s (%d mp)\n", song.artist, song.title, song.duration)
	}

	if err := MusicPlayer1.Next(); err != nil {
		fmt.Println("Hiba:", err)
	}

	song, err = MusicPlayer1.GetCurrentSong()
	if err != nil {
		fmt.Println("Hiba:", err)
	} else {
		fmt.Printf("Most szól: %s - %s (%d mp)\n", song.artist, song.title, song.duration)
	}

	fmt.Println(MusicPlayer1.GetPlaylistInfo())
}