package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalOrganizationDirectoryServiceImpl_Activity(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		organizationID, accountID string
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
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				accountID:      "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-id-sample/directory/users/account-id-sample/last-active-dates",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.UserProductAccessScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminAccountID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				accountID:      "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"admin/v1/orgs/organization-id-sample/directory/users/account-id-sample/last-active-dates",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newOrganizationDirectoryService := NewOrganizationDirectoryService(testCase.fields.c)

			gotResult, gotResponse, err := newOrganizationDirectoryService.Activity(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.accountID)

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

func Test_internalOrganizationDirectoryServiceImpl_Remove(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		organizationID, accountID string
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
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				accountID:      "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"admin/v1/orgs/organization-id-sample/directory/users/account-id-sample",
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
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminAccountID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				accountID:      "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"admin/v1/orgs/organization-id-sample/directory/users/account-id-sample",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newOrganizationDirectoryService := NewOrganizationDirectoryService(testCase.fields.c)

			gotResponse, err := newOrganizationDirectoryService.Remove(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.accountID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}

		})
	}
}

func Test_internalOrganizationDirectoryServiceImpl_Suspend(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		organizationID, accountID string
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
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				accountID:      "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"admin/v1/orgs/organization-id-sample/directory/users/account-id-sample/suspend-access",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.GenericActionSuccessScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminAccountID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				accountID:      "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"admin/v1/orgs/organization-id-sample/directory/users/account-id-sample/suspend-access",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newOrganizationDirectoryService := NewOrganizationDirectoryService(testCase.fields.c)

			gotResult, gotResponse, err := newOrganizationDirectoryService.Suspend(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.accountID)

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

func Test_internalOrganizationDirectoryServiceImpl_Restore(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		organizationID, accountID string
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
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				accountID:      "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"admin/v1/orgs/organization-id-sample/directory/users/account-id-sample/restore-access",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.GenericActionSuccessScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the organization id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminOrganization,
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminAccountID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: "organization-id-sample",
				accountID:      "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"admin/v1/orgs/organization-id-sample/directory/users/account-id-sample/restore-access",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newOrganizationDirectoryService := NewOrganizationDirectoryService(testCase.fields.c)

			gotResult, gotResponse, err := newOrganizationDirectoryService.Restore(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.accountID)

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
