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

func Test_internalGroupServiceImpl_Create(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx       context.Context
		groupName string
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/group",
					"",
					map[string]interface{}{"name": "jira-users"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.GroupScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/group",
					"",
					map[string]interface{}{"name": "jira-users"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.GroupScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the group name is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				groupName: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoGroupName,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/group",
					"",
					map[string]interface{}{"name": "jira-users"}).
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

			groupService, err := NewGroupService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := groupService.Create(testCase.args.ctx, testCase.args.groupName)

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

func Test_internalGroupServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx       context.Context
		groupName string
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/group?groupname=jira-users",
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/group?groupname=jira-users",
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
			name:   "when the group name is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				groupName: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoGroupName,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/group?groupname=jira-users",
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

			groupService, err := NewGroupService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := groupService.Delete(testCase.args.ctx, testCase.args.groupName)

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

func Test_internalGroupServiceImpl_Remove(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx       context.Context
		groupName string
		accountID string
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
				accountID: "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/group/user?accountId=account-id-sample&groupname=jira-users",
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
				accountID: "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/group/user?accountId=account-id-sample&groupname=jira-users",
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
			name:   "when the group name is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				groupName: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoGroupName,
		},

		{
			name:   "when the account id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
				accountID: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoAccountID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
				accountID: "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/group/user?accountId=account-id-sample&groupname=jira-users",
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

			groupService, err := NewGroupService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := groupService.Remove(testCase.args.ctx, testCase.args.groupName, testCase.args.accountID)

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

func Test_internalGroupServiceImpl_Add(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx       context.Context
		groupName string
		accountID string
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
				accountID: "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/group/user?groupname=jira-users",
					"",
					map[string]interface{}{"accountId": "account-id-sample"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.GroupScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
				accountID: "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/group/user?groupname=jira-users",
					"",
					map[string]interface{}{"accountId": "account-id-sample"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.GroupScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the group name is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				groupName: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoGroupName,
		},

		{
			name:   "when the account id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
				accountID: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoAccountID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:       context.Background(),
				groupName: "jira-users",
				accountID: "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/group/user?groupname=jira-users",
					"",
					map[string]interface{}{"accountId": "account-id-sample"}).
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

			groupService, err := NewGroupService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := groupService.Add(testCase.args.ctx, testCase.args.groupName, testCase.args.accountID)

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

func Test_internalGroupServiceImpl_Bulk(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		options             *model.GroupBulkOptionsScheme
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				options: &model.GroupBulkOptionsScheme{
					GroupIDs:   []string{"1001", "1002"},
					GroupNames: []string{"jira-users", "confluence-users"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/group/bulk?groupId=1001&groupId=1002&groupName=jira-users&groupName=confluence-users&maxResults=50&startAt=0",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BulkGroupScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
				options: &model.GroupBulkOptionsScheme{
					GroupIDs:   []string{"1001", "1002"},
					GroupNames: []string{"jira-users", "confluence-users"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/group/bulk?groupId=1001&groupId=1002&groupName=jira-users&groupName=confluence-users&maxResults=50&startAt=0",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BulkGroupScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
				options: &model.GroupBulkOptionsScheme{
					GroupIDs:   []string{"1001", "1002"},
					GroupNames: []string{"jira-users", "confluence-users"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/group/bulk?groupId=1001&groupId=1002&groupName=jira-users&groupName=confluence-users&maxResults=50&startAt=0",
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

			groupService, err := NewGroupService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := groupService.Bulk(testCase.args.ctx, testCase.args.options, testCase.args.startAt,
				testCase.args.maxResults)

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

func Test_internalGroupServiceImpl_Members(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		groupName           string
		inactive            bool
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.Background(),
				groupName:  "jira-users",
				inactive:   true,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/group/member?groupname=jira-users&includeInactiveUsers=true&maxResults=50&startAt=0",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.GroupMemberPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.Background(),
				groupName:  "jira-users",
				inactive:   true,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/group/member?groupname=jira-users&includeInactiveUsers=true&maxResults=50&startAt=0",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.GroupMemberPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the group name is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:        context.Background(),
				groupName:  "",
				inactive:   true,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoGroupName,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:        context.Background(),
				groupName:  "jira-users",
				inactive:   true,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/group/member?groupname=jira-users&includeInactiveUsers=true&maxResults=50&startAt=0",
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

			groupService, err := NewGroupService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := groupService.Members(testCase.args.ctx, testCase.args.groupName, testCase.args.inactive,
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

func Test_NewGroupService(t *testing.T) {

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
			got, err := NewGroupService(testCase.args.client, testCase.args.version)

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
