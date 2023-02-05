package spotify_api

import (
	"context"
	"log"
	"os"

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

// Connect to Spotify API and print artist info
// TODO: Consolidate redundant code (generate client once)
func Artist(artistName string) {
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

	playlist, err := client.Search(artistName, spotify.SearchTypeArtist)
	if err != nil {
		log.Fatalf("Error retrieving playlist data: %v", err)
	}

	log.Println(playlist)
}
