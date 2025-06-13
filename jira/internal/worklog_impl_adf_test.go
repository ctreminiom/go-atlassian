package internal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalWorklogAdfImpl_Gets(t *testing.T) {

	payloadMocked := map[string]interface{}{"ids": []int{1, 2, 3, 4}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx        context.Context
		worklogIDs []int
		expand     []string
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
				worklogIDs: []int{1, 2, 3, 4},
				expand:     []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/worklog/list?expand=properties",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.Background(),
				worklogIDs: []int{1, 2, 3, 4},
				expand:     []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/worklog/list?expand=properties",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the worklogs ids are not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNpWorklogs,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.Background(),
				worklogIDs: []int{1, 2, 3, 4},
				expand:     []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/worklog/list?expand=properties",
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

			newService, err := NewWorklogADFService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Gets(testCase.args.ctx, testCase.args.worklogIDs, testCase.args.expand)

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

func Test_internalWorklogAdfImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                     context.Context
		issueKeyOrID, worklogID string
		expand                  []string
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				worklogID:    "493939",
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-5/worklog/493939?expand=properties",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogADFScheme{}).
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				worklogID:    "493939",
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-5/worklog/493939?expand=properties",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogADFScheme{}).
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the worklog id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
			},
			wantErr: true,
			Err:     model.ErrNoWorklogID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				worklogID:    "493939",
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-5/worklog/493939?expand=properties",
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

			newService, err := NewWorklogADFService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.worklogID, testCase.args.expand)

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

func Test_internalWorklogAdfImpl_Issue(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                        context.Context
		issueKeyOrID               string
		startAt, maxResults, after int
		expand                     []string
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				startAt:      0,
				maxResults:   50,
				after:        1661101991,
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-5/worklog?expand=properties&maxResults=50&startAt=0&startedAfter=1661101991",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogADFPageScheme{}).
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				startAt:      0,
				maxResults:   50,
				after:        1661101991,
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-5/worklog?expand=properties&maxResults=50&startAt=0&startedAfter=1661101991",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogADFPageScheme{}).
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				startAt:      0,
				maxResults:   50,
				after:        1661101991,
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-5/worklog?expand=properties&maxResults=50&startAt=0&startedAfter=1661101991",
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

			newService, err := NewWorklogADFService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Issue(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.startAt,
				testCase.args.maxResults, testCase.args.after, testCase.args.expand)

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

func Test_internalWorklogAdfImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                     context.Context
		issueKeyOrID, worklogID string
		options                 *model.WorklogOptionsScheme
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				worklogID:    "h837372",
				options: &model.WorklogOptionsScheme{
					Notify:               true,
					AdjustEstimate:       "new",
					NewEstimate:          "2d",
					ReduceBy:             "manual",
					OverrideEditableFlag: true,
					Expand:               []string{"properties"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issue/DUMMY-5/worklog/h837372?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
					"", nil).
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				worklogID:    "h837372",
				options: &model.WorklogOptionsScheme{
					Notify:               true,
					AdjustEstimate:       "new",
					NewEstimate:          "2d",
					ReduceBy:             "manual",
					OverrideEditableFlag: true,
					Expand:               []string{"properties"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issue/DUMMY-5/worklog/h837372?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
					"", nil).
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
			name:   "when the options are not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				worklogID:    "h837372",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issue/DUMMY-5/worklog/h837372",
					"", nil).
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the worklog id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
			},
			wantErr: true,
			Err:     model.ErrNoWorklogID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				worklogID:    "h837372",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issue/DUMMY-5/worklog/h837372",
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

			newService, err := NewWorklogADFService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.worklogID,
				testCase.args.options)

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
			}

		})
	}
}

func Test_internalWorklogAdfImpl_Deleted(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx   context.Context
		since int
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
				ctx:   context.Background(),
				since: 928281811,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/worklog/deleted?since=928281811",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ChangedWorklogPageScheme{}).
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
				ctx:   context.Background(),
				since: 928281811,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/worklog/deleted?since=928281811",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ChangedWorklogPageScheme{}).
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
				ctx:   context.Background(),
				since: 928281811,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/worklog/deleted?since=928281811",
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

			newService, err := NewWorklogADFService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Deleted(testCase.args.ctx, testCase.args.since)

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

func Test_internalWorklogAdfImpl_Updated(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx    context.Context
		since  int
		expand []string
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
				ctx:    context.Background(),
				since:  928281811,
				expand: []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/worklog/updated?expand=properties&since=928281811",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ChangedWorklogPageScheme{}).
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
				ctx:    context.Background(),
				since:  928281811,
				expand: []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/worklog/updated?expand=properties&since=928281811",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ChangedWorklogPageScheme{}).
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
				ctx:    context.Background(),
				since:  928281811,
				expand: []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/worklog/updated?expand=properties&since=928281811",
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

			newService, err := NewWorklogADFService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Updated(testCase.args.ctx, testCase.args.since, testCase.args.expand)

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

func Test_internalWorklogAdfImpl_Add(t *testing.T) {

	worklogCommentMocked := &model.CommentNodeScheme{
		Type:    "doc",
		Version: 1,
	}

	worklogCommentMocked.AppendNode(&model.CommentNodeScheme{
		Type: "paragraph",
		Content: []*model.CommentNodeScheme{
			{
				Type: "text",
				Text: "I did some work here!",
			},
		},
	})

	payloadMocked := &model.WorklogADFPayloadScheme{
		Comment: worklogCommentMocked,
		Visibility: &model.IssueWorklogVisibilityScheme{
			Type:  "group",
			Value: "jira-project-admins",
		},
		Started:          "2021-01-17T12:34:00.000+0000",
		TimeSpent:        "3h",
		TimeSpentSeconds: 12000,
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
		payload      *model.WorklogADFPayloadScheme
		options      *model.WorklogOptionsScheme
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				payload:      payloadMocked,
				options: &model.WorklogOptionsScheme{
					Notify:               true,
					AdjustEstimate:       "new",
					NewEstimate:          "2d",
					ReduceBy:             "manual",
					OverrideEditableFlag: true,
					Expand:               []string{"properties"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-5/worklog?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogADFScheme{}).
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				payload:      payloadMocked,
				options: &model.WorklogOptionsScheme{
					Notify:               true,
					AdjustEstimate:       "new",
					NewEstimate:          "2d",
					ReduceBy:             "manual",
					OverrideEditableFlag: true,
					Expand:               []string{"properties"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-5/worklog?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogADFScheme{}).
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				payload:      payloadMocked,
				options: &model.WorklogOptionsScheme{
					Notify:               true,
					AdjustEstimate:       "new",
					NewEstimate:          "2d",
					ReduceBy:             "manual",
					OverrideEditableFlag: true,
					Expand:               []string{"properties"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-5/worklog?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
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

			newService, err := NewWorklogADFService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Add(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.payload,
				testCase.args.options)

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

func Test_internalWorklogAdfImpl_Update(t *testing.T) {

	worklogCommentMocked := &model.CommentNodeScheme{
		Type:    "doc",
		Version: 1,
	}

	worklogCommentMocked.AppendNode(&model.CommentNodeScheme{
		Type: "paragraph",
		Content: []*model.CommentNodeScheme{
			{
				Type: "text",
				Text: "I did some work here!",
			},
		},
	})

	payloadMocked := &model.WorklogADFPayloadScheme{
		Comment: worklogCommentMocked,
		Visibility: &model.IssueWorklogVisibilityScheme{
			Type:  "group",
			Value: "jira-project-admins",
		},
		Started:          "2021-01-17T12:34:00.000+0000",
		TimeSpent:        "3h",
		TimeSpentSeconds: 12000,
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                     context.Context
		issueKeyOrID, worklogID string
		payload                 *model.WorklogADFPayloadScheme
		options                 *model.WorklogOptionsScheme
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				worklogID:    "3933828822",
				payload:      payloadMocked,
				options: &model.WorklogOptionsScheme{
					Notify:               true,
					AdjustEstimate:       "new",
					NewEstimate:          "2d",
					ReduceBy:             "manual",
					OverrideEditableFlag: true,
					Expand:               []string{"properties"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issue/DUMMY-5/worklog/3933828822?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogADFScheme{}).
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				worklogID:    "3933828822",
				payload:      payloadMocked,
				options: &model.WorklogOptionsScheme{
					Notify:               true,
					AdjustEstimate:       "new",
					NewEstimate:          "2d",
					ReduceBy:             "manual",
					OverrideEditableFlag: true,
					Expand:               []string{"properties"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/DUMMY-5/worklog/3933828822?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogADFScheme{}).
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the worklog id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
			},
			wantErr: true,
			Err:     model.ErrNoWorklogID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-5",
				worklogID:    "3933828822",
				payload:      payloadMocked,
				options: &model.WorklogOptionsScheme{
					Notify:               true,
					AdjustEstimate:       "new",
					NewEstimate:          "2d",
					ReduceBy:             "manual",
					OverrideEditableFlag: true,
					Expand:               []string{"properties"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issue/DUMMY-5/worklog/3933828822?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
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

			newService, err := NewWorklogADFService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Update(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.worklogID,
				testCase.args.payload, testCase.args.options)

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

func Test_NewWorklogADFService(t *testing.T) {

	type args struct {
		client  service.Connector
		version string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				client:  nil,
				version: "3",
			},
			wantErr: false,
		},

		{
			name: "when the version is not provided",
			args: args{
				client:  nil,
				version: "",
			},
			wantErr: true,
			Err:     model.ErrNoVersionProvided,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := NewWorklogADFService(testCase.args.client, testCase.args.version)

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
				assert.NotEqual(t, got, nil)
			}
		})
	}
}
