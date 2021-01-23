package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

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

func setDefaultShareScope() (err error) {

	log.Println("------------- setDefaultShareScope -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Filter.Share.SetScope(context.Background(), "GLOBAL")
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println(response)

	return
}

func getFilterShareScope() (err error) {

	log.Println("------------- getFilterShareScope -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	scope, response, err := atlassian.Filter.Share.Scope(context.Background())
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("Scope", scope)

	return
}

func getFilterPermissions(filterID string) (err error) {

	log.Println("------------- getFilterPermissions -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	permissions, response, err := atlassian.Filter.Share.Gets(context.Background(), filterID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for index, permission := range *permissions {
		log.Println(index, permission.ID, permission.Type)
	}

	return
}

func addSharePermission(filterID string) (err error) {

	log.Println("------------- addSharePermission -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	/*
		We can add different share permissions, for example:

		---- Project ID only
		payload := jira.PermissionFilterBodyScheme{
				Type:      "project",
				ProjectID: "10000",
			}

		---- Project ID and role ID
		payload := jira.PermissionFilterBodyScheme{
				Type:          "project",
				ProjectID:     "10000",
				ProjectRoleID: "222222",
			}

		==== Group Name
		payload := jira.PermissionFilterBodyScheme{
				Type:          "group",
				GroupName: "jira-users",
			}
	*/

	payload := jira.PermissionFilterBodyScheme{
		Type:      "project",
		ProjectID: "10000",
	}

	permissions, response, err := atlassian.Filter.Share.Add(context.Background(), filterID, &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
			log.Println("Response HTTP Code", response.StatusCode)
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for index, permission := range *permissions {
		log.Println(index, permission.ID, permission.Type)
	}

	return
}

func getFilterPermission(filterID, permissionID string) (err error) {

	log.Println("------------- getFilterPermission -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	permission, response, err := atlassian.Filter.Share.Get(context.Background(), filterID, permissionID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(permission)

	return
}

func deleteFilterPermission(filterID, permissionID string) (err error) {

	log.Println("------------- deleteFilterPermission -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Filter.Share.Delete(context.Background(), filterID, permissionID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	return
}

func main() {

	if err := getFilterShareScope(); err != nil {
		log.Fatal(err)
	}

	if err := getFilterPermissions("10000"); err != nil {
		log.Fatal(err)
	}

	if err := addSharePermission("10000"); err != nil {
		log.Fatal(err)
	}

}
