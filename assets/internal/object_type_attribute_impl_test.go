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
	var attributeType int
	var defaultTypeID int
	var minimumCardinality int
	var maximumCardinality int
	attributeType = 0
	defaultTypeID = 0
	minimumCardinality = 0
	maximumCardinality = 0
	payloadMocked := &model.ObjectTypeAttributePayloadScheme{
		Name:        		"Geolocation",
		Label:  	 	 false,
		Type:       		 &attributeType,
		Description: 		 "",
		DefaultTypeId: 		 &defaultTypeID,
		TypeValue:               "",
		TypeValueMulti:          nil,
		AdditionalValue:         "",
		Summable:                false,
		MinimumCardinality:      &minimumCardinality,
		MaximumCardinality:      &maximumCardinality,
		Suffix:                  "",
		Hidden:                  false,
		IncludeChildObjectTypes: false,
		UniqueAttribute:         false,
		RegexValidation:         "",
		Iql:                     "",
		QlQuery:                 "",
		Options:                 "",
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                       context.Context
		workspaceID, objectTypeID string
		payload                   *model.ObjectTypeAttributePayloadScheme
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

	var attributeType int
	var defaultTypeID int
	var minimumCardinality int
	var maximumCardinality int
	attributeType = 0
	defaultTypeID = 0
	minimumCardinality = 0
	maximumCardinality = 0
	payloadMocked := &model.ObjectTypeAttributePayloadScheme{
		Name:        		"Geolocation",
		Label:  	 	 false,
		Type:       		 &attributeType,
		Description: 		 "",
		DefaultTypeId: 		 &defaultTypeID,
		TypeValue:               "",
		TypeValueMulti:          nil,
		AdditionalValue:         "",
		Summable:                false,
		MinimumCardinality:      &minimumCardinality,
		MaximumCardinality:      &maximumCardinality,
		Suffix:                  "",
		Hidden:                  false,
		IncludeChildObjectTypes: false,
		UniqueAttribute:         false,
		RegexValidation:         "",
		Iql:                     "",
		QlQuery:                 "",
		Options:                 "",
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
