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

func Test_internalCustomerImpl_Create(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                context.Context
		email, displayName string
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
				ctx:         context.Background(),
				email:       "carlos.treminio@example.com",
				displayName: "Carlos T",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/customer",
					"",
					map[string]interface{}{"displayName": "Carlos T", "email": "carlos.treminio@example.com"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:         context.Background(),
				email:       "carlos.treminio@example.com",
				displayName: "Carlos T",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/customer",
					"",
					map[string]interface{}{"displayName": "Carlos T", "email": "carlos.treminio@example.com"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:         context.Background(),
				email:       "carlos.treminio@example.com",
				displayName: "Carlos T",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/customer",
					"",
					map[string]interface{}{"displayName": "Carlos T", "email": "carlos.treminio@example.com"}).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			customerService := NewCustomerService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := customerService.Create(testCase.args.ctx, testCase.args.email, testCase.args.displayName)

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

func Test_internalCustomerImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx           context.Context
		serviceDeskID int
		query         string
		start, limit  int
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
				ctx:           context.Background(),
				serviceDeskID: 10001,
				query:         "Carlos T",
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/customer?limit=50&query=Carlos+T&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				query:         "Carlos T",
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/customer?limit=50&query=Carlos+T&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.CustomerPageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				query:         "Carlos T",
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/customer?limit=50&query=Carlos+T&start=100",
					"",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			customerService := NewCustomerService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := customerService.Gets(testCase.args.ctx, testCase.args.serviceDeskID, testCase.args.query,
				testCase.args.start, testCase.args.limit)

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

func Test_internalCustomerImpl_Add(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx           context.Context
		serviceDeskID int
		accountIDs    []string
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
				ctx:           context.Background(),
				serviceDeskID: 10001,
				accountIDs:    []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/customer",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				accountIDs:    []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/customer",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				accountIDs:    []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/customer",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the service desk id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoServiceDeskIDError,
		},

		{
			name: "when the account ids are not provided",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoAccountSliceError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			customerService := NewCustomerService(testCase.fields.c, "latest")

			gotResponse, err := customerService.Add(testCase.args.ctx, testCase.args.serviceDeskID, testCase.args.accountIDs)

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

func Test_internalCustomerImpl_Remove(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx           context.Context
		serviceDeskID int
		accountIDs    []string
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
				ctx:           context.Background(),
				serviceDeskID: 10001,
				accountIDs:    []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/servicedesk/10001/customer",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				accountIDs:    []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/servicedesk/10001/customer",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				accountIDs:    []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/servicedesk/10001/customer",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the service desk id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoServiceDeskIDError,
		},

		{
			name: "when the account ids are not provided",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoAccountSliceError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			customerService := NewCustomerService(testCase.fields.c, "latest")

			gotResponse, err := customerService.Remove(testCase.args.ctx, testCase.args.serviceDeskID, testCase.args.accountIDs)

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
