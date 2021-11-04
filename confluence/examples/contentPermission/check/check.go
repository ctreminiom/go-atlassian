package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"net/http"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	instance, err := confluence.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	instance.Auth.SetBasicAuth(mail, token)
	instance.Auth.SetUserAgent("curl/7.54.0")

	var (
		contentID = "76513281"
		payload   = &confluence.CheckPermissionScheme{
			Subject: &confluence.PermissionSubjectScheme{
				Type:       "user",
				Identifier: "5b86be50b8e3cb5895860d6d",
			},
			Operation: "read",
		}
	)

	check, response, err := instance.Content.Permission.Check(context.Background(), contentID, payload)
	if err != nil {

		if response != nil {
			if response.Code == http.StatusBadRequest {
				log.Println(response.API)
			}
		}
		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)

	log.Println(check.HasPermission)
}
