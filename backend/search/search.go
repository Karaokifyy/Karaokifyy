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

// A Song object (for front-end) with filtered JSON
type Song struct {
	SongID     string  `json:"SongID"`
	SongName   string  `json:"SongName"`
	ArtistName string  `json:"ArtistName"`
	SongLength float32 `json:"SongLength"`
	AlbumImg   string  `json:"AlbumImg"`
}

// An Artist object (for front-end) with filtered JSON
type Artist struct {
	ArtistID   string `json:"ArtistID"`
	ArtistName string `json:"ArtistName"`
}

// An Album object (for front-end) with filtered JSON
type Album struct {
	AlbumID    string `json:"AlbumID"`
	AlbumName  string `json:"AlbumName"`
	ArtistName string `json:"ArtistName"`
	AlbumImg   string `json:"AlbumImg"`
}

// A Playlist object (for front-end) with filtered JSON
type Playlist struct {
	PlaylistID   string `json:"PlaylistID"`
	PlaylistName string `json:"PlaylistName"`
}

// Load environment vars from file
func getEnvVars() {
	err := godotenv.Load("./config/cred.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Authenticates token and returns a Spotify client
func getSpotifyClient() spotify.Client {
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

	return spotify.Authenticator{}.NewClient(accessToken)
}

// Takes string and performs a Spotify song search
// Returns a slice of Song structs
func SearchBySong(songName string) []Song {
	client := getSpotifyClient()

	// Search for a song/track containing "songName"
	result, err := client.Search(songName, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatalf("Error retrieving track data: %v", err)
	}

	var songList []Song
	length := len(result.Tracks.Tracks)

	for i := 0; i < length; i++ {
		track := result.Tracks.Tracks[i]
		songID := track.SimpleTrack.ID.String()
		songName := track.SimpleTrack.Name
		artistName := track.SimpleTrack.Artists[0].Name
		songLength := float32(track.SimpleTrack.Duration) / 1000
		albumImg := track.Album.Images[1].URL
		songList = append(songList, Song{songID, songName, artistName, songLength, albumImg})
	}

	return songList
}

func SearchByAlbum(albumName string) []Album {
	client := getSpotifyClient()

	// Search for an album containing "albumName"
	result, err := client.Search(albumName, spotify.SearchTypeAlbum)
	if err != nil {
		log.Fatalf("Error retrieving album data: %v", err)
	}

	var albumList []Album
	length := len(result.Albums.Albums)

	for i := 0; i < length; i++ {
		album := result.Albums.Albums[i]
		albumID := album.ID.String()
		albumName := album.Name
		artistName := album.Artists[0].Name
		albumImg := album.Images[1].URL
		albumList = append(albumList, Album{albumID, albumName, artistName, albumImg})
	}

	return albumList
}

func SearchByArtist(artistName string) []Artist {
	client := getSpotifyClient()

	// Search for an artist containing "artistName"
	result, err := client.Search(artistName, spotify.SearchTypeArtist)
	if err != nil {
		log.Fatalf("Error retrieving artist data: %v", err)
	}

	var artistList []Artist
	length := len(result.Artists.Artists)

	for i := 0; i < length; i++ {
		artist := result.Artists.Artists[i]
		artistID := artist.SimpleArtist.ID.String()
		artistName := artist.SimpleArtist.Name
		artistList = append(artistList, Artist{artistID, artistName})
	}

	return artistList
}

func SearchByPlaylist(playlistName string) []Playlist {
	client := getSpotifyClient()

	// Search for an artist containing "artistName"
	result, err := client.Search(playlistName, spotify.SearchTypePlaylist)
	if err != nil {
		log.Fatalf("Error retrieving artist data: %v", err)
	}

	var playlistList []Playlist
	length := len(result.Playlists.Playlists)

	for i := 0; i < length; i++ {
		playlist := result.Playlists.Playlists[i]
		playlistID := playlist.ID.String()
		playlistName := playlist.Name
		playlistList = append(playlistList, Playlist{playlistID, playlistName})
	}

	return playlistList
}

func searchSong(song *gin.Context) {
	songName := song.Param("songName")
	songList := SearchBySong(songName)

	// Loop over the list of tracks, looking for a track
	// name that matches the parameter given.

	song.IndentedJSON(http.StatusOK, songList)

	// song.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Song not found"})
}

func searchAlbum(album *gin.Context) {
	albumName := album.Param("albumName")
	albumList := SearchByAlbum(albumName)

	album.IndentedJSON(http.StatusOK, albumList)
}

func searchArtist(artist *gin.Context) {
	artistName := artist.Param("artistName")
	artistList := SearchByArtist(artistName)

	artist.IndentedJSON(http.StatusOK, artistList)
}

func searchPlaylist(playlist *gin.Context) {
	playlistName := playlist.Param("playlistName")
	playlistList := SearchByPlaylist(playlistName)

	playlist.IndentedJSON(http.StatusOK, playlistList)
}

func Init(router *gin.Engine) {
	router.GET("/search/song/:songName", searchSong)
	router.GET("/search/album/:albumName", searchAlbum)
	router.GET("/search/artist/:artistName", searchArtist)
	router.GET("/search/playlist/:playlistName", searchPlaylist)

}
