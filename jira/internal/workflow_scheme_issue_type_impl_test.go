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

func Test_internalWorkflowSchemeIssueTypeImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx         context.Context
		schemeID    int
		issueTypeID string
		returnDraft bool
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
				ctx:         context.TODO(),
				schemeID:    10002,
				issueTypeID: "4",
				returnDraft: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflowscheme/10002/issuetype/4?returnDraftIfExists=true",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeWorkflowMappingScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the workflow scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkflowSchemeIDError,
		},

		{
			name:   "when the issue type id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.TODO(),
				schemeID: 10002,
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypeIDError,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:         context.TODO(),
				schemeID:    10002,
				issueTypeID: "4",
				returnDraft: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/workflowscheme/10002/issuetype/4?returnDraftIfExists=true",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeWorkflowMappingScheme{}).
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
				ctx:         context.TODO(),
				schemeID:    10002,
				issueTypeID: "4",
				returnDraft: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflowscheme/10002/issuetype/4?returnDraftIfExists=true",
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

			newService := NewWorkflowSchemeIssueTypeService(testCase.fields.c, testCase.fields.version)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.schemeID, testCase.args.issueTypeID,
				testCase.args.returnDraft)

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

func Test_internalWorkflowSchemeIssueTypeImpl_Mapping(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx          context.Context
		schemeID     int
		workflowName string
		returnDraft  bool
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
				schemeID:     10002,
				workflowName: "jira workflow ",
				returnDraft:  true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflowscheme/10002/workflow?returnDraftIfExists=true&workflowName=jira+workflow+",
					nil).
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
			name:   "when the workflow scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkflowSchemeIDError,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.TODO(),
				schemeID:     10002,
				workflowName: "jira workflow ",
				returnDraft:  true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/workflowscheme/10002/workflow?returnDraftIfExists=true&workflowName=jira+workflow+",
					nil).
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
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				schemeID:     10002,
				workflowName: "jira workflow ",
				returnDraft:  true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflowscheme/10002/workflow?returnDraftIfExists=true&workflowName=jira+workflow+",
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

			newService := NewWorkflowSchemeIssueTypeService(testCase.fields.c, testCase.fields.version)

			gotResult, gotResponse, err := newService.Mapping(testCase.args.ctx, testCase.args.schemeID, testCase.args.workflowName,
				testCase.args.returnDraft)

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

func Test_internalWorkflowSchemeIssueTypeImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx         context.Context
		schemeID    int
		issueTypeID string
		updateDraft bool
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
				ctx:         context.TODO(),
				schemeID:    10002,
				issueTypeID: "4",
				updateDraft: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/workflowscheme/10002/issuetype/4?updateDraftIfNeeded=true",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowSchemeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the workflow scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkflowSchemeIDError,
		},

		{
			name:   "when the issue type id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.TODO(),
				schemeID: 10002,
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypeIDError,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:         context.TODO(),
				schemeID:    10002,
				issueTypeID: "4",
				updateDraft: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/workflowscheme/10002/issuetype/4?updateDraftIfNeeded=true",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowSchemeScheme{}).
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
				ctx:         context.TODO(),
				schemeID:    10002,
				issueTypeID: "4",
				updateDraft: true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/workflowscheme/10002/issuetype/4?updateDraftIfNeeded=true",
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

			newService := NewWorkflowSchemeIssueTypeService(testCase.fields.c, testCase.fields.version)

			gotResult, gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.schemeID, testCase.args.issueTypeID,
				testCase.args.updateDraft)

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

func Test_internalWorkflowSchemeIssueTypeImpl_Set(t *testing.T) {

	payloadMocked := &model.IssueTypeWorkflowPayloadScheme{
		IssueType:           "193",
		UpdateDraftIfNeeded: true,
		Workflow:            "jira workflow sample",
	}

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx         context.Context
		schemeID    int
		issueTypeID string
		payload     *model.IssueTypeWorkflowPayloadScheme
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
				ctx:         context.TODO(),
				schemeID:    10002,
				issueTypeID: "4",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/workflowscheme/10002/issuetype/4",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowSchemeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the workflow scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkflowSchemeIDError,
		},

		{
			name:   "when the issue type id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.TODO(),
				schemeID: 10002,
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypeIDError,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:         context.TODO(),
				schemeID:    10002,
				issueTypeID: "4",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/workflowscheme/10002/issuetype/4",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowSchemeScheme{}).
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
				ctx:         context.TODO(),
				schemeID:    10002,
				issueTypeID: "4",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/workflowscheme/10002/issuetype/4",
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

			newService := NewWorkflowSchemeIssueTypeService(testCase.fields.c, testCase.fields.version)

			gotResult, gotResponse, err := newService.Set(testCase.args.ctx, testCase.args.schemeID, testCase.args.issueTypeID,
				testCase.args.payload)

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
