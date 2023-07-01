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

func Test_internalSCIMSchemaImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		directoryID string
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
				directoryID: "direction-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Schemas",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SCIMSchemasScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.TODO(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				directoryID: "direction-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Schemas",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newSCIMSchemaService := NewSCIMSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMSchemaService.Gets(testCase.args.ctx, testCase.args.directoryID)

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

func Test_internalSCIMSchemaImpl_Group(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		directoryID string
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
				directoryID: "direction-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SCIMSchemaScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.TODO(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				directoryID: "direction-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newSCIMSchemaService := NewSCIMSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMSchemaService.Group(testCase.args.ctx, testCase.args.directoryID)

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

func Test_internalSCIMSchemaImpl_User(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		directoryID string
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
				directoryID: "direction-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SCIMSchemaScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.TODO(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				directoryID: "direction-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newSCIMSchemaService := NewSCIMSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMSchemaService.User(testCase.args.ctx, testCase.args.directoryID)

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

func Test_internalSCIMSchemaImpl_Enterprise(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		directoryID string
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
				directoryID: "direction-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SCIMSchemaScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.TODO(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				directoryID: "direction-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newSCIMSchemaService := NewSCIMSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMSchemaService.Enterprise(testCase.args.ctx, testCase.args.directoryID)

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

func Test_internalSCIMSchemaImpl_Feature(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		directoryID string
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
				directoryID: "direction-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/ServiceProviderConfig",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ServiceProviderConfigScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.TODO(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryIDError,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.TODO(),
				directoryID: "direction-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/ServiceProviderConfig",
					"",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newSCIMSchemaService := NewSCIMSchemaService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMSchemaService.Feature(testCase.args.ctx, testCase.args.directoryID)

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
