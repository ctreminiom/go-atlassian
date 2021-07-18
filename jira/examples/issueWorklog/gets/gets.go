package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main()  {

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

	var (
		worklogsIDs = []int{10000}
		expand = []string{"all"}
	)

	worklogs, response, err := atlassian.Issue.Worklog.Gets(context.Background(), worklogsIDs, expand)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.Endpoint, response.Code)

	for _, worklog := range worklogs {
		log.Println(worklog)
	}
}
