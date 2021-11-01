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

	var userRecipients []*v3.IssueNotifyUserScheme
	userRecipients = append(userRecipients, &v3.IssueNotifyUserScheme{AccountID: "87dde939-73be-465f-83c5-2217fb9dd9de"})
	userRecipients = append(userRecipients, &v3.IssueNotifyUserScheme{AccountID: "8abc0d5f-5eb9-48af-bd8d-b83451828a40"})

	var groupsRecipients []*v3.IssueNotifyGroupScheme
	groupsRecipients = append(groupsRecipients, &v3.IssueNotifyGroupScheme{Name: "jira-users"})
	groupsRecipients = append(groupsRecipients, &v3.IssueNotifyGroupScheme{Name: "scrum-masters"})

	opts := &v3.IssueNotifyOptionsScheme{

		// The HTML body of the email notification for the issue.
		HTMLBody: "The <strong>latest</strong> test results for this ticket are now available.",

		// The subject of the email notification for the issue.
		// If this is not specified, then the subject is set to the issue key and summary.
		Subject: "SUBJECT EMAIL EXAMPLE",

		// The plain text body of the email notification for the issue.
		// TextBody: "lorem",

		// The recipients of the email notification for the issue.
		To: &v3.IssueNotifyToScheme{
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
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
}
