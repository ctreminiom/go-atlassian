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

	searchOptions := jira.DashboardSearchOptionsScheme{
		DashboardName:       "Bug",
		GroupPermissionName: "administrators",
		OrderBy:             "description",
		Expand:              []string{"description", "favourite", "sharePermissions"},
	}

	dashboards, response, err := jiraCloud.Dashboard.Search(context.Background(), &searchOptions, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, dashboard := range dashboards.Values {
		log.Printf("Dashboard Name: %v", dashboard.Name)
		log.Printf("Dashboard ID: %v", dashboard.ID)
		log.Printf("Dashboard View: %v", dashboard.View)
	}

}
