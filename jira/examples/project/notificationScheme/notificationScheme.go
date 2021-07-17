package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	/*
		----------- Set an environment variable in git bash -----------
		export HOST="https://ctreminiom.atlassian.net/"
		export MAIL="MAIL_ADDRESS"
		export TOKEN="TOKEN_API"

		Docs: https://stackoverflow.com/questions/34169721/set-an-environment-variable-in-git-bash
	*/

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

	notificationScheme, response, err := atlassian.Project.NotificationScheme(context.Background(), "KP", []string{"all"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(notificationScheme)

	log.Println(notificationScheme.Name, notificationScheme.ID)

	for _, event := range notificationScheme.NotificationSchemeEvents {

		log.Println(event.Event.ID, event.Event.Name)

		for index, notification := range event.Notifications {
			log.Println(index, event.Event.Name, notification.ID, notification.NotificationType, notification.Parameter)
		}

	}
}
