package v3

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestLabelService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
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
			name:               "GetInstanceLabelsWhenTheParametersAreCorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-instance-labels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/label?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetInstanceLabelsWhenTheParametersAreIncorrect",
			startAt:            0,
			maxResults:         60,
			mockFile:           "./mocks/get-instance-labels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/label?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetInstanceLabelsWhenTheEndpointIsIncorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-instance-labels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/labels?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetInstanceLabelsWhenTheContextIsNil",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-instance-labels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/label?maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetInstanceLabelsWhenTheResponseBodyHasADifferentFormat",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/label?maxResults=50&startAt=0",
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

			service := &LabelService{client: mockClient}
			getResult, gotResponse, err := service.Gets(testCase.context, testCase.startAt, testCase.maxResults)

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
