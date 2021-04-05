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

	var accountID = "5e5f6a63157ed50cd2b9eaca"

	/*
		-------------- NOTE --------------
		The fields the endpoint can edit depends how you configured the user provisioning, for example:
		If the provisioning is made using the GSuite integration or the SCIM connector, you won't be able to
		edit that field using this endpoint because that field is blocked.

		You can check the availability using the Permission method and searching on the tag permissions.EmailSet.Allowed
		-------------- NOTE --------------
	*/
	var payload = make(map[string]interface{})
	payload["nickname"] = "marshmallow"

	userUpdated, response, err := cloudAdmin.User.Update(context.Background(), accountID, payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("New NickName, ", userUpdated.Account.Nickname)

}
