package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestContentRestrictionOperationService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		expand             []string
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
			contentID:          "233838383",
			expand:             []string{"restrictions.user", "restrictions.group"},
			mockFile:           "./mocks/get-content-restriction-by-operation.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation?expand=restrictions.user%2Crestrictions.group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			expand:             []string{"restrictions.user", "restrictions.group"},
			mockFile:           "./mocks/get-content-restriction-by-operation.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation?expand=restrictions.user%2Crestrictions.group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			expand:             []string{"restrictions.user", "restrictions.group"},
			mockFile:           "./mocks/get-content-restriction-by-operation.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation?expand=restrictions.user%2Crestrictions.group",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is not correct",
			contentID:          "233838383",
			expand:             []string{"restrictions.user", "restrictions.group"},
			mockFile:           "./mocks/get-content-restriction-by-operation.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation?expand=restrictions.user%2Crestrictions.group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the response body is not empty",
			contentID:          "233838383",
			expand:             []string{"restrictions.user", "restrictions.group"},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation?expand=restrictions.user%2Crestrictions.group",
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

			service := &ContentRestrictionOperationService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.contentID, testCase.expand)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}
		})
	}
}

func TestContentRestrictionOperationService_Get(t *testing.T) {

	testCases := []struct {
		name                string
		contentID           string
		operationKey        string
		expand              []string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
		expectedError       string
	}{
		{
			name:               "when the parameters are correct",
			contentID:          "233838383",
			operationKey:       "read",
			expand:             []string{"restrictions.user", "restrictions.group"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-restrictions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:               "when the operation key is not provided",
			contentID:          "233838383",
			operationKey:       "",
			expand:             []string{"restrictions.user", "restrictions.group"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-restrictions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no content restriction operation key set",
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			operationKey:       "read",
			expand:             []string{"restrictions.user", "restrictions.group"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-restrictions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			expand:             []string{"restrictions.user", "restrictions.group"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-restrictions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is not correct",
			contentID:          "233838383",
			operationKey:       "read",
			expand:             []string{"restrictions.user", "restrictions.group"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-restrictions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the response body is not empty",
			contentID:          "233838383",
			operationKey:       "read",
			expand:             []string{"restrictions.user", "restrictions.group"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
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

			service := &ContentRestrictionOperationService{client: mockClient}

			gotResult, gotResponse, err := service.Get(testCase.context, testCase.contentID, testCase.operationKey, testCase.expand,
				testCase.startAt, testCase.maxResults)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}
		})
	}
}
