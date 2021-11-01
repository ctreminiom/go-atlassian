package main

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/jira/v3"
	"log"
	"math/rand"
	"os"
)

func main() {

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

	atlassian, err := v3.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var (
		projectCategoryID = 10000
		payload           = &v3.ProjectCategoryPayloadScheme{
			Name: fmt.Sprintf("Category #%v - updated", rand.Intn(100)),
		}
	)

	categoryUpdated, response, err := atlassian.Project.Category.Update(context.Background(), projectCategoryID, payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Printf("The project category %v has been updated", categoryUpdated.ID)
}
