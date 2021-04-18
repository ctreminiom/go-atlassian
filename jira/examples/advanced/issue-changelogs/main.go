package main

/*
import (
	"context"
	"encoding/csv"
	"github.com/AlecAivazis/survey/v2"
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
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	jql := ""
	prompt := &survey.Input{
		Message: "Please provide the JQL query to use",
		Default: "project = KP",
	}

	if err = survey.AskOne(prompt, &jql); err != nil {
		log.Fatal(err)
	}

	var (
		searchPages []*jira.IssueSearchScheme
		startAt     = 0
		fields      = []string{"status"}
		expand      = []string{"changelog"}
	)

	for {

		chunk, _, err := atlassian.Issue.Search.Get(context.Background(), jql, fields, expand, startAt, 50, "")
		if err != nil {
			log.Fatal(err)
		}

		searchPages = append(searchPages, chunk)

		if len(chunk.Issues) == 0 {
			break
		}

		startAt += 50
	}

	var records []changelogRecord
	for _, page := range searchPages {

		for _, issue := range page.Issues {

			for _, history := range issue.Changelog.Histories {

				for _, item := range history.Items {

					records = append(records, changelogRecord{
						IssueKey:   issue.Key,
						AuthorMail: history.Author.EmailAddress,
						Field:      item.Field,
						From:       item.From,
						To:         item.To,
						FieldType:  item.Fieldtype,
					})
				}
			}
		}
	}

	//Create the .csv file
	_ = os.Remove("issue-changelogs.csv")

	csvFile, err := os.Create("issue-changelogs.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	//Write the csv headers
	if err = writer.Write([]string{"issue-key", "author-mail", "field", "field-type", "from", "to"}); err != nil {
		log.Fatal(err)
	}

	//Write the records extracted
	for _, record := range records {

		if err = writer.Write([]string{
			record.IssueKey,
			record.AuthorMail,
			record.Field,
			record.FieldType, record.From,
			record.To,
		}); err != nil {
			log.Fatal(err)
		}
	}

}

type changelogRecord struct {
	IssueKey   string
	AuthorMail string
	Field      string
	FieldType  string
	From       string
	To         string
}
*/
