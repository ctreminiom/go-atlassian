package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira/v3"
	"log"
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

	var expand = []string{
		"transitions",         //For each workflow, returns information about the transitions inside the workflow.
		"transitions.rules",   //For each workflow transition, returns information about its rules. Transitions are included automatically if this expand is requested.
		"statuses",            //For each workflow, returns information about the statuses inside the workflow.
		"statuses.properties", //For each workflow status, returns information about its properties. Statuses are included automatically if this expand is requested.
		"default",             //For each workflow, returns information about whether this is the default workflow.
	}

	workflows, response, err := atlassian.Workflow.Gets(context.Background(), nil, expand, 0, 50)
	if err != nil {
		if response != nil {
			log.Println(response.Endpoint)
			log.Println(response.Code)
			log.Println(response.Bytes.String())
		}
	}

	for _, workflow := range workflows.Values {

		log.Println("--------------------")
		log.Printf("Workflow ID: %v", workflow.ID.EntityID)
		log.Printf("Workflow Name: %v", workflow.ID.Name)
		log.Printf("Workflow Description: %v", workflow.Description)

		if workflow.Statuses != nil {
			for index, status := range workflow.Statuses {
				log.Printf("Workflow Status #%v - Name: %v - ID: %v", index+1, status.Name, status.ID)

				if status.Properties != nil {
					log.Printf("Workflow Status #%v - Property IssueEditable: %v", index+1, status.Properties.IssueEditable)
				}

			}
		}

		if workflow.Transitions != nil {

			for index, transition := range workflow.Transitions {
				log.Printf("Workflow Transition #%v - Name: %v - ID: %v - Type: %v - From: %v - To: %v",
					index+1, transition.Name, transition.ID, transition.Type, transition.From, transition.To)

				if transition.Rules != nil {

					for index, condition := range transition.Rules.Conditions {
						log.Printf("Workflow Transition #%v Condition - Type: %v", index+1, condition.Type)
					}

					for index, validator := range transition.Rules.Validators {
						log.Printf("Workflow Transition #%v Validator - Type: %v", index+1, validator.Type)
					}

					for index, postFunction := range transition.Rules.PostFunctions {
						log.Printf("Workflow Transition #%v PostFunction - Type: %v", index+1, postFunction.Type)
					}
				}
			}
		}

		log.Println("--------------------")
	}

}
