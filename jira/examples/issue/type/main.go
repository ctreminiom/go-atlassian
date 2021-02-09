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

func getIssueTypes() (issueTypesIDs []string) {

	log.Println("---------------------- getIssueTypes ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	types, response, err := atlassian.Issue.Type.Gets(context.Background())
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, value := range *types {
		log.Println(value.ID, value.Name, value.Subtask)
		issueTypesIDs = append(issueTypesIDs, value.ID)
	}

	return
}

func getIssueType(id string) {

	log.Println("---------------------- getIssueType ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	issueType, response, err := atlassian.Issue.Type.Get(context.Background(), id)

	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(issueType.Name)

}

func createIssueType() (id string) {

	log.Println("---------------------- createIssueType ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	issueTypePayload := jira.IssueTypePayloadScheme{
		Name:        "Risk",
		Description: "this is the issue type description",
		Type:        "standard",
	}

	issueType, response, err := atlassian.Issue.Type.Create(context.Background(), &issueTypePayload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(issueType.Name)

	id = issueType.ID

	return
}

func updateIssueType(id string) {

	log.Println("---------------------- updateIssueType ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	issueTypePayload := jira.IssueTypePayloadScheme{
		Name:        "Risk UPDATED",
		Description: "this is the issue type description, UPDATED",
	}

	issueType, response, err := atlassian.Issue.Type.Update(context.Background(), id, &issueTypePayload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(issueType.Name)
}

func deleteIssueType(id string) {

	log.Println("---------------------- deleteIssueType ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Issue.Type.Delete(context.Background(), id)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
}

func getAlternativesIssueTypes(id string) {

	log.Println("---------------------- getAlternativesIssueTypes ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	issueTypes, response, err := atlassian.Issue.Type.Alternatives(context.Background(), id)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, issueType := range *issueTypes {
		log.Println(issueType)
	}

}

func main() {

	for _, value := range getIssueTypes() {
		getIssueType(value)
	}

	newIssueTypeID := createIssueType()
	log.Println("The issue type has been created", newIssueTypeID)

	getAlternativesIssueTypes(newIssueTypeID)
	updateIssueType(newIssueTypeID)
	deleteIssueType(newIssueTypeID)

}
