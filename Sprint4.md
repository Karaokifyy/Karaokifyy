# Sprint 4

## Video Links

---

## Front-end

We added unit tests using the ```.spec.ts``` files for each component of our app. These tests verified that the components were rendered correctly and our HTTP requests were recieving valid json data. We made sure to make a 1 to 1 ratio for these tests to make sure that everything is working as intended. 

For this Sprint, we used Cypress tests to verify an OAuth Login and Authentication with a 3rd party(in our case spotify). Utilizing this test, we were able simulate running our application in new tests by choosing a valid playlist and song to demo our work, which leads to our end kareoke screen with time synchronized lyrics and a correct youtube iframe. We also used cypress to look for the correct lyrics on the screen to validate that the output was correct(depending on the song).

The code displays lyrics for a YouTube video in sync with the video playback. The code uses Angular to make an HTTP GET request to a local server and fetch the lyrics for the YouTube video. The lyrics are fetched as a JSON object and are then parsed and stored in the httpData array. The timeArray array stores the start times of each line of the lyrics, and the textArray array stores the corresponding text for each line. The displaylyrics function is called repeatedly using setInterval at an interval of timeInterval milliseconds. This function is responsible for displaying the lyrics at the appropriate time in sync with the video playback. The counter variable is used to keep track of which line of the lyrics is currently being displayed. In the displaylyrics function, the start time of the current line and the next line are retrieved from the timeArray. The start time is converted from minutes and seconds to seconds only. If the current time of the video is greater than or equal to the start time of the next line, the displayText variable is updated with the text of the next line, and the counter variable is incremented to move on to the next line of lyrics. Once all the lines of lyrics have been displayed, the setInterval is cleared to stop calling the displaylyrics function.

### Issues Completed
   - Get lyrics of songs from database based off spotify song IDs  
   - Displayed time synchronized lyrics in accordance to iframe video 
   - Fixed CORS errors from Sprint 3
   - Updated UI of application
   - Unit and Cypress tests
---

## Back-end


### Issues Completed

---

<br />


# API Documentation
## _**Table of Contents**_

[General](#general)
   - [What is Karaokiffy?](#what-is-karaokiffy)

[How to use the API](#how-to-use-the-api)
   - [Get YouTube link and lyric file for a song](#get-youtube-link-and-lyric-file-for-a-song)
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

### __Get YouTube Link and Lyric File for a Song__

To get a YouTube embedded link and lyric file for a song, you use the ```/karaokifyy/:songID``` endpoint. <br/>
The results will include an audio YouTube link of the song and an LRC file with the lyrics. The information is displayed as:

> -  **_SongURL:_** YouTube embedded link to the audio of the song as a string
> -  **_SongLyrics:_** Song lyrics in LRC format

An example search and response is shown below.

<table><tr><th> Example Request </th><th> Example Response </th></tr><tr><td>

```sh
http://localhost:8080/karaokifyy/:songID
```

</td><td>

```sh
{
    {
      "SongURL":     "",
      "SongLyrics":   ""
    }, ...
}
```

</td></tr></table>

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
> See: https://developer.spotify.com/documentation/web-api/

<br/>

The API backend uses the ```gorilla/mux``` Go package for request routing and handling.<br/>
> See: https://www.gorillatoolkit.org/pkg/mux
