package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
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

	payload := jira.LinkTypeScheme{
		Inward:  "Clone/Duplicated by",
		Name:    "Clone/Duplicate",
		Outward: "Clone/Duplicates",
	}

	issueLinkType, response, err := atlassian.Issue.Link.Type.Create(context.Background(), &payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(issueLinkType)
}
