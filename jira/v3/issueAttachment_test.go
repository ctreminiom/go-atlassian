package v3

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

func TestAttachmentService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		attachmentID       string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteAttachmentWhenTheAttachmentIDIsCorrect",
			attachmentID:       "10006",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteAttachmentWhenTheAttachmentIDIsNotSet",
			attachmentID:       "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteAttachmentWhenTheAttachmentIDIsIncorrect",
			attachmentID:       "10007",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteAttachmentWhenTheContextIsNil",
			attachmentID:       "10006",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteAttachmentWhenTheResponseCodeIsDifferent",
			attachmentID:       "10006",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantErr:            true,
		},
		{
			name:               "DeleteAttachmentWhenTheRequestMethodIsDifferent",
			attachmentID:       "10006",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

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

			service := &AttachmentService{client: mockClient}

			gotResponse, err := service.Delete(testCase.context, testCase.attachmentID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func TestAttachmentService_Human(t *testing.T) {
	testCases := []struct {
		name               string
		attachmentID       string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetHumanReadableAttachmentWhenTheAttachmentIDIsCorrect",
			mockFile:           "./mocks/get-attachment-human-view.json",
			attachmentID:       "10006",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006/expand/human",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetHumanReadableAttachmentWhenTheAttachmentIDIsNotSet",
			mockFile:           "./mocks/get-attachment-human-view.json",
			attachmentID:       "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006/expand/human",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:         "GetHumanReadableAttachmentWhenTheAttachmentIDIsIncorrect",
			attachmentID: "10007",
			mockFile:     "./mocks/get-attachment-human-view.json",

			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006/expand/human",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:         "GetHumanReadableAttachmentWhenTheAttachmentIDIsEmpty",
			attachmentID: "",
			mockFile:     "./mocks/get-attachment-human-view.json",

			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006/expand/human",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:         "GetHumanReadableAttachmentWhenTheAttachmentIDHasSpecialCharacters",
			attachmentID: "((*^%%**",
			mockFile:     "./mocks/get-attachment-human-view.json",

			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006/expand/human",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:         "GetHumanReadableAttachmentWhenTheContextIsNil",
			attachmentID: "10006",
			mockFile:     "./mocks/get-attachment-human-view.json",

			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006/expand/human",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:         "GetHumanReadableAttachmentWhenTheRequestMethodIsDifferent",
			attachmentID: "10006",
			mockFile:     "./mocks/get-attachment-human-view.json",

			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/attachment/10006/expand/human",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetHumanReadableAttachmentWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "./mocks/empty_json.json",
			attachmentID:       "10006",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006/expand/human",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

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

			service := &AttachmentService{client: mockClient}

			getResult, gotResponse, err := service.Human(testCase.context, testCase.attachmentID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, getResult, nil)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func TestAttachmentService_Metadata(t *testing.T) {
	testCases := []struct {
		name               string
		attachmentID       string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetAttachmentMetadataWhenTheAttachmentIDIsCorrect",
			mockFile:           "./mocks/get-attachment-metadata.json",
			attachmentID:       "10006",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetAttachmentMetadataWhenTheAttachmentIDIsNotSet",
			mockFile:           "./mocks/get-attachment-metadata.json",
			attachmentID:       "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetAttachmentMetadataWhenTheAttachmentIDIsIncorrect",
			mockFile:           "./mocks/get-attachment-metadata.json",
			attachmentID:       "10007",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetAttachmentMetadataWhenTheAttachmentIDIsEmpty",
			mockFile:           "./mocks/get-attachment-metadata.json",
			attachmentID:       "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetAttachmentMetadataWhenTheContextIsNil",
			mockFile:           "./mocks/get-attachment-metadata.json",
			attachmentID:       "10006",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetAttachmentMetadataWhenTheRequestMethodIsDifferent",
			mockFile:           "./mocks/get-attachment-metadata.json",
			attachmentID:       "10006",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetAttachmentMetadataWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "./mocks/empty_json.json",
			attachmentID:       "10006",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/10006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

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

			service := &AttachmentService{client: mockClient}

			getResult, gotResponse, err := service.Metadata(testCase.context, testCase.attachmentID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, getResult, nil)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func TestAttachmentService_Settings(t *testing.T) {
	testCases := []struct {
		name               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetAttachmentSetting",
			mockFile:           "./mocks/get-attachment-settings.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/meta",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetAttachmentSettingWhenTheRequestMethodIsDifferent",
			mockFile:           "./mocks/get-attachment-settings.json",
			wantHTTPMethod:     http.MethodConnect,
			endpoint:           "/rest/api/3/attachment/meta",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetAttachmentSettingWhenTheContextIsNil",
			mockFile:           "./mocks/get-attachment-settings.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/meta",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetAttachmentSettingWhenTheResponseBodyLengthIsZero",
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/meta",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetAttachmentSettingWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/attachment/meta",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

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

			service := &AttachmentService{client: mockClient}

			getResult, gotResponse, err := service.Settings(testCase.context)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, getResult, nil)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func TestAttachmentService_Add(t *testing.T) {

	absolutePathMocked, err := filepath.Abs("./mocks/image.png")
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
		name               string
		mockFile           string
		issueKeyOrID       string
		fileName           string
		file               io.Reader
		context            context.Context
		wantHTTPMethod     string
		endpoint           string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddAttachmentWhenThePathIsAbsolute",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			fileName:           "image.png",
			file:               fileMocked,
			context:            context.Background(),
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "AddAttachmentWhenTheIssueKeyOrIDIsNotProvided",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "",
			fileName:           "image.png",
			file:               fileMocked,
			context:            context.Background(),
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddAttachmentWhenTheFileNameIsNotProvided",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			fileName:           "",
			file:               fileMocked,
			context:            context.Background(),
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddAttachmentWhenTheFileReaderIsNotProvided",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			fileName:           "image.png",
			file:               nil,
			context:            context.Background(),
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddAttachmentWhenTheReaderIsBlocked",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			fileName:           "image.png",
			file:               fileErrorMocked,
			context:            context.Background(),
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddAttachmentWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			fileName:           "image.png",
			file:               fileMocked,
			context:            context.Background(),
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddAttachmentWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			fileName:           "image.png",
			file:               fileMocked,
			context:            context.Background(),
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AddAttachmentWhenTheContextIsNotProvided",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			fileName:           "image.png",
			file:               fileMocked,
			context:            nil,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

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

			service := &AttachmentService{client: mockClient}

			getResult, gotResponse, err := service.Add(testCase.context, testCase.issueKeyOrID, testCase.fileName, testCase.file)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, getResult, nil)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}
