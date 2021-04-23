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

	payload := &admin.SCIMUserToPathScheme{
		Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
	}

	if err = payload.AddStringOperation("replace", "displayName", "Docs Atlassian DisplayName 2"); err != nil {
		log.Fatal(err)
	}

	if err = payload.AddStringOperation("replace", "userName", "user-name-updated2"); err != nil {
		log.Fatal(err)
	}

	if err = payload.AddBoolOperation("replace", "active", false); err != nil {
		log.Fatal(err)
	}

	if err = payload.AddComplexOperation("add", "emails", []*admin.SCIMUserComplexOperationScheme{
		{
			Value:     "primary@go-atlassian.io",
			ValueType: "work",
			Primary:   true,
		},
		{
			Value:     "second@go-atlassian.io",
			ValueType: "other",
			Primary:   false,
		},
	}); err != nil {
		log.Fatal(err)
	}

	/*
		payload := &admin.SCIMUserToPathScheme{
			Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
			Operations: []*admin.SCIMUserToPathOperationScheme{
				{
					Op:    "replace",
					Path:  "displayName",
					Value: "Docs Atlassian DisplayName",
				},

				{
					Op:    "replace",
					Path:  "userName",
					Value: "user-name-updated",
				},
				{
					Op:    "replace",
					Path:  "name.formatted",
					Value: "Ms. Barbara J Jensen, III",
				},
				{
					Op:    "replace",
					Path:  "name.familyName",
					Value: "Jensen",
				},
				{
					Op:    "replace",
					Path:  "name.givenName",
					Value: "Barbara",
				},
				{
					Op:    "replace",
					Path:  "name.middleName",
					Value: "Jane",
				},
				{
					Op:    "replace",
					Path:  "name.honorificPrefix",
					Value: "Ms.",
				},
				{
					Op:    "replace",
					Path:  "name.honorificSuffix",
					Value: "III",
				},

				{
					Op:    "replace",
					Path:  "nickName",
					Value: "Bobby",
				},

				{
					Op:    "replace",
					Path:  "title",
					Value: "Vice President.",
				},

				{
					Op:    "replace",
					Path:  "title",
					Value: "Vice President.",
				},
			},
		}
	*/

	userUpdated, response, err := cloudAdmin.SCIM.User.Path(context.Background(), directoryID, userID, payload, nil, nil)
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

	for _, mail := range userUpdated.Emails {
		log.Println(mail)
	}
}
