package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestRequestTypeService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		query              string
		start, limit       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetRequestApprovalWhenTheParametersAreCorrect",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/requesttype?limit=50&searchQuery=test&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetRequestApprovalWhenTheRequestMethodIsIncorrect",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/requesttype?limit=50&searchQuery=test&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalWhenTheStatusCodeIsIncorrect",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/requesttype?limit=50&searchQuery=test&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetRequestApprovalWhenTheContextIsNil",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/requesttype?limit=50&searchQuery=test&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalWhenTheEndpointIsIncorrect",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalWhenTheResponseBodyHasADifferentFormat",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/requesttype?limit=50&searchQuery=test&start=0",
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

			service := &RequestTypeService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.query, testCase.start, testCase.limit)

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

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				for _, requestType := range gotResult.Values {
					t.Log(requestType.ID, requestType.Name, requestType.Description)
				}
			}

		})
	}

}
