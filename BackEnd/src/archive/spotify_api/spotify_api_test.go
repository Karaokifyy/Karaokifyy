package spotify_api

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

// Load environment vars from file
func getEnvVars() {
	err := godotenv.Load("../config/cred.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestSearchSongInterface(t *testing.T) {
	getEnvVars()
	result := SearchSongInterface("Bohemian%20Rhapsody")[0]
	expected := Song{
		"7tFiyTwD0nx5a1eklYtX2J",
		"Bohemian Rhapsody - Remastered 2011",
		"Queen",
		354.32,
		"https://i.scdn.co/image/ab67616d00001e02ce4f1737bc8a646c8c4bd25a"}

	if (result.SongID != expected.SongID) ||
		(result.SongName != expected.SongName) ||
		(result.ArtistName != expected.ArtistName) ||
		(result.SongLength != expected.SongLength) ||
		(result.AlbumImg != expected.AlbumImg) {
		t.Errorf("result: %v\n expected: %v", result, expected)
	}
}

func TestSearchAlbumInterface(t *testing.T) {
	getEnvVars()
	result := SearchAlbumInterface("Party")[0]
	expected := Album{
		"1xwhNJCfTwuRia7Cpo7IbJ",
		"PARTYNEXTDOOR TWO",
		"PARTYNEXTDOOR",
		"https://i.scdn.co/image/ab67616d00001e026cfa297b0178fd91dca5c4f1"}

	if (result.AlbumID != expected.AlbumID) ||
		(result.AlbumName != expected.AlbumName) ||
		(result.ArtistName != expected.ArtistName) ||
		(result.AlbumImg != expected.AlbumImg) {
		t.Errorf("result: %v\n expected: %v", result, expected)
	}
}

func TestSearchArtistInterface(t *testing.T) {
	getEnvVars()
	result := SearchArtistInterface("Queen")[0]
	expected := Artist{"1dfeR4HaWDbWqFHLkxsg1d", "Queen"}

	if (result.ArtistID != expected.ArtistID) ||
		(result.ArtistName != expected.ArtistName) {
		t.Errorf("result: %v\n expected: %v", result, expected)
	}
}

func TestSearchPlaylistInterface(t *testing.T) {
	getEnvVars()
	result := SearchPlaylistInterface("Queen")[0]
	expected := Playlist2{"37i9dQZF1DWSIO2QWRavWZ", "Queen"}

	if (result.PlaylistID != expected.PlaylistID) ||
		(result.PlaylistName != expected.PlaylistName) {
		t.Errorf("result: %v\n expected: %v", result, expected)
	}
}
