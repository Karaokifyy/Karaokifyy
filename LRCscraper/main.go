package main

import (
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"encoding/json"

)

type Track struct {
	gorm.Model
    id 				int 	`json:"id"`
	name 			string 	`json:"name"`
	artistName 		string 	`json:"artistName"`
	albumName 		string 	`json:"albumName"`
	duration 		int 	`json:"duration"`
	instrumental 	bool 	`json:"instrumental"`
	lang 			string 	`json:"lang"`
	isrc 			string 	`json:"isrc"`
	spotifyId 		string 	`json:"spotifyId"`
	releaseDate 	string 	`json:"releaseDate"`
	plainLyrics 	string	`json:"plainLyrics"`
	syncedLyrics 	string 	`json:"syncedLyrics"`
}


func scrape() {
	db, err := gorm.Open(sqlite.Open("lyrics.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		}

	db.AutoMigrate(&Track{})

	for i:=1; i <= 579899; i++{
		var newTrack Track

		requestURL := fmt.Sprintf("https://lrclib.net/api/get/%d", i)
		res, err := http.Get(requestURL)
		if err != nil {
			log.Printf("error making http request for track ID: %d\n", i)
		}

		if(res.StatusCode == 200){
			if err := json.NewDecoder(res.Body).Decode(&newTrack); err != nil {
				log.Printf("error decoding lrclib response: %d\n", i)
			}

			db.Create(newTrack)//add new track to database
		}else{
			log.Printf("error making http request for track ID: %d, statusCode: %d\n", i, res.StatusCode)
		}
	}

}

