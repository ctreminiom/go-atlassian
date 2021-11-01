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

	settings, response, err := atlassian.Issue.Attachment.Settings(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("settings", settings.Enabled, settings.UploadLimit)
}
