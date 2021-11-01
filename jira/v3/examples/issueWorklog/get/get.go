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

	var (
		ctx       = context.Background()
		key       = "KP-1"
		worklogID = "10000"
		expand    = []string{"all"}
	)

	worklog, response, err := atlassian.Issue.Worklog.Get(ctx, key, worklogID, expand)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.Endpoint, response.Code)
	log.Println(worklog.ID, worklog.Self)
}
