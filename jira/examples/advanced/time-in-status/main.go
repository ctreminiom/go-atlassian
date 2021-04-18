package main

/*

import (
	"context"
	"github.com/AlecAivazis/survey/v2"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/ctreminiom/go-atlassian/jira"
	"github.com/hako/durafmt"
	"log"
	"os"
	"sort"
	"time"
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
		Default: "order by created desc",
	}

	if err = survey.AskOne(prompt, &jql); err != nil {
		log.Fatal(err)
	}

	var (
		searchPages []*jira.IssueSearchScheme
		startAt     = 0
		fields      = []string{"created"}
		expand      = []string{"changelog"}
	)

	//status

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

					if item.Field == "status" {

						whenAsTime, err := time.Parse(jira.DateFormatJira, history.Created)
						if err != nil {
							log.Fatal(err)
						}

						createdAsTime, err := time.Parse(jira.DateFormatJira, issue.Fields.Created)
						if err != nil {
							log.Fatal(err)
						}

						records = append(records, changelogRecord{
							IssueKey:           issue.Key,
							AuthorMail:         history.Author.EmailAddress,
							Field:              item.Field,
							FieldType:          item.Fieldtype,
							From:               item.From,
							To:                 item.To,
							When:               history.Created,
							Fieldtype:          item.Fieldtype,
							FieldID:            item.FieldID,
							FromString:         item.FromString,
							ToString:           item.ToString,
							WhenAsTime:         whenAsTime,
							IssueCreatedAsTime: createdAsTime,
						})

					}

				}
			}
		}
	}

	//group the records by issue key
	var recordsByIssue = make(map[string][]changelogRecord)
	for _, record := range records {
		recordsByIssue[record.IssueKey] = append(recordsByIssue[record.IssueKey], record)
	}

	var csvRows []csvRow
	for issue, records := range recordsByIssue {

		sort.Slice(records, func(i, j int) bool {
			return records[i].WhenAsTime.Before(records[j].WhenAsTime)
		})

		for index, record := range records {

			if index == 0 {

				currentTime := records[index].WhenAsTime //record.WhenAsTime
				diff := currentTime.Sub(record.IssueCreatedAsTime)

				csvRows = append(csvRows, csvRow{
					Key:     issue,
					Status:  record.FromString,
					Pretty:  durafmt.ParseShort(diff).LimitFirstN(4).String(),
					Days:    diff.Hours() / 24,
					Hours:   diff.Hours(),
					Minutes: diff.Minutes(),
					Seconds: diff.Seconds(),
				})

				log.Println(index, durafmt.ParseShort(diff).LimitFirstN(4).String())

			}

			if index == len(records)-1 {

				currentTime := time.Now()
				diff := currentTime.Sub(record.WhenAsTime)

				csvRows = append(csvRows, csvRow{
					Key:     issue,
					Status:  record.ToString,
					Pretty:  durafmt.ParseShort(diff).LimitFirstN(4).String(),
					Days:    diff.Hours() / 24,
					Hours:   diff.Hours(),
					Minutes: diff.Minutes(),
					Seconds: diff.Seconds(),
				})

				log.Println(index == len(records)-1, durafmt.ParseShort(diff).LimitFirstN(4).String())

				break
			}

			currentTime := records[index+1].WhenAsTime //record.WhenAsTime
			diff := currentTime.Sub(record.WhenAsTime)

			csvRows = append(csvRows, csvRow{
				Key:     issue,
				Status:  record.ToString,
				Pretty:  durafmt.ParseShort(diff).LimitFirstN(4).String(),
				Days:    diff.Hours() / 24,
				Hours:   diff.Hours(),
				Minutes: diff.Minutes(),
				Seconds: diff.Seconds(),
			})

			log.Println(index, durafmt.ParseShort(diff).LimitFirstN(4).String())

		}

	}

	//Create the .csv file
	_ = os.Remove("time-in-status.csv")

	csvFile, err := os.Create("time-in-status.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	//Write the csv headers
	if err = writer.Write([]string{"issue-key", "status", "pretty", "days", "hours", "minutes", "seconds"}); err != nil {
		log.Fatal(err)
	}

	for _, row := range csvRows {

		if err = writer.Write([]string{
			row.Key,
			row.Status,
			row.Pretty,
			fmt.Sprintf("%.2f", row.Days),
			fmt.Sprintf("%.2f", row.Hours),
			fmt.Sprintf("%.2f", row.Minutes),
			fmt.Sprintf("%.2f", row.Seconds),
		}); err != nil {
			log.Fatal(err)
		}

	}

	writer.Flush()

}

type changelogRecord struct {
	IssueCreatedAsTime time.Time
	IssueKey           string
	AuthorMail         string
	Field              string
	FieldType          string
	From               string
	To                 string
	When               string
	WhenAsTime         time.Time
	Fieldtype          string
	FieldID            string
	FromString         string
	ToString           string
}

type csvRow struct {
	Key     string
	Status  string
	Pretty  string
	Days    float64
	Hours   float64
	Minutes float64
	Seconds float64
}

*/
