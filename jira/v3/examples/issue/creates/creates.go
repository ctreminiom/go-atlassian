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

	var issue1 = v3.IssueBulkScheme{
		Payload: &v3.IssueScheme{
			Fields: &v3.IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &v3.ProjectScheme{ID: "10000"},
				IssueType: &v3.IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: &customFields,
	}

	var issue2 = v3.IssueBulkScheme{
		Payload: &v3.IssueScheme{
			Fields: &v3.IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &v3.ProjectScheme{ID: "10000"},
				IssueType: &v3.IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: &customFields,
	}

	var issue3 = v3.IssueBulkScheme{
		Payload: &v3.IssueScheme{
			Fields: &v3.IssueFieldsScheme{
				Summary:   "New summary test",
				Project:   &v3.ProjectScheme{ID: "10000"},
				IssueType: &v3.IssueTypeScheme{Name: "Story"},
			},
		},
		CustomFields: &customFields,
	}

	var payload []*v3.IssueBulkScheme
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
