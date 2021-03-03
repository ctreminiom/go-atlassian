package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestPermissionGrantSchemeService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		schemeID           int
		payload            *PermissionGrantPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:     "CreatePermissionGrantWhenTheParametersAreCorrect",
			schemeID: 1000,
			payload: &PermissionGrantPayloadScheme{
				Holder: &PermissionGrantHolderPayloadScheme{
					Parameter: "scrum-masters",
					Type:      "group",
				},
				Permission: "EDIT_ISSUES",
			},
			mockFile:           "./mocks/get-permission-grant.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreatePermissionGrantWhenThePayloadIsNil",
			schemeID:           1000,
			payload:            nil,
			mockFile:           "./mocks/get-permission-grant.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:     "CreatePermissionGrantWhenTheEndpointIsIncorrect",
			schemeID: 1000,
			payload: &PermissionGrantPayloadScheme{
				Holder: &PermissionGrantHolderPayloadScheme{
					Parameter: "scrum-masters",
					Type:      "group",
				},
				Permission: "EDIT_ISSUES",
			},
			mockFile:           "./mocks/get-permission-grant.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme/10001/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:     "CreatePermissionGrantWhenTheRequestMethodIsIncorrect",
			schemeID: 1000,
			payload: &PermissionGrantPayloadScheme{
				Holder: &PermissionGrantHolderPayloadScheme{
					Parameter: "scrum-masters",
					Type:      "group",
				},
				Permission: "EDIT_ISSUES",
			},
			mockFile:           "./mocks/get-permission-grant.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:     "CreatePermissionGrantWhenTheStatusCodeIsIncorrect",
			schemeID: 1000,
			payload: &PermissionGrantPayloadScheme{
				Holder: &PermissionGrantHolderPayloadScheme{
					Parameter: "scrum-masters",
					Type:      "group",
				},
				Permission: "EDIT_ISSUES",
			},
			mockFile:           "./mocks/get-permission-grant.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:     "CreatePermissionGrantWhenTheContextIsNil",
			schemeID: 1000,
			payload: &PermissionGrantPayloadScheme{
				Holder: &PermissionGrantHolderPayloadScheme{
					Parameter: "scrum-masters",
					Type:      "group",
				},
				Permission: "EDIT_ISSUES",
			},
			mockFile:           "./mocks/get-permission-grant.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:     "CreatePermissionGrantWhenTheResponseBodyHasADifferentFormat",
			schemeID: 1000,
			payload: &PermissionGrantPayloadScheme{
				Holder: &PermissionGrantHolderPayloadScheme{
					Parameter: "scrum-masters",
					Type:      "group",
				},
				Permission: "EDIT_ISSUES",
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission",
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

			i := &PermissionGrantSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.schemeID, testCase.payload)

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

func TestPermissionGrantSchemeService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		schemeID           int
		permissionID       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeletePermissionGrantWhenTheParametersAreCorrect",
			schemeID:           1000,
			permissionID:       10002,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission/10002",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeletePermissionGrantWhenTheEndpointIsIncorrect",
			schemeID:           1000,
			permissionID:       10002,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeletePermissionGrantWhenTheRequestMethodIsIncorrect",
			schemeID:           1000,
			permissionID:       10002,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission/10002",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeletePermissionGrantWhenTheStatusCodeIsIncorrect",
			schemeID:           1000,
			permissionID:       10002,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission/10002",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeletePermissionGrantWhenTheContextIsNil",
			schemeID:           1000,
			permissionID:       10002,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/permissionscheme/1000/permission/10002",
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

			i := &PermissionGrantSchemeService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.schemeID, testCase.permissionID)

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

func TestPermissionGrantSchemeService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		schemeID           int
		permissionID       int
		expands            []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetPermissionGrantWhenTheParametersAreCorrect",
			schemeID:           1000,
			permissionID:       10002,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-permission-grant.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission/10002?expand=field%2Cgroup%2Cpermissions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetPermissionGrantWhenTheExpandIsNil",
			schemeID:           1000,
			permissionID:       10002,
			expands:            nil,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-permission-grant.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission/10002",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetPermissionGrantWhenTheRequestMethodIsIncorrect",
			schemeID:           1000,
			permissionID:       10002,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "./mocks/get-permission-grant.json",
			endpoint:           "/rest/api/3/permissionscheme/1001/permission/10002?expand=field%2Cgroup%2Cpermissions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetPermissionGrantWhenTheStatusIsIncorrect",
			schemeID:           1000,
			permissionID:       10002,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-permission-grant.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission/10002?expand=field%2Cgroup%2Cpermissions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetPermissionGrantWhenTheContextIsNil",
			schemeID:           1000,
			permissionID:       10002,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-permission-grant.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission/10002?expand=field%2Cgroup%2Cpermissions",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetPermissionGrantWhenTheResponseBodyHasADifferentFormat",
			schemeID:           1000,
			permissionID:       10002,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/empty_json.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission/10002?expand=field%2Cgroup%2Cpermissions",
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

			i := &PermissionGrantSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.schemeID, testCase.permissionID, testCase.expands)

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

func TestPermissionGrantSchemeService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		schemeID           int
		expands            []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetPermissionGrantsWhenTheParametersAreCorrect",
			schemeID:           1000,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-permission-grants.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission?expand=field%2Cgroup%2Cpermissions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetPermissionGrantsWhenTheExpandIsNil",
			schemeID:           1000,
			expands:            nil,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-permission-grants.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetPermissionGrantsWhenTheEndpointIsIncorrect",
			schemeID:           1000,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-permission-grants.json",
			endpoint:           "/rest/api/3/permissionscheme/1001/permission?expand=field%2Cgroup%2Cpermissions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetPermissionGrantsWhenTheRequestMethodIsIncorrect",
			schemeID:           1000,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "./mocks/get-permission-grants.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission?expand=field%2Cgroup%2Cpermissions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetPermissionGrantsWhenTheStatusCodeIsIncorrect",
			schemeID:           1000,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-permission-grants.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission?expand=field%2Cgroup%2Cpermissions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetPermissionGrantsWhenTheContextIsNil",
			schemeID:           1000,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-permission-grants.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission?expand=field%2Cgroup%2Cpermissions",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetPermissionGrantsWhenTheResponseBodyHasADifferentFormat",
			schemeID:           1000,
			expands:            []string{"field", "group", "permissions"},
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/empty_json.json",
			endpoint:           "/rest/api/3/permissionscheme/1000/permission?expand=field%2Cgroup%2Cpermissions",
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

			i := &PermissionGrantSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.schemeID, testCase.expands)

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
