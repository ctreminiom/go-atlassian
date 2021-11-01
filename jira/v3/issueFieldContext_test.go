package v3

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestFieldContextService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		payload            *FieldContextPayloadScheme
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:    "CreateFieldContextsWhenTheParametersAreCorrect",
			fieldID: "100",
			payload: &FieldContextPayloadScheme{
				IssueTypeIDs: []int{10010},
				ProjectIDs:   nil,
				Name:         "Bug fields context",
				Description:  "A context used to define the custom field options for bugs.",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/create-issue-field-context.json",
			endpoint:           "/rest/api/3/field/100/context",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:    "CreateFieldContextsWhenTheFieldIsInEmpty",
			fieldID: "",
			payload: &FieldContextPayloadScheme{
				IssueTypeIDs: []int{10010},
				ProjectIDs:   nil,
				Name:         "Bug fields context",
				Description:  "A context used to define the custom field options for bugs.",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/create-issue-field-context.json",
			endpoint:           "/rest/api/3/field/100/context",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateFieldContextsWhenTheParametersAreNil",
			fieldID:            "100",
			payload:            nil,
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/create-issue-field-context.json",
			endpoint:           "/rest/api/3/field/100/context",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:    "CreateFieldContextsWhenTheFieldIDIsIncorrect",
			fieldID: "1001",
			payload: &FieldContextPayloadScheme{
				IssueTypeIDs: []int{10010},
				ProjectIDs:   nil,
				Name:         "Bug fields context",
				Description:  "A context used to define the custom field options for bugs.",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/create-issue-field-context.json",
			endpoint:           "/rest/api/3/field/100/context",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:    "CreateFieldContextsWhenTheFieldIDIsEmpty",
			fieldID: "",
			payload: &FieldContextPayloadScheme{
				IssueTypeIDs: []int{10010},
				ProjectIDs:   nil,
				Name:         "Bug fields context",
				Description:  "A context used to define the custom field options for bugs.",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/create-issue-field-context.json",
			endpoint:           "/rest/api/3/field/100/context",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:    "CreateFieldContextsWhenTheContextIsNil",
			fieldID: "100",
			payload: &FieldContextPayloadScheme{
				IssueTypeIDs: []int{10010},
				ProjectIDs:   nil,
				Name:         "Bug fields context",
				Description:  "A context used to define the custom field options for bugs.",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/create-issue-field-context.json",
			endpoint:           "/rest/api/3/field/100/context",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:    "CreateFieldContextsWhenTheProjectIDIsSet",
			fieldID: "100",
			payload: &FieldContextPayloadScheme{
				IssueTypeIDs: []int{10010},
				ProjectIDs:   []int{111111},
				Name:         "Bug fields context",
				Description:  "A context used to define the custom field options for bugs.",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/create-issue-field-context.json",
			endpoint:           "/rest/api/3/field/100/context",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:    "CreateFieldContextsWhenTheRequestMethodIsIncorrect",
			fieldID: "100",
			payload: &FieldContextPayloadScheme{
				IssueTypeIDs: []int{10010},
				ProjectIDs:   nil,
				Name:         "Bug fields context",
				Description:  "A context used to define the custom field options for bugs.",
			},
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "./mocks/create-issue-field-context.json",
			endpoint:           "/rest/api/3/field/100/context",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:    "CreateFieldContextsWhenTheStatusCodeIsIncorrect",
			fieldID: "100",
			payload: &FieldContextPayloadScheme{
				IssueTypeIDs: []int{10010},
				ProjectIDs:   nil,
				Name:         "Bug fields context",
				Description:  "A context used to define the custom field options for bugs.",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/create-issue-field-context.json",
			endpoint:           "/rest/api/3/field/100/context",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:    "CreateFieldContextsWhenTheResponseBodyHasADifferentFormat",
			fieldID: "100",
			payload: &FieldContextPayloadScheme{
				IssueTypeIDs: []int{10010},
				ProjectIDs:   nil,
				Name:         "Bug fields context",
				Description:  "A context used to define the custom field options for bugs.",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/invalid-json.json",
			endpoint:           "/rest/api/3/field/100/context",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:    "CreateFieldContextsWhenTheResponseBodyLengthIsZero",
			fieldID: "100",
			payload: &FieldContextPayloadScheme{
				IssueTypeIDs: []int{10010},
				ProjectIDs:   nil,
				Name:         "Bug fields context",
				Description:  "A context used to define the custom field options for bugs.",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "",
			endpoint:           "/rest/api/3/field/100/context",
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

			service := &FieldContextService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.fieldID, testCase.payload)

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
			}

		})
	}

}

func TestFieldContextService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		opts               *FieldContextOptionsScheme
		startAt            int
		maxResult          int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "GetsFieldContextsWhenTheParametersAreCorrect",
			opts: &FieldContextOptionsScheme{
				IsAnyIssueType:  true,
				IsGlobalContext: true,
				ContextID:       []int{10000},
			},
			fieldID:            "100",
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/issue-field-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name: "GetsFieldContextsWhenTheFieldIDIsEmpty",
			opts: &FieldContextOptionsScheme{
				IsAnyIssueType:  true,
				IsGlobalContext: true,
				ContextID:       []int{10000},
			},
			fieldID:            "",
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/issue-field-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetsFieldContextsWhenTheParametersAreIncorrect",
			opts: &FieldContextOptionsScheme{
				IsAnyIssueType:  true,
				IsGlobalContext: false,
				ContextID:       []int{10000},
			},
			fieldID:            "100",
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/issue-field-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "GetsFieldContextsWhenTheEndpointIsIncorrect",
			opts: &FieldContextOptionsScheme{
				IsAnyIssueType:  true,
				IsGlobalContext: true,
				ContextID:       []int{10000},
			},
			fieldID:            "100",
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/issue-field-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fields/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsFieldContextsWhenTheOptionsAreNil",
			opts:               nil,
			fieldID:            "100",
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/issue-field-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "GetsFieldContextsWhenTheFieldIDIsIncorrect",
			opts: &FieldContextOptionsScheme{
				IsAnyIssueType:  true,
				IsGlobalContext: true,
				ContextID:       []int{10000},
			},
			fieldID:            "10011",
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/issue-field-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "GetsFieldContextsWhenTheContextIsNil",
			opts: &FieldContextOptionsScheme{
				IsAnyIssueType:  true,
				IsGlobalContext: true,
				ContextID:       []int{10000},
			},
			fieldID:            "100",
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/issue-field-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetsFieldContextsWhenTheRequestMethodIsIncorrect",
			opts: &FieldContextOptionsScheme{
				IsAnyIssueType:  true,
				IsGlobalContext: true,
				ContextID:       []int{10000},
			},
			fieldID:            "100",
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/issue-field-contexts.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetsFieldContextsWhenTheStatusCodeIsIncorrect",
			opts: &FieldContextOptionsScheme{
				IsAnyIssueType:  true,
				IsGlobalContext: true,
				ContextID:       []int{10000},
			},
			fieldID:            "100",
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/issue-field-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "GetsFieldContextsWhenTheResponseBodyHasADifferentFormat",
			opts: &FieldContextOptionsScheme{
				IsAnyIssueType:  true,
				IsGlobalContext: true,
				ContextID:       []int{10000},
			},
			fieldID:            "100",
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/invalid-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetsFieldContextsWhenTheResponseBodyLengthIsZero",
			opts: &FieldContextOptionsScheme{
				IsAnyIssueType:  true,
				IsGlobalContext: true,
				ContextID:       []int{10000},
			},
			fieldID:            "100",
			startAt:            0,
			maxResult:          50,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100/context?contextId=10000&isAnyIssueType=true&isGlobalContext=true&maxResults=50&startAt=0",
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

			service := &FieldContextService{client: mockClient}
			getResult, gotResponse, err := service.Gets(testCase.context, testCase.fieldID, testCase.opts, testCase.startAt, testCase.maxResult)

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
				assert.NotEqual(t, getResult, nil)

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

func TestFieldContextService_AddIssueTypes(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		issueTypesIDs      []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddIssueTypesToFieldContextWhenTheParametersAreCorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "AddIssueTypesToFieldContextWhenTheFieldIDIsNotSet",
			fieldID:            "",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddIssueTypesToFieldContextWhenTheIssueTypeIDsAreNotSet",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddIssueTypesToFieldContextWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002"},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddIssueTypesToFieldContextWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "AddIssueTypesToFieldContextWhenTheEndpointIsEmpty",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddIssueTypesToFieldContextWhenTheContextIsNil",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype",
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

			service := &FieldContextService{client: mockClient}
			gotResponse, err := service.AddIssueTypes(testCase.context, testCase.fieldID, testCase.contextID, testCase.issueTypesIDs)

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

func TestFieldContextService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteCustomFieldContextWhenTheParametersAreCorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteCustomFieldContextWhenTheContextIDIsNotProvided",
			fieldID:            "customfield_10002",
			contextID:          0,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteCustomFieldContextWhenTheFieldIsNotSet",
			fieldID:            "",
			contextID:          2001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteCustomFieldContextWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteCustomFieldContextWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "DeleteCustomFieldContextWhenTheEndpointIsEmpty",
			fieldID:            "customfield_10002",
			contextID:          2001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteCustomFieldContextWhenTheContextIsNil",
			fieldID:            "customfield_10002",
			contextID:          2001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
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

			service := &FieldContextService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.fieldID, testCase.contextID)

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

func TestFieldContextService_GetDefaultValues(t *testing.T) {

	testCases := []struct {
		name                string
		fieldID             string
		contextIDs          []int
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetCustomFieldContextDefaultValuesWhenTheParametersAreCorrect",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-custom-field-default-values.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomFieldContextDefaultValuesWhenTheContextIDsAreNotSet",
			fieldID:            "customfield_10002",
			contextIDs:         nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-custom-field-default-values.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomFieldContextDefaultValuesWhenTheFieldIDIsNotSet",
			fieldID:            "",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-custom-field-default-values.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomFieldContextDefaultValuesWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-custom-field-default-values.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomFieldContextDefaultValuesWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-custom-field-default-values.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetCustomFieldContextDefaultValuesWhenTheEndpointIsEmpty",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-custom-field-default-values.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomFieldContextDefaultValuesWhenTheContextIsNil",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-custom-field-default-values.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomFieldContextDefaultValuesWhenResponseBodyIsEmpty",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue?contextId=1001&contextId=1002&maxResults=50&startAt=0",
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

			service := &FieldContextService{client: mockClient}
			gotResult, gotResponse, err := service.GetDefaultValues(testCase.context, testCase.fieldID, testCase.contextIDs,
				testCase.startAt, testCase.maxResults)

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

				for _, contextData := range gotResult.Values {

					t.Log("------------------------------")
					t.Logf("Field Type: %v", contextData.Type)
					t.Logf("Context ID : %v", contextData.ContextID)
					t.Logf("Context Default Value: %v", contextData.OptionID)
					t.Log("------------------------------ \n")

				}

			}

		})
	}

}

func TestFieldContextService_IssueTypesContext(t *testing.T) {

	testCases := []struct {
		name                string
		fieldID             string
		contextIDs          []int
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetIssueTypesContextWhenTheParametersAreCorrect",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-issue-type-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/issuetypemapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTypesContextWhenTheContextIDsAreNotSet",
			fieldID:            "customfield_10002",
			contextIDs:         nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-issue-type-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/issuetypemapping?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTypesContextWhenTheFieldIDIsNotSet",
			fieldID:            "",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-issue-type-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/issuetypemapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypesContextWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-issue-type-contexts.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/issuetypemapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypesContextWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-issue-type-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/issuetypemapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetIssueTypesContextWhenTheEndpointIsEmpty",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-issue-type-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypesContextWhenTheContextIsNil",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-issue-type-contexts.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/issuetypemapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypesContextWhenResponseBodyIsEmpty",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/issuetypemapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
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

			service := &FieldContextService{client: mockClient}
			gotResult, gotResponse, err := service.IssueTypesContext(testCase.context, testCase.fieldID, testCase.contextIDs,
				testCase.startAt, testCase.maxResults)

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

				for _, contextData := range gotResult.Values {

					t.Log("------------------------------")
					t.Logf("IssueTypeID Type: %v", contextData.IssueTypeID)
					t.Logf("Context ID : %v", contextData.ContextID)
					t.Logf("IsAnyIssueType Value: %v", contextData.IsAnyIssueType)
					t.Log("------------------------------ \n")

				}

			}

		})
	}

}

func TestFieldContextService_Projects(t *testing.T) {

	testCases := []struct {
		name                string
		fieldID             string
		contextIDs          []int
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetProjectContextWhenTheParametersAreCorrect",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-project-mapping.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/projectmapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectContextWhenTheContextIDsAreNotSet",
			fieldID:            "customfield_10002",
			contextIDs:         nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-project-mapping.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/projectmapping?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectContextWhenTheFieldIDIsNotSet",
			fieldID:            "",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-project-mapping.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/projectmapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectContextWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-project-mapping.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/projectmapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectContextWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-project-mapping.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/projectmapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetProjectContextWhenTheEndpointIsEmpty",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-project-mapping.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectContextWhenTheContextIsNil",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-field-context-project-mapping.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/projectmapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectContextWhenResponseBodyIsEmpty",
			fieldID:            "customfield_10002",
			contextIDs:         []int{1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10002/context/projectmapping?contextId=1001&contextId=1002&maxResults=50&startAt=0",
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

			service := &FieldContextService{client: mockClient}
			gotResult, gotResponse, err := service.ProjectsContext(testCase.context, testCase.fieldID, testCase.contextIDs,
				testCase.startAt, testCase.maxResults)

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

				for _, contextData := range gotResult.Values {

					t.Log("------------------------------")
					t.Logf("IsGlobalContext Type: %v", contextData.IsGlobalContext)
					t.Logf("Context ID : %v", contextData.ContextID)
					t.Logf("ProjectID Value: %v", contextData.ProjectID)
					t.Log("------------------------------ \n")

				}

			}

		})
	}

}

func TestFieldContextService_Link(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		projectIDs         []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "LinkCustomFieldContextToProjectWhenTheParametersAreCorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "LinkCustomFieldContextToProjectWhenTheFieldIDIsNotSet",
			fieldID:            "",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "LinkCustomFieldContextToProjectWhenTheProjectIDsAreNotSet",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "LinkCustomFieldContextToProjectWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "LinkCustomFieldContextToProjectWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "LinkCustomFieldContextToProjectWhenTheContextIsNil",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "LinkCustomFieldContextToProjectWhenTheEndpointIsEmpty",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPut,
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

			service := &FieldContextService{client: mockClient}
			gotResponse, err := service.Link(testCase.context, testCase.fieldID, testCase.contextID, testCase.projectIDs)

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

func TestFieldContextService_RemoveIssueTypes(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		issueTypesIDs      []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "RemoveIssueTypesFromFieldContextWhenTheParametersAreCorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype/remove",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "RemoveIssueTypesFromFieldContextWhenTheFieldIDIsNotSet",
			fieldID:            "",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype/remove",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveIssueTypesFromFieldContextWhenTheIssueTypeIDsAreNotSet",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      nil,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype/remove",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveIssueTypesFromFieldContextWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype/remove",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveIssueTypesFromFieldContextWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype/remove",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "RemoveIssueTypesFromFieldContextWhenTheEndpointIsEmpty",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveIssueTypesFromFieldContextWhenTheContextIsNil",
			fieldID:            "customfield_10002",
			contextID:          2001,
			issueTypesIDs:      []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/issuetype/remove",
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

			service := &FieldContextService{client: mockClient}
			gotResponse, err := service.RemoveIssueTypes(testCase.context, testCase.fieldID, testCase.contextID, testCase.issueTypesIDs)

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

func TestFieldContextService_SetDefaultValue(t *testing.T) {

	var defaultValuesMockedWithValues = &FieldContextDefaultPayloadScheme{
		DefaultValues: []*CustomFieldDefaultValueScheme{
			{
				ContextID: "10138",
				OptionID:  "10022",
				Type:      "option.single",
			},
		},
	}

	var defaultValuesMockedWithNotValues = &FieldContextDefaultPayloadScheme{}

	testCases := []struct {
		name               string
		fieldID            string
		payload            *FieldContextDefaultPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "SetDefaultValueToCustomFieldContextWhenTheParametersAreCorrect",
			fieldID:            "customfield_10002",
			payload:            defaultValuesMockedWithValues,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "SetDefaultValueToCustomFieldContextWhenTheValuesAreEmpty",
			fieldID:            "customfield_10002",
			payload:            defaultValuesMockedWithNotValues,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "SetDefaultValueToCustomFieldContextWhenTheFieldIDIsNotSet",
			fieldID:            "",
			payload:            defaultValuesMockedWithValues,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "SetDefaultValueToCustomFieldContextWhenTheValuesAreNotSet",
			fieldID:            "customfield_10002",
			payload:            nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "SetDefaultValueToCustomFieldContextWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10002",
			payload:            defaultValuesMockedWithValues,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "SetDefaultValueToCustomFieldContextWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10002",
			payload:            defaultValuesMockedWithValues,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "SetDefaultValueToCustomFieldContextWhenTheContextIsNil",
			fieldID:            "customfield_10002",
			payload:            defaultValuesMockedWithValues,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/defaultValue",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "SetDefaultValueToCustomFieldContextWhenTheEndpointIsEmpty",
			fieldID:            "customfield_10002",
			payload:            defaultValuesMockedWithValues,
			wantHTTPMethod:     http.MethodPut,
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

			service := &FieldContextService{client: mockClient}
			gotResponse, err := service.SetDefaultValue(testCase.context, testCase.fieldID, testCase.payload)

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

func TestFieldContextService_UnLink(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		projectIDs         []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "UnLinkCustomFieldContextToProjectWhenTheParametersAreCorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project/remove",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "UnLinkCustomFieldContextToProjectWhenTheFieldIDIsNotSet",
			fieldID:            "",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project/remove",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "UnLinkCustomFieldContextToProjectWhenTheProjectIDsAreNotSet",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         nil,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project/remove",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "UnLinkCustomFieldContextToProjectWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project/remove",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "UnLinkCustomFieldContextToProjectWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project/remove",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "UnLinkCustomFieldContextToProjectWhenTheContextIsNil",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001/project/remove",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "UnLinkCustomFieldContextToProjectWhenTheEndpointIsEmpty",
			fieldID:            "customfield_10002",
			contextID:          2001,
			projectIDs:         []string{"10001", "10002", "10003"},
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

			service := &FieldContextService{client: mockClient}
			gotResponse, err := service.UnLink(testCase.context, testCase.fieldID, testCase.contextID, testCase.projectIDs)

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

func TestFieldContextService_Update(t *testing.T) {

	testCases := []struct {
		name                          string
		fieldID                       string
		contextID                     int
		fieldContextName, description string
		mockFile                      string
		wantHTTPMethod                string
		endpoint                      string
		context                       context.Context
		wantHTTPCodeReturn            int
		wantErr                       bool
	}{
		{
			name:               "UpdateCustomFieldContextWhenTheParametersAreCorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			fieldContextName:   "new customfield context",
			description:        "new customfield description",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "UpdateCustomFieldContextWhenTheContextIDIsNotProvided",
			fieldID:            "customfield_10002",
			contextID:          0,
			fieldContextName:   "new customfield context",
			description:        "new customfield description",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "UpdateCustomFieldContextWhenTheFieldIDIsNotSet",
			fieldID:            "",
			contextID:          2001,
			fieldContextName:   "new customfield context",
			description:        "new customfield description",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "UpdateCustomFieldContextWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			fieldContextName:   "new customfield context",
			description:        "new customfield description",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "UpdateCustomFieldContextWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10002",
			contextID:          2001,
			fieldContextName:   "new customfield context",
			description:        "new customfield description",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "UpdateCustomFieldContextWhenTheEndpointIsEmpty",
			fieldID:            "customfield_10002",
			contextID:          2001,
			fieldContextName:   "new customfield context",
			description:        "new customfield description",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "UpdateCustomFieldContextWhenTheContextIsNil",
			fieldID:            "customfield_10002",
			contextID:          2001,
			fieldContextName:   "new customfield context",
			description:        "new customfield description",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/customfield_10002/context/2001",
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

			service := &FieldContextService{client: mockClient}
			gotResponse, err := service.Update(testCase.context, testCase.fieldID, testCase.contextID,
				testCase.fieldContextName, testCase.description)

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
