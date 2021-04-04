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

	cloudAdmin.Auth.SetAccessToken(apiKey)
	cloudAdmin.Auth.SetUserAgent("curl/7.54.0")

	var accountID = "5e5f6a63157ed50cd2b9eaca"

	tokens, response, err := cloudAdmin.User.Token.Gets(context.Background(), accountID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, token := range *tokens {
		log.Println(token.ID, token.Label, token.CreatedAt, token.LastAccess)
	}
}
