package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_Field_Configuration_Scheme_Service_Assign(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.FieldConfigurationSchemeAssignPayload
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name: "when the parameters are correct",
			payload: &models.FieldConfigurationSchemeAssignPayload{
				FieldConfigurationSchemeID: "10000",
				ProjectID:                  "10000",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "when the payload is not provided",
			payload:            nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "failed to parse the interface pointer, please provide a valid one",
		},

		{
			name: "when the request method is incorrect",
			payload: &models.FieldConfigurationSchemeAssignPayload{
				FieldConfigurationSchemeID: "10000",
				ProjectID:                  "10000",
			},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name: "when the context is not provided",
			payload: &models.FieldConfigurationSchemeAssignPayload{
				FieldConfigurationSchemeID: "10000",
				ProjectID:                  "10000",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
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

			service := &FieldConfigurationSchemeService{client: mockClient}
			gotResponse, err := service.Assign(testCase.context, testCase.payload)

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

func Test_Field_Configuration_Scheme_Service_Create(t *testing.T) {

	testCases := []struct {
		name                                string
		fieldConfigurationSchemeDescription string
		fieldConfigurationSchemeName        string
		mockFile                            string
		wantHTTPMethod                      string
		endpoint                            string
		context                             context.Context
		wantHTTPCodeReturn                  int
		wantErr                             bool
		expectedError                       string
	}{
		{
			name:                                "when the parameters are correct",
			fieldConfigurationSchemeDescription: "field configuration name",
			fieldConfigurationSchemeName:        "field configuration description",
			mockFile:                            "../mocks/create-field-configuration-scheme.json",
			wantHTTPMethod:                      http.MethodPost,
			endpoint:                            "/rest/api/3/fieldconfigurationscheme",
			context:                             context.Background(),
			wantHTTPCodeReturn:                  http.StatusOK,
			wantErr:                             false,
		},

		{
			name:                                "when the field configuration scheme name is not provided",
			fieldConfigurationSchemeDescription: "field configuration description",
			fieldConfigurationSchemeName:        "",
			mockFile:                            "../mocks/create-field-configuration-scheme.json",
			wantHTTPMethod:                      http.MethodPost,
			endpoint:                            "/rest/api/3/fieldconfigurationscheme",
			context:                             context.Background(),
			wantHTTPCodeReturn:                  http.StatusOK,
			wantErr:                             true,
			expectedError:                       "jira: no field configuration scheme name set",
		},

		{
			name:                                "when the response status code is invalid",
			fieldConfigurationSchemeDescription: "field configuration name",
			fieldConfigurationSchemeName:        "field configuration description",
			mockFile:                            "../mocks/create-field-configuration-scheme.json",
			wantHTTPMethod:                      http.MethodPost,
			endpoint:                            "/rest/api/3/fieldconfigurationscheme",
			context:                             context.Background(),
			wantHTTPCodeReturn:                  http.StatusBadRequest,
			wantErr:                             true,
			expectedError:                       "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:                                "when the context is not provided",
			fieldConfigurationSchemeDescription: "field configuration name",
			fieldConfigurationSchemeName:        "field configuration description",
			mockFile:                            "../mocks/create-field-configuration-scheme.json",
			wantHTTPMethod:                      http.MethodPost,
			endpoint:                            "/rest/api/3/fieldconfigurationscheme",
			context:                             nil,
			wantHTTPCodeReturn:                  http.StatusOK,
			wantErr:                             true,
			expectedError:                       "request creation failed: net/http: nil Context",
		},

		{
			name:                                "when the response body is empty",
			fieldConfigurationSchemeDescription: "field configuration name",
			fieldConfigurationSchemeName:        "field configuration description",
			mockFile:                            "../mocks/empty-json.json",
			wantHTTPMethod:                      http.MethodPost,
			endpoint:                            "/rest/api/3/fieldconfigurationscheme",
			context:                             context.Background(),
			wantHTTPCodeReturn:                  http.StatusOK,
			wantErr:                             true,
			expectedError:                       "unexpected end of JSON input",
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

			service := &FieldConfigurationSchemeService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.fieldConfigurationSchemeName,
				testCase.fieldConfigurationSchemeDescription)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func Test_Field_Configuration_Scheme_Service_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		schemeID           int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the parameters are correct",
			schemeID:           1000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "when the field configuration scheme id is not provided",
			schemeID:           0,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "jira: no field configuration scheme id set",
		},

		{
			name:               "when the request method is incorrect",
			schemeID:           1000,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name:               "when the context is not provided",
			schemeID:           1000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
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

			service := &FieldConfigurationSchemeService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.schemeID)

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

func Test_Field_Configuration_Scheme_Service_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		ids                 []int
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
		expectedError       string
	}{
		{
			name:               "when the parameters are correct",
			ids:                []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=1000&id=1001&id=1002&maxResults=100&startAt=50",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
			expectedError:      "",
		},

		{
			name:               "when the response status code is invalid",
			ids:                []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=1000&id=1001&id=1002&maxResults=100&startAt=50",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the context is not provided",
			ids:                []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=1000&id=1001&id=1002&maxResults=100&startAt=50",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response body is empty",
			ids:                []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme?id=1000&id=1001&id=1002&maxResults=100&startAt=50",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
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

			service := &FieldConfigurationSchemeService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.ids, testCase.startAt, testCase.maxResults)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func Test_Field_Configuration_Scheme_Service_Link(t *testing.T) {

	testCases := []struct {
		name               string
		schemeID           int
		payload            *models.FieldConfigurationToIssueTypeMappingPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:     "when the parameters are correct",
			schemeID: 1000,
			payload: &models.FieldConfigurationToIssueTypeMappingPayloadScheme{
				Mappings: []*models.FieldConfigurationToIssueTypeMappingScheme{
					{
						IssueTypeID:          "default",
						FieldConfigurationID: "10000",
					},
					{
						IssueTypeID:          "10001",
						FieldConfigurationID: "10002",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "when the paylod is not provided",
			schemeID:           1000,
			payload:            nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "failed to parse the interface pointer, please provide a valid one",
		},

		{
			name:     "when the field configuration scheme id is not provided",
			schemeID: 0,
			payload: &models.FieldConfigurationToIssueTypeMappingPayloadScheme{
				Mappings: []*models.FieldConfigurationToIssueTypeMappingScheme{
					{
						IssueTypeID:          "default",
						FieldConfigurationID: "10000",
					},
					{
						IssueTypeID:          "10001",
						FieldConfigurationID: "10002",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "jira: no field configuration scheme id set",
		},

		{
			name:     "when the request method is incorrect",
			schemeID: 1000,
			payload: &models.FieldConfigurationToIssueTypeMappingPayloadScheme{
				Mappings: []*models.FieldConfigurationToIssueTypeMappingScheme{
					{
						IssueTypeID:          "default",
						FieldConfigurationID: "10000",
					},
					{
						IssueTypeID:          "10001",
						FieldConfigurationID: "10002",
					},
				},
			},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name:     "when the context is not provided",
			schemeID: 1000,
			payload: &models.FieldConfigurationToIssueTypeMappingPayloadScheme{
				Mappings: []*models.FieldConfigurationToIssueTypeMappingScheme{
					{
						IssueTypeID:          "default",
						FieldConfigurationID: "10000",
					},
					{
						IssueTypeID:          "10001",
						FieldConfigurationID: "10002",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000/mapping",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
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

			service := &FieldConfigurationSchemeService{client: mockClient}
			gotResponse, err := service.Link(testCase.context, testCase.schemeID, testCase.payload)

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

func Test_Field_Configuration_Scheme_Service_Mapping(t *testing.T) {

	testCases := []struct {
		name                string
		fieldConfigIDs      []int
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
		expectedError       string
	}{
		{
			name:               "when the parameters are correct",
			fieldConfigIDs:     []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/get-field-configuration-mapping.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=1000&fieldConfigurationSchemeId=1001&fieldConfigurationSchemeId=1002&maxResults=100&startAt=50",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
			expectedError:      "",
		},

		{
			name:               "when the response status code is invalid",
			fieldConfigIDs:     []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=1000&fieldConfigurationSchemeId=1001&fieldConfigurationSchemeId=1002&maxResults=100&startAt=50",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the context is not provided",
			fieldConfigIDs:     []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=1000&fieldConfigurationSchemeId=1001&fieldConfigurationSchemeId=1002&maxResults=100&startAt=50",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response body is empty",
			fieldConfigIDs:     []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/mapping?fieldConfigurationSchemeId=1000&fieldConfigurationSchemeId=1001&fieldConfigurationSchemeId=1002&maxResults=100&startAt=50",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
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

			service := &FieldConfigurationSchemeService{client: mockClient}
			gotResult, gotResponse, err := service.Mapping(testCase.context, testCase.fieldConfigIDs, testCase.startAt, testCase.maxResults)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func Test_Field_Configuration_Scheme_Service_Project(t *testing.T) {

	testCases := []struct {
		name                string
		projectIDs          []int
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
		expectedError       string
	}{
		{
			name:               "when the parameters are correct",
			projectIDs:         []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/get-field-configuration-schemes-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=100&projectId=1000&projectId=1001&projectId=1002&startAt=50",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
			expectedError:      "",
		},

		{
			name:               "when the response status code is invalid",
			projectIDs:         []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=100&projectId=1000&projectId=1001&projectId=1002&startAt=50",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the context is not provided",
			projectIDs:         []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/get-field-configuration-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=100&projectId=1000&projectId=1001&projectId=1002&startAt=50",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response body is empty",
			projectIDs:         []int{1000, 1001, 1002},
			startAt:            50,
			maxResults:         100,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/project?maxResults=100&projectId=1000&projectId=1001&projectId=1002&startAt=50",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
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

			service := &FieldConfigurationSchemeService{client: mockClient}
			gotResult, gotResponse, err := service.Project(testCase.context, testCase.projectIDs, testCase.startAt, testCase.maxResults)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func Test_Field_Configuration_Scheme_Service_Unlink(t *testing.T) {

	testCases := []struct {
		name               string
		schemeID           int
		issueTypeIDs       []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the parameters are correct",
			schemeID:           1000,
			issueTypeIDs:       []string{"10001", "2"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000/mapping/delete",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "when the field configuration scheme id is not provided",
			schemeID:           0,
			issueTypeIDs:       []string{"10001", "2"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000/mapping/delete",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "jira: no field configuration scheme id set",
		},

		{
			name:               "when the request method is incorrect",
			schemeID:           1000,
			issueTypeIDs:       []string{"10001", "2"},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000/mapping/delete",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name:               "when the context is not provided",
			schemeID:           1000,
			issueTypeIDs:       []string{"10001", "2"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/fieldconfigurationscheme/1000/mapping/delete",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
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

			service := &FieldConfigurationSchemeService{client: mockClient}
			gotResponse, err := service.Unlink(testCase.context, testCase.schemeID, testCase.issueTypeIDs)

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

func Test_Field_Configuration_Scheme_Service_Update(t *testing.T) {

	testCases := []struct {
		name                                string
		schemeID                            int
		fieldConfigurationSchemeDescription string
		fieldConfigurationSchemeName        string
		mockFile                            string
		wantHTTPMethod                      string
		endpoint                            string
		context                             context.Context
		wantHTTPCodeReturn                  int
		wantErr                             bool
		expectedError                       string
	}{
		{
			name:                                "when the parameters are correct",
			schemeID:                            100,
			fieldConfigurationSchemeDescription: "field configuration name",
			fieldConfigurationSchemeName:        "field configuration description",
			wantHTTPMethod:                      http.MethodPut,
			endpoint:                            "/rest/api/3/fieldconfigurationscheme/100",
			context:                             context.Background(),
			wantHTTPCodeReturn:                  http.StatusOK,
			wantErr:                             false,
		},

		{
			name:                                "when the field configuration scheme id is not provided",
			schemeID:                            0,
			fieldConfigurationSchemeDescription: "field configuration name",
			fieldConfigurationSchemeName:        "field configuration description",
			wantHTTPMethod:                      http.MethodPut,
			endpoint:                            "/rest/api/3/fieldconfigurationscheme/100",
			context:                             context.Background(),
			wantHTTPCodeReturn:                  http.StatusOK,
			wantErr:                             true,
			expectedError:                       "jira: no field configuration scheme id set",
		},

		{
			name:                                "when the field configuration scheme name is not provided",
			schemeID:                            100,
			fieldConfigurationSchemeDescription: "field configuration description",
			fieldConfigurationSchemeName:        "",
			wantHTTPMethod:                      http.MethodPut,
			endpoint:                            "/rest/api/3/fieldconfigurationscheme/100",
			context:                             context.Background(),
			wantHTTPCodeReturn:                  http.StatusOK,
			wantErr:                             true,
			expectedError:                       "jira: no field configuration scheme name set",
		},

		{
			name:                                "when the response status code is invalid",
			schemeID:                            100,
			fieldConfigurationSchemeDescription: "field configuration name",
			fieldConfigurationSchemeName:        "field configuration description",
			wantHTTPMethod:                      http.MethodPut,
			endpoint:                            "/rest/api/3/fieldconfigurationscheme/100",
			context:                             context.Background(),
			wantHTTPCodeReturn:                  http.StatusBadRequest,
			wantErr:                             true,
			expectedError:                       "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:                                "when the context is not provided",
			schemeID:                            100,
			fieldConfigurationSchemeDescription: "field configuration name",
			fieldConfigurationSchemeName:        "field configuration description",
			wantHTTPMethod:                      http.MethodPut,
			endpoint:                            "/rest/api/3/fieldconfigurationscheme/100",
			context:                             nil,
			wantHTTPCodeReturn:                  http.StatusOK,
			wantErr:                             true,
			expectedError:                       "request creation failed: net/http: nil Context",
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

			service := &FieldConfigurationSchemeService{client: mockClient}
			gotResponse, err := service.Update(testCase.context, testCase.schemeID, testCase.fieldConfigurationSchemeName,
				testCase.fieldConfigurationSchemeDescription)

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
