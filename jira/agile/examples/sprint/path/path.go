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
		Name: "Sprint XX-Patched",
	}

	sprint, response, err := atlassian.Agile.Sprint.Path(context.Background(), 2, payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response",response.Bytes.String())
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(sprint.Name, sprint.Goal)
}
