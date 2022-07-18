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

func Test_internalIssueFieldServiceImpl_Gets(t *testing.T) {

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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field",
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
				ctx: context.TODO(),
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field",
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
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field",
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

			fieldService, err := NewIssueFieldService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldService.Gets(testCase.args.ctx)

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

func Test_internalIssueFieldServiceImpl_Create(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx     context.Context
		payload *model.CustomFieldScheme
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
				payload: &model.CustomFieldScheme{
					Name:        "Alliance",
					Description: "this is the alliance description field",
					FieldType:   "cascadingselect",
					SearcherKey: "cascadingselectsearcher",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.CustomFieldScheme{
						Name:        "Alliance",
						Description: "this is the alliance description field",
						FieldType:   "cascadingselect",
						SearcherKey: "cascadingselectsearcher",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueFieldScheme{}).
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
				payload: &model.CustomFieldScheme{
					Name:        "Alliance",
					Description: "this is the alliance description field",
					FieldType:   "cascadingselect",
					SearcherKey: "cascadingselectsearcher",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.CustomFieldScheme{
						Name:        "Alliance",
						Description: "this is the alliance description field",
						FieldType:   "cascadingselect",
						SearcherKey: "cascadingselectsearcher",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/field",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueFieldScheme{}).
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
				ctx:     context.TODO(),
				payload: nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.CustomFieldScheme)(nil)).
					Return(bytes.NewReader([]byte{}), model.ErrNilPayloadError)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrNilPayloadError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
				payload: &model.CustomFieldScheme{
					Name:        "Alliance",
					Description: "this is the alliance description field",
					FieldType:   "cascadingselect",
					SearcherKey: "cascadingselectsearcher",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.CustomFieldScheme{
						Name:        "Alliance",
						Description: "this is the alliance description field",
						FieldType:   "cascadingselect",
						SearcherKey: "cascadingselectsearcher",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/field",
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

			fieldService, err := NewIssueFieldService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldService.Create(testCase.args.ctx, testCase.args.payload)

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

func Test_internalIssueFieldServiceImpl_Search(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx                 context.Context
		options             *model.FieldSearchOptionsScheme
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
				options: &model.FieldSearchOptionsScheme{
					Types:   []string{"custom"},
					IDs:     []string{"111", "12222"},
					Query:   "query-sample",
					OrderBy: "lastUsed",
					Expand:  []string{"screensCount", "lastUsed"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/search?expand=screensCount%2ClastUsed&id=111%2C12222&maxResults=50&orderBy=lastUsed&query=query-sample&startAt=0&type=custom",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldSearchPageScheme{}).
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
				options: &model.FieldSearchOptionsScheme{
					Types:   []string{"custom"},
					IDs:     []string{"111", "12222"},
					Query:   "query-sample",
					OrderBy: "lastUsed",
					Expand:  []string{"screensCount", "lastUsed"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/field/search?expand=screensCount%2ClastUsed&id=111%2C12222&maxResults=50&orderBy=lastUsed&query=query-sample&startAt=0&type=custom",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.FieldSearchPageScheme{}).
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
				options: &model.FieldSearchOptionsScheme{
					Types:   []string{"custom"},
					IDs:     []string{"111", "12222"},
					Query:   "query-sample",
					OrderBy: "lastUsed",
					Expand:  []string{"screensCount", "lastUsed"},
				},
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/search?expand=screensCount%2ClastUsed&id=111%2C12222&maxResults=50&orderBy=lastUsed&query=query-sample&startAt=0&type=custom",
					nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx: context.TODO(),
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/field/search?maxResults=0&startAt=0",
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

			fieldService, err := NewIssueFieldService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldService.Search(testCase.args.ctx, testCase.args.options,
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

func Test_internalIssueFieldServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Client
		version string
	}

	type args struct {
		ctx     context.Context
		fieldId string
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
				fieldId: "10005",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/field/10005",
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
				ctx:     context.TODO(),
				fieldId: "10005",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/field/10005",
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
			name:   "when the field id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "",
			},
			wantErr: true,
			Err:     model.ErrNoFieldIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "3"},
			args: args{
				ctx:     context.TODO(),
				fieldId: "10005",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/field/10005",
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

			fieldService, err := NewIssueFieldService(testCase.fields.c, testCase.fields.version, nil)
			assert.NoError(t, err)

			gotResult, gotResponse, err := fieldService.Delete(testCase.args.ctx, testCase.args.fieldId)

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
