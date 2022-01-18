package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_Content_Restriction_Operation_Group_Service_Get(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		operationKey       string
		groupNameOrID      string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the group id is provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "61eb397f-3bb8-4fb8-843f-a290b9ff17d0",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/byGroupId/61eb397f-3bb8-4fb8-843f-a290b9ff17d0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the group name is provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the operation key is not provided",
			contentID:          "233838383",
			operationKey:       "",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content restriction operation key set",
		},

		{
			name:               "when the group name or id is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no group id or name set",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is invalid",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
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

			service := &ContentRestrictionOperationGroupService{client: mockClient}

			gotResponse, err := service.Get(testCase.context, testCase.contentID, testCase.operationKey, testCase.groupNameOrID)

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

func Test_Content_Restriction_Operation_Group_Service_Add(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		operationKey       string
		groupNameOrID      string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the group id is provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "61eb397f-3bb8-4fb8-843f-a290b9ff17d0",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/byGroupId/61eb397f-3bb8-4fb8-843f-a290b9ff17d0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the group name is provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the operation key is not provided",
			contentID:          "233838383",
			operationKey:       "",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content restriction operation key set",
		},

		{
			name:               "when the group name or id is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no group id or name set",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is invalid",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
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

			service := &ContentRestrictionOperationGroupService{client: mockClient}

			gotResponse, err := service.Add(testCase.context, testCase.contentID, testCase.operationKey, testCase.groupNameOrID)

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

func Test_Content_Restriction_Operation_Group_Service_Remove(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		operationKey       string
		groupNameOrID      string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the group id is provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "61eb397f-3bb8-4fb8-843f-a290b9ff17d0",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/byGroupId/61eb397f-3bb8-4fb8-843f-a290b9ff17d0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the group name is provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the operation key is not provided",
			contentID:          "233838383",
			operationKey:       "",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content restriction operation key set",
		},

		{
			name:               "when the group name or id is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no group id or name set",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is invalid",
			contentID:          "233838383",
			operationKey:       "read",
			groupNameOrID:      "confluence-system-admins",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/restriction/byOperation/read/group/confluence-system-admins",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
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

			service := &ContentRestrictionOperationGroupService{client: mockClient}

			gotResponse, err := service.Remove(testCase.context, testCase.contentID, testCase.operationKey, testCase.groupNameOrID)

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
