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
    ID 				int 	`json:"id"`
	Name 			string 	`json:"name"`
	ArtistName 		string 	`json:"artistName"`
	AlbumName 		string 	`json:"albumName"`
	Duration 		int 	`json:"duration"`
	Instrumental 	bool 	`json:"instrumental"`
	Lang 			string 	`json:"lang"`
	ISRC 			string 	`json:"isrc"`
	SpotifyId 		string 	`json:"spotifyId"`
	ReleaseDate 	string 	`json:"releaseDate"`
	PlainLyrics 	string	`json:"plainLyrics"`
	SyncedLyrics 	string 	`json:"syncedLyrics"`
}


func scrape() {
	fmt.Println("Scraping Lyrics")
	db, err := gorm.Open(sqlite.Open("lyrics.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
		}

	db.AutoMigrate(&Track{})

	for i:=1; i <= 579899; i++{
		fmt.Printf("Working on track #%d\n", i)
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

			db.Create(&newTrack)//add new track to database
		}else{
			log.Printf("error making http request for track ID: %d, statusCode: %d\n", i, res.StatusCode)
		}

		res.Body.Close()
	}

}

func main(){
	scrape()
}

