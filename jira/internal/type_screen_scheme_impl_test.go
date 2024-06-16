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

func Test_internalTypeScreenSchemeImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		options             *model.ScreenSchemeParamsScheme
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
				ctx: context.Background(),
				options: &model.ScreenSchemeParamsScheme{
					IDs:         []int{10001, 10002},
					QueryString: "query",
					OrderBy:     "id",
					Expand:      []string{"expand"},
				},
				startAt:    50,
				maxResults: 100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issuetypescreenscheme?=id&expand=expand&id=10001&id=10002&maxResults=100&queryString=query&startAt=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeScreenSchemePageScheme{}).
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
				ctx: context.Background(),
				options: &model.ScreenSchemeParamsScheme{
					IDs:         []int{10001, 10002},
					QueryString: "query",
					OrderBy:     "id",
					Expand:      []string{"expand"},
				},
				startAt:    50,
				maxResults: 100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issuetypescreenscheme?=id&expand=expand&id=10001&id=10002&maxResults=100&queryString=query&startAt=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeScreenSchemePageScheme{}).
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
				ctx: context.Background(),
				options: &model.ScreenSchemeParamsScheme{
					IDs:         []int{10001, 10002},
					QueryString: "query",
					OrderBy:     "id",
					Expand:      []string{"expand"},
				},
				startAt:    50,
				maxResults: 100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issuetypescreenscheme?=id&expand=expand&id=10001&id=10002&maxResults=100&queryString=query&startAt=50",
					"", nil).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Gets(testCase.args.ctx, testCase.args.options, testCase.args.startAt,
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

func Test_internalTypeScreenSchemeImpl_Projects(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                 context.Context
		projectIds          []int
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
				ctx:        context.Background(),
				projectIds: []int{29992, 349383},
				startAt:    50,
				maxResults: 100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issuetypescreenscheme/project?maxResults=100&projectId=29992&projectId=349383&startAt=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeProjectScreenSchemePageScheme{}).
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
				ctx:        context.Background(),
				projectIds: []int{29992, 349383},
				startAt:    50,
				maxResults: 100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issuetypescreenscheme/project?maxResults=100&projectId=29992&projectId=349383&startAt=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeProjectScreenSchemePageScheme{}).
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
				ctx:        context.Background(),
				projectIds: []int{29992, 349383},
				startAt:    50,
				maxResults: 100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issuetypescreenscheme/project?maxResults=100&projectId=29992&projectId=349383&startAt=50",
					"", nil).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Projects(testCase.args.ctx, testCase.args.projectIds, testCase.args.startAt,
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

func Test_internalTypeScreenSchemeImpl_Mapping(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                      context.Context
		issueTypeScreenSchemeIds []int
		startAt, maxResults      int
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
				ctx:                      context.Background(),
				issueTypeScreenSchemeIds: []int{29992, 349383},
				startAt:                  50,
				maxResults:               100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issuetypescreenscheme/mapping?issueTypeScreenSchemeId=29992&issueTypeScreenSchemeId=349383&maxResults=100&startAt=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeScreenSchemeMappingScheme{}).
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
				ctx:                      context.Background(),
				issueTypeScreenSchemeIds: []int{29992, 349383},
				startAt:                  50,
				maxResults:               100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issuetypescreenscheme/mapping?issueTypeScreenSchemeId=29992&issueTypeScreenSchemeId=349383&maxResults=100&startAt=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeScreenSchemeMappingScheme{}).
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
				ctx:                      context.Background(),
				issueTypeScreenSchemeIds: []int{29992, 349383},
				startAt:                  50,
				maxResults:               100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issuetypescreenscheme/mapping?issueTypeScreenSchemeId=29992&issueTypeScreenSchemeId=349383&maxResults=100&startAt=50",
					"", nil).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.Mapping(testCase.args.ctx, testCase.args.issueTypeScreenSchemeIds, testCase.args.startAt,
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

func Test_internalTypeScreenSchemeImpl_SchemesByProject(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                     context.Context
		issueTypeScreenSchemeId int
		startAt, maxResults     int
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
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: 29992,
				startAt:                 50,
				maxResults:              100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issuetypescreenscheme/29992/project?maxResults=100&startAt=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeScreenSchemeByProjectPageScheme{}).
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
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: 29992,
				startAt:                 50,
				maxResults:              100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issuetypescreenscheme/29992/project?maxResults=100&startAt=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeScreenSchemeByProjectPageScheme{}).
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
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: 29992,
				startAt:                 50,
				maxResults:              100,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issuetypescreenscheme/29992/project?maxResults=100&startAt=50",
					"", nil).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := newService.SchemesByProject(testCase.args.ctx, testCase.args.issueTypeScreenSchemeId, testCase.args.startAt,
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

func Test_internalTypeScreenSchemeImpl_Assign(t *testing.T) {

	payloadMocked := map[string]interface{}{"issueTypeScreenSchemeId": "20001", "projectId": "848483"}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                                context.Context
		issueTypeScreenSchemeId, projectId string
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
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				projectId:               "848483",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issuetypescreenscheme/project",
					"", payloadMocked).
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
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				projectId:               "848483",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issuetypescreenscheme/project",
					"", payloadMocked).
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
			name:   "when the issue type screen scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypeScreenSchemeIDError,
		},

		{
			name:   "when the project id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
			},
			wantErr: true,
			Err:     model.ErrNoProjectIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				projectId:               "848483",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issuetypescreenscheme/project",
					"", payloadMocked).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := newService.Assign(testCase.args.ctx, testCase.args.issueTypeScreenSchemeId, testCase.args.projectId)

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

func Test_internalTypeScreenSchemeImpl_Update(t *testing.T) {

	payloadMocked := map[string]interface{}{"description": "New issue type scheme description", "name": "New issue type scheme name"}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                                        context.Context
		issueTypeScreenSchemeId, name, description string
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
			name:   "when the issue type screen scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypeScreenSchemeIDError,
		},
		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				name:                    "New issue type scheme name",
				description:             "New issue type scheme description",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issuetypescreenscheme/20001",
					"", payloadMocked).
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
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				name:                    "New issue type scheme name",
				description:             "New issue type scheme description",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issuetypescreenscheme/20001",
					"", payloadMocked).
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
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				name:                    "New issue type scheme name",
				description:             "New issue type scheme description",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issuetypescreenscheme/20001",
					"", payloadMocked).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := newService.Update(testCase.args.ctx, testCase.args.issueTypeScreenSchemeId, testCase.args.name,
				testCase.args.description)

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

func Test_internalTypeScreenSchemeImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                     context.Context
		issueTypeScreenSchemeId string
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
			name:   "when the issue type screen scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypeScreenSchemeIDError,
		},
		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issuetypescreenscheme/20001",
					"", nil).
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
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/issuetypescreenscheme/20001",
					"", nil).
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
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/issuetypescreenscheme/20001",
					"", nil).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := newService.Delete(testCase.args.ctx, testCase.args.issueTypeScreenSchemeId)

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

func Test_internalTypeScreenSchemeImpl_Append(t *testing.T) {

	payloadMocked := &model.IssueTypeScreenSchemePayloadScheme{
		IssueTypeMappings: []*model.IssueTypeScreenSchemeMappingPayloadScheme{
			{
				IssueTypeID:    "10000",
				ScreenSchemeID: "10001",
			},
			{
				IssueTypeID:    "10001",
				ScreenSchemeID: "10002",
			},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                     context.Context
		issueTypeScreenSchemeId string
		payload                 *model.IssueTypeScreenSchemePayloadScheme
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
			name:   "when the issue type screen scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypeScreenSchemeIDError,
		},
		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				payload:                 payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issuetypescreenscheme/20001/mapping",
					"", payloadMocked).
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
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				payload:                 payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issuetypescreenscheme/20001/mapping",
					"", payloadMocked).
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
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				payload:                 payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issuetypescreenscheme/20001/mapping",
					"", payloadMocked).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := newService.Append(testCase.args.ctx, testCase.args.issueTypeScreenSchemeId,
				testCase.args.payload)

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

func Test_internalTypeScreenSchemeImpl_UpdateDefault(t *testing.T) {

	payloadMocked := map[string]interface{}{"screenSchemeId": "200202"}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                                     context.Context
		issueTypeScreenSchemeId, screenSchemeId string
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
			name:   "when the issue type screen scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypeScreenSchemeIDError,
		},
		{
			name:   "when the screen scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
			},
			wantErr: true,
			Err:     model.ErrNoScreenSchemeIDError,
		},
		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				screenSchemeId:          "200202",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issuetypescreenscheme/20001/mapping/default",
					"", payloadMocked).
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
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				screenSchemeId:          "200202",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issuetypescreenscheme/20001/mapping/default",
					"", payloadMocked).
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
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				screenSchemeId:          "200202",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/issuetypescreenscheme/20001/mapping/default",
					"", payloadMocked).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := newService.UpdateDefault(testCase.args.ctx, testCase.args.issueTypeScreenSchemeId,
				testCase.args.screenSchemeId)

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

func Test_internalTypeScreenSchemeImpl_Remove(t *testing.T) {

	payloadMocked := map[string]interface{}{"issueTypeIds": []string{"9", "43"}}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                     context.Context
		issueTypeScreenSchemeId string
		issueTypeIds            []string
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
			name:   "when the issue type screen scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "",
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypeScreenSchemeIDError,
		},
		{
			name:   "when the issue type id's are not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "2201",
			},
			wantErr: true,
			Err:     model.ErrNoIssueTypesError,
		},
		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				issueTypeIds:            []string{"9", "43"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issuetypescreenscheme/20001/mapping/remove",
					"", payloadMocked).
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
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				issueTypeIds:            []string{"9", "43"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issuetypescreenscheme/20001/mapping/remove",
					"", payloadMocked).
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
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:                     context.Background(),
				issueTypeScreenSchemeId: "20001",
				issueTypeIds:            []string{"9", "43"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issuetypescreenscheme/20001/mapping/remove",
					"", payloadMocked).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := newService.Remove(testCase.args.ctx, testCase.args.issueTypeScreenSchemeId,
				testCase.args.issueTypeIds)

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

func Test_internalTypeScreenSchemeImpl_Create(t *testing.T) {

	payloadMocked := &model.IssueTypeScreenSchemePayloadScheme{
		Name: "FX 2 Issue Type Screen Scheme",
		IssueTypeMappings: []*model.IssueTypeScreenSchemeMappingPayloadScheme{
			{
				IssueTypeID:    "default",
				ScreenSchemeID: "10000",
			},
			{
				IssueTypeID:    "10004", // Bug
				ScreenSchemeID: "10002",
			},
		},
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.IssueTypeScreenSchemePayloadScheme
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
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issuetypescreenscheme",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeScreenScreenCreatedScheme{}).
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
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issuetypescreenscheme",
					"", payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueTypeScreenScreenCreatedScheme{}).
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
				ctx:     context.Background(),
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issuetypescreenscheme",
					"", payloadMocked).
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

			newService, err := NewTypeScreenSchemeService(testCase.fields.c, testCase.fields.version)
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

func Test_NewTypeScreenSchemeService(t *testing.T) {

	type args struct {
		client  service.Connector
		version string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				client:  nil,
				version: "3",
			},
			wantErr: false,
		},

		{
			name: "when the version is not provided",
			args: args{
				client:  nil,
				version: "",
			},
			wantErr: true,
			err:     model.ErrNoVersionProvided,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := NewTypeScreenSchemeService(testCase.args.client, testCase.args.version)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, got, nil)
			}
		})
	}
}
