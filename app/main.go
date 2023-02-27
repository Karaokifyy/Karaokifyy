package main

import (
	"net/http"
	//"path/filepath"
	//"encoding/json"

	"github.com/gorilla/mux"
	//	"github.com/zmb3/spotify"
	//	"golang.org/x/oauth2/clientcredentials"
	"github.com/Karaokifyy/Karaokifyy/app/spotify_api"
)

func main() {
	//gorilla-mux router
	r := mux.NewRouter()

	//code to display front-end resources and files when requested
	client_fs := http.FileServer(http.Dir("client/"))
	r.PathPrefix("/client/").Handler(http.StripPrefix("/client/",client_fs))

	//redirects root requests to client homepage
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/client/Karaokify/dist/karaokify/index.html", http.StatusFound)
	})

	r.HandleFunc("/main.4835f67e1d016940.js", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/client/Karaokify/dist/karaokify/main.4835f67e1d016940.js", http.StatusFound)
	})

	r.HandleFunc("/polyfills.bf8970584519db6b.js", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/client/Karaokify/dist/karaokify/polyfills.bf8970584519db6b.js", http.StatusFound)
	})

	r.HandleFunc("/runtime.8a584406e953a5e1.js", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/client/Karaokify/dist/karaokify/runtime.8a584406e953a5e1.js", http.StatusFound)
	})

	r.HandleFunc("/styles.f37111f28d0f589f.css", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/client/Karaokify/dist/karaokify/styles.f37111f28d0f589f.css", http.StatusFound)
	})

	r.HandleFunc("/track/{trackName}", spotify_api.GetTrackByName).Methods("GET")
	r.HandleFunc("/albums", spotify_api.GetAlbums).Methods("GET")
	r.HandleFunc("/albums/{id}", spotify_api.GetAlbumByID).Methods("GET")
	r.HandleFunc("/albums", spotify_api.PostAlbums).Methods("POST")

	http.ListenAndServe(":80", r)
}

/*var spotify_redirect_uri string = "/"
var spotify_users []spotify_api.SpotifyUserSession


func newSpotifySession(w http.ResponseWriter, r *http.Request){
	var newUser spotify_api.SpotifyUserSession;

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		return
	}

	//requesting access_token


	spotify_users = append(spotify_users, newUser)
}*/