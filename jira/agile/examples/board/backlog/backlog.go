package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ctreminiom/go-atlassian/jira/agile"
	"github.com/ctreminiom/go-atlassian/jira/v3"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
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
		options = &agile.IssueOptionScheme{
			JQL:           "project = KP",
			ValidateQuery: true,
			Fields:        []string{"status", "issuetype", "summary"},
			Expand:        []string{"changelog"},
		}

		boardID   = 4
		startAt   = 0
		maxResult = 50
	)

	issues, response, err := atlassian.Agile.Board.Backlog(context.Background(), boardID, startAt, maxResult, options)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.Bytes.Bytes()))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, issue := range issues.Issues {
		log.Println(issue.Key)
	}

	//If you want to extract the field metadata
	var searchPage models.IssueSearchScheme
	if err = json.Unmarshal(response.Bytes.Bytes(), &searchPage); err != nil {
		log.Fatal(err)
	}

	for _, issue := range searchPage.Issues {
		log.Println(issue.Key, issue.Fields.Summary, issue.Fields.IssueType.Name)
	}

	fmt.Println(string(response.Bytes.Bytes()))
}
