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
type Playlist2 struct {
	PlaylistID   string `json:"PlaylistID"`
	PlaylistName string `json:"PlaylistName"`
}

type Playlist struct {
	ID     string          `json:"ID"`
	Name   string          `json:"name"`
	Images []spotify.Image `json:"images"`
	URI    spotify.URI     `json:"uri"`
}

// Authenticates token and returns a Spotify client
func getSpotifyClient() spotify.Client {
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

// Get a song name from a Spotify song ID
func GetSongName(songID string) string {
	client := getSpotifyClient()

	// Get track data from Spotify
	result, err := client.GetTrack(spotify.ID(songID))
	if err != nil {
		log.Fatalf("Error retrieving track data: %v", err)
	}

	return result.SimpleTrack.Name
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

func SearchByPlaylist(playlistName string) []Playlist2 {
	client := getSpotifyClient()

	// Search for an artist containing "artistName"
	result, err := client.Search(playlistName, spotify.SearchTypePlaylist)
	if err != nil {
		log.Fatalf("Error retrieving artist data: %v", err)
	}

	var playlistList []Playlist2
	length := len(result.Playlists.Playlists)

	for i := 0; i < length; i++ {
		playlist := result.Playlists.Playlists[i]
		playlistID := playlist.ID.String()
		playlistName := playlist.Name
		playlistList = append(playlistList, Playlist2{playlistID, playlistName})
	}

	return playlistList
}

// Fetches a playlist using playlist ID, and returns the tracks from that playlist
// as a slice of song structs
func GetByPlaylistID(playlistID string) []Song {
	client := getSpotifyClient()

	result, err := client.GetPlaylist(spotify.ID(playlistID))
	if err != nil {
		log.Fatalf("Error retrieving playlist data: %v", err)
	}

	var songList []Song
	length := len(result.Tracks.Tracks)

	for i := 0; i < length; i++ {
		track := result.Tracks.Tracks[i].Track
		songID := track.SimpleTrack.ID.String()
		songName := track.SimpleTrack.Name
		artistName := track.SimpleTrack.Artists[0].Name
		songLength := float32(track.SimpleTrack.Duration) / 1000
		albumImg := track.Album.Images[1].URL
		songList = append(songList, Song{songID, songName, artistName, songLength, albumImg})
	}

	return songList
}

func SearchSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	songName := vars["songName"]

	songList := SearchBySong(songName)

	// Loop over the list of tracks, looking for a track
	// name that matches the parameter given.

	json.NewEncoder(w).Encode(songList)
	// song.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Song not found"})
}

func SearchAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	albumName := vars["albumName"]
	albumList := SearchByAlbum(albumName)

	json.NewEncoder(w).Encode(albumList)
}

func SearchArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	artistName := vars["artistName"]
	artistList := SearchByArtist(artistName)

	json.NewEncoder(w).Encode(artistList)
}

func SearchPlaylist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistName := vars["playlistName"]
	playlistList := SearchByPlaylist(playlistName)

	json.NewEncoder(w).Encode(playlistList)
}

// Object to represent a spotify user and their unique/temp session
type SpotifyUserSession struct {
	Initial_oauth_code  string `json:"o_code"`
	Initial_oauth_state string `json:"o_state"`
	Redirect_uri        string `json:"redirect_uri"`

	Access_token  string `json:"access_token"`
	Token_type    string `json:"token_type"`
	Scope         string `json:"scope"`
	Expires_in    int    `json:"expires_in"`
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
			"grant_type":   {"authorization_code"},
			"code":         {user.Initial_oauth_code},
			"redirect_uri": {user.Redirect_uri}},
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

// Getter for a list of a user's Spotify playlists
func GetUserPlaylists(user *SpotifyUserSession) (output map[string]Playlist, err error) {
	playlists, err_pl := user.client.CurrentUsersPlaylists()
	output = make(map[string]Playlist)

	if err_pl != nil {
		err = errors.New("couldn't get user's playlists")
		return
	}
	for _, playlist := range playlists.Playlists {
		output[playlist.Name] = Playlist{string(playlist.ID), playlist.Name, playlist.Images, playlist.URI}
	}

	return
}

// Getter for playlist using a playlist ID
func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	vars := mux.Vars(r)
	playlistID := vars["playlistID"]
	trackList := GetByPlaylistID(playlistID)

	json.NewEncoder(w).Encode(trackList)
}

// Connect to Spotify api and returns track search results via:
// // spotify.SearchTypeTrack.Tracks
// func SearchByTrack(trackName string) *spotify.FullTrackPage {
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

// 	result, err := client.Search(trackName, spotify.SearchTypeTrack)
// 	if err != nil {
// 		log.Fatalf("Error retrieving track data: %v", err)
// 	}

// 	return result.Tracks
// }
