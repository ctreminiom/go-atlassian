package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalBoardBacklogImpl_Move(t *testing.T) {

	payloadMocked := map[string]interface{}{"issues": []string{"DUMMY-1", "12323"}}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx    context.Context
		issues []string
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
				ctx:    context.Background(),
				issues: []string{"DUMMY-1", "12323"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/backlog/issue",
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
				ctx:    context.Background(),
				issues: []string{"DUMMY-1", "12323"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/backlog/issue",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, fmt.Errorf("agile: %w", model.ErrNotFound))

				fields.c = client
			},
			Err:     model.ErrNotFound,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:    context.Background(),
				issues: []string{"DUMMY-1", "12323"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/backlog/issue",
					"",
					payloadMocked).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service := NewBoardBacklogService(testCase.fields.c, "1.0")

			gotResponse, err := service.Move(testCase.args.ctx, testCase.args.issues)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}
		})
	}
}

func Test_internalBoardBacklogImpl_MoveTo(t *testing.T) {

	payloadMocked := &model.BoardBacklogPayloadScheme{
		Issues:            []string{"PR-1", "10001", "PR-3"},
		RankBeforeIssue:   "PR-4",
		RankAfterIssue:    "",
		RankCustomFieldID: 10521,
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		boardID int
		payload *model.BoardBacklogPayloadScheme
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
				boardID: 56,
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/backlog/56/issue",
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
				ctx:     context.Background(),
				boardID: 56,
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/backlog/56/issue",
					"",
					payloadMocked).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, model.ErrNoExecHttpCall)

				fields.c = client
			},
			Err:     model.ErrNoExecHttpCall,
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:     context.Background(),
				boardID: 56,
				payload: payloadMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/agile/1.0/backlog/56/issue",
					"",
					payloadMocked).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			Err:     model.ErrCreateHttpReq,
			wantErr: true,
		},

		{
			name: "when the board id is not provided",
			args: args{
				ctx: context.Background(),
			},
			Err:     model.ErrNoBoardID,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			service := NewBoardBacklogService(testCase.fields.c, "1.0")

			gotResponse, err := service.MoveTo(testCase.args.ctx, testCase.args.boardID, testCase.args.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				// the first if statement is to handle wrapped errors from url and json packages for more accurate comparison
				var urlErr *url.Error
				var jsonErr *json.SyntaxError
				if errors.As(err, &urlErr) || errors.As(err, &jsonErr) {
					assert.Contains(t, err.Error(), testCase.Err.Error())
				} else {
					assert.True(t, errors.Is(err, testCase.Err), "expected error: %v, got: %v", testCase.Err, err)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}
		})
	}
}
