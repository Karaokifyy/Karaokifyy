package spotify_api

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

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
func GetAlbums(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(albums)
}

// getAlbumByID locates the album whose ID value matches the id
// Parameter sent by the client, then returns that album as a response
func GetAlbumByID(w http.ResponseWriter, r *http.Request) {
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
	//http.Error(w,"", http.StatusNotFound)
}

// postAlbums adds an album from JSON received in the request body
func PostAlbums(w http.ResponseWriter, r *http.Request) {
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
func GetTrackByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackName := vars["trackName"]
	trackList := SearchByTrack(trackName)

	// Loop over the list of tracks, looking for a track
	// name that matches the parameter given.

	for _, a := range trackList.Tracks {
		json.NewEncoder(w).Encode(a)
	}
	//json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Track not found"})
}


// Connect to Spotify API and print playlist info
func Playlist(id string) {

	// Initialize and assign environment variables
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

//Object to represent a spotify user and their unique/temp session
type SpotifyUserSession struct
{
	Initial_oauth_code string `json:"o_code"`
	Initial_oauth_state string `json:"o_state"`
	Redirect_uri string `json:"redirect_uri"`
	
	Access_token string `json:"access_token"`
	Token_type string `json:"token_type"`
	Scope string `json:"scope"`
	Expires_in int `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`

	client spotify.Client
}

func RequestAccessToken(user *SpotifyUserSession) error {
	if user.Initial_oauth_code == "" {
		return errors.New("spotify user OAUTH code was emppty")
	}

	// Initialize and assign environment variables
	cid := os.Getenv("ClientID")
	sec := os.Getenv("ClientSecret")

	// Authorize credentials
	authConfig := &clientcredentials.Config{
		ClientID:     cid,
		ClientSecret: sec,
		TokenURL:     spotify.TokenURL,
		EndpointParams: map[string][]string{
			"grant_type":{"authorization_code"},
			"code":{user.Initial_oauth_code},
			"redirect_uri":{user.Redirect_uri}},
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("Error retrieving access token: %v", err)
	}

	scope := accessToken.Extra("scope")
	log.Printf("Type of Spotify Scope Extra: %T", scope)

	user.client = spotify.NewAuthenticator(user.Redirect_uri).NewClient(accessToken)
	return nil
}

func GetUserPlaylists(user *SpotifyUserSession) (p_strings []string, err error) {
	playlists, err_pl := user.client.CurrentUsersPlaylists()

	if err_pl != nil{
		err = errors.New("couldn't get user's playlists")
		return
	}
	for _, playlist := range playlists.Playlists{
		p_strings = append(p_strings, playlist.Name)
	}

	return
}