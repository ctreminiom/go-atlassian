package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"os"
)

func main()  {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := confluence.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")


	response, err := atlassian.Content.Label.Remove(context.Background(), "80412692", "label-02")
	if err != nil {
		log.Println(response.Endpoint)
		log.Println(response.Bytes.String())
		log.Fatal(err)
	}

}