package youtube_service

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spatialtime/iso8601-go"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

/*
  - Should send back "embed" version of url

- Should  first select the videos with the same duration as the song and then return the url of the song in that set with the most views
*/
func GetYoutubeURL(query string, songDuration int) (resultURL string, err error) {
	resultURL = ""
	err = errors.New("RUh-ROh")
	var maxResults int64 = 25

	songDur, _ := time.ParseDuration(strconv.Itoa(songDuration) + "s")
	secDur, _ := time.ParseDuration("1s")

	developerKey := os.Getenv("YoutubeAPIKey")

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Printf("Error creating new YouTube client: %v", err)
		return resultURL, err
	}

	callHeaders := []string{"snippet"}
	call := service.Search.List(callHeaders).
		Q(query).
		MaxResults(maxResults).Type("video").VideoEmbeddable("true")
	response, err := call.Do()

	if err != nil {
		return resultURL, err
	}

	videos := make(map[string]string)

	var mostViewed *youtube.Video = nil
	var mostViews uint64 = 0
	for _, item := range response.Items {
		videos[item.Id.VideoId] = item.Snippet.Title
		callHeaders2 := []string{"statistics, contentDetails, player"}
		videoCall := service.Videos.List(callHeaders2).Id(item.Id.VideoId)
		videoResponse, err := videoCall.Do()

		if err != nil {
			log.Printf("Error getting stats: %v", err)
			return resultURL, err
		}

		for _, vid := range videoResponse.Items {
			videoDur, _ := iso8601.ParseDuration(vid.ContentDetails.Duration)
			if songDur-secDur <= videoDur && videoDur <= songDur+secDur && vid.Statistics.ViewCount > mostViews {
				mostViewed = vid
				mostViews = vid.Statistics.ViewCount
			}
		}
	}

	if mostViewed != nil {
		resultURL = "https://www.youtube.com/embed/" + mostViewed.Id
		err = nil
	}
	return resultURL, nil
}
