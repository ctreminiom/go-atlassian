package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestProjectPermissionSchemeService_Assign(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		permissionSchemeID int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AssignPermissionSchemeToProjectWhenTheParametersAreCorrect",
			projectKeyOrID:     "10001",
			permissionSchemeID: 001,
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/10001/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "AssignPermissionSchemeToProjectWhenTheProjectKeyOrIDIsEmpty",
			projectKeyOrID:     "",
			permissionSchemeID: 001,
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/10001/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AssignPermissionSchemeToProjectWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "10001",
			permissionSchemeID: 001,
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/10001/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AssignPermissionSchemeToProjectWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "10001",
			permissionSchemeID: 001,
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/10001/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AssignPermissionSchemeToProjectWhenTheContextIsNil",
			projectKeyOrID:     "10001",
			permissionSchemeID: 001,
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/10001/permissionscheme",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AssignPermissionSchemeToProjectWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "10001",
			permissionSchemeID: 001,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/10001/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AssignPermissionSchemeToProjectWhenTheEndpointIsIncorrect",
			projectKeyOrID:     "10001",
			permissionSchemeID: 001,
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/10001/permissionscheme",
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

			i := &ProjectPermissionSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Assign(testCase.context, testCase.projectKeyOrID, testCase.permissionSchemeID)

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

func TestProjectPermissionSchemeService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		expands            []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectPermissionSchemeWhenTheParametersAreCorrect",
			projectKeyOrID:     "10001",
			expands:            []string{"all", "field", "group"},
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/permissionscheme?expand=all%2Cfield%2Cgroup",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectPermissionSchemeWhenTheExpandParamIsEmpty",
			projectKeyOrID:     "10001",
			expands:            nil,
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectPermissionSchemeWhenTheProjectKeyOrIDIsEmpty",
			projectKeyOrID:     "",
			expands:            []string{"all", "field", "group"},
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/permissionscheme?expand=all%2Cfield%2Cgroup",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectPermissionSchemeWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "10001",
			expands:            []string{"all", "field", "group"},
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/10001/permissionscheme?expand=all%2Cfield%2Cgroup",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectPermissionSchemeWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "10001",
			expands:            []string{"all", "field", "group"},
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/permissionscheme?expand=all%2Cfield%2Cgroup",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectPermissionSchemeWhenTheContextIsNil",
			projectKeyOrID:     "10001",
			expands:            []string{"all", "field", "group"},
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/permissionscheme?expand=all%2Cfield%2Cgroup",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectPermissionSchemeWhenTheEndpointIsIncorrect",
			projectKeyOrID:     "10001",
			expands:            []string{"all", "field", "group"},
			mockFile:           "./mocks/get-project-permission-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/10001/permissionscheme?expand=all%2Cfield%2Cgroup",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectPermissionSchemeWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "10001",
			expands:            []string{"all", "field", "group"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/permissionscheme?expand=all%2Cfield%2Cgroup",
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

			i := &ProjectPermissionSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.projectKeyOrID, testCase.expands)

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

func TestProjectPermissionSchemeService_SecurityLevels(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectSecurityLevelsWhenTheParametersAreCorrect",
			projectKeyOrID:     "10001",
			mockFile:           "./mocks/get-project-security-levels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/securitylevel",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectSecurityLevelsWhenTheProjectKeyOrIDIsEmpty",
			projectKeyOrID:     "",
			mockFile:           "./mocks/get-project-security-levels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/securitylevel",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectSecurityLevelsWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "10001",
			mockFile:           "./mocks/get-project-security-levels.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/10001/securitylevel",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectSecurityLevelsWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "10001",
			mockFile:           "./mocks/get-project-security-levels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/securitylevel",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectSecurityLevelsWhenTheContextIsNil",
			projectKeyOrID:     "10001",
			mockFile:           "./mocks/get-project-security-levels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/securitylevel",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectSecurityLevelsWhenTheEndpointIsIncorrect",
			projectKeyOrID:     "10001",
			mockFile:           "./mocks/get-project-security-levels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/10001/securitylevel",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectSecurityLevelsWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "10001",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/10001/securitylevel",
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

			i := &ProjectPermissionSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.SecurityLevels(testCase.context, testCase.projectKeyOrID)

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
