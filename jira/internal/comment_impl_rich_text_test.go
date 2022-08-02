package internal

import (
	"bytes"
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalRichTextCommentImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Client
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
			name:   "when the document format is rich-text",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				orderBy:      "id",
				expand:       []string{"renderedBody"},
				startAt:      0,
				maxResults:   50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-1/comment?expand=renderedBody&maxResults=50&orderBy=id&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueCommentPageSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "2"},
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
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				orderBy:      "id",
				expand:       []string{"renderedBody"},
				startAt:      0,
				maxResults:   50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-1/comment?expand=renderedBody&maxResults=50&orderBy=id&startAt=0",
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

			_, commentService, err := NewCommentService(testCase.fields.c, testCase.fields.version)
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

func Test_internalRichTextCommentImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Client
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
			name:   "when the document format is rich-text",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				commentId:    "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-1/comment/10001",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueCommentSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the comment id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
			},
			wantErr: true,
			Err:     model.ErrNoCommentIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				commentId:    "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-1/comment/10001",
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

			_, commentService, err := NewCommentService(testCase.fields.c, testCase.fields.version)
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

func Test_internalRichTextCommentImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Client
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
			name:   "when the document format is rich-text",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				commentId:    "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issue/DUMMY-1/comment/10001",
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
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the comment id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
			},
			wantErr: true,
			Err:     model.ErrNoCommentIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				commentId:    "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issue/DUMMY-1/comment/10001",
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

			_, commentService, err := NewCommentService(testCase.fields.c, testCase.fields.version)
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

func Test_internalRichTextCommentImpl_Add(t *testing.T) {

	payloadMocked := &model.CommentPayloadSchemeV2{
		Visibility: &model.CommentVisibilityScheme{
			Type:  "role",
			Value: "Administrators",
		},
		Body: "test case",
	}

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrId string
		payload      *model.CommentPayloadSchemeV2
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
			name:   "when the document format is rich-text",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				payload:      payloadMocked,
				expand:       []string{"body"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/comment?expand=body",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueCommentSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "2"},
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
			name:   "when the payload is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				payload:      nil,
				expand:       []string{"body"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.CommentPayloadSchemeV2)(nil)).
					Return(bytes.NewReader([]byte{}), model.ErrNilPayloadError)
				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrNilPayloadError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-1",
				payload:      payloadMocked,
				expand:       []string{"body"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/comment?expand=body",
					bytes.NewReader([]byte{})).
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

			_, commentService, err := NewCommentService(testCase.fields.c, testCase.fields.version)
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
