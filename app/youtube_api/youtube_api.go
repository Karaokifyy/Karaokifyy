package youtube_api

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)


func GetYoutubeURL(query string) (resultURL string, err error) {
		var maxResults int64 = 25

		developerKey := os.Getenv("YoutubeAPIKey")

        client := &http.Client{
                Transport: &transport.APIKey{Key: developerKey},
        }

        service, err := youtube.New(client)
        if err != nil {
                log.Printf("Error creating new YouTube client: %v", err)
				return resultURL, err
        }

        // Make the API call to YouTube.
        call := service.Search.List("id,snippet").
                Q(query).
                MaxResults(maxResults)
        response, err := call.Do()
		
		if(err != nil)
		{
			return resultURL, err
		}

        // Group video, channel, and playlist results in separate lists.
        videos := make(map[string]string)

        // Iterate through each item and add it to the correct list.
        for _, item := range response.Items {
                switch item.Id.Kind {
                case "youtube#video":
                        videos[item.Id.VideoId] = item.Snippet.Title
                case "youtube#channel":
                        channels[item.Id.ChannelId] = item.Snippet.Title
                case "youtube#playlist":
                        playlists[item.Id.PlaylistId] = item.Snippet.Title
                }
        }
}