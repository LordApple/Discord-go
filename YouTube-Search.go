package main

import (
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func findVideo(name string) ([]string, error) {
	developerKey := os.Getenv("YouTubeKey")

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		return nil, err
	}

	call := service.Search.List("id,snippet").
		Q(name).
		MaxResults(3)
	response, err := call.Do()

	videos := make(map[string]string)

	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
		}
	}

	list := []string{}

	x := 0
	for id := range videos {
		list = append(list, "https://www.youtube.com/watch?v="+id)
		x++
		if x == 3 {
			break
		}
	}
	return list, nil
}
