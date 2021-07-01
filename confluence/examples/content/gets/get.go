package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"net/http"
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

	options := &confluence.GetContentOptionsScheme{
		//ContextType: "page",
		SpaceKey: "",
		Title:    "",
		Trigger:  "",
		OrderBy:  "",
		//Status:      []string{"any", "any"},
		Expand: []string{"childTypes.all", "operations"},
		//PostingDay:  time.Now(),
	}

	page, response, err := instance.Content.Gets(context.Background(), options, 0, 50)
	if err != nil {

		if response.Code == http.StatusBadRequest {
			log.Println(response.API)
		}
		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)

	for _, content := range page.Results {

		if content.ChildTypes != nil {
			log.Println("- Content ChildTypes -")
			log.Println(content.ChildTypes.Attachment.Links.Self)
			log.Println(content.ChildTypes.Comment.Links.Self)
			log.Println(content.ChildTypes.Page.Links.Self)
		}

		if content.Space != nil {
			log.Println("- Space -")
			log.Println(content.Space)
		}

		if content.Metadata != nil  && content.Metadata.Labels != nil {
			log.Println("- Content Labels -")

			for _, result := range content.Metadata.Labels.Results {
				log.Println(result)
			}
		}

		if content.Operations != nil {
			log.Println("- Operations -")
			for _, operation := range content.Operations {
				log.Println(operation)
			}
		}

		log.Println("--------------------------------------")
	}
}
