package main

import (
	"context"
	"fmt"
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

	options := &v3.IssueMetadataCreateOptions{
		ProjectIDs:     nil,
		ProjectKeys:    []string{"K2"},
		IssueTypeIDs:   nil,
		IssueTypeNames: nil,
		Expand:         "projects.issuetypes.fields",
	}

	metadata, response, err := atlassian.Issue.Metadata.Create(context.Background(), options)
	if err != nil {

		if response != nil {
			log.Println(response.Bytes.String())
		}

		log.Fatal(err)
	}

	for _, project := range metadata.Get("projects").Array() {

		fmt.Println("Project: ", project.Get("name"))
		for _, issueType := range project.Get("issuetypes").Array() {

			fmt.Println("Issue Type: ", issueType.Get("name"))
			for _, fields := range issueType.Get("fields").Array() {
				fmt.Println(fields.Raw)
			}
		}
	}
}
