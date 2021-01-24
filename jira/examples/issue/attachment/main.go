package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
	"path/filepath"
)

/*
----------- Set an environment variable in git bash -----------
export HOST="https://ctreminiom.atlassian.net/"
export MAIL="MAIL_ADDRESS"
export TOKEN="TOKEN_API"

Docs: https://stackoverflow.com/questions/34169721/set-an-environment-variable-in-git-bash
*/

var (
	host  = os.Getenv("HOST")
	mail  = os.Getenv("MAIL")
	token = os.Getenv("TOKEN")
)

func getAttachmentsSetting() (err error) {
	log.Println("------------- getAttachmentsSetting -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	settings, response, err := atlassian.Issue.Attachment.Settings(context.Background())
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("settings", settings.Enabled, settings.UploadLimit)

	return
}

func getAttachmentMetadata() (err error) {
	log.Println("------------- getAttachmentMetadata -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	metadata, response, err := atlassian.Issue.Attachment.Metadata(context.Background(), "10000")
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(metadata)
	return
}

func getAttachmentHuman() (err error) {
	log.Println("------------- getAttachmentHuman -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	humanMetadata, response, err := atlassian.Issue.Attachment.Human(context.Background(), "10000")
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(humanMetadata)

	return
}

func addAttachment() (id string, err error) {
	log.Println("------------- addAttachment -----------------")

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
		id = attachment.ID
	}

	return
}

func deleteAttachment(attachmentID string) (err error) {

	log.Println("------------- deleteAttachment -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Issue.Attachment.Delete(context.Background(), attachmentID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	return
}

func main() {

	var err error
	if err = getAttachmentsSetting(); err != nil {
		log.Fatal(err)
	}

	if err = getAttachmentMetadata(); err != nil {
		log.Fatal(err)
	}

	if err = getAttachmentHuman(); err != nil {
		log.Fatal(err)
	}

	log.Println("Creating the new attachment on the Jira Cloud instance")
	newAttachmentID, err := addAttachment()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The new attachment has been added on your Jira Cloud instance, ID: %v", newAttachmentID)

	if err = deleteAttachment(newAttachmentID); err != nil {
		log.Fatal(err)
	}
	log.Printf("The attachment has been removed")

}
