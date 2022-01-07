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

func Test_Field_Configuration_Service_Update(t *testing.T) {

	testCases := []struct {
		name                 string
		fieldConfigurationID int
		payload              *models.UpdateFieldConfigurationItemPayloadScheme
		mockFile             string
		wantHTTPMethod       string
		endpoint             string
		context              context.Context
		wantHTTPCodeReturn   int
		wantErr              bool
		expectedError        string
	}{
		{
			name:                 "when the parameters are correct",
			fieldConfigurationID: 1000,
			payload: &models.UpdateFieldConfigurationItemPayloadScheme{
				FieldConfigurationItems: []*models.FieldConfigurationItemScheme{
					{
						ID:          "customfield_10012",
						IsHidden:    false,
						Description: "The new description of this item.",
					},
					{
						ID:         "customfield_10011",
						IsRequired: true,
					},
					{
						ID:          "customfield_10010",
						IsHidden:    false,
						IsRequired:  false,
						Description: "Another new description.",
						Renderer:    "wiki-renderer",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/fieldconfiguration/1000/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:                 "when the field configuration id is not provided",
			fieldConfigurationID: 0,
			payload: &models.UpdateFieldConfigurationItemPayloadScheme{
				FieldConfigurationItems: []*models.FieldConfigurationItemScheme{
					{
						ID:          "customfield_10012",
						IsHidden:    false,
						Description: "The new description of this item.",
					},
					{
						ID:         "customfield_10011",
						IsRequired: true,
					},
					{
						ID:          "customfield_10010",
						IsHidden:    false,
						IsRequired:  false,
						Description: "Another new description.",
						Renderer:    "wiki-renderer",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/fieldconfiguration/1000/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "jira: no field configuration id set",
		},

		{
			name:                 "when the payload is not provided",
			fieldConfigurationID: 1000,
			payload:              nil,
			wantHTTPMethod:       http.MethodPut,
			endpoint:             "/rest/api/2/fieldconfiguration/1000/fields",
			context:              context.Background(),
			wantHTTPCodeReturn:   http.StatusNoContent,
			wantErr:              true,
			expectedError:        "failed to parse the interface pointer, please provide a valid one",
		},

		{
			name:                 "when the request method is incorrect",
			fieldConfigurationID: 1000,
			payload: &models.UpdateFieldConfigurationItemPayloadScheme{
				FieldConfigurationItems: []*models.FieldConfigurationItemScheme{
					{
						ID:          "customfield_10012",
						IsHidden:    false,
						Description: "The new description of this item.",
					},
					{
						ID:         "customfield_10011",
						IsRequired: true,
					},
					{
						ID:          "customfield_10010",
						IsHidden:    false,
						IsRequired:  false,
						Description: "Another new description.",
						Renderer:    "wiki-renderer",
					},
				},
			},
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/2/fieldconfiguration/1000/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name:                 "when the context is not provided",
			fieldConfigurationID: 1000,
			payload: &models.UpdateFieldConfigurationItemPayloadScheme{
				FieldConfigurationItems: []*models.FieldConfigurationItemScheme{
					{
						ID:          "customfield_10012",
						IsHidden:    false,
						Description: "The new description of this item.",
					},
					{
						ID:         "customfield_10011",
						IsRequired: true,
					},
					{
						ID:          "customfield_10010",
						IsHidden:    false,
						IsRequired:  false,
						Description: "Another new description.",
						Renderer:    "wiki-renderer",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/fieldconfiguration/1000/fields",
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

			service := &FieldConfigurationItemService{client: mockClient}
			gotResponse, err := service.Update(testCase.context, testCase.fieldConfigurationID, testCase.payload)

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

func Test_Field_Configuration_Service_Get(t *testing.T) {

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
		expectedError      string
	}{
		{
			name:               "when the parameters are correct",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/get-field-configuration-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "when the field configuration id is not provided",
			fieldConfigID:      0,
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/get-field-configuration-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "jira: no field configuration id set",
		},

		{
			name:               "when the request method is incorrect",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/get-field-configuration-items.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name:               "when the context is not provided",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/get-field-configuration-items.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response body is empty",
			fieldConfigID:      10000,
			startAt:            0,
			maxResult:          50,
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/fieldconfiguration/10000/fields?maxResults=50&startAt=0",
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

			service := &FieldConfigurationItemService{client: mockClient}
			getResult, gotResponse, err := service.Gets(testCase.context, testCase.fieldConfigID, testCase.startAt, testCase.maxResult)

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
