package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main()  {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	instance, err := confluence.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	instance.Auth.SetBasicAuth(mail, token)
	instance.Auth.SetUserAgent("curl/7.54.0")

	absolutePath, err := filepath.Abs("confluence/mocks/mock.png")
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
		attachmentID = "76513281"
		fileName = "mock-00.png"
	)

	attachmentsPage, response, err := instance.Content.Attachment.Create(context.Background(), attachmentID, "", fileName, reader)
	if err != nil {

		if response != nil {
			if response.Code == http.StatusBadRequest {
				log.Println(response.API)
			}
		}
		log.Println(response.Endpoint)
		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)

	for _, attachment := range attachmentsPage.Results {
		log.Println(attachment.ID, attachment.Title)
	}

}
