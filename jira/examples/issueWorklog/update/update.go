package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main()  {

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

	// Comment worklog
	commentBody := jira.CommentNodeScheme{}
	commentBody.Version = 1
	commentBody.Type = "doc"

	//Create the Tables Headers
	tableHeaders := &jira.CommentNodeScheme{
		Type: "tableRow",
		Content: []*jira.CommentNodeScheme{

			{
				Type: "tableHeader",
				Content: []*jira.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*jira.CommentNodeScheme{
							{
								Type: "text",
								Text: "Header 1",
								Marks: []*jira.MarkScheme{
									{
										Type: "strong",
									},
								},
							},
						},
					},

				},
			},

			{
				Type: "tableHeader",
				Content: []*jira.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*jira.CommentNodeScheme{
							{
								Type: "text",
								Text: "Header 2",
								Marks: []*jira.MarkScheme{
									{
										Type: "strong",
									},
								},
							},
						},
					},

				},
			},

			{
				Type: "tableHeader",
				Content: []*jira.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*jira.CommentNodeScheme{
							{
								Type: "text",
								Text: "Header 3",
								Marks: []*jira.MarkScheme{
									{
										Type: "strong",
									},
								},
							},
						},
					},

				},
			},

		},

	}

	row1 := &jira.CommentNodeScheme{
		Type: "tableRow",
		Content: []*jira.CommentNodeScheme{
			{
				Type: "tableCell",
				Content: []*jira.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*jira.CommentNodeScheme{
							{Type: "text", Text: "Row 00"},
						},
					},
				},
			},

			{
				Type: "tableCell",
				Content: []*jira.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*jira.CommentNodeScheme{
							{Type: "text", Text: "Row 01"},
						},
					},
				},
			},

			{
				Type: "tableCell",
				Content: []*jira.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*jira.CommentNodeScheme{
							{Type: "text", Text: "Row 02"},
						},
					},
				},
			},

		},

	}

	row2 := &jira.CommentNodeScheme{
		Type: "tableRow",
		Content: []*jira.CommentNodeScheme{
			{
				Type: "tableCell",
				Content: []*jira.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*jira.CommentNodeScheme{
							{Type: "text", Text: "Row 10"},
						},
					},
				},
			},

			{
				Type: "tableCell",
				Content: []*jira.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*jira.CommentNodeScheme{
							{Type: "text", Text: "Row 11"},
						},
					},
				},
			},

			{
				Type: "tableCell",
				Content: []*jira.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*jira.CommentNodeScheme{
							{Type: "text", Text: "Row 12"},
						},
					},
				},
			},

		},

	}

	commentBody.AppendNode(&jira.CommentNodeScheme{
		Type: "table",
		Attrs: map[string]interface{}{"isNumberColumnEnabled": false, "layout": "default"},
		Content: []*jira.CommentNodeScheme{tableHeaders, row1, row2},
	})

	var options = &jira.WorklogOptionsScheme{
		Notify:               true,
		AdjustEstimate:       "auto",
		ReduceBy:             "3h",
		//OverrideEditableFlag: true,
		Expand:               []string{"expand", "properties"},
		Payload:              &jira.WorklogPayloadScheme{
			Comment:          &commentBody,
			/*
				Visibility:       &jira.IssueWorklogVisibilityScheme{
					Type:  "group",
					Value: "jira-users",
				},
			*/
			Started:          "2021-07-16T07:01:10.774+0000",
			TimeSpentSeconds: 12000,
		},
	}

	worklog, response, err := atlassian.Issue.Worklog.Update(context.Background(), "KP-1", "10000",options)
	if err != nil {
		log.Println(response.Endpoint, response.Code)
		log.Fatal(err)
	}

	log.Println(response.Endpoint, response.Code)
	log.Println(worklog.ID, worklog.IssueID)

}
