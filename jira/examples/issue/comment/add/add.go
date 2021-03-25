package main

import (
	"context"
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
		Type:    "table",
		Attrs:   map[string]interface{}{"isNumberColumnEnabled": false, "layout": "default"},
		Content: []*jira.CommentNodeScheme{tableHeaders, row1, row2},
	})

	/*
		commentBody.AppendNode(&jira.CommentNodeScheme{
			Type: "paragraph",
			Content: []*jira.CommentNodeScheme{
				{
					Type: "text",
					Text: "Carlos Test",
				},
				{
					Type: "emoji",
					Attrs: map[string]interface{}{
						"shortName": ":grin",
						"id":        "1f601",
						"text":      "😁",
					},
				},
				{
					Type: "text",
					Text: " ",
				},
			},
		})
	*/

	newComment, response, err := atlassian.Issue.Comment.Add(context.Background(), "KP-2", "role", "Administrators", &commentBody, nil)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println(newComment.ID)
}
