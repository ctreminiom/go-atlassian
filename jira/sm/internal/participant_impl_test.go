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

func Test_internalServiceRequestParticipantImpl_Gets(t *testing.T) {

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
				issueKeyOrID: "DESK-1",
				start:        100,
				limit:        50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DESK-1/participant?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestParticipantPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				start:        100,
				limit:        50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DESK-1/participant?limit=50&start=100",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestParticipantPageScheme{}).
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
				issueKeyOrID: "DESK-1",
				start:        100,
				limit:        50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/request/DESK-1/participant?limit=50&start=100",
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
			Err:     model.ErrNoIssueKeyOrIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			participantService := NewParticipantService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := participantService.Gets(testCase.args.ctx, testCase.args.issueKeyOrID,
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

func Test_internalServiceRequestParticipantImpl_Add(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
		accountIDs   []string
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
				issueKeyOrID: "DESK-1",
				accountIDs:   []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request/DESK-1/participant",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestParticipantPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				accountIDs:   []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request/DESK-1/participant",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestParticipantPageScheme{}).
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
				issueKeyOrID: "DESK-1",
				accountIDs:   []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)
				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/servicedeskapi/request/DESK-1/participant",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
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
			Err:     model.ErrNoIssueKeyOrIDError,
			wantErr: true,
		},

		{
			name: "when the account ids are not provided",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
			},
			Err:     model.ErrNoAccountSliceError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			participantService := NewParticipantService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := participantService.Add(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.accountIDs)

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

func Test_internalServiceRequestParticipantImpl_Remove(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx          context.Context
		issueKeyOrID string
		accountIDs   []string
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
				issueKeyOrID: "DESK-1",
				accountIDs:   []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/request/DESK-1/participant",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestParticipantPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
				accountIDs:   []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/request/DESK-1/participant",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.RequestParticipantPageScheme{}).
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
				issueKeyOrID: "DESK-1",
				accountIDs:   []string{"uuid-sample-1", "uuid-sample-2"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/servicedeskapi/request/DESK-1/participant",
					"",
					map[string]interface{}{"accountIds": []string{"uuid-sample-1", "uuid-sample-2"}}).
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
			Err:     model.ErrNoIssueKeyOrIDError,
			wantErr: true,
		},

		{
			name: "when the account ids are not provided",
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DESK-1",
			},
			Err:     model.ErrNoAccountSliceError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			participantService := NewParticipantService(testCase.fields.c, "latest")

			gotResult, gotResponse, err := participantService.Remove(testCase.args.ctx, testCase.args.issueKeyOrID, testCase.args.accountIDs)

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
