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

func Test_internalSCIMUserImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx               context.Context
		directoryID       string
		opts              *model.SCIMUserGetsOptionsScheme
		startIndex, count int
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
				directoryID: "direction-id-sample",
				opts: &model.SCIMUserGetsOptionsScheme{
					Attributes:         []string{"attributes"},
					ExcludedAttributes: []string{"attributes"},
					Filter:             "users",
				},
				startIndex: 0,
				count:      50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Users?attributes=attributes&count=50&excludedAttributes=attributes&filter=users&startIndex=0",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SCIMUserPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
				opts: &model.SCIMUserGetsOptionsScheme{
					Attributes:         []string{"attributes"},
					ExcludedAttributes: []string{"attributes"},
					Filter:             "users",
				},
				startIndex: 0,
				count:      50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Users?attributes=attributes&count=50&excludedAttributes=attributes&filter=users&startIndex=0",
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

			newSCIMUserService := NewSCIMUserService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMUserService.Gets(testCase.args.ctx, testCase.args.directoryID, testCase.args.opts, testCase.args.startIndex,
				testCase.args.count)

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

func Test_internalSCIMUserImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                            context.Context
		directoryID, userID            string
		attributes, excludedAttributes []string
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
				ctx:                context.Background(),
				directoryID:        "direction-id-sample",
				userID:             "user-id-uuid-sample",
				attributes:         []string{"groups"},
				excludedAttributes: []string{"roles"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Users/user-id-uuid-sample?attributes=groups&excludedAttributes=roles",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SCIMUserScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryID,
		},

		{
			name: "when the user id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminUserID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:                context.Background(),
				directoryID:        "direction-id-sample",
				userID:             "user-id-uuid-sample",
				attributes:         []string{"groups"},
				excludedAttributes: []string{"roles"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Users/user-id-uuid-sample?attributes=groups&excludedAttributes=roles",
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

			newSCIMUserService := NewSCIMUserService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMUserService.Get(testCase.args.ctx, testCase.args.directoryID, testCase.args.userID, testCase.args.attributes,
				testCase.args.excludedAttributes)

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

func Test_internalSCIMUserImpl_Deactivate(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                 context.Context
		directoryID, userID string
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
				directoryID: "direction-id-sample",
				userID:      "user-id-uuid-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"scim/directory/direction-id-sample/Users/user-id-uuid-sample",
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
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryID,
		},

		{
			name: "when the user id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminUserID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
				userID:      "user-id-uuid-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"scim/directory/direction-id-sample/Users/user-id-uuid-sample",
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

			newSCIMUserService := NewSCIMUserService(testCase.fields.c)

			gotResponse, err := newSCIMUserService.Deactivate(testCase.args.ctx, testCase.args.directoryID, testCase.args.userID)

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

func Test_internalSCIMUserImpl_Path(t *testing.T) {

	payloadMocked := &model.SCIMUserToPathScheme{
		Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                            context.Context
		directoryID, userID            string
		payload                        *model.SCIMUserToPathScheme
		attributes, excludedAttributes []string
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
				ctx:                context.Background(),
				directoryID:        "direction-id-sample",
				userID:             "user-id-uuid-sample",
				payload:            payloadMocked,
				attributes:         []string{"groups"},
				excludedAttributes: []string{"roles"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPatch,
					"scim/directory/direction-id-sample/Users/user-id-uuid-sample?attributes=groups&excludedAttributes=roles",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SCIMUserScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryID,
		},

		{
			name: "when the user id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminUserID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:                context.Background(),
				directoryID:        "direction-id-sample",
				userID:             "user-id-uuid-sample",
				payload:            payloadMocked,
				attributes:         []string{"groups"},
				excludedAttributes: []string{"roles"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPatch,
					"scim/directory/direction-id-sample/Users/user-id-uuid-sample?attributes=groups&excludedAttributes=roles",
					"",
					payloadMocked).
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

			newSCIMUserService := NewSCIMUserService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMUserService.Path(testCase.args.ctx, testCase.args.directoryID, testCase.args.userID,
				testCase.args.payload, testCase.args.attributes, testCase.args.excludedAttributes)

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

func Test_internalSCIMUserImpl_Update(t *testing.T) {

	payloadMocked := &model.SCIMUserScheme{
		UserName:    "username-updated-with-overwrite-method",
		DisplayName: "AA",
		NickName:    "AA",
		Title:       "AA",
		Department:  "President",
		Emails: []*model.SCIMUserEmailScheme{
			{
				Value:   "carlos@go-atlassian.io",
				Type:    "work",
				Primary: true,
			},
		},
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                            context.Context
		directoryID, userID            string
		payload                        *model.SCIMUserScheme
		attributes, excludedAttributes []string
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
				ctx:                context.Background(),
				directoryID:        "direction-id-sample",
				userID:             "user-id-uuid-sample",
				payload:            payloadMocked,
				attributes:         []string{"groups"},
				excludedAttributes: []string{"roles"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"scim/directory/direction-id-sample/Users/user-id-uuid-sample?attributes=groups&excludedAttributes=roles",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SCIMUserScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryID,
		},

		{
			name: "when the user id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminUserID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:                context.Background(),
				directoryID:        "direction-id-sample",
				userID:             "user-id-uuid-sample",
				payload:            payloadMocked,
				attributes:         []string{"groups"},
				excludedAttributes: []string{"roles"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"scim/directory/direction-id-sample/Users/user-id-uuid-sample?attributes=groups&excludedAttributes=roles",
					"",
					payloadMocked).
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

			newSCIMUserService := NewSCIMUserService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMUserService.Update(testCase.args.ctx, testCase.args.directoryID, testCase.args.userID,
				testCase.args.payload, testCase.args.attributes, testCase.args.excludedAttributes)

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

func Test_internalSCIMUserImpl_Create(t *testing.T) {

	payloadMocked := &model.SCIMUserScheme{
		UserName:    "username-updated-with-overwrite-method",
		DisplayName: "AA",
		NickName:    "AA",
		Title:       "AA",
		Department:  "President",
		Emails: []*model.SCIMUserEmailScheme{
			{
				Value:   "carlos@go-atlassian.io",
				Type:    "work",
				Primary: true,
			},
		},
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                            context.Context
		directoryID                    string
		payload                        *model.SCIMUserScheme
		attributes, excludedAttributes []string
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
				ctx:                context.Background(),
				directoryID:        "direction-id-sample",
				payload:            payloadMocked,
				attributes:         []string{"groups"},
				excludedAttributes: []string{"roles"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"scim/directory/direction-id-sample/Users?attributes=groups&excludedAttributes=roles",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SCIMUserScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the directory id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "",
			},
			wantErr: true,
			Err:     model.ErrNoAdminDirectoryID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:                context.Background(),
				directoryID:        "direction-id-sample",
				payload:            payloadMocked,
				attributes:         []string{"groups"},
				excludedAttributes: []string{"roles"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"scim/directory/direction-id-sample/Users?attributes=groups&excludedAttributes=roles",
					"",
					payloadMocked).
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

			newSCIMUserService := NewSCIMUserService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMUserService.Create(testCase.args.ctx, testCase.args.directoryID,
				testCase.args.payload, testCase.args.attributes, testCase.args.excludedAttributes)

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
