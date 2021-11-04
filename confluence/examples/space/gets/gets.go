package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"net/http"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	instance, err := confluence.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	instance.Auth.SetBasicAuth(mail, token)
	instance.Auth.SetUserAgent("curl/7.54.0")

	var (
		options = &confluence.GetSpacesOptionScheme{
			SpaceKeys:       nil,
			SpaceIDs:        []int{80838666, 196613},
			SpaceType:       "",
			Status:          "",
			Labels:          nil,
			Favorite:        false,
			FavoriteUserKey: "",
			Expand:          nil,
		}
		startAt    = 0
		maxResults = 25
	)

	spaces, response, err := instance.Space.Gets(context.Background(), options, startAt, maxResults)
	if err != nil {

		if response != nil {
			if response.Code == http.StatusBadRequest {
				log.Println(response.API)
			}
		}
		log.Println("Endpoint:", response.Endpoint)
		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)

	for _, space := range spaces.Results {
		log.Println(space.Name, space.Key, space.ID)
	}
}
