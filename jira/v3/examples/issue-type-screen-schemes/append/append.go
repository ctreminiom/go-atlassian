package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira/v3"
	"log"
	"os"
)

func main() {

	/*
		----------- Set an environment variable in git bash -----------
		export HOST="https://ctreminiom.atlassian.net/"
		export MAIL="MAIL_ADDRESS"
		export TOKEN="TOKEN_API"

		Docs: https://stackoverflow.com/questions/34169721/set-an-environment-variable-in-git-bash
	*/

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

	var payload = &v3.IssueTypeScreenSchemePayloadScheme{
		IssueTypeMappings: []*v3.IssueTypeScreenSchemeMappingPayloadScheme{
			{
				IssueTypeID:    "10000", // Epic Issue Type
				ScreenSchemeID: "10002",
			},
			{
				IssueTypeID:    "10002", // Task Issue Type
				ScreenSchemeID: "10002",
			},
		},
	}

	var issueTypeScreenSchemeID = "10005"

	response, err := atlassian.Issue.Type.ScreenScheme.Append(context.Background(), issueTypeScreenSchemeID, payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
}
