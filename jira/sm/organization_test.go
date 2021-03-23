package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestOrganizationService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     int
		accountIDs         []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddUsersToOrganizationWhenTheParamsAreCorrect",
			organizationID:     2,
			accountIDs:         []string{"18878b34-f4db-4385-9768-93f562a96b53", "e4365b37-71d2-4711-95b1-b825b5b9d197"},
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/organization/2/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "AddUsersToOrganizationWhenTheAccountIDsAreNotSet",
			organizationID:     2,
			accountIDs:         nil,
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/organization/2/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddUsersToOrganizationWhenTheRequestMethodIsIncorrect",
			organizationID:     2,
			accountIDs:         []string{"18878b34-f4db-4385-9768-93f562a96b53", "e4365b37-71d2-4711-95b1-b825b5b9d197"},
			mockFile:           "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/organization/2/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddUsersToOrganizationWhenTheStatusCodeIsIncorrect",
			organizationID:     2,
			accountIDs:         []string{"18878b34-f4db-4385-9768-93f562a96b53", "e4365b37-71d2-4711-95b1-b825b5b9d197"},
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/organization/2/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AddUsersToOrganizationWhenTheContextIsNil",
			organizationID:     2,
			accountIDs:         []string{"18878b34-f4db-4385-9768-93f562a96b53", "e4365b37-71d2-4711-95b1-b825b5b9d197"},
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/organization/2/user",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddUsersToOrganizationWhenTheEndpointIsEmpty",
			organizationID:     2,
			accountIDs:         []string{"18878b34-f4db-4385-9768-93f562a96b53", "e4365b37-71d2-4711-95b1-b825b5b9d197"},
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
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

			service := &OrganizationService{client: mockClient}
			gotResponse, err := service.Add(testCase.context, testCase.organizationID, testCase.accountIDs)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestOrganizationService_Associate(t *testing.T) {

	testCases := []struct {
		name                string
		serviceDeskPortalID int
		organizationID      int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:                "AssociateOrganizationToServiceManagementPortalWhenTheParametersAreCorrect",
			organizationID:      2,
			serviceDeskPortalID: 1,
			mockFile:            "",
			wantHTTPMethod:      http.MethodPost,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusNoContent,
			wantErr:             false,
		},

		{
			name:                "AssociateOrganizationToServiceManagementPortalWhenTheRequestMethodIsIncorrect",
			organizationID:      2,
			serviceDeskPortalID: 1,
			mockFile:            "",
			wantHTTPMethod:      http.MethodGet,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusNoContent,
			wantErr:             true,
		},

		{
			name:                "AssociateOrganizationToServiceManagementPortalWhenTheStatusCodeIsIncorrect",
			organizationID:      2,
			serviceDeskPortalID: 1,
			mockFile:            "",
			wantHTTPMethod:      http.MethodPost,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusBadRequest,
			wantErr:             true,
		},

		{
			name:                "AssociateOrganizationToServiceManagementPortalWhenTheContextIsNil",
			organizationID:      2,
			serviceDeskPortalID: 1,
			mockFile:            "",
			wantHTTPMethod:      http.MethodPost,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization",
			context:             nil,
			wantHTTPCodeReturn:  http.StatusNoContent,
			wantErr:             true,
		},

		{
			name:                "AssociateOrganizationToServiceManagementPortalWhenTheEndpointIsEmpty",
			organizationID:      2,
			serviceDeskPortalID: 1,
			mockFile:            "",
			wantHTTPMethod:      http.MethodPost,
			endpoint:            "",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusNoContent,
			wantErr:             true,
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
			gotResponse, err := service.Associate(testCase.context, testCase.serviceDeskPortalID, testCase.organizationID)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestOrganizationService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		organizationName   string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CreateOrganizationWhenTheParametersAreCorrect",
			organizationName:   "HR's organization",
			mockFile:           "./mocks/create-organization.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/organization",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "CreateOrganizationWhenTheOrganizationNameIsNotSet",
			organizationName:   "",
			mockFile:           "./mocks/create-organization.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/organization",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateOrganizationWhenTheRequestMethodIsIncorrect",
			organizationName:   "HR's organization",
			mockFile:           "./mocks/create-organization.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateOrganizationWhenTheStatusCodeIsIncorrect",
			organizationName:   "HR's organization",
			mockFile:           "./mocks/create-organization.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/organization",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "CreateOrganizationWhenTheContextIsNil",
			organizationName:   "HR's organization",
			mockFile:           "./mocks/create-organization.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/organization",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "CreateOrganizationWhenTheEndpointIsEmpty",
			organizationName:   "HR's organization",
			mockFile:           "./mocks/create-organization.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateOrganizationWhenTheResponseBodyHasADifferentFormat",
			organizationName:   "HR's organization",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/organization",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.organizationName)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				t.Log("-------------------------------------------")
				t.Logf("New Organization Name: %v", gotResult.Name)
				t.Logf("New Organization ID: %v", gotResult.ID)
				t.Log("-------------------------------------------")
			}

		})
	}

}

func TestOrganizationService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteOrganizationWhenTheParametersAreCorrect",
			organizationID:     2,
			mockFile:           "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/organization/2",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteOrganizationWhenTheRequestMethodIsIncorrect",
			organizationID:     2,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization/2",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteOrganizationWhenTheStatusCodeIsIncorrect",
			organizationID:     2,
			mockFile:           "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/organization/2",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteOrganizationWhenTheContextIsNil",
			organizationID:     2,
			mockFile:           "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/organization/2",
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

			service := &OrganizationService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.organizationID)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestOrganizationService_Detach(t *testing.T) {

	testCases := []struct {
		name                                string
		serviceDeskPortalID, organizationID int
		mockFile                            string
		wantHTTPMethod                      string
		endpoint                            string
		context                             context.Context
		wantHTTPCodeReturn                  int
		wantErr                             bool
	}{
		{
			name:                "DetachOrganizationWhenTheParametersAreCorrect",
			serviceDeskPortalID: 1,
			organizationID:      2,
			mockFile:            "",
			wantHTTPMethod:      http.MethodDelete,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusNoContent,
			wantErr:             false,
		},

		{
			name:                "DetachOrganizationWhenTheRequestMethodIsIncorrect",
			serviceDeskPortalID: 1,
			organizationID:      2,
			mockFile:            "",
			wantHTTPMethod:      http.MethodGet,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusNoContent,
			wantErr:             true,
		},

		{
			name:                "DetachOrganizationWhenTheStatusCodeIsIncorrect",
			serviceDeskPortalID: 1,
			organizationID:      2,
			mockFile:            "",
			wantHTTPMethod:      http.MethodDelete,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusBadRequest,
			wantErr:             true,
		},

		{
			name:                "DetachOrganizationWhenTheContextIsNil",
			serviceDeskPortalID: 1,
			organizationID:      2,
			mockFile:            "",
			wantHTTPMethod:      http.MethodDelete,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization",
			context:             nil,
			wantHTTPCodeReturn:  http.StatusNoContent,
			wantErr:             true,
		},
		{
			name:                "DetachOrganizationWhenTheEndpointsIsEmpty",
			serviceDeskPortalID: 1,
			organizationID:      2,
			mockFile:            "",
			wantHTTPMethod:      http.MethodDelete,
			endpoint:            "",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusNoContent,
			wantErr:             true,
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
			gotResponse, err := service.Detach(testCase.context, testCase.serviceDeskPortalID, testCase.organizationID)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestOrganizationService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetOrganizationWhenTheParametersAreCorrect",
			organizationID:     1,
			mockFile:           "./mocks/get-organization.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetOrganizationWhenTheRequestMethodIsIncorrect",
			organizationID:     1,
			mockFile:           "./mocks/get-organization.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/organization/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationWhenTheStatusCodeIsIncorrect",
			organizationID:     1,
			mockFile:           "./mocks/get-organization.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationWhenTheContextIsNil",
			organizationID:     1,
			mockFile:           "./mocks/get-organization.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization/1",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationWhenTheEndpointIsEmpty",
			organizationID:     1,
			mockFile:           "./mocks/get-organization.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationWhenTheResponseBodyHasADifferentFormat",
			organizationID:     1,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization/1",
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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				t.Log("-------------------------------------------")
				t.Logf("Organization Name: %v", gotResult.Name)
				t.Logf("Organization ID: %v", gotResult.ID)
				t.Log("-------------------------------------------")
			}

		})
	}

}

func TestOrganizationService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		accountID          string
		start, limit       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetOrganizationsWhenTheParametersAreCorrect",
			accountID:          "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-organizations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization?accountId=8a920b17-afbc-49b1-a40d-c19f30465a2d&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetOrganizationsWhenTheRequestMethodIsIncorrect",
			accountID:          "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-organizations.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/organization?accountId=8a920b17-afbc-49b1-a40d-c19f30465a2d&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationsWhenTheStatusCodeIsIncorrect",
			accountID:          "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-organizations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization?accountId=8a920b17-afbc-49b1-a40d-c19f30465a2d&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationsWhenTheEndpointIsEmpty",
			accountID:          "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-organizations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationsWhenTheContextIsNil",
			accountID:          "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-organizations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization?accountId=8a920b17-afbc-49b1-a40d-c19f30465a2d&limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationsWhenTheResponseBodyHasADifferentFormat",
			accountID:          "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization?accountId=8a920b17-afbc-49b1-a40d-c19f30465a2d&limit=50&start=0",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.accountID, testCase.start, testCase.limit)

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

				for _, organization := range gotResult.Values {
					t.Log("-------------------------------------------")
					t.Logf("Organization Name: %v", organization.Name)
					t.Logf("Organization ID: %v", organization.ID)
					t.Log("-------------------------------------------")
				}
			}

		})
	}

}

func TestOrganizationService_Project(t *testing.T) {

	testCases := []struct {
		name                              string
		accountID                         string
		serviceDeskPortalID, start, limit int
		mockFile                          string
		wantHTTPMethod                    string
		endpoint                          string
		context                           context.Context
		wantHTTPCodeReturn                int
		wantErr                           bool
	}{
		{
			name:                "GetOrganizationsAssociatedWithAPortalWhenTheParametersAreCorrect",
			accountID:           "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			serviceDeskPortalID: 1,
			start:               0,
			limit:               50,
			mockFile:            "./mocks/get-project-organizations.json",
			wantHTTPMethod:      http.MethodGet,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization?accountId=8a920b17-afbc-49b1-a40d-c19f30465a2d&limit=50&start=0",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusOK,
			wantErr:             false,
		},

		{
			name:                "GetOrganizationsAssociatedWithAPortalWhenTheRequestMethodIsIncorrect",
			accountID:           "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			serviceDeskPortalID: 1,
			start:               0,
			limit:               50,
			mockFile:            "./mocks/get-project-organizations.json",
			wantHTTPMethod:      http.MethodPost,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization?accountId=8a920b17-afbc-49b1-a40d-c19f30465a2d&limit=50&start=0",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusOK,
			wantErr:             true,
		},

		{
			name:                "GetOrganizationsAssociatedWithAPortalWhenTheStatusCodeIsIncorrect",
			accountID:           "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			serviceDeskPortalID: 1,
			start:               0,
			limit:               50,
			mockFile:            "./mocks/get-project-organizations.json",
			wantHTTPMethod:      http.MethodGet,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization?accountId=8a920b17-afbc-49b1-a40d-c19f30465a2d&limit=50&start=0",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusBadRequest,
			wantErr:             true,
		},

		{
			name:                "GetOrganizationsAssociatedWithAPortalWhenTheContextIsNil",
			accountID:           "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			serviceDeskPortalID: 1,
			start:               0,
			limit:               50,
			mockFile:            "./mocks/get-project-organizations.json",
			wantHTTPMethod:      http.MethodGet,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization?accountId=8a920b17-afbc-49b1-a40d-c19f30465a2d&limit=50&start=0",
			context:             nil,
			wantHTTPCodeReturn:  http.StatusOK,
			wantErr:             true,
		},

		{
			name:                "GetOrganizationsAssociatedWithAPortalWhenTheEndpointIsEmpty",
			accountID:           "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			serviceDeskPortalID: 1,
			start:               0,
			limit:               50,
			mockFile:            "./mocks/get-project-organizations.json",
			wantHTTPMethod:      http.MethodGet,
			endpoint:            "",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusOK,
			wantErr:             true,
		},
		{
			name:                "GetOrganizationsAssociatedWithAPortalWhenTheResponseBodyHasADifferentFormat",
			accountID:           "8a920b17-afbc-49b1-a40d-c19f30465a2d",
			serviceDeskPortalID: 1,
			start:               0,
			limit:               50,
			mockFile:            "./mocks/empty_json.json",
			wantHTTPMethod:      http.MethodGet,
			endpoint:            "/rest/servicedeskapi/servicedesk/1/organization?accountId=8a920b17-afbc-49b1-a40d-c19f30465a2d&limit=50&start=0",
			context:             context.Background(),
			wantHTTPCodeReturn:  http.StatusOK,
			wantErr:             true,
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
			gotResult, gotResponse, err := service.Project(testCase.context, testCase.accountID, testCase.serviceDeskPortalID, testCase.start, testCase.limit)

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

				for _, organization := range gotResult.Values {
					t.Log("-------------------------------------------")
					t.Logf("Organization Name: %v", organization.Name)
					t.Logf("Organization ID: %v", organization.ID)
					t.Log("-------------------------------------------")
				}
			}

		})
	}

}

func TestOrganizationService_Remove(t *testing.T) {

	testCases := []struct {
		name               string
		organizationID     int
		accountIDs         []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "RemoveUsersToOrganizationWhenTheParamsAreCorrect",
			organizationID:     2,
			accountIDs:         []string{"18878b34-f4db-4385-9768-93f562a96b53", "e4365b37-71d2-4711-95b1-b825b5b9d197"},
			mockFile:           "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/organization/2/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "RemoveUsersToOrganizationWhenTheAccountIDsAreNotSet",
			organizationID:     2,
			accountIDs:         nil,
			mockFile:           "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/organization/2/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveUsersToOrganizationWhenTheRequestMethodIsIncorrect",
			organizationID:     2,
			accountIDs:         []string{"18878b34-f4db-4385-9768-93f562a96b53", "e4365b37-71d2-4711-95b1-b825b5b9d197"},
			mockFile:           "",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/servicedeskapi/organization/2/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveUsersToOrganizationWhenTheStatusCodeIsIncorrect",
			organizationID:     2,
			accountIDs:         []string{"18878b34-f4db-4385-9768-93f562a96b53", "e4365b37-71d2-4711-95b1-b825b5b9d197"},
			mockFile:           "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/organization/2/user",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "RemoveUsersToOrganizationWhenTheContextIsNil",
			organizationID:     2,
			accountIDs:         []string{"18878b34-f4db-4385-9768-93f562a96b53", "e4365b37-71d2-4711-95b1-b825b5b9d197"},
			mockFile:           "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/organization/2/user",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveUsersToOrganizationWhenTheEndpointIsEmpty",
			organizationID:     2,
			accountIDs:         []string{"18878b34-f4db-4385-9768-93f562a96b53", "e4365b37-71d2-4711-95b1-b825b5b9d197"},
			mockFile:           "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
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

			service := &OrganizationService{client: mockClient}
			gotResponse, err := service.Remove(testCase.context, testCase.organizationID, testCase.accountIDs)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestOrganizationService_Users(t *testing.T) {

	testCases := []struct {
		name                         string
		organizationID, start, limit int
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:               "GetOrganizationUsersWhenTheResponseBodyHasADifferentFormat",
			organizationID:     1,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-organization-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization/1/user?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetOrganizationUsersWhenTheRequestMethodIsIncorrect",
			organizationID:     1,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-organization-users.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/servicedeskapi/organization/1/user?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationUsersWhenTheStatusCodeIsIncorrect",
			organizationID:     1,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-organization-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization/1/user?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationUsersWhenTheContextIsNil",
			organizationID:     1,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-organization-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization/1/user?limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetOrganizationUsersWhenTheEndpointIsEmpty",
			organizationID:     1,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-organization-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetOrganizationUsersWhenTheResponseBodyHasADifferentFormat",
			organizationID:     1,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/organization/1/user?limit=50&start=0",
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

			service := &OrganizationService{client: mockClient}
			gotResult, gotResponse, err := service.Users(testCase.context, testCase.organizationID, testCase.start, testCase.limit)

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

				for _, organization := range gotResult.Values {
					t.Log("-------------------------------------------")
					t.Logf("Organization User Name: %v", organization.Name)
					t.Logf("Organization User Mail: %v", organization.EmailAddress)
					t.Log("-------------------------------------------")
				}
			}

		})
	}

}
