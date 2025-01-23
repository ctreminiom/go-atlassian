package internal

import (
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

func Test_internalObjectSchemaImpl_List(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		workspaceID string
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
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/list",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectSchemaPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/list",
					"",
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

			newObjectSchemaService := NewObjectSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectSchemaService.List(testCase.args.ctx, testCase.args.workspaceID)

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

func Test_internalObjectSchemaImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                         context.Context
		workspaceID, objectSchemaID string
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
				workspaceID:    "workspace-uuid-sample",
				objectSchemaID: "object-schema-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectSchemaScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				workspaceID:    "workspace-uuid-sample",
				objectSchemaID: "object-schema-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample",
					"",
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkspaceID,
		},

		{
			name: "when the object schema id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectSchemaID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectSchemaService := NewObjectSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectSchemaService.Get(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectSchemaID)

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

func Test_internalObjectSchemaImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                         context.Context
		workspaceID, objectSchemaID string
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
				workspaceID:    "workspace-uuid-sample",
				objectSchemaID: "object-schema-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectSchemaScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				workspaceID:    "workspace-uuid-sample",
				objectSchemaID: "object-schema-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample",
					"",
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkspaceID,
		},

		{
			name: "when the object schema id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectSchemaID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectSchemaService := NewObjectSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectSchemaService.Delete(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectSchemaID)

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

func Test_internalObjectSchemaImpl_Attributes(t *testing.T) {

	optionsMocked := &model.ObjectSchemaAttributesParamsScheme{
		OnlyValueEditable: true,
		Extended:          true,
		Query:             "query sample",
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                         context.Context
		workspaceID, objectSchemaID string
		options                     *model.ObjectSchemaAttributesParamsScheme
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
				workspaceID:    "workspace-uuid-sample",
				objectSchemaID: "object-schema-id-sample",
				options:        optionsMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample/attributes?extended=true&onlyValueEditable=true&onlyValueEditable=true&query=query+sample",
					"",
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
			name: "when the options are not provided",
			args: args{
				ctx:            context.Background(),
				workspaceID:    "workspace-uuid-sample",
				objectSchemaID: "object-schema-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample/attributes",
					"",
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
				ctx:            context.Background(),
				workspaceID:    "workspace-uuid-sample",
				objectSchemaID: "object-schema-id-sample",
				options:        optionsMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample/attributes?extended=true&onlyValueEditable=true&onlyValueEditable=true&query=query+sample",
					"",
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkspaceID,
		},

		{
			name: "when the object schema id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectSchemaID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectSchemaService := NewObjectSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectSchemaService.Attributes(
				testCase.args.ctx,
				testCase.args.workspaceID,
				testCase.args.objectSchemaID,
				testCase.args.options)

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

func Test_internalObjectSchemaImpl_ObjectTypes(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                         context.Context
		workspaceID, objectSchemaID string
		excludeAbstract             bool
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
				ctx:             context.Background(),
				workspaceID:     "workspace-uuid-sample",
				objectSchemaID:  "object-schema-id-sample",
				excludeAbstract: true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample/objecttypes?excludeAbstract=true",
					"",
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
				ctx:             context.Background(),
				workspaceID:     "workspace-uuid-sample",
				objectSchemaID:  "object-schema-id-sample",
				excludeAbstract: true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample/objecttypes?excludeAbstract=true",
					"",
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
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkspaceID,
		},

		{
			name: "when the object schema id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectSchemaID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectSchemaService := NewObjectSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectSchemaService.ObjectTypes(
				testCase.args.ctx,
				testCase.args.workspaceID,
				testCase.args.objectSchemaID,
				testCase.args.excludeAbstract)

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

func Test_internalObjectSchemaImpl_Update(t *testing.T) {

	payloadMocked := &model.ObjectSchemaPayloadScheme{
		Name:            "Computers",
		ObjectSchemaKey: "COMP",
		Description:     "The IT department schema",
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                         context.Context
		workspaceID, objectSchemaID string
		payload                     *model.ObjectSchemaPayloadScheme
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
				workspaceID:    "workspace-uuid-sample",
				objectSchemaID: "object-schema-id-sample",
				payload:        payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectSchemaScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:            context.Background(),
				workspaceID:    "workspace-uuid-sample",
				objectSchemaID: "object-schema-id-sample",
				payload:        payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/object-schema-id-sample",
					"",
					payloadMocked).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the workspace id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkspaceID,
		},

		{
			name: "when the object schema id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectSchemaID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectSchemaService := NewObjectSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectSchemaService.Update(
				testCase.args.ctx,
				testCase.args.workspaceID,
				testCase.args.objectSchemaID,
				testCase.args.payload,
			)

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

func Test_internalObjectSchemaImpl_Create(t *testing.T) {

	payloadMocked := &model.ObjectSchemaPayloadScheme{
		Name:            "Computers",
		ObjectSchemaKey: "COMP",
		Description:     "The IT department schema",
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		workspaceID string
		payload     *model.ObjectSchemaPayloadScheme
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
					http.MethodPost,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/create",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectSchemaScheme{}).
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
					http.MethodPost,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objectschema/create",
					"",
					payloadMocked).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
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

			newObjectSchemaService := NewObjectSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectSchemaService.Create(
				testCase.args.ctx,
				testCase.args.workspaceID,
				testCase.args.payload,
			)

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
