package internal

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalIssueFieldContextServiceImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldID             string
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
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
				options: &model.FieldContextOptionsScheme{
					IsAnyIssueType:  true,
					IsGlobalContext: false,
					ContextID:       []int{10001, 10002},
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context?contextId=10001&contextId=10002&isAnyIssueType=true&isGlobalContext=false&maxResults=50&startAt=50",
					"",
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
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
				options: &model.FieldContextOptionsScheme{
					IsAnyIssueType:  true,
					IsGlobalContext: false,
					ContextID:       []int{10001, 10002},
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context?contextId=10001&contextId=10002&isAnyIssueType=true&isGlobalContext=false&maxResults=50&startAt=50",
					"",
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
				ctx:     context.Background(),
				fieldID: "",
				options: &model.FieldContextOptionsScheme{
					IsAnyIssueType:  true,
					IsGlobalContext: false,
					ContextID:       []int{10001, 10002},
				},
				startAt:    50,
				maxResults: 50,
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
				options: &model.FieldContextOptionsScheme{
					IsAnyIssueType:  true,
					IsGlobalContext: false,
					ContextID:       []int{10001, 10002},
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context?contextId=10001&contextId=10002&isAnyIssueType=true&isGlobalContext=false&maxResults=50&startAt=50",
					"",
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Gets(testCase.args.ctx, testCase.args.fieldID, testCase.args.options,
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

	payloadMocked := &model.FieldContextPayloadScheme{
		IssueTypeIDs: []int{10010},
		ProjectIDs:   nil,
		Name:         "Bug fields context",
		Description:  "A context used to define the custom field options for bugs.",
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		fieldID string
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
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field/custom_field_10002/context",
					"",
					payloadMocked).
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
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context",
					"",
					payloadMocked).
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
				ctx:     context.Background(),
				fieldID: "",
				payload: payloadMocked,
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context",
					"",
					payloadMocked).
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Create(testCase.args.ctx, testCase.args.fieldID, testCase.args.payload)

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
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldID             string
		contextIDs          []int
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
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextIDs: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context/defaultValue?contextId=10001&maxResults=50&startAt=0",
					"",
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
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextIDs: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/defaultValue?contextId=10001&maxResults=50&startAt=0",
					"",
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
				ctx:     context.Background(),
				fieldID: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextIDs: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/defaultValue?contextId=10001&maxResults=50&startAt=0",
					"",
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.GetDefaultValues(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextIDs,
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

	payloadMocked := &model.FieldContextDefaultPayloadScheme{
		DefaultValues: []*model.CustomFieldDefaultValueScheme{
			{
				ContextID: "10128",
				OptionID:  "10022",
				Type:      "option.single",
			},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		fieldID string
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
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/defaultValue",
					"",
					payloadMocked).
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
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/defaultValue",
					"",
					payloadMocked).
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
				ctx:     context.Background(),
				fieldID: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/defaultValue",
					"",
					payloadMocked).
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.SetDefaultValue(testCase.args.ctx, testCase.args.fieldID, testCase.args.payload)

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
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldID             string
		contextIDs          []int
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
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextIDs: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context/issuetypemapping?contextId=10001&maxResults=50&startAt=0",
					"",
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
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextIDs: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/issuetypemapping?contextId=10001&maxResults=50&startAt=0",
					"",
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
				ctx:     context.Background(),
				fieldID: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextIDs: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/issuetypemapping?contextId=10001&maxResults=50&startAt=0",
					"",
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.IssueTypesContext(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextIDs,
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
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldID             string
		contextIDs          []int
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
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextIDs: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context/projectmapping?contextId=10001&maxResults=50&startAt=0",
					"",
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
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextIDs: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/projectmapping?contextId=10001&maxResults=50&startAt=0",
					"",
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
				ctx:     context.Background(),
				fieldID: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextIDs: []int{10001},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/projectmapping?contextId=10001&maxResults=50&startAt=0",
					"",
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.ProjectsContext(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextIDs,
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

	payloadMockedWithDescription := map[string]interface{}{"description": "new customfield context description", "name": "DUMMY - customfield_10002 Context"}

	payloadMocked := map[string]interface{}{"name": "DUMMY - customfield_10002 Context"}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx               context.Context
		fieldID           string
		contextID         int
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
				ctx:         context.Background(),
				fieldID:     "custom_field_10002",
				contextID:   10001,
				name:        "DUMMY - customfield_10002 Context",
				description: "new customfield context description",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001",
					"",
					payloadMockedWithDescription).
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
			name:   "when the description is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:         context.Background(),
				fieldID:     "custom_field_10002",
				contextID:   10001,
				name:        "DUMMY - customfield_10002 Context",
				description: "",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001",
					"",
					payloadMocked).
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
				ctx:         context.Background(),
				fieldID:     "custom_field_10002",
				contextID:   10001,
				name:        "DUMMY - customfield_10002 Context",
				description: "new customfield context description",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001",
					"",
					payloadMockedWithDescription).
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
				ctx:     context.Background(),
				fieldID: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:         context.Background(),
				fieldID:     "custom_field_10002",
				contextID:   10001,
				name:        "DUMMY - customfield_10002 Context",
				description: "new customfield context description",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001",
					"",
					payloadMockedWithDescription).
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Update(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID,
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
		c       service.Connector
		version string
	}

	type args struct {
		ctx       context.Context
		fieldID   string
		contextID int
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
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/field/custom_field_10002/context/10001",
					"",
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
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/field/custom_field_10002/context/10001",
					"",
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
				ctx:     context.Background(),
				fieldID: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/field/custom_field_10002/context/10001",
					"",
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Delete(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID)

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

	payloadMocked := map[string]interface{}{"issueTypeIds": []string{"4", "3"}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx           context.Context
		fieldID       string
		contextID     int
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
				ctx:           context.Background(),
				fieldID:       "custom_field_10002",
				contextID:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/issuetype",
					"",
					payloadMocked).
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
				ctx:           context.Background(),
				fieldID:       "custom_field_10002",
				contextID:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/issuetype",
					"",
					payloadMocked).
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
				ctx:     context.Background(),
				fieldID: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:           context.Background(),
				fieldID:       "custom_field_10002",
				contextID:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/issuetype",
					"",
					payloadMocked).
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.AddIssueTypes(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID,
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

	payloadMocked := map[string]interface{}{"issueTypeIds": []string{"4", "3"}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx           context.Context
		fieldID       string
		contextID     int
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
				ctx:           context.Background(),
				fieldID:       "custom_field_10002",
				contextID:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field/custom_field_10002/context/10001/issuetype/remove",
					"",
					payloadMocked).
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
				ctx:           context.Background(),
				fieldID:       "custom_field_10002",
				contextID:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context/10001/issuetype/remove",
					"",
					payloadMocked).
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
				ctx:     context.Background(),
				fieldID: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the issuetype id's are not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:           context.Background(),
				fieldID:       "custom_field_10002",
				issueTypesIds: nil,
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypes,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:           context.Background(),
				fieldID:       "custom_field_10002",
				contextID:     10001,
				issueTypesIds: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context/10001/issuetype/remove",
					"",
					payloadMocked).
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.RemoveIssueTypes(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID,
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

	payloadMocked := map[string]interface{}{"projectIds": []string{"4", "3"}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx        context.Context
		fieldID    string
		contextID  int
		projectIDs []string
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
				fieldID:    "custom_field_10002",
				contextID:  10001,
				projectIDs: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/project",
					"",
					payloadMocked).
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
				fieldID:    "custom_field_10002",
				contextID:  10001,
				projectIDs: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/project",
					"",
					payloadMocked).
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
				ctx:     context.Background(),
				fieldID: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the context id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
			},
			wantErr: true,
			Err:     model.ErrNoFieldContextID,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDs,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextID:  10001,
				projectIDs: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/project",
					"",
					payloadMocked).
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Link(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID,
				testCase.args.projectIDs)

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

func Test_internalIssueFieldContextServiceImpl_Unlink(t *testing.T) {

	payloadMocked := map[string]interface{}{"projectIds": []string{"4", "3"}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx        context.Context
		fieldID    string
		contextID  int
		projectIDs []string
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
				fieldID:    "custom_field_10002",
				contextID:  10001,
				projectIDs: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field/custom_field_10002/context/10001/project/remove",
					"",
					payloadMocked).
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
				fieldID:    "custom_field_10002",
				contextID:  10001,
				projectIDs: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context/10001/project/remove",
					"",
					payloadMocked).
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
				ctx:     context.Background(),
				fieldID: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the context id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.Background(),
				fieldID: "custom_field_10002",
			},
			wantErr: true,
			Err:     model.ErrNoFieldContextID,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDs,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.Background(),
				fieldID:    "custom_field_10002",
				contextID:  10001,
				projectIDs: []string{"4", "3"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context/10001/project/remove",
					"",
					payloadMocked).
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

			fieldConfigService, err := NewIssueFieldContextService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.UnLink(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID,
				testCase.args.projectIDs)

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
