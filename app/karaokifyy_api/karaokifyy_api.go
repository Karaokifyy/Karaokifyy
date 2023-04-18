package karaokifyy_api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Karaokifyy/Karaokifyy/app/spotify_api"
	"github.com/Karaokifyy/Karaokifyy/app/youtube_api"

	"github.com/gorilla/mux"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AudioLyrics struct {
	SongURL    string `json:"SongURL"`
	SongLyrics string `json:"SongLyrics"`
}

func GetAudioLyricsMux(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	vars := mux.Vars(r)
	songID := vars["songID"]
	audio, lyrics := GetAudioLyrics(songID)
	audioLyrics := AudioLyrics{audio, lyrics}
	json.NewEncoder(w).Encode(audioLyrics)
}

func GetAudioLyrics(songID string) (string, string) {
	songName := getSongName(songID)
	songLyrics, songDuration := getLRC(songID)

	songURL, err := youtube_api.GetYoutubeURL(songName, songDuration)
	if err != nil {
		log.Printf("Could not get youtubeURL")
	}
	return songURL, songLyrics
}

// TODO: change to getYoutubeQuery; should return song name with artist name appended to it
func getSongName(songID string) string {
	return spotify_api.GetSongName(songID)
}

func getYoutubeLink(songName string) string {
	if songName == "Landslide" {
		return "https://www.youtube.com/watch?v=radHy4HhhNg&ab_channel=FleetwoodMac-Topic"
	} else {
		return "nil"
	}
}

type Track struct {
	gorm.Model
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ArtistName   string `json:"artistName"`
	AlbumName    string `json:"albumName"`
	Duration     int    `json:"duration"`
	Instrumental bool   `json:"instrumental"`
	Lang         string `json:"lang"`
	ISRC         string `json:"isrc"`
	SpotifyId    string `json:"spotifyId"`
	ReleaseDate  string `json:"releaseDate"`
	PlainLyrics  string `json:"plainLyrics"`
	SyncedLyrics string `json:"syncedLyrics"`
}

func getLRC(songID string) (string, int) {

	db, err := gorm.Open(sqlite.Open("assets/lyrics_lite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Track{})

	var track Track
	// db.First(&track, "spotify_id = ?", songID)
	if db.First(&track, "spotify_id = ?", songID).Error != nil {
		return "", 0
	}

	return track.SyncedLyrics, track.Duration
}
