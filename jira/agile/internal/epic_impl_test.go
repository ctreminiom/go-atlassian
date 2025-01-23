package internal

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
)

func Test_EpicService_Get(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		epicIDOrKey string
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
				epicIDOrKey: "EPIC-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1",
					"",
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
				epicIDOrKey: "EPIC-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1",
					"",
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
				epicIDOrKey: "EPIC-1",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1",
					"",
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
				epicIDOrKey: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoEpicID,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			epicService := NewEpicService(testCase.fields.c, "1.0")

			gotResult, gotResponse, err := epicService.Get(testCase.args.ctx, testCase.args.epicIDOrKey)

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
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		epicIDOrKey string
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
				epicIDOrKey: "EPIC-1",
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

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1/issue?expand=changelogs&fields=status%2Csummary&jql=project+%3D+EPIC&maxResults=50&startAt=10&validateQuery=true",
					"",
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
				epicIDOrKey: "EPIC-1",
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

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1/issue?expand=changelogs&fields=status%2Csummary&jql=project+%3D+EPIC&maxResults=50&startAt=10&validateQuery=true",
					"",
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
				epicIDOrKey: "EPIC-1",
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

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/epic/EPIC-1/issue?expand=changelogs&fields=status%2Csummary&jql=project+%3D+EPIC&maxResults=50&startAt=10&validateQuery=true",
					"",
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
				epicIDOrKey: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoEpicID,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			epicService := NewEpicService(testCase.fields.c, "1.0")

			gotResult, gotResponse, err := epicService.Issues(testCase.args.ctx, testCase.args.epicIDOrKey, testCase.args.opts,
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

	payloadMocked := map[string]interface{}{"issues": []string{"EPIC-10"}}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx         context.Context
		epicIDOrKey string
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
				epicIDOrKey: "EPIC-1",
				issues:      []string{"EPIC-10"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/epic/EPIC-1/issue",
					"",
					payloadMocked).
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
				epicIDOrKey: "EPIC-1",
				issues:      []string{"EPIC-10"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/epic/EPIC-1/issue",
					"",
					payloadMocked).
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
				epicIDOrKey: "EPIC-1",
				issues:      []string{"EPIC-10"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/epic/EPIC-1/issue",
					"",
					payloadMocked).
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
				epicIDOrKey: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			Err:     model.ErrNoEpicID,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			epicService := NewEpicService(testCase.fields.c, "1.0")

			gotResponse, err := epicService.Move(testCase.args.ctx, testCase.args.epicIDOrKey, testCase.args.issues)

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
