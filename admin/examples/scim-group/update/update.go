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

	var (
		directoryID  = "bcdde508-ee40-4df2-89cc-d3f6292c5971"
		groupID      = "e18da5e4-ba2e-4039-9046-30000af6c0b7"
		newGroupName = "scim-jira-users-updated"
	)

	group, response, err := cloudAdmin.SCIM.Group.Update(context.Background(), directoryID, groupID, newGroupName)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(group)
}