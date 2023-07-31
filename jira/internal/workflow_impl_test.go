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

func Test_internalWorkflowImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		options             *model.WorkflowSearchOptions
		startAt, maxResults int
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
				ctx: context.TODO(),
				options: &model.WorkflowSearchOptions{
					WorkflowName: []string{"workflow-name"},
					Expand:       []string{"transitions"},
					QueryString:  "workflow",
					OrderBy:      "name",
					IsActive:     true,
				},
				startAt:    50,
				maxResults: 25,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflow/search?expand=transitions&isActive=true&maxResults=25&orderBy=name&queryString=workflow&startAt=50&workflowName=workflow-name",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowPageScheme{}).
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
				ctx: context.TODO(),
				options: &model.WorkflowSearchOptions{
					WorkflowName: []string{"workflow-name"},
					Expand:       []string{"transitions"},
					QueryString:  "workflow",
					OrderBy:      "name",
					IsActive:     true,
				},
				startAt:    50,
				maxResults: 25,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/workflow/search?expand=transitions&isActive=true&maxResults=25&orderBy=name&queryString=workflow&startAt=50&workflowName=workflow-name",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowPageScheme{}).
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
				ctx: context.TODO(),
				options: &model.WorkflowSearchOptions{
					WorkflowName: []string{"workflow-name"},
					Expand:       []string{"transitions"},
					QueryString:  "workflow",
					OrderBy:      "name",
					IsActive:     true,
				},
				startAt:    50,
				maxResults: 25,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflow/search?expand=transitions&isActive=true&maxResults=25&orderBy=name&queryString=workflow&startAt=50&workflowName=workflow-name",
					"", nil).
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

			newService, err := NewWorkflowService(testCase.fields.c, testCase.fields.version, nil, nil, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Gets(testCase.args.ctx, testCase.args.options, testCase.args.startAt,
				testCase.args.maxResults)

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

func Test_internalWorkflowImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx        context.Context
		workflowId string
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
				workflowId: "2838382882",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/workflow/2838382882",
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
				ctx:        context.TODO(),
				workflowId: "2838382882",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/workflow/2838382882",
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
			name:   "when the workflow id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkflowIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.TODO(),
				workflowId: "2838382882",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/workflow/2838382882",
					"", nil).
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

			newService, err := NewWorkflowService(testCase.fields.c, testCase.fields.version, nil, nil, nil)
			assert.NoError(t, err)

			gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.workflowId)

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

func Test_internalWorkflowImpl_Create(t *testing.T) {

	payloadMocked := &model.WorkflowPayloadScheme{
		Name:        "DUMMY - Epic Workflow",
		Description: "The workflows represents the process for the Epic issue types",
		Statuses: []*model.WorkflowTransitionScreenScheme{
			{ID: "1"},
			{ID: "2"},
			{ID: "3"},
			{ID: "4"},
		},
		Transitions: []*model.WorkflowTransitionPayloadScheme{
			{
				Name: "In Progress",
				From: []string{"1"},
				To:   "2",
				Type: "global",
				Screen: &model.WorkflowTransitionScreenPayloadScheme{
					ID: "29394",
				},
			},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.WorkflowPayloadScheme
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflow",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowCreatedResponseScheme{}).
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/workflow",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowCreatedResponseScheme{}).
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflow",
					"", payloadMocked).
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

			newService, err := NewWorkflowService(testCase.fields.c, testCase.fields.version, nil, nil, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Create(testCase.args.ctx, testCase.args.payload)

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

func Test_NewWorkflowService(t *testing.T) {

	type args struct {
		client  service.Connector
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
			got, err := NewWorkflowService(testCase.args.client, testCase.args.version, nil, nil, nil)

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

func Test_internalWorkflowImpl_Bulk(t *testing.T) {

	payloadMocked := &model.WorkflowBulkOptionsScheme{
		ProjectAndIssueTypes: []*model.ProjectAndIssueTypePairScheme{
			{
				IssueTypeID: "444403",
				ProjectID:   "32424",
			},
		},
		WorkflowIds:   []string{"333"},
		WorkflowNames: []string{"workflow default"},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		options *model.WorkflowBulkOptionsScheme
		expand  []string
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
				ctx:     context.TODO(),
				options: payloadMocked,
				expand:  []string{"workflows.usages", "statuses.usages"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows?expand=workflows.usages%2Cstatuses.usages",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowReadResponseScheme{}).
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
				ctx:     context.TODO(),
				options: payloadMocked,
				expand:  []string{"workflows.usages", "statuses.usages"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/workflows?expand=workflows.usages%2Cstatuses.usages",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowReadResponseScheme{}).
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
				ctx:     context.TODO(),
				options: payloadMocked,
				expand:  []string{"workflows.usages", "statuses.usages"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows?expand=workflows.usages%2Cstatuses.usages",
					"", payloadMocked).
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

			newService, err := NewWorkflowService(testCase.fields.c, testCase.fields.version, nil, nil, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Bulk(testCase.args.ctx, testCase.args.options, testCase.args.expand)

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

func Test_internalWorkflowImpl_Creates(t *testing.T) {

	payloadMocked := &model.WorkflowCreatesPayloadScheme{
		Scope: &model.WorkflowStatusScopeScheme{
			Type: "GLOBAL",
		},
		Statuses: []*model.WorkflowStatusUpdateScheme{
			{
				Name:            "To Do",
				StatusCategory:  "TODO",
				StatusReference: "1",
			},
			{
				Name:            "In Progress",
				StatusCategory:  "IN_PROGRESS",
				StatusReference: "2",
			},
			{
				Name:            "Done",
				StatusCategory:  "DONE",
				StatusReference: "3",
			},
		},
		Workflows: []*model.WorkflowCreatePayloadScheme{
			{
				Name:        "Software workflow 1",
				Description: "workflow description sample",
				StartPointLayout: &model.StartPointLayoutScheme{
					X: -100.00030899047852,
					Y: -153.00020599365234,
				},
				Statuses: []*model.StatusLayoutUpdateScheme{
					{
						Layout: &model.StartPointLayoutScheme{
							X: 114.99993896484375,
							Y: -16,
						},
						StatusReference: "1",
					},

					{
						Layout: &model.StartPointLayoutScheme{
							X: 317.0000915527344,
							Y: -16,
						},
						StatusReference: "2",
					},

					{
						Layout: &model.StartPointLayoutScheme{
							X: 508.000244140625,
							Y: -16,
						},
						StatusReference: "3",
					},
				},
				Transitions: []*model.TransitionUpdateScheme{
					{
						ID:   "1",
						Name: "Create",
						To: &model.StatusReferenceAndPortScheme{
							StatusReference: "1",
						},
						Type: "INITIAL",
					},
					{
						ID:   "11",
						Name: "To Do",
						To: &model.StatusReferenceAndPortScheme{
							StatusReference: "2",
						},
						Type: "GLOBAL",
					},
					{
						ID:   "21",
						Name: "In Progress",
						To: &model.StatusReferenceAndPortScheme{
							StatusReference: "3",
						},
						Type: "GLOBAL",
					},
				},
			},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.WorkflowCreatesPayloadScheme
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/create",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowCreateResponseScheme{}).
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/workflows/create",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowCreateResponseScheme{}).
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/create",
					"", payloadMocked).
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

			newService, err := NewWorkflowService(testCase.fields.c, testCase.fields.version, nil, nil, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Creates(testCase.args.ctx, testCase.args.payload)

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

func Test_internalWorkflowImpl_Updates(t *testing.T) {

	payloadMocked := &model.WorkflowUpdatesPayloadScheme{
		Statuses: []*model.WorkflowStatusUpdateScheme{
			{
				Name:            "To Do",
				StatusCategory:  "TODO",
				StatusReference: "1",
			},
			{
				Name:            "In Progress",
				StatusCategory:  "IN_PROGRESS",
				StatusReference: "2",
			},
			{
				Name:            "Done",
				StatusCategory:  "DONE",
				StatusReference: "3",
			},
		},
		Workflows: []*model.WorkflowUpdatePayloadScheme{
			{
				DefaultStatusMappings: []*model.StatusMigrationScheme{
					{
						NewStatusReference: "10011",
						OldStatusReference: "10010",
					},
				},
				Description: "",
				ID:          "10001",
				StartPointLayout: &model.StartPointLayoutScheme{
					X: -100.00030899047852,
					Y: -153.00020599365234,
				},
				StatusMappings: []*model.StatusMappingScheme{
					{
						IssueTypeID: "10002",
						ProjectID:   "10003",
						StatusMigrations: []*model.StatusMigrationScheme{
							{
								NewStatusReference: "10011",
								OldStatusReference: "10010",
							},
						},
					},
				},
				Statuses: []*model.StatusLayoutUpdateScheme{
					{
						Layout: &model.StartPointLayoutScheme{
							X: 114.99993896484375,
							Y: -16,
						},
						StatusReference: "f0b24de5-25e7-4fab-ab94-63d81db6c0c0",
					},

					{
						Layout: &model.StartPointLayoutScheme{
							X: 317.0000915527344,
							Y: -16,
						},
						StatusReference: "c7a35bf0-c127-4aa6-869f-4033730c61d8",
					},

					{
						Layout: &model.StartPointLayoutScheme{
							X: 508.000244140625,
							Y: -16,
						},
						StatusReference: "6b3fc04d-3316-46c5-a257-65751aeb8849",
					},
				},
				Transitions: []*model.TransitionUpdateScheme{
					{
						ID:   "1",
						Name: "Create",
						To: &model.StatusReferenceAndPortScheme{
							StatusReference: "1",
						},
						Type: "INITIAL",
					},
					{
						ID:   "11",
						Name: "To Do",
						To: &model.StatusReferenceAndPortScheme{
							StatusReference: "2",
						},
						Type: "GLOBAL",
					},
					{
						ID:   "21",
						Name: "In Progress",
						To: &model.StatusReferenceAndPortScheme{
							StatusReference: "3",
						},
						Type: "GLOBAL",
					},
				},
			},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.WorkflowUpdatesPayloadScheme
		expand  []string
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
				ctx:     context.TODO(),
				payload: payloadMocked,
				expand:  []string{"workflows.usages", "statuses.usages"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/update?expand=workflows.usages%2Cstatuses.usages",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowUpdateResponseScheme{}).
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/workflows/update",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowUpdateResponseScheme{}).
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/update",
					"", payloadMocked).
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

			newService, err := NewWorkflowService(testCase.fields.c, testCase.fields.version, nil, nil, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Updates(testCase.args.ctx, testCase.args.payload, testCase.args.expand)

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

func Test_internalWorkflowImpl_Capabilities(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                    context.Context
		workflowID             string
		projectID, issueTypeID string
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
				workflowID:  "uuid-sample",
				projectID:   "473772",
				issueTypeID: "23",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflows/capabilities?issueTypeId=23&projectId=473772&workflowId=uuid-sample",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowCapabilitiesScheme{}).
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
				ctx:         context.TODO(),
				workflowID:  "uuid-sample",
				projectID:   "473772",
				issueTypeID: "23",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/workflows/capabilities?issueTypeId=23&projectId=473772&workflowId=uuid-sample",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowCapabilitiesScheme{}).
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
				workflowID:  "uuid-sample",
				projectID:   "473772",
				issueTypeID: "23",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflows/capabilities?issueTypeId=23&projectId=473772&workflowId=uuid-sample",
					"", nil).
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

			newService, err := NewWorkflowService(testCase.fields.c, testCase.fields.version, nil, nil, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Capabilities(testCase.args.ctx,
				testCase.args.workflowID,
				testCase.args.projectID,
				testCase.args.issueTypeID,
			)

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
