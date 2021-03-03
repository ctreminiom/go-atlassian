package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
	"strconv"
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

func getIssueTypeSchemes() {

	log.Println("---------------------- getIssueTypeSchemes ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var startAt int
	for {

		issueTypesSchemes, response, err := atlassian.Issue.Type.Scheme.Gets(context.Background(), nil, startAt, 50)
		if err != nil {
			if response != nil {
				log.Println("Response HTTP Response", string(response.BodyAsBytes))
			}
			log.Fatal(err)
		}

		log.Println("Response HTTP Code", response.StatusCode)
		log.Println("HTTP Endpoint Used", response.Endpoint)

		for _, scheme := range issueTypesSchemes.Values {
			log.Println(scheme.ID, scheme.Name, scheme.DefaultIssueTypeID)
		}

		if issueTypesSchemes.IsLast {
			break
		}

		startAt += 50
	}

}

func createIssueTypeScheme() (id string) {

	log.Println("---------------------- createIssueTypeScheme ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	payload := jira.IssueTypeSchemePayloadScheme{
		DefaultIssueTypeID: "10001",
		IssueTypeIds:       []string{"10001", "10002", "10005"},
		Name:               "Kanban Issue Type Scheme 1",
		Description:        "A collection of issue types suited to use in a kanban style project.",
	}

	issueTypeSchemeID, response, err := atlassian.Issue.Type.Scheme.Create(context.Background(), &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("issueTypeSchemeID", issueTypeSchemeID)

	return issueTypeSchemeID
}

func updateIssueTypeScheme(id int) {

	log.Println("---------------------- updateIssueTypeScheme ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	payload := jira.IssueTypeSchemePayloadScheme{
		Name:        "Kanban Issue Type Scheme - UPDATED",
		Description: "A collection of issue types suited to use in a kanban style project.- UPDATED",
	}

	response, err := atlassian.Issue.Type.Scheme.Update(context.Background(), id, &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

}

func deleteIssueTypeScheme(id int) {

	log.Println("---------------------- deleteIssueTypeScheme ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Issue.Type.Scheme.Delete(context.Background(), id)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
}

func getIssueTypeSchemeItems() {

	log.Println("---------------------- getIssueTypeSchemeItems ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	items, response, err := atlassian.Issue.Type.Scheme.Items(context.Background(), nil, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, item := range items.Values {
		log.Println(item)
	}

}

func getIssueTypeSchemesForProjects() {

	log.Println("---------------------- getIssueTypeSchemesForProjects ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	issueTypesSchemes, response, err := atlassian.Issue.Type.Scheme.Projects(context.Background(), []int{10000}, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, issueTypeScheme := range issueTypesSchemes.Values {
		log.Println(issueTypeScheme.IssueTypeScheme.Name, issueTypeScheme.ProjectIds)
	}

}

func assignIssueTypeToProject(schemeID, projectID string) {

	log.Println("---------------------- assignIssueTypeToProject ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Issue.Type.Scheme.Assign(context.Background(), schemeID, projectID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
}

func addIssueTypesToScheme(schemeID int) {
	log.Println("---------------------- addIssueTypesToScheme ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Issue.Type.Scheme.AddIssueTypes(context.Background(), schemeID, []int{10003, 10000})
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

}

func removeIssueTypeFromIssueScheme(schemeID int) {

	log.Println("---------------------- removeIssueTypeFromIssueScheme ----------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Issue.Type.Scheme.RemoveIssueType(context.Background(), schemeID, 10003)
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
	getIssueTypeSchemes()

	id := createIssueTypeScheme()
	idAsInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	updateIssueTypeScheme(idAsInt)
	addIssueTypesToScheme(idAsInt)
	removeIssueTypeFromIssueScheme(idAsInt)
	deleteIssueTypeScheme(idAsInt)
	getIssueTypeSchemeItems()
	getIssueTypeSchemesForProjects()
	assignIssueTypeToProject("10131", "10000")
	assignIssueTypeToProject("10131", "10000")

}
