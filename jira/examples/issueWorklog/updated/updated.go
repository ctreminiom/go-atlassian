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

	result, response, err := atlassian.Issue.Worklog.Updated(context.Background(), 0, nil)
	if err != nil {
		log.Println(response.Endpoint, response.Code)
		log.Fatal(err)
	}

	log.Println(response.Endpoint, response.Code)
	log.Println(result)
}
