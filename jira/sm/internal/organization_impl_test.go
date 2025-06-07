package internal

import (
	"context"
	"encoding/json"
	"errors"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_internalOrganizationImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx          context.Context
		accountID    string
		start, limit int
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
				accountID: "account-id-sample",
				start:     100,
				limit:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/organization?accountId=account-id-sample&limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:       context.Background(),
				accountID: "account-id-sample",
				start:     100,
				limit:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/organization?accountId=account-id-sample&limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationPageScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:       context.Background(),
				accountID: "account-id-sample",
				start:     100,
				limit:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/organization?accountId=account-id-sample&limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			organizationService := NewOrganizationService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := organizationService.Gets(testCase.args.ctx, testCase.args.accountID,
				testCase.args.start, testCase.args.limit)

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

func Test_internalOrganizationImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx            context.Context
		organizationID int
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
				organizationID: 100001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/organization/100001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:            context.Background(),
				organizationID: 100001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/organization/100001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: 100001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/organization/100001",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			organizationService := NewOrganizationService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := organizationService.Get(testCase.args.ctx, testCase.args.organizationID)

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

func Test_internalOrganizationImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx            context.Context
		organizationID int
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
				organizationID: 100001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/organization/100001",
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
			name: "when the http call cannot be executed",
			args: args{
				ctx:            context.Background(),
				organizationID: 100001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/organization/100001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: 100001,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/organization/100001",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			organizationService := NewOrganizationService(testCase.fields.c, "latest")

			gotResponse, err := organizationService.Delete(testCase.args.ctx, testCase.args.organizationID)

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

func Test_internalOrganizationImpl_Create(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx  context.Context
		name string
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
				ctx:  context.Background(),
				name: "Client A - Organization Users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/organization",
					"",
					map[string]interface{}{"name": "Client A - Organization Users"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:  context.Background(),
				name: "Client A - Organization Users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/organization",
					"",
					map[string]interface{}{"name": "Client A - Organization Users"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:  context.Background(),
				name: "Client A - Organization Users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/organization",
					"",
					map[string]interface{}{"name": "Client A - Organization Users"}).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			organizationService := NewOrganizationService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := organizationService.Create(testCase.args.ctx, testCase.args.name)

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

func Test_internalOrganizationImpl_Users(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                          context.Context
		organizationID, start, limit int
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
				organizationID: 10001,
				start:          100,
				limit:          50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/organization/10001/user?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationUsersPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:            context.Background(),
				organizationID: 10001,
				start:          100,
				limit:          50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/organization/10001/user?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationUsersPageScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: 10001,
				start:          100,
				limit:          50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/organization/10001/user?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			organizationService := NewOrganizationService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := organizationService.Users(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.start, testCase.args.limit)

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

func Test_internalOrganizationImpl_Add(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx            context.Context
		organizationID int
		accountIDs     []string
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
				organizationID: 10001,
				accountIDs:     []string{"account-id-sample", "account-id-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/organization/10001/user",
					"",
					map[string]interface{}{"accountIds": []string{"account-id-sample", "account-id-sample-2"}}).
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
				ctx:            context.Background(),
				organizationID: 10001,
				accountIDs:     []string{"account-id-sample", "account-id-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/organization/10001/user",
					"",
					map[string]interface{}{"accountIds": []string{"account-id-sample", "account-id-sample-2"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: 10001,
				accountIDs:     []string{"account-id-sample", "account-id-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/organization/10001/user",
					"",
					map[string]interface{}{"accountIds": []string{"account-id-sample", "account-id-sample-2"}}).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			organizationService := NewOrganizationService(testCase.fields.c, "latest")

			gotResponse, err := organizationService.Add(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.accountIDs)

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

func Test_internalOrganizationImpl_Remove(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx            context.Context
		organizationID int
		accountIDs     []string
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
				organizationID: 10001,
				accountIDs:     []string{"account-id-sample", "account-id-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/organization/10001/user",
					"",
					map[string]interface{}{"accountIds": []string{"account-id-sample", "account-id-sample-2"}}).
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
				ctx:            context.Background(),
				organizationID: 10001,
				accountIDs:     []string{"account-id-sample", "account-id-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/organization/10001/user",
					"",
					map[string]interface{}{"accountIds": []string{"account-id-sample", "account-id-sample-2"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: 10001,
				accountIDs:     []string{"account-id-sample", "account-id-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/organization/10001/user",
					"",
					map[string]interface{}{"accountIds": []string{"account-id-sample", "account-id-sample-2"}}).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			organizationService := NewOrganizationService(testCase.fields.c, "latest")

			gotResponse, err := organizationService.Remove(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.accountIDs)

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

func Test_internalOrganizationImpl_Project(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                         context.Context
		accountID                   string
		serviceDeskID, start, limit int
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
				accountID:     "account-id-sample",
				serviceDeskID: 1001,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/1001/organization?accountId=account-id-sample&limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				accountID:     "account-id-sample",
				serviceDeskID: 1001,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/1001/organization?accountId=account-id-sample&limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.OrganizationPageScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				accountID:     "account-id-sample",
				serviceDeskID: 1001,
				start:         100,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/1001/organization?accountId=account-id-sample&limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			organizationService := NewOrganizationService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := organizationService.Project(testCase.args.ctx, testCase.args.accountID,
				testCase.args.serviceDeskID,
				testCase.args.start, testCase.args.limit)

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

func Test_internalOrganizationImpl_Associate(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                           context.Context
		serviceDeskID, organizationID int
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
				organizationID: 10001,
				serviceDeskID:  100022,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/organization",
					"",
					map[string]interface{}{"organizationId": 100022}).
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
				ctx:            context.Background(),
				organizationID: 10001,
				serviceDeskID:  100022,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/organization",
					"",
					map[string]interface{}{"organizationId": 100022}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: 10001,
				serviceDeskID:  100022,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/servicedesk/10001/organization",
					"",
					map[string]interface{}{"organizationId": 100022}).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			organizationService := NewOrganizationService(testCase.fields.c, "latest")

			gotResponse, err := organizationService.Associate(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.serviceDeskID)

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

func Test_internalOrganizationImpl_Detach(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                           context.Context
		serviceDeskID, organizationID int
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
				organizationID: 10001,
				serviceDeskID:  100022,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/servicedesk/10001/organization",
					"",
					map[string]interface{}{"organizationId": 100022}).
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
				ctx:            context.Background(),
				organizationID: 10001,
				serviceDeskID:  100022,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/servicedesk/10001/organization",
					"",
					map[string]interface{}{"organizationId": 100022}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoHttpResponse)

				fields.c = client
			},
			Err:     model.ErrNoHttpResponse,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:            context.Background(),
				organizationID: 10001,
				serviceDeskID:  100022,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/servicedesk/10001/organization",
					"",
					map[string]interface{}{"organizationId": 100022}).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			organizationService := NewOrganizationService(testCase.fields.c, "latest")

			gotResponse, err := organizationService.Detach(testCase.args.ctx, testCase.args.organizationID,
				testCase.args.serviceDeskID)

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
