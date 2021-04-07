package admin

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestSCIMSchemeService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetSCIMSchemasWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetSCIMSchemasWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			mockFile:           "./mocks/scim-get-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMSchemasWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-schemas.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMSchemasWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetSCIMSchemasWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMSchemasWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMSchemasWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas",
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

			service := &SCIMSchemeService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.directoryID)

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

				for _, schema := range gotResult.Resources {
					t.Log(schema)
				}

			}

		})
	}

}

func TestSCIMSchemeService_Groups(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetSCIMGroupsWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-group-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetSCIMGroupsWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			mockFile:           "./mocks/scim-get-group-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMGroupsWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-group-schemas.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMGroupsWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-group-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetSCIMGroupsWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-group-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMGroupsWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-group-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMGroupsWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:Group",
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

			service := &SCIMSchemeService{client: mockClient}
			gotResult, gotResponse, err := service.Group(testCase.context, testCase.directoryID)

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

				for _, group := range gotResult.Attributes {
					t.Log(group)
				}

			}

		})
	}

}

func TestSCIMSchemeService_User(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetSCIMUserSchemaWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-user-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetSCIMUserSchemaWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			mockFile:           "./mocks/scim-get-user-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserSchemaWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-user-schemas.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserSchemaWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-user-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserSchemaWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-user-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserSchemaWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-user-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserSchemaWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:core:2.0:User",
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

			service := &SCIMSchemeService{client: mockClient}
			gotResult, gotResponse, err := service.User(testCase.context, testCase.directoryID)

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

				for _, group := range gotResult.Attributes {
					t.Log(group)
				}

			}

		})
	}

}

func TestSCIMSchemeService_Enterprise(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetSCIMUserEnterpriseSchemaWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-user-enterprise-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetSCIMUserEnterpriseSchemaWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			mockFile:           "./mocks/scim-get-user-enterprise-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserEnterpriseSchemaWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-user-enterprise-schemas.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserEnterpriseSchemaWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-user-enterprise-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserEnterpriseSchemaWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-user-enterprise-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserEnterpriseSchemaWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-user-enterprise-schemas.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserEnterpriseSchemaWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Schemas/urn:ietf:params:scim:schemas:extension:enterprise:2.0:User",
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

			service := &SCIMSchemeService{client: mockClient}
			gotResult, gotResponse, err := service.Enterprise(testCase.context, testCase.directoryID)

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

				for _, group := range gotResult.Attributes {
					t.Log(group)
				}

			}

		})
	}

}

func TestSCIMSchemeService_Feature(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetSCIMFeatureMetadataWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-service-provider.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/ServiceProviderConfig",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetSCIMFeatureMetadataWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			mockFile:           "./mocks/scim-get-service-provider.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/ServiceProviderConfig",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMFeatureMetadataWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-service-provider.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/ServiceProviderConfig",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMFeatureMetadataWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-service-provider.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/ServiceProviderConfig",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetSCIMFeatureMetadataWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-service-provider.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/ServiceProviderConfig",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMFeatureMetadataWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/scim-get-service-provider.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMFeatureMetadataWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/ServiceProviderConfig",
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

			service := &SCIMSchemeService{client: mockClient}
			gotResult, gotResponse, err := service.Feature(testCase.context, testCase.directoryID)

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

				for _, group := range gotResult.AuthenticationSchemes {
					t.Log(group)
				}

			}

		})
	}

}
