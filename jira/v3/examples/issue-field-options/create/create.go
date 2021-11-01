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

	var (
		fieldID   = "customfield_10038"
		contextID = 10180

		payload = &v3.FieldContextOptionListScheme{
			Options: []*v3.CustomFieldContextOptionScheme{

				// Single/Multiple Choice example
				{
					Value:    "Option 3",
					Disabled: false,
				},
				{
					Value:    "Option 4",
					Disabled: false,
				},

				///////////////////////////////////////////
				/*
					// Cascading Choice example
					{
						OptionID: "1027",
						Value:    "Argentina",
						Disabled: false,
					},
					{
						OptionID: "1027",
						Value:    "Uruguay",
						Disabled: false,
					},
				*/

			}}
	)

	fieldOptions, response, err := atlassian.Issue.Field.Context.Option.Create(context.Background(), fieldID, contextID, payload)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, option := range fieldOptions.Options {
		log.Println(option)
	}

}
