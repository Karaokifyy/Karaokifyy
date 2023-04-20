package sql_service

import (
	types "github.com/Karaokifyy/BackEnd/src/models/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetLRC(songID string) (string, int) {

	db, err := gorm.Open(sqlite.Open("assets/lyrics_lite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&types.Track{})

	var track types.Track
	if db.First(&track, "spotify_id = ?", songID).Error != nil {
		return "", 0
	}

	return track.SyncedLyrics, track.Duration
}
