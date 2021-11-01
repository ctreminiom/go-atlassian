package main

import (
	"context"
	"fmt"
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
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var (
		serviceDeskID = 1
		queueID       = 1
		start, limit  = 0, 50
	)

	issues, response, err := atlassian.ServiceManagement.ServiceDesk.Queue.Issues(context.Background(), serviceDeskID, queueID, start, limit)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	for _, issue := range issues.Values {
		fmt.Println(issue)
	}

}
