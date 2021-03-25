package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	jiraCloud, err := jira.New(nil, host)
	if err != nil {
		return
	}

	jiraCloud.Auth.SetBasicAuth(mail, token)
	jiraCloud.Auth.SetUserAgent("curl/7.54.0")

	var sharePermissions []jira.SharePermissionScheme

	projectPermission := &jira.SharePermissionScheme{
		Type: "project",
		Project: &jira.SharePermissionProjectScheme{
			ID: "10000",
		},
	}

	groupPermission := &jira.SharePermissionScheme{
		Type:  "group",
		Group: &jira.SharePermissionGroupScheme{Name: "jira-administrators"},
	}

	sharePermissions = append(sharePermissions, *projectPermission, *groupPermission)

	dashboard, response, err := jiraCloud.Dashboard.Create(context.Background(), "Team Tracking 1", "", &sharePermissions)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Printf("Dashboard Name: %v", dashboard.Name)
	log.Printf("Dashboard ID: %v", dashboard.ID)
	log.Printf("Dashboard View: %v", dashboard.View)
}
