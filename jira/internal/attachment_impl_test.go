package internal

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func Test_internalIssueAttachmentServiceImpl_Settings(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx context.Context
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/attachment/meta",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.AttachmentSettingScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/attachment/meta",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.AttachmentSettingScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx: context.Background(),
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/attachment/meta",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService, err := NewIssueAttachmentService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := attachmentService.Settings(testCase.args.ctx)

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
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalIssueAttachmentServiceImpl_Metadata(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		attachmentID string
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/attachment/1110",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueAttachmentMetadataScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/attachment/1110",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueAttachmentMetadataScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name:   "when the attachment id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoAttachmentID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/attachment/1110",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService, err := NewIssueAttachmentService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := attachmentService.Metadata(testCase.args.ctx, testCase.args.attachmentID)

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
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalIssueAttachmentServiceImpl_Human(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		attachmentID string
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/attachment/1110/expand/human",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueAttachmentHumanMetadataScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/attachment/1110/expand/human",
					"",
					nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.IssueAttachmentHumanMetadataScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name:   "when the attachment id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoAttachmentID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/attachment/1110/expand/human",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService, err := NewIssueAttachmentService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := attachmentService.Human(testCase.args.ctx, testCase.args.attachmentID)

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
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalIssueAttachmentServiceImpl_Delete(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		attachmentID string
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/attachment/1110",
					"",
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/3/attachment/1110",
					"",
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
			name:   "when the attachment id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoAttachmentID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodDelete,
					"rest/api/2/attachment/1110",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService, err := NewIssueAttachmentService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := attachmentService.Delete(testCase.args.ctx, testCase.args.attachmentID)

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

func Test_internalIssueAttachmentServiceImpl_Add(t *testing.T) {

	absolutePathMocked, err := filepath.Abs("../../LICENSE")
	if err != nil {
		t.Fatal(err)
	}

	fileMocked, err := os.Open(absolutePathMocked)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx                    context.Context
		issueKeyOrID, fileName string
		file                   io.Reader
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				fileName:     "LICENSE",
				file:         fileMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/attachments",
					mock.Anything,
					mock.Anything).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				fileName:     "LICENSE",
				file:         fileMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/3/issue/DUMMY-1/attachments",
					mock.Anything,
					mock.Anything).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					mock.Anything).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client
			},
		},

		{
			name:   "when the issue key or id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "",
				fileName:     "LICENSE",
				file:         fileMocked,
			},
			wantErr: true,
			Err:     model.ErrNoIssueKeyOrID,
		},

		{
			name:   "when the file name is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				fileName:     "",
				file:         fileMocked,
			},
			wantErr: true,
			Err:     model.ErrNoAttachmentName,
		},

		{
			name:   "when the field reader is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				fileName:     "LICENSE",
				file:         nil,
			},
			wantErr: true,
			Err:     model.ErrNoReader,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				issueKeyOrID: "DUMMY-1",
				fileName:     "LICENSE",
				file:         fileMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"rest/api/2/issue/DUMMY-1/attachments",
					mock.Anything,
					mock.Anything).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client
			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService, err := NewIssueAttachmentService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResult, gotResponse, err := attachmentService.Add(testCase.args.ctx, testCase.args.issueKeyOrID,
				testCase.args.fileName, testCase.args.file)

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
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func Test_internalIssueAttachmentServiceImpl_Download(t *testing.T) {

	type fields struct {
		c       service.Connector
		version string
	}

	type args struct {
		ctx          context.Context
		attachmentID string
		redirect     bool
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
			name:   "when the api version is v2",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
				redirect:     false,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/attachment/content/1110?redirect=false",
					"",
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
			name:   "when the api version is v3",
			fields: fields{version: "3"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
				redirect:     true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/3/attachment/content/1110",
					"",
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
			name:   "when the attachment id is not provided",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "",
			},
			on: func(fields *fields) {
				fields.c = mocks.NewConnector(t)
			},
			wantErr: true,
			Err:     model.ErrNoAttachmentID,
		},

		{
			name:   "when the http request cannot be created",
			fields: fields{version: "2"},
			args: args{
				ctx:          context.Background(),
				attachmentID: "1110",
				redirect:     true,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"rest/api/2/attachment/content/1110",
					"",
					nil).
					Return(&http.Request{}, model.ErrCreateHttpReq)

				fields.c = client

			},
			wantErr: true,
			Err:     model.ErrCreateHttpReq,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService, err := NewIssueAttachmentService(testCase.fields.c, testCase.fields.version)
			assert.NoError(t, err)

			gotResponse, err := attachmentService.Download(testCase.args.ctx, testCase.args.attachmentID, testCase.args.redirect)

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

func TestNewIssueAttachmentService(t *testing.T) {

	type args struct {
		client  service.Connector
		version string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		Err     error
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
			Err:     model.ErrNoVersionProvided,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := NewIssueAttachmentService(testCase.args.client, testCase.args.version)

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
				assert.NotEqual(t, got, nil)
			}
		})
	}
}
