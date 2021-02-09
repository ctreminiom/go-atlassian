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

func getIssueLinkTypes() (ids []string) {

	log.Println("-------------------- getIssueLinkTypes --------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	linkTypes, response, err := atlassian.Issue.Link.Type.Gets(context.Background())
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, linkType := range linkTypes.IssueLinkTypes {
		log.Println(linkType.Name)
		ids = append(ids, linkType.ID)
	}

	return
}

func getIssueLinkType(ids []string) {

	log.Println("-------------------- getIssueLinkType --------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	for _, id := range ids {
		linkType, response, err := atlassian.Issue.Link.Type.Get(context.Background(), id)
		if err != nil {
			if response != nil {
				log.Println("Response HTTP Response", string(response.BodyAsBytes))
			}
			log.Fatal(err)
		}

		log.Println("Response HTTP Code", response.StatusCode)
		log.Println("HTTP Endpoint Used", response.Endpoint)
		log.Println(linkType)

	}
}

func createIssueLinkType() (id string) {

	log.Println("-------------------- createIssueLinkType --------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	payload := jira.IssueLinkTypePayloadScheme{
		Inward:  "Clone/Duplicated by",
		Name:    "Clone/Duplicate",
		Outward: "Clone/Duplicates",
	}

	issueLinkType, response, err := atlassian.Issue.Link.Type.Create(context.Background(), &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	id = issueLinkType.ID
	return

}

func updateIssueTypeLink(id string) {

	log.Println("-------------------- updateIssueTypeLink --------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	payload := jira.IssueLinkTypePayloadScheme{
		Inward:  "Clone/Duplicated by - Updated",
		Name:    "Clone/Duplicate - Updated",
		Outward: "Clone/Duplicates - Updated",
	}

	issueLinkType, response, err := atlassian.Issue.Link.Type.Update(context.Background(), id, &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(issueLinkType)

}

func deleteIssueTypeLink(id string) {

	log.Println("-------------------- deleteIssueTypeLink --------------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Issue.Link.Type.Delete(context.Background(), id)
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
	ids := getIssueLinkTypes()
	getIssueLinkType(ids)

	id := createIssueLinkType()
	updateIssueTypeLink(id)
	deleteIssueTypeLink(id)
}
