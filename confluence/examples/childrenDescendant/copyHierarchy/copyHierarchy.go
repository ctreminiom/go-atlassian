package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"os"
)

func main() {

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

	var (
		contentID = "76513281"
		options   = &confluence.CopyOptionsScheme{
			CopyAttachments:    true,
			CopyPermissions:    true,
			CopyProperties:     true,
			CopyLabels:         true,
			CopyCustomContents: true,
			DestinationPageID:  "80412692",
			TitleOptions: &confluence.CopyTitleOptionScheme{
				Prefix: "copy-01-",
			},
		}
	)

	task, response, err := instance.Content.ChildrenDescendant.CopyHierarchy(context.Background(), contentID, options)
	if err != nil {

		if response != nil {
			log.Println(response.API)
		}

		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)

	log.Println(task.ID)
	log.Println(task.Links.Status)

}
