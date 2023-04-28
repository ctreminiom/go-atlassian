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

func Test_internalUserImpl_Permissions(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx        context.Context
		accountID  string
		privileges []string
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
				accountID:  "account-id-sample",
				privileges: []string{"privileges-sample"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"users/account-id-sample/manage?privileges=privileges-sample",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.AdminUserPermissionScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:       context.TODO(),
				accountID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminAccountIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:        context.TODO(),
				accountID:  "account-id-sample",
				privileges: []string{"privileges-sample"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"users/account-id-sample/manage?privileges=privileges-sample",
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

			service := NewUserService(testCase.fields.c, nil)

			gotResult, gotResponse, err := service.Permissions(testCase.args.ctx, testCase.args.accountID, testCase.args.privileges)

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

func Test_internalUserImpl_Get(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx       context.Context
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
			name: "when the parameters are correct",
			args: args{
				ctx:       context.TODO(),
				accountID: "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"users/account-id-sample/manage/profile",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.AdminUserScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:       context.TODO(),
				accountID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminAccountIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.TODO(),
				accountID: "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"users/account-id-sample/manage/profile",
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

			service := NewUserService(testCase.fields.c, nil)

			gotResult, gotResponse, err := service.Get(testCase.args.ctx, testCase.args.accountID)

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

func Test_internalUserImpl_Enable(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx       context.Context
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
			name: "when the parameters are correct",
			args: args{
				ctx:       context.TODO(),
				accountID: "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"users/account-id-sample/manage/lifecycle/enable",
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
			name: "when the account id is not provided",
			args: args{
				ctx:       context.TODO(),
				accountID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminAccountIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.TODO(),
				accountID: "account-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"users/account-id-sample/manage/lifecycle/enable",
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

			service := NewUserService(testCase.fields.c, nil)

			gotResponse, err := service.Enable(testCase.args.ctx, testCase.args.accountID)

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

func Test_internalUserImpl_Disable(t *testing.T) {

	payloadMocked := &struct {
		Message string "json:\"message\""
	}{Message: "Your account has been disabled :("}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                context.Context
		accountID, message string
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
				ctx:       context.TODO(),
				accountID: "account-id-sample",
				message:   "Your account has been disabled :(",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"users/account-id-sample/manage/lifecycle/disable",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:       context.TODO(),
				accountID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminAccountIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.TODO(),
				accountID: "account-id-sample",
				message:   "Your account has been disabled :(",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"users/account-id-sample/manage/lifecycle/disable",
					bytes.NewReader([]byte{})).
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

			service := NewUserService(testCase.fields.c, nil)

			gotResponse, err := service.Disable(testCase.args.ctx, testCase.args.accountID, testCase.args.message)

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

func Test_internalUserImpl_Update(t *testing.T) {

	payloadMocked := map[string]interface{}{"nickname": "marshmallow"}

	var payload = make(map[string]interface{})
	payload["nickname"] = "marshmallow"

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx       context.Context
		accountID string
		payload   map[string]interface{}
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
				ctx:       context.TODO(),
				accountID: "account-id-sample",
				payload:   payload,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPatch,
					"users/account-id-sample/manage/profile",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.AdminUserScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:       context.TODO(),
				accountID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminAccountIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:       context.TODO(),
				accountID: "account-id-sample",
				payload:   payload,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPatch,
					"users/account-id-sample/manage/profile",
					bytes.NewReader([]byte{})).
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

			service := NewUserService(testCase.fields.c, nil)

			gotResult, gotResponse, err := service.Update(testCase.args.ctx, testCase.args.accountID, testCase.args.payload)

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
