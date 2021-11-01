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

	payload := &v3.LinkPayloadScheme{

		Comment: &v3.CommentPayloadScheme{

			Body: &v3.CommentNodeScheme{
				Version: 1,
				Type:    "doc",
				Content: []*v3.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*v3.CommentNodeScheme{
							{
								Type: "text",
								Text: "Carlos Test",
							},
							{
								Type: "emoji",
								Attrs: map[string]interface{}{
									"shortName": ":grin",
									"id":        "1f601",
									"text":      "😁",
								},
							},
							{
								Type: "text",
								Text: " ",
							},
						},
					},
				},
			},
		},

		InwardIssue: &v3.LinkedIssueScheme{
			Key: "KP-1",
		},
		OutwardIssue: &v3.LinkedIssueScheme{
			Key: "KP-2",
		},
		Type: &v3.LinkTypeScheme{
			Name: "Duplicate",
		},
	}

	response, err := atlassian.Issue.Link.Create(context.Background(), payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
}
