package v2

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
			mockFile:           "../v3/mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
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
			mockFile:           "../v3/mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration?id=10000&id=100001&maxResults=50&startAts=0",
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
			mockFile:           "../v3/mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
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
			mockFile:           "../v3/mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration?maxResults=50&startAt=0",
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
			mockFile:           "../v3/mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
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
			mockFile:           "../v3/mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
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
			mockFile:           "../v3/mocks/get-field-configurations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration?id=10000&id=100001&isDefault=true&maxResults=50&startAt=0",
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
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
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
			endpoint:           "/rest/api/2/fieldconfiguration?id=10000&id=100001&maxResults=50&startAt=0",
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

func TestFieldConfigurationService_Create(t *testing.T) {

	testCases := []struct {
		name                          string
		fieldConfigurationName        string
		fieldConfigurationDescription string
		mockFile                      string
		wantHTTPMethod                string
		endpoint                      string
		context                       context.Context
		wantHTTPCodeReturn            int
		wantErr                       bool
		expectedError                 string
	}{
		{
			name:                          "CreateFieldConfigurationWhenTheParametersAreCorrect",
			fieldConfigurationName:        "Field Configuration Name Sample",
			fieldConfigurationDescription: "Field Configuration Name Description",
			mockFile:                      "../v3/mocks/get-field-configuration.json",
			wantHTTPMethod:                http.MethodPost,
			endpoint:                      "/rest/api/2/fieldconfiguration",
			context:                       context.Background(),
			wantHTTPCodeReturn:            http.StatusOK,
			wantErr:                       false,
		},

		{
			name:                          "CreateFieldConfigurationWhenTheFieldConfigurationNameIsNotSet",
			fieldConfigurationName:        "",
			fieldConfigurationDescription: "Field Configuration Name Description",
			mockFile:                      "../v3/mocks/get-field-configuration.json",
			wantHTTPMethod:                http.MethodPost,
			endpoint:                      "/rest/api/2/fieldconfiguration",
			context:                       context.Background(),
			wantHTTPCodeReturn:            http.StatusOK,
			wantErr:                       true,
			expectedError:                 "jira: no field configuration name set",
		},

		{
			name:                          "CreateFieldConfigurationWhenTheRequestMethodIsIncorrect",
			fieldConfigurationName:        "Field Configuration Name Sample",
			fieldConfigurationDescription: "Field Configuration Name Description",
			mockFile:                      "../v3/mocks/get-field-configuration.json",
			wantHTTPMethod:                http.MethodHead,
			endpoint:                      "/rest/api/2/fieldconfiguration",
			context:                       context.Background(),
			wantHTTPCodeReturn:            http.StatusOK,
			wantErr:                       true,
			expectedError:                 "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name:                          "CreateFieldConfigurationWhenTheContextIsNotProvided",
			fieldConfigurationName:        "Field Configuration Name Sample",
			fieldConfigurationDescription: "Field Configuration Name Description",
			mockFile:                      "../v3/mocks/get-field-configuration.json",
			wantHTTPMethod:                http.MethodPost,
			endpoint:                      "/rest/api/2/fieldconfiguration",
			context:                       nil,
			wantHTTPCodeReturn:            http.StatusOK,
			wantErr:                       true,
			expectedError:                 "request creation failed: net/http: nil Context",
		},

		{
			name:                          "CreateFieldConfigurationWhenTheResponseBodyIsEmpty",
			fieldConfigurationName:        "Field Configuration Name Sample",
			fieldConfigurationDescription: "Field Configuration Name Description",
			mockFile:                      "../v3/mocks/empty_json.json",
			wantHTTPMethod:                http.MethodPost,
			endpoint:                      "/rest/api/2/fieldconfiguration",
			context:                       context.Background(),
			wantHTTPCodeReturn:            http.StatusOK,
			wantErr:                       true,
			expectedError:                 "unexpected end of JSON input",
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
			getResult, gotResponse, err := service.Create(testCase.context, testCase.fieldConfigurationName, testCase.fieldConfigurationDescription)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

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

func TestFieldConfigurationService_Update(t *testing.T) {

	testCases := []struct {
		name                          string
		fieldConfigurationID          int
		fieldConfigurationName        string
		fieldConfigurationDescription string
		mockFile                      string
		wantHTTPMethod                string
		endpoint                      string
		context                       context.Context
		wantHTTPCodeReturn            int
		wantErr                       bool
		expectedError                 string
	}{
		{
			name:                          "UpdateFieldConfigurationWhenTheParametersAreCorrect",
			fieldConfigurationID:          1000,
			fieldConfigurationName:        "Field Configuration Name Sample",
			fieldConfigurationDescription: "Field Configuration Name Description",
			wantHTTPMethod:                http.MethodPut,
			endpoint:                      "/rest/api/2/fieldconfiguration/1000",
			context:                       context.Background(),
			wantHTTPCodeReturn:            http.StatusNoContent,
			wantErr:                       false,
		},

		{
			name:                          "UpdateFieldConfigurationWhenTheFieldConfigurationIDIsNotSet",
			fieldConfigurationID:          0,
			fieldConfigurationName:        "Field Configuration Name Sample",
			fieldConfigurationDescription: "Field Configuration Name Description",
			wantHTTPMethod:                http.MethodPut,
			endpoint:                      "/rest/api/2/fieldconfiguration/1000",
			context:                       context.Background(),
			wantHTTPCodeReturn:            http.StatusNoContent,
			wantErr:                       true,
			expectedError:                 "jira: no field configuration id set",
		},

		{
			name:                          "UpdateFieldConfigurationWhenTheFieldConfigurationNameIsNotSet",
			fieldConfigurationID:          1000,
			fieldConfigurationName:        "",
			fieldConfigurationDescription: "Field Configuration Name Description",
			wantHTTPMethod:                http.MethodPut,
			endpoint:                      "/rest/api/2/fieldconfiguration/1000",
			context:                       context.Background(),
			wantHTTPCodeReturn:            http.StatusNoContent,
			wantErr:                       true,
			expectedError:                 "jira: no field configuration name set",
		},

		{
			name:                          "UpdateFieldConfigurationWhenTheRequestMethodIsIncorrect",
			fieldConfigurationID:          1000,
			fieldConfigurationName:        "Field Configuration Name Sample",
			fieldConfigurationDescription: "Field Configuration Name Description",
			mockFile:                      "../v3/mocks/get-field-configuration.json",
			wantHTTPMethod:                http.MethodHead,
			context:                       context.Background(),
			wantHTTPCodeReturn:            http.StatusNoContent,
			wantErr:                       true,
			expectedError:                 "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name:                          "UpdateFieldConfigurationWhenTheContextIsNotProvided",
			fieldConfigurationID:          1000,
			fieldConfigurationName:        "Field Configuration Name Sample",
			fieldConfigurationDescription: "Field Configuration Name Description",
			wantHTTPMethod:                http.MethodPut,
			endpoint:                      "/rest/api/2/fieldconfiguration/1000",
			context:                       nil,
			wantHTTPCodeReturn:            http.StatusNoContent,
			wantErr:                       true,
			expectedError:                 "request creation failed: net/http: nil Context",
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
			gotResponse, err := service.Update(testCase.context, testCase.fieldConfigurationID, testCase.fieldConfigurationName, testCase.fieldConfigurationDescription)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

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

func TestFieldConfigurationService_Delete(t *testing.T) {

	testCases := []struct {
		name                 string
		fieldConfigurationID int
		mockFile             string
		wantHTTPMethod       string
		endpoint             string
		context              context.Context
		wantHTTPCodeReturn   int
		wantErr              bool
		expectedError        string
	}{
		{
			name:                 "DeleteFieldConfigurationWhenTheParametersAreCorrect",
			fieldConfigurationID: 1000,
			wantHTTPMethod:       http.MethodDelete,
			endpoint:             "/rest/api/2/fieldconfiguration/1000",
			context:              context.Background(),
			wantHTTPCodeReturn:   http.StatusNoContent,
			wantErr:              false,
		},

		{
			name:                 "DeleteFieldConfigurationWhenTheFieldConfigurationIDIsNotSet",
			fieldConfigurationID: 0,
			wantHTTPMethod:       http.MethodDelete,
			endpoint:             "/rest/api/2/fieldconfiguration/1000",
			context:              context.Background(),
			wantHTTPCodeReturn:   http.StatusNoContent,
			wantErr:              true,
			expectedError:        "jira: no field configuration id set",
		},

		{
			name:                 "DeleteFieldConfigurationWhenTheRequestMethodIsIncorrect",
			fieldConfigurationID: 1000,
			wantHTTPMethod:       http.MethodHead,
			context:              context.Background(),
			wantHTTPCodeReturn:   http.StatusNoContent,
			wantErr:              true,
			expectedError:        "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name:                 "DeleteFieldConfigurationWhenTheContextIsNotProvided",
			fieldConfigurationID: 1000,
			wantHTTPMethod:       http.MethodDelete,
			endpoint:             "/rest/api/2/fieldconfiguration/1000",
			context:              nil,
			wantHTTPCodeReturn:   http.StatusNoContent,
			wantErr:              true,
			expectedError:        "request creation failed: net/http: nil Context",
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
			gotResponse, err := service.Delete(testCase.context, testCase.fieldConfigurationID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

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
