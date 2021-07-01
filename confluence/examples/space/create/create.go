package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"net/http"
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

	var payload = &confluence.CreateSpaceScheme{
		Key:              "DUM",
		Name:             "Dum Confluence Space",
		Description:      &confluence.CreateSpaceDescriptionScheme{
			Plain: &confluence.CreateSpaceDescriptionPlainScheme{
				Value:          "Confluence Space Description Sample",
				Representation: "plain",
			},
		},
		AnonymousAccess:  true,
		UnlicensedAccess: false,
	}

	space, response, err := instance.Space.Create(context.Background(), payload, false)
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
	log.Println(space)
}
