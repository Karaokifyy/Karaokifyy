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
	audio, lyrics := GetAudioLyrics("2lXDbdbBw7lJVCWHkYxdep")
	result := AudioLyrics{audio, lyrics}
	expected := AudioLyrics{
		"FyZybpwOBoY", "[00:13.18] He said, \"Let's get out of this town\n[00:16.59] Drive out of the city, away from the crowds\"\n[00:20.00] I thought Heaven can't help me now\n[00:23.30] Nothing lasts forever\n[00:25.76] But this is gonna take me down\n[00:27.94] He's so tall and handsome as hell\n[00:31.17] He's so bad, but he does it so well\n[00:34.42] I can see the end as it begins\n[00:37.85] My one condition is\n[00:41.25] Say you'll remember me\n[00:43.78] Standing in a nice dress\n[00:45.48] Staring at the sunset, babe\n[00:48.17] Red lips and rosy cheeks\n[00:50.67] Say you'll see me again\n[00:52.57] Even if it's just in your wildest dreams, ah-ah, ha\n[01:01.92] Wildest dreams, ah-ah, ha\n[01:08.58] \n[01:11.51] I said, \"No one has to know what we do\"\n[01:14.56] His hands are in my hair, his clothes are in my room\n[01:18.75] And his voice is a familiar sound\n[01:21.78] Nothing lasts forever\n[01:23.95] But this is getting good now\n[01:26.20] He's so tall and handsome as hell\n[01:29.43] He's so bad, but he does it so well\n[01:32.88] And when we've had our very last kiss\n[01:36.03] My last request is\n[01:39.48] Say you'll remember me\n[01:42.15] Standing in a nice dress\n[01:43.73] Staring at the sunset, babe\n[01:46.58] Red lips and rosy cheeks\n[01:48.97] Say you'll see me again\n[01:50.76] Even if it's just in your wildest dreams, ah-ah, ha (ha-ah, ha)\n[02:00.26] Wildest dreams, ah-ah, ha\n[02:06.86] You'll see me in hindsight\n[02:08.76] Tangled up with you all night\n[02:10.47] Burning it down\n[02:13.82] Someday when you leave me\n[02:15.69] I bet these memories\n[02:17.44] Follow you around\n[02:20.46] You'll see me in hindsight\n[02:22.44] Tangled up with you all night\n[02:24.19] Burning (burning) it (it) down (down)\n[02:27.43] Someday when you leave me\n[02:29.33] I bet these memories\n[02:31.04] Follow (follow) you (you) around (follow you around)\n[02:37.71] Say you'll remember me\n[02:40.68] Standing in a nice dress\n[02:42.28] Staring at the sunset, babe\n[02:44.82] Red lips and rosy cheeks\n[02:47.34] Say you'll see me again\n[02:49.13] Even if it's just pretend\n[02:53.22] Say you'll remember me\n[02:55.89] Standing in a nice dress\n[02:57.51] Staring at the sunset, babe\n[03:00.24] Red lips and rosy cheeks\n[03:02.78] Say you'll see me again\n[03:04.46] Even if it's just (pretend, just pretend) in your wildest dreams, ah-ah, ha (ah)\n[03:13.03] In your wildest dreams, ah-ah, ha\n[03:18.32] Even if it's just stayed in your wildest dreams, ah-ah, ha\n[03:26.78] In your wildest dreams, ah-ah, ha\n[03:33.43] "}

	if (result.SongURL != expected.SongURL) ||
		(result.SongLyrics != expected.SongLyrics) {
		t.Errorf("result: %v\n expected: %v", result, expected)
	}
}
