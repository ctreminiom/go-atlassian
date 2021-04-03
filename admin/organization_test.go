package admin

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestOrganizationService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		cursor             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetOrganizationsWhenTheParametersAreCorrect",
			mockFile:           "./mocks/get-organizations.json",
			cursor:             "assays22222",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs?cursor=assays22222",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationsWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-organizations.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationsWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-organizations.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationsWhenTheContextIsNil",
			mockFile:           "./mocks/get-organizations.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationsWhenTheEndpointIsEmpty",
			mockFile:           "./mocks/get-organizations.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationsWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-organizations.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationsWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.cursor)

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

				for _, organization := range gotResult.Data {
					t.Log(organization.ID, organization.Attributes.Name)
				}
			}

		})
	}

}

func TestOrganizationService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetOrganizationWhenTheParametersAreCorrect",
			mockFile:           "./mocks/get-organization.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationWhenTheOrganizationIDIsNotSet",
			mockFile:           "./mocks/get-organization.json",
			organizationID:     "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "GetOrganizationWhenTheRequestMethodIsIncorrect",
			mockFile:       "./mocks/get-organization.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",

			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "GetOrganizationWhenTheStatusCodeIsIncorrect",
			mockFile:       "./mocks/get-organization.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",

			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "GetOrganizationWhenTheContextIsNil",
			mockFile:       "./mocks/get-organization.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",

			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267",
			context:            nil,
			wantErr:            true,
		},

		{
			name:           "GetOrganizationWhenTheEndpointIsEmpty",
			mockFile:       "./mocks/get-organization.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",

			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "GetOrganizationWhenTheEndpointIsIncorrect",
			mockFile:       "./mocks/get-organization.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",

			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:           "GetOrganizationWhenTheResponseBodyIsEmpty",
			mockFile:       "./mocks/empty.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",

			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.organizationID)

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

				t.Log(gotResult.Data.ID, gotResult.Data.Attributes.Name)
			}

		})
	}

}

func TestOrganizationService_Users(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     string
		cursor             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetOrganizationUsersWhenTheParametersAreCorrect",
			mockFile:           "./mocks/get-organization-users.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			cursor:             "xxxx",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/users?cursor=xxxx",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationUsersWhenTheOrganizationIDIsNotSet",
			mockFile:           "./mocks/get-organization-users.json",
			organizationID:     "",
			cursor:             "xxxx",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/users?cursor=xxxx",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationUsersWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-organization-users.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/users",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationUsersWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-organization-users.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/users",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationUsersWhenTheContextIsNil",
			mockFile:           "./mocks/get-organization-users.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/users",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationUsersWhenTheEndpointIsEmpty",
			mockFile:           "./mocks/get-organization-users.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationUsersWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-organization-users.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationUsersWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/users",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Users(testCase.context, testCase.organizationID, testCase.cursor)

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

				for _, user := range gotResult.Data {
					t.Log(user.Name, user.Email)
				}
			}

		})
	}

}

func TestOrganizationService_Domains(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     string
		cursor             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetOrganizationDomainsWhenTheParametersAreCorrect",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			cursor:             "xxxx",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains?cursor=xxxx",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationDomainsWhenTheOrganizationIDIsNotSet",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "",
			cursor:             "xxxx",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains?cursor=xxxx",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainsWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainsWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainsWhenTheContextIsNil",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainsWhenTheEndpointIsEmpty",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainsWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainsWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Domains(testCase.context, testCase.organizationID, testCase.cursor)

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

				for _, domain := range gotResult.Data {
					t.Log(domain.ID, domain.Attributes.Name, domain.Attributes.Claim.Status)
				}
			}

		})
	}

}

func TestOrganizationService_Domain(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     string
		domainID           string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetOrganizationDomainWhenTheParametersAreCorrect",
			mockFile:           "./mocks/get-organization-domain.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			domainID:           "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains/go-atlassian.io",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationDomainWhenTheDomainIDIsNotSet",
			mockFile:           "./mocks/get-organization-domain.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			domainID:           "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains/go-atlassian.io",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainWhenTheOrganizationIDIsNotSet",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "",
			domainID:           "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains/go-atlassian.io",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			domainID:           "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains/go-atlassian.io",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			domainID:           "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains/go-atlassian.io",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainWhenTheContextIsNil",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			domainID:           "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains/go-atlassian.io",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainWhenTheEndpointIsEmpty",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			domainID:           "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-organization-domains.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			domainID:           "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationDomainWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			domainID:           "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/domains/go-atlassian.io",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Domain(testCase.context, testCase.organizationID, testCase.domainID)

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

				t.Log(gotResult.Data.ID, gotResult.Data.Attributes.Name, gotResult.Data.Attributes.Claim.Status)
			}

		})
	}

}
