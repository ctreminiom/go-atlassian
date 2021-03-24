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
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	commentBody := jira.CommentNodeScheme{}
	commentBody.Version = 1
	commentBody.Type = "doc"

	commentBody.AppendNode(&jira.CommentNodeScheme{
		Type: "paragraph",
		Content: []*jira.CommentNodeScheme{
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
	})

	newComment, response, err := atlassian.Issue.Comment.Add(context.Background(), "KP-2", "role", "Administrators", &commentBody, nil)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println(newComment.ID)
}
