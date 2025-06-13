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

func Test_SprintService_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx      context.Context
		sprintID int
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
				ctx:      context.Background(),
				sprintID: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintID: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			Err:     model.ErrNoExecHttpCall,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintID: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},

		{
			name: "when the sprint id is not provided",
			args: args{
				ctx:      context.Background(),
				sprintID: 0,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoSprintID,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			sprintService := NewSprintService(testCase.fields.c, "1.0")

			gotResult, gotResponse, err := sprintService.Get(testCase.args.ctx, testCase.args.sprintID)

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

func Test_SprintService_Create(t *testing.T) {

	payloadMocked := &model.SprintPayloadScheme{
		Name:          "Board Name Sample",
		StartDate:     "2015-04-20T01:22:00.000+10:00",
		EndDate:       "2015-04-11T15:22:00.000+10:00",
		OriginBoardID: 5,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		payload *model.SprintPayloadScheme
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
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			Err:     model.ErrNoExecHttpCall,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint",
					"",
					payloadMocked).
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

			sprintService := NewSprintService(testCase.fields.c, "1.0")

			gotResult, gotResponse, err := sprintService.Create(testCase.args.ctx, testCase.args.payload)

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

func Test_SprintService_Update(t *testing.T) {

	payloadMocked := &model.SprintPayloadScheme{
		Name:          "Board Name Sample",
		StartDate:     "2015-04-20T01:22:00.000+10:00",
		EndDate:       "2015-04-11T15:22:00.000+10:00",
		OriginBoardID: 5,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx      context.Context
		sprintID int
		payload  *model.SprintPayloadScheme
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
				ctx:      context.Background(),
				sprintID: 1001,
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/agile/1.0/sprint/1001",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/agile/1.0/sprint/1001",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			Err:     model.ErrNoExecHttpCall,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/agile/1.0/sprint/1001",
					"",
					payloadMocked).
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

			sprintService := NewSprintService(testCase.fields.c, "1.0")

			gotResult, gotResponse, err := sprintService.Update(testCase.args.ctx, testCase.args.sprintID, testCase.args.payload)

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

func Test_SprintService_Path(t *testing.T) {

	payloadMocked := &model.SprintPayloadScheme{
		Name:          "Board Name Sample",
		StartDate:     "2015-04-20T01:22:00.000+10:00",
		EndDate:       "2015-04-11T15:22:00.000+10:00",
		OriginBoardID: 5,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx      context.Context
		sprintID int
		payload  *model.SprintPayloadScheme
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
				ctx:      context.Background(),
				sprintID: 1001,
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			Err:     model.ErrNoExecHttpCall,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					"",
					payloadMocked).
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

			sprintService := NewSprintService(testCase.fields.c, "1.0")

			gotResult, gotResponse, err := sprintService.Path(testCase.args.ctx, testCase.args.sprintID, testCase.args.payload)

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

func Test_SprintService_Delete(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx      context.Context
		sprintID int
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
				ctx:      context.Background(),
				sprintID: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/agile/1.0/sprint/1001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the sprintId is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoSprintID,
			wantErr: true,
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/agile/1.0/sprint/1001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			Err:     model.ErrNoExecHttpCall,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/agile/1.0/sprint/1001",
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

			sprintService := NewSprintService(testCase.fields.c, "1.0")

			gotResponse, err := sprintService.Delete(testCase.args.ctx, testCase.args.sprintID)

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

func Test_SprintService_Issues(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                 context.Context
		sprintID            int
		opts                *model.IssueOptionScheme
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
			name: "when the parameters are correct",
			args: args{
				ctx:      context.Background(),
				sprintID: 10001,
				opts: &model.IssueOptionScheme{
					JQL:           "project = ABC",
					ValidateQuery: true,
					Fields:        []string{"summary", "status"},
					Expand:        []string{"changelog"},
				},
				startAt:    100,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001/issue?expand=changelog&fields=summary%2Cstatus&jql=project+%3D+ABC&maxResults=50&startAt=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintIssuePageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintID: 10001,
				opts: &model.IssueOptionScheme{
					JQL:           "project = ABC",
					ValidateQuery: true,
					Fields:        []string{"summary", "status"},
					Expand:        []string{"changelog"},
				},
				startAt:    100,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001/issue?expand=changelog&fields=summary%2Cstatus&jql=project+%3D+ABC&maxResults=50&startAt=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintIssuePageScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			Err:     model.ErrNoExecHttpCall,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintID: 10001,
				opts: &model.IssueOptionScheme{
					JQL:           "project = ABC",
					ValidateQuery: true,
					Fields:        []string{"summary", "status"},
					Expand:        []string{"changelog"},
				},
				startAt:    100,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001/issue?expand=changelog&fields=summary%2Cstatus&jql=project+%3D+ABC&maxResults=50&startAt=100",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},

		{
			name: "when the sprint id is not provided",
			args: args{
				ctx:      context.Background(),
				sprintID: 0,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoSprintID,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			sprintService := NewSprintService(testCase.fields.c, "1.0")

			gotResult, gotResponse, err := sprintService.Issues(testCase.args.ctx, testCase.args.sprintID, testCase.args.opts,
				testCase.args.startAt, testCase.args.maxResults)

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

func Test_SprintService_Start(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx      context.Context
		sprintID int
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
				ctx:      context.Background(),
				sprintID: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					"",
					&model.SprintPayloadScheme{State: "Active"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the sprintId is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoSprintID,
			wantErr: true,
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					"",
					&model.SprintPayloadScheme{State: "Active"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			Err:     model.ErrNoExecHttpCall,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					"",
					&model.SprintPayloadScheme{State: "Active"}).
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

			sprintService := NewSprintService(testCase.fields.c, "1.0")

			gotResponse, err := sprintService.Start(testCase.args.ctx, testCase.args.sprintID)

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

func Test_SprintService_Close(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx      context.Context
		sprintID int
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
				ctx:      context.Background(),
				sprintID: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					"",
					&model.SprintPayloadScheme{State: "Closed"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the sprintId is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoSprintID,
			wantErr: true,
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					"",
					&model.SprintPayloadScheme{State: "Closed"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			Err:     model.ErrNoExecHttpCall,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					"",
					&model.SprintPayloadScheme{State: "Closed"}).
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

			sprintService := NewSprintService(testCase.fields.c, "1.0")

			gotResponse, err := sprintService.Close(testCase.args.ctx, testCase.args.sprintID)

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

func Test_SprintService_Move(t *testing.T) {

	payloadMocked := &model.SprintMovePayloadScheme{
		Issues:            []string{"DUMMY-1", "DUMMY-2"},
		RankBeforeIssue:   "DUMMY-4",
		RankAfterIssue:    "DUMMY-12",
		RankCustomFieldID: 10521,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx      context.Context
		sprintID int
		payload  *model.SprintMovePayloadScheme
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
				ctx:      context.Background(),
				sprintID: 1001,
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"/rest/agile/1.0/sprint/1001/issue",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the sprintId is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoSprintID,
			wantErr: true,
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"/rest/agile/1.0/sprint/1001/issue",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			Err:     model.ErrNoExecHttpCall,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintID: 1001,
				payload:  payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"/rest/agile/1.0/sprint/1001/issue",
					"",
					payloadMocked).
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

			sprintService := NewSprintService(testCase.fields.c, "1.0")

			gotResponse, err := sprintService.Move(testCase.args.ctx, testCase.args.sprintID, testCase.args.payload)

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
