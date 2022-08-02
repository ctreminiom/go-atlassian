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

func Test_internalIssueFieldConfigSchemeServiceImpl_Gets(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                 context.Context
		ids                 []int
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
				ctx:        context.TODO(),
				ids:        []int{10001},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/fieldconfigurationscheme?id=10001&maxResults=50&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldConfigurationSchemePageScheme{}).
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
				ctx:        context.TODO(),
				ids:        []int{10001},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/fieldconfigurationscheme?id=10001&maxResults=50&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldConfigurationSchemePageScheme{}).
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
				ctx:        context.TODO(),
				ids:        []int{10001},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/fieldconfigurationscheme?id=10001&maxResults=50&startAt=50",
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

			fieldConfigService, err := NewIssueFieldConfigurationSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Gets(testCase.args.ctx, testCase.args.ids,
				testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalIssueFieldConfigSchemeServiceImpl_Create(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx               context.Context
		name, description string
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
				ctx:         context.TODO(),
				name:        "field scheme sample",
				description: "field scheme sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						Name        string "json:\"name\""
						Description string "json:\"description,omitempty\""
					}{
						Name:        "field scheme sample",
						Description: "field scheme sample"},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/fieldconfigurationscheme",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldConfigurationSchemeScheme{}).
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
				ctx:         context.TODO(),
				name:        "field scheme sample",
				description: "field scheme sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						Name        string "json:\"name\""
						Description string "json:\"description,omitempty\""
					}{
						Name:        "field scheme sample",
						Description: "field scheme sample"},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/fieldconfigurationscheme",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldConfigurationSchemeScheme{}).
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
				ctx:         context.TODO(),
				name:        "field scheme sample",
				description: "field scheme sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						Name        string "json:\"name\""
						Description string "json:\"description,omitempty\""
					}{
						Name:        "field scheme sample",
						Description: "field scheme sample"},
				).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/fieldconfigurationscheme",
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

			fieldConfigService, err := NewIssueFieldConfigurationSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Create(testCase.args.ctx, testCase.args.name,
				testCase.args.description)

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

func Test_internalIssueFieldConfigSchemeServiceImpl_Mapping(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                 context.Context
		fieldConfigIds      []int
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
				ctx:            context.TODO(),
				fieldConfigIds: []int{10001},
				startAt:        50,
				maxResults:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=10001&maxResults=50&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldConfigurationIssueTypeItemPageScheme{}).
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
				fieldConfigIds: []int{10001},
				startAt:        50,
				maxResults:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=10001&maxResults=50&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldConfigurationIssueTypeItemPageScheme{}).
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
				fieldConfigIds: []int{10001},
				startAt:        50,
				maxResults:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=10001&maxResults=50&startAt=50",
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

			fieldConfigService, err := NewIssueFieldConfigurationSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Mapping(testCase.args.ctx, testCase.args.fieldConfigIds,
				testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalIssueFieldConfigSchemeServiceImpl_Project(t *testing.T) {

	type fields struct {
		c       service.Client
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
				ctx:        context.TODO(),
				projectIds: []int{10001},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/fieldconfigurationscheme/project?maxResults=50&projectId=10001&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldConfigurationSchemeProjectPageScheme{}).
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
				ctx:        context.TODO(),
				projectIds: []int{10001},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/fieldconfigurationscheme/project?maxResults=50&projectId=10001&startAt=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldConfigurationSchemeProjectPageScheme{}).
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
				ctx:        context.TODO(),
				projectIds: []int{10001},
				startAt:    50,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/fieldconfigurationscheme/project?maxResults=50&projectId=10001&startAt=50",
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

			fieldConfigService, err := NewIssueFieldConfigurationSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldConfigService.Project(testCase.args.ctx, testCase.args.projectIds,
				testCase.args.startAt, testCase.args.maxResults)

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

func Test_internalIssueFieldConfigSchemeServiceImpl_Assign(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.FieldConfigurationSchemeAssignPayload
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
				payload: &model.FieldConfigurationSchemeAssignPayload{
					FieldConfigurationSchemeID: "10000",
					ProjectID:                  "10000",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldConfigurationSchemeAssignPayload{
						FieldConfigurationSchemeID: "10000",
						ProjectID:                  "10000",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/fieldconfigurationscheme/project",
					bytes.NewReader([]byte{})).
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
				ctx: context.TODO(),
				payload: &model.FieldConfigurationSchemeAssignPayload{
					FieldConfigurationSchemeID: "10000",
					ProjectID:                  "10000",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldConfigurationSchemeAssignPayload{
						FieldConfigurationSchemeID: "10000",
						ProjectID:                  "10000",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/fieldconfigurationscheme/project",
					bytes.NewReader([]byte{})).
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
			name:   "when the payload is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
				payload: &model.FieldConfigurationSchemeAssignPayload{
					FieldConfigurationSchemeID: "10000",
					ProjectID:                  "10000",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldConfigurationSchemeAssignPayload{
						FieldConfigurationSchemeID: "10000",
						ProjectID:                  "10000",
					}).
					Return(bytes.NewReader([]byte{}), model.ErrNonPayloadPointerError)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrNonPayloadPointerError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldConfigurationSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Assign(testCase.args.ctx, testCase.args.payload)

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

func Test_internalIssueFieldConfigSchemeServiceImpl_Update(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx               context.Context
		schemeId          int
		name, description string
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
				ctx:         context.TODO(),
				schemeId:    10001,
				name:        "name sample",
				description: "description sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						Name        string "json:\"name\""
						Description string "json:\"description,omitempty\""
					}{Name: "name sample", Description: "description sample"}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/fieldconfigurationscheme/10001",
					bytes.NewReader([]byte{})).
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
			name:   "when the scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:         context.TODO(),
				schemeId:    0,
				name:        "name sample",
				description: "description sample",
			},
			wantErr: true,
			Err:     model.ErrNoFieldConfigurationSchemeIDError,
		},

		{
			name:   "when the scheme name is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:         context.TODO(),
				schemeId:    10001,
				name:        "",
				description: "description sample",
			},
			wantErr: true,
			Err:     model.ErrNoFieldConfigurationSchemeNameError,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:         context.TODO(),
				schemeId:    10001,
				name:        "name sample",
				description: "description sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						Name        string "json:\"name\""
						Description string "json:\"description,omitempty\""
					}{Name: "name sample", Description: "description sample"}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/fieldconfigurationscheme/10001",
					bytes.NewReader([]byte{})).
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
			name:   "when the payload is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:         context.TODO(),
				schemeId:    10001,
				name:        "name sample",
				description: "description sample",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						Name        string "json:\"name\""
						Description string "json:\"description,omitempty\""
					}{Name: "name sample", Description: "description sample"}).
					Return(bytes.NewReader([]byte{}), model.ErrNonPayloadPointerError)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrNonPayloadPointerError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldConfigurationSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Update(testCase.args.ctx, testCase.args.schemeId, testCase.args.name, testCase.args.description)

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

func Test_internalIssueFieldConfigSchemeServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx      context.Context
		schemeId int
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
				ctx:      context.TODO(),
				schemeId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/fieldconfigurationscheme/10001",
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
			name:   "when the scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.TODO(),
				schemeId: 0,
			},
			wantErr: true,
			Err:     model.ErrNoFieldConfigurationSchemeIDError,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:      context.TODO(),
				schemeId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/fieldconfigurationscheme/10001",
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			fieldConfigService, err := NewIssueFieldConfigurationSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Delete(testCase.args.ctx, testCase.args.schemeId)

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

func Test_internalIssueFieldConfigSchemeServiceImpl_Link(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx      context.Context
		schemeId int
		payload  *model.FieldConfigurationToIssueTypeMappingPayloadScheme
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
				ctx:      context.TODO(),
				schemeId: 10001,
				payload: &model.FieldConfigurationToIssueTypeMappingPayloadScheme{
					Mappings: []*model.FieldConfigurationToIssueTypeMappingScheme{
						{
							IssueTypeID:          "default",
							FieldConfigurationID: "10000",
						},
						{
							IssueTypeID:          "10001",
							FieldConfigurationID: "10002",
						},
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldConfigurationToIssueTypeMappingPayloadScheme{
						Mappings: []*model.FieldConfigurationToIssueTypeMappingScheme{
							{
								IssueTypeID:          "default",
								FieldConfigurationID: "10000",
							},
							{
								IssueTypeID:          "10001",
								FieldConfigurationID: "10002",
							},
						},
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/fieldconfigurationscheme/10001/mapping",
					bytes.NewReader([]byte{})).
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
			name:   "when the scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.TODO(),
				schemeId: 0,
				payload: &model.FieldConfigurationToIssueTypeMappingPayloadScheme{
					Mappings: []*model.FieldConfigurationToIssueTypeMappingScheme{
						{
							IssueTypeID:          "default",
							FieldConfigurationID: "10000",
						},
						{
							IssueTypeID:          "10001",
							FieldConfigurationID: "10002",
						},
					},
				},
			},
			wantErr: true,
			Err:     model.ErrNoFieldConfigurationSchemeIDError,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:      context.TODO(),
				schemeId: 10001,
				payload: &model.FieldConfigurationToIssueTypeMappingPayloadScheme{
					Mappings: []*model.FieldConfigurationToIssueTypeMappingScheme{
						{
							IssueTypeID:          "default",
							FieldConfigurationID: "10000",
						},
						{
							IssueTypeID:          "10001",
							FieldConfigurationID: "10002",
						},
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldConfigurationToIssueTypeMappingPayloadScheme{
						Mappings: []*model.FieldConfigurationToIssueTypeMappingScheme{
							{
								IssueTypeID:          "default",
								FieldConfigurationID: "10000",
							},
							{
								IssueTypeID:          "10001",
								FieldConfigurationID: "10002",
							},
						},
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/fieldconfigurationscheme/10001/mapping",
					bytes.NewReader([]byte{})).
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
				ctx:      context.TODO(),
				schemeId: 10001,
				payload: &model.FieldConfigurationToIssueTypeMappingPayloadScheme{
					Mappings: []*model.FieldConfigurationToIssueTypeMappingScheme{
						{
							IssueTypeID:          "default",
							FieldConfigurationID: "10000",
						},
						{
							IssueTypeID:          "10001",
							FieldConfigurationID: "10002",
						},
					},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.FieldConfigurationToIssueTypeMappingPayloadScheme{
						Mappings: []*model.FieldConfigurationToIssueTypeMappingScheme{
							{
								IssueTypeID:          "default",
								FieldConfigurationID: "10000",
							},
							{
								IssueTypeID:          "10001",
								FieldConfigurationID: "10002",
							},
						},
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/3/fieldconfigurationscheme/10001/mapping",
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

			fieldConfigService, err := NewIssueFieldConfigurationSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Link(testCase.args.ctx, testCase.args.schemeId, testCase.args.payload)

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

func Test_internalIssueFieldConfigSchemeServiceImpl_Unlink(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx          context.Context
		schemeId     int
		issueTypeIDs []string
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
				ctx:          context.TODO(),
				schemeId:     10001,
				issueTypeIDs: []string{"1001", "1002"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						IssueTypeIds []string "json:\"issueTypeIds\""
					}{IssueTypeIds: []string{"1001", "1002"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/fieldconfigurationscheme/10001/mapping/delete",
					bytes.NewReader([]byte{})).
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
				ctx:          context.TODO(),
				schemeId:     10001,
				issueTypeIDs: []string{"1001", "1002"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						IssueTypeIds []string "json:\"issueTypeIds\""
					}{IssueTypeIds: []string{"1001", "1002"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/fieldconfigurationscheme/10001/mapping/delete",
					bytes.NewReader([]byte{})).
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
			name:   "when the scheme id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:      context.TODO(),
				schemeId: 0,
			},
			wantErr: true,
			Err:     model.ErrNoFieldConfigurationSchemeIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.TODO(),
				schemeId:     10001,
				issueTypeIDs: []string{"1001", "1002"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&struct {
						IssueTypeIds []string "json:\"issueTypeIds\""
					}{IssueTypeIds: []string{"1001", "1002"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/fieldconfigurationscheme/10001/mapping/delete",
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

			fieldConfigService, err := NewIssueFieldConfigurationSchemeService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := fieldConfigService.Unlink(testCase.args.ctx, testCase.args.schemeId, testCase.args.issueTypeIDs)

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
