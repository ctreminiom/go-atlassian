package models

import "testing"

func TestCommentNodeScheme_AppendNode(t *testing.T) {
	type fields struct {
		Version int
		Type    string
		Content []*CommentNodeScheme
		Text    string
		Attrs   map[string]interface{}
		Marks   []*MarkScheme
	}
	type args struct {
		node *CommentNodeScheme
	}
	testCases := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				node: &CommentNodeScheme{
					Type: "text",
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			n := &CommentNodeScheme{
				Version: testCase.fields.Version,
				Type:    testCase.fields.Type,
				Content: testCase.fields.Content,
				Text:    testCase.fields.Text,
				Attrs:   testCase.fields.Attrs,
				Marks:   testCase.fields.Marks,
			}
			n.AppendNode(testCase.args.node)
		})
	}
}
