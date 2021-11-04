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

	var contentID = "74219528"

	response, err := instance.Content.Delete(context.Background(), contentID, "")
	if err != nil {

		if response != nil {
			log.Println(response.API)
			log.Println(response)
		}

		log.Fatal(err)
	}

	log.Println(response)

}
