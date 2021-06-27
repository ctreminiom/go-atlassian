package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

func TestContentAttachmentService_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		contentID           string
		startAt, maxResults int
		options             *GetContentAttachmentsOptionsScheme
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:       "GetContentAttachmentsWhenTheParametersAreCorrect",
			contentID:  "5949392",
			startAt:    0,
			maxResults: 50,
			options: &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "GetContentAttachmentsWhenTheContentIDIsNotProvided",
			contentID:  "",
			startAt:    0,
			maxResults: 50,
			options: &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:       "GetContentAttachmentsWhenTheContextIsNotProvided",
			contentID:  "5949392",
			startAt:    0,
			maxResults: 50,
			options: &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:       "GetContentAttachmentsWhenTheRequestMethodIsIncorrect",
			contentID:  "5949392",
			startAt:    0,
			maxResults: 50,
			options: &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentAttachmentsWhenTheOptionsAreNotProvided",
			contentID:          "5949392",
			startAt:            0,
			maxResults:         50,
			options:            nil,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "GetContentAttachmentsWhenTheStatusCodeIsIncorrect",
			contentID:  "5949392",
			startAt:    0,
			maxResults: 50,
			options: &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:       "GetContentAttachmentsWhenTheResponseBodyIsEmpty",
			contentID:  "5949392",
			startAt:    0,
			maxResults: 50,
			options: &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			service := &ContentAttachmentService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.contentID, testCase.startAt, testCase.maxResults, testCase.options)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}

		})

	}

}

func TestContentAttachmentService_CreateOrUpdate(t *testing.T) {

	absolutePathMocked, err := filepath.Abs("./mocks/mock.png")
	if err != nil {
		t.Fatal(err)
	}

	fileMocked, err := os.Open(absolutePathMocked)
	if err != nil {
		t.Fatal(err)
	}

	fileErrorMocked, err := os.OpenFile(absolutePathMocked, os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name                 string
		attachmentID, status string
		fileName             string
		file                 io.Reader
		mockFile             string
		wantHTTPMethod       string
		endpoint             string
		context              context.Context
		wantHTTPCodeReturn   int
		wantErr              bool
	}{
		{
			name:               "CreateOrUpdateAttachmentWhenTheParametersAreCorrect",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "CreateOrUpdateAttachmentWhenTheReaderIsBlocked",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileErrorMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateOrUpdateAttachmentWhenTheAttachmentIDIsNotProvided",
			attachmentID:       "",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateOrUpdateAttachmentWhenTheFileNameIsNotProvided",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateOrUpdateAttachmentWhenTheFileReaderIsNotProvided",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               nil,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateOrUpdateAttachmentWhenTheContextIsNotProvided",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateOrUpdateAttachmentWhenTheRequestMethodIsIncorrect",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateOrUpdateAttachmentWhenTheStatusCodeIsIncorrect",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantErr:            true,
		},

		{
			name:               "CreateOrUpdateAttachmentWhenTheResponseBodyIsEmpty",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			service := &ContentAttachmentService{client: mockClient}

			gotResult, gotResponse, err := service.CreateOrUpdate(testCase.context, testCase.attachmentID, testCase.status, testCase.fileName, testCase.file)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}

		})

	}

}

func TestContentAttachmentService_Create(t *testing.T) {

	absolutePathMocked, err := filepath.Abs("./mocks/mock.png")
	if err != nil {
		t.Fatal(err)
	}

	fileMocked, err := os.Open(absolutePathMocked)
	if err != nil {
		t.Fatal(err)
	}

	fileErrorMocked, err := os.OpenFile(absolutePathMocked, os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name                 string
		attachmentID, status string
		fileName             string
		file                 io.Reader
		mockFile             string
		wantHTTPMethod       string
		endpoint             string
		context              context.Context
		wantHTTPCodeReturn   int
		wantErr              bool
	}{
		{
			name:               "CreateAttachmentWhenTheParametersAreCorrect",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "CreateAttachmentWhenTheReaderIsBlocked",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileErrorMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateAttachmentWhenTheAttachmentIDIsNotProvided",
			attachmentID:       "",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateAttachmentWhenTheFileNameIsNotProvided",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateAttachmentWhenTheFileReaderIsNotProvided",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               nil,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateAttachmentWhenTheContextIsNotProvided",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateAttachmentWhenTheRequestMethodIsIncorrect",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodConnect,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateAttachmentWhenTheStatusCodeIsIncorrect",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantErr:            true,
		},

		{
			name:               "CreateAttachmentWhenTheResponseBodyIsEmpty",
			attachmentID:       "490949439393",
			status:             "current",
			fileName:           "mock.png",
			file:               fileMocked,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/490949439393/child/attachment?status=current",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			service := &ContentAttachmentService{client: mockClient}

			gotResult, gotResponse, err := service.Create(testCase.context, testCase.attachmentID, testCase.status, testCase.fileName, testCase.file)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}

		})

	}

}
