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

	comment, response, err := atlassian.Issue.Comment.Get(context.Background(), "KP-2", "10011")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response.Endpoint)
	log.Println(comment.ID, comment.Created)
}
