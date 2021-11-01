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

	var payload = &v3.WorkflowSchemePayloadScheme{
		DefaultWorkflow: "jira",
		Name:            "Example workflow scheme",
		Description:     "The description of the example workflow scheme.",
		IssueTypeMappings: map[string]string{
			"10000": "scrum workflow",
			"10001": "builds workflow",
		},
	}

	workflowScheme, response, err := atlassian.Workflow.Scheme.Create(context.Background(), payload)
	if err != nil {
		if response != nil {
			log.Println(response.Endpoint)
			log.Println(response.Code)
			log.Println(response.Bytes.String())
		}
	}

	log.Println(workflowScheme)
}
