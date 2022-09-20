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

func Test_internalKnowledgebaseImpl_Search(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx          context.Context
		query        string
		highlight    bool
		start, limit int
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
				ctx:       context.Background(),
				query:     "how to login to Jira?",
				highlight: true,
				start:     50,
				limit:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/knowledgebase/article?highlight=true&limit=50&query=how+to+login+to+Jira%3F&start=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ArticlePageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:       context.Background(),
				query:     "how to login to Jira?",
				highlight: true,
				start:     50,
				limit:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/knowledgebase/article?highlight=true&limit=50&query=how+to+login+to+Jira%3F&start=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ArticlePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:       context.Background(),
				query:     "how to login to Jira?",
				highlight: true,
				start:     50,
				limit:     50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/knowledgebase/article?highlight=true&limit=50&query=how+to+login+to+Jira%3F&start=50",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the query filter is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoKBQueryError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewKnowledgebaseService(testCase.fields.c, "latest")
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Search(testCase.args.ctx, testCase.args.query, testCase.args.highlight,
				testCase.args.start, testCase.args.limit)

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

func Test_internalKnowledgebaseImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Client
	}

	type args struct {
		ctx           context.Context
		serviceDeskID int
		query         string
		highlight     bool
		start, limit  int
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
				ctx:           context.Background(),
				serviceDeskID: 10001,
				query:         "how to login to Jira?",
				highlight:     true,
				start:         50,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/knowledgebase/article?highlight=true&limit=50&query=how+to+login+to+Jira%3F&start=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ArticlePageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name: "when the http call cannot be executed",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				query:         "how to login to Jira?",
				highlight:     true,
				start:         50,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/knowledgebase/article?highlight=true&limit=50&query=how+to+login+to+Jira%3F&start=50",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ArticlePageScheme{}).
					Return(&model.ResponseScheme{}, errors.New("client: no http response found"))

				fields.c = client
			},
			Err:     errors.New("client: no http response found"),
			wantErr: true,
		},

		{
			name: "when the request cannot be created",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
				query:         "how to login to Jira?",
				highlight:     true,
				start:         50,
				limit:         50,
			},
			on: func(fields *fields) {

				client := mocks.NewClient(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/servicedeskapi/servicedesk/10001/knowledgebase/article?highlight=true&limit=50&query=how+to+login+to+Jira%3F&start=50",
					nil).
					Return(&http.Request{}, errors.New("client: no http request created"))

				fields.c = client
			},
			Err:     errors.New("client: no http request created"),
			wantErr: true,
		},

		{
			name: "when the service desk id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoServiceDeskIDError,
		},

		{
			name: "when the query filter is not provided",
			args: args{
				ctx:           context.Background(),
				serviceDeskID: 10001,
			},
			wantErr: true,
			Err:     model.ErrNoKBQueryError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			smService, err := NewKnowledgebaseService(testCase.fields.c, "latest")
			assert.NoError(t, err)

			gotResult, gotResponse, err := smService.Gets(testCase.args.ctx, testCase.args.serviceDeskID, testCase.args.query, testCase.args.highlight,
				testCase.args.start, testCase.args.limit)

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
