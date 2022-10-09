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

func Test_internalRestrictionOperationUserImpl_Get(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                                context.Context
		contentID, operationKey, accountID string
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
				ctx:          context.TODO(),
				contentID:    "100001",
				operationKey: "read",
				accountID:    "06db0c76-115b-498e-9cd6-921d6f6dde46",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100001/restriction/byOperation/read/user?accountId=06db0c76-115b-498e-9cd6-921d6f6dde46",
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
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.TODO(),
				contentID:    "100001",
				operationKey: "read",
				accountID:    "06db0c76-115b-498e-9cd6-921d6f6dde46",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100001/restriction/byOperation/read/user?accountId=06db0c76-115b-498e-9cd6-921d6f6dde46",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},

		{
			name: "when the property key is not provided",
			args: args{
				ctx:       context.TODO(),
				contentID: "1111",
			},
			wantErr: true,
			Err:     model.ErrNoContentRestrictionKeyError,
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:          context.TODO(),
				contentID:    "1111",
				operationKey: "read",
			},
			wantErr: true,
			Err:     model.ErrNoAccountIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewRestrictionOperationUserService(testCase.fields.c)

			gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.contentID, testCase.args.operationKey,
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

func Test_internalRestrictionOperationUserImpl_Add(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                                context.Context
		contentID, operationKey, accountID string
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
				ctx:          context.TODO(),
				contentID:    "100001",
				operationKey: "read",
				accountID:    "06db0c76-115b-498e-9cd6-921d6f6dde46",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/100001/restriction/byOperation/read/user?accountId=06db0c76-115b-498e-9cd6-921d6f6dde46",
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
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.TODO(),
				contentID:    "100001",
				operationKey: "read",
				accountID:    "06db0c76-115b-498e-9cd6-921d6f6dde46",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/100001/restriction/byOperation/read/user?accountId=06db0c76-115b-498e-9cd6-921d6f6dde46",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},

		{
			name: "when the property key is not provided",
			args: args{
				ctx:       context.TODO(),
				contentID: "1111",
			},
			wantErr: true,
			Err:     model.ErrNoContentRestrictionKeyError,
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:          context.TODO(),
				contentID:    "1111",
				operationKey: "read",
			},
			wantErr: true,
			Err:     model.ErrNoAccountIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewRestrictionOperationUserService(testCase.fields.c)

			gotResponse, err := newService.Add(testCase.args.ctx, testCase.args.contentID, testCase.args.operationKey,
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

func Test_internalRestrictionOperationUserImpl_Remove(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                                context.Context
		contentID, operationKey, accountID string
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
				ctx:          context.TODO(),
				contentID:    "100001",
				operationKey: "read",
				accountID:    "06db0c76-115b-498e-9cd6-921d6f6dde46",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/content/100001/restriction/byOperation/read/user?accountId=06db0c76-115b-498e-9cd6-921d6f6dde46",
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
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.TODO(),
				contentID:    "100001",
				operationKey: "read",
				accountID:    "06db0c76-115b-498e-9cd6-921d6f6dde46",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/content/100001/restriction/byOperation/read/user?accountId=06db0c76-115b-498e-9cd6-921d6f6dde46",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},

		{
			name: "when the property key is not provided",
			args: args{
				ctx:       context.TODO(),
				contentID: "1111",
			},
			wantErr: true,
			Err:     model.ErrNoContentRestrictionKeyError,
		},

		{
			name: "when the account id is not provided",
			args: args{
				ctx:          context.TODO(),
				contentID:    "1111",
				operationKey: "read",
			},
			wantErr: true,
			Err:     model.ErrNoAccountIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewRestrictionOperationUserService(testCase.fields.c)

			gotResponse, err := newService.Remove(testCase.args.ctx, testCase.args.contentID, testCase.args.operationKey,
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
