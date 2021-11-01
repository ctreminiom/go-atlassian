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
		maxResult = 50
		after     = 0
		expand    = []string{""}
	)

	worklogs, response, err := atlassian.Issue.Worklog.Issue(ctx, key, 0, maxResult, after, expand)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.Endpoint, response.Code)

	for _, worklog := range worklogs.Worklogs {
		log.Println(worklog.ID)
	}
}
