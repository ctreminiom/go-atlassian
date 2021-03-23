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

	var (
		issueKeyOrID = "DESK-12"
		commentID    = 10012
		expands      = []string{"attachment", "renderedBody"}
	)

	comment, response, err := atlassian.ServiceManagement.Request.Comment.Get(context.Background(), issueKeyOrID, commentID, expands)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println("----------------------------------")
	log.Printf("Comment, ID: %v", comment.ID)
	log.Printf("Comment, Creator Name: %v", comment.Author.DisplayName)
	log.Printf("Comment, Created Date: %v", comment.Created.Friendly)
	log.Printf("Comment, # of attachments: %v", comment.Attachments.Size)
	log.Printf("Comment, is public?: %v", comment.Public)
	log.Println("----------------------------------")

}
