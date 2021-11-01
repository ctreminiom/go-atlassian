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

	var payload = &v3.IssueScheme{
		Fields: &v3.IssueFieldsScheme{
			Summary: "New summary test test",
		},
	}

	//CustomFields
	var customFields = &v3.CustomFields{}
	err = customFields.Groups("customfield_10052", []string{"jira-administrators", "jira-administrators-system"})
	if err != nil {
		log.Fatal(err)
	}

	//Issue Update Operations
	var operations = &v3.UpdateOperations{}

	err = operations.AddArrayOperation("labels", map[string]string{
		"triaged":   "add",
		"triaged-2": "add",
		"triaged-1": "add",
		"blocker":   "add",
	})

	if err != nil {
		log.Fatal(err)
	}

	options := &v3.IssueMoveOptions{
		Fields:       payload,
		Operations:   operations,
		CustomFields: customFields,
	}

	response, err := atlassian.Issue.Move(context.Background(), "KP-7", "41", options)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
}
