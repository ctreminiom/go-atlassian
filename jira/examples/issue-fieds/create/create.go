package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
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

	fieldNewCreate := jira.CustomFieldScheme{
		Name:        "Alliance 2",
		Description: "this is the alliance description field",
		FieldType:   "cascadingselect",
		SearcherKey: "cascadingselectsearcher",
	}

	field, response, err := atlassian.Issue.Field.Create(context.Background(), &fieldNewCreate)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("field", field)

}
