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

	var payload = &admin.SCIMUserScheme{
		UserName: "Example Username 3",
		Emails: []*admin.SCIMUserEmailScheme{
			{
				Value:   "example-2@go-atlassian.io",
				Type:    "work",
				Primary: true,
			},
		},
		Name: &admin.SCIMUserNameScheme{
			Formatted:       "Example Full Name with Last Name",
			FamilyName:      "Example Family Name",
			GivenName:       "Example Name",
			MiddleName:      "Name",
			HonorificPrefix: "",
			HonorificSuffix: "",
		},

		DisplayName:       "Example Display Name 3",
		NickName:          "Example NickName",
		Title:             "Atlassian Administrator",
		PreferredLanguage: "en-US",
		Active:            true,
	}

	newUser, response, err := cloudAdmin.SCIM.User.Create(context.Background(), directoryID, payload, nil, nil)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(newUser.ID)

}
