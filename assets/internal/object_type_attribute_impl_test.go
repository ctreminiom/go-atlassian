package internal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
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
		Name:                    "Geolocation",
		Label:                   false,
		Type:                    &attributeType,
		Description:             "",
		DefaultTypeID:           &defaultTypeID,
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
				ctx:          context.Background(),
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
				ctx:          context.Background(),
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

		{
			name: "when the object type id id is not provided",
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
		Name:                    "Geolocation",
		Label:                   false,
		Type:                    &attributeType,
		Description:             "",
		DefaultTypeID:           &defaultTypeID,
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
		payload                                *model.ObjectTypeAttributePayloadScheme
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
				ctx:          context.Background(),
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

		{
			name: "when the object type id id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeID,
		},

		{
			name: "when the attribute id id is not provided",
			args: args{
				ctx:          context.Background(),
				workspaceID:  "workspace-uuid-sample",
				objectTypeID: "object-type-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeAttributeID,
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
				ctx:         context.Background(),
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
				ctx:         context.Background(),
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

		{
			name: "when the attribute id id is not provided",
			args: args{
				ctx:         context.Background(),
				workspaceID: "workspace-uuid-sample",
			},
			wantErr: true,
			Err:     model.ErrNoObjectTypeAttributeID,
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
			}
		})
	}
}
