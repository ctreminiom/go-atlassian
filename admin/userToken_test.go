package admin

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestUserTokenService_Gets(t *testing.T) {

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
			name:               "GetUserTokensWhenTheParametersAreCorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/get-user-tokens.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetUserTokensWhenTheAccountIDIsNotSet",
			accountID:          "",
			mockFile:           "./mocks/get-user-tokens.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserTokensWhenTheRequestMethodIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/get-user-tokens.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserTokensWhenTheStatusCodeIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/get-user-tokens.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetUserTokensWhenTheContextIsNil",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/get-user-tokens.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserTokensWhenTheEndpointIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/get-user-tokens.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetUserTokensWhenTheResponseBodyIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens",
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

			service := &UserTokenService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.accountID)

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

				for _, token := range *gotResult {
					t.Log(token)
				}
			}

		})
	}

}

func TestUserTokenService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		accountID          string
		tokenID            string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteUserTokenWhenTheParametersAreCorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			tokenID:            "asp_A0T8NbvnwgyVkw3K",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens/asp_A0T8NbvnwgyVkw3K",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteUserTokenWhenTheTokenIDIsNotSet",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			tokenID:            "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens/asp_A0T8NbvnwgyVkw3K",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteUserTokenWhenTheAccountIDIsNotSet",
			accountID:          "",
			tokenID:            "asp_A0T8NbvnwgyVkw3K",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens/asp_A0T8NbvnwgyVkw3K",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteUserTokenWhenTheRequestMethodIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			tokenID:            "asp_A0T8NbvnwgyVkw3K",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens/asp_A0T8NbvnwgyVkw3K",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteUserTokenWhenTheStatusCodeIsIncorrect",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			tokenID:            "asp_A0T8NbvnwgyVkw3K",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens/asp_A0T8NbvnwgyVkw3K",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteUserTokenWhenTheEndpointIsEmpty",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			tokenID:            "asp_A0T8NbvnwgyVkw3K",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteUserTokenWhenTheContextIsNil",
			accountID:          "651c2e11-afea-4475-a0c4-422b89683e0f",
			tokenID:            "asp_A0T8NbvnwgyVkw3K",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/users/651c2e11-afea-4475-a0c4-422b89683e0f/manage/api-tokens/asp_A0T8NbvnwgyVkw3K",
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

			service := &UserTokenService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.accountID, testCase.tokenID)

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
