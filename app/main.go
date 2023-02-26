package main

import (
	"net/http"
    "github.com/gorilla/mux"
	"github.com/Karaokifyy/Karaokifyy/app/spotify_api"
)

//Object to represent a spotify user and their unique/temp session
type SpotifyUserSession struct
{
	inital_oauth_code string
	inital_oauth_state string
}

func newSpotifySession(w http.ResponseWriter, r *http.Request){

}

func main() {
	//gorilla-mux router
	r := mux.NewRouter()

	//code to display front-end resources and files when requested
	client_fs := http.FileServer(http.Dir("client/"))
	r.Handle("/client/", http.StripPrefix("/client/", client_fs))

	//redirects root requests to client homepage
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/client/project1/src/app/app.component.html", http.StatusFound)
	})

	r.HandleFunc("/track/{trackName}", spotify_api.getTrackByName).Methods("GET")
	r.HandleFunc("/albums", spotify_api.getAlbums).Methods("GET")
	r.HandleFunc("/albums/{id}", spotify_api.getAlbumByID).Methods("GET")
	r.HandleFunc("/albums", spotify_api.postAlbums).Methods("POST")

	http.ListenAndServe(":80", r)
}
