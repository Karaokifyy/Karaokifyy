package karaokifyy_api

import (
	"encoding/json"
	"net/http"

	"github.com/Karaokifyy/Karaokifyy/app/spotify_api"

	"github.com/gorilla/mux"
)

// Struct for song audio via youtube link and LRC file for song lyrics
type AudioLyrics struct {
	SongURL    string `json:"SongURL"`
	SongLyrics string `json:"SongLyrics"`
}

// Gets audio link and song lyrics function for gorilla mux
func GetAudioLyricsMux(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songID := vars["songID"]
	audio, lyrics := GetAudioLyrics(songID)
	audioLyrics := AudioLyrics{audio, lyrics}
	json.NewEncoder(w).Encode(audioLyrics)
}

// Gets audio link and song lyrics from a Spotify song ID
func GetAudioLyrics(songID string) (string, string) {
	songName := getSongName(songID)
	songURL := getYoutubeLink(songName)
	songLyrics := getLRC(songID)
	return songURL, songLyrics
}

// Gets a song's name from its song ID
func getSongName(songID string) string {
	return spotify_api.GetSongName(songID)
}

// Gets audio for a song via youtube using song name
func getYoutubeLink(songName string) string {
	if songName == "Landslide" {
		return "https://www.youtube.com/watch?v=radHy4HhhNg&ab_channel=FleetwoodMac-Topic"
	} else {
		return "nil"
	}
}

// Gets LRC lyrics for a song using its unique Spotify song ID
func getLRC(songID string) string {
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

}
