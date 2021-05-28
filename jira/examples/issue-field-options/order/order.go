package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
	"sort"
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

	var (
		fieldID   = "customfield_10038"
		contextID = 10180
	)

	log.Println("Getting the field context options")
	options, response, err := atlassian.Issue.Field.Context.Option.Gets(context.Background(), fieldID, contextID, nil, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	var (
		optionsAsMap = make(map[string]string)
		optionsAsList []string
	)

	for _, option := range options.Values {
		optionsAsList = append(optionsAsList, option.Value)
		optionsAsMap[option.Value] = option.ID
	}

	log.Println("Sorting the fields")
	sort.Strings(optionsAsList)

	log.Println("Creating the new option ID's payload to order")
	var optionsIDsAsList []string
	for _, option := range optionsAsList {
		optionsIDsAsList = append(optionsIDsAsList, optionsAsMap[option])
	}

	var payload = &jira.OrderFieldOptionPayloadScheme{
		Position:             "First",
		CustomFieldOptionIds: optionsIDsAsList,
	}

	log.Println("Ordering the options")
	response, err = atlassian.Issue.Field.Context.Option.Order(context.Background(), fieldID, contextID, payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Code", response.StatusCode)
			log.Println("HTTP Endpoint Used", response.Endpoint)
			log.Fatal(err)
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
}