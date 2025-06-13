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
	"time"
)

func Test_internalAuditRecordImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		options *model.AuditRecordGetOptions
		offSet  int
		limit   int
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
				options: &model.AuditRecordGetOptions{
					Filter: "summary",
					From:   time.Date(2015, 11, 17, 20, 34, 58, 651387237, time.UTC),
					To:     time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC),
				},
				offSet: 2000,
				limit:  1000,
			},
			fields: fields{version: "2"},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/auditing/record?=summary&from=2015-11-17&limit=1000&offset=2000&to=2019-11-17",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.AuditRecordPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api version is v3",
			args: args{
				ctx: context.Background(),
				options: &model.AuditRecordGetOptions{
					Filter: "summary",
					From:   time.Date(2015, 11, 17, 20, 34, 58, 651387237, time.UTC),
					To:     time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC),
				},
				offSet: 2000,
				limit:  1000,
			},
			fields: fields{version: "3"},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/auditing/record?=summary&from=2015-11-17&limit=1000&offset=2000&to=2019-11-17",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.AuditRecordPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				options: &model.AuditRecordGetOptions{
					Filter: "summary",
					From:   time.Date(2015, 11, 17, 20, 34, 58, 651387237, time.UTC),
					To:     time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC),
				},
				offSet: 2000,
				limit:  1000,
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/auditing/record?=summary&from=2015-11-17&limit=1000&offset=2000&to=2019-11-17",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name:   "when the http call cannot be executed",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				options: &model.AuditRecordGetOptions{
					Filter: "summary",
					From:   time.Date(2015, 11, 17, 20, 34, 58, 651387237, time.UTC),
					To:     time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC),
				},
				offSet: 2000,
				limit:  1000,
			},
			on: func(fields *fields) {
				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/auditing/record?=summary&from=2015-11-17&limit=1000&offset=2000&to=2019-11-17",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.AuditRecordPageScheme{}).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
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

			newAuditRecordService, err := NewAuditRecordService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newAuditRecordService.Get(testCase.args.ctx, testCase.args.options, testCase.args.offSet,
				testCase.args.limit)

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

func TestNewAuditRecordService(t *testing.T) {

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
			got, err := NewAuditRecordService(testCase.args.client, testCase.args.version)

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
