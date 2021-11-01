package v3

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestFieldConfigurationService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		IDs                []int
		isDefault          bool
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
			name:               "GetsFieldConfigurationsWhenTheParametersAreCorrect",
			IDs:                []int{10000, 100001},
			isDefault:          false,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetsFieldConfigurationsWhenTheEndpointIsIncorrect",
			IDs:                []int{10000, 100001},
			isDefault:          false,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration?id=10000&id=100001&maxResults=50&startAts=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsWhenTheContextIsNil",
			IDs:                []int{10000, 100001},
			isDefault:          false,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsWhenTheIDsAreNil",
			IDs:                nil,
			isDefault:          false,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetsFieldConfigurationsWhenTheRequestMethodIsIncorrect",
			IDs:                []int{10000, 100001},
			isDefault:          false,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsWhenTheStatusCodeIsIncorrect",
			IDs:                []int{10000, 100001},
			isDefault:          false,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsWhenTheDefaultOptionIsSet",
			IDs:                []int{10000, 100001},
			isDefault:          true,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration?id=10000&id=100001&isDefault=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetsFieldConfigurationsWhenTheResponseBodyHasADifferentFormat",
			IDs:                []int{10000, 100001},
			isDefault:          false,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsWhenTheResponseBodyLengthIsZero",
			IDs:                []int{10000, 100001},
			isDefault:          false,
			startAt:            0,
			maxResult:          50,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
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

			service := &FieldConfigurationService{client: mockClient}
			getResult, gotResponse, err := service.Gets(testCase.context, testCase.IDs, testCase.isDefault, testCase.startAt, testCase.maxResult)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func TestFieldConfigurationService_IssueTypeItems(t *testing.T) {

	testCases := []struct {
		name               string
		fieldConfigIDs     []int
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
			name:               "GetsFieldConfigurationsIssueTypeItemsWhenTheParametersAreCorrect",
			fieldConfigIDs:     []int{10000, 10001},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-issue-field-configuration-issue-type-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=10000&fieldConfigurationSchemeId=10001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsFieldConfigurationsIssueTypeItemsWhenTheParametersAreIncorrect",
			fieldConfigIDs:     []int{10000, 10002},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-issue-field-configuration-issue-type-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=10000&fieldConfigurationSchemeId=10001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsIssueTypeItemsWhenTheEndpointIsIncorrect",
			fieldConfigIDs:     []int{10000, 10001},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-issue-field-configuration-issue-type-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mappings?fieldConfigurationSchemeId=10000&fieldConfigurationSchemeId=10001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsIssueTypeItemsWhenTheContextIsNil",
			fieldConfigIDs:     []int{10000, 10001},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-issue-field-configuration-issue-type-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=10000&fieldConfigurationSchemeId=10001&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsIssueTypeItemsWhenTheFieldConfigIDsAreNil",
			fieldConfigIDs:     nil,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-issue-field-configuration-issue-type-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsFieldConfigurationsIssueTypeItemsWhenTheRequestMethodIsIncorrect",
			fieldConfigIDs:     []int{10000, 10001},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-issue-field-configuration-issue-type-items.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=10000&fieldConfigurationSchemeId=10001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsIssueTypeItemsWhenTheStatusCodeIsIncorrect",
			fieldConfigIDs:     []int{10000, 10001},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-issue-field-configuration-issue-type-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=10000&fieldConfigurationSchemeId=10001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsIssueTypeItemsWhenTheResponseBodyLengthIsZero",
			fieldConfigIDs:     []int{10000, 10001},
			startAt:            0,
			maxResult:          50,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=10000&fieldConfigurationSchemeId=10001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsIssueTypeItemsWhenTheResponseBodyHasADifferentFormat",
			fieldConfigIDs:     []int{10000, 10001},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=10000&fieldConfigurationSchemeId=10001&maxResults=50&startAt=0",
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

			service := &FieldConfigurationService{client: mockClient}
			getResult, gotResponse, err := service.IssueTypeItems(testCase.context, testCase.fieldConfigIDs, testCase.startAt, testCase.maxResult)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestFieldConfigurationService_Items(t *testing.T) {

	testCases := []struct {
		name               string
		fieldConfigID      int
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
			name:               "GetsFieldConfigurationsItemsWhenTheParametersAreCorrect",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetsFieldConfigurationsItemsWhenTheParametersAreIncorrect",
			fieldConfigID:      10001,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsItemsWhenTheEndpointIsIncorrect",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurations/10000/fields?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsItemsWhenTheContextIsNil",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsItemsWhenTheRequestMethodIsIncorrect",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-items.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsItemsWhenTheStatusCodeIsIncorrect",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsItemsWhenTheResponseBodyLengthIsZero",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsItemsWhenTheResponseBodyHasADifferentFormat",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
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

			service := &FieldConfigurationService{client: mockClient}
			getResult, gotResponse, err := service.Items(testCase.context, testCase.fieldConfigID, testCase.startAt, testCase.maxResult)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestFieldConfigurationService_Schemes(t *testing.T) {

	testCases := []struct {
		name               string
		ids                []int
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
			name:               "GetsFieldConfigurationsItemsWhenTheParametersAreCorrect",
			ids:                []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=10000&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetsFieldConfigurationsItemsWhenTheParametersAreIncorrect",
			ids:                []int{10001},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=10000&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsItemsWhenTheEndpointIsIncorrect",
			ids:                []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationschemes?id=10000&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsItemsWhenTheContextIsNil",
			ids:                []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=10000&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsItemsWhenTheIDsAreNil",
			ids:                nil,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetsFieldConfigurationsItemsWhenTheRequestMethodIsIncorrect",
			ids:                []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=10000&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsItemsWhenTheStatusCodeIsIncorrect",
			ids:                []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=10000&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsItemsWhenTheResponseBodyLengthIsZero",
			ids:                []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=10000&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsFieldConfigurationsItemsWhenTheResponseBodyHasADifferentFormat",
			ids:                []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/empty_json.json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=10000&maxResults=50&startAt=0",
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

			service := &FieldConfigurationService{client: mockClient}
			getResult, gotResponse, err := service.Schemes(testCase.context, testCase.ids, testCase.startAt, testCase.maxResult)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestFieldConfigurationService_SchemesByProject(t *testing.T) {

	testCases := []struct {
		name               string
		projectIDs         []int
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
			name:               "GetsFieldConfigurationsSchemesByProjectWhenTheParametersAreCorrect",
			projectIDs:         []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=50&projectId=10000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsFieldConfigurationsSchemesByProjectWhenTheParametersAreIncorrect",
			projectIDs:         []int{10001},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=50&projectId=10000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsSchemesByProjectWhenTheEndpointIsIncorrect",
			projectIDs:         []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationschemes/project?maxResults=50&projectId=10000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsSchemesByProjectWhenTheEndpointIsEmpty",
			projectIDs:         []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsSchemesByProjectWhenTheContextIsNil",
			projectIDs:         []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=50&projectId=10000&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsSchemesByProjectWhenTheProjectIDsAreNil",
			projectIDs:         nil,
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsFieldConfigurationsSchemesByProjectWhenTheRequestMethodIsIncorrect",
			projectIDs:         []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=50&projectId=10000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsSchemesByProjectWhenTheStatusCodeIsIncorrect",
			projectIDs:         []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=50&projectId=10000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsSchemesByProjectWhenTheResponseBodyLengthIsZero",
			projectIDs:         []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=50&projectId=10000&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsFieldConfigurationsSchemesByProjectWhenTheResponseBodyHasADifferentFormat",
			projectIDs:         []int{10000},
			startAt:            0,
			maxResult:          50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=50&projectId=10000&startAt=0",
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

			service := &FieldConfigurationService{client: mockClient}
			getResult, gotResponse, err := service.SchemesByProject(testCase.context, testCase.projectIDs, testCase.startAt, testCase.maxResult)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestFieldConfigurationService_Assign(t *testing.T) {

	testCases := []struct {
		name                       string
		fieldConfigurationSchemeID string
		projectID                  string
		mockFile                   string
		wantHTTPMethod             string
		endpoint                   string
		context                    context.Context
		wantHTTPCodeReturn         int
		wantErr                    bool
	}{
		{
			name:                       "AssignFieldConfigurationToProjectWhenTheParametersAreCorrect",
			fieldConfigurationSchemeID: "1000",
			projectID:                  "1001",
			wantHTTPMethod:             http.MethodPut,
			endpoint:                   "/rest/api/3/fieldconfigurationscheme/project",
			context:                    context.Background(),
			wantHTTPCodeReturn:         http.StatusNoContent,
			wantErr:                    false,
		},

		{
			name:                       "AssignFieldConfigurationToProjectWhenTheProjectIDIsNotProvided",
			fieldConfigurationSchemeID: "1000",
			projectID:                  "",
			wantHTTPMethod:             http.MethodPut,
			endpoint:                   "/rest/api/3/fieldconfigurationscheme/project",
			context:                    context.Background(),
			wantHTTPCodeReturn:         http.StatusNoContent,
			wantErr:                    true,
		},

		{
			name:                       "AssignFieldConfigurationToProjectWhenTheRequestMethodIsIncorrect",
			fieldConfigurationSchemeID: "1000",
			projectID:                  "1001",
			wantHTTPMethod:             http.MethodHead,
			endpoint:                   "/rest/api/3/fieldconfigurationscheme/project",
			context:                    context.Background(),
			wantHTTPCodeReturn:         http.StatusNoContent,
			wantErr:                    true,
		},

		{
			name:                       "AssignFieldConfigurationToProjectWhenTheStatusCodeIsIncorrect",
			fieldConfigurationSchemeID: "1000",
			projectID:                  "1001",
			wantHTTPMethod:             http.MethodPut,
			endpoint:                   "/rest/api/3/fieldconfigurationscheme/project",
			context:                    context.Background(),
			wantHTTPCodeReturn:         http.StatusBadRequest,
			wantErr:                    true,
		},

		{
			name:                       "AssignFieldConfigurationToProjectWhenTheContextIsNil",
			fieldConfigurationSchemeID: "1000",
			projectID:                  "1001",
			wantHTTPMethod:             http.MethodPut,
			endpoint:                   "/rest/api/3/fieldconfigurationscheme/project",
			context:                    nil,
			wantHTTPCodeReturn:         http.StatusNoContent,
			wantErr:                    true,
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

			service := &FieldConfigurationService{client: mockClient}
			gotResponse, err := service.Assign(testCase.context, testCase.fieldConfigurationSchemeID, testCase.projectID)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}
