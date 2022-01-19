package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_Content_Restriction_Operation_User_Service_Get(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		operationKey       string
		accountID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the params are correct",
			contentID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			operationKey:       "read",
			accountID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			operationKey:       "read",
			accountID:          "confluence-system-admins",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the operation key is not provided",
			contentID:          "233838383",
			operationKey:       "",
			accountID:          "confluence-system-admins",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content restriction operation key set",
		},

		{
			name:               "when the accountID is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			accountID:          "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "jira: no account id set",
		},

		{
			name:               "when the response status is not correct",
			contentID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			operationKey:       "read",
			accountID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			accountID:          "confluence-system-admins",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
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

			service := &ContentRestrictionOperationUserService{client: mockClient}

			gotResponse, err := service.Get(testCase.context, testCase.contentID, testCase.operationKey, testCase.accountID)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}
		})
	}
}

func Test_Content_Restriction_Operation_User_Service_Add(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		operationKey       string
		accountID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the params are correct",
			contentID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			operationKey:       "read",
			accountID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			operationKey:       "read",
			accountID:          "confluence-system-admins",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the operation key is not provided",
			contentID:          "233838383",
			operationKey:       "",
			accountID:          "confluence-system-admins",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content restriction operation key set",
		},

		{
			name:               "when the accountID is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			accountID:          "",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "jira: no account id set",
		},

		{
			name:               "when the response status is not correct",
			contentID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			operationKey:       "read",
			accountID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			accountID:          "confluence-system-admins",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
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

			service := &ContentRestrictionOperationUserService{client: mockClient}

			gotResponse, err := service.Add(testCase.context, testCase.contentID, testCase.operationKey, testCase.accountID)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}
		})
	}
}

func Test_Content_Restriction_Operation_User_Service_Remove(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		operationKey       string
		accountID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the params are correct",
			contentID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			operationKey:       "read",
			accountID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			operationKey:       "read",
			accountID:          "confluence-system-admins",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the operation key is not provided",
			contentID:          "233838383",
			operationKey:       "",
			accountID:          "confluence-system-admins",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content restriction operation key set",
		},

		{
			name:               "when the accountID is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			accountID:          "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "jira: no account id set",
		},

		{
			name:               "when the response status is not correct",
			contentID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			operationKey:       "read",
			accountID:          "7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			accountID:          "confluence-system-admins",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/7747bbd7-ade1-401f-bcc2-eb24403e75f3/restriction/byOperation/read/user?accountId=7747bbd7-ade1-401f-bcc2-eb24403e75f3",
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

			service := &ContentRestrictionOperationUserService{client: mockClient}

			gotResponse, err := service.Remove(testCase.context, testCase.contentID, testCase.operationKey, testCase.accountID)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}
		})
	}
}
