package model

import (
	types "github.com/Karaokifyy/BackEnd/src/models/types"
	"github.com/zmb3/spotify"
)

// Generates song list from types according to search result
func SongListModel(result *spotify.SearchResult) []types.Song {
	var songList []types.Song
	length := len(result.Tracks.Tracks)

	for i := 0; i < length; i++ {
		track := result.Tracks.Tracks[i]
		songID := track.SimpleTrack.ID.String()
		songName := track.SimpleTrack.Name
		artistName := track.SimpleTrack.Artists[0].Name
		songLength := float32(track.SimpleTrack.Duration) / 1000
		albumImg := track.Album.Images[1].URL
		songList = append(songList, types.Song{SongID: songID, SongName: songName, ArtistName: artistName, SongLength: songLength, AlbumImg: albumImg})
	}

	return songList
}

// Generates song list from types according to search result
func AlbumListModel(result *spotify.SearchResult) []types.Album {
	var albumList []types.Album
	length := len(result.Albums.Albums)

	for i := 0; i < length; i++ {
		album := result.Albums.Albums[i]
		albumID := album.ID.String()
		albumName := album.Name
		artistName := album.Artists[0].Name
		albumImg := album.Images[1].URL
		albumList = append(albumList, types.Album{AlbumID: albumID, AlbumName: albumName, ArtistName: artistName, AlbumImg: albumImg})
	}

	return albumList
}

// Generates song list from types according to search result
func ArtistListModel(result *spotify.SearchResult) []types.Artist {
	var artistList []types.Artist
	length := len(result.Artists.Artists)

	for i := 0; i < length; i++ {
		artist := result.Artists.Artists[i]
		artistID := artist.SimpleArtist.ID.String()
		artistName := artist.SimpleArtist.Name
		artistList = append(artistList, types.Artist{ArtistID: artistID, ArtistName: artistName})
	}

	return artistList
}

// Generates song list from types according to search result
func PlaylistListModel(result *spotify.SearchResult) []types.Playlist {
	var playlistList []types.Playlist
	length := len(result.Playlists.Playlists)

	for i := 0; i < length; i++ {
		playlist := result.Playlists.Playlists[i]
		playlistID := playlist.ID.String()
		playlistName := playlist.Name
		playlistList = append(playlistList, types.Playlist{PlaylistID: playlistID, PlaylistName: playlistName})
	}

	return playlistList
}

// Generates song list from types according to search result
func UserPlaylistModel(result *spotify.FullPlaylist) []types.Song {
	var songList []types.Song
	length := len(result.Tracks.Tracks)

	for i := 0; i < length; i++ {
		track := result.Tracks.Tracks[i].Track
		songID := track.SimpleTrack.ID.String()
		songName := track.SimpleTrack.Name
		artistName := track.SimpleTrack.Artists[0].Name
		songLength := float32(track.SimpleTrack.Duration) / 1000
		albumImg := track.Album.Images[1].URL
		songList = append(songList, types.Song{SongID: songID, SongName: songName, ArtistName: artistName, SongLength: songLength, AlbumImg: albumImg})
	}

	return songList
}

// Generates audio and lyrics struct
func AudioLyricsModel(songURL string, songLyrics string) types.AudioLyrics {
	return types.AudioLyrics{SongURL: songURL, SongLyrics: songLyrics}
}
