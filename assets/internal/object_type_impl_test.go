package internal

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
)

func Test_internalObjectTypeImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		workspaceID, objectTypeID string
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
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/object-type-id-sample",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectTypeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.Background(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/object-type-id-sample",
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
			name: "when the object type id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectTypeService := NewObjectTypeService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectTypeService.Get(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectTypeID)

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

func Test_internalObjectTypeImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		workspaceID, objectTypeID string
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
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/object-type-id-sample",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectTypeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.Background(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/object-type-id-sample",
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
			name: "when the object type id is not provided",
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

			newObjectTypeService := NewObjectTypeService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectTypeService.Delete(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectTypeID)

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

func Test_internalObjectTypeImpl_Attributes(t *testing.T) {

	optionsMocked := &model.ObjectTypeAttributesParamsScheme{
		OnlyValueEditable:       true,
		OrderByName:             true,
		Query:                   "aql-query-sample",
		IncludeValuesExist:      true,
		ExcludeParentAttributes: true,
		IncludeChildren:         true,
		OrderByRequired:         true,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		workspaceID, objectTypeID string
		options                   *model.ObjectTypeAttributesParamsScheme
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
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
				options:      optionsMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/object-type-id-sample/attributes?excludeParentAttributes=true&includeChildren=true&includeValuesExist=true&onlyValueEditable=true&orderByName=true&orderByRequired=true&query=aql-query-sample",
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
				ctx:          context.Background(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
				options:      optionsMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/object-type-id-sample/attributes?excludeParentAttributes=true&includeChildren=true&includeValuesExist=true&onlyValueEditable=true&orderByName=true&orderByRequired=true&query=aql-query-sample",
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
			name: "when the object type id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectTypeService := NewObjectTypeService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectTypeService.Attributes(testCase.args.ctx, testCase.args.workspaceID, testCase.args.objectTypeID,
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

func Test_internalObjectTypeImpl_Update(t *testing.T) {

	payloadMocked := &model.ObjectTypePayloadScheme{
		Name:               "Office",
		Description:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin nec ex.",
		IconID:             "13",
		ObjectSchemaID:     "",
		ParentObjectTypeID: "",
		Inherited:          true,
		AbstractObjectType: true,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		workspaceID, objectTypeID string
		payload                   *model.ObjectTypePayloadScheme
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
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
				payload:      payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/object-type-id-sample",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectTypeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.Background(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
				payload:      payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/object-type-id-sample",
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
			name: "when the object type id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectTypeService := NewObjectTypeService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectTypeService.Update(
				testCase.args.ctx,
				testCase.args.workspaceID,
				testCase.args.objectTypeID,
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

func Test_internalObjectTypeImpl_Create(t *testing.T) {

	payloadMocked := &model.ObjectTypePayloadScheme{
		Name:               "Office",
		Description:        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin nec ex.",
		IconID:             "13",
		ObjectSchemaID:     "",
		ParentObjectTypeID: "",
		Inherited:          true,
		AbstractObjectType: true,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		workspaceID string
		payload     *model.ObjectTypePayloadScheme
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
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/create",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectTypeScheme{}).
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
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/create",
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

			newObjectTypeService := NewObjectTypeService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectTypeService.Create(
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

func Test_internalObjectTypeImpl_Position(t *testing.T) {

	payloadMocked := &model.ObjectTypePositionPayloadScheme{
		ToObjectTypeID: "2",
		Position:       0,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		workspaceID, objectTypeID string
		payload                   *model.ObjectTypePositionPayloadScheme
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
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
				payload:      payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/object-type-id-sample/position",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectTypeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.Background(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
				payload:      payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttype/object-type-id-sample/position",
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
			name: "when the object type id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeID,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectTypeService := NewObjectTypeService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectTypeService.Position(
				testCase.args.ctx,
				testCase.args.workspaceID,
				testCase.args.objectTypeID,
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
