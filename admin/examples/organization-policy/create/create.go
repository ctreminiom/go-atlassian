package main

import (
	"context"
	"encoding/json"
	"fmt"
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

	cloudAdmin.Auth.SetAccessToken(apiKey)
	cloudAdmin.Auth.SetUserAgent("curl/7.54.0")

	payload := &admin.OrganizationPolicyData{
		Type: "policy",
		Attributes: admin.OrganizationPolicyAttributes{
			Type:   "data-residency", //ip-allowlist
			Name:   "Name of this Policy",
			Status: "enabled", //disabled
		},
	}

	var organizationID = "9a1jj823-jac8-123d-jj01-63315k059cb2"

	newPolicy, response, err := cloudAdmin.Organization.Policy.Create(context.Background(), organizationID, payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	policyAsJSONKeys, err := json.MarshalIndent(newPolicy, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MarshalIndent Struct keys output\n %s\n", string(policyAsJSONKeys))

	fmt.Println(payload)

}
