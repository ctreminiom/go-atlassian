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

	var fieldID = "customfield_10038"

	var payload = &jira.FieldContextDefaultPayloadScheme{
		DefaultValues: []*jira.CustomFieldDefaultValueScheme{
			{
				ContextID: "10138",
				OptionID:  "10022",
				Type:      "option.single",
			},
		},
	}

	response, err := atlassian.Issue.Field.Context.SetDefaultValue(context.Background(), fieldID, payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
}
