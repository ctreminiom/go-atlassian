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

	var workflowID = "47e1994b-b84a-476c-84e1-2c8b0e5bb181"

	response, err := atlassian.Workflow.Delete(context.Background(), workflowID)
	if err != nil {
		if response != nil {
			log.Println(response.Endpoint)
			log.Println(response.Code)
			log.Println(response.Bytes.String())
		}
	}

	log.Println(response.Code)
	log.Println(response.Endpoint)
}
