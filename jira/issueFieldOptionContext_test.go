package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestFieldOptionContextService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		payload            *CreateCustomFieldOptionPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:      "GetsFieldContextsWhenTheParametersAreCorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:      "GetsFieldContextsWhenTheContextIsNil",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheFieldIDIncorrect",
			fieldID:   "100002",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheContextIDIsIncorrect",
			fieldID:   "100001",
			contextID: 01110,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheFieldIDIsEmpty",
			fieldID:   "",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheEndpointIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/fields/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "GetsFieldContextsWhenThePayloadIsNil",
			fieldID:            "100001",
			contextID:          01111,
			payload:            nil,
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheRequestMethodIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheStatusCodeIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheTheResponseBodyLengthIsZero",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:      "GetsFieldContextsWhenTheTheResponseBodyHasADifferentFormat",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/invalid-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
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

			service := &FieldOptionContextService{client: mockClient}
			getResult, gotResponse, err := service.Create(testCase.context, testCase.fieldID, testCase.contextID, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
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

func TestFieldOptionContextService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		optionID           int
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteFieldContextsWhenTheParametersAreCorrect",
			fieldID:            "0001",
			contextID:          0,
			optionID:           0,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/0001/context/0/option/0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteFieldContextsWhenTheFieldIDIsEmpty",
			fieldID:            "",
			contextID:          0,
			optionID:           0,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/0001/context/0/option/0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteFieldContextsWhenTheContextIsNil",
			fieldID:            "0001",
			contextID:          0,
			optionID:           0,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/0001/context/0/option/0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteFieldContextsWhenTheRequestMethodIsIncorrect",
			fieldID:            "0001",
			contextID:          0,
			optionID:           0,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/0001/context/0/option/0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteFieldContextsWhenTheEndpointIsIncorrect",
			fieldID:            "0001",
			contextID:          0,
			optionID:           0,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/0001/contexts/0/option/0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteFieldContextsWhenTheStatusCodeIsIncorrect",
			fieldID:            "0001",
			contextID:          0,
			optionID:           0,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/0001/context/0/option/0",
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

			service := &FieldOptionContextService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.fieldID, testCase.contextID, testCase.optionID)

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

func TestFieldOptionContextService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		opts               *FieldOptionContextParams
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
			opts: &FieldOptionContextParams{
				FieldID:     "100001",
				ContextID:   10001,
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100001/context/10001/option?contextId=10001&fieldId=100001&maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name: "GetsFieldContextsWhenTheParametersAreIncorrect",
			opts: &FieldOptionContextParams{
				FieldID:     "100000",
				ContextID:   10001,
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100001/context/10001/option?contextId=10001&fieldId=100001&maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldContextsWhenTheParametersAreNil",
			opts:               nil,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100001/context/10001/option?contextId=10001&fieldId=100001&maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "GetsFieldContextsWhenTheEndpointIsIncorrect",
			opts: &FieldOptionContextParams{
				FieldID:     "100001",
				ContextID:   10001,
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fields/100001/context/10001/option?contextId=10001&fieldId=100001&maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetsFieldContextsWhenTheContextIsNil",
			opts: &FieldOptionContextParams{
				FieldID:     "100001",
				ContextID:   10001,
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100001/context/10001/option?contextId=10001&fieldId=100001&maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetsFieldContextsWhenTheRequestMethodIsIncorrect",
			opts: &FieldOptionContextParams{
				FieldID:     "100001",
				ContextID:   10001,
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/100001/context/10001/option?contextId=10001&fieldId=100001&maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetsFieldContextsWhenTheStatusCodeIsIncorrect",
			opts: &FieldOptionContextParams{
				FieldID:     "100001",
				ContextID:   10001,
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100001/context/10001/option?contextId=10001&fieldId=100001&maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "GetsFieldContextsWhenTheResponseBodyHasADifferentFormat",
			opts: &FieldOptionContextParams{
				FieldID:     "100001",
				ContextID:   10001,
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/invalid-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100001/context/10001/option?contextId=10001&fieldId=100001&maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetsFieldContextsWhenTheResponseBodyLengthIsZero",
			opts: &FieldOptionContextParams{
				FieldID:     "100001",
				ContextID:   10001,
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/100001/context/10001/option?contextId=10001&fieldId=100001&maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
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

			service := &FieldOptionContextService{client: mockClient}
			getResult, gotResponse, err := service.Gets(testCase.context, testCase.opts, testCase.startAt, testCase.maxResult)

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

func TestFieldOptionContextService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		payload            *CreateCustomFieldOptionPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:      "UpdateFieldContextsWhenTheParametersAreCorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:      "UpdateFieldContextsWhenTheContextIDIsNil",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheFieldIDIncorrect",
			fieldID:   "100002",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheContextIDIsIncorrect",
			fieldID:   "100001",
			contextID: 01110,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheFieldIDIsEmpty",
			fieldID:   "",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheEndpointIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fields/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateFieldContextsWhenThePayloadIsNil",
			fieldID:            "100001",
			contextID:          01111,
			payload:            nil,
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheRequestMethodIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheStatusCodeIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheTheResponseBodyLengthIsZero",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:      "UpdateFieldContextsWhenTheTheResponseBodyHasADifferentFormat",
			fieldID:   "100001",
			contextID: 01111,
			payload: &CreateCustomFieldOptionPayloadScheme{Options: []FieldContextOptionValueScheme{
				{Value: "Argentina", Disabled: false, OptionID: "10027"},
				{Value: "Canada", Disabled: false, OptionID: "10027"},
			}},
			mockFile:           "./mocks/invalid-json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/field/100001/context/585/option",
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

			service := &FieldOptionContextService{client: mockClient}
			getResult, gotResponse, err := service.Update(testCase.context, testCase.fieldID, testCase.contextID, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
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
