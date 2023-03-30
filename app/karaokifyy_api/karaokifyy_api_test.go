package karaokifyy_api

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

// Load environment vars from file
func getEnvVars() {
	err := godotenv.Load("../config/cred.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestGetAudioLyrics(t *testing.T) {
	getEnvVars()
	audio, lyrics := GetAudioLyrics("5ihS6UUlyQAfmp48eSkxuQ")
	result := AudioLyrics{audio, lyrics}
	expected := AudioLyrics{
		"https://www.youtube.com/watch?v=radHy4HhhNg&ab_channel=FleetwoodMac-Topic", `[00:12.62]I took my love, I took it down
[00:18.40]Climbed a mountain and I turned around
[00:23.11]And I saw my reflection in the snow covered hills
[00:29.47]'Til the landslide brought it down

[00:35.98]Oh, mirror in the sky, What is love?
[00:41.85]Can the child within my heart rise above?
[00:48.01]Can I sail thru the changin' ocean tides?
[00:53.80]Can I handle the seasons of my life?

[01:11.77]Well, I've been afraid of changin'
[01:16.71]'Cause I've built my life around you
[01:23.39]But time makes you bolder
[01:26.65]Even children get older
[01:30.02]And I'm getting older too
[02:00.42]Well, I've been afraid of changin'
[02:05.67]'Cause I've built my life around you
[02:12.48]But time makes you bolder
[02:15.52]Even children get older
[02:18.70]And I'm getting older too
[02:24.88]And I'm getting older too

[02:31.02]Ahh, take my love, take it down
[02:38.34]Ahh, Climb a mountain and turn around
[02:43.03]And if you see my reflection in the snow covered hills
[02:49.30]Well the landslide will bring it down
[02:54.85]And if you see my reflection in the snow covered hills
[03:02.73]Well the landslide will bring it down

[03:09.14]Oh, The landslide will bring it down`}

	if (result.SongURL != expected.SongURL) ||
		(result.SongLyrics != expected.SongLyrics) {
		t.Errorf("result: %v\n expected: %v", result, expected)
	}
}
