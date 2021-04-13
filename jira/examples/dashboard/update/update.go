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

	var sharePermissions = []jira.SharePermissionScheme{
		{
			Type: "project",
			Project: &jira.ProjectScheme{
				ID: "10000",
			},
			Role:  nil,
			Group: nil,
		},
		{
			Type:  "group",
			Group: &jira.GroupScheme{Name: "jira-administrators"},
		},
	}

	dashboard, response, err := jiraCloud.Dashboard.Update(context.Background(), "10001", "Team Tracking #1111", "", &sharePermissions)
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
