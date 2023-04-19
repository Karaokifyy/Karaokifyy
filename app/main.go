package main

import (
	"log"
	"net/http"

	//"path/filepath"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	//	"github.com/zmb3/spotify"
	//	"golang.org/x/oauth2/clientcredentials"

	"github.com/Karaokifyy/Karaokifyy/app/karaokifyy_api"
	"github.com/Karaokifyy/Karaokifyy/app/spotify_api"
)

func main() {
	// load environment variables for client authentication
	getEnvVars()
	// gorilla-mux router
	r := mux.NewRouter()

	//code to display front-end resources and files when requested
	// client_fs := http.FileServer(http.Dir("client/"))
	// r.PathPrefix("/client/").Handler(http.StripPrefix("/client/",client_fs))

	//redirects root requests to client homepage
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "/client/Karaokify/dist/karaokify/index.html", http.StatusFound)
	// })

	r.HandleFunc("/search/song/{songName}", spotify_api.SearchSongRouter).Methods("GET")
	r.HandleFunc("/search/album/{albumName}", spotify_api.SearchAlbumRouter).Methods("GET")
	r.HandleFunc("/search/artist/{artistName}", spotify_api.SearchArtistRouter).Methods("GET")
	r.HandleFunc("/search/playlist/{playlistName}", spotify_api.SearchPlaylistRouter).Methods("GET")
	r.HandleFunc("/user/playlist/{playlistID}", spotify_api.GetPlaylist).Methods("GET")
	r.HandleFunc("/karaokifyy/{songID}", karaokifyy_api.GetAudioLyricsMux).Methods("GET")

	// r.HandleFunc("/track/{trackName}", spotify_api.GetTrackByName).Methods("GET")
	// r.HandleFunc("/albums", spotify_api.GetAlbums).Methods("GET")
	// r.HandleFunc("/albums/{id}", spotify_api.GetAlbumByID).Methods("GET")
	// r.HandleFunc("/albums", spotify_api.PostAlbums).Methods("POST")

	r.HandleFunc("/newSpotifySession", newSpotifySession)
	http.ListenAndServe(":8080", r)
}

var spotify_users map[string]spotify_api.SpotifyUserSession = make(map[string]spotify_api.SpotifyUserSession)

func newSpotifySession(w http.ResponseWriter, r *http.Request) {
	log.Printf("Called newSpotifySession\n")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	//retrieving initial json paramaters and setting redirect0rui
	var newUser spotify_api.SpotifyUserSession

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		return
	}
	newUser.Redirect_uri = "http://localhost:4200/screen-search"

	//requesting access_token
	log.Printf("requesting access_token\n")
	if err := spotify_api.RequestAccessToken(&newUser); err != nil {
		log.Fatalf("couldn't retrive access token") //convert to non-fatal or return
	}

	log.Printf("getting user playlist\n")
	playlists, err := spotify_api.GetUserPlaylists(&newUser)

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

// Load environment vars from file
func getEnvVars() {
	err := godotenv.Load("./config/cred.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
