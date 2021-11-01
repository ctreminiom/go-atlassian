package v3

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestWatcherService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddIssueWatcherWhenTheIssueKeyIsProvided",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "AddIssueWatcherWhenTheIssueKeyIsEmpty",
			issueKeyOrID:       "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddIssueWatcherWhenTheEndpointIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watcher",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "AddIssueWatcherWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddIssueWatcherWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AddIssueWatcherWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			context:            nil,
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

			i := &WatcherService{client: mockClient}

			gotResponse, err := i.Add(testCase.context, testCase.issueKeyOrID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
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

func TestWatcherService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		accountID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteIssueWatcherWhenTheIssueKeyIsProvided",
			issueKeyOrID:       "DUMMY-3",
			accountID:          "e75fa4ab-c86b-4500-b522-de963f31b928",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers?accountId=e75fa4ab-c86b-4500-b522-de963f31b928",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteIssueWatcherWhenTheEmptyIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			accountID:          "e75fa4ab-c86b-4500-b522-de963f31b928",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWatcherWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			accountID:          "e75fa4ab-c86b-4500-b522-de963f31b928",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers?accountId=e75fa4ab-c86b-4500-b522-de963f31b928",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWatcherWhenTheIssueKeyIsEmpty",
			issueKeyOrID:       "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers?accountId=e75fa4ab-c86b-4500-b522-de963f31b928",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWatcherWhenTheEndpointIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers?aaccountId=e75fa4ab-c86b-4500-b522-de963f31b928",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteIssueWatcherWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers?accountId=e75fa4ab-c86b-4500-b522-de963f31b928",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWatcherWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers?accountId=e75fa4ab-c86b-4500-b522-de963f31b928",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWatcherWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers?accountId=e75fa4ab-c86b-4500-b522-de963f31b928",
			context:            nil,
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

			i := &WatcherService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.issueKeyOrID, testCase.accountID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
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

func TestWatcherService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetsIssueWatchersWhenTheIssueKeyIsProvided",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			mockFile:           "./mocks/get-issue-watchers.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsIssueWatchersWhenTheResponseBodyIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			mockFile:           "./mocks/empty_json.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueWatchersWhenTheIssueKeyIsEmpty",
			issueKeyOrID:       "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			mockFile:           "./mocks/get-issue-watchers.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueWatchersWhenTheEndpointIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/vote",
			mockFile:           "./mocks/get-issue-watchers.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsIssueWatchersWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			mockFile:           "./mocks/get-issue-watchers.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueWatchersWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			mockFile:           "./mocks/get-issue-watchers.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetsIssueWatchersWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/watchers",
			mockFile:           "./mocks/get-issue-watchers.json",
			context:            nil,
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

			i := &WatcherService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.issueKeyOrID)

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
