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

func Test_internalIssueFieldConfigItemServiceImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                     context.Context
		id, startAt, maxResults int
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
				id:         10001,
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/fieldconfiguration/10001/fields?maxResults=50&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldConfigurationItemPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field config is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.TODO(),
				startAt:    50,
				maxResults: 50,
			},
			wantErr: true,
			Err:     model.ErrNoFieldConfigurationIDError,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.TODO(),
				id:         10001,
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/fieldconfiguration/10001/fields?maxResults=50&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldConfigurationItemPageScheme{}).
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
				ctx:        context.TODO(),
				id:         10001,
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/fieldconfiguration/10001/fields?maxResults=50&startAt=50",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http created"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http created"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldConfigurationItemService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Gets(testCase.args.ctx, testCase.args.id,
				testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalIssueFieldConfigItemServiceImpl_Update(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx     context.Context
		id      int
		payload *model.UpdateFieldConfigurationItemPayloadScheme
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
				id:  10001,
				payload: &model.UpdateFieldConfigurationItemPayloadScheme{
					FieldConfigurationItems: []*model.FieldConfigurationItemScheme{
						{
							ID:          "customfield_10012",
							IsHidden:    false,
							Description: "The new description of this item.",
						},
						{
							ID:         "customfield_10011",
							IsRequired: true,
						},
						{
							ID:          "customfield_10010",
							IsHidden:    false,
							IsRequired:  false,
							Description: "Another new description.",
							Renderer:    "wiki-renderer",
						},
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.UpdateFieldConfigurationItemPayloadScheme{
						FieldConfigurationItems: []*model.FieldConfigurationItemScheme{
							{
								ID:          "customfield_10012",
								IsHidden:    false,
								Description: "The new description of this item.",
							},
							{
								ID:         "customfield_10011",
								IsRequired: true,
							},
							{
								ID:          "customfield_10010",
								IsHidden:    false,
								IsRequired:  false,
								Description: "Another new description.",
								Renderer:    "wiki-renderer",
							},
						},
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/fieldconfiguration/10001/fields",
					bytes.NewReader([]byte{})).
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
			name:   "when the field config is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoFieldConfigurationIDError,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx: context.TODO(),
				id:  10001,
				payload: &model.UpdateFieldConfigurationItemPayloadScheme{
					FieldConfigurationItems: []*model.FieldConfigurationItemScheme{
						{
							ID:          "customfield_10012",
							IsHidden:    false,
							Description: "The new description of this item.",
						},
						{
							ID:         "customfield_10011",
							IsRequired: true,
						},
						{
							ID:          "customfield_10010",
							IsHidden:    false,
							IsRequired:  false,
							Description: "Another new description.",
							Renderer:    "wiki-renderer",
						},
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.UpdateFieldConfigurationItemPayloadScheme{
						FieldConfigurationItems: []*model.FieldConfigurationItemScheme{
							{
								ID:          "customfield_10012",
								IsHidden:    false,
								Description: "The new description of this item.",
							},
							{
								ID:         "customfield_10011",
								IsRequired: true,
							},
							{
								ID:          "customfield_10010",
								IsHidden:    false,
								IsRequired:  false,
								Description: "Another new description.",
								Renderer:    "wiki-renderer",
							},
						},
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/fieldconfiguration/10001/fields",
					bytes.NewReader([]byte{})).
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
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
				id:  10001,
				payload: &model.UpdateFieldConfigurationItemPayloadScheme{
					FieldConfigurationItems: []*model.FieldConfigurationItemScheme{
						{
							ID:          "customfield_10012",
							IsHidden:    false,
							Description: "The new description of this item.",
						},
						{
							ID:         "customfield_10011",
							IsRequired: true,
						},
						{
							ID:          "customfield_10010",
							IsHidden:    false,
							IsRequired:  false,
							Description: "Another new description.",
							Renderer:    "wiki-renderer",
						},
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.UpdateFieldConfigurationItemPayloadScheme{
						FieldConfigurationItems: []*model.FieldConfigurationItemScheme{
							{
								ID:          "customfield_10012",
								IsHidden:    false,
								Description: "The new description of this item.",
							},
							{
								ID:         "customfield_10011",
								IsRequired: true,
							},
							{
								ID:          "customfield_10010",
								IsHidden:    false,
								IsRequired:  false,
								Description: "Another new description.",
								Renderer:    "wiki-renderer",
							},
						},
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/fieldconfiguration/10001/fields",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http created"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http created"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldConfigurationItemService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Update(testCase.args.ctx, testCase.args.id, testCase.args.payload)

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
