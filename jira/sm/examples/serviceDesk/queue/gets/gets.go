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
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var (
		serviceDeskID      = 1
		includeCount  bool = true
		start, limit  int  = 0, 50
	)

	queues, response, err := atlassian.ServiceManagement.ServiceDesk.Queue.Gets(context.Background(), serviceDeskID, includeCount, start, limit)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	for pos, queue := range queues.Values {

		log.Println("------------------------------------")
		log.Printf("Queue ID #%v: %v", pos+1, queue.ID)
		log.Printf("Queue Name #%v: %v", pos+1, queue.Name)
		log.Printf("Queue JQL #%v: %v", pos+1, queue.Jql)
		log.Printf("Queue Issue Count #%v: %v", pos+1, queue.IssueCount)
		log.Printf("Queue Fields #%v: %v", pos+1, queue.Fields)
		log.Println("------------------------------------")
	}

}
