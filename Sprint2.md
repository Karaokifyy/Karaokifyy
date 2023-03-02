# Sprint 2

## Video Links

## Issues

### Issues Completed
- Front-end and back-end integration
- Create an API for our front-end to display Spotify search results
  - We completed an API for our application to send filtered JSON results from the Spotify API to our front-end, including searches for songs, albums, artists, and playlists.
- Parse Spotify API results to make them usable for the front-end
  - Structs were implemented in order to create objects for search results that were usable and readable since the Spotify API returns lots of information we didn't need.
- Complete unit testing for our API

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


### API Documentation
