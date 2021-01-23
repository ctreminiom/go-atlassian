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

func createGroup() (groupName string, err error) {

	log.Println("------------- createGroup -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	group, response, err := atlassian.Group.Create(context.Background(), "jira-users")
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("Group created", group.Name)

	groupName = group.Name

	return
}

func deleteGroup(groupName string) (err error) {

	log.Println("------------- deleteGroup -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Group.Delete(context.Background(), groupName)
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

func getGroups() (groupsNames []string, err error) {

	log.Println("------------- getGroups -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	options := jira.GroupBulkOptionsScheme{
		GroupIDs:   nil,
		GroupNames: nil,
	}

	groups, response, err := atlassian.Group.Bulk(context.Background(), &options, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(groups.IsLast)

	for index, group := range groups.Values {
		groupsNames = append(groupsNames, group.Name)
		log.Printf("#%v, Group: %v", index, group.Name)
	}

	return
}

func getGroupMembers(groupName string) (err error) {

	log.Println("------------- getGroupMembers -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	members, response, err := atlassian.Group.Members(context.Background(), groupName, false, 0, 100)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(members.IsLast)

	for index, member := range members.Values {
		log.Printf("#%v - Group %v - Member Mail %v - Member AccountID %v", index, groupName, member.EmailAddress, member.AccountID)
	}

	return
}

func addUserToGroup(groupName, accountID string) (err error) {

	log.Println("------------- addUserToGroup -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	_, response, err := atlassian.Group.Add(context.Background(), groupName, accountID)
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

func removeUserFromGroup(groupName, accountID string) (err error) {

	log.Println("------------- removeUserFromGroup -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Group.Remove(context.Background(), groupName, accountID)
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

	log.Println("Creating the new Jira Cloud group")
	groupName, err := createGroup()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("The group has been created, %v", groupName)

	log.Println("Adding user to a group")
	if err := addUserToGroup("jira-users", "5b86be50b8e3cb5895860d6d"); err != nil {
		log.Fatal(err)
	}

	log.Println("Removing user to a group")
	if err := removeUserFromGroup("jira-users", "5b86be50b8e3cb5895860d6d"); err != nil {
		log.Fatal(err)
	}

	log.Printf("Delete the group %v", groupName)
	if err := deleteGroup(groupName); err != nil {
		log.Fatal(err)
	}

	log.Println("Getting the Jira Groups")
	groups, err := getGroups()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("We've found %v groups in your Jira Cloud instance", len(groups))
	log.Println("Extracting the group members")

	for _, group := range groups {
		if err = getGroupMembers(group); err != nil {
			log.Println(err)
		}
	}

}
