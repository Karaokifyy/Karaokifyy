package spotify_api

import (
	"context"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// Load environment vars from file
func getEnvVars() {
	err := godotenv.Load("./config/cred.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Create Albums struct to represent list of albums
type Albums struct {
	Albums []Album `json:"albums"`
}

// Create Album struct to represent data about a record album
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Load file and store albums data
var byteValue, _ = os.ReadFile("./server/playlist.json")
var albums Albums
var _ = json.Unmarshal(byteValue, &albums)


// getAlbums responds with the list of all albums as JSON
func getAlbums(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(albums)
}

// getAlbumByID locates the album whose ID value matches the id
// Parameter sent by the client, then returns that album as a response
func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter
	for _, a := range albums.Albums {
		if a.ID == id {
			json.NewEncoder(w).Encode(a)
			return
		}
	}
	
	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Album not found"})
	http.Error(w,"", http.StatusNotFound)
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(w http.ResponseWriter, r *http.Request) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum
	if err := json.NewDecoder(r.Body).Decode(&newAlbum); err != nil {
		return
	}

	// Add the new album to the list
	albums.Albums = append(albums.Albums, newAlbum)
	json.NewEncoder(w).Encode(newAlbum)
}

// getTrackByName searches for a track by name using spotify_api.SearchTrackByName method and
// serves resulting json data.
func getTrackByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackName := vars["trackName"]
	trackList := SearchByTrack(trackName)

	// Loop over the list of tracks, looking for a track
	// name that matches the parameter given.

	for _, a := range trackList.Tracks {
		json.NewEncoder(w).Encode(a)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Track not found"})
}


// Connect to Spotify API and print playlist info
func Playlist(id string) {

	// Initialize and assign environment variables
	getEnvVars()
	cid := os.Getenv("ClientID")
	sec := os.Getenv("ClientSecret")

	// Authorize credentials
	authConfig := &clientcredentials.Config{
		ClientID:     cid,
		ClientSecret: sec,
		TokenURL:     spotify.TokenURL,
	}

	// Get token from authorized credentials
	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("Error retrieving access token: %v", err)
	}

	// Initialize client from token
	client := spotify.Authenticator{}.NewClient(accessToken)

	// Get playlist
	playlistID := spotify.ID(id)
	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		log.Fatalf("Error retrieving playlist data: %v", err)
	}

	// Display playlist information
	log.Println("Playlist id:", playlist.ID)
	log.Println("Playlist name:", playlist.Name)
	log.Println("Playlist description:", playlist.Description)
}

// // Connect to Spotify API and print artist info
// // TODO: Consolidate redundant code (generate client once)
// func Artist(artistName string) {
// 	getEnvVars()
// 	cid := os.Getenv("ClientID")
// 	sec := os.Getenv("ClientSecret")

// 	authConfig := &clientcredentials.Config{
// 		ClientID:     cid,
// 		ClientSecret: sec,
// 		TokenURL:     spotify.TokenURL,
// 	}

// 	accessToken, err := authConfig.Token(context.Background())
// 	if err != nil {
// 		log.Fatalf("Error retrieving access token: %v", err)
// 	}

// 	client := spotify.Authenticator{}.NewClient(accessToken)

// 	artist, err := client.Search(artistName, spotify.SearchTypeArtist)
// 	if err != nil {
// 		log.Fatalf("Error retrieving artist data: %v", err)
// 	}

// 	log.Println(artist)
// }

// Connect to Spotify api and returns track search results via:
// spotify.SearchTypeTrack.Tracks
func SearchByTrack(trackName string) *spotify.FullTrackPage {
	getEnvVars()
	cid := os.Getenv("ClientID")
	sec := os.Getenv("ClientSecret")

	authConfig := &clientcredentials.Config{
		ClientID:     cid,
		ClientSecret: sec,
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("Error retrieving access token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)

	result, err := client.Search(trackName, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatalf("Error retrieving track data: %v", err)
	}

	return result.Tracks
}
