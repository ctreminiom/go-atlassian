package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"os"
)

func main()  {

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
		Title: "Confluence Page Title - Updated",
		Body: &confluence.BodyScheme{
			Storage: &confluence.BodyNodeScheme{
				Value:          "<p>This is <br/> a new page - updated</p>",
				Representation: "storage",
			},
		},
		Version: &confluence.VersionScheme{Number: 2},
	}

	content, response, err := instance.Content.Update(context.Background(), "64290828", payload)
	if err != nil {

		if response != nil {
			log.Println(response.API)
		}

		log.Fatal(err)
	}

	log.Println("Endpoint:",	 response.Endpoint)
	log.Println("Status Code:", response.Code)
	log.Println(content)
}