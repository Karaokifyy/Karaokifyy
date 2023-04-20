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
http://localhost:8080/karaokifyy/:59HjlYCeBsxdI0fcm3zglw
```

</td><td>

```sh
{
    {
      "SongURL":     "https://www.youtube.com/embed/CUr_UwUUXzU",
      "SongLyrics":   "[00:13.18] He said, \"Let's get out of this town\n[00:16.59] Drive out of the city, away from the crowds\"\n[00:20.00] I thought Heaven can't help me now\n[00:23.30] Nothing lasts forever\n[00:25.76] But this is gonna take me down\n[00:27.94] He's so tall and handsome as hell\n[00:31.17] He's so bad, but he does it so well\n[00:34.42] I can see the end as it begins\n[00:37.85] My one condition is\n[00:41.25] Say you'll remember me\n[00:43.78] Standing in a nice dress\n[00:45.48] Staring at the sunset, babe\n[00:48.17] Red lips and rosy cheeks\n[00:50.67] Say you'll see me again\n[00:52.57] Even if it's just in your wildest dreams, ah-ah, ha\n[01:01.92] Wildest dreams, ah-ah, ha\n[01:08.58] \n[01:11.51] I said, \"No one has to know what we do\"\n[01:14.56] His hands are in my hair, his clothes are in my room\n[01:18.75] And his voice is a familiar sound\n[01:21.78] Nothing lasts forever\n[01:23.95] But this is getting good now\n[01:26.20] He's so tall and handsome as hell\n[01:29.43] He's so bad, but he does it so well\n[01:32.88] And when we've had our very last kiss\n[01:36.03] My last request is\n[01:39.48] Say you'll remember me\n[01:42.15] Standing in a nice dress\n[01:43.73] Staring at the sunset, babe\n[01:46.58] Red lips and rosy cheeks\n[01:48.97] Say you'll see me again\n[01:50.76] Even if it's just in your wildest dreams, ah-ah, ha (ha-ah, ha)\n[02:00.26] Wildest dreams, ah-ah, ha\n[02:06.86] You'll see me in hindsight\n[02:08.76] Tangled up with you all night\n[02:10.47] Burning it down\n[02:13.82] Someday when you leave me\n[02:15.69] I bet these memories\n[02:17.44] Follow you around\n[02:20.46] You'll see me in hindsight\n[02:22.44] Tangled up with you all night\n[02:24.19] Burning (burning) it (it) down (down)\n[02:27.43] Someday when you leave me\n[02:29.33] I bet these memories\n[02:31.04] Follow (follow) you (you) around (follow you around)\n[02:37.71] Say you'll remember me\n[02:40.68] Standing in a nice dress\n[02:42.28] Staring at the sunset, babe\n[02:44.82] Red lips and rosy cheeks\n[02:47.34] Say you'll see me again\n[02:49.13] Even if it's just pretend\n[02:53.22] Say you'll remember me\n[02:55.89] Standing in a nice dress\n[02:57.51] Staring at the sunset, babe\n[03:00.24] Red lips and rosy cheeks\n[03:02.78] Say you'll see me again\n[03:04.46] Even if it's just (pretend, just pretend) in your wildest dreams, ah-ah, ha (ah)\n[03:13.03] In your wildest dreams, ah-ah, ha\n[03:18.32] Even if it's just stayed in your wildest dreams, ah-ah, ha\n[03:26.78] In your wildest dreams, ah-ah, ha\n[03:33.43]"
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
