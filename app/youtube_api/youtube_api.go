package youtube_api

import (
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
		// var callHeaders []string 
		// callHeaders[0] - "id"
		// callHeaders[1] = "snippet"
		callHeaders  :=  []string{"id,snippet"}
        call := service.Search.List(callHeaders).
                Q(query).
                MaxResults(maxResults)
        response, err := call.Do()
		
		if(err != nil){
			return resultURL, err
		}

        // Group video, channel, and playlist results in separate lists.
        videos := make(map[string]string)

		mostViewed := ""
		highestCount := 0
        // Iterate through each item and add it to the correct list.
        for _, item := range response.Items {
                switch item.Id.Kind {
                case "youtube#video":
                        videos[item.Id.VideoId] = item.Snippet.Title
						callHeaders2  :=  []string{"statistics"}
						videoCall := service.Videos.List(callHeaders2).Id(item.Id.VideoId)
						videoResponse, err := videoCall.Do()

						if err != nil {
							log.Printf("Error getting stats: %v", err)
							return resultURL, err
						}

						for _, field := range videoResponse.Items{
							if field.Statistics.ViewCount > uint64(highestCount){
								mostViewed = item.Id.VideoId
							}
						}
                }
        }

		return mostViewed, nil
}