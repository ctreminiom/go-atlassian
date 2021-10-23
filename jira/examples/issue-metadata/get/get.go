package main

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	var (
		issueKeyOrID           = "KP-19"
		overrideScreenSecurity = false
		overrideEditableFlag   = false
		ctx                    = context.Background()
	)

	metadata, response, err := atlassian.Issue.Metadata.Get(ctx, issueKeyOrID, overrideScreenSecurity, overrideEditableFlag)
	if err != nil {

		if response != nil {
			log.Println(response.Bytes.String())
		}

		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	// =========================================================================
	// If you need to extract the issue type metadata
	// =========================================================================
	fmt.Println(metadata.Get("fields.issuetype.required"))
	fmt.Println(metadata.Get("fields.issuetype.name"))
	fmt.Println(metadata.Get("fields.issuetype.key"))

	for _, key := range metadata.Get("fields.issuetype.allowedValues").Array() {
		fmt.Println(key)
	}

	// =========================================================================
	// If you need to extract the components metadata
	// =========================================================================
	fmt.Println(metadata.Get("fields.components.required"))
	fmt.Println(metadata.Get("fields.components.name"))
	fmt.Println(metadata.Get("fields.components.key"))

	for _, key := range metadata.Get("fields.components.allowedValues").Array() {
		fmt.Println(key)
	}

}
