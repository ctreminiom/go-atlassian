package jira

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
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
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
