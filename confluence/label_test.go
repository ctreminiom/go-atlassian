package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_Label_Service_Get(t *testing.T) {

	testCases := []struct {
		name               string
		labelName          string
		labelType          string
		start              int
		limit              int
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
			start:              0,
			limit:              50,
			labelName:          "tracking",
			labelType:          "page",
			mockFile:           "./mocks/get-label-details.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/label?limit=50&name=tracking&start=0&type=page",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:               "when the label name is not provided",
			start:              0,
			limit:              50,
			labelName:          "",
			labelType:          "page",
			mockFile:           "./mocks/get-label-details.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/label?limit=50&name=tracking&start=0&type=page",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no label name set",
		},

		{
			name:               "when the context is not provided",
			start:              0,
			limit:              50,
			labelName:          "tracking",
			labelType:          "page",
			mockFile:           "./mocks/get-label-details.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/label?limit=50&name=tracking&start=0&type=page",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is not correct",
			start:              0,
			limit:              50,
			labelName:          "tracking",
			labelType:          "page",
			mockFile:           "./mocks/get-label-details.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/label?limit=50&name=tracking&start=0&type=page",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the response body is not empty",
			start:              0,
			limit:              50,
			labelName:          "tracking",
			labelType:          "page",
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/label?limit=50&name=tracking&start=0&type=page",
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

			service := &LabelService{client: mockClient}

			gotResult, gotResponse, err := service.Get(testCase.context, testCase.labelName, testCase.labelType,
				testCase.start, testCase.limit)

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
