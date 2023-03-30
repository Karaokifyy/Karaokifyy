# Sprint 3

## Video Links

---

## Front-end

We added unit tests using the .spec.ts files for each component of our app. These tests verified that the components were rendered correctly.

For this Sprint, we used Cypress tests to verify are routing was working and to simulate going through the application. The simulation goes through the entire cycle of logging in, picking your desired playlist, picking your desired song, and recieving a kareoke video from youtube.
It was also used to look for elements on the screen to validate that the output was correct.

### Issues Completed
   - Display data on songs in a playlist the user selects
   - Display a youtube video with the kareoke version of the song the user selects through an IFrame
   - Created and Utilized an injectable service to send user data from component to component
   - Created unit and cypress tests for work we completed this sprint
   - Continued integration between recieving data from backend

### Issues in Progress
   - Recieving lyrics data from the database to display to user
   - Cleaning up design of webpages
   - Implementing a search feature for a user to find a song without their spotify playlist

---

## Back-end
For the back-end, in Sprint 3 we implemented a GetPlaylistByID function in ```spotify_api.go``` that receives a playlist ID from the front-end and serves the songs from that playlist to the front-end using the ```/user/playlist/:playlistID``` endpoint. 

We also implemented a new API functionality ```karaokifyy_api.go``` that includes our functions for serving a YouTube audio link and LRC lyric file of a specific song using its Spotify song ID to the front-end via the ```/karaokifyy/:songID``` endpoint. Unit tests were also implemented to test our API functionality that we completed in Sprint 3. 

### Issues Completed
   - Create API functions and a router endpoint for serving songs from a user's playlist to the front-end
   - Create API functions and router endpoint for serving youtube link and LRC files to the front-end
   - Complete unit testing of API functions

### Issues in Progress
   - Implementation of database queries into our API
   - Access and utilize YouTube's search using YouTube's API

---

<br />


# API Documentation
## _**Table of Contents**_

[General](#general)
   - [What is Karaokiffy?](#what-is-karaokiffy)

[How to use the API](#how-to-use-the-api)
   - [Get audio link and lyric file for a song](#get-audio-link-and-lyric-file-for-a-song)
   - [Get tracks from a playlist](#get-tracks-from-a-playlist)
   - [Search for a song](#search-for-a-song)
   - [Search for an album](#search-for-an-album)
   - [Search for an artist](#search-for-an-artist)
   - [Search for a playlist](#search-for-a-playlist)

[Other](#other)



---


## __General__
This documentation can be used to access the Karaokifyy API.<br/>
Descriptions of each available endpoint and results are included below.

### __What is Karaokiffy?__

Karaokiffy is an app that utilizes Spotify and Youtube's APIs in conjunction with LRC files from a database for karaoke from a user's Spotify Playlist.<br/> 

---

## __How to use the API__

To perform various searches, you will send a ```GET``` request to the relevant endpoint.
> _Note: Make sure to properly convert invalid characters (e.g. a space ```" "``` should be converted to ```%20```)_

---

### __Get Audio Link and Lyric File for a Song__

To get an audio link and lyric file for a song, you use the ```/karaokifyy/:songID``` endpoint. <br/>
The results will include an audio YouTube link of the song and an LRC file with the lyrics. The information is displayed as:

> -  **_SongURL:_** YouTube link to the audio of the song as a string
> -  **_SongLyrics:_** Song lyrics in LRC format

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

