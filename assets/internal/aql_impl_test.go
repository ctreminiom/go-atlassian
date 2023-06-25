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

func Test_internalAQLImpl_Filter(t *testing.T) {

	payloadMocked := &model.AQLSearchParamsScheme{
		Query:                 "objectType = Office AND Name LIKE SYD",
		Page:                  2,
		ResultPerPage:         25,
		IncludeAttributes:     true,
		IncludeAttributesDeep: true,
		IncludeTypeAttributes: true,
		IncludeExtendedInfo:   true,
	}

	type fields struct {
		c service.Client
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
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"jsm/assets/workspace/workspace-uuid-sample/v1/aql/objects",
					bytes.NewReader([]byte{})).
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
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"jsm/assets/workspace/workspace-uuid-sample/v1/aql/objects",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the workspace id is not provided",
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkspaceIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewAQLService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Filter(testCase.args.ctx, testCase.args.workspaceID, testCase.args.payload)

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
