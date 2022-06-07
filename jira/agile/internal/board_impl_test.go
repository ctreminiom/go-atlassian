package internal

import (
	"bytes"
	"context"
	"errors"
	"github.com/ctreminiom/go-atlassian/jira/agile/internal/mocks"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/agile"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func Test_BoardService_Get(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx     context.Context
		boardId int
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
				ctx:     context.Background(),
				boardId: 1,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx:     context.Background(),
				boardId: 1,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardScheme{}).
					Return(&model.ResponseScheme{}, errors.New("error, unable to execute the http call"))

				fields.c = client
			},
			Err:     errors.New("error, unable to execute the http call"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:     context.Background(),
				boardId: 1,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1",
					nil).
					Return(&http.Request{}, errors.New("unable to create the http request"))

				fields.c = client
			},
			Err:     errors.New("unable to create the http request"),
			wantErr: true,
		},

		{
			name: "when the board id is not provided",
			args: args{
				ctx:     context.Background(),
				boardId: 0,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoBoardIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.Get(testCase.args.ctx, testCase.args.boardId)

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

func Test_BoardService_Create(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx     context.Context
		payload *model.BoardPayloadScheme
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
				payload: &model.BoardPayloadScheme{
					Name:     "Board Name Sample",
					Type:     "scrum",
					FilterID: 1002,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.BoardPayloadScheme{
						Name:     "Board Name Sample",
						Type:     "scrum",
						FilterID: 1002,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"/rest/agile/1.0/board",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the api cannot be executed",
			args: args{
				ctx: context.Background(),
				payload: &model.BoardPayloadScheme{
					Name:     "Board Name Sample",
					Type:     "scrum",
					FilterID: 1002,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.BoardPayloadScheme{
						Name:     "Board Name Sample",
						Type:     "scrum",
						FilterID: 1002,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"/rest/agile/1.0/board",
					bytes.NewReader([]byte{})).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardScheme{}).
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
				payload: &model.BoardPayloadScheme{
					Name:     "Board Name Sample",
					Type:     "scrum",
					FilterID: 1002,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("TransformStructToReader",
					&model.BoardPayloadScheme{
						Name:     "Board Name Sample",
						Type:     "scrum",
						FilterID: 1002,
					}).
					Return(bytes.NewReader([]byte{}), nil)

				client.On("NewJsonRequest",
					context.Background(),
					http.MethodPost,
					"/rest/agile/1.0/board",
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
					(*model.BoardPayloadScheme)(nil)).
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

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.Create(testCase.args.ctx, testCase.args.payload)

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

func Test_BoardService_Backlog(t *testing.T) {
	type fields struct {
		c service.Client
	}

	type args struct {
		ctx        context.Context
		boardId    int
		startAt    int
		maxResults int
		opts       *model.IssueOptionScheme
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
				ctx:        context.Background(),
				boardId:    1001,
				startAt:    0,
				maxResults: 50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = ACA",
					ValidateQuery: true,
					Fields:        []string{"status", "description"},
					Expand:        []string{"changelogs "},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1001/backlog?expand=changelogs+&fields=status%2Cdescription&jql=project+%3D+ACA&maxResults=50&startAt=0&validateQuery=true",
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
			name: "when the http call cannot be executed",
			args: args{
				ctx:        context.Background(),
				boardId:    1001,
				startAt:    0,
				maxResults: 50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = ACA",
					ValidateQuery: true,
					Fields:        []string{"status", "description"},
					Expand:        []string{"changelogs "},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1001/backlog?expand=changelogs+&fields=status%2Cdescription&jql=project+%3D+ACA&maxResults=50&startAt=0&validateQuery=true",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardIssuePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:        context.Background(),
				boardId:    1001,
				startAt:    0,
				maxResults: 50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = ACA",
					ValidateQuery: true,
					Fields:        []string{"status", "description"},
					Expand:        []string{"changelogs "},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1001/backlog?expand=changelogs+&fields=status%2Cdescription&jql=project+%3D+ACA&maxResults=50&startAt=0&validateQuery=true",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the board id is not provided",
			args: args{
				ctx:        context.Background(),
				startAt:    0,
				maxResults: 50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = ACA",
					ValidateQuery: true,
					Fields:        []string{"status", "description"},
					Expand:        []string{"changelogs "},
				},
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoBoardIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.Backlog(testCase.args.ctx, testCase.args.boardId, testCase.args.startAt,
				testCase.args.maxResults, testCase.args.opts)

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

func Test_BoardService_Configuration(t *testing.T) {
	type fields struct {
		c service.Client
	}

	type args struct {
		ctx     context.Context
		boardId int
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
				ctx:     context.Background(),
				boardId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1001/configuration",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardConfigurationScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:     context.Background(),
				boardId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1001/configuration",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardConfigurationScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:     context.Background(),
				boardId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1001/configuration",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the board id is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoBoardIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.Configuration(testCase.args.ctx, testCase.args.boardId)

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

func Test_BoardService_Epics(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx        context.Context
		boardId    int
		startAt    int
		maxResults int
		done       bool
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
				ctx:        context.Background(),
				boardId:    1001,
				startAt:    0,
				maxResults: 50,
				done:       false,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1001/epic?done=false&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardEpicPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:        context.Background(),
				boardId:    1001,
				startAt:    0,
				maxResults: 50,
				done:       false,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1001/epic?done=false&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardEpicPageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:        context.Background(),
				boardId:    1001,
				startAt:    0,
				maxResults: 50,
				done:       false,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1001/epic?done=false&maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the board id is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoBoardIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.Epics(testCase.args.ctx, testCase.args.boardId, testCase.args.startAt,
				testCase.args.maxResults, testCase.args.done)

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

func Test_BoardService_Delete(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx     context.Context
		boardId int
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
				ctx:     context.Background(),
				boardId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"/rest/agile/1.0/board/1001",
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
			name: "when the http call cannot be executed",
			args: args{
				ctx:     context.Background(),
				boardId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"/rest/agile/1.0/board/1001",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:     context.Background(),
				boardId: 1001,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"/rest/agile/1.0/board/1001",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the board id is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoBoardIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResponse, err := service.Delete(testCase.args.ctx, testCase.args.boardId)

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

func Test_BoardService_Filter(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx        context.Context
		filterId   int
		startAt    int
		maxResults int
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
				ctx:        context.Background(),
				filterId:   1001,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/filter/1001?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:        context.Background(),
				filterId:   1001,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/filter/1001?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardPageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:        context.Background(),
				filterId:   1001,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/filter/1001?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the filter id is not provided",
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoFilterIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.Filter(testCase.args.ctx, testCase.args.filterId, testCase.args.startAt,
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

func Test_BoardService_Gets(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx        context.Context
		opts       *model.GetBoardsOptions
		startAt    int
		maxResults int
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
				ctx:        context.Background(),
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the search options are provided",
			args: args{
				ctx:        context.Background(),
				startAt:    0,
				maxResults: 50,
				opts: &model.GetBoardsOptions{
					BoardType:               "scrum",
					BoardName:               "Sample Name",
					ProjectKeyOrID:          "DUMMY",
					AccountIDLocation:       "uuid-sample",
					ProjectIDLocation:       "uuid-sample",
					IncludePrivate:          true,
					NegateLocationFiltering: true,
					OrderBy:                 "issues",
					Expand:                  "issues",
					FilterID:                100,
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board?accountIdLocation=uuid-sample&expand=issues&filterId=100&includePrivate=true&"+
						"maxResults=50&name=Sample+Name&negateLocationFiltering=true&orderBy=issues&projectKeyOrId=DUMMY&proj"+
						"ectLocation=uuid-sample&startAt=0&type=scrum",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:        context.Background(),
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardPageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:        context.Background(),
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.Gets(testCase.args.ctx, testCase.args.opts, testCase.args.startAt,
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

func Test_BoardService_Issues(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx        context.Context
		boardId    int
		opts       *model.IssueOptionScheme
		startAt    int
		maxResults int
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
				ctx:        context.Background(),
				boardId:    1000,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/issue?maxResults=50&startAt=0",
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
			name: "when the search options are provided",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				startAt:    0,
				maxResults: 50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = DUMMY",
					ValidateQuery: true,
					Fields:        []string{"fields"},
					Expand:        []string{"orders"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/issue?expand=orders&fields=fields&jql=project+%3D+DUMMY&maxResults=50&startAt=0",
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
			name: "when the http call cannot be executed",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/issue?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardIssuePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/issue?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the board id is not provided",
			args: args{
				ctx:        context.Background(),
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoBoardIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.Issues(testCase.args.ctx, testCase.args.boardId, testCase.args.startAt,
				testCase.args.maxResults, testCase.args.opts)

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

func Test_BoardService_IssuesByEpic(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx        context.Context
		boardId    int
		epicId     int
		opts       *model.IssueOptionScheme
		startAt    int
		maxResults int
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
				ctx:        context.Background(),
				boardId:    1000,
				epicId:     102,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/epic/102/issue?maxResults=50&startAt=0",
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
			name: "when the search options are provided",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				epicId:     102,
				startAt:    0,
				maxResults: 50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = DUMMY",
					ValidateQuery: true,
					Fields:        []string{"fields"},
					Expand:        []string{"orders"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/epic/102/issue?expand=orders&fields=fields&jql=project+%3D+DUMMY&maxResults=50&startAt=0",
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
			name: "when the http call cannot be executed",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				epicId:     102,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/epic/102/issue?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardIssuePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				epicId:     102,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/epic/102/issue?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the board id is not provided",
			args: args{
				ctx:        context.Background(),
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoBoardIDError,
			wantErr: true,
		},

		{
			name: "when the epic id is not provided",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				startAt:    0,
				maxResults: 50,
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

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.IssuesByEpic(testCase.args.ctx, testCase.args.boardId, testCase.args.epicId,
				testCase.args.startAt,
				testCase.args.maxResults, testCase.args.opts)

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

func Test_BoardService_IssuesBySprint(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx        context.Context
		boardId    int
		sprintId   int
		opts       *model.IssueOptionScheme
		startAt    int
		maxResults int
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
				ctx:        context.Background(),
				boardId:    1000,
				sprintId:   102,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/sprint/102/issue?maxResults=50&startAt=0",
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
			name: "when the search options are provided",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				sprintId:   102,
				startAt:    0,
				maxResults: 50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = DUMMY",
					ValidateQuery: false,
					Fields:        []string{"fields"},
					Expand:        []string{"orders"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/sprint/102/issue?expand=orders&fields=fields&jql=project+%3D+DUMMY&maxResults=50&startAt=0&validateQuery+=false",
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
			name: "when the http call cannot be executed",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				sprintId:   102,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/sprint/102/issue?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardIssuePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				sprintId:   102,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/sprint/102/issue?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the board id is not provided",
			args: args{
				ctx:        context.Background(),
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoBoardIDError,
			wantErr: true,
		},

		{
			name: "when the sprint id is not provided",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				startAt:    0,
				maxResults: 50,
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

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.IssuesBySprint(testCase.args.ctx, testCase.args.boardId, testCase.args.sprintId,
				testCase.args.startAt,
				testCase.args.maxResults, testCase.args.opts)

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

func Test_BoardService_IssuesWithoutEpic(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx        context.Context
		boardId    int
		opts       *model.IssueOptionScheme
		startAt    int
		maxResults int
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
				ctx:        context.Background(),
				boardId:    1000,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/epic/none/issue?maxResults=50&startAt=0",
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
			name: "when the search options are provided",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				startAt:    0,
				maxResults: 50,
				opts: &model.IssueOptionScheme{
					JQL:           "project = DUMMY",
					ValidateQuery: false,
					Fields:        []string{"fields"},
					Expand:        []string{"orders"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/epic/none/issue?expand=orders&fields=fields&jql=project+%3D+DUMMY&maxResults=50&startAt=0&validateQuery=false",
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
			name: "when the http call cannot be executed",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/epic/none/issue?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.BoardIssuePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:        context.Background(),
				boardId:    1000,
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"/rest/agile/1.0/board/1000/epic/none/issue?maxResults=50&startAt=0",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the board id is not provided",
			args: args{
				ctx:        context.Background(),
				startAt:    0,
				maxResults: 50,
			},
			on: func(fields *fields) {
				fields.c = mocks.NewClient(t)
			},
			Err:     model.ErrNoBoardIDError,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service, err := NewBoardService(testCase.fields.c, "1.0")
			assert.NoError(t, err)

			gotResult, gotResponse, err := service.IssuesWithoutEpic(testCase.args.ctx, testCase.args.boardId, testCase.args.startAt,
				testCase.args.maxResults, testCase.args.opts)

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

func TestBoardService_Move(t *testing.T) {
	type fields struct {
		c       service.Client
		version string
	}
	type args struct {
		ctx     context.Context
		boardId int
		payload *model.BoardMovementPayloadScheme
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ResponseScheme
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BoardService{
				c:       tt.fields.c,
				version: tt.fields.version,
			}
			got, err := b.Move(tt.args.ctx, tt.args.boardId, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("Move() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Move() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoardService_Projects(t *testing.T) {
	type fields struct {
		c       service.Client
		version string
	}
	type args struct {
		ctx        context.Context
		boardId    int
		startAt    int
		maxResults int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.BoardProjectPageScheme
		want1   *model.ResponseScheme
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BoardService{
				c:       tt.fields.c,
				version: tt.fields.version,
			}
			got, got1, err := b.Projects(tt.args.ctx, tt.args.boardId, tt.args.startAt, tt.args.maxResults)
			if (err != nil) != tt.wantErr {
				t.Errorf("Projects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Projects() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Projects() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBoardService_Sprints(t *testing.T) {
	type fields struct {
		c       service.Client
		version string
	}
	type args struct {
		ctx        context.Context
		boardId    int
		startAt    int
		maxResults int
		states     []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.BoardSprintPageScheme
		want1   *model.ResponseScheme
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BoardService{
				c:       tt.fields.c,
				version: tt.fields.version,
			}
			got, got1, err := b.Sprints(tt.args.ctx, tt.args.boardId, tt.args.startAt, tt.args.maxResults, tt.args.states)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sprints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sprints() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Sprints() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestBoardService_Versions(t *testing.T) {
	type fields struct {
		c       service.Client
		version string
	}
	type args struct {
		ctx        context.Context
		boardID    int
		startAt    int
		maxResults int
		released   bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.BoardVersionPageScheme
		want1   *model.ResponseScheme
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BoardService{
				c:       tt.fields.c,
				version: tt.fields.version,
			}
			got, got1, err := b.Versions(tt.args.ctx, tt.args.boardID, tt.args.startAt, tt.args.maxResults, tt.args.released)
			if (err != nil) != tt.wantErr {
				t.Errorf("Versions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Versions() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Versions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewBoardService(t *testing.T) {
	type args struct {
		client  service.Client
		version string
	}
	tests := []struct {
		name    string
		args    args
		want    agile.Board
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewBoardService(tt.args.client, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBoardService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoardService() got = %v, want %v", got, tt.want)
			}
		})
	}
}
