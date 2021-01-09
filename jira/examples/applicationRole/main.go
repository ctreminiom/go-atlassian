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

func getApplicationRoles() (err error) {

	log.Println("------------- getApplicationRoles -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	roles, response, err := atlassian.Role.Gets(context.Background())
	if err != nil {
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, role := range *roles {
		log.Println(role.Key)
	}

	return
}

func getApplicationRole() (err error) {

	log.Println("------------- getApplicationRole -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	role, response, err := atlassian.Role.Get(context.Background(), "jira-software")
	if err != nil {

		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}

		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(role.Key)

	return
}

func main() {

	if err := getApplicationRoles(); err != nil {
		log.Fatal(err)
	}

	if err := getApplicationRole(); err != nil {
		log.Fatal(err)
	}
}
