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

func Test_EpicService_Get(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx         context.Context
		epicIdOrKey string
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
				epicIdOrKey: "EPIC-1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.EpicScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:         context.Background(),
				epicIdOrKey: "EPIC-1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.EpicScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			Err:     errors.New("error, unable to execute the http call"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:         context.Background(),
				epicIdOrKey: "EPIC-1",
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1",
					nil).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},

		{
			name: "when the epic id is not provided",
			args: args{
				ctx:         context.Background(),
				epicIdOrKey: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoEpicIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewEpicService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.Get(testCase.args.ctx, testCase.args.epicIdOrKey)

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

func Test_EpicService_Issues(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx         context.Context
		epicIdOrKey string
		startAt     int
		maxResults  int
		opts        *model.IssueOptionScheme
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
				epicIdOrKey: "EPIC-1",
				startAt:     10,
				maxResults:  50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = EPIC",
					ValidateQuery: true,
					Fields:        []string{"status", "summary"},
					Expand:        []string{"changelogs"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1/issue?expand=changelogs&fields=status%2Csummary&jql=project+%3D+EPIC&maxResults=50&startAt=10&validateQuery=true",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardIssuePageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:         context.Background(),
				epicIdOrKey: "EPIC-1",
				startAt:     10,
				maxResults:  50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = EPIC",
					ValidateQuery: true,
					Fields:        []string{"status", "summary"},
					Expand:        []string{"changelogs"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1/issue?expand=changelogs&fields=status%2Csummary&jql=project+%3D+EPIC&maxResults=50&startAt=10&validateQuery=true",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardIssuePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			Err:     errors.New("error, unable to execute the http call"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:         context.Background(),
				epicIdOrKey: "EPIC-1",
				startAt:     10,
				maxResults:  50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = EPIC",
					ValidateQuery: true,
					Fields:        []string{"status", "summary"},
					Expand:        []string{"changelogs"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1/issue?expand=changelogs&fields=status%2Csummary&jql=project+%3D+EPIC&maxResults=50&startAt=10&validateQuery=true",
					nil).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},

		{
			name: "when the epic id is not provided",
			args: args{
				ctx:         context.Background(),
				epicIdOrKey: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoEpicIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewEpicService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.Issues(testCase.args.ctx, testCase.args.epicIdOrKey, testCase.args.opts,
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

func Test_EpicService_Move(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx         context.Context
		epicIdOrKey string
		issues      []string
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
				epicIdOrKey: "EPIC-1",
				issues:      []string{"EPIC-10"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					map[string]interface{}{"issues": []string{"EPIC-10"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/epic/EPIC-1/issue",
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
			name: "when the api cannot be executed",
			args: args{
				ctx:         context.Background(),
				epicIdOrKey: "EPIC-1",
				issues:      []string{"EPIC-10"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					map[string]interface{}{"issues": []string{"EPIC-10"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/epic/EPIC-1/issue",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			Err:     errors.New("error, unable to execute the http call"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:         context.Background(),
				epicIdOrKey: "EPIC-1",
				issues:      []string{"EPIC-10"},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					map[string]interface{}{"issues": []string{"EPIC-10"}}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/epic/EPIC-1/issue",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},

		{
			name: "when the epic id is not provided",
			args: args{
				ctx:         context.Background(),
				epicIdOrKey: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoEpicIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewEpicService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResponse, err := service.Move(testCase.args.ctx, testCase.args.epicIdOrKey, testCase.args.issues)

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
