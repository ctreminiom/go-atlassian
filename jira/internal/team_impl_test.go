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

func Test_internalTeamServiceImpl_Gets(t *testing.T) {

	payloadMocked := &struct {
		MaxResults int "json:\"maxResults\""
	}{MaxResults: 1000}

	type fields struct {
		c       service.Client
		version string
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
				ctx:        context.TODO(),
				maxResults: 1000,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/find",
					bytes.NewReader([]byte{})).
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
				ctx:        context.TODO(),
				maxResults: 1000,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/find",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the http call returns an error",
			args: args{
				ctx:        context.TODO(),
				maxResults: 1000,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/find",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.JiraTeamPageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to execute the http call"),
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

				assert.EqualError(t, err, testCase.Err.Error())

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
				PersonId: 6,
			},
		},
	}

	type fields struct {
		c       service.Client
		version string
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
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/create",
					bytes.NewReader([]byte{})).
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
			name: "when the payload is not provided",
			args: args{
				ctx:     context.TODO(),
				payload: nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.JiraTeamCreatePayloadScheme)(nil)).
					Return(nil, model.ErrNilPayloadError)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrNilPayloadError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/create",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/teams/1.0/teams/create",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.JiraTeamCreateResponseScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http transition"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to execute the http transition"),
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

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})
	}
}
