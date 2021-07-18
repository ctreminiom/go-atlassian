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

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	//CustomFields
	var customFields = jira.CustomFields{}
	err = customFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		log.Fatal(err)
	}

	err = customFields.Number("customfield_10043", 1000.3232)
	if err != nil {
		log.Fatal(err)
	}

	var issue1 = jira.IssueBulkScheme{
		Payload: &jira.IssueScheme{
			Fields: &jira.IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &jira.ProjectScheme{ID: "10000"},
				IssueType: &jira.IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: &customFields,
	}

	var issue2 = jira.IssueBulkScheme{
		Payload: &jira.IssueScheme{
			Fields: &jira.IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &jira.ProjectScheme{ID: "10000"},
				IssueType: &jira.IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: &customFields,
	}

	var issue3 = jira.IssueBulkScheme{
		Payload: &jira.IssueScheme{
			Fields: &jira.IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &jira.ProjectScheme{ID: "10000"},
				IssueType: &jira.IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: &customFields,
	}

	var payload []*jira.IssueBulkScheme
	payload = append(payload, &issue1, &issue2, &issue3)

	newIssues, response, err := atlassian.Issue.Creates(context.Background(), payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, issue := range newIssues.Issues {
		log.Printf(issue.Key)
	}

	for _, apiError := range newIssues.Errors {
		log.Println(apiError.Status, apiError.Status)
	}
}
