package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

/*
----------- Set an environment variable in git bash -----------
export HOST="https://ctreminiom.atlassian.net/"
export MAIL="MAIL_ADDRESS"
export TOKEN="TOKEN_API"

Docs: https://stackoverflow.com/questions/34169721/set-an-environment-variable-in-git-bash
*/

var (
	host  = os.Getenv("HOST")
	mail  = os.Getenv("MAIL")
	token = os.Getenv("TOKEN")
)

func getFieldOptions() (err error) {
	log.Println("------------- getFieldOptions -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var (
		fieldID   = "customfield_10038"
		contextID = 10138
	)

	options := jira.FieldOptionContextParams{
		FieldID:     fieldID,
		ContextID:   contextID,
		OnlyOptions: false,
	}

	fieldOptions, response, err := atlassian.Issue.Field.Context.Option.Gets(context.Background(), &options, 0, 60)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, fieldOption := range fieldOptions.Values {
		log.Println(fieldOption)
	}

	return
}

func createCustomFieldOptionSelectLisCascading() (err error) {

	log.Println("------------- createCustomFieldOptionSelectLisCascading -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var valuesToAdd []jira.FieldContextOptionValueScheme

	argentinaOption := jira.FieldContextOptionValueScheme{
		OptionID: "10027",
		Value:    "Argentina",
		Disabled: false,
	}

	uruguayOption := jira.FieldContextOptionValueScheme{
		OptionID: "10027",
		Value:    "Uruguay",
		Disabled: false,
	}

	valuesToAdd = append(valuesToAdd, argentinaOption, uruguayOption)

	var payload = jira.CreateCustomFieldOptionPayloadScheme{Options: valuesToAdd}
	var fieldID = "customfield_10037"
	var contextID = 10141

	fieldOptions, response, err := atlassian.Issue.Field.Context.Option.Create(context.Background(), fieldID, contextID, &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, option := range fieldOptions.Options {
		log.Println(option)
	}

	return
}

func createCustomFieldOptionSelectLisSingleChoice() (err error) {

	log.Println("------------- createCustomFieldOptionSelectLisSingleChoice -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var valuesToAdd []jira.FieldContextOptionValueScheme

	value00 := jira.FieldContextOptionValueScheme{
		Value:    "Scranton 1",
		Disabled: false,
	}

	valuesToAdd = append(valuesToAdd, value00)

	var payload = jira.CreateCustomFieldOptionPayloadScheme{Options: valuesToAdd}
	var fieldID = "customfield_10038"
	var contextID = 10140

	fieldOptions, response, err := atlassian.Issue.Field.Context.Option.Create(context.Background(), fieldID, contextID, &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, option := range fieldOptions.Options {
		log.Println(option)
	}
	return
}

func main() {

	var err error

	err = getFieldOptions()
	if err != nil {
		log.Fatal(err)
	}

	err = createCustomFieldOptionSelectLisSingleChoice()
	if err != nil {
		log.Fatal(err)
	}

	err = createCustomFieldOptionSelectLisCascading()
	if err != nil {
		log.Fatal(err)
	}

}
