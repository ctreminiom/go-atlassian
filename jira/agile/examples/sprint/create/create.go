package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira/agile"
	"github.com/ctreminiom/go-atlassian/jira/v3"
	"log"
	"os"
)

func main() {
	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := v3.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	payload := &agile.SprintPayloadScheme{
		Name:          "Sprint XX",
		StartDate:     "2015-04-11T15:22:00.000+10:00",
		EndDate:       "2015-04-20T01:22:00.000+10:00",
		OriginBoardID: 4,
		Goal:          "Sprint XX goal",
	}

	sprint, response, err := atlassian.Agile.Sprint.Create(context.Background(), payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Bytes.String())
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(sprint)
}
