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

	var userRecipients []*jira.IssueNotifyUserScheme
	userRecipients = append(userRecipients, &jira.IssueNotifyUserScheme{AccountID: "87dde939-73be-465f-83c5-2217fb9dd9de"})
	userRecipients = append(userRecipients, &jira.IssueNotifyUserScheme{AccountID: "8abc0d5f-5eb9-48af-bd8d-b83451828a40"})

	var groupsRecipients []*jira.IssueNotifyGroupScheme
	groupsRecipients = append(groupsRecipients, &jira.IssueNotifyGroupScheme{Name: "jira-users"})
	groupsRecipients = append(groupsRecipients, &jira.IssueNotifyGroupScheme{Name: "scrum-masters"})

	opts := &jira.IssueNotifyOptionsScheme{

		// The HTML body of the email notification for the issue.
		HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",

		// The subject of the email notification for the issue.
		// If this is not specified, then the subject is set to the issue key and summary.
		Subject: "SUBJECT EMAIL EXAMPLE",

		// The plain text body of the email notification for the issue.
		// TextBody: "lorem",

		// The recipients of the email notification for the issue.
		To: &jira.IssueNotifyToScheme{
			Reporter: true,
			Assignee: true,
			Watchers: true,
			Voters:   true,
			Users:    userRecipients,
			Groups:   groupsRecipients,
		},

		// Restricts the notifications to users with the specified permissions.
		Restrict: nil,
	}

	response, err := atlassian.Issue.Notify(context.Background(), "KP-2", opts)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
			log.Println(response.StatusCode)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
}
