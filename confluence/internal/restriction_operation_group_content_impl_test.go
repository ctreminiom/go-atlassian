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

func Test_internalRestrictionOperationGroupImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                                    context.Context
		contentID, operationKey, groupNameOrID string
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
				contentID:     "100001",
				operationKey:  "read",
				groupNameOrID: "confluence-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100001/restriction/byOperation/read/group/confluence-users",
					"", nil).
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
				ctx:           context.Background(),
				contentID:     "100001",
				operationKey:  "read",
				groupNameOrID: "confluence-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100001/restriction/byOperation/read/group/confluence-users",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentID,
		},

		{
			name: "when the property key is not provided",
			args: args{
				ctx:       context.Background(),
				contentID: "1111",
			},
			wantErr: true,
			Err:     model.ErrNoContentRestrictionKey,
		},

		{
			name: "when the group name or id is not provided",
			args: args{
				ctx:          context.Background(),
				contentID:    "1111",
				operationKey: "read",
			},
			wantErr: true,
			Err:     model.ErrNoConfluenceGroup,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewRestrictionOperationGroupService(testCase.fields.c)

			gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.contentID, testCase.args.operationKey,
				testCase.args.groupNameOrID)

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

func Test_internalRestrictionOperationGroupImpl_Add(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                                    context.Context
		contentID, operationKey, groupNameOrID string
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
				contentID:     "100001",
				operationKey:  "read",
				groupNameOrID: "confluence-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/100001/restriction/byOperation/read/group/confluence-users",
					"", nil).
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
				ctx:           context.Background(),
				contentID:     "100001",
				operationKey:  "read",
				groupNameOrID: "confluence-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/100001/restriction/byOperation/read/group/confluence-users",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentID,
		},

		{
			name: "when the property key is not provided",
			args: args{
				ctx:       context.Background(),
				contentID: "1111",
			},
			wantErr: true,
			Err:     model.ErrNoContentRestrictionKey,
		},

		{
			name: "when the group name or id is not provided",
			args: args{
				ctx:          context.Background(),
				contentID:    "1111",
				operationKey: "read",
			},
			wantErr: true,
			Err:     model.ErrNoConfluenceGroup,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewRestrictionOperationGroupService(testCase.fields.c)

			gotResponse, err := newService.Add(testCase.args.ctx, testCase.args.contentID, testCase.args.operationKey,
				testCase.args.groupNameOrID)

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

func Test_internalRestrictionOperationGroupImpl_Remove(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                                    context.Context
		contentID, operationKey, groupNameOrID string
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
				contentID:     "100001",
				operationKey:  "read",
				groupNameOrID: "confluence-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/content/100001/restriction/byOperation/read/group/confluence-users",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the group provided is an uuid type",
			args: args{
				ctx:           context.Background(),
				contentID:     "100001",
				operationKey:  "read",
				groupNameOrID: "5185574c-4008-49bf-803c-e71baecf37d3",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/content/100001/restriction/byOperation/read/byGroupId/5185574c-4008-49bf-803c-e71baecf37d3",
					"", nil).
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
				ctx:           context.Background(),
				contentID:     "100001",
				operationKey:  "read",
				groupNameOrID: "confluence-users",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"wiki/rest/api/content/100001/restriction/byOperation/read/group/confluence-users",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentID,
		},

		{
			name: "when the property key is not provided",
			args: args{
				ctx:       context.Background(),
				contentID: "1111",
			},
			wantErr: true,
			Err:     model.ErrNoContentRestrictionKey,
		},

		{
			name: "when the group name or id is not provided",
			args: args{
				ctx:          context.Background(),
				contentID:    "1111",
				operationKey: "read",
			},
			wantErr: true,
			Err:     model.ErrNoConfluenceGroup,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewRestrictionOperationGroupService(testCase.fields.c)

			gotResponse, err := newService.Remove(testCase.args.ctx, testCase.args.contentID, testCase.args.operationKey,
				testCase.args.groupNameOrID)

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
