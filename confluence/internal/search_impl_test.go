package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_internalSearchImpl_Content(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx     context.Context
		cql     string
		options *model.SearchContentOptions
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
				cql: "type=page",
				options: &model.SearchContentOptions{
					Context:                  "spaceKey",
					Cursor:                   "raNDoMsTRiNg",
					Next:                     true,
					Prev:                     true,
					Limit:                    20,
					Start:                    10,
					IncludeArchivedSpaces:    true,
					ExcludeCurrentSpaces:     true,
					SitePermissionTypeFilter: "externalCollaborator",
					Excerpt:                  "indexed",
					Expand:                   []string{"space"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/search?cql=type%3Dpage&cqlcontext=spaceKey&cursor=raNDoMsTRiNg&excerpt=indexed&excludeCurrentSpaces=true&expand=space&includeArchivedSpaces=true&limit=20&next=true&prev=true&sitePermissionTypeFilter=externalCollaborator&start=10",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SearchPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx: context.Background(),
				cql: "type=page",
				options: &model.SearchContentOptions{
					Context:                  "spaceKey",
					Cursor:                   "raNDoMsTRiNg",
					Next:                     true,
					Prev:                     true,
					Limit:                    20,
					Start:                    10,
					IncludeArchivedSpaces:    true,
					ExcludeCurrentSpaces:     true,
					SitePermissionTypeFilter: "externalCollaborator",
					Excerpt:                  "indexed",
					Expand:                   []string{"space"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/search?cql=type%3Dpage&cqlcontext=spaceKey&cursor=raNDoMsTRiNg&excerpt=indexed&excludeCurrentSpaces=true&expand=space&includeArchivedSpaces=true&limit=20&next=true&prev=true&sitePermissionTypeFilter=externalCollaborator&start=10",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the cql is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoCQL,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSearchService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Content(testCase.args.ctx, testCase.args.cql, testCase.args.options)

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

func Test_internalSearchImpl_Users(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx    context.Context
		cql    string
		start  int
		limit  int
		expand []string
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
				cql:    "type=page",
				start:  20,
				limit:  50,
				expand: []string{"operations", "personalSpace"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/search/user?cql=type%3Dpage&expand=operations%2CpersonalSpace&limit=50&start=20",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.SearchPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:    context.Background(),
				cql:    "type=page",
				start:  20,
				limit:  50,
				expand: []string{"operations", "personalSpace"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/search/user?cql=type%3Dpage&expand=operations%2CpersonalSpace&limit=50&start=20",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the cql is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoCQL,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			newService := NewSearchService(testCase.fields.c)

			gotResult, gotResponse, err := newService.Users(testCase.args.ctx, testCase.args.cql, testCase.args.start,
				testCase.args.limit, testCase.args.expand)

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
