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

	fieldConfigurationSchemes, response, err := atlassian.Issue.Field.Configuration.Schemes(context.Background(), nil, 0, 50)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, fieldConfigurationScheme := range fieldConfigurationSchemes.Values {
		log.Println(fieldConfigurationScheme)
	}

}
