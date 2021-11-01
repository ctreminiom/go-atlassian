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
		boardID    = 4
		startAt    = 0
		maxResults = 50
		released   = false
	)

	versionsPage, response, err := atlassian.Agile.Board.Versions(context.Background(), boardID, startAt, maxResults, released)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, version := range versionsPage.Values {
		log.Println(version)
	}

	fmt.Println(response.Bytes.String())
}
