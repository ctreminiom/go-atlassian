package v3

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestResolutionService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		resolutionID       string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
		errorMessage       string
	}{
		{
			name:               "GetIssueResolutionWhenTheIDIsCorrect",
			mockFile:           "./mocks/get_resolution_10000.json",
			resolutionID:       "10000",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/resolution/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueResolutionWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty_json.json",
			resolutionID:       "10000",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/resolution/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueResolutionWhenTheIDIsNotSet",
			mockFile:           "./mocks/get_resolution_10000.json",
			resolutionID:       "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/resolution/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueResolutionWhenTheIDIsIncorrect",
			mockFile:           "./mocks/get_resolution_10000.json",
			resolutionID:       "10001",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/resolution/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssueResolutionWhenTheIDIsEmpty",
			mockFile:           "./mocks/get_resolution_10000.json",
			resolutionID:       "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/resolution/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssueResolutionWhenTheIDIsANumber",
			mockFile:           "./mocks/get_resolution_10000.json",
			resolutionID:       "2222",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/resolution/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssueResolutionWhenTheContextIsNil",
			mockFile:           "./mocks/get_resolution_10000.json",
			resolutionID:       "10000",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/resolution/10000",
			context:            nil,
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssueResolutionWhenTheHTTPMethodIsIncorrect",
			mockFile:           "./mocks/get_resolution_10000.json",
			resolutionID:       "10000",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/resolution/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
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
				Headers:            testCase.wantHTTPHeaders,
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

			service := &ResolutionService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.resolutionID)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func TestResolutionService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
		errorMessage       string
	}{
		{
			name:               "GetsIssueResolutionWhenTheIDIsCorrect",
			mockFile:           "./mocks/get_resolutions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/resolution",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsIssueResolutionWhenTheResponseBodyIsEmpty",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/resolution",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueResolutionWhenTheMethodIsDifferent",
			mockFile:           "./mocks/get_resolutions.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/resolution",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsIssueResolutionWhenTheContextIsNil",
			mockFile:           "./mocks/get_resolutions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/resolution",
			context:            nil,
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
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
				Headers:            testCase.wantHTTPHeaders,
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

			service := &ResolutionService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}
