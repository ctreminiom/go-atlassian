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
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var payload = v3.IssueScheme{
		Fields: &v3.IssueFieldsScheme{
			Summary:   "New summary test",
			Project:   &v3.ProjectScheme{ID: "10000"},
			IssueType: &v3.IssueTypeScheme{Name: "Story"},
		},
	}

	//CustomFields
	var customFields = v3.CustomFields{}
	err = customFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		log.Fatal(err)
	}

	err = customFields.Number("customfield_10043", 1000.3232)
	if err != nil {
		log.Fatal(err)
	}

	newIssue, response, err := atlassian.Issue.Create(context.Background(), &payload, &customFields)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Printf("The new issue %v has been created with the ID %v", newIssue.Key, newIssue.ID)
}
