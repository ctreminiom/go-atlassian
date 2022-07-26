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

func Test_internalIssueFieldContextServiceImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldId             string
		options             *model.FieldContextOptionsScheme
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
				ctx:     context.TODO(),
				fieldId: "custom_field_10002",
				options: &model.FieldContextOptionsScheme{
					IsAnyIssueType:  true,
					IsGlobalContext: false,
					ContextID:       []int{10001, 10002},
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context?contextId=10001&contextId=10002&isAnyIssueType=true&isGlobalContext=false&maxResults=50&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldContextPageScheme{}).
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
				fieldId: "custom_field_10002",
				options: &model.FieldContextOptionsScheme{
					IsAnyIssueType:  true,
					IsGlobalContext: false,
					ContextID:       []int{10001, 10002},
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context?contextId=10001&contextId=10002&isAnyIssueType=true&isGlobalContext=false&maxResults=50&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldContextPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
				options: &model.FieldContextOptionsScheme{
					IsAnyIssueType:  true,
					IsGlobalContext: false,
					ContextID:       []int{10001, 10002},
				},
				startAt:    50,
				maxResults: 50,
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "custom_field_10002",
				options: &model.FieldContextOptionsScheme{
					IsAnyIssueType:  true,
					IsGlobalContext: false,
					ContextID:       []int{10001, 10002},
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context?contextId=10001&contextId=10002&isAnyIssueType=true&isGlobalContext=false&maxResults=50&startAt=50",
					nil).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Gets(testCase.args.ctx, testCase.args.fieldId, testCase.args.options,
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

func Test_internalIssueFieldContextServiceImpl_Create(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx     context.Context
		fieldId string
		payload *model.FieldContextPayloadScheme
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
				fieldId: "custom_field_10002",
				payload: &model.FieldContextPayloadScheme{
					IssueTypeIDs: []int{10010},
					ProjectIDs:   nil,
					Name:         "Bug fields context",
					Description:  "A context used to define the custom field options for bugs.",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextPayloadScheme{
						IssueTypeIDs: []int{10010},
						ProjectIDs:   nil,
						Name:         "Bug fields context",
						Description:  "A context used to define the custom field options for bugs.",
					},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field/custom_field_10002/context",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldContextScheme{}).
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
				fieldId: "custom_field_10002",
				payload: &model.FieldContextPayloadScheme{
					IssueTypeIDs: []int{10010},
					ProjectIDs:   nil,
					Name:         "Bug fields context",
					Description:  "A context used to define the custom field options for bugs.",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextPayloadScheme{
						IssueTypeIDs: []int{10010},
						ProjectIDs:   nil,
						Name:         "Bug fields context",
						Description:  "A context used to define the custom field options for bugs.",
					},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldContextScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
				payload: &model.FieldContextPayloadScheme{
					IssueTypeIDs: []int{10010},
					ProjectIDs:   nil,
					Name:         "Bug fields context",
					Description:  "A context used to define the custom field options for bugs.",
				},
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "custom_field_10002",
				payload: &model.FieldContextPayloadScheme{
					IssueTypeIDs: []int{10010},
					ProjectIDs:   nil,
					Name:         "Bug fields context",
					Description:  "A context used to define the custom field options for bugs.",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextPayloadScheme{
						IssueTypeIDs: []int{10010},
						ProjectIDs:   nil,
						Name:         "Bug fields context",
						Description:  "A context used to define the custom field options for bugs.",
					},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Create(testCase.args.ctx, testCase.args.fieldId, testCase.args.payload)

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

func Test_internalIssueFieldContextServiceImpl_GetDefaultValues(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldId             string
		contextIds          []int
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
				ctx:        context.TODO(),
				fieldId:    "custom_field_10002",
				contextIds: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context/defaultValue?contextId=10001&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldDefaultValuePageScheme{}).
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
				fieldId:    "custom_field_10002",
				contextIds: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/defaultValue?contextId=10001&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldDefaultValuePageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.TODO(),
				fieldId:    "custom_field_10002",
				contextIds: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/defaultValue?contextId=10001&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldDefaultValuePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.GetDefaultValues(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextIds,
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

func Test_internalIssueFieldContextServiceImpl_SetDefaultValue(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx     context.Context
		fieldId string
		payload *model.FieldContextDefaultPayloadScheme
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
				fieldId: "custom_field_10002",
				payload: &model.FieldContextDefaultPayloadScheme{
					DefaultValues: []*model.CustomFieldDefaultValueScheme{
						{
							ContextID: "10128",
							OptionID:  "10022",
							Type:      "option.single",
						},
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextDefaultPayloadScheme{
						DefaultValues: []*model.CustomFieldDefaultValueScheme{
							{
								ContextID: "10128",
								OptionID:  "10022",
								Type:      "option.single",
							},
						},
					},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/defaultValue",
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "custom_field_10002",
				payload: &model.FieldContextDefaultPayloadScheme{
					DefaultValues: []*model.CustomFieldDefaultValueScheme{
						{
							ContextID: "10128",
							OptionID:  "10022",
							Type:      "option.single",
						},
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextDefaultPayloadScheme{
						DefaultValues: []*model.CustomFieldDefaultValueScheme{
							{
								ContextID: "10128",
								OptionID:  "10022",
								Type:      "option.single",
							},
						},
					},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/defaultValue",
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
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "custom_field_10002",
				payload: &model.FieldContextDefaultPayloadScheme{
					DefaultValues: []*model.CustomFieldDefaultValueScheme{
						{
							ContextID: "10128",
							OptionID:  "10022",
							Type:      "option.single",
						},
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldContextDefaultPayloadScheme{
						DefaultValues: []*model.CustomFieldDefaultValueScheme{
							{
								ContextID: "10128",
								OptionID:  "10022",
								Type:      "option.single",
							},
						},
					},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/defaultValue",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.SetDefaultValue(testCase.args.ctx, testCase.args.fieldId, testCase.args.payload)

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

func Test_internalIssueFieldContextServiceImpl_IssueTypesContext(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldId             string
		contextIds          []int
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
				ctx:        context.TODO(),
				fieldId:    "custom_field_10002",
				contextIds: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context/issuetypemapping?contextId=10001&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeToContextMappingPageScheme{}).
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
				fieldId:    "custom_field_10002",
				contextIds: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/issuetypemapping?contextId=10001&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeToContextMappingPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.TODO(),
				fieldId:    "custom_field_10002",
				contextIds: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/issuetypemapping?contextId=10001&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.IssueTypesContext(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextIds,
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

func Test_internalIssueFieldContextServiceImpl_ProjectsContext(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldId             string
		contextIds          []int
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
				ctx:        context.TODO(),
				fieldId:    "custom_field_10002",
				contextIds: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context/projectmapping?contextId=10001&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldContextProjectMappingPageScheme{}).
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
				fieldId:    "custom_field_10002",
				contextIds: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/projectmapping?contextId=10001&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldContextProjectMappingPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.TODO(),
				fieldId:    "custom_field_10002",
				contextIds: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/projectmapping?contextId=10001&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.ProjectsContext(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextIds,
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

func Test_internalIssueFieldContextServiceImpl_Update(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx               context.Context
		fieldId           string
		contextId         int
		name, description string
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
				fieldId:     "custom_field_10002",
				contextId:   10001,
				name:        "DUMMY - customfield_10002 Context",
				description: "new customfield context description",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						Name        string "json:\"name,omitempty\""
						Description string "json:\"description,omitempty\""
					}{
						Name:        "DUMMY - customfield_10002 Context",
						Description: "new customfield context description"},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001",
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:         context.TODO(),
				fieldId:     "custom_field_10002",
				contextId:   10001,
				name:        "DUMMY - customfield_10002 Context",
				description: "new customfield context description",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						Name        string "json:\"name,omitempty\""
						Description string "json:\"description,omitempty\""
					}{
						Name:        "DUMMY - customfield_10002 Context",
						Description: "new customfield context description"},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001",
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
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:         context.TODO(),
				fieldId:     "custom_field_10002",
				contextId:   10001,
				name:        "DUMMY - customfield_10002 Context",
				description: "new customfield context description",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						Name        string "json:\"name,omitempty\""
						Description string "json:\"description,omitempty\""
					}{
						Name:        "DUMMY - customfield_10002 Context",
						Description: "new customfield context description"},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Update(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId,
				testCase.args.name, testCase.args.description)

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

func Test_internalIssueFieldContextServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx       context.Context
		fieldId   string
		contextId int
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
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/field/custom_field_10002/context/10001",
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
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/field/custom_field_10002/context/10001",
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
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.TODO(),
				fieldId:   "custom_field_10002",
				contextId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/field/custom_field_10002/context/10001",
					nil).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Delete(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId)

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

func Test_internalIssueFieldContextServiceImpl_AddIssueTypes(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx           context.Context
		fieldId       string
		contextId     int
		issueTypesIds []string
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
				ctx:           context.TODO(),
				fieldId:       "custom_field_10002",
				contextId:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						IssueTypeIds []string "json:\"issueTypeIds\""
					}{IssueTypeIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/issuetype",
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:           context.TODO(),
				fieldId:       "custom_field_10002",
				contextId:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						IssueTypeIds []string "json:\"issueTypeIds\""
					}{IssueTypeIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/issuetype",
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
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:           context.TODO(),
				fieldId:       "custom_field_10002",
				contextId:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						IssueTypeIds []string "json:\"issueTypeIds\""
					}{IssueTypeIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/issuetype",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.AddIssueTypes(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId,
				testCase.args.issueTypesIds)

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

func Test_internalIssueFieldContextServiceImpl_RemoveIssueTypes(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx           context.Context
		fieldId       string
		contextId     int
		issueTypesIds []string
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
				ctx:           context.TODO(),
				fieldId:       "custom_field_10002",
				contextId:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						IssueTypeIds []string "json:\"issueTypeIds\""
					}{IssueTypeIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field/custom_field_10002/context/10001/issuetype/remove",
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:           context.TODO(),
				fieldId:       "custom_field_10002",
				contextId:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						IssueTypeIds []string "json:\"issueTypeIds\""
					}{IssueTypeIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context/10001/issuetype/remove",
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
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:           context.TODO(),
				fieldId:       "custom_field_10002",
				contextId:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						IssueTypeIds []string "json:\"issueTypeIds\""
					}{IssueTypeIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context/10001/issuetype/remove",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.RemoveIssueTypes(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId,
				testCase.args.issueTypesIds)

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

func Test_internalIssueFieldContextServiceImpl_Link(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx        context.Context
		fieldId    string
		contextId  int
		projectIds []string
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
				fieldId:    "custom_field_10002",
				contextId:  10001,
				projectIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						ProjectIds []string "json:\"projectIds\""
					}{ProjectIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/project",
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.TODO(),
				fieldId:    "custom_field_10002",
				contextId:  10001,
				projectIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						ProjectIds []string "json:\"projectIds\""
					}{ProjectIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/project",
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
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.TODO(),
				fieldId:    "custom_field_10002",
				contextId:  10001,
				projectIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						ProjectIds []string "json:\"projectIds\""
					}{ProjectIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/project",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Link(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId,
				testCase.args.projectIds)

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

func Test_internalIssueFieldContextServiceImpl_UnLink(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx        context.Context
		fieldId    string
		contextId  int
		projectIds []string
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
				fieldId:    "custom_field_10002",
				contextId:  10001,
				projectIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						ProjectIds []string "json:\"projectIds\""
					}{ProjectIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field/custom_field_10002/context/10001/project/remove",
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.TODO(),
				fieldId:    "custom_field_10002",
				contextId:  10001,
				projectIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						ProjectIds []string "json:\"projectIds\""
					}{ProjectIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context/10001/project/remove",
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
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.TODO(),
				fieldId:    "custom_field_10002",
				contextId:  10001,
				projectIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						ProjectIds []string "json:\"projectIds\""
					}{ProjectIds: []string{"4", "3"}},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context/10001/project/remove",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.UnLink(testCase.args.ctx, testCase.args.fieldId, testCase.args.contextId,
				testCase.args.projectIds)

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
