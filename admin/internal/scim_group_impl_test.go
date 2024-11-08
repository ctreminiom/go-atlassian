package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalSCIMGroupImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                 context.Context
		directoryID, filter string
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
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
				filter:      "filter-sample",
				startAt:     50,
				maxResults:  50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Groups?count=50&filter=filter-sample&startIndex=50",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ScimGroupPageScheme{}).
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
				filter:      "filter-sample",
				startAt:     50,
				maxResults:  50,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Groups?count=50&filter=filter-sample&startIndex=50",
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

			newSCIMGroupService := NewSCIMGroupService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMGroupService.Gets(testCase.args.ctx, testCase.args.directoryID,
				testCase.args.filter, testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalSCIMGroupImpl_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                  context.Context
		directoryID, groupID string
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
				groupID:     "group-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Groups/group-id-sample",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ScimGroupScheme{}).
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
			name: "when the group id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "directory-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminGroupID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
				groupID:     "group-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"scim/directory/direction-id-sample/Groups/group-id-sample",
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

			newSCIMGroupService := NewSCIMGroupService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMGroupService.Get(testCase.args.ctx, testCase.args.directoryID,
				testCase.args.groupID)

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

func Test_internalSCIMGroupImpl_Delete(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                  context.Context
		directoryID, groupID string
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
				groupID:     "group-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"scim/directory/direction-id-sample/Groups/group-id-sample",
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
			name: "when the group id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "directory-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminGroupID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
				groupID:     "group-id-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"scim/directory/direction-id-sample/Groups/group-id-sample",
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

			newSCIMGroupService := NewSCIMGroupService(testCase.fields.c)

			gotResponse, err := newSCIMGroupService.Delete(testCase.args.ctx, testCase.args.directoryID,
				testCase.args.groupID)

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

func Test_internalSCIMGroupImpl_Create(t *testing.T) {

	payloadMocked := map[string]interface{}{"displayName": "group-name-sample"}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                    context.Context
		directoryID, groupName string
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
				groupName:   "group-name-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"scim/directory/direction-id-sample/Groups",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ScimGroupScheme{}).
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
			name: "when the group name is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "directory-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminGroupName,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
				groupName:   "group-name-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"scim/directory/direction-id-sample/Groups",
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

			newSCIMGroupService := NewSCIMGroupService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMGroupService.Create(testCase.args.ctx, testCase.args.directoryID,
				testCase.args.groupName)

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

func Test_internalSCIMGroupImpl_Update(t *testing.T) {

	payloadMocked := map[string]interface{}{"displayName": "group-name-sample"}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                                context.Context
		directoryID, groupID, newGroupName string
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
				directoryID:  "direction-id-sample",
				groupID:      "group-id-sample",
				newGroupName: "group-name-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"scim/directory/direction-id-sample/Groups/group-id-sample",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ScimGroupScheme{}).
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
			name: "when the group id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "directory-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminGroupID,
		},

		{
			name: "when the group name is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "directory-id-sample",
				groupID:     "group-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminGroupName,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.Background(),
				directoryID:  "direction-id-sample",
				groupID:      "group-id-sample",
				newGroupName: "group-name-sample",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"scim/directory/direction-id-sample/Groups/group-id-sample",
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

			newSCIMGroupService := NewSCIMGroupService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMGroupService.Update(testCase.args.ctx, testCase.args.directoryID,
				testCase.args.groupID, testCase.args.newGroupName)

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

func Test_internalSCIMGroupImpl_Path(t *testing.T) {

	payloadMocked := &model.SCIMGroupPathScheme{
		Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
		Operations: []*model.SCIMGroupOperationScheme{
			{
				Op:   "add",
				Path: "members",
				Value: []*model.SCIMGroupOperationValueScheme{
					{
						Value:   "account-id-sample",
						Display: "Example Display Name",
					},
				},
			},
		},
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                  context.Context
		directoryID, groupID string
		payload              *model.SCIMGroupPathScheme
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
				groupID:     "group-id-sample",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPatch,
					"scim/directory/direction-id-sample/Groups/group-id-sample",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ScimGroupScheme{}).
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
			name: "when the group id is not provided",
			args: args{
				ctx:         context.Background(),
				directoryID: "directory-id-sample",
			},
			wantErr: true,
			Err:     model.ErrNoAdminGroupID,
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:         context.Background(),
				directoryID: "direction-id-sample",
				groupID:     "group-id-sample",
				payload:     payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPatch,
					"scim/directory/direction-id-sample/Groups/group-id-sample",
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

			newSCIMGroupService := NewSCIMGroupService(testCase.fields.c)

			gotResult, gotResponse, err := newSCIMGroupService.Path(testCase.args.ctx, testCase.args.directoryID,
				testCase.args.groupID, testCase.args.payload)

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
