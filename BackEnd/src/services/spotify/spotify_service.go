package spotify_service

import (
	"context"
	"errors"
	"log"
	"os"

	types "github.com/Karaokifyy/BackEnd/src/models/types"
	spotify_client "github.com/Karaokifyy/BackEnd/src/services/spotify/spotify_client"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// Takes string and returns a Spotify song search
func SearchSongService(songName string) *spotify.SearchResult {
	client := spotify_client.GetSpotifyClient()

	// Search for a song/track containing "songName"
	result, err := client.Search(songName, spotify.SearchTypeTrack)
	if err != nil {
		log.Fatalf("Error retrieving track data: %v", err)
	}

	return result
}

// Takes string and returns a Spotify album search
func SearchAlbumService(albumName string) *spotify.SearchResult {
	client := spotify_client.GetSpotifyClient()

	// Search for an album containing "albumName"
	result, err := client.Search(albumName, spotify.SearchTypeAlbum)
	if err != nil {
		log.Fatalf("Error retrieving album data: %v", err)
	}

	return result
}

// Takes string and returns a Spotify artist search
func SearchArtistService(artistName string) *spotify.SearchResult {
	client := spotify_client.GetSpotifyClient()

	// Search for an album containing "artistName"
	result, err := client.Search(artistName, spotify.SearchTypeArtist)
	if err != nil {
		log.Fatalf("Error retrieving artist data: %v", err)
	}

	return result
}

// Takes string and returns a Spotify playlist search
func SearchPlaylistService(playlistName string) *spotify.SearchResult {
	client := spotify_client.GetSpotifyClient()

	// Search for an album containing "playlistName"
	result, err := client.Search(playlistName, spotify.SearchTypePlaylist)
	if err != nil {
		log.Fatalf("Error retrieving artist data: %v", err)
	}

	return result
}

// Takes string and get Spotify playlist by ID
func GetPlaylistService(playlistID string) *spotify.FullPlaylist {
	client := spotify_client.GetSpotifyClient()

	// Search for an album containing "playlistName"
	result, err := client.GetPlaylist(spotify.ID(playlistID))
	if err != nil {
		log.Fatalf("Error retrieving artist data: %v", err)
	}

	return result
}

// Get a song name from a Spotify song ID
func GetSongName(songID string) string {
	client := spotify_client.GetSpotifyClient()

	// Get track data from Spotify
	result, err := client.GetTrack(spotify.ID(songID))
	if err != nil {
		log.Fatalf("Error retrieving track data: %v", err)
	}

	return result.SimpleTrack.Name
}

func RequestAccessToken(user *types.SpotifyUserSession) error {
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

	user.Client = spotify.NewAuthenticator(user.Redirect_uri).NewClient(accessToken)
	return nil
}

// Getter for a list of a user's Spotify playlists
func GetUserPlaylists(user *types.SpotifyUserSession) (output map[string]types.UserPlaylist, err error) {
	playlists, err_pl := user.Client.CurrentUsersPlaylists()
	output = make(map[string]types.UserPlaylist)

	if err_pl != nil {
		err = errors.New("couldn't get user's playlists")
		return
	}
	for _, playlist := range playlists.Playlists {
		output[playlist.Name] = types.UserPlaylist{ID: string(playlist.ID), Name: playlist.Name, Images: playlist.Images, URI: playlist.URI}
	}

	return
}
