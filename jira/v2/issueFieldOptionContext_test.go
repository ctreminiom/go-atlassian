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

func TestFieldOptionContextService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		payload            *models.FieldContextOptionListScheme
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
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{

					// Single/Multiple Choice example
					{
						Value:    "Option 2",
						Disabled: false,
					},
					{
						Value:    "Option 4",
						Disabled: false,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:      "GetsFieldContextsWhenTheContextIsNil",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{

					// Single/Multiple Choice example
					{
						Value:    "Option 2",
						Disabled: false,
					},
					{
						Value:    "Option 4",
						Disabled: false,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheFieldIDIncorrect",
			fieldID:   "100002",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{

					// Single/Multiple Choice example
					{
						Value:    "Option 2",
						Disabled: false,
					},
					{
						Value:    "Option 4",
						Disabled: false,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheContextIDIsIncorrect",
			fieldID:   "100001",
			contextID: 01110,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{

					// Single/Multiple Choice example
					{
						Value:    "Option 2",
						Disabled: false,
					},
					{
						Value:    "Option 4",
						Disabled: false,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheFieldIDIsEmpty",
			fieldID:   "",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{

					// Single/Multiple Choice example
					{
						Value:    "Option 2",
						Disabled: false,
					},
					{
						Value:    "Option 4",
						Disabled: false,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheEndpointIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{

					// Single/Multiple Choice example
					{
						Value:    "Option 2",
						Disabled: false,
					},
					{
						Value:    "Option 4",
						Disabled: false,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/fields/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "GetsFieldContextsWhenThePayloadIsNil",
			fieldID:            "100001",
			contextID:          01111,
			payload:            nil,
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheRequestMethodIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{

					// Single/Multiple Choice example
					{
						Value:    "Option 2",
						Disabled: false,
					},
					{
						Value:    "Option 4",
						Disabled: false,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheStatusCodeIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{

					// Single/Multiple Choice example
					{
						Value:    "Option 2",
						Disabled: false,
					},
					{
						Value:    "Option 4",
						Disabled: false,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheTheResponseBodyLengthIsZero",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{

					// Single/Multiple Choice example
					{
						Value:    "Option 2",
						Disabled: false,
					},
					{
						Value:    "Option 4",
						Disabled: false,
					},
				}},
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:      "GetsFieldContextsWhenTheResponseBodyHasADifferentFormat",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{

					// Single/Multiple Choice example
					{
						Value:    "Option 2",
						Disabled: false,
					},
					{
						Value:    "Option 4",
						Disabled: false,
					},
				}},
			mockFile:           "../v3/mocks/invalid-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
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
			contextID:          100,
			optionID:           222,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/field/0001/context/100/option/222",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},
		{
			name:               "DeleteFieldContextsWhenTheOptionIDIsNotProvided",
			fieldID:            "0001",
			contextID:          100,
			optionID:           0,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/field/0001/context/100/option/222",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteFieldContextsWhenTheContextIDIsNotProvided",
			fieldID:            "0001",
			contextID:          0,
			optionID:           222,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/field/0001/context/100/option/222",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteFieldContextsWhenTheFieldIDIsEmpty",
			fieldID:            "",
			contextID:          100,
			optionID:           222,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/field/0001/context/100/option/222",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteFieldContextsWhenTheContextIsNil",
			fieldID:            "0001",
			contextID:          100,
			optionID:           222,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/field/0001/context/100/option/222",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteFieldContextsWhenTheRequestMethodIsIncorrect",
			fieldID:            "0001",
			contextID:          100,
			optionID:           222,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/0001/context/100/option/222",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteFieldContextsWhenTheEndpointIsIncorrect",
			fieldID:            "0001",
			contextID:          100,
			optionID:           222,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/field/0001/context/100/option/000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteFieldContextsWhenTheStatusCodeIsIncorrect",
			fieldID:            "0001",
			contextID:          100,
			optionID:           222,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/field/0001/context/100/option/222",
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

func TestFieldOptionContextService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		FieldID            string
		ContextID          int
		opts               *models.FieldOptionContextParams
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
			name:      "GetsFieldContextsWhenTheParametersAreCorrect",
			FieldID:   "100001",
			ContextID: 10001,
			opts: &models.FieldOptionContextParams{
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/field/100001/context/10001/option?maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsFieldContextsWhenTheParametersAreNil",
			opts:               nil,
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/field/100001/context/10001/option?maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:      "GetsFieldContextsWhenTheEndpointIsIncorrect",
			FieldID:   "100001",
			ContextID: 10001,
			opts: &models.FieldOptionContextParams{
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/field/100001/context/10001/option?maxResults=50&onlyOptions=true&optionId=1000&saaaaaartAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheContextIsNil",
			FieldID:   "100001",
			ContextID: 10001,
			opts: &models.FieldOptionContextParams{
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/field/100001/context/10001/option?maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheRequestMethodIsIncorrect",
			FieldID:   "100001",
			ContextID: 10001,
			opts: &models.FieldOptionContextParams{
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/field/100001/context/10001/option?maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheStatusCodeIsIncorrect",
			FieldID:   "100001",
			ContextID: 10001,
			opts: &models.FieldOptionContextParams{
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/get-custom-field-context-options.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/field/100001/context/10001/option?maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheResponseBodyHasADifferentFormat",
			FieldID:   "100001",
			ContextID: 10001,
			opts: &models.FieldOptionContextParams{
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/invalid-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/field/100001/context/10001/option?maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "GetsFieldContextsWhenTheResponseBodyLengthIsZero",
			FieldID:   "100001",
			ContextID: 10001,
			opts: &models.FieldOptionContextParams{
				OptionID:    1000,
				OnlyOptions: true,
			},
			startAt:            0,
			maxResult:          50,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/field/100001/context/10001/option?maxResults=50&onlyOptions=true&optionId=1000&startAt=0",
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
			getResult, gotResponse, err := service.Gets(testCase.context, testCase.FieldID, testCase.ContextID, testCase.opts, testCase.startAt, testCase.maxResult)

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

func TestFieldOptionContextService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		payload            *models.FieldContextOptionListScheme
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
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{
					{
						ID:       "10064",
						Value:    "Option 2 - Updated",
						Disabled: false,
					},
					{
						ID:       "10065",
						Value:    "Option 4 - Updated",
						Disabled: true,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:      "UpdateFieldContextsWhenTheContextIDIsNil",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{
					{
						ID:       "10064",
						Value:    "Option 2 - Updated",
						Disabled: false,
					},
					{
						ID:       "10065",
						Value:    "Option 4 - Updated",
						Disabled: true,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheFieldIDIncorrect",
			fieldID:   "100002",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{
					{
						ID:       "10064",
						Value:    "Option 2 - Updated",
						Disabled: false,
					},
					{
						ID:       "10065",
						Value:    "Option 4 - Updated",
						Disabled: true,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheContextIDIsIncorrect",
			fieldID:   "100001",
			contextID: 01110,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{
					{
						ID:       "10064",
						Value:    "Option 2 - Updated",
						Disabled: false,
					},
					{
						ID:       "10065",
						Value:    "Option 4 - Updated",
						Disabled: true,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheFieldIDIsEmpty",
			fieldID:   "",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{
					{
						ID:       "10064",
						Value:    "Option 2 - Updated",
						Disabled: false,
					},
					{
						ID:       "10065",
						Value:    "Option 4 - Updated",
						Disabled: true,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheEndpointIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{
					{
						ID:       "10064",
						Value:    "Option 2 - Updated",
						Disabled: false,
					},
					{
						ID:       "10065",
						Value:    "Option 4 - Updated",
						Disabled: true,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/fields/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateFieldContextsWhenThePayloadIsNil",
			fieldID:            "100001",
			contextID:          01111,
			payload:            nil,
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheRequestMethodIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{
					{
						ID:       "10064",
						Value:    "Option 2 - Updated",
						Disabled: false,
					},
					{
						ID:       "10065",
						Value:    "Option 4 - Updated",
						Disabled: true,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheStatusCodeIsIncorrect",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{
					{
						ID:       "10064",
						Value:    "Option 2 - Updated",
						Disabled: false,
					},
					{
						ID:       "10065",
						Value:    "Option 4 - Updated",
						Disabled: true,
					},
				}},
			mockFile:           "../v3/mocks/create-custom-field-context-option.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:      "UpdateFieldContextsWhenTheTheResponseBodyLengthIsZero",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{
					{
						ID:       "10064",
						Value:    "Option 2 - Updated",
						Disabled: false,
					},
					{
						ID:       "10065",
						Value:    "Option 4 - Updated",
						Disabled: true,
					},
				}},
			mockFile:           "",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:      "UpdateFieldContextsWhenTheTheResponseBodyHasADifferentFormat",
			fieldID:   "100001",
			contextID: 01111,
			payload: &models.FieldContextOptionListScheme{
				Options: []*models.CustomFieldContextOptionScheme{
					{
						ID:       "10064",
						Value:    "Option 2 - Updated",
						Disabled: false,
					},
					{
						ID:       "10065",
						Value:    "Option 4 - Updated",
						Disabled: true,
					},
				}},
			mockFile:           "../v3/mocks/invalid-json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/100001/context/585/option",
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

func TestFieldOptionContextService_Order(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		contextID          int
		payload            *models.OrderFieldOptionPayloadScheme
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:      "OrderFieldContextsWhenTheParametersAreCorrect",
			fieldID:   "0001",
			contextID: 100,
			payload: &models.OrderFieldOptionPayloadScheme{
				Position:             "Last",
				CustomFieldOptionIds: []string{"111"},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/0001/context/100/option/move",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},
		{
			name:               "OrderFieldContextsWhenThePayloadIsNotProvided",
			fieldID:            "0001",
			contextID:          100,
			payload:            nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/0001/context/100/option/move",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:      "OrderFieldContextsWhenTheContextIDIsNotProvided",
			fieldID:   "0001",
			contextID: 0,
			payload: &models.OrderFieldOptionPayloadScheme{
				Position:             "Last",
				CustomFieldOptionIds: []string{"111"},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/0001/context/100/option/move",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:      "OrderFieldContextsWhenTheFieldIDIsEmpty",
			fieldID:   "",
			contextID: 100,
			payload: &models.OrderFieldOptionPayloadScheme{
				Position:             "Last",
				CustomFieldOptionIds: []string{"111"},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/0001/context/100/option/move",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:      "OrderFieldContextsWhenTheContextIsNil",
			fieldID:   "0001",
			contextID: 100,
			payload: &models.OrderFieldOptionPayloadScheme{
				Position:             "Last",
				CustomFieldOptionIds: []string{"111"},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/0001/context/100/option/move",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:      "OrderFieldContextsWhenTheRequestMethodIsIncorrect",
			fieldID:   "0001",
			contextID: 100,
			payload: &models.OrderFieldOptionPayloadScheme{
				Position:             "Last",
				CustomFieldOptionIds: []string{"111"},
			},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/field/0001/context/100/option/move",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:      "OrderFieldContextsWhenTheEndpointIsIncorrect",
			fieldID:   "0001",
			contextID: 100,
			payload: &models.OrderFieldOptionPayloadScheme{
				Position:             "Last",
				CustomFieldOptionIds: []string{"111"},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/0001/context/100/option/000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:      "OrderFieldContextsWhenTheStatusCodeIsIncorrect",
			fieldID:   "0001",
			contextID: 100,
			payload: &models.OrderFieldOptionPayloadScheme{
				Position:             "Last",
				CustomFieldOptionIds: []string{"111"},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/field/0001/context/100/option/move",
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
			gotResponse, err := service.Order(testCase.context, testCase.fieldID, testCase.contextID, testCase.payload)

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
