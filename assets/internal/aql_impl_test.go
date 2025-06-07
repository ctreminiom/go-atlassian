package internal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_internalAQLImpl_Filter(t *testing.T) {

	payloadMocked := &model.AQLSearchParamsScheme{
		Query:                 "Name LIKE Test",
		Page:                  2,
		ResultPerPage:         25,
		IncludeAttributes:     true,
		IncludeAttributesDeep: 12,
		IncludeTypeAttributes: true,
		IncludeExtendedInfo:   true,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		workspaceID string
		payload     *model.AQLSearchParamsScheme
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
				workspaceID: "workspace-uuid-sample",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/aql/objects?includeAttributes=true&includeAttributesDeep=12&includeExtendedInfo=true&includeTypeAttributes=true&page=2&qlQuery=Name+LIKE+Test&resultPerPage=25",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectListScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/aql/objects?includeAttributes=true&includeAttributesDeep=12&includeExtendedInfo=true&includeTypeAttributes=true&page=2&qlQuery=Name+LIKE+Test&resultPerPage=25",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},

		{
			name: "when the workspace id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkspaceID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newAQLService := NewAQLService(testCase.fields.c)

			gotResult, gotResponse, err := newAQLService.Filter(testCase.args.ctx, testCase.args.workspaceID, testCase.args.payload)

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
