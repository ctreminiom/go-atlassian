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

	var fieldID = "customfield_10038"

	var payload = &v3.FieldContextDefaultPayloadScheme{
		DefaultValues: []*v3.CustomFieldDefaultValueScheme{
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
