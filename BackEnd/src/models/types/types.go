package types

import (
	"github.com/zmb3/spotify"

	"gorm.io/gorm"
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

// A User Playlist object (for front-end) with filtered JSON
type UserPlaylist struct {
	ID     string          `json:"ID"`
	Name   string          `json:"name"`
	Images []spotify.Image `json:"images"`
	URI    spotify.URI     `json:"uri"`
}

// An LRC Track object (for front-end) with filtered JSON
type Track struct {
	gorm.Model
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ArtistName   string `json:"artistName"`
	AlbumName    string `json:"albumName"`
	Duration     int    `json:"duration"`
	Instrumental bool   `json:"instrumental"`
	Lang         string `json:"lang"`
	ISRC         string `json:"isrc"`
	SpotifyId    string `json:"spotifyId"`
	ReleaseDate  string `json:"releaseDate"`
	PlainLyrics  string `json:"plainLyrics"`
	SyncedLyrics string `json:"syncedLyrics"`
}

// LRC lyrics and Youtube URL object (for front-end) with filtered JSON
type AudioLyrics struct {
	SongURL    string `json:"SongURL"`
	SongLyrics string `json:"SongLyrics"`
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

	Client spotify.Client
}
