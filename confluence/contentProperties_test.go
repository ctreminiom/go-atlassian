package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestContentPropertyService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		payload            *ContentPropertyPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:      "CreateContentPropertyWhenTheParametersAreCorrect",
			contentID: "9438383838383",
			payload: &ContentPropertyPayloadScheme{
				Key:   "key",
				Value: "value",
			},
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/9438383838383/property",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:      "CreateContentPropertyWhenTheContentIDIsNotSet",
			contentID: "",
			payload: &ContentPropertyPayloadScheme{
				Key:   "key",
				Value: "value",
			},
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/9438383838383/property",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateContentPropertyWhenThePayloadIsNotSet",
			contentID:          "9438383838383",
			payload:            nil,
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/9438383838383/property",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "CreateContentPropertyWhenTheContextIsNotSet",
			contentID: "9438383838383",
			payload: &ContentPropertyPayloadScheme{
				Key:   "key",
				Value: "value",
			},
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/9438383838383/property",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "CreateContentPropertyWhenTheParametersAreCorrect",
			contentID: "9438383838383",
			payload: &ContentPropertyPayloadScheme{
				Key:   "key",
				Value: "value",
			},
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/9438383838383/property",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:      "CreateContentPropertyWhenTheRequestMethodIsIncorrect",
			contentID: "9438383838383",
			payload: &ContentPropertyPayloadScheme{
				Key:   "key",
				Value: "value",
			},
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/wiki/rest/api/content/9438383838383/property",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "CreateContentPropertyWhenTheStatusCodeIsIncorrect",
			contentID: "9438383838383",
			payload: &ContentPropertyPayloadScheme{
				Key:   "key",
				Value: "value",
			},
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/9438383838383/property",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &ContentPropertyService{client: mockClient}

			gotResult, gotResponse, err := service.Create(testCase.context, testCase.contentID, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
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

func TestContentPropertyService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		contentID, key     string
		expand             []string
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
			name:               "GetContentPropertiesWhenTheParametersAreCorrect",
			contentID:          "9438383838383",
			expand:             []string{"expand-00", "expand-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-properties.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/9438383838383/property?expand=expand-00%2Cexpand-01&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentPropertiesWhenTheContentIDIsNotSet",
			contentID:          "",
			expand:             []string{"expand-00", "expand-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-properties.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/9438383838383/property?expand=expand-00%2Cexpand-01&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentPropertiesWhenTheContextIsNotSet",
			contentID:          "9438383838383",
			expand:             []string{"expand-00", "expand-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-properties.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/9438383838383/property?expand=expand-00%2Cexpand-01&limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentPropertiesWhenTheRequestMethodIsIncorrect",
			contentID:          "9438383838383",
			expand:             []string{"expand-00", "expand-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-properties.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/wiki/rest/api/content/9438383838383/property?expand=expand-00%2Cexpand-01&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentPropertiesWhenTheParametersAreCorrect",
			contentID:          "9438383838383",
			expand:             []string{"expand-00", "expand-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-properties.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/9438383838383/property?expand=expand-00%2Cexpand-01&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &ContentPropertyService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.contentID, testCase.expand, testCase.startAt, testCase.maxResults)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
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

func TestContentPropertyService_Delete(t *testing.T) {
	testCases := []struct {
		name               string
		contentID, key     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "RemoveContentPropertyWhenTheParametersAreCorrect",
			contentID:          "9438383838383",
			key:                "editor",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "RemoveContentPropertyWhenTheContentIDIsNotSet",
			contentID:          "",
			key:                "editor",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "RemoveContentPropertyWhenTheKeyIsNotSet",
			contentID:          "9438383838383",
			key:                "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveContentPropertyWhenTheContextIsNotSet",
			contentID:          "9438383838383",
			key:                "editor",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveContentPropertyWhenTheRequestMethodIsIncorrect",
			contentID:          "9438383838383",
			key:                "editor",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &ContentPropertyService{client: mockClient}

			gotResponse, err := service.Delete(testCase.context, testCase.contentID, testCase.key)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
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

func TestContentPropertyService_Get(t *testing.T) {
	testCases := []struct {
		name               string
		contentID, key     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetContentPropertyWhenTheParametersAreCorrect",
			contentID:          "9438383838383",
			key:                "editor",
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentPropertyWhenTheContentIDIsNotSet",
			contentID:          "",
			key:                "editor",
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentPropertyWhenTheKeyIsNotSet",
			contentID:          "9438383838383",
			key:                "",
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentPropertyWhenTheContextIsNotSet",
			contentID:          "9438383838383",
			key:                "editor",
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentPropertyWhenTheRequestMethodIsIncorrect",
			contentID:          "9438383838383",
			key:                "editor",
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentPropertyWhenTheStatusCodeIsIncorrect",
			contentID:          "9438383838383",
			key:                "editor",
			mockFile:           "./mocks/get-content-property.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/9438383838383/property/editor",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &ContentPropertyService{client: mockClient}

			gotResult, gotResponse, err := service.Get(testCase.context, testCase.contentID, testCase.key)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
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
