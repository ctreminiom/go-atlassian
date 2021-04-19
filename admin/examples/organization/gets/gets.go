package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/admin"
	"log"
	"os"
)

func main() {

	//ATLASSIAN_ADMIN_TOKEN
	var apiKey = os.Getenv("ATLASSIAN_ADMIN_TOKEN")

	cloudAdmin, err := admin.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	cloudAdmin.Auth.SetBearerToken(apiKey)
	cloudAdmin.Auth.SetUserAgent("curl/7.54.0")

	organizations, response, err := cloudAdmin.Organization.Gets(context.Background(), "")
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, organization := range organizations.Data {
		log.Println(organization.ID, organization.Attributes.Name)
	}
}
