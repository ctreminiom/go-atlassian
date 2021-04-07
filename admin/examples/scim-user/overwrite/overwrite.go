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
		directoryID = "bcdde508-ee40-4df2-89cc-d3f6292c5971"
		userID      = "ef5ff80e-9ca6-449c-8cca-5b621085c6c9"
	)

	payload := &admin.SCIMUserScheme{
		UserName:    "username-updated-with-overwrite-method",
		DisplayName: "AA",
		NickName:    "AA",
		Title:       "AA",
		Department:  "President",
		Emails: []*admin.SCIMUserEmailScheme{
			{
				Value:   "carlos@go-atlassian.io",
				Type:    "work",
				Primary: true,
			},
		},
	}
	userUpdated, response, err := cloudAdmin.SCIM.User.Overwrite(context.Background(), directoryID, userID, payload, nil, nil)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(userUpdated.DisplayName)
	log.Println(userUpdated.Active)
	log.Println(userUpdated.Department)

}
