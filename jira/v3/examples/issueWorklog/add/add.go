package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira/v3"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := v3.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	// Comment worklog
	commentBody := v3.CommentNodeScheme{}
	commentBody.Version = 1
	commentBody.Type = "doc"

	//Create the Tables Headers
	tableHeaders := &v3.CommentNodeScheme{
		Type: "tableRow",
		Content: []*v3.CommentNodeScheme{

			{
				Type: "tableHeader",
				Content: []*v3.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*v3.CommentNodeScheme{
							{
								Type: "text",
								Text: "Header 1",
								Marks: []*v3.MarkScheme{
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
				Content: []*v3.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*v3.CommentNodeScheme{
							{
								Type: "text",
								Text: "Header 2",
								Marks: []*v3.MarkScheme{
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
				Content: []*v3.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*v3.CommentNodeScheme{
							{
								Type: "text",
								Text: "Header 3",
								Marks: []*v3.MarkScheme{
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

	row1 := &v3.CommentNodeScheme{
		Type: "tableRow",
		Content: []*v3.CommentNodeScheme{
			{
				Type: "tableCell",
				Content: []*v3.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*v3.CommentNodeScheme{
							{Type: "text", Text: "Row 00"},
						},
					},
				},
			},

			{
				Type: "tableCell",
				Content: []*v3.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*v3.CommentNodeScheme{
							{Type: "text", Text: "Row 01"},
						},
					},
				},
			},

			{
				Type: "tableCell",
				Content: []*v3.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*v3.CommentNodeScheme{
							{Type: "text", Text: "Row 02"},
						},
					},
				},
			},
		},
	}

	row2 := &v3.CommentNodeScheme{
		Type: "tableRow",
		Content: []*v3.CommentNodeScheme{
			{
				Type: "tableCell",
				Content: []*v3.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*v3.CommentNodeScheme{
							{Type: "text", Text: "Row 10"},
						},
					},
				},
			},

			{
				Type: "tableCell",
				Content: []*v3.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*v3.CommentNodeScheme{
							{Type: "text", Text: "Row 11"},
						},
					},
				},
			},

			{
				Type: "tableCell",
				Content: []*v3.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*v3.CommentNodeScheme{
							{Type: "text", Text: "Row 12"},
						},
					},
				},
			},
		},
	}

	commentBody.AppendNode(&v3.CommentNodeScheme{
		Type:    "table",
		Attrs:   map[string]interface{}{"isNumberColumnEnabled": false, "layout": "default"},
		Content: []*v3.CommentNodeScheme{tableHeaders, row1, row2},
	})

	var options = &v3.WorklogOptionsScheme{
		Notify:         true,
		AdjustEstimate: "auto",
		ReduceBy:       "3h",
		//OverrideEditableFlag: true,
		Expand: []string{"expand", "properties"},
		Payload: &v3.WorklogPayloadScheme{
			Comment: &commentBody,
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

	worklog, response, err := atlassian.Issue.Worklog.Add(context.Background(), "KP-1", options)
	if err != nil {
		log.Println(response.Endpoint, response.Code)
		log.Fatal(err)
	}

	log.Println(response.Endpoint, response.Code)
	log.Println(worklog.ID, worklog.IssueID)
}
