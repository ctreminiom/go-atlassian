package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func createCustomer() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	jiraCloud, err := jira.New(nil, host)
	if err != nil {
		return
	}

	jiraCloud.Auth.SetBasicAuth(mail, token)
	jiraCloud.Auth.SetUserAgent("curl/7.54.0")

	var (
		email       = "example@gmail.com"
		displayName = "Example Customer"
	)

	newCustomer, response, err := jiraCloud.ServiceManagement.Customer.Create(context.Background(), email, displayName)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println("The new customer has been created!!")
	log.Println("-------------------------")
	log.Println(newCustomer.Name)
	log.Println(newCustomer.DisplayName)
	log.Println(newCustomer.AccountID)
	log.Println(newCustomer.EmailAddress)
	log.Println(newCustomer.Links)
	log.Println("-------------------------")
}
