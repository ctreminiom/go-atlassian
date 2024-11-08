package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestApplicationRoleService_Get(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx context.Context
		key string
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
			name: "when the api version is v2",
			args: args{
				ctx: context.Background(),
				key: "jira-users",
			},
			fields: fields{version: "2"},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/applicationrole/jira-users",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ApplicationRoleScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api version is v3",
			args: args{
				ctx: context.Background(),
				key: "jira-users",
			},
			fields: fields{version: "3"},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/applicationrole/jira-users",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ApplicationRoleScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the application role is not provided",
			args: args{
				ctx: context.Background(),
				key: "",
			},
			fields: fields{version: "2"},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoApplicationRole,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				key: "jira-users",
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/applicationrole/jira-users",
					"",
					nil).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("unable to create the http request"),
		},

		{
			name:   "when the http call cannot be executed",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				key: "jira-users",
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/applicationrole/jira-users",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ApplicationRoleScheme{}).
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

			applicationService, err := NewApplicationRoleService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := applicationService.Get(testCase.args.ctx, testCase.args.key)

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

func TestApplicationRoleService_Gets(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx context.Context
	}

	testCases := []struct {
		name    string
		fields  fields
		on      func(*fields)
		args    args
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/applicationrole",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/applicationrole",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},

		{
			name:   "when the request cannot be created",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/applicationrole",
					"",
					nil).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     errors.New("unable to create the http request"),
		},

		{
			name:   "when the http call cannot be executed",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/applicationrole",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     errors.New("error, unable to execute the http call"),
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/applicationrole",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			applicationService, err := NewApplicationRoleService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := applicationService.Gets(testCase.args.ctx)

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

func TestNewApplicationRoleService(t *testing.T) {

	type args struct {
		client  service.Connector
		version string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
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
			err:     model.ErrNoVersionProvided,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := NewApplicationRoleService(testCase.args.client, testCase.args.version)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, got, nil)
			}
		})
	}
}
