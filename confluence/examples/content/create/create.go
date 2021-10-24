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

	var payload = &confluence.ContentScheme{
		Type:  "page", // Valid values: page, blogpost, comment
		Title: "Confluence Page Title",
		Space: &confluence.SpaceScheme{Key: "DUMMY"},
		Body: &confluence.BodyScheme{
			Storage: &confluence.BodyNodeScheme{
				Value:          "<p>This is <br/> a new page</p>",
				Representation: "storage",
			},
		},
		Ancestors: []*confluence.ContentScheme{
			{
				ID: "78643265",
			},
		},
	}

	newConfluence, response, err := instance.Content.Create(context.Background(), payload)
	if err != nil {

		if response != nil {
			log.Fatal(response.API)
		}
		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)

	log.Println("The new content has been created")
	log.Println(newConfluence.ID)
	log.Println(newConfluence.Links.Self)
	log.Println(newConfluence.Title)
	log.Println(newConfluence.Space.Name)
}
