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

func searchUsingGET() {

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var (
		jql    = "order by created DESC"
		fields = []string{"status"}
		expand = []string{"changelog", "renderedFields", "names", "schema", "transitions", "operations", "editmeta"}
	)

	issues, response, err := atlassian.Issue.Search.Get(context.Background(), jql, fields, expand, 0, 50, "")
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println(issues.Total)

	for _, issue := range issues.Issues {
		for _, history := range issue.Changelog.Histories {

			for _, item := range history.Items {
				log.Println(issue.Key, item.Field, history.Author.DisplayName)
			}
		}
	}

	for _, issue := range issues.Issues {
		for _, transition := range issue.Transitions {
			log.Println(issue.Key, transition.Name, transition.ID)
		}

	}

}

func searchUsingPost() {

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var (
		jql    = "order by created DESC"
		fields = []string{"status"}
		expand = []string{"changelog", "renderedFields", "names", "schema", "transitions", "operations", "editmeta"}
	)

	issues, response, err := atlassian.Issue.Search.Post(context.Background(), jql, fields, expand, 0, 50, "")
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println(issues.Total)

}

func main() {

	searchUsingGET()
	log.Println("---------------- searchUsingPost --------------")
	searchUsingPost()

}
