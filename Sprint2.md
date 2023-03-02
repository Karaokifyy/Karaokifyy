# Sprint 2

## Video Links

## Issues

### Issues Completed
- Front-end and back-end integration
- Routing to different pages with redirection after login
- Used spotify API to access users' private playlists
- Create an API for our front-end to display Spotify search results
  - We completed an API for our application to send filtered JSON results from the Spotify API to our front-end, including searches for songs, albums, artists, and playlists.
- Parse Spotify API results to make them usable for the front-end
  - Structs were implemented in order to create objects for search results that were usable and readable since the Spotify API returns lots of information we didn't need.
- Complete unit testing for our API
- Frontend Cypress and Jasmine Unit tests completed
- Frontend cosmetic code implemented for search screen

### Issues in Progress
- Error handling for Spotify API results to front-end

## Front-end

### Cypress Testing
Utilizing the Cypress software we have implemented 3 tests to run in sprint2.cy.ts. The first test is a simple test designed to visit the initial login page and uses the contains function to find the string "Welcome". The 2nd test simply check to see if the search-screen dashboard is reachable when routed to the search-screen component page. Finally, the third test fills out the login form of the login screen with a sample username/email and password to login. The tests all pass in functionality.

### Unit Testing
-The describe function is used to group multiple tests together that are testing the same thing. In this case, it's the HomeScreenComponent. The let keyword is used to declare variables that will be used within the tests. In this case, component and fixture are declared. The beforeEach function is a hook provided by the testing framework that runs before each test. It sets up the testing environment by configuring the TestBed with the necessary components and services needed for the HomeScreenComponent. It also creates an instance of the HomeScreenComponent and assigns it to the component variable.The first test it('should create the component', () => {...}) is testing whether the HomeScreenComponent was successfully created. It uses the expect function to check whether the component variable has a truthy value, which means it exists and was successfully created. The second test it('should have empty username and password fields on initialization', () => {...}) is testing whether the username and password fields of the HomeScreenComponent are empty when it is initialized. It again uses the expect function to check whether the username and password properties of the component variable are empty strings. The providers array within the TestBed.configureTestingModule method is used to provide any dependencies needed for the HomeScreenComponent. In this case, it provides the UsersService and Router services. The useValue property is used to provide a mock implementation of the services so that the tests can run without actually using the real services. Overall, this unit test ensures that the HomeScreenComponent is successfully created and that its username and password fields are empty when it is initialized.

## Back-end

### Unit Testing 
**List of unit tests completed for back-end located in *spotify_api_test.go*:**
- TestSearchBySong
  - Tests the song search function by passing the string of a song name, and comparing the first result retrieved from the Spotify API against the inputted expected result.
- TestSearchByAlbum
  - Tests the album search function by passing the string of an album name, and comparing the first result retrieved from the Spotify API against the inputted expected result.
- TestSearchByArtist
  - Tests the artist search function by passing the string of an artist name, and comparing the first result retrieved from the Spotify API against the inputted expected result.
- TestSearchByPlaylist
  - Tests the playlist search function by passing the string of a playlist name, and comparing the first result retrieved from the Spotify API against the inputted expected result.


# API Documentation
## _**Table of Contents**_

[General](#general)
   - [What is Karaokiffy?](#what-is-karaokiffy)

[How to use the API](#how-to-use-the-api)
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
