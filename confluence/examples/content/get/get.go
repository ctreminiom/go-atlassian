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
		contentID = "78643301"
		expand    = []string{"any", "ancestors"}
		version   = 1
	)

	content, response, err := instance.Content.Get(context.Background(), contentID, expand, version)
	if err != nil {

		if response != nil {
			log.Println(response.API)
		}

		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)
	log.Println(content)

	for index, ancestors := range content.Ancestors {
		log.Println(index, ancestors.Title)
	}
}
