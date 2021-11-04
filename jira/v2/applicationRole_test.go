package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_ApplicationRoleService_V2_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetApplicationRolesWhenTheFormatIsCorrect",
			mockFile:           "../v3/mocks/get_application_roles.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole",
			context:            context.Background(),
			wantErr:            false,
		},
		{
			name:               "GetApplicationRolesWhenTheHTTPMethodIsIncorrect",
			mockFile:           "../v3/mocks/get_application_roles.json",
			wantHTTPCodeReturn: http.StatusMethodNotAllowed,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/applicationrole",
			context:            context.Background(),
			wantErr:            true,
		},
		{
			name:               "GetApplicationRolesWhenTheHTTPResponseCodeIsIncorrect",
			mockFile:           "../v3/mocks/get_application_roles.json",
			wantHTTPCodeReturn: http.StatusMethodNotAllowed,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole",
			context:            context.Background(),
			wantErr:            true,
		},
		{
			name:               "GetApplicationRolesWhenTheContextIsNil",
			mockFile:           "../v3/mocks/get_application_roles.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole",
			context:            nil,
			wantErr:            true,
		},
		{
			name:               "GetApplicationRolesWhenTheEndpointIsIncorrect",
			mockFile:           "../v3/mocks/get_application_roles.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationroleSAX",
			context:            context.Background(),
			wantErr:            true,
		},
		{
			name:               "GetApplicationRolesWhenTheResponseBodyLengthIsZero",
			mockFile:           "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole",
			context:            context.Background(),
			wantErr:            true,
		},
		{
			name:               "GetApplicationRolesWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "../v3/mocks/get-attachment-settings.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole",
			context:            context.Background(),
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

			service := &ApplicationRoleService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func Test_ApplicationRoleService_V2_Get(t *testing.T) {

	testCases := []struct {
		name               string
		applicationRoleKey string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetApplicationRoleWhenTheKeyIsSet",
			applicationRoleKey: "jira-software",
			mockFile:           "../v3/mocks/get_application_role_by_key.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole/jira-software",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetApplicationRoleWhenTheKeyIsNotSet",
			applicationRoleKey: "",
			mockFile:           "../v3/mocks/get_application_role_by_key.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole/jira-software",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetApplicationRoleWhenTheKeyIsIncorrect",
			applicationRoleKey: "jira-scrum-masters",
			mockFile:           "../v3/mocks/get_application_role_by_key.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole/jira-software",
			context:            context.Background(),
			wantErr:            true,
		},
		{
			name:               "GetApplicationRoleWhenTheHTTPMethodIsIncorrect",
			applicationRoleKey: "jira-software",
			mockFile:           "../v3/mocks/get_application_role_by_key.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/applicationrole/jira-software",
			context:            context.Background(),
			wantErr:            true,
		},
		{
			name:               "GetApplicationRoleWhenTheHTTPResponseCodeIsIncorrect",
			applicationRoleKey: "jira-software",
			mockFile:           "../v3/mocks/get_application_role_by_key.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole/jira-software",
			context:            context.Background(),
			wantErr:            true,
		},
		{
			name:               "GetApplicationRoleWhenTheContextIsNil",
			applicationRoleKey: "jira-software",
			mockFile:           "../v3/mocks/get_application_role_by_key.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole/jira-software",
			context:            nil,
			wantErr:            true,
		},
		{
			name:               "GetApplicationRoleWhenTheResponseBodyLengthIsZero",
			applicationRoleKey: "jira-software",
			mockFile:           "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole/jira-software",
			context:            context.Background(),
			wantErr:            true,
		},
		{
			name:               "GetApplicationRoleWhenTheResponseBodyHasADifferentFormat",
			applicationRoleKey: "jira-software",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/applicationrole/jira-software",
			context:            context.Background(),
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

			service := &ApplicationRoleService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.applicationRoleKey)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}
