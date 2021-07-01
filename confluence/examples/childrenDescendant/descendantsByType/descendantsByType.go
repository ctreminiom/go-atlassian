package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
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
		contentID   = "76513281"
		contentType = "page"
		expand      = []string{"childTypes.all"}
		depth       = "1"
		starAt      = 0
		maxResult   = 50
	)

	contents, response, err := instance.Content.ChildrenDescendant.DescendantsByType(context.Background(), contentID, contentType,
		depth, expand, starAt, maxResult)

	if err != nil {

		if response != nil {
			log.Println(response.API)
		}

		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)

	for _, content := range contents.Results {
		log.Println(content.Type, content.Title, content.ID)
	}

}
