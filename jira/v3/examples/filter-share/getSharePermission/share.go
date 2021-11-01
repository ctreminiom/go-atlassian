package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira/v3"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := v3.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	permission, response, err := atlassian.Filter.Share.Get(context.Background(), 10001, 100012)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(permission)
}
