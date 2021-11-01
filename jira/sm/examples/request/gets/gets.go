package main

import (
	"context"
	"encoding/json"
	"github.com/ctreminiom/go-atlassian/jira/sm"
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

	options := &sm.RequestGetOptionsScheme{
		SearchTerm:        "",
		RequestOwnerships: []string{"OWNED_REQUESTS"},
		RequestStatus:     "ALL_REQUESTS",
		ApprovalStatus:    "",
		OrganizationId:    0,
		ServiceDeskID:     0,
		RequestTypeID:     0,
		Expand:            []string{"serviceDesk", "requestType", "status", "action"},
	}

	customerRequests, response, err := atlassian.ServiceManagement.Request.Gets(context.Background(), options, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, customRequest := range customerRequests.Values {

		dataAsJson, err := json.MarshalIndent(customRequest, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(dataAsJson))
	}

}
