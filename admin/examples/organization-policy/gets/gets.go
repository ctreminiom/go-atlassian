package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/admin"
	"log"
	"net/url"
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

	var (
		organizationID = "9a1jj823-jac8-123d-jj01-63315k059cb2"
		policyType     = ""
		policyChunks   []*admin.OrganizationPolicyPageScheme
		cursor         string
	)

	for {

		policies, response, err := cloudAdmin.Organization.Policy.Gets(context.Background(), organizationID, policyType, cursor)
		if err != nil {
			if response != nil {
				log.Println("Response HTTP Response", string(response.BodyAsBytes))
			}
			log.Fatal(err)
		}

		log.Println("Response HTTP Code", response.StatusCode)
		log.Println("HTTP Endpoint Used", response.Endpoint)
		policyChunks = append(policyChunks, policies)

		if len(policies.Links.Next) == 0 {
			break
		}

		//extract the next cursor pagination
		nextAsURL, err := url.Parse(policies.Links.Next)
		if err != nil {
			log.Fatal(err)
		}

		cursor = nextAsURL.Query().Get("cursor")
	}

	for _, chunk := range policyChunks {

		for _, policy := range chunk.Data {

			log.Println(policy.ID, policy.Type, policy.Attributes.Status)
		}

	}

}
