package internal

import (
	"context"
	"encoding/json"
	"errors"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_internalQueueServiceImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx           context.Context
		serviceDeskID int
		includeCount  bool
		start, limit  int
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
				serviceDeskID: 10001,
				includeCount:  true,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/queue?includeCount=true&limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskQueuePageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				includeCount:  true,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/queue?includeCount=true&limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskQueuePageScheme{}).
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
				serviceDeskID: 10001,
				includeCount:  true,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/queue?includeCount=true&limit=50&start=100",
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

			queueService := NewQueueService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := queueService.Gets(testCase.args.ctx, testCase.args.serviceDeskID, testCase.args.includeCount,
				testCase.args.start, testCase.args.limit)

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

func Test_internalQueueServiceImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx           context.Context
		serviceDeskID int
		queueID       int
		includeCount  bool
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
				serviceDeskID: 10001,
				queueID:       29,
				includeCount:  true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/queue/29?includeCount=true",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskQueueScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				queueID:       29,
				includeCount:  true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/queue/29?includeCount=true",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskQueueScheme{}).
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
				serviceDeskID: 10001,
				queueID:       29,
				includeCount:  true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/queue/29?includeCount=true",
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

		{
			name: "when the queue id is not provided",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
			},
			Err:     model.ErrNoQueueID,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			queueService := NewQueueService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := queueService.Get(testCase.args.ctx, testCase.args.serviceDeskID,
				testCase.args.queueID, testCase.args.includeCount)

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

func Test_internalQueueServiceImpl_Issues(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx           context.Context
		serviceDeskID int
		queueID       int
		includeCount  bool
		start, limit  int
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
				serviceDeskID: 10001,
				queueID:       29,
				includeCount:  true,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/queue/29/issue?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskIssueQueueScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				queueID:       29,
				includeCount:  true,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/queue/29/issue?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceDeskIssueQueueScheme{}).
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
				serviceDeskID: 10001,
				queueID:       29,
				includeCount:  true,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/queue/29/issue?limit=50&start=100",
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

			queueService := NewQueueService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := queueService.Issues(testCase.args.ctx, testCase.args.serviceDeskID, testCase.args.queueID,
				testCase.args.start, testCase.args.limit)

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
