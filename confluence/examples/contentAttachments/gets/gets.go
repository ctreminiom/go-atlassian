package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"net/http"
	"os"
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

	var options = &confluence.GetContentAttachmentsOptionsScheme{
		Expand:    nil,
		FileName:  "",
		MediaType: "",
	}

	attachments, response, err := instance.Content.Attachment.Gets(context.Background(), "76513281", 0, 50, options)
	if err != nil {

		if response != nil {
			if response.Code == http.StatusBadRequest {
				log.Println(response.API)
			}
		}
		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)

	for _, attachment := range attachments.Results {
		log.Println(attachment.Metadata.MediaType, attachment.Title)
	}
}
