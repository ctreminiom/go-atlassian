package internal

import (
	"context"
	"errors"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func Test_internalContentAttachmentImpl_Gets(t *testing.T) {

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                 context.Context
		contentID           string
		startAt, maxResults int
		options             *model.GetContentAttachmentsOptionsScheme
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
				contentID:  "100100101",
				startAt:    50,
				maxResults: 50,
				options: &model.GetContentAttachmentsOptionsScheme{
					Expand:    []string{"childTypes.all", "metadata.currentuser"},
					FileName:  "report_CCID",
					MediaType: "excel",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/child/attachment?expand=childTypes.all%2Cmetadata.currentuser&filename=report_CCID&limit=50&mediaType=excel&start=50",
					"", nil).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:        context.Background(),
				contentID:  "100100101",
				startAt:    50,
				maxResults: 50,
				options: &model.GetContentAttachmentsOptionsScheme{
					Expand:    []string{"childTypes.all", "metadata.currentuser"},
					FileName:  "report_CCID",
					MediaType: "excel",
				},
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodGet,
					"wiki/rest/api/content/100100101/child/attachment?expand=childTypes.all%2Cmetadata.currentuser&filename=report_CCID&limit=50&mediaType=excel&start=50",
					"", nil).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the content id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentIDError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService := NewContentAttachmentService(testCase.fields.c)

			gotResult, gotResponse, err := attachmentService.Gets(testCase.args.ctx, testCase.args.contentID, testCase.args.startAt,
				testCase.args.maxResults, testCase.args.options)

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

func Test_internalContentAttachmentImpl_CreateOrUpdate(t *testing.T) {

	absolutePathMocked, err := filepath.Abs("../../LICENSE")
	if err != nil {
		t.Fatal(err)
	}

	fileMocked, err := os.Open(absolutePathMocked)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                            context.Context
		attachmentID, status, fileName string
		file                           io.Reader
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
				ctx:          context.Background(),
				attachmentID: "3837272",
				status:       "current",
				fileName:     "LICENSE",
				file:         fileMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/3837272/child/attachment?status=current",
					mock.Anything,
					mock.Anything).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.Background(),
				attachmentID: "3837272",
				status:       "current",
				fileName:     "LICENSE",
				file:         fileMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPut,
					"wiki/rest/api/content/3837272/child/attachment?status=current",
					mock.Anything,
					mock.Anything).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the attachment id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentAttachmentIDError,
		},

		{
			name: "when the file name is not provided",
			args: args{
				ctx:          context.Background(),
				attachmentID: "3837272",
			},
			wantErr: true,
			Err:     model.ErrNoContentAttachmentNameError,
		},

		{
			name: "when the file reader is not provided",
			args: args{
				ctx:          context.Background(),
				attachmentID: "3837272",
				fileName:     "LICENSE",
			},
			wantErr: true,
			Err:     model.ErrNoContentReaderError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService := NewContentAttachmentService(testCase.fields.c)

			gotResult, gotResponse, err := attachmentService.CreateOrUpdate(testCase.args.ctx, testCase.args.attachmentID,
				testCase.args.status, testCase.args.fileName, testCase.args.file)

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

func Test_internalContentAttachmentImpl_Create(t *testing.T) {

	absolutePathMocked, err := filepath.Abs("../../LICENSE")
	if err != nil {
		t.Fatal(err)
	}

	fileMocked, err := os.Open(absolutePathMocked)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		c service.Connector
	}

	type args struct {
		ctx                            context.Context
		attachmentID, status, fileName string
		file                           io.Reader
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
				ctx:          context.Background(),
				attachmentID: "3837272",
				status:       "current",
				fileName:     "LICENSE",
				file:         fileMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/3837272/child/attachment?status=current",
					mock.Anything,
					mock.Anything).
					Return(&http.Request{}, nil)

				client.On("Call",
					&http.Request{},
					&model.ContentPageScheme{}).
					Return(&model.ResponseScheme{}, nil)

				fields.c = client

			},
		},

		{
			name: "when the http request cannot be created",
			args: args{
				ctx:          context.Background(),
				attachmentID: "3837272",
				status:       "current",
				fileName:     "LICENSE",
				file:         fileMocked,
			},
			on: func(fields *fields) {

				client := mocks.NewConnector(t)

				client.On("NewRequest",
					context.Background(),
					http.MethodPost,
					"wiki/rest/api/content/3837272/child/attachment?status=current",
					mock.Anything,
					mock.Anything).
					Return(&http.Request{}, errors.New("error, unable to create the http request"))

				fields.c = client

			},
			wantErr: true,
			Err:     errors.New("error, unable to create the http request"),
		},

		{
			name: "when the attachment id is not provided",
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
			Err:     model.ErrNoContentAttachmentIDError,
		},

		{
			name: "when the file name is not provided",
			args: args{
				ctx:          context.Background(),
				attachmentID: "3837272",
			},
			wantErr: true,
			Err:     model.ErrNoContentAttachmentNameError,
		},

		{
			name: "when the file reader is not provided",
			args: args{
				ctx:          context.Background(),
				attachmentID: "3837272",
				fileName:     "LICENSE",
			},
			wantErr: true,
			Err:     model.ErrNoContentReaderError,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			attachmentService := NewContentAttachmentService(testCase.fields.c)

			gotResult, gotResponse, err := attachmentService.Create(testCase.args.ctx, testCase.args.attachmentID,
				testCase.args.status, testCase.args.fileName, testCase.args.file)

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
