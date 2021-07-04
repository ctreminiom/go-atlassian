package sm

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestInfoService_Get(t *testing.T) {

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
			name:               "GetServiceManagementInfoWhenTheParamsAreCorrect",
			mockFile:           "./mocks/get-service-management-info.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/info",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetServiceManagementInfoWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-service-management-info.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/info",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetServiceManagementInfoWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-service-management-info.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/info",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetServiceManagementInfoWhenTheContextIsNil",
			mockFile:           "./mocks/get-service-management-info.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/info",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetServiceManagementInfoWhenTheEndpointIsEmpty",
			mockFile:           "./mocks/get-service-management-info.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetServiceManagementInfoWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/info",
			context:            context.Background(),
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

			service := &InfoService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context)

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
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}
