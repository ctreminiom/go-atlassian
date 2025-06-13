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

func Test_internalProjectRoleActorImpl_Add(t *testing.T) {

	payloadMocked := map[string]interface{}{"group": []string{"jira-users"}, "user": []string{"uuid"}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                context.Context
		projectKeyOrID     string
		roleID             int
		accountIDs, groups []string
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
				ctx:            context.Background(),
				projectKeyOrID: "DUMMY",
				roleID:         10001,
				accountIDs:     []string{"uuid"},
				groups:         []string{"jira-users"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/project/DUMMY/role/10001",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectRoleScheme{}).
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
				ctx:            context.Background(),
				projectKeyOrID: "DUMMY",
				roleID:         10001,
				accountIDs:     []string{"uuid"},
				groups:         []string{"jira-users"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/project/DUMMY/role/10001",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectRoleScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the project key or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDOrKey,
		},

		{
			name:   "when the project role id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.Background(),
				projectKeyOrID: "DUMMY",
			},
			wantErr: true,
			Err:     model.ErrNoProjectRoleID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.Background(),
				projectKeyOrID: "DUMMY",
				roleID:         10001,
				accountIDs:     []string{"uuid"},
				groups:         []string{"jira-users"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/project/DUMMY/role/10001",
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

			newService, err := NewProjectRoleActorService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Add(testCase.args.ctx, testCase.args.projectKeyOrID, testCase.args.roleID,
				testCase.args.accountIDs, testCase.args.groups)

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

func Test_internalProjectRoleActorImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx              context.Context
		projectKeyOrID   string
		roleID           int
		accountID, group string
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
				ctx:            context.Background(),
				projectKeyOrID: "DUMMY",
				roleID:         10001,
				accountID:      "uuid",
				group:          "jira-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/project/DUMMY/role/10001?group=jira-users&user=uuid",
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
				ctx:            context.Background(),
				projectKeyOrID: "DUMMY",
				roleID:         10001,
				accountID:      "uuid",
				group:          "jira-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/project/DUMMY/role/10001?group=jira-users&user=uuid",
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
			name:   "when the project key or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDOrKey,
		},

		{
			name:   "when the project role id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.Background(),
				projectKeyOrID: "DUMMY",
			},
			wantErr: true,
			Err:     model.ErrNoProjectRoleID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.Background(),
				projectKeyOrID: "DUMMY",
				roleID:         10001,
				accountID:      "uuid",
				group:          "jira-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/project/DUMMY/role/10001?group=jira-users&user=uuid",
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

			newService, err := NewProjectRoleActorService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.projectKeyOrID, testCase.args.roleID,
				testCase.args.accountID, testCase.args.group)

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

func Test_NewProjectRoleActorService(t *testing.T) {

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
			got, err := NewProjectRoleActorService(testCase.args.client, testCase.args.version)

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
