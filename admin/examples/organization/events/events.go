package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/admin"
	"log"
	"net/url"
	"os"
	"time"
)

func main() {

	//ATLASSIAN_ADMIN_TOKEN
	var apiKey = os.Getenv("ATLASSIAN_ADMIN_TOKEN")

	cloudAdmin, err := admin.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	cloudAdmin.Auth.SetAccessToken(apiKey)
	cloudAdmin.Auth.SetUserAgent("curl/7.54.0")

	var (
		organizationID = "9a1jj823-jac8-123d-jj01-63315k059cb2"
		cursor         string
		eventChunks    []*admin.OrganizationEventPageScheme
	)

	for {

		opts := &admin.OrganizationEventOptScheme{
			Q:      "",
			From:   time.Now().Add(time.Duration(-24) * time.Hour),
			To:     time.Time{},
			Action: "",
		}

		events, response, err := cloudAdmin.Organization.Events(context.Background(), organizationID, opts, cursor)
		if err != nil {
			if response != nil {
				log.Println("Response HTTP Response", string(response.BodyAsBytes))
			}
			log.Fatal(err)
		}

		log.Println("Response HTTP Code", response.StatusCode)
		log.Println("HTTP Endpoint Used", response.Endpoint)
		eventChunks = append(eventChunks, events)

		if len(events.Links.Next) == 0 {
			break
		}

		//extract the next cursor pagination
		nextAsURL, err := url.Parse(events.Links.Next)
		if err != nil {
			log.Fatal(err)
		}

		cursor = nextAsURL.Query().Get("cursor")
	}

	for _, chunk := range eventChunks {

		for _, event := range chunk.Data {
			log.Println(event.ID, event.Attributes.Action, event.Attributes.Time)
		}

	}

}
