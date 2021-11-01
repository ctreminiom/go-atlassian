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
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

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

	payload := &v3.CommentPayloadScheme{
		Visibility: &v3.CommentVisibilityScheme{
			Type:  "role",
			Value: "Administrators",
		},
		Body: &commentBody,
	}

	newComment, response, err := atlassian.Issue.Comment.Add(context.Background(), "KP-2", payload, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(newComment.ID)
}
