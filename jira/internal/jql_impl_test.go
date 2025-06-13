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

func Test_internalJQLServiceImpl_Parse(t *testing.T) {

	payloadMocked := map[string]interface{}{
		"queries": []string{
			"summary ~ test AND (labels in (urgent, blocker) OR lastCommentedBy = currentUser()) AND status CHANGED AFTER startOfMonth(-1M) ORDER BY updated DESC",
			"invalid query",
			"summary = test",
			"summary in test",
			"project = INVALID",
			"universe = 42",
		}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx            context.Context
		validationType string
		JqlQueries     []string
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
				validationType: "strict",
				JqlQueries: []string{
					"summary ~ test AND (labels in (urgent, blocker) OR lastCommentedBy = currentUser()) AND status CHANGED AFTER startOfMonth(-1M) ORDER BY updated DESC",
					"invalid query",
					"summary = test",
					"summary in test",
					"project = INVALID",
					"universe = 42"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"/rest/api/3/jql/parse?validation=strict",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ParsedQueryPageScheme{}).
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
				validationType: "strict",
				JqlQueries: []string{
					"summary ~ test AND (labels in (urgent, blocker) OR lastCommentedBy = currentUser()) AND status CHANGED AFTER startOfMonth(-1M) ORDER BY updated DESC",
					"invalid query",
					"summary = test",
					"summary in test",
					"project = INVALID",
					"universe = 42"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"/rest/api/2/jql/parse?validation=strict",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ParsedQueryPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.Background(),
				validationType: "strict",
				JqlQueries: []string{
					"summary ~ test AND (labels in (urgent, blocker) OR lastCommentedBy = currentUser()) AND status CHANGED AFTER startOfMonth(-1M) ORDER BY updated DESC",
					"invalid query",
					"summary = test",
					"summary in test",
					"project = INVALID",
					"universe = 42"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"/rest/api/3/jql/parse?validation=strict",
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

			fieldService, err := NewJQLService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldService.Parse(testCase.args.ctx, testCase.args.validationType,
				testCase.args.JqlQueries)

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

func Test_NewJQLService(t *testing.T) {

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
			got, err := NewJQLService(testCase.args.client, testCase.args.version)

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
