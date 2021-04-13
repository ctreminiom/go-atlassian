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
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	issueLink, response, err := atlassian.Issue.Link.Get(context.Background(), "10002")
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes), response.StatusCode)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println(issueLink.ID)
	log.Println("----------------")
	log.Println(issueLink.Type.Name)
	log.Println(issueLink.Type.ID)
	log.Println(issueLink.Type.Self)
	log.Println(issueLink.Type.Inward)
	log.Println(issueLink.Type.Outward)
	log.Println("----------------")
	log.Println(issueLink.InwardIssue)
	log.Println(issueLink.OutwardIssue)
}
