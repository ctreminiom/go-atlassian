package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"net/http"
	"testing"
)

func Test_internalMetadataImpl_Get(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                    context.Context
		issueKeyOrId           string
		overrideScreenSecurity bool
		overrideEditableFlag   bool
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		want    gjson.Result
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:                    context.Background(),
				issueKeyOrId:           "DUMMY-4",
				overrideScreenSecurity: true,
				overrideEditableFlag:   false,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/DUMMY-4/editmeta?overrideEditableFlag=false&overrideScreenSecurity=true",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			want:    gjson.Result{},
			wantErr: false,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:                    context.Background(),
				issueKeyOrId:           "DUMMY-4",
				overrideScreenSecurity: true,
				overrideEditableFlag:   false,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-4/editmeta?overrideEditableFlag=false&overrideScreenSecurity=true",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			want:    gjson.Result{},
			wantErr: false,
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "3"},
			args: args{
				ctx:                    context.Background(),
				issueKeyOrId:           "",
				overrideScreenSecurity: true,
				overrideEditableFlag:   false,
			},
			want:    gjson.Result{},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrIDError,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:                    context.Background(),
				issueKeyOrId:           "DUMMY-4",
				overrideScreenSecurity: true,
				overrideEditableFlag:   false,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/DUMMY-4/editmeta?overrideEditableFlag=false&overrideScreenSecurity=true",
					"",
					nil).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			want:    gjson.Result{},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			metadataService, err := NewMetadataService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := metadataService.Get(testCase.args.ctx, testCase.args.issueKeyOrId, testCase.args.overrideScreenSecurity,
				testCase.args.overrideEditableFlag)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
				assert.Equal(t, gotResult, testCase.want)
			}
		})
	}
}

func Test_internalMetadataImpl_Create(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx  context.Context
		opts *model.IssueMetadataCreateOptions
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		want    gjson.Result
		wantErr bool
		Err     error
	}{
		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
				opts: &model.IssueMetadataCreateOptions{
					ProjectIDs:     []string{"1002"},
					ProjectKeys:    []string{"DUMMY"},
					IssueTypeIDs:   []string{"1", "2"},
					IssueTypeNames: []string{"Story", "Bug"},
					Expand:         "projects.issuetypes.fields",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/issue/createmeta?expand=projects.issuetypes.fields&issuetypeIds=1&issuetypeIds=2&issuetypeNames=Story&issuetypeNames=Bug&projectIds=1002&projectKeys=DUMMY",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			want:    gjson.Result{},
			wantErr: false,
		},

		{
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				opts: &model.IssueMetadataCreateOptions{
					ProjectIDs:     []string{"1002"},
					ProjectKeys:    []string{"DUMMY"},
					IssueTypeIDs:   []string{"1", "2"},
					IssueTypeNames: []string{"Story", "Bug"},
					Expand:         "projects.issuetypes.fields",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/createmeta?expand=projects.issuetypes.fields&issuetypeIds=1&issuetypeIds=2&issuetypeNames=Story&issuetypeNames=Bug&projectIds=1002&projectKeys=DUMMY",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
			want:    gjson.Result{},
			wantErr: false,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				opts: &model.IssueMetadataCreateOptions{
					ProjectIDs:     []string{"1002"},
					ProjectKeys:    []string{"DUMMY"},
					IssueTypeIDs:   []string{"1", "2"},
					IssueTypeNames: []string{"Story", "Bug"},
					Expand:         "projects.issuetypes.fields",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/issue/createmeta?expand=projects.issuetypes.fields&issuetypeIds=1&issuetypeIds=2&issuetypeNames=Story&issuetypeNames=Bug&projectIds=1002&projectKeys=DUMMY",
					"",
					nil).
					Return(&http.Request{}, errors.New("error"))

				fields.c = client
			},
			want:    gjson.Result{},
			wantErr: true,
			Err:     errors.New("error"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			metadataService, err := NewMetadataService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := metadataService.Create(testCase.args.ctx, testCase.args.opts)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
				assert.Equal(t, gotResult, testCase.want)
			}
		})
	}
}

func Test_NewMetadataService(t *testing.T) {

	type args struct {
		client  service.Connector
		version string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				client:  nil,
				version: "3",
			},
			wantErr: false,
		},

		{
			name: "when the version is not provided",
			args: args{
				client:  nil,
				version: "",
			},
			wantErr: true,
			err:     model.ErrNoVersionProvided,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := NewMetadataService(testCase.args.client, testCase.args.version)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.err.Error())

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, got, nil)
			}
		})
	}
}
