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
		log.Fatal(err)
	}

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
