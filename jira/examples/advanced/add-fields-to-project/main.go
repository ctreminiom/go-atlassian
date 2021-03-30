package main

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	/*
		Workflow:

		1. Select the project to use
		2. Select the fields you want to add
		3. Check the issue field category
	*/

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	projectChunks, err := fetchProjects(atlassian, 0)
	if err != nil {
		log.Fatal(err)
	}

	var projectKeys []string
	for _, chunk := range projectChunks {
		for _, project := range chunk.Values {
			projectKeys = append(projectKeys, project.Key)
		}
	}

	projectSelected := ""
	prompt := &survey.Select{
		Message: "Choose a project:",
		Options: projectKeys,
	}

	if err = survey.AskOne(prompt, &projectSelected); err != nil {
		log.Fatal(err)
	}

	// 3.
	customFields, _, err := atlassian.Issue.Field.Gets(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var fieldsOptions []string
	for _, field := range *customFields {
		fieldsOptions = append(fieldsOptions, strings.Join([]string{field.Name, field.ID}, " | "))
	}

	var fieldsToAdd []string
	fieldsPrompt := &survey.MultiSelect{
		Message: "What field do you want to add:",
		Options: fieldsOptions,
	}

	if err = survey.AskOne(fieldsPrompt, &fieldsToAdd); err != nil {
		log.Fatal(err)
	}

	//Split the fields ID's
	var fieldIDs []string
	for _, fieldSelected := range fieldsToAdd {
		fieldIDs = append(fieldIDs, strings.Split(fieldSelected, " | ")[1])
	}

	log.Println(fieldIDs)

	project, _, err := atlassian.Project.Get(context.Background(), projectSelected, nil)
	if err != nil {
		log.Fatal(err)
	}

	var issueTypeNames []string
	for _, issueType := range project.IssueTypes {
		issueTypeNames = append(issueTypeNames, strings.Join([]string{issueType.Name, issueType.ID}, " | "))
	}

	var issueTypeNamesToAdd []string
	typePrompt := &survey.MultiSelect{
		Message: "What issue type do you want to use:",
		Options: issueTypeNames,
	}

	if err = survey.AskOne(typePrompt, &issueTypeNamesToAdd); err != nil {
		log.Fatal(err)
	}

	fmt.Println(issueTypeNamesToAdd)

	//Get issue type screen schemes for projects
	projectIDAsInt, err := strconv.Atoi(project.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(projectIDAsInt)
	//Get issue type schemes for projects
	//screenSchemes, _, err := atlassian.Issue.Type.ScreenScheme.Gets(context.Background(), []int{projectIDAsInt}, 0, 50)
}

func fetchProjects(atlassian *jira.Client, startAt int) (chunks []*jira.ProjectSearchScheme, err error) {

	options := &jira.ProjectSearchOptionsScheme{
		OrderBy: "issueCount",
		Action:  "edit",
	}

	projects, _, err := atlassian.Project.Search(context.Background(), options, startAt, 50)
	if err != nil {
		return nil, err
	}

	chunks = append(chunks, projects)

	if projects.IsLast {
		return
	} else {
		startAt = +50
		chunk, err := fetchProjects(atlassian, startAt)
		if err != nil {
			return nil, err
		}

		chunks = append(chunks, chunk...)
	}

	return
}
