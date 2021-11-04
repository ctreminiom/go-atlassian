package v2

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestIssueTypeSchemeService_AddIssueTypes(t *testing.T) {

	testCases := []struct {
		name               string
		issueTypeSchemeID  int
		issueTypeIDs       []int
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddIssueTypesToIssueTypeSchemeWhenTheParametersAreCorrect",
			issueTypeSchemeID:  1000,
			issueTypeIDs:       []int{10001, 10002},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetypescheme/1000/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "AddIssueTypesToIssueTypeSchemeWhenTheIssueTypeIDsValueIsEmpty",
			issueTypeSchemeID:  1000,
			wantHTTPMethod:     http.MethodPut,
			issueTypeIDs:       nil,
			endpoint:           "/rest/api/2/issuetypescheme/1000/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddIssueTypesToIssueTypeSchemeWhenTheEndpointIsIncorrect",
			issueTypeSchemeID:  1000,
			issueTypeIDs:       []int{10001, 10002},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetypescheme/1000/issuetypes",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddIssueTypesToIssueTypeSchemeWhenTheRequestMethodIsIncorrect",
			issueTypeSchemeID:  1000,
			issueTypeIDs:       []int{10001, 10002},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme/1000/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddIssueTypesToIssueTypeSchemeWhenTheStatusCodeIsIncorrect",
			issueTypeSchemeID:  1000,
			issueTypeIDs:       []int{10001, 10002},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetypescheme/1000/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AddIssueTypesToIssueTypeSchemeWhenTheContextIsNil",
			issueTypeSchemeID:  1000,
			issueTypeIDs:       []int{10001, 10002},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetypescheme/1000/issuetype",
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

			i := &IssueTypeSchemeService{client: mockClient}

			gotResponse, err := i.Append(testCase.context, testCase.issueTypeSchemeID, testCase.issueTypeIDs)

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

func TestIssueTypeSchemeService_Assign(t *testing.T) {

	testCases := []struct {
		name               string
		issueTypeSchemeID  string
		projectID          string
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AssignIssueTypeSchemeToProjectWhenTheParametersAreCorrect",
			issueTypeSchemeID:  "10001",
			projectID:          "10001",
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "",
			endpoint:           "/rest/api/2/issuetypescheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "AssignIssueTypeSchemeToProjectWhenTheContextIsNil",
			issueTypeSchemeID:  "10001",
			projectID:          "10001",
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "",
			endpoint:           "/rest/api/2/issuetypescheme/project",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignIssueTypeSchemeToProjectWhenTheProjectIDIsNotProvided",
			issueTypeSchemeID:  "10001",
			projectID:          "",
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "",
			endpoint:           "/rest/api/2/issuetypescheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignIssueTypeSchemeToProjectWhenTheIssueTypeSchemeIDIsNotProvided",
			issueTypeSchemeID:  "",
			projectID:          "10001",
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "",
			endpoint:           "/rest/api/2/issuetypescheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignIssueTypeSchemeToProjectWhenTheEndpointDoesNotHaveACorrectFormat",
			issueTypeSchemeID:  "10001",
			projectID:          "10001",
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "",
			endpoint:           "est/api/2/issuetypescheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignIssueTypeSchemeToProjectWhenTheEndpointIsIncorrect",
			issueTypeSchemeID:  "10001",
			projectID:          "10001",
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "",
			endpoint:           "/rest/api/2/issuetypescheme/projects",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignIssueTypeSchemeToProjectWhenTheRequestMethodIsIncorrect",
			issueTypeSchemeID:  "10001",
			projectID:          "10001",
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "",
			endpoint:           "/rest/api/2/issuetypescheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignIssueTypeSchemeToProjectWhenTheStatusCodeIsIncorrect",
			issueTypeSchemeID:  "10001",
			projectID:          "10001",
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "",
			endpoint:           "/rest/api/2/issuetypescheme/project",
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

			i := &IssueTypeSchemeService{client: mockClient}

			gotResponse, err := i.Assign(testCase.context, testCase.issueTypeSchemeID, testCase.projectID)

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

func TestIssueTypeSchemeService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *IssueTypeSchemePayloadScheme
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateIssueTypeSchemeWhenThePayloadIsCorrect",
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			mockFile:           "../v3/mocks/create-issue-type-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issuetypescheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheDefaultIssueTypeIDIsEmpty",
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			mockFile:           "../v3/mocks/create-issue-type-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issuetypescheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheDefaultIssueTypeIDParamIsIncorrect",
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "10055",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			mockFile:           "../v3/mocks/create-issue-type-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issuetypescheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssueTypeSchemeWhenThePayloadIsNil",
			payload:            nil,
			mockFile:           "../v3/mocks/create-issue-type-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issuetypescheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheEndpointIsIncorrect",
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			mockFile:           "../v3/mocks/create-issue-type-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issuetypeschemes",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheRequestMethodIsIncorrect",
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			mockFile:           "../v3/mocks/create-issue-type-scheme.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheStatusCodeIsIncorrect",
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			mockFile:           "../v3/mocks/create-issue-type-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issuetypescheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheContextIsNil",
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			mockFile:           "../v3/mocks/create-issue-type-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issuetypescheme",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheResponseBodyHasADifferentFormat",
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issuetypescheme",
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

			i := &IssueTypeSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.payload)

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

func TestIssueTypeSchemeService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		issueTypeSchemeID  int
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteIssueTypeSchemeWhenTheIDIsCorrect",
			issueTypeSchemeID:  10001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteIssueTypeSchemeWhenTheIssueTypeSchemeIDIsNotProvided",
			issueTypeSchemeID:  0,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueTypeSchemeWhenTheIDIsIncorrect",
			issueTypeSchemeID:  10000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueTypeSchemeWhenTheContextIsNil",
			issueTypeSchemeID:  10001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme/10001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueTypeSchemeWhenTheRequestMethodIsIncorrect",
			issueTypeSchemeID:  10001,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issuetypescheme/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueTypeSchemeWhenTheStatusCodeIsIncorrect",
			issueTypeSchemeID:  10001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme/10001",
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

			i := &IssueTypeSchemeService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.issueTypeSchemeID)

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

func TestIssueTypeSchemeService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		issueTypeSchemeIDs []int
		statAt             int
		maxResults         int
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetsIssueTypeSchemesWhenThePayloadIsCorrect",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-schemes.json",
			endpoint:           "/rest/api/2/issuetypescheme?id=10000&id=10001&id=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "GetsIssueTypeSchemesWhenTheSchemeIDsAreNotSet",
			issueTypeSchemeIDs: nil,
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-schemes.json",
			endpoint:           "/rest/api/2/issuetypescheme?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "GetsIssueTypeSchemesWhenTheEndpointIsIncorrect",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-schemes.json",
			endpoint:           "/rest/api/2/issuetypeschemes?id=10000&id=10001&id=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "GetsIssueTypeSchemesWhenTheContextIsNil",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-schemes.json",
			endpoint:           "/rest/api/2/issuetypescheme?id=10000&id=10001&id=10002&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "GetsIssueTypeSchemesWhenTheRequestMethodIsIncorrect",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "../v3/mocks/get-issue-type-schemes.json",
			endpoint:           "/rest/api/2/issuetypescheme?id=10000&id=10001&id=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "GetsIssueTypeSchemesWhenTheStatusCodeIsIncorrect",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-schemes.json",
			endpoint:           "/rest/api/2/issuetypescheme?id=10000&id=10001&id=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemesWhenTheResponseBodyHasADifferentFormat",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/empty_json.json",
			endpoint:           "/rest/api/2/issuetypescheme?id=10000&id=10001&id=10002&maxResults=50&startAt=0",
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

			i := &IssueTypeSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.issueTypeSchemeIDs, testCase.statAt, testCase.maxResults)

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

func TestIssueTypeSchemeService_Items(t *testing.T) {

	testCases := []struct {
		name               string
		issueTypeSchemeIDs []int
		statAt             int
		maxResults         int
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetsIssueTypeSchemeItemsWhenThePayloadIsCorrect",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-scheme-items.json",
			endpoint:           "/rest/api/2/issuetypescheme/mapping?issueTypeSchemeId=10000&issueTypeSchemeId=10001&issueTypeSchemeId=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsIssueTypeSchemeItemsWhenTheIssueTypeSchemeIDsParamIsNil",
			issueTypeSchemeIDs: nil,
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-scheme-items.json",
			endpoint:           "/rest/api/2/issuetypescheme/mapping?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsIssueTypeSchemeItemsWhenTheEndpointIsIncorrect",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-scheme-items.json",
			endpoint:           "/rest/api/2/issuetypeschemes/mapping?issueTypeSchemeId=10000&issueTypeSchemeId=10001&issueTypeSchemeId=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemeItemsWhenTheContextIsNil",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-scheme-items.json",
			endpoint:           "/rest/api/2/issuetypescheme/mapping?issueTypeSchemeId=10000&issueTypeSchemeId=10001&issueTypeSchemeId=10002&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemeItemsWhenTheRequestMethodIsIncorrect",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "../v3/mocks/get-issue-type-scheme-items.json",
			endpoint:           "/rest/api/2/issuetypescheme/mapping?issueTypeSchemeId=10000&issueTypeSchemeId=10001&issueTypeSchemeId=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemeItemsWhenTheStatusCodeIsIncorrect",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-scheme-items.json",
			endpoint:           "/rest/api/2/issuetypescheme/mapping?issueTypeSchemeId=10000&issueTypeSchemeId=10001&issueTypeSchemeId=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemeItemsWhenTheContextIsNil",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-type-scheme-items.json",
			endpoint:           "/rest/api/2/issuetypescheme/mapping?issueTypeSchemeId=10000&issueTypeSchemeId=10001&issueTypeSchemeId=10002&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemeItemsWhenTheResponseBodyHasADifferentFormat",
			issueTypeSchemeIDs: []int{10000, 10001, 10002},
			statAt:             0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/empty_json.json",
			endpoint:           "/rest/api/2/issuetypescheme/mapping?issueTypeSchemeId=10000&issueTypeSchemeId=10001&issueTypeSchemeId=10002&maxResults=50&startAt=0",
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

			i := &IssueTypeSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Items(testCase.context, testCase.issueTypeSchemeIDs, testCase.statAt, testCase.maxResults)

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

func TestIssueTypeSchemeService_Projects(t *testing.T) {

	testCases := []struct {
		name               string
		projectIDs         []int
		startAt            int
		maxResults         int
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetsIssueTypeSchemesUsedByAnyProjectsWhenTheParametersAreCorrect",
			projectIDs:         []int{10000, 10001, 10002},
			startAt:            0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-types-schemes-for-projects.json",
			endpoint:           "/rest/api/2/issuetypescheme/project?maxResults=50&projectId=10000&projectId=10001&projectId=10002&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsIssueTypeSchemesUsedByAnyProjectsWhenTheProjectIDsParamIsNil",
			projectIDs:         nil,
			startAt:            0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-types-schemes-for-projects.json",
			endpoint:           "/rest/api/2/issuetypescheme/project?maxResults=50&projectId=10000&projectId=10001&projectId=10002&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemesUsedByAnyProjectsWhenTheEndpointIsIncorrect",
			projectIDs:         []int{10000, 10001, 10002},
			startAt:            0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-types-schemes-for-projects.json",
			endpoint:           "/rest/api/2/issuetypeschemes/project?maxResults=50&projectId=10000&projectId=10001&projectId=10002&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemesUsedByAnyProjectsWhenTheRequestMethodIsIncorrect",
			projectIDs:         []int{10000, 10001, 10002},
			startAt:            0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "../v3/mocks/get-issue-types-schemes-for-projects.json",
			endpoint:           "/rest/api/2/issuetypescheme/project?maxResults=50&projectId=10000&projectId=10001&projectId=10002&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemesUsedByAnyProjectsWhenTheStatusCodeIsIncorrect",
			projectIDs:         []int{10000, 10001, 10002},
			startAt:            0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-types-schemes-for-projects.json",
			endpoint:           "/rest/api/2/issuetypescheme/project?maxResults=50&projectId=10000&projectId=10001&projectId=10002&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemesUsedByAnyProjectsWhenTheContextIsNil",
			projectIDs:         []int{10000, 10001, 10002},
			startAt:            0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-types-schemes-for-projects.json",
			endpoint:           "/rest/api/2/issuetypescheme/project?maxResults=50&projectId=10000&projectId=10001&projectId=10002&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueTypeSchemesUsedByAnyProjectsWhenTheResponseBodyHasADifferentFormat",
			projectIDs:         []int{10000, 10001, 10002},
			startAt:            0,
			maxResults:         50,
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/empty_json.json",
			endpoint:           "/rest/api/2/issuetypescheme/project?maxResults=50&projectId=10000&projectId=10001&projectId=10002&startAt=0",
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

			i := &IssueTypeSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Projects(testCase.context, testCase.projectIDs, testCase.startAt, testCase.maxResults)

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

func TestIssueTypeSchemeService_RemoveIssueType(t *testing.T) {

	testCases := []struct {
		name               string
		issueTypeSchemeID  int
		issueTypeID        int
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "RemoveIssueTypesToFromIssueTypeSchemeWhenTheParametersAreCorrect",
			issueTypeSchemeID:  1000,
			issueTypeID:        10001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme/1000/issuetype/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "RemoveIssueTypesToFromIssueTypeSchemeWhenTheParametersAreIncorrect",
			issueTypeSchemeID:  1000,
			issueTypeID:        10000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme/1000/issuetype/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveIssueTypesToFromIssueTypeSchemeWhenTheContextIsNil",
			issueTypeSchemeID:  1000,
			issueTypeID:        10001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme/1000/issuetype/10001",
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

			i := &IssueTypeSchemeService{client: mockClient}

			gotResponse, err := i.Remove(testCase.context, testCase.issueTypeSchemeID, testCase.issueTypeID)

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

func TestIssueTypeSchemeService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		issueTypeSchemeID  int
		payload            *IssueTypeSchemePayloadScheme
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:              "UpdateIssueTypeSchemeWhenThePayloadIsCorrect",
			issueTypeSchemeID: 12,
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetypescheme/12",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:              "UpdateIssueTypeSchemeWhenTheContextIsNil",
			issueTypeSchemeID: 12,
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetypescheme/12",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:              "UpdateIssueTypeSchemeWhenTheIssueTypeSchemeIDIsNotProvided",
			issueTypeSchemeID: 0,
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetypescheme/12",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "UpdateIssueTypeSchemeWhenThePayloadIsNil",
			issueTypeSchemeID:  12,
			payload:            nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetypescheme/12",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:              "UpdateIssueTypeSchemeWhenTheRequestMethodIsIncorrect",
			issueTypeSchemeID: 12,
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetypescheme/12",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:              "UpdateIssueTypeSchemeWhenTheStatusCodeIsIncorrect",
			issueTypeSchemeID: 12,
			payload: &IssueTypeSchemePayloadScheme{
				DefaultIssueTypeID: "1000",
				IssueTypeIds:       []string{"1000", "1001", "1002"},
				Name:               "Default Issue Type Scheme",
				Description:        "Issue Type Scheme description",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetypescheme/12",
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

			i := &IssueTypeSchemeService{client: mockClient}

			gotResponse, err := i.Update(testCase.context, testCase.issueTypeSchemeID, testCase.payload)

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
