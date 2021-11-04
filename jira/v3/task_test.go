package v3

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestTaskService_Cancel(t1 *testing.T) {

	testCases := []struct {
		name               string
		taskID             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CancelTaskWhenTheParamsAreCorrect",
			taskID:             "1",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/task/1/cancel",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            false,
		},

		{
			name:               "CancelTaskWhenTheTaskIDIsEmpty",
			taskID:             "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/task/1/cancel",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},

		{
			name:               "CancelTaskWhenTheRequestMethodIsIncorrect",
			taskID:             "1",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/task/1/cancel",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},

		{
			name:               "CancelTaskWhenTheStatusCodeIsIncorrect",
			taskID:             "1",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/task/1/cancel",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "CancelTaskWhenTheContextIsNil",
			taskID:             "1",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/task/1/cancel",
			context:            nil,
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t1.Run(testCase.name, func(t *testing.T) {

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

			i := &TaskService{client: mockClient}

			gotResponse, err := i.Cancel(testCase.context, testCase.taskID)

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

func TestTaskService_Get(t1 *testing.T) {

	testCases := []struct {
		name               string
		taskID             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetTaskWhenTheParamsAreCorrect",
			taskID:             "1",
			mockFile:           "./mocks/task.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/task/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetTaskWhenTheTaskIDIsEmpty",
			taskID:             "",
			mockFile:           "./mocks/task.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/task/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetTaskWhenTheRequestMethodIsIncorrect",
			taskID:             "1",
			mockFile:           "./mocks/task.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/task/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetTaskWhenTheStatusCodeIsIncorrect",
			taskID:             "1",
			mockFile:           "./mocks/task.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/task/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetTaskWhenTheContextIsNil",
			taskID:             "1",
			mockFile:           "./mocks/task.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/task/1",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetTaskWhenTheEndpointIsIncorrect",
			taskID:             "1",
			mockFile:           "./mocks/task.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/task/2",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetTaskWhenTheResponseBodyHasADifferentFormat",
			taskID:             "1",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/task/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t1.Run(testCase.name, func(t *testing.T) {

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

			i := &TaskService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.taskID)

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
