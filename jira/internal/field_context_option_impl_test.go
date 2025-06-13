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

func Test_internalIssueFieldContextOptionServiceImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldID             string
		contextID           int
		options             *model.FieldOptionContextParams
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
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
				options: &model.FieldOptionContextParams{
					OptionID:    3022,
					OnlyOptions: false,
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context/10001/option?maxResults=50&onlyOptions=false&optionId=3022&startAt=50",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldContextOptionPageScheme{}).
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
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
				options: &model.FieldOptionContextParams{
					OptionID:    3022,
					OnlyOptions: false,
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/custom_field_10002/context/10001/option?maxResults=50&onlyOptions=false&optionId=3022&startAt=50",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomFieldContextOptionPageScheme{}).
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
				ctx:       context.Background(),
				fieldID:   "",
				contextID: 10001,
				options: &model.FieldOptionContextParams{
					OptionID:    3022,
					OnlyOptions: false,
				},
				startAt:    50,
				maxResults: 50,
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
				options: &model.FieldOptionContextParams{
					OptionID:    3022,
					OnlyOptions: false,
				},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/custom_field_10002/context/10001/option?maxResults=50&onlyOptions=false&optionId=3022&startAt=50",
					"",
					nil).
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

			fieldConfigService, err := NewIssueFieldContextOptionService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Gets(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID,
				testCase.args.options, testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalIssueFieldContextOptionServiceImpl_Create(t *testing.T) {

	payloadMocked := &model.FieldContextOptionListScheme{
		Options: []*model.CustomFieldContextOptionScheme{

			// Single/Multiple Choice example
			{
				Value:    "Option 2",
				Disabled: false,
			},
			{
				Value:    "Option 4",
				Disabled: false,
			},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx       context.Context
		fieldID   string
		contextID int
		payload   *model.FieldContextOptionListScheme
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
				payload:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field/custom_field_10002/context/10001/option",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldContextOptionListScheme{}).
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
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
				payload:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field/custom_field_10002/context/10001/option",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldContextOptionListScheme{}).
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
				ctx:       context.Background(),
				fieldID:   "",
				contextID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the context id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "customfield_1000",
				contextID: 0,
			},
			wantErr: true,
			Err:     model.ErrNoFieldContextID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
				payload:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field/custom_field_10002/context/10001/option",
					"",
					payloadMocked).
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

			fieldConfigService, err := NewIssueFieldContextOptionService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Create(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID,
				testCase.args.payload)

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

func Test_internalIssueFieldContextOptionServiceImpl_Update(t *testing.T) {

	payloadMocked := &model.FieldContextOptionListScheme{
		Options: []*model.CustomFieldContextOptionScheme{

			// Single/Multiple Choice example
			{
				Value:    "Option 2",
				Disabled: false,
			},
			{
				Value:    "Option 4",
				Disabled: false,
			},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx       context.Context
		fieldID   string
		contextID int
		payload   *model.FieldContextOptionListScheme
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
				payload:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/option",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldContextOptionListScheme{}).
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
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
				payload:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/option",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldContextOptionListScheme{}).
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
				ctx:       context.Background(),
				fieldID:   "",
				contextID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the context id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "customfield_1000",
				contextID: 0,
			},
			wantErr: true,
			Err:     model.ErrNoFieldContextID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
				payload:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/option",
					"",
					payloadMocked).
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

			fieldConfigService, err := NewIssueFieldContextOptionService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Update(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID,
				testCase.args.payload)

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

func Test_internalIssueFieldContextOptionServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx       context.Context
		fieldID   string
		contextID int
		optionID  int
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
				optionID:  1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/field/custom_field_10002/context/10001/option/1001",
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
				optionID:  1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/field/custom_field_10002/context/10001/option/1001",
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
				ctx:       context.Background(),
				fieldID:   "",
				contextID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the context id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "customfield_1000",
				contextID: 0,
			},
			wantErr: true,
			Err:     model.ErrNoFieldContextID,
		},

		{
			name:   "when the option id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "customfield_1000",
				contextID: 1000,
				optionID:  0,
			},

			wantErr: true,
			Err:     model.ErrNoContextOptionID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
				optionID:  1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/field/custom_field_10002/context/10001/option/1001",
					"",
					nil).
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

			fieldConfigService, err := NewIssueFieldContextOptionService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Delete(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID,
				testCase.args.optionID)

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

func Test_internalIssueFieldContextOptionServiceImpl_Order(t *testing.T) {

	payloadMocked := &model.OrderFieldOptionPayloadScheme{
		Position:             "Last",
		CustomFieldOptionIDs: []string{"111"},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx       context.Context
		fieldID   string
		contextID int
		payload   *model.OrderFieldOptionPayloadScheme
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
				payload:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/option/move",
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
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
				payload:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/field/custom_field_10002/context/10001/option/move",
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
				ctx:       context.Background(),
				fieldID:   "",
				contextID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the context id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "customfield_1000",
				contextID: 0,
			},
			wantErr: true,
			Err:     model.ErrNoFieldContextID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				fieldID:   "custom_field_10002",
				contextID: 10001,
				payload:   payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/field/custom_field_10002/context/10001/option/move",
					"",
					payloadMocked).
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

			fieldConfigService, err := NewIssueFieldContextOptionService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Order(testCase.args.ctx, testCase.args.fieldID, testCase.args.contextID,
				testCase.args.payload)

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

func Test_NewIssueFieldContextOptionService(t *testing.T) {

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
			got, err := NewIssueFieldContextOptionService(testCase.args.client, testCase.args.version)

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
