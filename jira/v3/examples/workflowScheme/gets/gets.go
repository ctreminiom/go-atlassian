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

	workflowsSchemes, response, err := atlassian.Workflow.Scheme.Gets(context.Background(), 0, 50)
	if err != nil {
		if response != nil {
			log.Println(response.Endpoint)
			log.Println(response.Code)
			log.Println(response.Bytes.String())
		}
	}

	for _, workflowScheme := range workflowsSchemes.Values {

		log.Println("------------")
		log.Printf("Workflow Scheme ID: %v", workflowScheme.ID)
		log.Printf("Workflow Scheme Name: %v", workflowScheme.Name)
		log.Printf("Workflow Scheme Description: %v", workflowScheme.Description)
		log.Printf("Workflow Scheme DefaultWorkflow: %v", workflowScheme.DefaultWorkflow)
		log.Printf("Workflow Scheme Draft: %v", workflowScheme.Draft)
		log.Println("------------")
	}

}
