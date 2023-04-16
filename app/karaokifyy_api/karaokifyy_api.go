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
<<<<<<< HEAD
	// songURL := getYoutubeLink(songName)
	songURL, err := youtube_api.GetYoutubeURL(songName)
	if err != nil {
		log.Printf("Could not get yotubeURL")
=======
	songLyrics, songDuration := getLRC(songID)

	songURL, err := youtube_api.GetYoutubeURL(songName, songDuration)
	if err != nil{
		log.Printf("Could not get youtubeURL")
>>>>>>> main
	}
	return songURL, songLyrics
}

//TODO: change to getYoutubeQuery; should return song name with artist name appended to it 
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
<<<<<<< HEAD
	if db.First(&track, "spotify_id = ?", songID).Error != nil {
		if songID == "5ihS6UUlyQAfmp48eSkxuQ" {
			return `[00:12.62]I took my love, I took it down
	[00:18.40]Climbed a mountain and I turned around
	[00:23.11]And I saw my reflection in the snow covered hills
	[00:29.47]'Til the landslide brought it down

	[00:35.98]Oh, mirror in the sky, What is love?
	[00:41.85]Can the child within my heart rise above?
	[00:48.01]Can I sail thru the changin' ocean tides?
	[00:53.80]Can I handle the seasons of my life?

	[01:11.77]Well, I've been afraid of changin'
	[01:16.71]'Cause I've built my life around you
	[01:23.39]But time makes you bolder
	[01:26.65]Even children get older
	[01:30.02]And I'm getting older too
	[02:00.42]Well, I've been afraid of changin'
	[02:05.67]'Cause I've built my life around you
	[02:12.48]But time makes you bolder
	[02:15.52]Even children get older
	[02:18.70]And I'm getting older too
	[02:24.88]And I'm getting older too

	[02:31.02]Ahh, take my love, take it down
	[02:38.34]Ahh, Climb a mountain and turn around
	[02:43.03]And if you see my reflection in the snow covered hills
	[02:49.30]Well the landslide will bring it down
	[02:54.85]And if you see my reflection in the snow covered hills
	[03:02.73]Well the landslide will bring it down

	[03:09.14]Oh, The landslide will bring it down`
		} else {
			return "nil"
		}
=======
	if db.First(&track, "spotify_id = ?", songID).Error != nil{
		return "",0
>>>>>>> main
	}

	return track.SyncedLyrics, track.Duration
}
