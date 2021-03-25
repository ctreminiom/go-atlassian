package main

import (
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
		return
	}

	attachments, response, err := atlassian.Issue.Attachment.Add("KP-1", absolutePath)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Printf("We've found %v attachments", len(*attachments))

	for _, attachment := range *attachments {
		log.Println(attachment.ID, attachment.Filename)
	}
}
