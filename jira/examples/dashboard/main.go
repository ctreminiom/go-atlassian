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

func getDashboards() (err error) {

	log.Println("------------- getDashboards -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	dashboards, response, err := atlassian.Dashboard.Gets(context.Background(), 0, 50, "")

	if err != nil {

		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}

		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(len(dashboards.Dashboards))

	return
}

func main() {
	if err := getDashboards(); err != nil {
		log.Fatal(err)
	}
}
