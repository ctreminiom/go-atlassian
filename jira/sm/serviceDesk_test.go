package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"path/filepath"
	"testing"
)

func TestServiceDeskService_Attach(t *testing.T) {

	absolutePath, err := filepath.Abs("./mocks/image.png")
	if err != nil {
		return
	}

	testCases := []struct {
		name               string
		serviceDeskID      int
		path               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AttachAttachmentToServiceDeskProjectWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			path:               absolutePath,
			mockFile:           "./mocks/attach-file-to-service-desk-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/attachTemporaryFile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "AttachAttachmentToServiceDeskProjectWhenThePathIsAFolder",
			serviceDeskID:      1,
			path:               "./mocks",
			mockFile:           "./mocks/attach-file-to-service-desk-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/attachTemporaryFile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AttachAttachmentToServiceDeskProjectWhenThePathIsEmpty",
			serviceDeskID:      1,
			path:               "",
			mockFile:           "./mocks/attach-file-to-service-desk-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/attachTemporaryFile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AttachAttachmentToServiceDeskProjectWhenThePathIsNotAbsolute",
			serviceDeskID:      1,
			path:               "./mocks/image.png",
			mockFile:           "./mocks/attach-file-to-service-desk-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/attachTemporaryFile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AttachAttachmentToServiceDeskProjectWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			path:               absolutePath,
			mockFile:           "./mocks/attach-file-to-service-desk-project.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/attachTemporaryFile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AttachAttachmentToServiceDeskProjectWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			path:               absolutePath,
			mockFile:           "./mocks/attach-file-to-service-desk-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/attachTemporaryFile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AttachAttachmentToServiceDeskProjectWhenTheEndpointIsEmpty",
			serviceDeskID:      1,
			path:               absolutePath,
			mockFile:           "./mocks/attach-file-to-service-desk-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AttachAttachmentToServiceDeskProjectWhenTheContextIsNil",
			serviceDeskID:      1,
			path:               absolutePath,
			mockFile:           "./mocks/attach-file-to-service-desk-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/attachTemporaryFile",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
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

			service := &ServiceDeskService{client: mockClient}
			gotResult, gotResponse, err := service.Attach(testCase.context, testCase.serviceDeskID, testCase.path)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				for _, attachment := range gotResult.TemporaryAttachments {
					t.Log(attachment.FileName, attachment.TemporaryAttachmentID)
				}
			}

		})
	}

}

func TestServiceDeskService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		serviceDeskID      int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetServiceDeskProjectWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			mockFile:           "./mocks/get-service-desk-project.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetServiceDeskProjectWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			mockFile:           "./mocks/get-service-desk-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetServiceDeskProjectWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			mockFile:           "./mocks/get-service-desk-project.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskProjectWhenTheContextIsNil",
			serviceDeskID:      1,
			mockFile:           "./mocks/get-service-desk-project.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskProjectWhenTheEndpointIsEmpty",
			serviceDeskID:      1,
			mockFile:           "./mocks/get-service-desk-project.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskProjectWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1",
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

			service := &ServiceDeskService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.serviceDeskID)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				t.Log(gotResult.ID, gotResult.ProjectName, gotResult.ProjectKey)
			}

		})
	}

}

func TestServiceDeskService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		start, limit       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetServiceDeskProjectsWhenTheParametersAreCorrect",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-service-desk-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetServiceDeskProjectsWhenTheRequestMethodIsIncorrect",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-service-desk-projects.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskProjectsWhenTheStatusCodeIsIncorrect",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-service-desk-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskProjectsWhenTheContextIsNil",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-service-desk-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk?limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskProjectsWhenTheEndpointIsEmpty",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-service-desk-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetServiceDeskProjectsWhenTheResponseBodyHasADifferentFormat",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk?limit=50&start=0",
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

			service := &ServiceDeskService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.start, testCase.limit)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				for _, serviceDesk := range gotResult.Values {
					t.Log(serviceDesk.ID, serviceDesk.ProjectName, serviceDesk.ProjectKey)
				}
			}

		})
	}

}
