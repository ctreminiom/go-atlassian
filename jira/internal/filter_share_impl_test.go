package internal

import (
	"bytes"
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFilterShareService_Scope(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx context.Context
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/filter/defaultShareScope",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ShareFilterScopeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/filter/defaultShareScope",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ShareFilterScopeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/filter/defaultShareScope",
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

			shareService, err := NewFilterShareService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := shareService.Scope(testCase.args.ctx)

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

func TestFilterShareService_SetScope(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx   context.Context
		scope string
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:   context.Background(),
				scope: "PRIVATE",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.ShareFilterScopeScheme{Scope: "PRIVATE"}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/filter/defaultShareScope",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:   context.Background(),
				scope: "PRIVATE",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.ShareFilterScopeScheme{Scope: "PRIVATE"}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/filter/defaultShareScope",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:   context.Background(),
				scope: "PRIVATE",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.ShareFilterScopeScheme{Scope: "PRIVATE"}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/filter/defaultShareScope",
					bytes.NewReader([]byte{})).
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

			shareService, err := NewFilterShareService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := shareService.SetScope(testCase.args.ctx, testCase.args.scope)

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

func TestFilterShareService_Gets(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx      context.Context
		filterId int
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:      context.Background(),
				filterId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/filter/10001/permission",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					[]*model.SharePermissionScheme(nil)).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				filterId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/filter/10001/permission",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					[]*model.SharePermissionScheme(nil)).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the fieldId is not provied",
			fields: fields{version: "2"},
			args: args{
				ctx:      context.Background(),
				filterId: 0,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoFilterIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:      context.Background(),
				filterId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/filter/10001/permission",
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

			shareService, err := NewFilterShareService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := shareService.Gets(testCase.args.ctx, testCase.args.filterId)

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

func TestFilterShareService_Add(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx      context.Context
		filterId int
		payload  *model.PermissionFilterPayloadScheme
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:      context.Background(),
				filterId: 10001,
				payload: &model.PermissionFilterPayloadScheme{
					Type:      "group",
					GroupName: "jira-administrators",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.PermissionFilterPayloadScheme{Type: "group", GroupName: "jira-administrators"}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/filter/10001/permission",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					[]*model.SharePermissionScheme(nil)).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.Background(),
				filterId: 10001,
				payload: &model.PermissionFilterPayloadScheme{
					Type:      "group",
					GroupName: "jira-administrators",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.PermissionFilterPayloadScheme{Type: "group", GroupName: "jira-administrators"}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/filter/10001/permission",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					[]*model.SharePermissionScheme(nil)).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the fieldId is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:      context.Background(),
				filterId: 0,
				payload: &model.PermissionFilterPayloadScheme{
					Type:      "group",
					GroupName: "jira-administrators",
				},
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoFilterIDError,
		},

		{
			name:   "when the payload id not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:      context.Background(),
				filterId: 10001,
				payload:  nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.PermissionFilterPayloadScheme)(nil)).
					Return(bytes.NewReader([]byte{}), model.ErrNilPayloadError)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrNilPayloadError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:      context.Background(),
				filterId: 10001,
				payload: &model.PermissionFilterPayloadScheme{
					Type:      "group",
					GroupName: "jira-administrators",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.PermissionFilterPayloadScheme{Type: "group", GroupName: "jira-administrators"}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/filter/10001/permission",
					bytes.NewReader([]byte{})).
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

			shareService, err := NewFilterShareService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := shareService.Add(testCase.args.ctx, testCase.args.filterId, testCase.args.payload)

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

func TestFilterShareService_Get(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                    context.Context
		filterId, permissionId int
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				filterId:     10001,
				permissionId: 20,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/filter/10001/permission/20",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SharePermissionScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				filterId:     10001,
				permissionId: 20,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/filter/10001/permission/20",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SharePermissionScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the filterId is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				filterId:     0,
				permissionId: 20,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoFilterIDError,
		},

		{
			name:   "when the permissionId is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				filterId:     10001,
				permissionId: 0,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoPermissionGrantIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				filterId:     10001,
				permissionId: 20,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/filter/10001/permission/20",
					nil).
					Return(&http.Request{}, errors.New("error, unable to creat the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to creat the http request"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			shareService, err := NewFilterShareService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := shareService.Get(testCase.args.ctx, testCase.args.filterId, testCase.args.permissionId)

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

func TestFilterShareService_Delete(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                    context.Context
		filterId, permissionId int
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				filterId:     10001,
				permissionId: 20,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/filter/10001/permission/20",
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				filterId:     10001,
				permissionId: 20,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/filter/10001/permission/20",
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
			name:   "when the filterId is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				filterId:     0,
				permissionId: 20,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoFilterIDError,
		},

		{
			name:   "when the permissionId is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				filterId:     10001,
				permissionId: 0,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			wantErr: true,
			Err:     model.ErrNoPermissionGrantIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				filterId:     10001,
				permissionId: 20,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/filter/10001/permission/20",
					nil).
					Return(&http.Request{}, errors.New("error, unable to creat the http request"))

				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to creat the http request"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			shareService, err := NewFilterShareService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := shareService.Delete(testCase.args.ctx, testCase.args.filterId, testCase.args.permissionId)

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
