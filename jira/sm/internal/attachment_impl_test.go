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

func Test_internalServiceRequestAttachmentImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
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
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-2",
				start:        100,
				limit:        50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DUMMY-2/attachment?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestAttachmentPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-2",
				start:        100,
				limit:        50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DUMMY-2/attachment?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestAttachmentPageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-2",
				start:        100,
				limit:        50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DUMMY-2/attachment?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the issue key or id is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoIssueKeyOrIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService := NewAttachmentService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := attachmentService.Gets(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.start,
				testCase.args.limit)

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

func Test_internalServiceRequestAttachmentImpl_Create(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                    context.Context
		issueKeyOrID           string
		temporaryAttachmentIDs []string
		public                 bool
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
				ctx:                    context.Background(),
				issueKeyOrID:           "DUMMY-2",
				temporaryAttachmentIDs: []string{"10001"},
				public:                 true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request/DUMMY-2/attachment",
					"",
					map[string]interface{}{"public": true, "temporaryAttachmentIds": []string{"10001"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestAttachmentCreationScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:                    context.Background(),
				issueKeyOrID:           "DUMMY-2",
				temporaryAttachmentIDs: []string{"10001"},
				public:                 true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request/DUMMY-2/attachment",
					"",
					map[string]interface{}{"public": true, "temporaryAttachmentIds": []string{"10001"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestAttachmentCreationScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:                    context.Background(),
				issueKeyOrID:           "DUMMY-2",
				temporaryAttachmentIDs: []string{"10001"},
				public:                 true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request/DUMMY-2/attachment",
					"",
					map[string]interface{}{"public": true, "temporaryAttachmentIds": []string{"10001"}}).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the issue key or id is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoIssueKeyOrIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService := NewAttachmentService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := attachmentService.Create(testCase.args.ctx, testCase.args.issueKeyOrID,
				testCase.args.temporaryAttachmentIDs, testCase.args.public)

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
