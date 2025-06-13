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

func Test_internalScreenTabFieldImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx             context.Context
		screenID, tabID int
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
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/screens/10002/tabs/18272/fields",
					"", nil).
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/screens/10002/tabs/18272/fields",
					"", nil).
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
			name:   "when the screen id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoScreenID,
		},

		{
			name:   "when the tab id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoScreenTabID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/screens/10002/tabs/18272/fields",
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

			resolutionService, err := NewScreenTabFieldService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := resolutionService.Gets(testCase.args.ctx, testCase.args.screenID, testCase.args.tabID)

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

func Test_internalScreenTabFieldImpl_Add(t *testing.T) {

	payloadMocked := map[string]interface{}{"fieldId": "customfield_10001"}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx             context.Context
		screenID, tabID int
		fieldID         string
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
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
				fieldID:  "customfield_10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/screens/10002/tabs/18272/fields",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ScreenTabFieldScheme{}).
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
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
				fieldID:  "customfield_10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/screens/10002/tabs/18272/fields",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ScreenTabFieldScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the screen id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoScreenID,
		},

		{
			name:   "when the tab id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoScreenTabID,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10001,
				tabID:    10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
				fieldID:  "customfield_10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/screens/10002/tabs/18272/fields",
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

			resolutionService, err := NewScreenTabFieldService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := resolutionService.Add(testCase.args.ctx, testCase.args.screenID, testCase.args.tabID,
				testCase.args.fieldID)

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

func Test_internalScreenTabFieldImpl_Remove(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx             context.Context
		screenID, tabID int
		fieldID         string
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
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
				fieldID:  "customfield_10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/screens/10002/tabs/18272/fields/customfield_10001",
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
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
				fieldID:  "customfield_10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/screens/10002/tabs/18272/fields/customfield_10001",
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
			name:   "when the screen id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoScreenID,
		},

		{
			name:   "when the tab id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoScreenTabID,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10001,
				tabID:    10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
				fieldID:  "customfield_10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/screens/10002/tabs/18272/fields/customfield_10001",
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

			resolutionService, err := NewScreenTabFieldService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := resolutionService.Remove(testCase.args.ctx, testCase.args.screenID, testCase.args.tabID,
				testCase.args.fieldID)

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

func Test_internalScreenTabFieldImpl_Move(t *testing.T) {

	payloadMocked := map[string]interface{}{"after": "", "position": "First"}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx             context.Context
		screenID, tabID int
		fieldID         string
		after, position string
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
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
				fieldID:  "customfield_10001",
				position: "First",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/screens/10002/tabs/18272/fields/customfield_10001/move",
					"", payloadMocked).
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
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
				fieldID:  "customfield_10001",
				position: "First",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/screens/10002/tabs/18272/fields/customfield_10001/move",
					"", payloadMocked).
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
			name:   "when the screen id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoScreenID,
		},

		{
			name:   "when the tab id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoScreenTabID,
		},

		{
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10001,
				tabID:    10001,
			},
			wantErr: true,
			Err:     model.ErrNoFieldID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				screenID: 10002,
				tabID:    18272,
				fieldID:  "customfield_10001",
				position: "First",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/screens/10002/tabs/18272/fields/customfield_10001/move",
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

			resolutionService, err := NewScreenTabFieldService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := resolutionService.Move(testCase.args.ctx, testCase.args.screenID, testCase.args.tabID,
				testCase.args.fieldID, testCase.args.after, testCase.args.position)

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

func Test_NewScreenTabFieldService(t *testing.T) {

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
			got, err := NewScreenTabFieldService(testCase.args.client, testCase.args.version)

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
