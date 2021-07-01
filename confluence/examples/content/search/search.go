package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/confluence"
	"log"
	"net/http"
	"os"
)

func main()  {

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
		cql = "type=page"
		cqlContext = ""
		expand = []string{"childTypes.all", "metadata.labels"}
		maxResults = 50
	)

	contentPage, response, err := instance.Content.Search(context.Background(), cql, cqlContext, expand, "", maxResults)
	if err != nil {

		if response.Code == http.StatusBadRequest {
			log.Println(response.API)
		}
		log.Fatal(err)
	}

	log.Println("Endpoint:", response.Endpoint)
	log.Println("Status Code:", response.Code)
	log.Println(contentPage.Links.Next)


	for _, content := range contentPage.Results {
		log.Println(content.Title, content.ID)
	}
}
