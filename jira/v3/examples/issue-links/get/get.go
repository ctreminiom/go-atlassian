package main

import (
	"context"
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
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	issueLink, response, err := atlassian.Issue.Link.Get(context.Background(), "10002")
	if err != nil {
		log.Fatal(err)
	}

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
