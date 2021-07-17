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

	payload := &jira.ProjectComponentPayloadScheme{
		IsAssigneeTypeValid: false,
		Name:                "Component 2",
		Description:         "This is a Jira component",
		Project:             "KP",
		AssigneeType:        "PROJECT_LEAD",
		LeadAccountID:       "5b86be50b8e3cb5895860d6d",
	}

	newComponent, response, err := atlassian.Project.Component.Create(context.Background(), payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Printf("The new component has been created with the ID %v", newComponent.ID)
}
