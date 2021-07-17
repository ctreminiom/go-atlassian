package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
	"path/filepath"
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


	absolutePath, err := filepath.Abs("jira/mocks/image.png")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Using the path", absolutePath)

	reader, err := os.Open(absolutePath)
	if err != nil {
		log.Fatal(err)
	}

	defer reader.Close()

	var (
		issueKeyOrID = "KP-2"
		fileName = "image-mock-00.png"
	)

	attachments, response, err := atlassian.Issue.Attachment.Add(context.Background(), issueKeyOrID, fileName, reader)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Printf("We've found %v attachments", len(attachments))

	for _, attachment := range attachments {
		log.Println(attachment.ID, attachment.Filename)
	}
}
