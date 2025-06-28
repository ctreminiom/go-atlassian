package internal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalSearchRichTextImpl_Checks(t *testing.T) {

	payloadMocked := &model.IssueSearchCheckPayloadScheme{
		IssueIDs: []int{10001, 1000, 10042},
		JQLs: []string{
			"project = FOO",
			"issuetype = Bug",
			"summary ~ \\\"some text\\\" AND project in (FOO, BAR)",
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.IssueSearchCheckPayloadScheme
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/jql/match",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueMatchesPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/jql/match",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueMatchesPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/jql/match",
					"", payloadMocked).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			_, newService, err := NewSearchService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Checks(testCase.args.ctx, testCase.args.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalSearchRichTextImpl_Post(t *testing.T) {

	payloadMocked := struct {
		Expand        []string "json:\"expand,omitempty\""
		Jql           string   "json:\"jql,omitempty\""
		MaxResults    int      "json:\"maxResults,omitempty\""
		Fields        []string "json:\"fields,omitempty\""
		StartAt       int      "json:\"startAt,omitempty\""
		ValidateQuery string   "json:\"validateQuery,omitempty\""
	}{Expand: []string{"operations"}, Jql: "project = FOO", MaxResults: 100, Fields: []string{"status", "summary"}, StartAt: 50, ValidateQuery: "strict"}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		jql                 string
		fields, expands     []string
		startAt, maxResults int
		validate            string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.Background(),
				jql:        "project = FOO",
				fields:     []string{"status", "summary"},
				expands:    []string{"operations"},
				startAt:    50,
				maxResults: 100,
				validate:   "strict",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/search",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueSearchSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.Background(),
				jql:        "project = FOO",
				fields:     []string{"status", "summary"},
				expands:    []string{"operations"},
				startAt:    50,
				maxResults: 100,
				validate:   "strict",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/search",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueSearchSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.Background(),
				jql:        "project = FOO",
				fields:     []string{"status", "summary"},
				expands:    []string{"operations"},
				startAt:    50,
				maxResults: 100,
				validate:   "strict",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/search",
					"", payloadMocked).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			_, newService, err := NewSearchService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Post(testCase.args.ctx, testCase.args.jql, testCase.args.fields,
				testCase.args.expands, testCase.args.startAt, testCase.args.maxResults, testCase.args.validate)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalSearchRichTextImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		jql                 string
		fields, expands     []string
		startAt, maxResults int
		validate            string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.Background(),
				jql:        "project = FOO",
				fields:     []string{"status", "summary"},
				expands:    []string{"operations"},
				startAt:    50,
				maxResults: 100,
				validate:   "strict",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/search?expand=operations&fields=status%2Csummary&jql=project+%3D+FOO&maxResults=100&startAt=50&validateQuery=strict",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueSearchSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.Background(),
				jql:        "project = FOO",
				fields:     []string{"status", "summary"},
				expands:    []string{"operations"},
				startAt:    50,
				maxResults: 100,
				validate:   "strict",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/search?expand=operations&fields=status%2Csummary&jql=project+%3D+FOO&maxResults=100&startAt=50&validateQuery=strict",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueSearchSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.Background(),
				jql:        "project = FOO",
				fields:     []string{"status", "summary"},
				expands:    []string{"operations"},
				startAt:    50,
				maxResults: 100,
				validate:   "strict",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/search?expand=operations&fields=status%2Csummary&jql=project+%3D+FOO&maxResults=100&startAt=50&validateQuery=strict",
					"", nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			_, newService, err := NewSearchService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.jql, testCase.args.fields,
				testCase.args.expands, testCase.args.startAt, testCase.args.maxResults, testCase.args.validate)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalSearchRichTextImpl_SearchJQL(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx           context.Context
		jql           string
		fields        []string
		expands       []string
		maxResults    int
		nextPageToken string
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
			name:   "when the search jql operation is successful",
			fields: fields{version: "2"},
			args: args{
				ctx:           context.Background(),
				jql:           "project = FOO",
				fields:        []string{"summary", "status"},
				expands:       []string{"changelog", "names"},
				maxResults:    50,
				nextPageToken: "CAEaAggD",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				payload := struct {
					Jql           string   `json:"jql,omitempty"`
					MaxResults    int      `json:"maxResults,omitempty"`
					Fields        []string `json:"fields,omitempty"`
					Expand        string   `json:"expand,omitempty"`
					NextPageToken string   `json:"nextPageToken,omitempty"`
				}{
					Jql:           "project = FOO",
					MaxResults:    50,
					Fields:        []string{"summary", "status"},
					Expand:        "changelog,names",
					NextPageToken: "CAEaAggD",
				}

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/search/jql",
					"", payload).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueSearchJQLSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the request returns an error",
			fields: fields{version: "2"},
			args: args{
				ctx:           context.Background(),
				jql:           "project = FOO",
				fields:        []string{"summary", "status"},
				expands:       []string{"changelog", "names"},
				maxResults:    50,
				nextPageToken: "CAEaAggD",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				payload := struct {
					Jql           string   `json:"jql,omitempty"`
					MaxResults    int      `json:"maxResults,omitempty"`
					Fields        []string `json:"fields,omitempty"`
					Expand        string   `json:"expand,omitempty"`
					NextPageToken string   `json:"nextPageToken,omitempty"`
				}{
					Jql:           "project = FOO",
					MaxResults:    50,
					Fields:        []string{"summary", "status"},
					Expand:        "changelog,names",
					NextPageToken: "CAEaAggD",
				}

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/search/jql",
					"", payload).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			searchImpl := &internalSearchRichTextImpl{
				c:       testCase.fields.c,
				version: testCase.fields.version,
			}

			gotResult, gotResponse, err := searchImpl.SearchJQL(
				testCase.args.ctx,
				testCase.args.jql,
				testCase.args.fields,
				testCase.args.expands,
				testCase.args.maxResults,
				testCase.args.nextPageToken,
			)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})
	}
}

func Test_internalSearchRichTextImpl_ApproximateCount(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx context.Context
		jql string
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
			name:   "when the approximate count operation is successful",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				jql: "project = FOO",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				payload := struct {
					Jql string `json:"jql,omitempty"`
				}{
					Jql: "project = FOO",
				}

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/search/approximate-count",
					"", payload).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueSearchApproximateCountScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the request returns an error",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				jql: "project = FOO",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				payload := struct {
					Jql string `json:"jql,omitempty"`
				}{
					Jql: "project = FOO",
				}

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/search/approximate-count",
					"", payload).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			searchImpl := &internalSearchRichTextImpl{
				c:       testCase.fields.c,
				version: testCase.fields.version,
			}

			gotResult, gotResponse, err := searchImpl.ApproximateCount(
				testCase.args.ctx,
				testCase.args.jql,
			)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})
	}
}

func Test_internalSearchRichTextImpl_BulkFetch(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx            context.Context
		issueIDsOrKeys []string
		fields         []string
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
			name:   "when the bulk fetch operation is successful",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueIDsOrKeys: []string{"FOO-1", "10067", "BAR-1"},
				fields:         []string{"summary", "status", "priority"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				payload := struct {
					IssueIDsOrKeys []string `json:"issueIdsOrKeys,omitempty"`
					Fields         []string `json:"fields,omitempty"`
				}{
					IssueIDsOrKeys: []string{"FOO-1", "10067", "BAR-1"},
					Fields:         []string{"summary", "status", "priority"},
				}

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/bulkfetch",
					"", payload).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueBulkFetchSchemeV2{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the request returns an error",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueIDsOrKeys: []string{"FOO-1", "10067", "BAR-1"},
				fields:         []string{"summary", "status", "priority"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				payload := struct {
					IssueIDsOrKeys []string `json:"issueIdsOrKeys,omitempty"`
					Fields         []string `json:"fields,omitempty"`
				}{
					IssueIDsOrKeys: []string{"FOO-1", "10067", "BAR-1"},
					Fields:         []string{"summary", "status", "priority"},
				}

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/bulkfetch",
					"", payload).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			searchImpl := &internalSearchRichTextImpl{
				c:       testCase.fields.c,
				version: testCase.fields.version,
			}

			gotResult, gotResponse, err := searchImpl.BulkFetch(
				testCase.args.ctx,
				testCase.args.issueIDsOrKeys,
				testCase.args.fields,
			)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})
	}
}
