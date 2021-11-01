package main

import (
	"context"
	"encoding/json"
	"github.com/ctreminiom/go-atlassian/jira/v3"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := v3.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var (
		issueKey   = "DESK-12"
		approvalID = 2
	)

	approvalsMembers, response, err := atlassian.ServiceManagement.Request.Approval.Get(context.Background(), issueKey, approvalID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	dataAsJson, err := json.MarshalIndent(approvalsMembers, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(dataAsJson))

}
