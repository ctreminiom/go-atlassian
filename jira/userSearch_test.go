package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestUserSearchService_Projects(t *testing.T) {

	testCases := []struct {
		name                string
		accountID           string
		projectKeys         []string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "FindUsersAssignableToProjectsWhenTheParamsAreCorrect",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			projectKeys:        []string{"DUMMY", "FK", "PS"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users-assignable-to-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/assignable/multiProjectSearch?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&projectKeys=DUMMY%2CFK%2CPS&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "FindUsersAssignableToProjectsWhenTheProjectKeysIsNil",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			projectKeys:        nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users-assignable-to-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/assignable/multiProjectSearch?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&projectKeys=DUMMY%2CFK%2CPS&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "FindUsersAssignableToProjectsWhenTheRequestMethodIsIncorrect",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			projectKeys:        []string{"DUMMY", "FK", "PS"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users-assignable-to-projects.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/user/assignable/multiProjectSearch?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&projectKeys=DUMMY%2CFK%2CPS&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "FindUsersAssignableToProjectsWhenTheStatusCodeIsIncorrect",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			projectKeys:        []string{"DUMMY", "FK", "PS"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users-assignable-to-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/assignable/multiProjectSearch?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&projectKeys=DUMMY%2CFK%2CPS&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "FindUsersAssignableToProjectsWhenTheContextIsNil",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			projectKeys:        []string{"DUMMY", "FK", "PS"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users-assignable-to-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/assignable/multiProjectSearch?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&projectKeys=DUMMY%2CFK%2CPS&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "FindUsersAssignableToProjectsWhenTheResponseBodyHasADifferentFormat",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			projectKeys:        []string{"DUMMY", "FK", "PS"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/assignable/multiProjectSearch?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&projectKeys=DUMMY%2CFK%2CPS&startAt=0",
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

			i := &UserSearchService{client: mockClient}

			gotResult, gotResponse, err := i.Projects(testCase.context, testCase.accountID, testCase.projectKeys,
				testCase.startAt,
				testCase.maxResults)

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

				for pos, user := range *gotResult {

					t.Log("------------------------------")
					t.Logf("User Display Name #%v: %v", pos+1, user.DisplayName)
					t.Logf("User Email #%v: %v", pos+1, user.EmailAddress)
					t.Logf("User AccountID #%v: %v", pos+1, user.AccountID)
					t.Logf("User Type #%v: %v", pos+1, user.AccountType)
					t.Log("------------------------------")
				}
			}
		})

	}

}

func TestUserSearchService_Do(t *testing.T) {

	testCases := []struct {
		name                string
		accountID, query    string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "SearchUsersWhenTheParamsAreCorrect",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			query:              "charlie",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/search?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&query=charlie&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "SearchUsersWhenTheRequestMethodIsIncorrect",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-users.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/user/search?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchUsersWhenTheStatusCodeIsIncorrect",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/search?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "SearchUsersWhenTheContextIsNil",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/search?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "SearchUsersWhenTheResponseBodyHasADifferentFormat",
			accountID:          "594b47b5-c774-4d51-9ee8-b604013e9d9a",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/search?accountId=594b47b5-c774-4d51-9ee8-b604013e9d9a&maxResults=50&startAt=0",
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

			i := &UserSearchService{client: mockClient}

			gotResult, gotResponse, err := i.Do(testCase.context, testCase.accountID, testCase.query, testCase.startAt,
				testCase.maxResults)

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

				for pos, user := range *gotResult {

					t.Log("------------------------------")
					t.Logf("User Display Name #%v: %v", pos+1, user.DisplayName)
					t.Logf("User Email #%v: %v", pos+1, user.EmailAddress)
					t.Logf("User AccountID #%v: %v", pos+1, user.AccountID)
					t.Logf("User Type #%v: %v", pos+1, user.AccountType)
					t.Log("------------------------------")
				}
			}
		})

	}

}
