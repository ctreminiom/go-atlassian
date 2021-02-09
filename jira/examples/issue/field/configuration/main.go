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

func getIssueFieldConfiguration() (err error) {
	log.Println("------------- getIssueFieldConfiguration -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	fieldConfigurations, response, err := atlassian.Issue.Field.Configuration.Gets(
		context.Background(),
		nil,
		false,
		0,
		50,
	)

	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, fieldConfiguration := range fieldConfigurations.Values {
		log.Println(fieldConfiguration)
	}

	return
}

func getIssueFieldConfigurationItems() (err error) {

	log.Println("------------- getIssueFieldConfigurationItems -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	configurationItems, response, err := atlassian.Issue.Field.Configuration.Items(context.Background(), "10000", 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, item := range configurationItems.Values {
		log.Println(item)
	}

	return
}

func getIssueFieldConfigurationSchemes() (err error) {
	log.Println("------------- getIssueFieldConfigurationSchemes -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	schemes, response, err := atlassian.Issue.Field.Configuration.Schemes(context.Background(), nil, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, scheme := range schemes.Values {
		log.Println(scheme)
	}

	return
}

func getIssueFieldConfigurationIssueTypeItems() (err error) {
	log.Println("------------- getIssueFieldConfigurationIssueTypeItems -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	configurationIssueTypeItems, response, err := atlassian.Issue.Field.Configuration.IssueTypeItems(context.Background(), nil, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, issueTypeItem := range configurationIssueTypeItems.Values {
		log.Println(issueTypeItem)
	}

	return
}

func getIssueFieldConfigurationSchemesByProject() (err error) {
	log.Println("------------- getIssueFieldConfigurationSchemesByProject -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	schemes, response, err := atlassian.Issue.Field.Configuration.SchemesByProject(context.Background(), []int{10001, 10000}, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, scheme := range schemes.Values {
		log.Println(scheme)
	}

	return
}

func main() {

	var err error
	if err = getIssueFieldConfiguration(); err != nil {
		log.Fatal(err)
	}

	if err = getIssueFieldConfigurationItems(); err != nil {
		log.Fatal(err)
	}

	if err = getIssueFieldConfigurationSchemes(); err != nil {
		log.Fatal(err)
	}

	if err = getIssueFieldConfigurationIssueTypeItems(); err != nil {
		log.Fatal(err)
	}

	if err = getIssueFieldConfigurationSchemesByProject(); err != nil {
		log.Fatal(err)
	}

}
