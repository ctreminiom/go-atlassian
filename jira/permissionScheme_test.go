package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestPermissionSchemeService_Create(t *testing.T) {

	testCases := []struct {
		name                        string
		permissionSchemeName        string
		permissionSchemeDescription string
		permissionGrants            *[]PermissionGrantPayloadScheme
		mockFile                    string
		wantHTTPMethod              string
		endpoint                    string
		context                     context.Context
		wantHTTPCodeReturn          int
		wantErr                     bool
	}{
		{
			name:                        "CreatePermissionSchemeWhenTheParametersAreCorrect",
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:                        "CreatePermissionSchemeWhenThePermissionGrantsIsNil",
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants:            nil,
			mockFile:                    "./mocks/get-permission-scheme.json",
			wantHTTPMethod:              http.MethodPost,
			endpoint:                    "/rest/api/3/permissionscheme",
			context:                     context.Background(),
			wantHTTPCodeReturn:          http.StatusCreated,
			wantErr:                     true,
		},

		{
			name:                        "CreatePermissionSchemeWhenThePermissionSchemeNameParamIsEmpty",
			permissionSchemeName:        "",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:                        "CreatePermissionSchemeWhenTheEndpointIsIncorrect",
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/apis/3/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:                        "CreatePermissionSchemeWhenTheRequestMethodIsIncorrect",
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:                        "CreatePermissionSchemeWhenTheStatusCodeIsIncorrect",
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:                        "CreatePermissionSchemeWhenTheContextIsNil",
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:                        "CreatePermissionSchemeWhenTheResponseBodyHasADifferentFormat",
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme",
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

			i := &PermissionSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context,
				testCase.permissionSchemeName,
				testCase.permissionSchemeDescription, testCase.permissionGrants)

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

func TestPermissionSchemeService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		permissionSchemeID int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeletePermissionSchemeWhenTheParametersAreCorrect",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/permissionscheme/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeletePermissionSchemeWhenTheEndpointIsIncorrect",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/permissionscheme/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeletePermissionSchemeWhenTheContextIsNil",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/permissionscheme/1000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeletePermissionSchemeWhenTheRequestMethodIsIncorrect",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/permissionscheme/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeletePermissionSchemeWhenTheStatusCodeIsIncorrect",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/permissionscheme/1000",
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

			i := &PermissionSchemeService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.permissionSchemeID)

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

func TestPermissionSchemeService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		permissionSchemeID int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetPermissionSchemeWhenTheParametersAreCorrect",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/permissionscheme/1000",
			mockFile:           "./mocks/get-permission-scheme.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetPermissionSchemeWhenTheEndpointIsIncorrect",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/permissionschemes/1000",
			mockFile:           "./mocks/get-permission-scheme.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetPermissionSchemeWhenTheRequestMethodIsIncorrect",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme/1000",
			mockFile:           "./mocks/get-permission-scheme.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetPermissionSchemeWhenTheStatusCodeIsIncorrect",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/permissionscheme/1000",
			mockFile:           "./mocks/get-permission-scheme.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetPermissionSchemeWhenTheContextIsNil",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/permissionscheme/1000",
			mockFile:           "./mocks/get-permission-scheme.json",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetPermissionSchemeWhenTheResponseBodyHasADifferentFormat",
			permissionSchemeID: 1000,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/permissionscheme/1000",
			mockFile:           "./mocks/empty_json.json",
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

			i := &PermissionSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.permissionSchemeID)

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

func TestPermissionSchemeService_Gets(t *testing.T) {

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
			name:               "GetPermissionSchemesWhenTheContextIsValid",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/permissionscheme",
			mockFile:           "./mocks/get-permission-schemes.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetPermissionSchemesWhenTheRequestMethodIsIncorrect",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/permissionscheme",
			mockFile:           "./mocks/get-permission-schemes.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetPermissionSchemesWhenTheStatusCodeIsIncorrect",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/permissionscheme",
			mockFile:           "./mocks/get-permission-schemes.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetPermissionSchemesWhenTheContextIsNil",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/permissionscheme",
			mockFile:           "./mocks/get-permission-schemes.json",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetPermissionSchemesWhenTheResponseBodyHasADifferentFormat",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/permissionscheme",
			mockFile:           "./mocks/empty_json.json",
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

			i := &PermissionSchemeService{client: mockClient}

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

func TestPermissionSchemeService_Update(t *testing.T) {
	testCases := []struct {
		name                        string
		schemeID                    int
		permissionSchemeName        string
		permissionSchemeDescription string
		permissionGrants            *[]PermissionGrantPayloadScheme
		mockFile                    string
		wantHTTPMethod              string
		endpoint                    string
		context                     context.Context
		wantHTTPCodeReturn          int
		wantErr                     bool
	}{
		{
			name:                        "UpdatePermissionSchemeWhenTheParametersAreCorrect",
			schemeID:                    1000,
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/permissionscheme/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:                        "UpdatePermissionSchemeWhenThePermissionGrantsIsNil",
			schemeID:                    1000,
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants:            nil,
			mockFile:                    "./mocks/get-permission-scheme.json",
			wantHTTPMethod:              http.MethodPost,
			endpoint:                    "/rest/api/3/permissionscheme/1000",
			context:                     context.Background(),
			wantHTTPCodeReturn:          http.StatusCreated,
			wantErr:                     true,
		},

		{
			name:                        "CreatePermissionSchemeWhenTheEndpointIsIncorrect",
			schemeID:                    1000,
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/apis/3/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:                        "CreatePermissionSchemeWhenTheRequestMethodIsIncorrect",
			schemeID:                    1000,
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:                        "CreatePermissionSchemeWhenTheStatusCodeIsIncorrect",
			schemeID:                    1000,
			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/permissionscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:     "CreatePermissionSchemeWhenTheContextIsNil",
			schemeID: 1000,

			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/get-permission-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/permissionscheme",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:     "CreatePermissionSchemeWhenTheResponseBodyHasADifferentFormat",
			schemeID: 1000,

			permissionSchemeName:        "DUMMY - Permission Scheme",
			permissionSchemeDescription: "Dummy permission scheme description",
			permissionGrants: &[]PermissionGrantPayloadScheme{
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
				{
					Holder: &PermissionGrantHolderPayloadScheme{
						Parameter: "scrum-masters",
						Type:      "group",
					},
					Permission: "EDIT_ISSUES",
				},
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/permissionscheme/1000",
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

			i := &PermissionSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.schemeID, testCase.permissionSchemeName,
				testCase.permissionSchemeDescription, testCase.permissionGrants)

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
