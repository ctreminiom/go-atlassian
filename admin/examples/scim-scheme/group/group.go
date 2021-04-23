package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/admin"
	"log"
	"os"
)

func main() {

	//ATLASSIAN_ADMIN_TOKEN
	var scimApiKey = os.Getenv("ATLASSIAN_SCIM_API_KEY")

	cloudAdmin, err := admin.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	cloudAdmin.Auth.SetBearerToken(scimApiKey)
	cloudAdmin.Auth.SetUserAgent("curl/7.54.0")

	var directoryID = "bcdde508-ee40-4df2-89cc-d3f6292c5971"

	schemas, response, err := cloudAdmin.SCIM.Scheme.Group(context.Background(), directoryID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, attribute := range schemas.Attributes {
		log.Println("----------------------")
		log.Println("Type", attribute.Type)
		log.Println("Description", attribute.Description)
		log.Println("Name", attribute.Name)
		log.Println("Required", attribute.Required)
		log.Println("Returned", attribute.Returned)
		log.Println("Mutability", attribute.Mutability)
		log.Println("SubAttributes", len(attribute.SubAttributes))

		for _, subAttribute := range attribute.SubAttributes {
			log.Println("==============================")
			log.Println("==", subAttribute.Uniqueness)
			log.Println("==", subAttribute.Mutability)
			log.Println("==", subAttribute.Returned)
			log.Println("==", subAttribute.Required)
			log.Println("==", subAttribute.Name)
			log.Println("==", subAttribute.Description)
			log.Println("==============================")

		}

		log.Println("----------------------")

	}
}
