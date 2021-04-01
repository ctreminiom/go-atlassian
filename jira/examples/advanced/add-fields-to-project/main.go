package main

import (
	"context"
	"encoding/json"
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
		3. Select the issue types you want to change (assuming each issue type has their own screen scheme)
		4. Map the screen scheme with the issue types
		5. For each issue type
		5.1. Extract the screen ID's (view, create, edit, default)
		5.1. For each screen found
		5.1.1 Extract the default tab ID
		5.1.2 Add the customfield ID on the default tab
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
		Message: "What issue types do you want to use:",
		Options: issueTypeNames,
	}

	if err = survey.AskOne(typePrompt, &issueTypeNamesToAdd); err != nil {
		log.Fatal(err)
	}

	//Get issue type screen schemes for projects
	projectIDAsInt, err := strconv.Atoi(project.ID)
	if err != nil {
		log.Fatal(err)
	}

	screenSchemes, _, err := atlassian.Issue.Type.ScreenScheme.Projects(context.Background(), []int{projectIDAsInt}, 0, 50)
	if err != nil {
		log.Fatal(err)
	}

	if screenSchemes.Total != 1 {
		log.Fatalf("error!, the project %v does not have an IssueType Screen Scheme associated", project.Key)
	}

	projectIssueTypeScreenScheme := screenSchemes.Values[0].IssueTypeScreenScheme.ID

	projectIssueTypeScreenSchemeAsInt, err := strconv.Atoi(projectIssueTypeScreenScheme)
	if err != nil {
		log.Fatal(err)
	}

	projectScreenScreenMapping, _, err := atlassian.Issue.Type.ScreenScheme.Mapping(context.Background(), []int{projectIssueTypeScreenSchemeAsInt}, 0, 50)
	if err != nil {
		log.Fatal(err)
	}

	//group the screen schemes with the project issue types
	var issueTypeRelationship = make(map[string]string)
	for _, mapping := range projectScreenScreenMapping.Values {

		log.Printf("ScreenSchemeID = %v - IssueTypeID = %v", mapping.ScreenSchemeID, mapping.IssueTypeID)
		issueTypeRelationship[mapping.IssueTypeID] = mapping.ScreenSchemeID
	}

	//Split the issue types ID's
	var issueTypeIDs []string
	for _, issueType := range issueTypeNamesToAdd {
		issueTypeIDs = append(issueTypeIDs, strings.Split(issueType, " | ")[1])
	}

	//For each map, process the issue types
	for _, screenScheme := range issueTypeRelationship {

		if err = addFieldToScreen(atlassian, screenScheme, project.Key, fieldIDs); err != nil {
			log.Fatal(err)
		}

	}

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

func addFieldToScreen(atlassian *jira.Client, screenSchemeID, projectKey string, fieldIDs []string) (err error) {

	screenSchemeIDAsInt, err := strconv.Atoi(screenSchemeID)
	if err != nil {
		return err
	}

	//The Screen Schemes have screens associated,
	//we need to extract those IDs
	screenSchemes, _, err := atlassian.Screen.Scheme.Gets(context.Background(), []int{screenSchemeIDAsInt}, 0, 50)
	if err != nil {
		return err
	}

	for _, screenScheme := range screenSchemes.Values {

		// Group the Screen IDs
		var screenIDs []int
		screenIDs = append(screenIDs, screenScheme.Screens.Create, screenScheme.Screens.View, screenScheme.Screens.Default, screenScheme.Screens.Edit)

		for _, screenID := range removeDuplicateValues(screenIDs) {

			if screenID == 0 {
				continue
			}

			for _, fieldID := range fieldIDs {

				//Get the default screen tab
				tabs, _, err := atlassian.Screen.Tab.Gets(context.Background(), screenID, projectKey)
				if err != nil {
					return err
				}

				var tabID int
				for index, tab := range *tabs {

					if index == 0 {
						tabID = tab.ID
					}
				}

				_, response, err := atlassian.Screen.Tab.Field.Add(context.Background(), screenID, tabID, fieldID)
				if err != nil {
					if response != nil {

						if response.StatusCode == 400 {

							var apiErrorResponse map[string]interface{}
							if err := json.Unmarshal(response.BodyAsBytes, &apiErrorResponse); err != nil {
								log.Fatal(err)
							}

							log.Println(apiErrorResponse["errors"].(map[string]interface{})["fieldId"])
							continue
						}

					}
					log.Fatal(err)
				}

				log.Printf("the field %v has been added on the screen %v", fieldID, screenID)
			}
		}
	}

	return
}

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	var list []int

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
