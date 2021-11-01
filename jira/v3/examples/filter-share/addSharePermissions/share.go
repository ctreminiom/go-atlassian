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

	/*
		We can add different share permissions, for example:

		---- Project ID only
		payload := jira.PermissionFilterBodyScheme{
				Type:      "project",
				ProjectID: "10000",
			}

		---- Project ID and role ID
		payload := jira.PermissionFilterBodyScheme{
				Type:          "project",
				ProjectID:     "10000",
				ProjectRoleID: "222222",
			}

		==== Group Name
		payload := jira.PermissionFilterBodyScheme{
				Type:          "group",
				GroupName: "jira-users",
			}
	*/

	payload := v3.PermissionFilterPayloadScheme{
		Type:      "project",
		ProjectID: "10000",
	}

	permissions, response, err := atlassian.Filter.Share.Add(context.Background(), 1001, &payload)
	if err != nil {
		return
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	for index, permission := range permissions {
		log.Println(index, permission.ID, permission.Type)
	}
}
