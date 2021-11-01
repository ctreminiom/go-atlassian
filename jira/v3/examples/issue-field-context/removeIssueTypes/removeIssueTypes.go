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

	var (
		customFieldID = "customfield_10038"
		contextID     = 10180
		issueTypesIDs = []string{"10007", "10002"}
	)

	response, err := atlassian.Issue.Field.Context.RemoveIssueTypes(context.Background(), customFieldID, contextID, issueTypesIDs)
	if err != nil {
		return
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
}
