package main

import (
	"context"
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

	options := v3.FilterSearchOptionScheme{
		Name:      "",
		AccountID: "",
		Group:     "",
		ProjectID: 0,
		IDs:       nil,
		OrderBy:   "description",
		Expand:    nil,
	}

	filters, response, err := atlassian.Filter.Search(context.Background(), &options, 0, 10)
	if err != nil {
		if response != nil {
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		return
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("Filters found", len(filters.Values))
}
