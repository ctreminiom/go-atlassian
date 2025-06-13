package internal

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalServiceDeskImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx          context.Context
		start, limit int
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
			name: "when the parameters are correct",
			args: args{
				ctx:   context.Background(),
				start: 100,
				limit: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:   context.Background(),
				start: 100,
				limit: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskPageScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:   context.Background(),
				start: 100,
				limit: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewServiceDeskService(testCase.fields.c, "latest", nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Gets(testCase.args.ctx, testCase.args.start, testCase.args.limit)

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

func Test_internalServiceDeskImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx           context.Context
		serviceDeskID string
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
			name: "when the parameters are correct",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: "10001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},

		{
			name: "when the service desk id is not provided",
			args: args{
				ctx: context.Background(),
			},
			Err:     model.ErrNoServiceDeskID,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewServiceDeskService(testCase.fields.c, "latest", nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Get(testCase.args.ctx, testCase.args.serviceDeskID)

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

func Test_internalServiceDeskImpl_Attach(t *testing.T) {

	absolutePathMocked, err := filepath.Abs("../../../LICENSE")
	if err != nil {
		t.Fatal(err)
	}

	fileMocked, err := os.Open(absolutePathMocked)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx           context.Context
		serviceDeskID string
		fileName      string
		file          io.Reader
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
			name: "when the parameters are correct",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: "10001",
				fileName:      "LICENSE",
				file:          fileMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/attachTemporaryFile",
					mock.AnythingOfType("string"),
					mock.Anything).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskTemporaryFileScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: "10001",
				fileName:      "LICENSE",
				file:          fileMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/attachTemporaryFile",
					mock.AnythingOfType("string"),
					mock.Anything).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskTemporaryFileScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: "10001",
				fileName:      "LICENSE",
				file:          fileMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/attachTemporaryFile",
					mock.AnythingOfType("string"),
					mock.Anything).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},

		{
			name: "when the service desk id is not provided",
			args: args{
				ctx: context.Background(),
			},
			Err:     model.ErrNoServiceDeskID,
			wantErr: true,
		},

		{
			name: "when the file name is not provided",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: "10001",
			},
			Err:     model.ErrNoFileName,
			wantErr: true,
		},

		{
			name: "when the file reader is not provided",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: "10001",
				fileName:      "LICENSE",
			},
			Err:     model.ErrNoFileReader,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewServiceDeskService(testCase.fields.c, "latest", nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Attach(testCase.args.ctx, testCase.args.serviceDeskID, testCase.args.fileName,
				testCase.args.file)

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
