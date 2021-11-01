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

	response, err := atlassian.Issue.Worklog.Delete(context.Background(), "KP-1", "10000", nil)
	if err != nil {
		log.Println(response.Endpoint, response.Code)
		log.Fatal(err)
	}

	log.Println(response.Endpoint, response.Code)
}
