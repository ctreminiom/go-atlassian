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
		expand    = []string{"page", "comments", "attachment"}
	)

	contents, response, err := instance.Content.ChildrenDescendant.Children(context.Background(), contentID, expand, 0)
	if err != nil {

		if response != nil {
			log.Println(response.API)
		}

		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)

	if contents.Attachment != nil {
		for _, attachment := range contents.Attachment.Results {
			log.Println(attachment.Type, attachment.Title)
		}
	}

	if contents.Comments != nil {
		for _, comment := range contents.Comments.Results {
			log.Println(comment.Type, comment.Title)
		}
	}

	if contents.Page != nil {
		for _, page := range contents.Page.Results {
			log.Println(page.Type, page.Title)
		}
	}

}
