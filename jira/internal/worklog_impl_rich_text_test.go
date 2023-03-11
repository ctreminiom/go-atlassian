package internal

import (
	"bytes"
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func Test_internalWorklogRichTextImpl_Gets(t *testing.T) {

	payloadMocked := &struct {
		Ids []int "json:\"ids\""
	}{Ids: []int{1, 2, 3, 4}}

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx        context.Context
		worklogIds []int
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
				ctx:        context.TODO(),
				worklogIds: []int{1, 2, 3, 4},
				expand:     []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/worklog/list?expand=properties",
					bytes.NewReader([]byte{})).
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
				ctx:        context.TODO(),
				worklogIds: []int{1, 2, 3, 4},
				expand:     []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/worklog/list?expand=properties",
					bytes.NewReader([]byte{})).
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
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNpWorklogsError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.TODO(),
				worklogIds: []int{1, 2, 3, 4},
				expand:     []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/worklog/list?expand=properties",
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

			newService, err := NewWorklogRichTextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Gets(testCase.args.ctx, testCase.args.worklogIds, testCase.args.expand)

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

func Test_internalWorklogRichTextImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                     context.Context
		issueKeyOrId, worklogId string
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
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
				worklogId:    "493939",
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-5/worklog/493939?expand=properties",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogScheme{}).
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
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
				worklogId:    "493939",
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-5/worklog/493939?expand=properties",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogScheme{}).
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
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the worklog id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
			},
			wantErr: true,
			Err:     model.ErrNoWorklogIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
				worklogId:    "493939",
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-5/worklog/493939?expand=properties",
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

			newService, err := NewWorklogRichTextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.worklogId, testCase.args.expand)

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

func Test_internalWorklogRichTextImpl_Issue(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                        context.Context
		issueKeyOrId               string
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
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
				startAt:      0,
				maxResults:   50,
				after:        1661101991,
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-5/worklog?expand=properties&maxResults=50&startAt=0&startedAfter=1661101991",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogPageScheme{}).
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
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
				startAt:      0,
				maxResults:   50,
				after:        1661101991,
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-5/worklog?expand=properties&maxResults=50&startAt=0&startedAfter=1661101991",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogPageScheme{}).
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
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
				startAt:      0,
				maxResults:   50,
				after:        1661101991,
				expand:       []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-5/worklog?expand=properties&maxResults=50&startAt=0&startedAfter=1661101991",
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

			newService, err := NewWorklogRichTextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Issue(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.startAt,
				testCase.args.maxResults, testCase.args.after, testCase.args.expand)

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

func Test_internalWorklogRichTextImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                     context.Context
		issueKeyOrId, worklogId string
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
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
				worklogId:    "h837372",
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

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issue/DUMMY-5/worklog/h837372?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
				worklogId:    "h837372",
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

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issue/DUMMY-5/worklog/h837372?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
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
			name:   "when the options are not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
				worklogId:    "h837372",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issue/DUMMY-5/worklog/h837372",
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
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the worklog id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
			},
			wantErr: true,
			Err:     model.ErrNoWorklogIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrId: "DUMMY-5",
				worklogId:    "h837372",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issue/DUMMY-5/worklog/h837372",
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

			newService, err := NewWorklogRichTextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.worklogId,
				testCase.args.options)

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

func Test_internalWorklogRichTextImpl_Deleted(t *testing.T) {

	type fields struct {
		c       service.Client
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
				ctx:   context.TODO(),
				since: 928281811,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/worklog/deleted?since=928281811",
					nil).
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
				ctx:   context.TODO(),
				since: 928281811,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/worklog/deleted?since=928281811",
					nil).
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
				ctx:   context.TODO(),
				since: 928281811,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/worklog/deleted?since=928281811",
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

			newService, err := NewWorklogRichTextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Deleted(testCase.args.ctx, testCase.args.since)

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

func Test_internalWorklogRichTextImpl_Updated(t *testing.T) {

	type fields struct {
		c       service.Client
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
				ctx:    context.TODO(),
				since:  928281811,
				expand: []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/worklog/updated?expand=properties&since=928281811",
					nil).
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
				ctx:    context.TODO(),
				since:  928281811,
				expand: []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/worklog/updated?expand=properties&since=928281811",
					nil).
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
				ctx:    context.TODO(),
				since:  928281811,
				expand: []string{"properties"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/worklog/updated?expand=properties&since=928281811",
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

			newService, err := NewWorklogRichTextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Updated(testCase.args.ctx, testCase.args.since, testCase.args.expand)

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

func Test_internalWorklogRichTextImpl_Add(t *testing.T) {

	payloadMocked := &model.WorklogPayloadSchemeV2{
		Comment: &model.CommentPayloadSchemeV2{
			Body: "Hello!!",
		},
		Visibility: &model.IssueWorklogVisibilityScheme{
			Type:  "group",
			Value: "jira-project-admins",
		},
		Started:          "2021-01-17T12:34:00.000+0000",
		TimeSpent:        "3h",
		TimeSpentSeconds: 12000,
	}

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
		payload      *model.WorklogPayloadSchemeV2
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
				ctx:          context.TODO(),
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

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-5/worklog?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogScheme{}).
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
				ctx:          context.TODO(),
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

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-5/worklog?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the payload is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrID: "DUMMY-5",
				payload:      nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.WorklogPayloadSchemeV2)(nil)).
					Return(bytes.NewReader([]byte{}), model.ErrNonPayloadPointerError)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrNonPayloadPointerError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
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

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-5/worklog?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
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

			newService, err := NewWorklogRichTextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Add(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.payload,
				testCase.args.options)

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

func Test_internalWorklogRichTextImpl_Update(t *testing.T) {

	payloadMocked := &model.WorklogPayloadSchemeV2{
		Comment: nil,
		Visibility: &model.IssueWorklogVisibilityScheme{
			Type:  "group",
			Value: "jira-project-admins",
		},
		Started:          "2021-01-17T12:34:00.000+0000",
		TimeSpent:        "3h",
		TimeSpentSeconds: 12000,
	}

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                     context.Context
		issueKeyOrID, worklogId string
		payload                 *model.WorklogPayloadSchemeV2
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
				ctx:          context.TODO(),
				issueKeyOrID: "DUMMY-5",
				worklogId:    "3933828822",
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

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issue/DUMMY-5/worklog/3933828822?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogScheme{}).
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
				ctx:          context.TODO(),
				issueKeyOrID: "DUMMY-5",
				worklogId:    "3933828822",
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

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/DUMMY-5/worklog/3933828822?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueWorklogScheme{}).
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
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the worklog id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrID: "DUMMY-5",
			},
			wantErr: true,
			Err:     model.ErrNoWorklogIDError,
		},

		{
			name:   "when the payload is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrID: "DUMMY-5",
				worklogId:    "38388382",
				payload:      nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.WorklogPayloadSchemeV2)(nil)).
					Return(bytes.NewReader([]byte{}), model.ErrNonPayloadPointerError)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrNonPayloadPointerError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				issueKeyOrID: "DUMMY-5",
				worklogId:    "3933828822",
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

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issue/DUMMY-5/worklog/3933828822?adjustEstimate=new&expand=properties&newEstimate=2d&notifyUsers=true&overrideEditableFlag=true&reduceBy=manual",
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

			newService, err := NewWorklogRichTextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Update(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.worklogId,
				testCase.args.payload, testCase.args.options)

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

func Test_NewWorklogRichTextService(t *testing.T) {

	type args struct {
		client  service.Client
		version string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
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
			err:     model.ErrNoVersionProvided,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := NewWorklogRichTextService(testCase.args.client, testCase.args.version)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, got, nil)
			}
		})
	}
}
