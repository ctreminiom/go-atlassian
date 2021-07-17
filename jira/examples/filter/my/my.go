package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	myFilters, response, err := atlassian.Filter.My(context.Background(), false, []string{"sharedUsers", "subscriptions"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("my filters", len(myFilters))

	for _, filter := range myFilters {
		log.Println(filter.ID)

		for _, shareUser := range filter.ShareUsers.Items {
			log.Println(shareUser.Name, shareUser.DisplayName)
		}

		for _, subscription := range filter.Subscriptions.Items {
			log.Println(subscription.ID)
		}

	}

}
