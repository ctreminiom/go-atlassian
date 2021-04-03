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

	cloudAdmin.Auth.SetAccessToken(apiKey)
	cloudAdmin.Auth.SetUserAgent("curl/7.54.0")

	var (
		organizationID = "9a1jj823-jac8-123d-jj01-63315k059cb2"
		cursor         string
		domainChunks   []*admin.OrganizationDomainPageScheme
	)

	for {

		domains, response, err := cloudAdmin.Organization.Domains(context.Background(), organizationID, cursor)
		if err != nil {
			if response != nil {
				log.Println("Response HTTP Response", string(response.BodyAsBytes))
			}
			log.Fatal(err)
		}

		log.Println("Response HTTP Code", response.StatusCode)
		log.Println("HTTP Endpoint Used", response.Endpoint)
		domainChunks = append(domainChunks, domains)

		if len(domains.Links.Next) == 0 {
			break
		}

		//extract the next cursor pagination
		nextAsURL, err := url.Parse(domains.Links.Next)
		if err != nil {
			log.Fatal(err)
		}

		cursor = nextAsURL.Query().Get("cursor")
	}

	for _, chunk := range domainChunks {

		for _, domain := range chunk.Data {
			log.Println(domain.ID, domain.Attributes.Name, domain.Attributes.Claim.Status)
		}

	}

}
