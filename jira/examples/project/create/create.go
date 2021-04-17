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

	payload := &jira.ProjectPayloadScheme{
		NotificationScheme:  10021,
		Description:         "Example Project description",
		LeadAccountID:       "5b86be50b8e3cb5895860d6d",
		URL:                 "http://atlassian.com",
		ProjectTemplateKey:  "com.pyxis.greenhopper.jira:gh-simplified-agility-kanban",
		AvatarID:            10200,
		IssueSecurityScheme: 10001,
		Name:                "Project DUMMY #3",
		PermissionScheme:    10011,
		AssigneeType:        "PROJECT_LEAD",
		ProjectTypeKey:      "software",
		Key:                 "DUMMY3",
		CategoryID:          10120,
	}

	newProject, response, err := atlassian.Project.Create(context.Background(), payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println("-------------------")
	log.Println(newProject.ID)
	log.Println(newProject.Self)
	log.Println(newProject.Key)
	log.Println("-------------------")
}
