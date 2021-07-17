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
		Name:        "EF Permission Scheme - UPDATED",
		Description: "EF Permission Scheme description - UPDATED",

		Permissions: []*jira.PermissionGrantScheme{
			{
				Permission: "CLOSE_ISSUES",
				Holder: &jira.PermissionGrantHolderScheme{
					Parameter: "jira-administrators-system",
					Type:      "group",
				},
			},
		},
	}

	permissionScheme, response, err := atlassian.Permission.Scheme.Update(context.Background(), 10004, payload)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println(permissionScheme.Name)
	log.Println(permissionScheme.ID)
	log.Println(permissionScheme.Description)
	log.Println(permissionScheme.Self)

	for _, permissionGrant := range permissionScheme.Permissions {
		log.Println(permissionGrant.ID, permissionGrant.Permission)
	}
}
