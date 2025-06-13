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

func Test_internalTeamServiceImpl_Gets(t *testing.T) {

	payloadMocked := map[string]interface{}{"maxResults": 1000}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx        context.Context
		maxResults int
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
				ctx:        context.Background(),
				maxResults: 1000,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/find",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.JiraTeamPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:        context.Background(),
				maxResults: 1000,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/find",
					"", payloadMocked).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name: "when the http call returns an error",
			args: args{
				ctx:        context.Background(),
				maxResults: 1000,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/find",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.JiraTeamPageScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrNoExecHttpCall,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			teamService := NewTeamService(testCase.fields.c)

			gotResult, gotResponse, err := teamService.Gets(testCase.args.ctx, testCase.args.maxResults)

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

func Test_internalTeamServiceImpl_Create(t *testing.T) {

	payloadMocked := &model.JiraTeamCreatePayloadScheme{
		Title:     "Team Name Sample",
		Shareable: true,
		Resources: []*model.JiraTeamResourceScheme{
			{
				PersonID: 6,
			},
		},
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		payload *model.JiraTeamCreatePayloadScheme
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
					"rest/teams/1.0/teams/create",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.JiraTeamCreateResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/create",
					"", payloadMocked).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/create",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.JiraTeamCreateResponseScheme{}).
					Return(&model.ResponseScheme{}, model.ErrHttpTransition)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrHttpTransition,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			teamService := NewTeamService(testCase.fields.c)

			gotResult, gotResponse, err := teamService.Create(testCase.args.ctx, testCase.args.payload)

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
