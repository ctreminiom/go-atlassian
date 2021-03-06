package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
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
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
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
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
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
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
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
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func TestAttachmentService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		issueKeyOrID       string
		path               string
		wannaAbsolutePath  bool
		wantHTTPMethod     string
		endpoint           string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddAttachmentWhenThePathIsAbsolute",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			path:               "./mocks/image.png",
			wannaAbsolutePath:  true,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "AddAttachmentWhenThePathIsNotSet",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			path:               "",
			wannaAbsolutePath:  true,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddAttachmentWhenTheIssueKeyOrIDIsNotSet",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "",
			path:               "./mocks/image.png",
			wannaAbsolutePath:  true,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddAttachmentWhenTheFilePathDoesNotExists",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			path:               "./mocks",
			wannaAbsolutePath:  true,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "AddAttachmentWhenTheFilePathIsEmpty",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			path:               "",
			wannaAbsolutePath:  true,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "AddAttachmentWhenTheFilePathIsAFolder",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			path:               "./mocks/",
			wannaAbsolutePath:  true,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "AddAttachmentWhenThePathIsNotAbsolute",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			path:               "./mocks/image.png",
			wannaAbsolutePath:  false,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "AddAttachmentWhenTheRequestMethodIsInvalid",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			path:               "./mocks/image.png",
			wannaAbsolutePath:  true,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "AddAttachmentWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-attachments.json",
			issueKeyOrID:       "KP-1",
			path:               "./mocks/image.png",
			wannaAbsolutePath:  true,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/KP-1/attachments",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "AddAttachmentWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "./mocks/empty_json.json",
			issueKeyOrID:       "KP-1",
			path:               "./mocks/image.png",
			wannaAbsolutePath:  true,
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

			var pathToUse string
			if testCase.wannaAbsolutePath {

				pathToUse, err = filepath.Abs(testCase.path)
				if err != nil {
					t.Fatal(err)
				}

			} else {
				pathToUse = testCase.path
			}

			getResult, gotResponse, err := service.Add(testCase.issueKeyOrID, pathToUse)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
			}

		})
	}
}
