package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestPermissionService_Gets(t *testing.T) {

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
			name:               "GetPermissionsWhenTheContextIsValid",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/permissions",
			mockFile:           "../v3/mocks/get-permissions.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetPermissionsWhenTheRequestMethodIsIncorrect",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/permissions",
			mockFile:           "../v3/mocks/get-permissions.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetPermissionsWhenTheStatusCodeIsIncorrect",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/permissions",
			mockFile:           "../v3/mocks/get-permissions.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetPermissionsWhenTheContextIsNil",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/permissions",
			mockFile:           "../v3/mocks/get-permissions.json",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetPermissionsWhenTheResponseBodyHasADifferentFormat",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/permissions",
			mockFile:           "../v3/mocks/empty_json.json",
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

			i := &PermissionService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context)

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

func TestPermissionService_Check(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.PermissionCheckPayload
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CheckPermissionsWhenTheParametersAreCorrect",
			payload: &models.PermissionCheckPayload{
				GlobalPermissions: []string{"ADMINISTER"},
				AccountID:         "", //
				ProjectPermissions: []*models.BulkProjectPermissionsScheme{
					{
						Issues:      nil,
						Projects:    []int{10000},
						Permissions: []string{"EDIT_ISSUES"},
					},
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/permissions/check",
			mockFile:           "../v3/mocks/check-permissions.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "CheckPermissionsWhenThePayloadIsNotSet",
			payload:            nil,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/permissions/check",
			mockFile:           "../v3/mocks/check-permissions.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "CheckPermissionsWhenTheProjectPermissionIsNotSet",
			payload: &models.PermissionCheckPayload{
				GlobalPermissions:  []string{"ADMINISTER"},
				AccountID:          "", //
				ProjectPermissions: nil,
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/permissions/check",
			mockFile:           "../v3/mocks/check-permissions.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "CheckPermissionsWhenTheRequestMethodIsIncorrect",
			payload: &models.PermissionCheckPayload{
				GlobalPermissions: []string{"ADMINISTER"},
				AccountID:         "", //
				ProjectPermissions: []*models.BulkProjectPermissionsScheme{
					{
						Issues:      nil,
						Projects:    []int{10000},
						Permissions: []string{"EDIT_ISSUES"},
					},
				},
			},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/permissions/check",
			mockFile:           "../v3/mocks/check-permissions.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "CheckPermissionsWhenTheStatusCodeIsIncorrect",
			payload: &models.PermissionCheckPayload{
				GlobalPermissions: []string{"ADMINISTER"},
				AccountID:         "", //
				ProjectPermissions: []*models.BulkProjectPermissionsScheme{
					{
						Issues:      nil,
						Projects:    []int{10000},
						Permissions: []string{"EDIT_ISSUES"},
					},
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/permissions/check",
			mockFile:           "../v3/mocks/check-permissions.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CheckPermissionsWhenTheContextIsNil",
			payload: &models.PermissionCheckPayload{
				GlobalPermissions: []string{"ADMINISTER"},
				AccountID:         "", //
				ProjectPermissions: []*models.BulkProjectPermissionsScheme{
					{
						Issues:      nil,
						Projects:    []int{10000},
						Permissions: []string{"EDIT_ISSUES"},
					},
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/permissions/check",
			mockFile:           "../v3/mocks/check-permissions.json",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "CheckPermissionsWhenTheResponseBodyIsEmpty",
			payload: &models.PermissionCheckPayload{
				GlobalPermissions: []string{"ADMINISTER"},
				AccountID:         "", //
				ProjectPermissions: []*models.BulkProjectPermissionsScheme{
					{
						Issues:      nil,
						Projects:    []int{10000},
						Permissions: []string{"EDIT_ISSUES"},
					},
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/permissions/check",
			mockFile:           "../v3/mocks/empty_json.json",
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

			i := &PermissionService{client: mockClient}

			gotResult, gotResponse, err := i.Check(testCase.context, testCase.payload)

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
