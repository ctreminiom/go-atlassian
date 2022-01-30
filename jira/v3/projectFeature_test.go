package v3

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_ProjectFeatureService_Gets(t *testing.T) {

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
			mockFile:           "../mocks/get-project-features.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/features",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},
		{
			name:               "when the project key or id not provided",
			projectKeyOrID:     "",
			mockFile:           "../mocks/get-project-features.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/features",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "jira: no project id set",
		},
		{
			name:               "when the context is not provided",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../mocks/get-project-features.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/features",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is incorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../mocks/get-project-features.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/features",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the response body is empty",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/features",
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

			i := &ProjectFeatureService{client: mockClient}

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

func Test_ProjectFeatureService_Set(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		featureKey, state  string
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
			featureKey:         "jsw.classic.roadmap",
			state:              "ENABLED",
			mockFile:           "../mocks/get-project-features.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY/features/jsw.classic.roadmap",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:               "when the feature key is not provided",
			projectKeyOrID:     "DUMMY",
			featureKey:         "",
			state:              "ENABLED",
			mockFile:           "../mocks/get-project-features.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY/features/jsw.classic.roadmap",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "jira: no project feature key set",
		},
		{
			name:               "when the project key or id not provided",
			projectKeyOrID:     "",
			featureKey:         "jsw.classic.roadmap",
			state:              "ENABLED",
			mockFile:           "../mocks/get-project-features.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY/features/jsw.classic.roadmap",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "jira: no project id set",
		},
		{
			name:               "when the context is not provided",
			projectKeyOrID:     "DUMMY",
			featureKey:         "jsw.classic.roadmap",
			state:              "ENABLED",
			mockFile:           "../mocks/get-project-features.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY/features/jsw.classic.roadmap",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is incorrect",
			projectKeyOrID:     "DUMMY",
			featureKey:         "jsw.classic.roadmap",
			state:              "ENABLED",
			mockFile:           "../mocks/get-project-features.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY/features/jsw.classic.roadmap",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the response body is empty",
			projectKeyOrID:     "DUMMY",
			featureKey:         "jsw.classic.roadmap",
			state:              "ENABLED",
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY/features/jsw.classic.roadmap",
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

			i := &ProjectFeatureService{client: mockClient}

			gotResult, gotResponse, err := i.Set(testCase.context, testCase.projectKeyOrID, testCase.featureKey, testCase.state)

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
