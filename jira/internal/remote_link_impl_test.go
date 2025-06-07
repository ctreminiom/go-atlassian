package internal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalRemoteLinkImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                    context.Context
		issueKeyOrID, globalID string
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
					"rest/api/2/issue/KP-23/remotelink?globalId=system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1",
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
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				globalID:     "system=http://www.mycompany.com/support&id=1",
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
					"rest/api/3/issue/KP-23/remotelink?globalId=system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1",
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
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				globalID:     "system=http://www.mycompany.com/support&id=1",
			},
			wantErr: false,
		},

		{
			name:   "when the issue key or not provided",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the request cannot be created",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/KP-23/remotelink?globalId=system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				globalID:     "system=http://www.mycompany.com/support&id=1",
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name:   "when the http call cannot be executed",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/KP-23/remotelink?globalId=system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				globalID:     "system=http://www.mycompany.com/support&id=1",
			},
			wantErr: true,
			Err:     model.ErrNoExecHttpCall,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			applicationService, err := NewRemoteLinkService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := applicationService.Gets(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.globalID)

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

func Test_internalRemoteLinkImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                  context.Context
		issueKeyOrID, linkID string
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
					"rest/api/2/issue/KP-23/remotelink/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RemoteLinkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
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
					"rest/api/3/issue/KP-23/remotelink/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RemoteLinkScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
			},
			wantErr: false,
		},

		{
			name:   "when the issue key or not provided",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the remote link is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-3",
			},
			wantErr: true,
			Err:     model.ErrNoRemoteLinkID,
		},

		{
			name:   "when the request cannot be created",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/KP-23/remotelink/10001",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name:   "when the http call cannot be executed",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/KP-23/remotelink/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RemoteLinkScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
			},
			wantErr: true,
			Err:     model.ErrNoExecHttpCall,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			applicationService, err := NewRemoteLinkService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := applicationService.Get(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.linkID)

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

func Test_internalRemoteLinkImpl_Update(t *testing.T) {

	payloadMocked := &model.RemoteLinkScheme{
		Application: &model.RemoteLinkApplicationScheme{
			Name: "My Acme Tracker",
			Type: "com.acme.tracker",
		},
		GlobalID: "system=http://www.mycompany.com/support&id=1",
		Object: &model.RemoteLinkObjectScheme{
			Icon: &model.RemoteLinkObjectLinkScheme{
				Title:    "Support Ticket",
				URL16X16: "http://www.mycompany.com/support/ticket.png",
			},
			Status: &model.RemoteLinkObjectStatusScheme{
				Icon: &model.RemoteLinkObjectLinkScheme{
					Link:     "http://www.mycompany.com/support?id=1&details=closed",
					Title:    "Case Closed",
					URL16X16: "http://www.mycompany.com/support/resolved.png",
				},
				Resolved: true,
			},
			Summary: "Customer support issue",
			Title:   "TSTSUP-111",
			URL:     "http://www.mycompany.com/support?id=1",
		},
		Relationship: "causes",
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                  context.Context
		issueKeyOrID, linkID string
		payload              *model.RemoteLinkScheme
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
					http.MethodPut,
					"rest/api/2/issue/KP-23/remotelink/10001",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
				payload:      payloadMocked,
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
					http.MethodPut,
					"rest/api/3/issue/KP-23/remotelink/10001",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
				payload:      payloadMocked,
			},
			wantErr: false,
		},

		{
			name:   "when the issue key or not provided",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the remote link is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-3",
			},
			wantErr: true,
			Err:     model.ErrNoRemoteLinkID,
		},

		{
			name:   "when the request cannot be created",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/KP-23/remotelink/10001",
					"",
					payloadMocked).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
				payload:      payloadMocked,
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name:   "when the http call cannot be executed",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/KP-23/remotelink/10001",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
				payload:      payloadMocked,
			},
			wantErr: true,
			Err:     model.ErrNoExecHttpCall,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			applicationService, err := NewRemoteLinkService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := applicationService.Update(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.linkID,
				testCase.args.payload)

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

func Test_internalRemoteLinkImpl_Create(t *testing.T) {

	payloadMocked := &model.RemoteLinkScheme{
		Application: &model.RemoteLinkApplicationScheme{
			Name: "My Acme Tracker",
			Type: "com.acme.tracker",
		},
		GlobalID: "system=http://www.mycompany.com/support&id=1",
		Object: &model.RemoteLinkObjectScheme{
			Icon: &model.RemoteLinkObjectLinkScheme{
				Title:    "Support Ticket",
				URL16X16: "http://www.mycompany.com/support/ticket.png",
			},
			Status: &model.RemoteLinkObjectStatusScheme{
				Icon: &model.RemoteLinkObjectLinkScheme{
					Link:     "http://www.mycompany.com/support?id=1&details=closed",
					Title:    "Case Closed",
					URL16X16: "http://www.mycompany.com/support/resolved.png",
				},
				Resolved: true,
			},
			Summary: "Customer support issue",
			Title:   "TSTSUP-111",
			URL:     "http://www.mycompany.com/support?id=1",
		},
		Relationship: "causes",
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
		payload      *model.RemoteLinkScheme
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
					http.MethodPost,
					"rest/api/2/issue/KP-23/remotelink",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RemoteLinkIdentify{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				payload:      payloadMocked,
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
					http.MethodPost,
					"rest/api/3/issue/KP-23/remotelink",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RemoteLinkIdentify{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				payload:      payloadMocked,
			},
			wantErr: false,
		},

		{
			name:   "when the issue key or not provided",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the request cannot be created",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/KP-23/remotelink",
					"",
					payloadMocked).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				payload:      payloadMocked,
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name:   "when the http call cannot be executed",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/KP-23/remotelink",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RemoteLinkIdentify{}).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				payload:      payloadMocked,
			},
			wantErr: true,
			Err:     model.ErrNoExecHttpCall,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			applicationService, err := NewRemoteLinkService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := applicationService.Create(testCase.args.ctx, testCase.args.issueKeyOrID,
				testCase.args.payload)

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

func Test_internalRemoteLinkImpl_DeleteByID(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                  context.Context
		issueKeyOrID, linkID string
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
					http.MethodDelete,
					"rest/api/2/issue/KP-23/remotelink/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
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
					http.MethodDelete,
					"rest/api/3/issue/KP-23/remotelink/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
			},
			wantErr: false,
		},

		{
			name:   "when the issue key or not provided",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the remote link is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-3",
			},
			wantErr: true,
			Err:     model.ErrNoRemoteLinkID,
		},

		{
			name:   "when the request cannot be created",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issue/KP-23/remotelink/10001",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name:   "when the http call cannot be executed",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issue/KP-23/remotelink/10001",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				linkID:       "10001",
			},
			wantErr: true,
			Err:     model.ErrNoExecHttpCall,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			applicationService, err := NewRemoteLinkService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := applicationService.DeleteByID(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.linkID)

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

func Test_internalRemoteLinkImpl_DeleteByGlobalID(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                    context.Context
		issueKeyOrID, globalID string
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
					http.MethodDelete,
					"rest/api/2/issue/KP-23/remotelink?globalId=system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				globalID:     "system=http://www.mycompany.com/support&id=1",
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
					http.MethodDelete,
					"rest/api/3/issue/KP-23/remotelink?globalId=system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				globalID:     "system=http://www.mycompany.com/support&id=1",
			},
			wantErr: false,
		},

		{
			name:   "when the issue key or not provided",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the global link is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-3",
			},
			wantErr: true,
			Err:     model.ErrNoRemoteLinkGlobalID,
		},

		{
			name:   "when the request cannot be created",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issue/KP-23/remotelink?globalId=system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				globalID:     "system=http://www.mycompany.com/support&id=1",
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name:   "when the http call cannot be executed",
			fields: fields{version: "2"},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issue/KP-23/remotelink?globalId=system%3Dhttp%3A%2F%2Fwww.mycompany.com%2Fsupport%26id%3D1",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "KP-23",
				globalID:     "system=http://www.mycompany.com/support&id=1",
			},
			wantErr: true,
			Err:     model.ErrNoExecHttpCall,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			applicationService, err := NewRemoteLinkService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := applicationService.DeleteByGlobalID(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.globalID)

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
