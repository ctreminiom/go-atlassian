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

func Test_internalProjectImpl_Create(t *testing.T) {

	payloadMocked := &model.ProjectPayloadScheme{
		NotificationScheme:  10021,
		Description:         "Example Project description",
		LeadAccountID:       "5b86be50b8e3cb5895860d6d",
		URL:                 "https://atlassian.com",
		ProjectTemplateKey:  "com.pyxis.greenhopper.jira:gh-simplified-agility-kanban",
		AvatarID:            10200,
		IssueSecurityScheme: 10001,
		Name:                "Project DUMMY #3",
		PermissionScheme:    10011,
		AssigneeType:        "PROJECT_LEAD",
		ProjectTypeKey:      "software",
		Key:                 "DUMMY3",
		CategoryID:          10120,
	}

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.ProjectPayloadScheme
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/project",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.NewProjectCreatedScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/project",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.NewProjectCreatedScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/project",
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

			newService, err := NewProjectService(testCase.fields.c, testCase.fields.version, &ProjectChildServices{})
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Create(testCase.args.ctx, testCase.args.payload)

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

func Test_internalProjectImpl_Search(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                 context.Context
		options             *model.ProjectSearchOptionsScheme
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
				options: &model.ProjectSearchOptionsScheme{
					OrderBy:        "category",
					Query:          "dummy project",
					Action:         "view",
					ProjectKeyType: "[thepropertykey].something.nested=1",
					CategoryID:     400,
					Expand:         []string{"description", "projectKeys"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/project/search?action=view&categoryId=400&expand=description%2CprojectKeys&maxResults=50&orderBy=category&query=dummy+project&startAt=0&typeKey=%5Bthepropertykey%5D.something.nested%3D1",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectSearchScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx: context.TODO(),
				options: &model.ProjectSearchOptionsScheme{
					OrderBy:        "category",
					Query:          "dummy project",
					Action:         "view",
					ProjectKeyType: "[thepropertykey].something.nested=1",
					CategoryID:     400,
					Expand:         []string{"description", "projectKeys"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/project/search?action=view&categoryId=400&expand=description%2CprojectKeys&maxResults=50&orderBy=category&query=dummy+project&startAt=0&typeKey=%5Bthepropertykey%5D.something.nested%3D1",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectSearchScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
				options: &model.ProjectSearchOptionsScheme{
					OrderBy:        "category",
					Query:          "dummy project",
					Action:         "view",
					ProjectKeyType: "[thepropertykey].something.nested=1",
					CategoryID:     400,
					Expand:         []string{"description", "projectKeys"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/project/search?action=view&categoryId=400&expand=description%2CprojectKeys&maxResults=50&orderBy=category&query=dummy+project&startAt=0&typeKey=%5Bthepropertykey%5D.something.nested%3D1",
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

			newService, err := NewProjectService(testCase.fields.c, testCase.fields.version, &ProjectChildServices{})
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Search(testCase.args.ctx, testCase.args.options, testCase.args.startAt,
				testCase.args.maxResults)

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

func Test_internalProjectImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx            context.Context
		projectKeyOrId string
		expand         []string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				expand:         []string{"issueTypes", "lead"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/project/KP?expand=issueTypes%2Clead",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				expand:         []string{"issueTypes", "lead"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/project/KP?expand=issueTypes%2Clead",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				expand:         []string{"issueTypes", "lead"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/project/KP?expand=issueTypes%2Clead",
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

			newService, err := NewProjectService(testCase.fields.c, testCase.fields.version, &ProjectChildServices{})
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Get(testCase.args.ctx, testCase.args.projectKeyOrId, testCase.args.expand)

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

func Test_internalProjectImpl_Update(t *testing.T) {

	payloadMocked := &model.ProjectUpdateScheme{
		NotificationScheme:  10021,
		Description:         "Example Project description",
		URL:                 "https://atlassian.com",
		ProjectTemplateKey:  "com.pyxis.greenhopper.jira:gh-simplified-agility-kanban",
		AvatarID:            10200,
		IssueSecurityScheme: 10001,
		Name:                "Project DUMMY #3",
		PermissionScheme:    10011,
		AssigneeType:        "PROJECT_LEAD",
		ProjectTypeKey:      "software",
		Key:                 "DUMMY3",
		CategoryID:          10120,
	}

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx            context.Context
		projectKeyOrId string
		payload        *model.ProjectUpdateScheme
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				payload:        payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/project/KP",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				payload:        payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/project/KP",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the project key or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDOrKeyError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				payload:        payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					payloadMocked).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/project/KP",
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

			newService, err := NewProjectService(testCase.fields.c, testCase.fields.version, &ProjectChildServices{})
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Update(testCase.args.ctx, testCase.args.projectKeyOrId, testCase.args.payload)

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

func Test_internalProjectImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx            context.Context
		projectKeyOrId string
		enableUndo     bool
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				enableUndo:     true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/project/KP?enableUndo=true",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				enableUndo:     true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/project/KP?enableUndo=true",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the project key or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDOrKeyError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				enableUndo:     true,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/project/KP?enableUndo=true",
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

			newService, err := NewProjectService(testCase.fields.c, testCase.fields.version, &ProjectChildServices{})
			assert.NoError(t, err)

			gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.projectKeyOrId, testCase.args.enableUndo)

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

func Test_internalProjectImpl_DeleteAsynchronously(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx            context.Context
		projectKeyOrId string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/project/KP/delete",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.TaskScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/project/KP/delete",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.TaskScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the project or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDOrKeyError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/project/KP/delete",
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

			newService, err := NewProjectService(testCase.fields.c, testCase.fields.version, &ProjectChildServices{})
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.DeleteAsynchronously(testCase.args.ctx, testCase.args.projectKeyOrId)

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

func Test_internalProjectImpl_Archive(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx            context.Context
		projectKeyOrId string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/project/KP/archive",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/project/KP/archive",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the project key or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDOrKeyError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/project/KP/archive",
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

			newService, err := NewProjectService(testCase.fields.c, testCase.fields.version, &ProjectChildServices{})
			assert.NoError(t, err)

			gotResponse, err := newService.Archive(testCase.args.ctx, testCase.args.projectKeyOrId)

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

func Test_internalProjectImpl_Restore(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx            context.Context
		projectKeyOrId string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/project/KP/restore",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/project/KP/restore",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ProjectScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the project or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDOrKeyError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/project/KP/restore",
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

			newService, err := NewProjectService(testCase.fields.c, testCase.fields.version, &ProjectChildServices{})
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Restore(testCase.args.ctx, testCase.args.projectKeyOrId)

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

func Test_internalProjectImpl_Statuses(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx            context.Context
		projectKeyOrId string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/project/KP/statuses",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/project/KP/statuses",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the project or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDOrKeyError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/project/KP/statuses",
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

			newService, err := NewProjectService(testCase.fields.c, testCase.fields.version, &ProjectChildServices{})
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Statuses(testCase.args.ctx, testCase.args.projectKeyOrId)

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

func Test_internalProjectImpl_NotificationScheme(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx            context.Context
		projectKeyOrId string
		expand         []string
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				expand:         []string{"all", "projectRole"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/project/KP/notificationscheme?expand=all%2CprojectRole",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.NotificationSchemeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				expand:         []string{"all", "projectRole"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/project/KP/notificationscheme?expand=all%2CprojectRole",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.NotificationSchemeScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			wantErr: false,
			Err:     nil,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:            context.TODO(),
				projectKeyOrId: "KP",
				expand:         []string{"all", "projectRole"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/project/KP/notificationscheme?expand=all%2CprojectRole",
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

			newService, err := NewProjectService(testCase.fields.c, testCase.fields.version, &ProjectChildServices{})
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.NotificationScheme(testCase.args.ctx, testCase.args.projectKeyOrId, testCase.args.expand)

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
