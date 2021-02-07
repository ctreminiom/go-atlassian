package main

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
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

func createFilter() (filterID string, err error) {

	log.Println("------------- createFilter -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	newFilterBody := jira.FilterBodyScheme{
		Name:        fmt.Sprintf("Filter #%v", uuid.New().String()),
		Description: "Filter's description",
		JQL:         "issuetype = Bug",
		Favorite:    false,
	}

	filter, response, err := atlassian.Filter.Create(context.Background(), &newFilterBody)

	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println(string(response.BodyAsBytes))

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Printf("The filter has been created: %v - %v", filter.ID, filter.Name)
	filterID = filter.ID

	return
}

func getFilter(filterID int) (err error) {

	log.Println("------------- getFilter -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	filter, response, err := atlassian.Filter.Get(context.Background(), filterID, []string{})
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("Get Filter result", filter.Name, filter.Name)

	return
}

func updateFilter(filterID int) (err error) {

	log.Println("------------- updateFilter -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	payload := jira.FilterBodyScheme{
		JQL: "issuetype = Story",
	}

	filter, response, err := atlassian.Filter.Update(context.Background(), filterID, &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("new JQL filter value", filter.Jql)

	return
}

func deleteFilter(filterID int) (err error) {

	log.Println("------------- deleteFilter -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	response, err := atlassian.Filter.Delete(context.Background(), filterID)
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

func getMyFilters() (err error) {

	log.Println("------------- getMyFilters -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	myFilters, response, err := atlassian.Filter.My(context.Background(), false, nil)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("my filters", len(*myFilters))

	for _, filter := range *myFilters {
		log.Println(filter)
	}

	return
}

func getFavoriteFilters() (err error) {

	log.Println("------------- getFavoriteFilters -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	filters, response, err := atlassian.Filter.Favorite(context.Background())
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("favorite filters", len(*filters))

	for _, filter := range *filters {
		log.Println(filter)
	}

	return
}

func searchFilters() (err error) {

	log.Println("------------- searchFilters -----------------")

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	options := jira.FilterSearchOptionScheme{
		Name:      "",
		AccountID: "",
		Group:     "",
		ProjectID: 0,
		IDs:       nil,
		OrderBy:   "description",
		Expand:    nil,
	}

	filters, response, err := atlassian.Filter.Search(context.Background(), &options, 0, 10)
	if err != nil {
		if response != nil {
			log.Println("HTTP Endpoint Used", response.Endpoint)
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("Filters found", len(filters.Values))

	return
}

func main() {

	if err := getMyFilters(); err != nil {
		log.Fatal(err)
	}

	if err := searchFilters(); err != nil {
		log.Fatal(err)
	}

	if err := getFavoriteFilters(); err != nil {
		log.Fatal(err)
	}

	filterID, err := createFilter()
	if err != nil {
		log.Fatal(err)
	}

	filterIDAsInt, err := strconv.Atoi(filterID)
	if err != nil {
		log.Fatal(err)
	}

	if err := getFilter(filterIDAsInt); err != nil {
		log.Fatal(err)
	}

	if err := updateFilter(filterIDAsInt); err != nil {
		log.Fatal(err)
	}

	if err := deleteFilter(filterIDAsInt); err != nil {
		log.Fatal(err)
	}

}
