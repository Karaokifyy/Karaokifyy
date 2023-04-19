package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Karaokifyy/Karaokifyy/BackEnd/src/api/helpers/envVars"
)

func Run() {
	// Load environment variables for client authentication
	GetEnvVars()

	// Gorilla-mux router
	r := mux.NewRouter()

	// Handlers for endpoints
	r.HandleFunc("/search/song/{songName}", spotify_api.SearchSongRouter).Methods("GET")
	r.HandleFunc("/search/album/{albumName}", spotify_api.SearchAlbumRouter).Methods("GET")
	r.HandleFunc("/search/artist/{artistName}", spotify_api.SearchArtistRouter).Methods("GET")
	r.HandleFunc("/search/playlist/{playlistName}", spotify_api.SearchPlaylistRouter).Methods("GET")
	r.HandleFunc("/user/playlist/{playlistID}", spotify_api.GetPlaylist).Methods("GET")
	r.HandleFunc("/karaokifyy/{songID}", karaokifyy_api.GetAudioLyricsMux).Methods("GET")

	r.HandleFunc("/newSpotifySession", newSpotifySession)
	http.ListenAndServe(":8080", r)
}