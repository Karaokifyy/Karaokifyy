package controller

import (
	"encoding/json"
	"log"
	"net/http"

	model "github.com/Karaokifyy/BackEnd/src/models"
	types "github.com/Karaokifyy/BackEnd/src/models/types"
	spotify_service "github.com/Karaokifyy/BackEnd/src/services/spotify"
	sql_service "github.com/Karaokifyy/BackEnd/src/services/sql"
	youtube_service "github.com/Karaokifyy/BackEnd/src/services/youtube"
	view "github.com/Karaokifyy/BackEnd/src/views"

	"github.com/gorilla/mux"
)

// Handle song search http request, send data to service, pass result to generate model, send to view encoder.
func SearchSongController(w http.ResponseWriter, r *http.Request) {
	songName := mux.Vars(r)["songName"]
	searchResult := spotify_service.SearchSongService(songName)
	songList := model.SongListModel(searchResult)
	view.EncodeSongList(w, songList)
}

// Handle album search http request, send data to service, pass result to generate model, send to view encoder.
func SearchAlbumController(w http.ResponseWriter, r *http.Request) {
	albumName := mux.Vars(r)["albumName"]
	searchResult := spotify_service.SearchAlbumService(albumName)
	albumList := model.AlbumListModel(searchResult)
	view.EncodeSongList(w, albumList)
}

// Handle artist search http request, send data to service, pass result to generate model, send to view encoder.
func SearchArtistController(w http.ResponseWriter, r *http.Request) {
	artistName := mux.Vars(r)["artistName"]
	searchResult := spotify_service.SearchArtistService(artistName)
	artistList := model.ArtistListModel(searchResult)
	view.EncodeSongList(w, artistList)
}

// Handle playlist search http request, send data to service, pass result to generate model, send to view encoder.
func SearchPlaylistController(w http.ResponseWriter, r *http.Request) {
	playlistName := mux.Vars(r)["playlistName"]
	searchResult := spotify_service.SearchPlaylistService(playlistName)
	playlistList := model.PlaylistListModel(searchResult)
	view.EncodeSongList(w, playlistList)
}

// Handle playlist id search http request, send data to service, pass result to generate model, send to view encoder.
func GetPlaylistController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	playlistName := mux.Vars(r)["playlistID"]
	searchResult := spotify_service.GetPlaylistService(playlistName)
	playlistList := model.UserPlaylistModel(searchResult)
	view.EncodeSongList(w, playlistList)
}

// Handle karaokifyy search http request, send data to service, pass result to generate model, send to view encoder.
func AudioLyricsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	songID := mux.Vars(r)["songID"]
	songName := spotify_service.GetSongName(songID)

	songLyrics, songDuration := sql_service.GetLRC(songID)

	songURL, err := youtube_service.GetYoutubeURL(songName, songDuration)
	if err != nil {
		log.Printf("Could not get youtubeURL")
	}

	audioLyrics := model.AudioLyricsModel(songURL, songLyrics)
	view.EncodeSongList(w, audioLyrics)
}

var spotify_users map[string]types.SpotifyUserSession = make(map[string]types.SpotifyUserSession)

// generate new spotify session
func NewSpotifySession(w http.ResponseWriter, r *http.Request) {
	log.Printf("Called newSpotifySession\n")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	//retrieving initial json paramaters and setting redirect0rui
	var newUser types.SpotifyUserSession
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		return
	}
	newUser.Redirect_uri = "http://localhost:4200/screen-search"
	//requesting access_token
	log.Printf("requesting access_token\n")
	if err := spotify_service.RequestAccessToken(&newUser); err != nil {
		log.Fatalf("couldn't retrive access token") //convert to non-fatal or return
	}

	log.Printf("getting user playlist\n")
	playlists, err := spotify_service.GetUserPlaylists(&newUser)

	if err != nil {
		log.Fatalf("Couldn't get user's playlist")
	}

	json.NewEncoder(log.Default().Writer()).Encode(playlists)
	json.NewEncoder(w).Encode(playlists)
	// for _, playlist := range playlists {
	// 	json.NewEncoder(log.Default().Writer()).Encode(playlist)
	// 	json.NewEncoder(w).Encode(playlist)
	// }

	spotify_users[r.RemoteAddr] = newUser
}
