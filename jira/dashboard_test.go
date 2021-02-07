package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestDashboardService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		startAt            int
		maxResults         int
		filter             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetDashboardsWhenIsCorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "./mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/dashboard?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetDashboardsWhenTheFilterQueryIsIncorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "xxxx",
			mockFile:           "./mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/dashboard?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheHTTPResponseCodeIsIncorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "./mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/dashboard?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheRequestMethodIsIncorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "./mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/dashboard?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheQueryParametersAreIncorrect",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "./mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/dashboard?maxResults=50&startAt=11111",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheContextIsNil",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "./mocks/get_dashboards.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/dashboard?maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheResponseBodyLengthIsZero",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/dashboard?maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetDashboardsWhenTheResponseBodyHasADifferentFormat",
			startAt:            0,
			maxResults:         50,
			filter:             "",
			mockFile:           "./mocks/invalid-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/dashboard?maxResults=50&startAt=0",
			context:            nil,
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

			service := &DashboardService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.startAt, testCase.maxResults, testCase.filter)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode()))
				assert.Equal(t, testCase.endpoint, fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode()))

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)
			}

		})
	}
}
