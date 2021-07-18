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

	payload := &jira.PermissionSchemeScheme{
		Name:        "EF Permission Scheme",
		Description: "EF Permission Scheme description",

		Permissions: []*jira.PermissionGrantScheme{
			{
				Permission: "ADMINISTER_PROJECTS",
				Holder: &jira.PermissionGrantHolderScheme{
					Parameter: "jira-administrators-system",
					Type:      "group",
				},
			},
			{
				Permission: "CLOSE_ISSUES",
				Holder: &jira.PermissionGrantHolderScheme{
					Type: "assignee",
				},
			},
		},
	}

	permissionScheme, response, err := atlassian.Permission.Scheme.Create(context.Background(), payload)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println(permissionScheme)
}
