package admin

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestUserService_Permissions(t *testing.T) {

	testCases := []struct {
		name               string
		accountID          string
		privileges         []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetUserPermissionsWhenTheParametersAreCorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			privileges:         []string{"email", "token", "etc"},
			mockFile:           "./mocks/get-user-permissions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage?privileges=email%2Ctoken%2Cetc",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetUserPermissionsWhenThePrivilegesAreNotSet",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			privileges:         nil,
			mockFile:           "./mocks/get-user-permissions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetUserPermissionsWhenTheAccountIDIsNotSet",
			accountID:          "",
			privileges:         []string{"email, token, etc"},
			mockFile:           "./mocks/get-user-permissions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage?privileges=email%2C+token%2C+etc",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserPermissionsWhenTheRequestMethodIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			privileges:         []string{"email, token, etc"},
			mockFile:           "./mocks/get-user-permissions.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage?privileges=email%2C+token%2C+etc",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserPermissionsWhenTheStatusCodeIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			privileges:         []string{"email, token, etc"},
			mockFile:           "./mocks/get-user-permissions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage?privileges=email%2C+token%2C+etc",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetUserPermissionsWhenTheEndpointIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			privileges:         []string{"email, token, etc"},
			mockFile:           "./mocks/get-user-permissions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserPermissionsWhenTheContextIsNil",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			privileges:         []string{"email, token, etc"},
			mockFile:           "./mocks/get-user-permissions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage?privileges=email%2C+token%2C+etc",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserPermissionsWhenTheRequestBodyIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			privileges:         []string{"email, token, etc"},
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage?privileges=email%2C+token%2C+etc",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &UserService{client: mockClient}
			gotResult, gotResponse, err := service.Permissions(testCase.context, testCase.accountID, testCase.privileges)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

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

				t.Log(gotResult.EmailSet.Allowed)
			}

		})
	}

}

func TestUserService_Get(t *testing.T) {

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
			name:               "GetUserWhenTheParametersAreCorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetUserWhenTheAccountIDIsNotSet",
			accountID:          "",
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserWhenTheRequestMethodIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserWhenTheStatusCodeIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetUserWhenTheEndpointIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserWhenTheContextIsNil",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserWhenTheRequestBodyIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &UserService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.accountID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

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

				t.Log(gotResult.Account)
			}

		})
	}

}

func TestUserService_Update(t *testing.T) {

	var mockedPayload = make(map[string]interface{})
	mockedPayload["nickname"] = "marshmallow"

	var mockedPayloadWithOutKeys = make(map[string]interface{})

	testCases := []struct {
		name               string
		accountID          string
		payload            map[string]interface{}
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "UpdateUserWhenTheParametersAreCorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload:            mockedPayload,
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "UpdateUserWhenThePayloadDoesNotContainsKeys",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload:            mockedPayloadWithOutKeys,
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateUserWhenThePayloadIsNil",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload:            nil,
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateUserWhenTheAccountIDIsNotSet",
			accountID:          "",
			payload:            mockedPayload,
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateUserWhenTheRequestMethodIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload:            mockedPayload,
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateUserWhenTheStatusCodeIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload:            mockedPayload,
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "UpdateUserWhenTheEndpointIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload:            mockedPayload,
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateUserWhenTheContextIsNil",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload:            mockedPayload,
			mockFile:           "./mocks/get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateUserWhenTheRequestBodyIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload:            mockedPayload,
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/profile",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &UserService{client: mockClient}
			gotResult, gotResponse, err := service.Update(testCase.context, testCase.accountID, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

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

				t.Log(gotResult.Account)
			}

		})
	}

}

func TestUserService_Disable(t *testing.T) {

	testCases := []struct {
		name               string
		accountID          string
		message            string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DisableUserWhenTheParametersAreCorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			message:            "Sample message",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/disable",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DisableUserWhenTheAccountIDIsNotSet",
			accountID:          "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/disable",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DisableUserWhenTheRequestMethodIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/disable",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DisableUserWhenTheStatusCodeIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/disable",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DisableUserWhenTheEndpointIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DisableUserWhenTheContextIsNil",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/disable",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DisableUserWhenTheContextIsNilAndHasAMessage",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			message:            "Sample message",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/disable",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &UserService{client: mockClient}
			gotResponse, err := service.Disable(testCase.context, testCase.accountID, testCase.message)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

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

func TestUserService_Enable(t *testing.T) {

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
			name:               "EnableUserWhenTheParametersAreCorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/enable",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "EnableUserWhenTheAccountIDIsNotSet",
			accountID:          "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/enable",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "EnableUserWhenTheRequestMethodIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/enable",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "EnableUserWhenTheStatusCodeIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/enable",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "EnableUserWhenTheEndpointIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "EnableUserWhenTheContextIsNil",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/lifecycle/enable",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &UserService{client: mockClient}
			gotResponse, err := service.Enable(testCase.context, testCase.accountID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

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
