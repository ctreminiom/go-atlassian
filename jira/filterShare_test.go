package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestFilterShareService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *PermissionFilterBodyScheme
		mockFile           string
		filterID           int
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "AddShareFilterPermissionWhenTheTypeIsAGroup",
			payload: &PermissionFilterBodyScheme{
				Type:      "group",
				ProjectID: "jira-administrators",
			},
			filterID:           10010,
			mockFile:           "./mocks/add_share_filter_permission_group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/10010/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},
		{
			name:               "AddShareFilterPermissionWhenThePayloadIsNil",
			payload:            nil,
			filterID:           10010,
			mockFile:           "./mocks/add_share_filter_permission_group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/10010/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "AddShareFilterPermissionWhenTheResponseBodyHasADifferentFormat",
			payload: &PermissionFilterBodyScheme{
				Type:      "group",
				ProjectID: "jira-administrators",
			},
			filterID:           10010,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/10010/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "AddShareFilterPermissionWhenTheTypeIsAProject",
			payload: &PermissionFilterBodyScheme{
				Type:      "project",
				ProjectID: "EX",
			},
			filterID:           10010,
			mockFile:           "./mocks/add_share_filter_permission_project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/10010/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},
		{
			name: "AddShareFilterPermissionWhenTheTypeIsAProjectAndHasARole",
			payload: &PermissionFilterBodyScheme{
				Type:          "project",
				ProjectID:     "EX",
				ProjectRoleID: "10360",
			},
			filterID:           10010,
			mockFile:           "./mocks/add_share_filter_permission_project_and_role.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/10010/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},
		{
			name: "AddShareFilterPermissionWhenTheContextIsNil",
			payload: &PermissionFilterBodyScheme{
				Type:          "project",
				ProjectID:     "EX",
				ProjectRoleID: "10360",
			},
			filterID:           10010,
			mockFile:           "./mocks/add_share_filter_permission_project_and_role.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/10010/permission",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "AddShareFilterPermissionWhenTheEndpointIsIncorrect",
			payload: &PermissionFilterBodyScheme{
				Type:          "project",
				ProjectID:     "EX",
				ProjectRoleID: "10360",
			},
			filterID:           10010,
			mockFile:           "./mocks/add_share_filter_permission_project_and_role.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/10010/permissiosdasdn",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
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

			service := &FilterShareService{client: mockClient}
			gotResult, gotResponse, err := service.Add(testCase.context, testCase.filterID, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				for _, filterPermission := range *gotResult {
					t.Logf("Filter ID Wanted: %v, Filter ID Returned: %v", filterPermission.ID, testCase.filterID)

					assert.Equal(t, filterPermission.ID, testCase.filterID)
				}
			}

		})
	}
}

func TestFilterShareService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		filterID           int
		permissionID       int
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteShareFilterPermissionWhenTheFilterIsCorrect",
			filterID:           10010,
			permissionID:       1111,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/filter/10010/permission/1111",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},
		{
			name:               "DeleteShareFilterPermissionWhenTheFilterIDIsIncorrect",
			filterID:           10011,
			permissionID:       1111,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/filter/10010/permission/1111",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteShareFilterPermissionWhenTheContextIsNil",
			filterID:           10010,
			permissionID:       1111,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/filter/10010/permission/1111",
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

			service := &FilterShareService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.filterID, testCase.permissionID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func TestFilterShareService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		filterID           int
		permissionID       int
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetShareFilterPermissionWhenTheFilterIsCorrect",
			filterID:           10000,
			mockFile:           "./mocks/get_share_permission.json",
			permissionID:       1111,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000/permission/1111",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetShareFilterPermissionWhenTheResponseBodyHasADifferentFormat",
			filterID:           10000,
			mockFile:           "./mocks/empty_json.json",
			permissionID:       1111,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000/permission/1111",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetShareFilterPermissionWhenTheFilterIDIsIncorrect",
			filterID:           10001,
			mockFile:           "./mocks/get_share_permission.json",
			permissionID:       1111,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000/permission/1111",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetShareFilterPermissionWhenTheContextIsNil",
			filterID:           10000,
			mockFile:           "./mocks/get_share_permission.json",
			permissionID:       1111,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000/permission/1111",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetShareFilterPermissionWhenThePermissionIDIsIncorrect",
			filterID:           10000,
			mockFile:           "./mocks/get_share_permission.json",
			permissionID:       111123,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000/permission/1111",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetShareFilterPermissionWhenThePermissionIDIsEmpty",
			filterID:           10000,
			mockFile:           "./mocks/get_share_permission.json",
			permissionID:       0,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000/permission/1111",
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

			service := &FilterShareService{client: mockClient}
			getResult, gotResponse, err := service.Get(testCase.context, testCase.filterID, testCase.permissionID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, getResult, nil)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				t.Logf("HTTP Filter ID Wanted: %v, HTTP Filter ID Returned: %v", testCase.filterID, getResult.ID)

				assert.Equal(t, testCase.filterID, getResult.ID)

			}

		})
	}
}

func TestFilterShareService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		filterID           int
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetShareFilterPermissionsWhenTheFilterIDIsCorrect",
			filterID:           10000,
			mockFile:           "./mocks/get_share_permissions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetShareFilterPermissionsWhenTheResponseBodyHasADifferentFormat",
			filterID:           10000,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetShareFilterPermissionsWhenTheFilterIDIsIncorrect",
			filterID:           10001,
			mockFile:           "./mocks/get_share_permissions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetShareFilterPermissionsWhenTheContextIsNil",
			filterID:           10000,
			mockFile:           "./mocks/get_share_permissions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000/permission",
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

			service := &FilterShareService{client: mockClient}
			getResult, gotResponse, err := service.Gets(testCase.context, testCase.filterID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, getResult, nil)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				for index, sharePermission := range *getResult {
					t.Logf("Share Permission #%v, | ID: %v, Type: %v", index, sharePermission.ID, sharePermission.Type)
				}
			}

		})
	}
}

func TestFilterShareService_Scope(t *testing.T) {

	testCases := []struct {
		name               string
		defaultScope       string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetDefaultFilterScopeWhenTheParametersAreCorrect",
			defaultScope:       "GLOBAL",
			mockFile:           "./mocks/get_default_share_scope.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/defaultShareScope",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetDefaultFilterScopeWhenTheContextIsNil",
			defaultScope:       "GLOBAL",
			mockFile:           "./mocks/get_default_share_scope.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/defaultShareScope",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDefaultFilterScopeWhenTheRequestMethodIsIncorrect",
			defaultScope:       "GLOBAL",
			mockFile:           "./mocks/get_default_share_scope.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/defaultShareScope",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDefaultFilterScopeWhenTheStatusCodeIsIncorrect",
			defaultScope:       "GLOBAL",
			mockFile:           "./mocks/get_default_share_scope.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/defaultShareScope",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetDefaultFilterScopeWhenTheResponseBodyHasADifferentFormat",
			defaultScope:       "GLOBAL",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/defaultShareScope",
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

			service := &FilterShareService{client: mockClient}
			getResult, gotResponse, err := service.Scope(testCase.context)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				}
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, getResult, nil)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				t.Logf("HTTP Default Scope Wanted: %v, HTTP Default Scope Returned: %v", testCase.defaultScope, getResult)
				assert.Equal(t, testCase.defaultScope, getResult)
			}

		})
	}
}

func TestFilterShareService_SetScope(t *testing.T) {

	testCases := []struct {
		name               string
		defaultScope       string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "SetDefaultFilterScopeWhenTheParametersAreCorrect",
			defaultScope:       "GLOBAL",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/defaultShareScope",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "SetDefaultFilterScopeWhenTheScopeIsIncorrect",
			defaultScope:       "GLOBAL_NO_VALID",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/defaultShareScope",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "SetDefaultFilterScopeWhenTheContextIsNil",
			defaultScope:       "GLOBAL",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/defaultShareScope",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "SetDefaultFilterScopeWhenTheRequestMethodIsIncorrect",
			defaultScope:       "GLOBAL",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/defaultShareScope",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "SetDefaultFilterScopeWhenTheStatusCodeIsIncorrect",
			defaultScope:       "GLOBAL",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/defaultShareScope",
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

			service := &FilterShareService{client: mockClient}
			gotResponse, err := service.SetScope(testCase.context, testCase.defaultScope)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
			}
		})
	}
}
