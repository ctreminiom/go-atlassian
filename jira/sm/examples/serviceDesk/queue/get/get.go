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
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var (
		serviceDeskID = 1
		queueID       = 1
		includeCount  = true
	)

	queue, response, err := atlassian.ServiceManagement.ServiceDesk.Queue.Get(context.Background(), serviceDeskID, queueID, includeCount)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	log.Println("------------------------------------")
	log.Printf("Queue ID:  %v", queue.ID)
	log.Printf("Queue Name: %v", queue.Name)
	log.Printf("Queue JQL: %v", queue.Jql)
	log.Printf("Queue Issue Count: %v", queue.IssueCount)
	log.Printf("Queue Fields: %v", queue.Fields)
	log.Println("------------------------------------")

}
