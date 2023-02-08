package server

import (
	"encoding/json"
	"net/http"
	"os"

	"Karaokifyy/Back-End/spotify_api"

	"github.com/gin-gonic/gin"
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
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// Parameter sent by the client, then returns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter
	for _, a := range albums.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Album not found"})
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to
	// newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the list
	albums.Albums = append(albums.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getTrackByName searches for a track by name using spotify_api.SearchTrackByName method and
// serves resulting json data.
func getTrackByName(c *gin.Context) {
	trackName := c.Param("trackName")
	trackList := spotify_api.SearchByTrack(trackName)

	// Loop over the list of tracks, looking for a track
	// name that matches the parameter given.

	for _, a := range trackList.Tracks {
		c.IndentedJSON(http.StatusOK, a)
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "track not found"})
}

func RunServer() {
	router := gin.Default()
	router.GET("/track/:trackName", getTrackByName)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
