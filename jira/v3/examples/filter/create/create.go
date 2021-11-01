package main

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/jira/v3"
	"github.com/google/uuid"
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

	newFilterBody := v3.FilterPayloadScheme{
		Name:        fmt.Sprintf("Filter #%v", uuid.New().String()),
		Description: "Filter's description",
		JQL:         "issuetype = Bug",
		Favorite:    false,
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

	filter, response, err := atlassian.Filter.Create(context.Background(), &newFilterBody)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Printf("The filter has been created: %v - %v", filter.ID, filter.Name)

}
