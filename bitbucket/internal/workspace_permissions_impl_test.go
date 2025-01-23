package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalWorkspacePermissionServiceImpl_Members(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx       context.Context
		workspace string
		query     string
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
				ctx:       context.Background(),
				workspace: "work-space-name-sample",
				query:     "permission=\"owner\"",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"2.0/workspaces/work-space-name-sample/permissions?q=permission%3D%22owner%22",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.WorkspaceMembershipPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				workspace: "work-space-name-sample",
				query:     "permission=\"owner\"",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"2.0/workspaces/work-space-name-sample/permissions?q=permission%3D%22owner%22",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the workspace is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "",
			},
			wantErr: true,
			Err:     model.ErrNoWorkspace,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewWorkspacePermissionService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Members(testCase.args.ctx, testCase.args.workspace, testCase.args.query)

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

func Test_internalWorkspacePermissionServiceImpl_Repositories(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx       context.Context
		workspace string
		query     string
		sort      string
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
				ctx:       context.Background(),
				workspace: "work-space-name-sample",
				query:     "permission=\"owner\"",
				sort:      "user.display_name",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"2.0/workspaces/work-space-name-sample/permissions/repositories?q=permission%3D%22owner%22&sort=user.display_name",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RepositoryPermissionPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.Background(),
				workspace: "work-space-name-sample",
				query:     "permission=\"owner\"",
				sort:      "user.display_name",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"2.0/workspaces/work-space-name-sample/permissions/repositories?q=permission%3D%22owner%22&sort=user.display_name",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the workspace is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "",
			},
			wantErr: true,
			Err:     model.ErrNoWorkspace,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewWorkspacePermissionService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Repositories(testCase.args.ctx, testCase.args.workspace, testCase.args.query,
				testCase.args.sort)

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

func Test_internalWorkspacePermissionServiceImpl_Repository(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx        context.Context
		workspace  string
		repository string
		query      string
		sort       string
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
				workspace:  "work-space-name-sample",
				query:      "permission=\"owner\"",
				sort:       "user.display_name",
				repository: "microservice-a",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"2.0/workspaces/work-space-name-sample/permissions/repositories/microservice-a?q=permission%3D%22owner%22&sort=user.display_name",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RepositoryPermissionPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:        context.Background(),
				workspace:  "work-space-name-sample",
				query:      "permission=\"owner\"",
				sort:       "user.display_name",
				repository: "microservice-a",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"2.0/workspaces/work-space-name-sample/permissions/repositories/microservice-a?q=permission%3D%22owner%22&sort=user.display_name",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the workspace is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "",
			},
			wantErr: true,
			Err:     model.ErrNoWorkspace,
		},

		{
			name: "when the repository is not provided",
			args: args{
				ctx:       context.Background(),
				workspace: "work-space-name-sample",
			},
			wantErr: true,
			Err:     model.ErrNoRepository,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewWorkspacePermissionService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Repository(testCase.args.ctx, testCase.args.workspace, testCase.args.repository,
				testCase.args.query, testCase.args.sort)

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
