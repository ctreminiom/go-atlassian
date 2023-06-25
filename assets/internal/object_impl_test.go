package internal

import (
	"bytes"
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func Test_internalObjectImpl_Get(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                   context.Context
		workspaceID, objectID string
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
				objectID:    "1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/1",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
				objectID:    "1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/1",
					nil).
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

		{
			name: "when the object id is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewObjectService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectID)

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

func Test_internalObjectImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                   context.Context
		workspaceID, objectID string
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
				objectID:    "1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/1",
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
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
				objectID:    "1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/1",
					nil).
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

		{
			name: "when the object id is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewObjectService(testCase.fields.c)

			gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectID)

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

func Test_internalObjectImpl_Attributes(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                   context.Context
		workspaceID, objectID string
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
				objectID:    "1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/1/attributes",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
				objectID:    "1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/1/attributes",
					nil).
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

		{
			name: "when the object id is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewObjectService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Attributes(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectID)

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

func Test_internalObjectImpl_History(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                   context.Context
		workspaceID, objectID string
		ascOrder              bool
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
				objectID:    "1",
				ascOrder:    true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/1/history?asc=true",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
				objectID:    "1",
				ascOrder:    true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/1/history?asc=true",
					nil).
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

		{
			name: "when the object id is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewObjectService(testCase.fields.c)

			gotResult, gotResponse, err := newService.History(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectID,
				testCase.args.ascOrder)

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

func Test_internalObjectImpl_References(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                   context.Context
		workspaceID, objectID string
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
				objectID:    "1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/1/referenceinfo",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
				objectID:    "1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/1/referenceinfo",
					nil).
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

		{
			name: "when the object id is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewObjectService(testCase.fields.c)

			gotResult, gotResponse, err := newService.References(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectID)

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

func Test_internalObjectImpl_Relation(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                   context.Context
		workspaceID, objectID string
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
				objectID:    "1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectconnectedtickets/1/tickets",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.TicketPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
				objectID:    "1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectconnectedtickets/1/tickets",
					nil).
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

		{
			name: "when the object id is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewObjectService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Relation(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectID)

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

func Test_internalObjectImpl_Update(t *testing.T) {

	payloadMocked := &model.ObjectPayloadScheme{
		ObjectTypeID: "23",
		HasAvatar:    false,
		Attributes: []*model.ObjectPayloadAttributeScheme{
			{
				ObjectTypeAttributeID: "265",
				ObjectAttributeValues: []*model.ObjectPayloadAttributeValueScheme{
					{
						Value: "A placeholder value",
					},
				},
			},
		},
	}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx         context.Context
		workspaceID string
		objectID    string
		payload     *model.ObjectPayloadScheme
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
				objectID:    "object-uuid-sample",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/object-uuid-sample",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
				objectID:    "object-uuid-sample",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/object-uuid-sample",
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

		{
			name: "when the object id is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewObjectService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Update(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectID,
				testCase.args.payload)

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

func Test_internalObjectImpl_Create(t *testing.T) {

	payloadMocked := &model.ObjectPayloadScheme{
		ObjectTypeID: "23",
		HasAvatar:    false,
		Attributes: []*model.ObjectPayloadAttributeScheme{
			{
				ObjectTypeAttributeID: "135",
				ObjectAttributeValues: []*model.ObjectPayloadAttributeValueScheme{
					{
						Value: "NY-1",
					},
				},
			},
			{
				ObjectTypeAttributeID: "144",
				ObjectAttributeValues: []*model.ObjectPayloadAttributeValueScheme{
					{
						Value: "99",
					},
				},
			},
		},
	}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx         context.Context
		workspaceID string
		payload     *model.ObjectPayloadScheme
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
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/create",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectScheme{}).
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
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/create",
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

			newService := NewObjectService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Create(testCase.args.ctx, testCase.args.workspaceID, testCase.args.payload)

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

func Test_internalObjectImpl_Filter(t *testing.T) {

	payloadMocked := &struct {
		Aql string "json:\"qlQuery\""
	}{Aql: "objectType = Office AND Name LIKE SYD"}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                 context.Context
		workspaceID, aql    string
		attributes          bool
		startAt, maxResults int
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
				aql:         "objectType = Office AND Name LIKE SYD",
				attributes:  false,
				startAt:     50,
				maxResults:  50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/aql?includeAttributes=false&maxResults=50&startAt=50",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectListResultScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
				aql:         "objectType = Office AND Name LIKE SYD",
				attributes:  false,
				startAt:     50,
				maxResults:  50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/aql?includeAttributes=false&maxResults=50&startAt=50",
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

		{
			name: "when the aql query is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAqlQueryError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewObjectService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Filter(testCase.args.ctx, testCase.args.workspaceID, testCase.args.aql,
				testCase.args.attributes, testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalObjectImpl_Search(t *testing.T) {

	payloadMocked := &model.ObjectSearchParamsScheme{
		Query:             "objectType = Office AND Name LIKE SYD",
		ObjectTypeID:      "23",
		Page:              1,
		ResultPerPage:     25,
		ObjectSchemaID:    "6",
		IncludeAttributes: false,
		AttributesToDisplay: &model.AttributesToDisplayScheme{
			AttributesToDisplayIds: []int{135, 144},
		},
	}

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx         context.Context
		workspaceID string
		payload     *model.ObjectSearchParamsScheme
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
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/navlist/aql",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectPageScheme{}).
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
					"jsm/assets/workspace/workspace-uuid-sample/v1/object/navlist/aql",
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

			objectService := NewObjectService(testCase.fields.c)

			gotResult, gotResponse, err := objectService.Search(testCase.args.ctx, testCase.args.workspaceID, testCase.args.payload)

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
