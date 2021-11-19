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

func TestIssueLinkTypeService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		payload            *models.LinkTypeScheme
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateIssueLinksTypeWhenThePayloadIsCorrect",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            false,
		},
		{
			name: "CreateIssueLinksTypeWhenTheResponseBodyHasADifferentFormat",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/invalid-json.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name: "CreateIssueLinksTypeWhenTheContextIsNil",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            nil,
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name: "CreateIssueLinksTypeWhenTheStatusCodeIsIncorrect",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name: "CreateIssueLinksTypeWhenTheMethodIsIncorrect",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "CreateIssueLinksTypeWhenThePayloadIsNil",
			payload:            nil,
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "CreateIssueLinksTypeWhenThePayloadIsEmpty",
			payload:            nil,
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
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
				Headers:            testCase.wantHTTPHeaders,
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

			i := &IssueLinkTypeService{client: mockClient}

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
				assert.Equal(t, gotResult.Name, testCase.payload.Name)
			}
		})

	}

}

func TestIssueLinkTypeService_Get(t *testing.T) {

	testCases := []struct {
		name                   string
		mockFile               string
		IssueLinkTypeServiceID string
		wantHTTPMethod         string
		endpoint               string
		context                context.Context
		wantHTTPHeaders        map[string]string
		wantHTTPCodeReturn     int
		wantErr                bool
	}{
		{
			name:                   "GetIssueLinksTypeWhenTheJSONIsCorrect",
			mockFile:               "../v3/mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/2/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                false,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheContextIsNil",
			mockFile:               "../v3/mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/2/issueLinkType/10000",
			context:                nil,
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheTheResponseBodyHasADifferentFormat",
			mockFile:               "../v3/mocks/invalid-json.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/2/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheJSONIsEmpty",
			mockFile:               "../v3/mocks/empty_json.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/2/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheMockedIDIsDifferent",
			mockFile:               "../v3/mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10001",
			endpoint:               "/rest/api/2/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheMockedIDIsEmpty",
			mockFile:               "../v3/mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "",
			endpoint:               "/rest/api/2/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheHTTPResponseCodeIsOK",
			mockFile:               "../v3/mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/2/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                false,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheHTTPResponseCodeIsNotValid",
			mockFile:               "../v3/mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodPut,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/2/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				Headers:            testCase.wantHTTPHeaders,
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

			i := &IssueLinkTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.IssueLinkTypeServiceID)

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
				assert.Equal(t, gotResult.ID, testCase.IssueLinkTypeServiceID)
			}
		})

	}

}

func TestIssueLinkTypeService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueLinksTypesWhenTheJSONIsCorrect",
			mockFile:           "../v3/mocks/get_issue_link_types.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            false,
		},

		{
			name:               "GetIssueLinksTypesWhenTheContextIsNil",
			mockFile:           "../v3/mocks/get_issue_link_types.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            nil,
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},

		{
			name:               "GetIssueLinksTypesWhenTheStatusCodeIsIncorrect",
			mockFile:           "../v3/mocks/get_issue_link_types.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheRequestMethodIsIncorrect",
			mockFile:           "../v3/mocks/get_issue_link_types.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheEndpointProvidedIsInvalid",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issueLinkTypeasasdadsads",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheEndpointProvidedIsEmpty",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheEndpointProvidedIsANumber",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "111",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheHTTPResponseCodeIsInvalid",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheHTTPResponseCodeIsDifferent",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPCodeReturn: 499,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheHTTPMethodRequestIsNotGET",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
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
				Headers:            testCase.wantHTTPHeaders,
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

			i := &IssueLinkTypeService{client: mockClient}

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

func TestIssueLinkTypeService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		issueLinkTypeID    string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteIssueLinkTypeWhenTheIDIsCorrect",
			issueLinkTypeID:    "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheIssueLinkTypeIDIsNotProvided",
			issueLinkTypeID:    "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheIDIsIncorrect",
			issueLinkTypeID:    "10002",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheEndpointIsIncorrect",
			issueLinkTypeID:    "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issueLinkTypes/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheContextIsNil",
			issueLinkTypeID:    "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheRequestMethodIsIncorrect",
			issueLinkTypeID:    "10001",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheStatusCodeIsIncorrect",
			issueLinkTypeID:    "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issueLinkType/10001",
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
				Headers:            testCase.wantHTTPHeaders,
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

			i := &IssueLinkTypeService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.issueLinkTypeID)

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

func TestIssueLinkTypeService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		issueLinkTypeID    string
		payload            *models.LinkTypeScheme
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:            "UpdateIssueLinksTypeWhenThePayloadIsCorrect",
			issueLinkTypeID: "10001",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheIssueLinkTypeIDIsNotProvided",
			issueLinkTypeID: "",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheIDInsIncorrect",
			issueLinkTypeID: "10000",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheContextIsNil",
			issueLinkTypeID: "10001",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "UpdateIssueLinksTypeWhenThePayloadIsNil",
			issueLinkTypeID:    "10001",
			payload:            nil,
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheRequestMethodIsIncorrect",
			issueLinkTypeID: "10001",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheStatusCodeIsIncorrect",
			issueLinkTypeID: "10001",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheResponseBodyHasADifferentFormat",
			issueLinkTypeID: "10001",
			payload: &models.LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "../v3/mocks/invalid-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issueLinkType/10001",
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
				Headers:            testCase.wantHTTPHeaders,
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

			i := &IssueLinkTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.issueLinkTypeID, testCase.payload)

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
