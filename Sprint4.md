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
