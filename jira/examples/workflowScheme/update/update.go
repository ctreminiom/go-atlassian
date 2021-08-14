package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
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

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var (
		payload = &jira.WorkflowSchemePayloadScheme{
			DefaultWorkflow: "jira",
			Name:            "Example workflow scheme - UPDATED",
			Description:     "The description of the example workflow scheme. - UPDATED"}
		workflowSchemeID = 10006
	)

	workflowScheme, response, err := atlassian.Workflow.Scheme.Update(context.Background(), workflowSchemeID, payload)
	if err != nil {
		if response != nil {
			log.Println(response.Endpoint)
			log.Println(response.Code)
			log.Println(response.Bytes.String())
		}
		return
	}

	log.Println("------------")
	log.Printf("Workflow Scheme ID: %v", workflowScheme.ID)
	log.Printf("Workflow Scheme Name: %v", workflowScheme.Name)
	log.Printf("Workflow Scheme Description: %v", workflowScheme.Description)
	log.Printf("Workflow Scheme DefaultWorkflow: %v", workflowScheme.DefaultWorkflow)
	log.Printf("Workflow Scheme Draft: %v", workflowScheme.Draft)
	log.Println("------------")

}
