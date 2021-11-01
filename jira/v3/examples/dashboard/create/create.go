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

	jiraCloud, err := v3.New(nil, host)
	if err != nil {
		return
	}

	jiraCloud.Auth.SetBasicAuth(mail, token)
	jiraCloud.Auth.SetUserAgent("curl/7.54.0")

	var payload = &v3.DashboardPayloadScheme{
		Name:        "Team Tracking 3",
		Description: "description sample",
		SharePermissions: []*v3.SharePermissionScheme{
			{
				Type: "project",
				Project: &v3.ProjectScheme{
					ID: "10000",
				},
				Role:  nil,
				Group: nil,
			},
			{
				Type:  "group",
				Group: &v3.GroupScheme{Name: "jira-administrators"},
			},
		},
	}

	dashboard, response, err := jiraCloud.Dashboard.Create(context.Background(), payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Printf("Dashboard Name: %v", dashboard.Name)
	log.Printf("Dashboard ID: %v", dashboard.ID)
	log.Printf("Dashboard View: %v", dashboard.View)
}
