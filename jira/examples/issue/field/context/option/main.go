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

	var payload = jira.CustomFieldOptionPayloadScheme{Options: valuesToAdd}
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

	var payload = jira.CustomFieldOptionPayloadScheme{Options: valuesToAdd}
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

func updateCustomFieldOption() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var optionsToUpdate []jira.FieldContextOptionValueScheme

	Option1 := jira.FieldContextOptionValueScheme{
		ID:       "10058",
		Value:    "Scranton 1",
		Disabled: false,
	}

	Option2 := jira.FieldContextOptionValueScheme{
		ID:       "10059",
		Disabled: true,
	}

	optionsToUpdate = append(optionsToUpdate, Option1, Option2)

	var payload = jira.CustomFieldOptionPayloadScheme{Options: optionsToUpdate}
	var fieldID = "customfield_10047"
	var contextID = 10175

	contextOptions, response, err := atlassian.Issue.Field.Context.Option.Update(context.Background(), fieldID, contextID, &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, option := range contextOptions.Options {
		log.Println(option)
	}

}

func deleteCustomFieldOption() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var fieldID = "customfield_10047"
	var contextID = 10175
	var optionID = 10061

	response, err := atlassian.Issue.Field.Context.Option.Delete(context.Background(), fieldID, contextID, optionID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
}

func main() {
	deleteCustomFieldOption()
}
