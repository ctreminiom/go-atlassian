package jira

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIssueLinkTypeService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		payload            *IssueLinkTypePayloadScheme
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateIssueLinksTypeWhenThePayloadIsCorrect",
			payload: &IssueLinkTypePayloadScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            false,
		},
		{
			name:               "CreateIssueLinksTypeWhenThePayloadIsNil",
			payload:            nil,
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name: "CreateIssueLinksTypeWhenThePayloadIsEmpty",
			payload: &IssueLinkTypePayloadScheme{
				Inward:  "",
				Name:    "",
				Outward: "",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
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
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

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
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                false,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheJSONIsEmpty",
			mockFile:               "./mocks/empty_json.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheMockedIDIsDifferent",
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10001",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheMockedIDIsEmpty",
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheHTTPResponseCodeIsOK",
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                false,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheHTTPResponseCodeIsNotValid",
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodPut,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/3/issueLinkType/10000",
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
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

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
			mockFile:           "./mocks/get_issue_link_types.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            false,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONIsEmpty",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheEndpointProvidedIsInvalid",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkTypeasasdadsads",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheEndpointProvidedIsEmpty",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheEndpointProvidedIsANumber",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "111",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheHTTPResponseCodeIsInvalid",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheHTTPResponseCodeIsDifferent",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: 499,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheHTTPMethodRequestIsNotGET",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
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
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}
		})

	}

}
