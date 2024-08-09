package internal

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
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
				ctx: context.Background(),
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
				ctx: context.Background(),
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
				ctx: context.Background(),
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

			newService, err := NewWorkflowService(testCase.fields.c, testCase.fields.version, nil, nil)
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
		workflowID string
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
				workflowID: "2838382882",
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
				ctx:        context.Background(),
				workflowID: "2838382882",
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkflowIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.Background(),
				workflowID: "2838382882",
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

			newService, err := NewWorkflowService(testCase.fields.c, testCase.fields.version, nil, nil)
			assert.NoError(t, err)

			gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.workflowID)

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
				ctx:     context.Background(),
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
				ctx:     context.Background(),
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
				ctx:     context.Background(),
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

			newService, err := NewWorkflowService(testCase.fields.c, testCase.fields.version, nil, nil)
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
			got, err := NewWorkflowService(testCase.args.client, testCase.args.version, nil, nil)

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

func Test_internalWorkflowImpl_Search(t *testing.T) {
	type fields struct {
		c       service.Connector
		version string
	}
	type args struct {
		ctx             context.Context
		options         *model.WorkflowSearchCriteria
		expand          []string
		transitionLinks bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		want    *model.WorkflowReadResponseScheme
		want1   *model.ResponseScheme
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "when search is successful",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:             context.Background(),
				options:         &model.WorkflowSearchCriteria{},
				expand:          []string{"transitions"},
				transitionLinks: true,
			},

			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows?expand=transitions&useTransitionLinksFormat=true",
					"", &model.WorkflowSearchCriteria{}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowReadResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			want:    &model.WorkflowReadResponseScheme{},
			want1:   &model.ResponseScheme{},
			wantErr: assert.NoError,
		},
		{
			name: "when search fails due to request creation error",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:             context.Background(),
				options:         &model.WorkflowSearchCriteria{},
				expand:          []string{"transitions"},
				transitionLinks: true,
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows?expand=transitions&useTransitionLinksFormat=true",
					"", &model.WorkflowSearchCriteria{}).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
		{
			name: "when search fails due to API call error",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:             context.Background(),
				options:         &model.WorkflowSearchCriteria{},
				expand:          []string{"transitions"},
				transitionLinks: true,
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows?expand=transitions&useTransitionLinksFormat=true",
					"", &model.WorkflowSearchCriteria{}).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowReadResponseScheme{}).
					Return(nil, errors.New("error, API call failed"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
		{
			name: "when search returns no results",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:             context.Background(),
				options:         &model.WorkflowSearchCriteria{},
				expand:          []string{"transitions"},
				transitionLinks: true,
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows?expand=transitions&useTransitionLinksFormat=true",
					"", &model.WorkflowSearchCriteria{}).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowReadResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)
			},
			want:    &model.WorkflowReadResponseScheme{},
			want1:   &model.ResponseScheme{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.on != nil {
				tt.on(&tt.fields)
			}

			newService, err := NewWorkflowService(tt.fields.c, tt.fields.version, nil, nil)
			assert.NoError(t, err)

			got, got1, err := newService.Search(tt.args.ctx, tt.args.options, tt.args.expand, tt.args.transitionLinks)
			if !tt.wantErr(t, err, fmt.Sprintf("Search(%v, %v, %v, %v)", tt.args.ctx, tt.args.options, tt.args.expand, tt.args.transitionLinks)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Search(%v, %v, %v, %v)", tt.args.ctx, tt.args.options, tt.args.expand, tt.args.transitionLinks)
			assert.Equalf(t, tt.want1, got1, "Search(%v, %v, %v, %v)", tt.args.ctx, tt.args.options, tt.args.expand, tt.args.transitionLinks)
		})
	}
}

func Test_internalWorkflowImpl_Capabilities(t *testing.T) {
	type fields struct {
		c       service.Connector
		version string
	}
	type args struct {
		ctx         context.Context
		workflowID  string
		projectID   string
		issueTypeID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		want    *model.WorkflowCapabilitiesScheme
		want1   *model.ResponseScheme
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "when capabilities are successfully retrieved",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:         context.Background(),
				workflowID:  "123",
				projectID:   "456",
				issueTypeID: "789",
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflows/capabilities?issueTypeId=789&projectId=456&workflowId=123",
					"", nil).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowCapabilitiesScheme{}).
					Return(&model.ResponseScheme{}, nil)
			},
			want:    &model.WorkflowCapabilitiesScheme{},
			want1:   &model.ResponseScheme{},
			wantErr: assert.NoError,
		},
		{
			name: "when request creation fails",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:         context.Background(),
				workflowID:  "123",
				projectID:   "456",
				issueTypeID: "789",
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflows/capabilities?issueTypeId=789&projectId=456&workflowId=123",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
		{
			name: "when API call fails",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:         context.Background(),
				workflowID:  "123",
				projectID:   "456",
				issueTypeID: "789",
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflows/capabilities?issueTypeId=789&projectId=456&workflowId=123",
					"", nil).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowCapabilitiesScheme{}).
					Return(nil, errors.New("error, API call failed"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
		{
			name: "when no parameters are provided",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/workflows/capabilities",
					"", nil).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowCapabilitiesScheme{}).
					Return(&model.ResponseScheme{}, nil)
			},
			want:    &model.WorkflowCapabilitiesScheme{},
			want1:   &model.ResponseScheme{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.on != nil {
				tt.on(&tt.fields)
			}

			newService, err := NewWorkflowService(tt.fields.c, tt.fields.version, nil, nil)
			assert.NoError(t, err)

			got, got1, err := newService.Capabilities(tt.args.ctx, tt.args.workflowID, tt.args.projectID, tt.args.issueTypeID)
			if !tt.wantErr(t, err, fmt.Sprintf("Capabilities(%v, %v, %v, %v)", tt.args.ctx, tt.args.workflowID, tt.args.projectID, tt.args.issueTypeID)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Capabilities(%v, %v, %v, %v)", tt.args.ctx, tt.args.workflowID, tt.args.projectID, tt.args.issueTypeID)
			assert.Equalf(t, tt.want1, got1, "Capabilities(%v, %v, %v, %v)", tt.args.ctx, tt.args.workflowID, tt.args.projectID, tt.args.issueTypeID)
		})
	}
}

func Test_internalWorkflowImpl_Creates(t *testing.T) {

	payload := &model.WorkflowCreatesPayload{
		Scope: &model.WorkflowScopeScheme{Type: "GLOBAL"},
	}

	// Add the status references on the payload
	statuses := []struct {
		ID, Name, StatusCategory, StatusReference string
	}{
		{"10012", "To Do", "TODO", "f0b24de5-25e7-4fab-ab94-63d81db6c0c0"},
		{"3", "In Progress", "IN_PROGRESS", "c7a35bf0-c127-4aa6-869f-4033730c61d8"},
		{"10002", "Done", "DONE", "6b3fc04d-3316-46c5-a257-65751aeb8849"},
	}

	for _, status := range statuses {
		payload.AddStatus(&model.WorkflowStatusUpdateScheme{
			ID:              status.ID,
			Name:            status.Name,
			StatusCategory:  status.StatusCategory,
			StatusReference: status.StatusReference,
		})
	}

	epicWorkflow := &model.WorkflowCreateScheme{
		Description: "This workflow represents the process of software development related to epics.",
		Name:        "Epic Software Development Workflow V4",
	}

	// Add the statuses to the workflow using the referenceID and the layout
	layouts := []struct {
		X, Y            float64
		StatusReference string
	}{
		{114.99993896484375, -16, "f0b24de5-25e7-4fab-ab94-63d81db6c0c0"},
		{317.0000915527344, -16, "c7a35bf0-c127-4aa6-869f-4033730c61d8"},
		{508.000244140625, -16, "6b3fc04d-3316-46c5-a257-65751aeb8849"},
	}

	for _, layout := range layouts {
		epicWorkflow.AddStatus(&model.StatusLayoutUpdateScheme{
			Layout:          &model.WorkflowLayoutScheme{X: layout.X, Y: layout.Y},
			StatusReference: layout.StatusReference,
		})
	}

	// Add the transitions to the workflow
	transitions := []struct {
		ID, Type, Name, StatusReference string
	}{
		{"1", "INITIAL", "Create", "f0b24de5-25e7-4fab-ab94-63d81db6c0c0"},
		{"21", "GLOBAL", "In Progress", "c7a35bf0-c127-4aa6-869f-4033730c61d8"},
		{"31", "GLOBAL", "Done", "6b3fc04d-3316-46c5-a257-65751aeb8849"},
	}

	for _, transition := range transitions {
		err := epicWorkflow.AddTransition(&model.TransitionUpdateDTOScheme{
			ID:   transition.ID,
			Type: transition.Type,
			Name: transition.Name,
			To: &model.StatusReferenceAndPortScheme{
				StatusReference: transition.StatusReference,
			},
		})

		if err != nil {
			log.Fatal(err)
		}
	}

	// You can multiple workflows on the same payload
	if err := payload.AddWorkflow(epicWorkflow); err != nil {
		log.Fatal(err)
	}

	type fields struct {
		c       service.Connector
		version string
	}
	type args struct {
		ctx     context.Context
		payload *model.WorkflowCreatesPayload
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		want    *model.WorkflowCreateResponseScheme
		want1   *model.ResponseScheme
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "when workflow is successfully created",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: payload,
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/create",
					"", payload).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowCreateResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)
			},
			want:    &model.WorkflowCreateResponseScheme{},
			want1:   &model.ResponseScheme{},
			wantErr: assert.NoError,
		},
		{
			name: "when request creation fails",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: payload,
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/create",
					"", payload).
					Return(nil, errors.New("error, unable to create the http request"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
		{
			name: "when API call fails",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: payload,
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/create",
					"", payload).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowCreateResponseScheme{}).
					Return(nil, errors.New("error, API call failed"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.on != nil {
				tt.on(&tt.fields)
			}

			newService, err := NewWorkflowService(tt.fields.c, tt.fields.version, nil, nil)
			assert.NoError(t, err)

			got, got1, err := newService.Creates(tt.args.ctx, tt.args.payload)
			if !tt.wantErr(t, err, fmt.Sprintf("Creates(%v, %v)", tt.args.ctx, tt.args.payload)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Creates(%v, %v)", tt.args.ctx, tt.args.payload)
			assert.Equalf(t, tt.want1, got1, "Creates(%v, %v)", tt.args.ctx, tt.args.payload)
		})
	}
}

func Test_internalWorkflowImpl_Updates(t *testing.T) {
	type fields struct {
		c       service.Connector
		version string
	}
	type args struct {
		ctx     context.Context
		payload *model.WorkflowUpdatesPayloadScheme
		expand  []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		want    *model.WorkflowUpdateResponseScheme
		want1   *model.ResponseScheme
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "when workflow is successfully updated",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.WorkflowUpdatesPayloadScheme{
					// populate the payload fields
				},
				expand: []string{"transitions"},
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/update?expand=transitions",
					"", &model.WorkflowUpdatesPayloadScheme{}).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowUpdateResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)
			},
			want:    &model.WorkflowUpdateResponseScheme{},
			want1:   &model.ResponseScheme{},
			wantErr: assert.NoError,
		},
		{
			name: "when request creation fails",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.WorkflowUpdatesPayloadScheme{
					// populate the payload fields
				},
				expand: []string{"transitions"},
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/update?expand=transitions",
					"", &model.WorkflowUpdatesPayloadScheme{}).
					Return(nil, errors.New("error, unable to create the http request"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
		{
			name: "when API call fails",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.WorkflowUpdatesPayloadScheme{
					// populate the payload fields
				},
				expand: []string{"transitions"},
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/update?expand=transitions",
					"", &model.WorkflowUpdatesPayloadScheme{}).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowUpdateResponseScheme{}).
					Return(nil, errors.New("error, API call failed"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
		{
			name: "when no expand parameters are provided",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.WorkflowUpdatesPayloadScheme{
					// populate the payload fields
				},
				expand: []string{},
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/update",
					"", &model.WorkflowUpdatesPayloadScheme{}).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowUpdateResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)
			},
			want:    &model.WorkflowUpdateResponseScheme{},
			want1:   &model.ResponseScheme{},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.on != nil {
				tt.on(&tt.fields)
			}

			newService, err := NewWorkflowService(tt.fields.c, tt.fields.version, nil, nil)
			assert.NoError(t, err)

			got, got1, err := newService.Updates(tt.args.ctx, tt.args.payload, tt.args.expand)
			if !tt.wantErr(t, err, fmt.Sprintf("Updates(%v, %v, %v)", tt.args.ctx, tt.args.payload, tt.args.expand)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Updates(%v, %v, %v)", tt.args.ctx, tt.args.payload, tt.args.expand)
			assert.Equalf(t, tt.want1, got1, "Updates(%v, %v, %v)", tt.args.ctx, tt.args.payload, tt.args.expand)
		})
	}
}

func Test_internalWorkflowImpl_ValidateCreateWorkflows(t *testing.T) {
	type fields struct {
		c       service.Connector
		version string
	}
	type args struct {
		ctx     context.Context
		payload *model.ValidationOptionsForCreateScheme
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		want    *model.WorkflowValidationErrorListScheme
		want1   *model.ResponseScheme
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "when validation is successful",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.ValidationOptionsForCreateScheme{
					// populate the payload fields
				},
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/create/validation",
					"", &model.ValidationOptionsForCreateScheme{}).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowValidationErrorListScheme{}).
					Return(&model.ResponseScheme{}, nil)
			},
			want:    &model.WorkflowValidationErrorListScheme{},
			want1:   &model.ResponseScheme{},
			wantErr: assert.NoError,
		},
		{
			name: "when request creation fails",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.ValidationOptionsForCreateScheme{
					// populate the payload fields
				},
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/create/validation",
					"", &model.ValidationOptionsForCreateScheme{}).
					Return(nil, errors.New("error, unable to create the http request"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
		{
			name: "when API call fails",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.ValidationOptionsForCreateScheme{
					// populate the payload fields
				},
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/create/validation",
					"", &model.ValidationOptionsForCreateScheme{}).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowValidationErrorListScheme{}).
					Return(nil, errors.New("error, API call failed"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.on != nil {
				tt.on(&tt.fields)
			}

			newService, err := NewWorkflowService(tt.fields.c, tt.fields.version, nil, nil)
			assert.NoError(t, err)

			got, got1, err := newService.ValidateCreateWorkflows(tt.args.ctx, tt.args.payload)
			if !tt.wantErr(t, err, fmt.Sprintf("ValidateCreateWorkflows(%v, %v)", tt.args.ctx, tt.args.payload)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ValidateCreateWorkflows(%v, %v)", tt.args.ctx, tt.args.payload)
			assert.Equalf(t, tt.want1, got1, "ValidateCreateWorkflows(%v, %v)", tt.args.ctx, tt.args.payload)
		})
	}
}

func Test_internalWorkflowImpl_ValidateUpdateWorkflows(t *testing.T) {
	type fields struct {
		c       service.Connector
		version string
	}
	type args struct {
		ctx     context.Context
		payload *model.ValidationOptionsForUpdateScheme
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		want    *model.WorkflowValidationErrorListScheme
		want1   *model.ResponseScheme
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "when validation is successful",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.ValidationOptionsForUpdateScheme{
					// populate the payload fields
				},
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/update/validation",
					"", &model.ValidationOptionsForUpdateScheme{}).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowValidationErrorListScheme{}).
					Return(&model.ResponseScheme{}, nil)
			},
			want:    &model.WorkflowValidationErrorListScheme{},
			want1:   &model.ResponseScheme{},
			wantErr: assert.NoError,
		},
		{
			name: "when request creation fails",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.ValidationOptionsForUpdateScheme{
					// populate the payload fields
				},
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/update/validation",
					"", &model.ValidationOptionsForUpdateScheme{}).
					Return(nil, errors.New("error, unable to create the http request"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
		{
			name: "when API call fails",
			fields: fields{
				c:       mocks.NewConnector(t),
				version: "3",
			},
			args: args{
				ctx:     context.Background(),
				payload: &model.ValidationOptionsForUpdateScheme{
					// populate the payload fields
				},
			},
			on: func(fields *fields) {
				client := fields.c.(*mocks.Connector)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/workflows/update/validation",
					"", &model.ValidationOptionsForUpdateScheme{}).
					Return(&http.Request{}, nil)
				client.On("Call",
					&http.Request{},
					&model.WorkflowValidationErrorListScheme{}).
					Return(nil, errors.New("error, API call failed"))
			},
			want:    nil,
			want1:   nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.on != nil {
				tt.on(&tt.fields)
			}

			newService, err := NewWorkflowService(tt.fields.c, tt.fields.version, nil, nil)
			assert.NoError(t, err)

			got, got1, err := newService.ValidateUpdateWorkflows(tt.args.ctx, tt.args.payload)
			if !tt.wantErr(t, err, fmt.Sprintf("ValidateUpdateWorkflows(%v, %v)", tt.args.ctx, tt.args.payload)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ValidateUpdateWorkflows(%v, %v)", tt.args.ctx, tt.args.payload)
			assert.Equalf(t, tt.want1, got1, "ValidateUpdateWorkflows(%v, %v)", tt.args.ctx, tt.args.payload)
		})
	}
}
