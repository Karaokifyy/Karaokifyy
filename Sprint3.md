# Sprint 3

## Video Links

---

## Issues

### Issues Completed

### Issues in Progress

---

## Front-end

## Back-end

---

<br />


# API Documentation
## _**Table of Contents**_

[General](#general)
   - [What is Karaokiffy?](#what-is-karaokiffy)

[How to use the API](#how-to-use-the-api)
   - [Get tracks from a playlist](#get-tracks-from-a-playlist)
   - [Search for a song](#search-for-a-song)
   - [Search for an album](#search-for-an-album)
   - [Search for an artist](#search-for-an-artist)
   - [Search for a playlist](#search-for-a-playlist)

[Other](#other)



---


## __General__
This documentation can be used to access the Karaokiffy API.<br/>
Descriptions of each available endpoint and results are included below.

### __What is Karaokiffy?__

Karaokiffy is an app for searching Spotify to make and modify a karaoke queue.<br/>
Each user can store and share their queues, using Spotify's Playlist functionality.

---

## __How to use the API__

To perform various searches, you will send a ```GET``` request to the relevant endpoint.
> _Note: Make sure to properly convert invalid characters (e.g. a space ```" "``` should be converted to ```%20```)_

---

### __Get Tracks from a Playlist__

To get tracks from a playlist, you use the ```/user/playlist/:playlistID``` endpoint. <br/>
The results will include all of the tracks in a user's playlist, with the following information:

  
> -  **_SongID:_** Unique identifier for the song (string)
> -  **_SongName:_** Name of the song (string)
> -  **_ArtistName:_** Name of the artist (string)
> -  **_SongLength:_** Length of the song, in seconds (float)
> -  **_AlbumImg:_** URL to an image of the album (string)

An example search and response is shown below.

<table><tr><th> Example Request </th><th> Example Response </th></tr><tr><td>

```sh
http://localhost:8080/user/playlist/1qgvTUX3QEOAJ9HMMxpW2a
```

</td><td>

```sh
{
    {
      "SongID":     "5ihS6UUlyQAfmp48eSkxuQ",
      "SongName":   "Landslide",
      "ArtistName": "Fleetwood Mac",
      "SongLength": 199.493,
      "AlbumImg":   "https://i.scdn.co/image/ab67616d00001e024fb043195e8d07e72edc7226"
    }, ...
}
```

</td></tr></table>

---

### __Search for a Song__

To search for a specific song, you use the ```/search/song/:songName``` endpoint. <br/>
The response will include the top 20 search results, with the following information:

  
> -  **_SongID:_** Unique identifier for the song (string)
> -  **_SongName:_** Name of the song (string)
> -  **_ArtistName:_** Name of the artist (string)
> -  **_SongLength:_** Length of the song, in seconds (float)
> -  **_AlbumImg:_** URL to an image of the album (string)

An example search and response is shown below.

<table><tr><th> Example Request </th><th> Example Response </th></tr><tr><td>

```sh
http://localhost:40/search/song/Bohemian%20Rhapsody
```

</td><td>

```sh
{
    {
      "SongID":     "7tFiyTwD0nx5a1eklYtX2J",
      "SongName":   "Bohemian Rhapsody - Remastered 2011",
      "ArtistName": "Queen",
      "SongLength": 354.32,
      "AlbumImg":   "https://i.scdn.co/image/ab67616d00001e02ce4f1737bc8a646c8c4bd25a"
    }, ...
}
```

</td></tr></table>

---

### __Search for an album__

To search for a specific album, you use the ```/search/album/:albumName``` endpoint. <br/>
The response will include the top 20 search results, with the following information:

  
> -  **_AlbumID:_** Unique identifier for the album (string)
> -  **_AlbumName:_** Name of the album (string)
> -  **_ArtistName:_** Name of the artist (string)
> -  **_AlbumImg:_** URL to an image of the album (string)

An example search and response is shown below.

<table><tr><th> Example Request </th><th> Example Response </th></tr><tr><td>

```sh
http://localhost:40/search/album/Party
```

</td><td>

```sh
{
    {
      "AlbumID":    "1xwhNJCfTwuRia7Cpo7IbJ",
      "AlbumName":  "PARTYNEXTDOOR TWO",
      "ArtistName": "PARTYNEXTDOOR",
      "AlbumImg":   "https://i.scdn.co/image/ab67616d00001e026cfa297b0178fd91dca5c4f1"
    }, ...
}
```

</td></tr></table>

---

### __Search for an artist__

To search for a specific artist, you use the ```/search/artist/:artistName``` endpoint. <br/>
The response will include the top 20 search results, with the following information:

  
> -  **_ArtistID:_** Unique identifier for the artist (string)
> -  **_ArtistName:_** Name of the artist (string)

An example search and response is shown below.

<table><tr><th> Example Request </th><th> Example Response </th></tr><tr><td>

```sh
http://localhost:40/search/artist/Queen
```

</td><td>

```sh
{
    {
      "ArtistID":    "1dfeR4HaWDbWqFHLkxsg1d",
      "ArtistName":  "Queen"
    }, ...
}
```

</td></tr></table>

---

### __Search for a playlist__

To search for a specific playlist, you use the ```/search/playlist/:playlistName``` endpoint. <br/>
The response will include the top 20 search results, with the following information:

  
> -  **_PlaylistID:_** Unique identifier for the playlist (string)
> -  **_PlaylistName:_** Name of the playlist (string)

An example search and response is shown below.

<table><tr><th> Example Request </th><th> Example Response </th></tr><tr><td>

```sh
http://localhost:40/search/playlist/Party
```

</td><td>

```sh
{
    {
      "PlaylistID":    "37i9dQZF1DWSIO2QWRavWZ",
      "PlaylistName":  "Queen"
    }, ...
}
```

</td></tr></table>

---

## __Other__
The API backend uses the ```Spotify Web API``` to retrieve song-related data.<br/>
<img src="https://developer.spotify.com/assets/branding-guidelines/logos.svg"
     alt="Spotify"/>
> See: https://developer.spotify.com/documentation/web-api/

<br/>

The API backend uses the ```gorilla/mux``` Go package for request routing and handling.<br/>
<img src="https://cloud-cdn.questionable.services/gorilla-icon-64.png"
     alt="GorillaMux"/>

> See: https://www.gorillatoolkit.org/pkg/mux
