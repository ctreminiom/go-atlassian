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

func Test_internalObjectTypeAttributeImpl_Create(t *testing.T) {

	payloadMocked := &model.ObjectTypeAttributeScheme{
		WorkspaceId: "g2778e1d-939d-581d-c8e2-9d5g59de456b",
		GlobalId:    "g2778e1d-939d-581d-c8e2-9d5g59de456b:1330",
		ID:          "1330",
		ObjectType:  nil,
		Name:        "Geolocation",
		Label:       false,
		Type:        0,
		Description: "",
		DefaultType: &model.ObjectTypeAssetAttributeDefaultTypeScheme{
			ID:   0,
			Name: "Text",
		},
		TypeValue:               "",
		TypeValueMulti:          nil,
		AdditionalValue:         "",
		ReferenceType:           nil,
		ReferenceObjectTypeId:   "",
		ReferenceObjectType:     nil,
		Editable:                false,
		System:                  false,
		Indexed:                 false,
		Sortable:                false,
		Summable:                false,
		MinimumCardinality:      0,
		MaximumCardinality:      0,
		Suffix:                  "",
		Removable:               false,
		ObjectAttributeExists:   false,
		Hidden:                  false,
		IncludeChildObjectTypes: false,
		UniqueAttribute:         false,
		RegexValidation:         "",
		Iql:                     "",
		QlQuery:                 "",
		Options:                 "",
		Position:                6,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		workspaceID, objectTypeID string
		payload                   *model.ObjectTypeAttributeScheme
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
				ctx:          context.TODO(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-uuid-sample",
				payload:      payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttypeattribute/object-type-uuid-sample",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectTypeAttributeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.TODO(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-uuid-sample",
				payload:      payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttypeattribute/object-type-uuid-sample",
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
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkspaceIDError,
		},

		{
			name: "when the object type id id is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectTypeAttributeService := NewObjectTypeAttributeService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectTypeAttributeService.Create(
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

func Test_internalObjectTypeAttributeImpl_Update(t *testing.T) {

	payloadMocked := &model.ObjectTypeAttributeScheme{
		WorkspaceId: "g2778e1d-939d-581d-c8e2-9d5g59de456b",
		GlobalId:    "g2778e1d-939d-581d-c8e2-9d5g59de456b:1330",
		ID:          "1330",
		ObjectType:  nil,
		Name:        "Geolocation",
		Label:       false,
		Type:        0,
		Description: "",
		DefaultType: &model.ObjectTypeAssetAttributeDefaultTypeScheme{
			ID:   0,
			Name: "Text",
		},
		TypeValue:               "",
		TypeValueMulti:          nil,
		AdditionalValue:         "",
		ReferenceType:           nil,
		ReferenceObjectTypeId:   "",
		ReferenceObjectType:     nil,
		Editable:                false,
		System:                  false,
		Indexed:                 false,
		Sortable:                false,
		Summable:                false,
		MinimumCardinality:      0,
		MaximumCardinality:      0,
		Suffix:                  "",
		Removable:               false,
		ObjectAttributeExists:   false,
		Hidden:                  false,
		IncludeChildObjectTypes: false,
		UniqueAttribute:         false,
		RegexValidation:         "",
		Iql:                     "",
		QlQuery:                 "",
		Options:                 "",
		Position:                6,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                                    context.Context
		workspaceID, objectTypeID, attributeID string
		payload                                *model.ObjectTypeAttributeScheme
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
				ctx:          context.TODO(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-uuid-sample",
				attributeID:  "attribute-id-sample",
				payload:      payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttypeattribute/object-type-uuid-sample/attribute-id-sample",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ObjectTypeAttributeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.TODO(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-uuid-sample",
				attributeID:  "attribute-id-sample",
				payload:      payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttypeattribute/object-type-uuid-sample/attribute-id-sample",
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
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkspaceIDError,
		},

		{
			name: "when the object type id id is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeIDError,
		},

		{
			name: "when the attribute id id is not provided",
			args: args{
				ctx:          context.TODO(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeAttributeIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectTypeAttributeService := NewObjectTypeAttributeService(testCase.fields.c)

			gotResult, gotResponse, err := newObjectTypeAttributeService.Update(
				testCase.args.ctx,
				testCase.args.workspaceID,
				testCase.args.objectTypeID,
				testCase.args.attributeID,
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

func Test_internalObjectTypeAttributeImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                      context.Context
		workspaceID, attributeID string
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
				attributeID: "attribute-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttypeattribute/attribute-id-sample",
					"",
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
				attributeID: "attribute-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"jsm/assets/workspace/workspace-uuid-sample/v1/objecttypeattribute/attribute-id-sample",
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
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoWorkspaceIDError,
		},

		{
			name: "when the attribute id id is not provided",
			args: args{
				ctx:         context.TODO(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeAttributeIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newObjectTypeAttributeService := NewObjectTypeAttributeService(testCase.fields.c)

			gotResponse, err := newObjectTypeAttributeService.Delete(
				testCase.args.ctx,
				testCase.args.workspaceID,
				testCase.args.attributeID,
			)

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
