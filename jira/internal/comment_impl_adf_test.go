package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalAdfCommentImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrId string
		orderBy      string
		expand       []string
		startAt      int
		maxResults   int
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the document format is adf (atlassian document format)",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				orderBy:      "id",
				expand:       []string{"renderedBody"},
				startAt:      0,
				maxResults:   50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-1/comment?expand=renderedBody&maxResults=50&orderBy=id&startAt=0",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueCommentPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
				orderBy:      "id",
				expand:       []string{"renderedBody"},
				startAt:      0,
				maxResults:   50,
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				orderBy:      "id",
				expand:       []string{"renderedBody"},
				startAt:      0,
				maxResults:   50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-1/comment?expand=renderedBody&maxResults=50&orderBy=id&startAt=0",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			commentService, _, err := NewCommentService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := commentService.Gets(testCase.args.ctx, testCase.args.issueKeyOrId,
				testCase.args.orderBy, testCase.args.expand, testCase.args.startAt, testCase.args.maxResults)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalAdfCommentImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                     context.Context
		issueKeyOrId, commentId string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the document format is adf (atlassian document format)",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				commentId:    "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-1/comment/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueCommentScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the comment id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
			},
			wantErr: true,
			Err:     model.ErrNoCommentIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				commentId:    "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-1/comment/10001",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			commentService, _, err := NewCommentService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := commentService.Get(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.commentId)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalAdfCommentImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                     context.Context
		issueKeyOrId, commentId string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the document format is adf (atlassian document format)",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				commentId:    "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issue/DUMMY-1/comment/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the comment id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
			},
			wantErr: true,
			Err:     model.ErrNoCommentIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				commentId:    "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issue/DUMMY-1/comment/10001",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			commentService, _, err := NewCommentService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := commentService.Delete(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.commentId)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}

		})
	}
}

func Test_internalAdfCommentImpl_Add(t *testing.T) {

	commentBody := model.CommentNodeScheme{}
	commentBody.Version = 1
	commentBody.Type = "doc"

	commentBody.AppendNode(&model.CommentNodeScheme{
		Type: "paragraph",
		Content: []*model.CommentNodeScheme{
			{
				Type: "text",
				Text: "Carlos Test",
				Marks: []*model.MarkScheme{
					{
						Type: "strong",
					},
				},
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

	payloadMocked := &model.CommentPayloadScheme{
		Visibility: &model.CommentVisibilityScheme{
			Type:  "role",
			Value: "Administrators",
		},
		Body: &commentBody,
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrId string
		payload      *model.CommentPayloadScheme
		expand       []string
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "when the document format is adf (atlassian document format)",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				payload:      payloadMocked,
				expand:       []string{"body"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-1/comment?expand=body",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueCommentScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
				payload:      payloadMocked,
				expand:       []string{"body"},
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				payload:      payloadMocked,
				expand:       []string{"body"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-1/comment?expand=body",
					"",
					payloadMocked).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			commentService, _, err := NewCommentService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := commentService.Add(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.payload,
				testCase.args.expand)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}
