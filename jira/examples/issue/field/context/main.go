package main

import (
	"context"
	"fmt"
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

func getFieldContexts() (err error) {
	log.Println("------------- getFieldContexts -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var fieldID = "customfield_10038"

	options := jira.FieldContextOptionsScheme{
		IsAnyIssueType:  true,
		IsGlobalContext: false,
		ContextID:       nil,
	}

	contexts, response, err := atlassian.Issue.Field.Context.Gets(context.Background(), fieldID, &options, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, fieldContext := range contexts.Values {
		log.Println(fieldContext)
	}

	return
}

func createFieldContexts() (err error) {

	log.Println("------------- createFieldContexts -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var fieldID = "customfield_10038"

	payload := jira.FieldContextPayloadScheme{
		IssueTypeIDs: []int{10004},
		ProjectIDs:   []int{10000},
		Name:         "Bug fields context",
		Description:  "A context used to define the custom field options for bugs.",
	}

	fmt.Println(payload)

	_, response, err := atlassian.Issue.Field.Context.Create(context.Background(), fieldID, &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	return
}

func main() {

	var err error

	err = getFieldContexts()
	if err != nil {
		log.Fatal(err)
	}

	err = createFieldContexts()
	if err != nil {
		log.Fatal(err)
	}

}
