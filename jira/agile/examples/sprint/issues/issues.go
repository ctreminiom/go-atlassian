package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/ctreminiom/go-atlassian/jira/agile"
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

	options := &agile.IssueOptionScheme{
		JQL:           "",
		Fields:        nil,
		Expand:        nil,
		ValidateQuery: false,
	}

	issues, response, err := atlassian.Agile.Sprint.Issues(context.Background(), 2, options, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println(issues.Total)
	log.Println(issues.Expand)
	log.Println(issues.MaxResults)
	log.Println(issues.StartAt)

	for _, issue := range issues.Issues {
		log.Println(issue.Key, issue.ID)
	}

	//If you want to extract the field metadata
	var searchPage jira.IssueSearchScheme
	if err = json.Unmarshal(response.Bytes.Bytes(), &searchPage); err != nil {
		log.Fatal(err)
	}

	for _, issue := range searchPage.Issues {
		log.Println(issue.Key, issue.Fields.Summary, issue.Fields.IssueType.Name)
	}

	fmt.Println(response.Bytes.String())

}
