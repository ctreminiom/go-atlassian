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
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var (
		serviceDeskProjectID = 1
		filePath             = "jira/sm/mocks/image.png"
	)

	absolutePath, err := filepath.Abs(filePath)
	if err != nil {
		return
	}

	attachments, response, err := atlassian.ServiceManagement.ServiceDesk.Attach(context.Background(), serviceDeskProjectID, absolutePath)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, attachment := range attachments.TemporaryAttachments {
		log.Println(attachment.FileName, attachment.TemporaryAttachmentID)
	}

}
