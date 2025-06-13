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

func Test_internalIssueArchivalImpl_Preserve(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}
	type args struct {
		ctx            context.Context
		issueIdsOrKeys []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "happy path - when the issue list is archived successfully",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueIdsOrKeys: []string{"KP-1"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/archive", "", map[string]interface{}{"issueIdsOrKeys": []string{"KP-1"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueArchivalSyncResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "fail path - when the issue list is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueIdsOrKeys: nil,
			},
			wantErr: true,
			Err:     model.ErrNoIssuesSlice,
		},

		{
			name:   "fail path - when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueIdsOrKeys: []string{"KP-1"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/archive", "", map[string]interface{}{"issueIdsOrKeys": []string{"KP-1"}}).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))
				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.on != nil {
				tt.on(&tt.fields)
			}

			archiveService := NewIssueArchivalService(tt.fields.c, tt.fields.version)

			gotResult, gotResponse, err := archiveService.internalClient.Preserve(tt.args.ctx, tt.args.issueIdsOrKeys)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.Err.Error())
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, gotResponse)
			assert.NotNil(t, gotResult)
		})
	}
}

func Test_internalIssueArchivalImpl_PreserveByJQL(t *testing.T) {
	type fields struct {
		c       service.Connector
		version string
	}
	type args struct {
		ctx context.Context
		jql string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "happy path - when the issues are archived successfully",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				jql: "project = KP",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/archive", "", map[string]interface{}{"jql": "project = KP"}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					nil).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},
		{
			name:   "fail path - when the jql is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				jql: "",
			},
			wantErr: true,
			Err:     model.ErrNoJQL,
		},
		{
			name:   "fail path - when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				jql: "project = KP",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/archive", "", map[string]interface{}{"jql": "project = KP"}).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))
				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.on != nil {
				tt.on(&tt.fields)
			}

			archiveService := NewIssueArchivalService(tt.fields.c, tt.fields.version)

			gotResult, gotResponse, err := archiveService.internalClient.PreserveByJQL(tt.args.ctx, tt.args.jql)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.Err.Error())
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, gotResponse)
			assert.NotNil(t, gotResult)
		})
	}
}

func Test_internalIssueArchivalImpl_Restore(t *testing.T) {
	type fields struct {
		c       service.Connector
		version string
	}
	type args struct {
		ctx            context.Context
		issueIDsOrKeys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "happy path - when the issues are restored successfully",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueIDsOrKeys: []string{"KP-1"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/unarchive", "", map[string]interface{}{"issueIdsOrKeys": []string{"KP-1"}}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueArchivalSyncResponseScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},
		{
			name:   "fail path - when the issue list is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueIDsOrKeys: nil,
			},
			wantErr: true,
			Err:     model.ErrNoIssuesSlice,
		},
		{
			name:   "fail path - when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:            context.Background(),
				issueIDsOrKeys: []string{"KP-1"},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issue/unarchive", "", map[string]interface{}{"issueIdsOrKeys": []string{"KP-1"}}).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))
				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.on != nil {
				tt.on(&tt.fields)
			}

			archiveService := NewIssueArchivalService(tt.fields.c, tt.fields.version)

			gotResult, gotResponse, err := archiveService.internalClient.Restore(tt.args.ctx, tt.args.issueIDsOrKeys)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.Err.Error())
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, gotResponse)
			assert.NotNil(t, gotResult)

		})
	}
}

func Test_internalIssueArchivalImpl_Export(t *testing.T) {
	type fields struct {
		c       service.Connector
		version string
	}
	type args struct {
		ctx     context.Context
		payload *model.IssueArchivalExportPayloadScheme
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		wantErr bool
		Err     error
	}{
		{
			name:   "happy path - when the issues are exported successfully",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
				payload: &model.IssueArchivalExportPayloadScheme{

					ArchivedBy: []string{
						"uuid-sample",
						"uuid-sample",
					},
					ArchivedDateRange: &model.DateRangeFilterRequestScheme{
						DateAfter:  "2023-01-01",
						DateBefore: "2023-01-12",
					},
					IssueTypes: []string{"Bug", "Story"},
					Projects:   []string{"WORK"},
					Reporters:  []string{"uuid-sample"},
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issues/archive/export", "", &model.IssueArchivalExportPayloadScheme{

						ArchivedBy: []string{
							"uuid-sample",
							"uuid-sample",
						},
						ArchivedDateRange: &model.DateRangeFilterRequestScheme{
							DateAfter:  "2023-01-01",
							DateBefore: "2023-01-12",
						},
						IssueTypes: []string{"Bug", "Story"},
						Projects:   []string{"WORK"},
						Reporters:  []string{"uuid-sample"},
					}).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueArchiveExportResultScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},
		{
			name:   "fail path - when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:     context.Background(),
				payload: &model.IssueArchivalExportPayloadScheme{},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"rest/api/2/issues/archive/export", "", &model.IssueArchivalExportPayloadScheme{}).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))
				fields.c = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.on != nil {
				tt.on(&tt.fields)
			}

			archiveService := NewIssueArchivalService(tt.fields.c, tt.fields.version)

			gotResult, gotResponse, err := archiveService.internalClient.Export(tt.args.ctx, tt.args.payload)

			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.Err.Error())
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, gotResponse)
			assert.NotNil(t, gotResult)
		})
	}
}
