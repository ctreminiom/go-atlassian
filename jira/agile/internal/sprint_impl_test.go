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

func Test_SprintService_Get(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx      context.Context
		sprintId int
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
				ctx:      context.Background(),
				sprintId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			Err:     errors.New("error, unable to execute the http call"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintId: 10001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001",
					nil).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},

		{
			name: "when the sprint id is not provided",
			args: args{
				ctx:      context.Background(),
				sprintId: 0,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoSprintIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			sprintService, err := NewSprintService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := sprintService.Get(testCase.args.ctx, testCase.args.sprintId)

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

func Test_SprintService_Create(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx     context.Context
		payload *model.SprintPayloadScheme
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
				ctx: context.Background(),
				payload: &model.SprintPayloadScheme{
					Name:          "Board Name Sample",
					StartDate:     "2015-04-20T01:22:00.000+10:00",
					EndDate:       "2015-04-11T15:22:00.000+10:00",
					OriginBoardID: 5,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.SprintPayloadScheme{
						Name:          "Board Name Sample",
						StartDate:     "2015-04-20T01:22:00.000+10:00",
						EndDate:       "2015-04-11T15:22:00.000+10:00",
						OriginBoardID: 5,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx: context.Background(),
				payload: &model.SprintPayloadScheme{
					Name:          "Board Name Sample",
					StartDate:     "2015-04-20T01:22:00.000+10:00",
					EndDate:       "2015-04-11T15:22:00.000+10:00",
					OriginBoardID: 5,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.SprintPayloadScheme{
						Name:          "Board Name Sample",
						StartDate:     "2015-04-20T01:22:00.000+10:00",
						EndDate:       "2015-04-11T15:22:00.000+10:00",
						OriginBoardID: 5,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			Err:     errors.New("error, unable to execute the http call"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx: context.Background(),
				payload: &model.SprintPayloadScheme{
					Name:          "Board Name Sample",
					StartDate:     "2015-04-20T01:22:00.000+10:00",
					EndDate:       "2015-04-11T15:22:00.000+10:00",
					OriginBoardID: 5,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.SprintPayloadScheme{
						Name:          "Board Name Sample",
						StartDate:     "2015-04-20T01:22:00.000+10:00",
						EndDate:       "2015-04-11T15:22:00.000+10:00",
						OriginBoardID: 5,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},

		{
			name: "when the payload is not provided",
			args: args{
				ctx:     context.Background(),
				payload: nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.SprintPayloadScheme)(nil)).
					Return(nil, errors.New("client: no payload provided"))

				fields.c = client
			},
			Err:     errors.New("client: no payload provided"),
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			sprintService, err := NewSprintService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := sprintService.Create(testCase.args.ctx, testCase.args.payload)

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

func Test_SprintService_Update(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx      context.Context
		sprintId int
		payload  *model.SprintPayloadScheme
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
				ctx:      context.Background(),
				sprintId: 1001,
				payload: &model.SprintPayloadScheme{
					Name:          "Board Name Sample",
					StartDate:     "2015-04-20T01:22:00.000+10:00",
					EndDate:       "2015-04-11T15:22:00.000+10:00",
					OriginBoardID: 5,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.SprintPayloadScheme{
						Name:          "Board Name Sample",
						StartDate:     "2015-04-20T01:22:00.000+10:00",
						EndDate:       "2015-04-11T15:22:00.000+10:00",
						OriginBoardID: 5,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/agile/1.0/sprint/1001",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintId: 1001,
				payload: &model.SprintPayloadScheme{
					Name:          "Board Name Sample",
					StartDate:     "2015-04-20T01:22:00.000+10:00",
					EndDate:       "2015-04-11T15:22:00.000+10:00",
					OriginBoardID: 5,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.SprintPayloadScheme{
						Name:          "Board Name Sample",
						StartDate:     "2015-04-20T01:22:00.000+10:00",
						EndDate:       "2015-04-11T15:22:00.000+10:00",
						OriginBoardID: 5,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/agile/1.0/sprint/1001",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			Err:     errors.New("error, unable to execute the http call"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintId: 1001,
				payload: &model.SprintPayloadScheme{
					Name:          "Board Name Sample",
					StartDate:     "2015-04-20T01:22:00.000+10:00",
					EndDate:       "2015-04-11T15:22:00.000+10:00",
					OriginBoardID: 5,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.SprintPayloadScheme{
						Name:          "Board Name Sample",
						StartDate:     "2015-04-20T01:22:00.000+10:00",
						EndDate:       "2015-04-11T15:22:00.000+10:00",
						OriginBoardID: 5,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/agile/1.0/sprint/1001",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},

		{
			name: "when the payload is not provided",
			args: args{
				ctx:      context.Background(),
				sprintId: 1001,
				payload:  nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.SprintPayloadScheme)(nil)).
					Return(nil, errors.New("client: no payload provided"))

				fields.c = client
			},
			Err:     errors.New("client: no payload provided"),
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			sprintService, err := NewSprintService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := sprintService.Update(testCase.args.ctx, testCase.args.sprintId, testCase.args.payload)

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

func Test_SprintService_Path(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx      context.Context
		sprintId int
		payload  *model.SprintPayloadScheme
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
				ctx:      context.Background(),
				sprintId: 1001,
				payload: &model.SprintPayloadScheme{
					Name:          "Board Name Sample",
					StartDate:     "2015-04-20T01:22:00.000+10:00",
					EndDate:       "2015-04-11T15:22:00.000+10:00",
					OriginBoardID: 5,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.SprintPayloadScheme{
						Name:          "Board Name Sample",
						StartDate:     "2015-04-20T01:22:00.000+10:00",
						EndDate:       "2015-04-11T15:22:00.000+10:00",
						OriginBoardID: 5,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintId: 1001,
				payload: &model.SprintPayloadScheme{
					Name:          "Board Name Sample",
					StartDate:     "2015-04-20T01:22:00.000+10:00",
					EndDate:       "2015-04-11T15:22:00.000+10:00",
					OriginBoardID: 5,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.SprintPayloadScheme{
						Name:          "Board Name Sample",
						StartDate:     "2015-04-20T01:22:00.000+10:00",
						EndDate:       "2015-04-11T15:22:00.000+10:00",
						OriginBoardID: 5,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			Err:     errors.New("error, unable to execute the http call"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintId: 1001,
				payload: &model.SprintPayloadScheme{
					Name:          "Board Name Sample",
					StartDate:     "2015-04-20T01:22:00.000+10:00",
					EndDate:       "2015-04-11T15:22:00.000+10:00",
					OriginBoardID: 5,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.SprintPayloadScheme{
						Name:          "Board Name Sample",
						StartDate:     "2015-04-20T01:22:00.000+10:00",
						EndDate:       "2015-04-11T15:22:00.000+10:00",
						OriginBoardID: 5,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},

		{
			name: "when the payload is not provided",
			args: args{
				ctx:      context.Background(),
				sprintId: 1001,
				payload:  nil,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					(*model.SprintPayloadScheme)(nil)).
					Return(nil, errors.New("client: no payload provided"))

				fields.c = client
			},
			Err:     errors.New("client: no payload provided"),
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			sprintService, err := NewSprintService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := sprintService.Path(testCase.args.ctx, testCase.args.sprintId, testCase.args.payload)

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

func Test_SprintService_Delete(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx      context.Context
		sprintId int
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
				ctx:      context.Background(),
				sprintId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/agile/1.0/sprint/1001",
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
			name: "when the sprintId is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoSprintIDError,
			wantErr: true,
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/agile/1.0/sprint/1001",
					nil).
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
				ctx:      context.Background(),
				sprintId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/agile/1.0/sprint/1001",
					nil).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			sprintService, err := NewSprintService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResponse, err := sprintService.Delete(testCase.args.ctx, testCase.args.sprintId)

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

func Test_SprintService_Issues(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx                 context.Context
		sprintId            int
		opts                *model.IssueOptionScheme
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
				ctx:      context.Background(),
				sprintId: 10001,
				opts: &model.IssueOptionScheme{
					JQL:           "project = ABC",
					ValidateQuery: true,
					Fields:        []string{"summary", "status"},
					Expand:        []string{"changelog"},
				},
				startAt:    100,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001/issue?expand=changelog&fields=summary%2Cstatus&jql=project+%3D+ABC&maxResults=50&startAt=100",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintIssuePageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintId: 10001,
				opts: &model.IssueOptionScheme{
					JQL:           "project = ABC",
					ValidateQuery: true,
					Fields:        []string{"summary", "status"},
					Expand:        []string{"changelog"},
				},
				startAt:    100,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001/issue?expand=changelog&fields=summary%2Cstatus&jql=project+%3D+ABC&maxResults=50&startAt=100",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SprintIssuePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			Err:     errors.New("error, unable to execute the http call"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:      context.Background(),
				sprintId: 10001,
				opts: &model.IssueOptionScheme{
					JQL:           "project = ABC",
					ValidateQuery: true,
					Fields:        []string{"summary", "status"},
					Expand:        []string{"changelog"},
				},
				startAt:    100,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/agile/1.0/sprint/10001/issue?expand=changelog&fields=summary%2Cstatus&jql=project+%3D+ABC&maxResults=50&startAt=100",
					nil).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},

		{
			name: "when the sprint id is not provided",
			args: args{
				ctx:      context.Background(),
				sprintId: 0,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoSprintIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			sprintService, err := NewSprintService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := sprintService.Issues(testCase.args.ctx, testCase.args.sprintId, testCase.args.opts,
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

func Test_SprintService_Start(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx      context.Context
		sprintId int
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
				ctx:      context.Background(),
				sprintId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					model.SprintPayloadScheme{
						State: "Active",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
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
			name: "when the sprintId is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoSprintIDError,
			wantErr: true,
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					model.SprintPayloadScheme{
						State: "Active",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
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
				ctx:      context.Background(),
				sprintId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					model.SprintPayloadScheme{
						State: "Active",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			sprintService, err := NewSprintService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResponse, err := sprintService.Start(testCase.args.ctx, testCase.args.sprintId)

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

func Test_SprintService_Close(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx      context.Context
		sprintId int
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
				ctx:      context.Background(),
				sprintId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					model.SprintPayloadScheme{
						State: "Closed",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
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
			name: "when the sprintId is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoSprintIDError,
			wantErr: true,
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:      context.Background(),
				sprintId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					model.SprintPayloadScheme{
						State: "Closed",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
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
				ctx:      context.Background(),
				sprintId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					model.SprintPayloadScheme{
						State: "Closed",
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/sprint/1001",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			sprintService, err := NewSprintService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResponse, err := sprintService.Close(testCase.args.ctx, testCase.args.sprintId)

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
