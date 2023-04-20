package router

import (
	"net/http"

	"github.com/gorilla/mux"

	controller "github.com/Karaokifyy/BackEnd/src/controllers"
)

func Run() {
	//gorilla-mux router
	r := mux.NewRouter()
	r.HandleFunc("/search/song/{songName}", controller.SearchSongController).Methods("GET")
	r.HandleFunc("/search/album/{albumName}", controller.SearchAlbumController).Methods("GET")
	r.HandleFunc("/search/artist/{artistName}", controller.SearchArtistController).Methods("GET")
	r.HandleFunc("/search/playlist/{playlistName}", controller.SearchPlaylistController).Methods("GET")
	r.HandleFunc("/user/playlist/{playlistID}", controller.GetPlaylistController).Methods("GET")
	r.HandleFunc("/karaokifyy/{songID}", controller.AudioLyricsController).Methods("GET")
	r.HandleFunc("/newSpotifySession", controller.NewSpotifySession)
	http.ListenAndServe(":8080", r)
}
