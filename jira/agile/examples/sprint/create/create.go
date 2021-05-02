package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/ctreminiom/go-atlassian/jira/agile"
	"log"
	"os"
)

func main() {
	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
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
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(sprint)
}
