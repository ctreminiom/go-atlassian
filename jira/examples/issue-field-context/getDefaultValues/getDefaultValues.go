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

	defaultValues, response, err := atlassian.Issue.Field.Context.GetDefaultValues(context.Background(), fieldID, nil, 0, 50)
	if err != nil {
		return
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, value := range defaultValues.Values {

		/*
			For singleOption customField type, use value.OptionID
			For multipleOption customField type, use value.OptionIDs
			For cascadingOption customField type, use value.OptionID and value.CascadingOptionID
		*/
		log.Println(value)
	}

}
