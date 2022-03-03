package v2

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_ProjectPropertyService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the parameters are correct",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../mocks/get-project-properties.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/properties",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:               "when project key or id is not provided",
			projectKeyOrID:     "",
			mockFile:           "../mocks/get-project-properties.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/properties",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "jira: no project id set",
		},

		{
			name:               "when the context is not provided",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../mocks/get-project-properties.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/properties",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response body is empty",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/properties",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
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

			i := &ProjectPropertyService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.projectKeyOrID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
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
			}
		})

	}

}

func Test_ProjectPropertyService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		propertyKey        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the parameters are correct",
			projectKeyOrID:     "DUMMY",
			propertyKey:        "property-key-sample",
			mockFile:           "../mocks/get-project-property.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:               "when the property key is not provided",
			projectKeyOrID:     "DUMMY",
			propertyKey:        "",
			mockFile:           "../mocks/get-project-property.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "jira: no property key set",
		},

		{
			name:               "when project key or id is not provided",
			projectKeyOrID:     "",
			mockFile:           "../mocks/get-project-property.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "jira: no project id set",
		},

		{
			name:               "when the context is not provided",
			projectKeyOrID:     "DUMMY",
			propertyKey:        "property-key-sample",
			mockFile:           "../mocks/get-project-property.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response body is empty",
			projectKeyOrID:     "DUMMY",
			propertyKey:        "property-key-sample",
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
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

			i := &ProjectPropertyService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.projectKeyOrID, testCase.propertyKey)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
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
			}
		})

	}

}

func Test_ProjectPropertyService_Set(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		propertyKey        string
		payload            interface{}
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:           "when the parameters are correct",
			projectKeyOrID: "DUMMY",
			propertyKey:    "property-key-sample",
			payload: map[string]interface{}{
				"system.conversation.id": "b1bf38be-5e94-4b40-a3b8-9278735ee1e6",
				"system.support.time":    "1m",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:           "when the response code is incorrect",
			projectKeyOrID: "DUMMY",
			propertyKey:    "property-key-sample",
			payload: map[string]interface{}{
				"system.conversation.id": "b1bf38be-5e94-4b40-a3b8-9278735ee1e6",
				"system.support.time":    "1m",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the payload is not provided",
			projectKeyOrID:     "DUMMY",
			propertyKey:        "property-key-sample",
			payload:            nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "failed to parse the interface pointer, please provide a valid one",
		},

		{
			name:           "when the property key is not provided",
			projectKeyOrID: "DUMMY",
			propertyKey:    "",
			payload: map[string]interface{}{
				"system.conversation.id": "b1bf38be-5e94-4b40-a3b8-9278735ee1e6",
				"system.support.time":    "1m",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "jira: no property key set",
		},

		{
			name:           "when project key or id is not provided",
			projectKeyOrID: "",
			payload: map[string]interface{}{
				"system.conversation.id": "b1bf38be-5e94-4b40-a3b8-9278735ee1e6",
				"system.support.time":    "1m",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "jira: no project id set",
		},

		{
			name:           "when the context is not provided",
			projectKeyOrID: "DUMMY",
			propertyKey:    "property-key-sample",
			payload: map[string]interface{}{
				"system.conversation.id": "b1bf38be-5e94-4b40-a3b8-9278735ee1e6",
				"system.support.time":    "1m",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
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

			i := &ProjectPropertyService{client: mockClient}

			gotResponse, err := i.Set(testCase.context, testCase.projectKeyOrID, testCase.propertyKey, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

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
			}
		})

	}

}

func Test_ProjectPropertyService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		propertyKey        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the parameters are correct",
			projectKeyOrID:     "DUMMY",
			propertyKey:        "property-key-sample",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the property key is not provided",
			projectKeyOrID:     "DUMMY",
			propertyKey:        "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "jira: no property key set",
		},

		{
			name:               "when project key or id is not provided",
			projectKeyOrID:     "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "jira: no project id set",
		},

		{
			name:               "when the context is not provided",
			projectKeyOrID:     "DUMMY",
			propertyKey:        "property-key-sample",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/project/DUMMY/properties/property-key-sample",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
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

			i := &ProjectPropertyService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.projectKeyOrID, testCase.propertyKey)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

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
			}
		})

	}

}
