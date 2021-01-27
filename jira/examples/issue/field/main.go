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

func getIssueFields() (err error) {
	log.Println("------------- getIssueFields -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	fields, response, err := atlassian.Issue.Field.Gets(context.Background())
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, field := range *fields {
		log.Println(field.ID, field.Name, field.Schema.Type)
	}

	return
}

func createCustomFields() (err error) {
	log.Println("------------- createCustomFields -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	fieldNewCreate := jira.CustomFieldScheme{
		Name:        "Alliance",
		Description: "this is the alliance description field",
		FieldType:   "cascadingselect",
		SearcherKey: "cascadingselectsearcher",
	}

	field, response, err := atlassian.Issue.Field.Create(context.Background(), &fieldNewCreate)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("field", field)

	return
}

func searchFields() (err error) {
	log.Println("------------- searchFields -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	options := jira.FieldSearchOptionsScheme{
		Types:   []string{"custom"},
		OrderBy: "lastUsed",
		Expand:  []string{"screensCount", "lastUsed"},
	}

	fields, response, err := atlassian.Issue.Field.Search(context.Background(), &options, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(fields.IsLast, fields.Total, fields.MaxResults, fields.StartAt)

	for _, field := range fields.Values {
		log.Println(field.ID, field.Description)
	}

	return
}

func main() {

	var err error

	err = getIssueFields()
	if err != nil {
		log.Fatal(err)
	}

	/*
		err = createCustomFields()
		if err != nil {
			log.Fatal(err)
		}
	*/

	err = searchFields()
	if err != nil {
		log.Fatal(err)
	}

}
