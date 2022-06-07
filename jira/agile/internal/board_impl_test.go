package internal

import (
	"bytes"
	"context"
	"errors"
	"github.com/ctreminiom/go-atlassian/jira/agile/internal/mocks"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

/*

func TestProductService_Insert(t *testing.T) {

	ctx := context.Background()
	id := 1

	client := mocks.NewClient(t)

	client.On("NewJsonRequest",
		ctx,
		http.MethodGet,
		"/rest/agile/1.0/board/1",
		nil).
		Return(&http.Request{}, nil)

	client.On("Call",
		&http.Request{},
		&model.BoardScheme{}).
		Return(&model.ResponseScheme{}, nil)

	service, err := newBoardService(client, "1.0")
	assert.NoError(t, err)

	gotResult, gotResponse, err := service.Get(ctx, id)

	assert.NoError(t, err)
	assert.NotEqual(t, gotResponse, nil)
	assert.NotEqual(t, gotResult, nil)
}

*/

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

func TestBoardService_Create(t *testing.T) {

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
