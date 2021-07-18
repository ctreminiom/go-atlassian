package admin

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"time"
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.cursor)

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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.organizationID)

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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Users(testCase.context, testCase.organizationID, testCase.cursor)

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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Domains(testCase.context, testCase.organizationID, testCase.cursor)

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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Domain(testCase.context, testCase.organizationID, testCase.domainID)

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

				t.Log(gotResult.Data.ID, gotResult.Data.Attributes.Name, gotResult.Data.Attributes.Claim.Status)
			}

		})
	}

}

func TestOrganizationService_Events(t *testing.T) {

	fromMocked, err := time.Parse(time.RFC3339Nano, "2020-05-12T11:45:26.371Z")
	if err != nil {
		t.Fatal(err)
	}

	toMocked, err := time.Parse(time.RFC3339Nano, "2020-11-12T11:45:26.371Z")
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name               string
		organizationID     string
		opts               *OrganizationEventOptScheme
		cursor             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:           "GetOrganizationAuditEventsWhenTheParametersAreCorrect",
			mockFile:       "./mocks/get-organization-audit-events.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			opts: &OrganizationEventOptScheme{
				Q:      "qq",
				From:   fromMocked.Add(time.Duration(-24) * time.Hour),
				To:     toMocked.Add(time.Duration(-1) * time.Hour),
				Action: "user_added_to_group",
			},
			cursor:             "d57e-483a",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events?action=user_added_to_group&cursor=d57e-483a&from=1589197526&q=qq&to=1605177926",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:           "GetOrganizationAuditEventsWhenTheResponseBodyIsEmpty",
			mockFile:       "./mocks/empty.json",
			organizationID: "d094d850-d57e-483a-bd03-ca8855919267",
			opts: &OrganizationEventOptScheme{
				Q:      "qq",
				From:   fromMocked.Add(time.Duration(-24) * time.Hour),
				To:     toMocked.Add(time.Duration(-1) * time.Hour),
				Action: "user_added_to_group",
			},
			cursor:             "d57e-483a",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events?action=user_added_to_group&cursor=d57e-483a&from=1589197526&q=qq&to=1605177926",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditEventsWhenTheDomainIDIsNotSet",
			mockFile:           "./mocks/get-organization-audit-events.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			cursor:             "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events?action=user_added_to_group&cursor=d57e-483a&from=1589197526&q=qq&to=1605177926",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditEventsWhenTheOrganizationIDIsNotSet",
			mockFile:           "./mocks/get-organization-audit-events.json",
			organizationID:     "",
			cursor:             "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events?action=user_added_to_group&cursor=d57e-483a&from=1589197526&q=qq&to=1605177926",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditEventsWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-organization-audit-events.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			cursor:             "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events?action=user_added_to_group&cursor=d57e-483a&from=1589197526&q=qq&to=1605177926",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditEventsWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-organization-audit-events.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			cursor:             "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events?action=user_added_to_group&cursor=d57e-483a&from=1589197526&q=qq&to=1605177926",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditEventsWhenTheContextIsNil",
			mockFile:           "./mocks/get-organization-audit-events.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			cursor:             "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events?action=user_added_to_group&cursor=d57e-483a&from=1589197526&q=qq&to=1605177926",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditEventsWhenTheEndpointIsEmpty",
			mockFile:           "./mocks/get-organization-audit-events.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			cursor:             "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditEventsWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-organization-audit-events.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			cursor:             "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditEventsWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			cursor:             "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events?action=user_added_to_group&cursor=d57e-483a&from=1589197526&q=qq&to=1605177926",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditEventsWhenTheResponseBodyIsIncorrect",
			mockFile:           "./mocks/get-organization-domain.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			cursor:             "go-atlassian.io",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events?action=user_added_to_group&cursor=d57e-483a&from=1589197526&q=qq&to=1605177926",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Events(testCase.context, testCase.organizationID, testCase.opts, testCase.cursor)

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

				for _, event := range gotResult.Data {
					t.Log(event.ID, event.Attributes.Action, event.Attributes.Time)
				}
			}

		})
	}

}

func TestOrganizationService_Event(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     string
		eventID            string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetOrganizationEventWhenTheParametersAreCorrect",
			mockFile:           "./mocks/get-organization-audit-event.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			eventID:            "0b86a87f-f376-46f8-bdb3-25bd5200e161",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events/0b86a87f-f376-46f8-bdb3-25bd5200e161",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationEventWhenTheEventIDIsNotSet",
			mockFile:           "./mocks/get-organization-audit-event.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			eventID:            "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events/0b86a87f-f376-46f8-bdb3-25bd5200e161",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationEventWhenTheOrganizationIDIsNotSet",
			mockFile:           "./mocks/get-organization-audit-event.json",
			organizationID:     "",
			eventID:            "0b86a87f-f376-46f8-bdb3-25bd5200e161",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events/0b86a87f-f376-46f8-bdb3-25bd5200e161",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationEventWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-organization-audit-event.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			eventID:            "0b86a87f-f376-46f8-bdb3-25bd5200e161",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events/0b86a87f-f376-46f8-bdb3-25bd5200e161",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationEventWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-organization-audit-event.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			eventID:            "0b86a87f-f376-46f8-bdb3-25bd5200e161",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events/0b86a87f-f376-46f8-bdb3-25bd5200e161",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationEventWhenTheContextIsNil",
			mockFile:           "./mocks/get-organization-audit-event.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			eventID:            "0b86a87f-f376-46f8-bdb3-25bd5200e161",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events/0b86a87f-f376-46f8-bdb3-25bd5200e161",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationEventWhenTheEndpointIsEmpty",
			mockFile:           "./mocks/get-organization-audit-event.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			eventID:            "0b86a87f-f376-46f8-bdb3-25bd5200e161",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationEventWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-organization-audit-event.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			eventID:            "0b86a87f-f376-46f8-bdb3-25bd5200e161",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationEventWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			eventID:            "0b86a87f-f376-46f8-bdb3-25bd5200e161",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/events/0b86a87f-f376-46f8-bdb3-25bd5200e161",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Event(testCase.context, testCase.organizationID, testCase.eventID)

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

				t.Log(gotResult.Data.ID, gotResult.Data.Attributes.Time, gotResult.Data.Attributes.Action)
			}

		})
	}

}

func TestOrganizationService_Actions(t *testing.T) {

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
			name:               "GetOrganizationAuditLogActionsWhenTheParametersAreCorrect",
			mockFile:           "./mocks/get-organization-audit-event-actions.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/event-actions",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetOrganizationAuditLogActionsWhenTheOrganizationIDIsNotSet",
			mockFile:           "./mocks/get-organization-audit-event-actions.json",
			organizationID:     "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/event-actions",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditLogActionsWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-organization-audit-event-actions.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/event-actions",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditLogActionsWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-organization-audit-event-actions.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/event-actions",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditLogActionsWhenTheContextIsNil",
			mockFile:           "./mocks/get-organization-audit-event-actions.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/event-actions",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditLogActionsWhenTheEndpointIsEmpty",
			mockFile:           "./mocks/get-organization-audit-event-actions.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditLogActionsWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-organization-audit-event-actions.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetOrganizationAuditLogActionsWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty.json",
			organizationID:     "d094d850-d57e-483a-bd03-ca8855919267",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/d094d850-d57e-483a-bd03-ca8855919267/event-actions",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Actions(testCase.context, testCase.organizationID)

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

				for _, action := range gotResult.Data {
					t.Log(action.ID, action.Type, action.Attributes.DisplayName)
				}
			}

		})
	}

}
