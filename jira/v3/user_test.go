package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestUserService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.UserPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateUserWhenTheParamsAreCorrect",
			payload: &models.UserPayloadScheme{
				Password:     "password1234",
				EmailAddress: "example@example.com",
				DisplayName:  "Example User",
				Notification: true,
			},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateUserWhenThePayloadIsNil",
			payload:            nil,
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateUserWhenTheRequestMethodIsIncorrect",
			payload: &models.UserPayloadScheme{
				Password:     "password1234",
				EmailAddress: "example@example.com",
				DisplayName:  "Example User",
				Notification: true,
			},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateUserWhenTheStatusCodeIsIncorrect",
			payload: &models.UserPayloadScheme{
				Password:     "password1234",
				EmailAddress: "example@example.com",
				DisplayName:  "Example User",
				Notification: true,
			},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CreateUserWhenTheContextIsNil",
			payload: &models.UserPayloadScheme{
				Password:     "password1234",
				EmailAddress: "example@example.com",
				DisplayName:  "Example User",
				Notification: true,
			},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/user",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateUserWhenTheEndpointIsIncorrect",
			payload: &models.UserPayloadScheme{
				Password:     "password1234",
				EmailAddress: "example@example.com",
				DisplayName:  "Example User",
				Notification: true,
			},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateUserWhenTheResponseBodyHasADifferentFormat",
			payload: &models.UserPayloadScheme{
				Password:     "password1234",
				EmailAddress: "example@example.com",
				DisplayName:  "Example User",
				Notification: true,
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/user",
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

			i := &UserService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.payload)

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

func TestUserService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		accountID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CreateUserWhenTheParamsAreCorrect",
			accountID:          "c5f6fccf-9195-4a50-82d0-a7c4ce5a5b78",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/user?accountId=c5f6fccf-9195-4a50-82d0-a7c4ce5a5b78",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "CreateUserWhenTheAccountIDIsEmpty",
			accountID:          "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/user?accountId=c5f6fccf-9195-4a50-82d0-a7c4ce5a5b78",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "CreateUserWhenTheContextIsNil",
			accountID:          "c5f6fccf-9195-4a50-82d0-a7c4ce5a5b78",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/user?accountId=c5f6fccf-9195-4a50-82d0-a7c4ce5a5b78",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "CreateUserWhenTheRequestMethodIsIncorrect",
			accountID:          "c5f6fccf-9195-4a50-82d0-a7c4ce5a5b78",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/user?accountId=c5f6fccf-9195-4a50-82d0-a7c4ce5a5b78",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "CreateUserWhenTheStatusCodeIsIncorrect",
			accountID:          "c5f6fccf-9195-4a50-82d0-a7c4ce5a5b78",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/user?accountId=c5f6fccf-9195-4a50-82d0-a7c4ce5a5b78",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
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

			i := &UserService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.accountID)

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

func TestUserService_Find(t *testing.T) {

	testCases := []struct {
		name                string
		accountIDs          []string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "FindsUsersWhenTheParamsAreCorrect",
			accountIDs:         []string{"aa2098e3-8c1c-498a-b67f-5c2f53c271cd", "f79c7c0c-9366-4acc-9a79-2e21763b1d29"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/bulk?accountId=aa2098e3-8c1c-498a-b67f-5c2f53c271cd&accountId=f79c7c0c-9366-4acc-9a79-2e21763b1d29&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "FindsUsersWhenTheAccountIDsIsEmpty",
			accountIDs:         nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/bulk?accountId=aa2098e3-8c1c-498a-b67f-5c2f53c271cd&accountId=f79c7c0c-9366-4acc-9a79-2e21763b1d29&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "FindsUsersWhenTheRequestMethodIsIncorrect",
			accountIDs:         []string{"aa2098e3-8c1c-498a-b67f-5c2f53c271cd", "f79c7c0c-9366-4acc-9a79-2e21763b1d29"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/user/bulk?accountId=aa2098e3-8c1c-498a-b67f-5c2f53c271cd&accountId=f79c7c0c-9366-4acc-9a79-2e21763b1d29&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "FindsUsersWhenTheStatusCodeIsIncorrect",
			accountIDs:         []string{"aa2098e3-8c1c-498a-b67f-5c2f53c271cd", "f79c7c0c-9366-4acc-9a79-2e21763b1d29"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/bulk?accountId=aa2098e3-8c1c-498a-b67f-5c2f53c271cd&accountId=f79c7c0c-9366-4acc-9a79-2e21763b1d29&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "FindsUsersWhenTheContextIsNil",
			accountIDs:         []string{"aa2098e3-8c1c-498a-b67f-5c2f53c271cd", "f79c7c0c-9366-4acc-9a79-2e21763b1d29"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/bulk?accountId=aa2098e3-8c1c-498a-b67f-5c2f53c271cd&accountId=f79c7c0c-9366-4acc-9a79-2e21763b1d29&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "FindsUsersWhenTheEndpointIsIncorrect",
			accountIDs:         []string{"aa2098e3-8c1c-498a-b67f-5c2f53c271cd", "f79c7c0c-9366-4acc-9a79-2e21763b1d29"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/find-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/user/bulk?accountId=aa2098e3-8c1c-498a-b67f-5c2f53c271cd&accountId=f79c7c0c-9366-4acc-9a79-2e21763b1d29&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "FindsUsersWhenTheResponseBodyHasADifferentFormat",
			accountIDs:         []string{"aa2098e3-8c1c-498a-b67f-5c2f53c271cd", "f79c7c0c-9366-4acc-9a79-2e21763b1d29"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/bulk?accountId=aa2098e3-8c1c-498a-b67f-5c2f53c271cd&accountId=f79c7c0c-9366-4acc-9a79-2e21763b1d29&maxResults=50&startAt=0",
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

			i := &UserService{client: mockClient}

			gotResult, gotResponse, err := i.Find(testCase.context, testCase.accountIDs, testCase.startAt,
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

				for position, user := range gotResult.Values {

					t.Log("---------------")
					t.Logf("User name #%v, %v", position, user.DisplayName)
					t.Logf("User mail #%v, %v", position, user.EmailAddress)
					t.Logf("User accountID #%v, %v", position, user.AccountID)
				}

			}
		})

	}

}

func TestUserService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		accountID          string
		expands            []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetUserWhenTheParamsAreCorrect",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			expands:            []string{"groups", "applicationRoles"},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user?accountId=5b10ac8d82e05b22cc7d4ef5&expand=groups%2CapplicationRoles",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetUserWhenTheAccountIDIsEmpty",
			accountID:          "",
			expands:            []string{"groups", "applicationRoles"},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user?accountId=5b10ac8d82e05b22cc7d4ef5&expand=groups%2CapplicationRoles",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserWhenTheExpandIsNil",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			expands:            nil,
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user?accountId=5b10ac8d82e05b22cc7d4ef5",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetUserWhenTheRequestMethodIsIncorrect",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			expands:            []string{"groups", "applicationRoles"},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/user?accountId=5b10ac8d82e05b22cc7d4ef5&expand=groups%2CapplicationRoles",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserWhenTheStatusCodeIsIncorrect",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			expands:            []string{"groups", "applicationRoles"},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user?accountId=5b10ac8d82e05b22cc7d4ef5&expand=groups%2CapplicationRoles",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetUserWhenTheContextIsNil",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			expands:            []string{"groups", "applicationRoles"},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user?accountId=5b10ac8d82e05b22cc7d4ef5&expand=groups%2CapplicationRoles",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetUserWhenTheEndpointsIsIncorrect",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			expands:            []string{"groups", "applicationRoles"},
			mockFile:           "./mocks/create-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/user?accountId=5b10ac8d82e05b22cc7d4ef5&expand=groups%2CapplicationRoles",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserWhenTheResponseBodyHasADifferentFormat",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			expands:            []string{"groups", "applicationRoles"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user?accountId=5b10ac8d82e05b22cc7d4ef5&expand=groups%2CapplicationRoles",
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

			i := &UserService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.accountID, testCase.expands)

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

				t.Log("---------------")
				t.Logf("User name %v", gotResult.DisplayName)
				t.Logf("User mail %v", gotResult.EmailAddress)
				t.Logf("User accountID %v", gotResult.AccountID)
				t.Log("---------------")

			}
		})

	}

}

func TestUserService_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetUsersWhenTheParamsAreCorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/users/search?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetUsersWhenTheRequestMethodIsIncorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-users.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/users/search?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUsersWhenTheStatusCodeIsIncorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/users/search?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetUsersWhenTheEndpointIsIncorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/users/search?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUsersWhenTheContextIsNil",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/users/search?maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUsersWhenTheResponseBodyHasADifferentFormat",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/users/search?maxResults=50&startAt=0",
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

			i := &UserService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.startAt, testCase.maxResults)

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

				for _, user := range gotResult {
					t.Log("---------------")
					t.Logf("User name %v", user.DisplayName)
					t.Logf("User mail %v", user.EmailAddress)
					t.Logf("User accountID %v", user.AccountID)
					t.Log("---------------")
				}

			}
		})

	}

}

func TestUserService_Groups(t *testing.T) {

	testCases := []struct {
		name               string
		accountID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetUserGroupsWhenTheParamsAreCorrect",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			mockFile:           "./mocks/get-user-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/groups?accountId=5b10ac8d82e05b22cc7d4ef5",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetUserGroupsWhenTheAccountIDIsEmpty",
			accountID:          "",
			mockFile:           "./mocks/get-user-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/groups?accountId=5b10ac8d82e05b22cc7d4ef5",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetUserGroupsWhenTheRequestMethodIsIncorrect",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			mockFile:           "./mocks/get-user-groups.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/user/groups?accountId=5b10ac8d82e05b22cc7d4ef5",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserGroupsWhenTheStatusCodeIsIncorrect",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			mockFile:           "./mocks/get-user-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/groups?accountId=5b10ac8d82e05b22cc7d4ef5",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetUserGroupsWhenTheEndpointIsIncorrect",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			mockFile:           "./mocks/get-user-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/user/groups?5b10ac8d82e05b22cc7d4ef5=",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserGroupsWhenTheContextIsNil",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			mockFile:           "./mocks/get-user-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/groups?accountId=5b10ac8d82e05b22cc7d4ef5",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserGroupsWhenTheResponseBodyHasADifferentFormat",
			accountID:          "5b10ac8d82e05b22cc7d4ef5",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/user/groups?accountId=5b10ac8d82e05b22cc7d4ef5",
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

			i := &UserService{client: mockClient}

			gotResult, gotResponse, err := i.Groups(testCase.context, testCase.accountID)

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

				for _, group := range gotResult {
					t.Log("---------------")
					t.Logf("Group name %v", group.Name)
					t.Logf("Group URL %v", group.Self)
					t.Log("---------------")
				}

			}
		})

	}

}
