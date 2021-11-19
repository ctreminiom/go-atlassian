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

func TestFieldService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.CustomFieldScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateFieldWhenThePayloadIsCorrect",
			payload: &models.CustomFieldScheme{
				Name:        "Alliance",
				Description: "this is the alliance description field",
				FieldType:   "cascadingselect",
				SearcherKey: "cascadingselectsearcher",
			},
			mockFile:           "./mocks/field-created.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},
		{
			name: "CreateFieldWhenTheEndpointIsIncorrect",
			payload: &models.CustomFieldScheme{
				Name:        "Alliance",
				Description: "this is the alliance description field",
				FieldType:   "cascadingselect",
				SearcherKey: "cascadingselectsearcher",
			},
			mockFile:           "./mocks/field-created.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "CreateFieldWhenThePayloadIsNil",
			payload:            nil,
			mockFile:           "./mocks/field-created.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "CreateFieldWhenTheContextIsNil",
			payload: &models.CustomFieldScheme{
				Name:        "Alliance",
				Description: "this is the alliance description field",
				FieldType:   "cascadingselect",
				SearcherKey: "cascadingselectsearcher",
			},
			mockFile:           "./mocks/field-created.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "CreateFieldWhenTheRequestMethodIsIncorrect",
			payload: &models.CustomFieldScheme{
				Name:        "Alliance",
				Description: "this is the alliance description field",
				FieldType:   "cascadingselect",
				SearcherKey: "cascadingselectsearcher",
			},
			mockFile:           "./mocks/field-created.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "CreateFieldWhenTheStatusCodeIsIncorrect",
			payload: &models.CustomFieldScheme{
				Name:        "Alliance",
				Description: "this is the alliance description field",
				FieldType:   "cascadingselect",
				SearcherKey: "cascadingselectsearcher",
			},
			mockFile:           "./mocks/field-created.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name: "CreateFieldWhenTheResponseBodyHasADifferentFormat",
			payload: &models.CustomFieldScheme{
				Name:        "Alliance",
				Description: "this is the alliance description field",
				FieldType:   "cascadingselect",
				SearcherKey: "cascadingselectsearcher",
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field",
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

			service := &FieldService{client: mockClient}

			getResult, gotResponse, err := service.Create(testCase.context, testCase.payload)

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

func TestFieldService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetFields",
			mockFile:           "./mocks/get-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetFieldsWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-fields.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetFieldsWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetFieldsWhenTheContextIsNil",
			mockFile:           "./mocks/get-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetFieldsWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetFieldsWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/fields",
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

			service := &FieldService{client: mockClient}

			getResult, gotResponse, err := service.Gets(testCase.context)

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

func TestFieldService_Search(t *testing.T) {

	testCases := []struct {
		name               string
		options            *models.FieldSearchOptionsScheme
		startAt            int
		maxResults         int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "SearchFieldsWhenTheOptionsAreCorrect",
			options: &models.FieldSearchOptionsScheme{
				Types:   []string{"custom"},
				IDs:     []string{"111", "12223"},
				Query:   "query-sample",
				OrderBy: "lastUsed",
				Expand:  []string{"screensCount", "lastUsed"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/search?expand=screensCount%2ClastUsed&id=111%2C12223&maxResults=50&orderBy=lastUsed&query=query-sample&startAt=0&type=custom",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "SearchFieldsWhenTheOptionsIsNil",
			options:            nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/search?expand=screensCount%2ClastUsed&maxResults=50&orderBy=lastUsed&startAt=0&type=custom",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "SearchFieldsWhenTheRequestMethodIsIncorrect",
			options: &models.FieldSearchOptionsScheme{
				Types:   []string{"custom"},
				OrderBy: "lastUsed",
				Expand:  []string{"screensCount", "lastUsed"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-fields.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/field/search?expand=screensCount%2ClastUsed&maxResults=50&orderBy=lastUsed&startAt=0&type=custom",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "SearchFieldsWhenTheStatusCodeIsIncorrect",
			options: &models.FieldSearchOptionsScheme{
				Types:   []string{"custom"},
				OrderBy: "lastUsed",
				Expand:  []string{"screensCount", "lastUsed"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/search?expand=screensCount%2ClastUsed&maxResults=50&orderBy=lastUsed&startAt=0&type=custom",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name: "SearchFieldsWhenTheContextIsNil",
			options: &models.FieldSearchOptionsScheme{
				Types:   []string{"custom"},
				OrderBy: "lastUsed",
				Expand:  []string{"screensCount", "lastUsed"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/search?expand=screensCount%2ClastUsed&maxResults=50&orderBy=lastUsed&startAt=0&type=custom",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "SearchFieldsWhenTheResponseBodyHasADifferentFormat",
			options: &models.FieldSearchOptionsScheme{
				Types:   []string{"custom"},
				OrderBy: "lastUsed",
				Expand:  []string{"screensCount", "lastUsed"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-fields.json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/search?expand=screensCount%2ClastUsed&maxResults=50&orderBy=lastUsed&startAt=0&type=custom",
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

			service := &FieldService{client: mockClient}

			getResult, gotResponse, err := service.Search(testCase.context, testCase.options, testCase.startAt, testCase.maxResults)

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
