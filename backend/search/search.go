package search

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

func searchBySong(songName string) *spotify.FullTrackPage {
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

	result, err := client.Search(songName, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatalf("Error retrieving track data: %v", err)
	}

	return result.Tracks
}

func searchSong(song *gin.Context) {
	songName := song.Param("songName")
	songList := searchBySong(songName)

	// Loop over the list of tracks, looking for a track
	// name that matches the parameter given.

	for _, a := range songList.Tracks {
		song.IndentedJSON(http.StatusOK, a)
	}
	song.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Song not found"})
}

func Init(router *gin.Engine) {
	router.GET("/search/song/:songName", searchSong)
}
