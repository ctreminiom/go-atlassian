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

func Test_internalWorkflowValidatorImpl_Creation(t *testing.T) {

	payloadMocked := &model.WorkflowCreateValidatorPayloadScheme{
		Payload: &model.WorkflowCreatesPayloadScheme{
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
		},
		ValidationOptions: &model.ValidationOptionsForCreateScheme{
			Levels: []string{"ERROR", "WARNING"},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.WorkflowCreateValidatorPayloadScheme
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
					"rest/api/3/workflows/create/validation",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowValidationErrorListScheme{}).
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
					"rest/api/2/workflows/create/validation",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowValidationErrorListScheme{}).
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
					"rest/api/3/workflows/create/validation",
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

			newService, err := NewWorkflowValidatorService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Creation(testCase.args.ctx, testCase.args.payload)

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

func Test_internalWorkflowValidatorImpl_Modification(t *testing.T) {

	payloadMocked := &model.WorkflowUpdateValidatorPayloadScheme{
		Payload: &model.WorkflowUpdatesPayloadScheme{
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
		},
		ValidationOptions: &model.ValidationOptionsForCreateScheme{
			Levels: []string{"ERROR", "WARNING"},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.WorkflowUpdateValidatorPayloadScheme
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
					"rest/api/3/workflows/update/validation",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowValidationErrorListScheme{}).
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
					"rest/api/2/workflows/update/validation",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkflowValidationErrorListScheme{}).
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
					"rest/api/3/workflows/update/validation",
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

			newService, err := NewWorkflowValidatorService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Modification(testCase.args.ctx, testCase.args.payload)

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

func Test_NewWorkflowValidatorService(t *testing.T) {

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
			got, err := NewWorkflowValidatorService(testCase.args.client, testCase.args.version)

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
