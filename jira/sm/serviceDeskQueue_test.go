package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestServiceDeskQueueService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		serviceDeskID      int
		includeCount       bool
		start, limit       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetServiceDeskQueuesWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			includeCount:       true,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-service-desk-queues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue?includeCount=true&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetServiceDeskQueuesWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			mockFile:           "./mocks/get-service-desk-queues.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue?limit=0&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskQueuesWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			mockFile:           "./mocks/get-service-desk-queues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue?limit=0&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskQueuesWhenTheContextIsNil",
			serviceDeskID:      1,
			mockFile:           "./mocks/get-service-desk-queues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue?limit=0&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskQueuesWhenTheEndpointIsIncorrect",
			serviceDeskID:      1,
			mockFile:           "./mocks/get-service-desk-queues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "AS*(saisaidhiahihihihi494949xzx.x[sa[.[;''''",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskQueuesWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue?limit=0&start=0",
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

			service := &ServiceDeskQueueService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.serviceDeskID, testCase.includeCount, testCase.start, testCase.limit)

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

				for pos, queue := range gotResult.Values {

					t.Log("------------------------------------")
					t.Logf("Queue ID #%v: %v", pos+1, queue.ID)
					t.Logf("Queue Name #%v: %v", pos+1, queue.Name)
					t.Logf("Queue JQL #%v: %v", pos+1, queue.Jql)
					t.Logf("Queue Issue Count #%v: %v", pos+1, queue.IssueCount)
					t.Logf("Queue Fields #%v: %v", pos+1, queue.Fields)
					t.Log("------------------------------------")
				}
			}

		})
	}

}

func TestServiceDeskQueueService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		serviceDeskID      int
		queueID            int
		includeCount       bool
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetServiceDeskQueueWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			queueID:            1,
			includeCount:       true,
			mockFile:           "./mocks/get-service-desk-queue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue/1?includeCount=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetServiceDeskQueueWhenTheIncludeIssueIsNotSet",
			serviceDeskID:      1,
			queueID:            1,
			includeCount:       false,
			mockFile:           "./mocks/get-service-desk-queue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetServiceDeskQueueWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			queueID:            1,
			includeCount:       true,
			mockFile:           "./mocks/get-service-desk-queue.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue/1?includeCount=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskQueuesWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			queueID:            1,
			includeCount:       true,
			mockFile:           "./mocks/get-service-desk-queue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue/1?includeCount=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskQueuesWhenTheContextIsNil",
			serviceDeskID:      1,
			queueID:            1,
			includeCount:       true,
			mockFile:           "./mocks/get-service-desk-queue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue/1?includeCount=true",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskQueuesWhenTheEndpointIsIncorrect",
			serviceDeskID:      1,
			queueID:            1,
			includeCount:       true,
			mockFile:           "./mocks/get-service-desk-queue.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "AS*(saisaidhiahihihihi494949xzx.x[sa[.[;''''",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskQueuesWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			queueID:            1,
			includeCount:       true,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue/1?includeCount=true",
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

			service := &ServiceDeskQueueService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.serviceDeskID, testCase.queueID, testCase.includeCount)

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

				t.Log("------------------------------------")
				t.Logf("Queue ID %v", gotResult.ID)
				t.Logf("Queue Name %v", gotResult.Name)
				t.Logf("Queue JQL %v", gotResult.Jql)
				t.Logf("Queue Issue Count %v", gotResult.IssueCount)
				t.Logf("Queue Fields %v", gotResult.Fields)
				t.Log("------------------------------------")
			}

		})
	}

}

func TestServiceDeskQueueService_Issues(t *testing.T) {

	testCases := []struct {
		name               string
		serviceDeskID      int
		queueID            int
		start, limit       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetServiceDeskQueueIssuesWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			queueID:            1,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-service-desk-queue-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue/1/issue?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetServiceDeskQueueIssuesWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			queueID:            1,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-service-desk-queue-issues.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue/1/issue?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:           "GetServiceDeskQueueIssuesWhenTheStatusCodeIsIncorrect",
			serviceDeskID:  1,
			queueID:        1,
			start:          0,
			limit:          50,
			mockFile:       "./mocks/get-service-desk-queue-issues.json",
			wantHTTPMethod: http.MethodGet,
			endpoint:       "/rest/servicedeskapi/servicedesk/1/queue/1/issue?limit=50&start=0",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:           "GetServiceDeskQueueIssuesWhenTheContextIsNil",
			serviceDeskID:  1,
			queueID:        1,
			start:          0,
			limit:          50,
			mockFile:       "./mocks/get-service-desk-queue-issues.json",
			wantHTTPMethod: http.MethodGet,
			endpoint:       "/rest/servicedeskapi/servicedesk/1/queue/1/issue?limit=50&start=0",

			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskQueueIssuesWhenTheEndpointIsIncorrect",
			serviceDeskID:      1,
			queueID:            1,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-service-desk-queue-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "AS*(saisaidhiahihihihi494949xzx.x[sa[.[;''''",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetServiceDeskQueueIssuesWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			queueID:            1,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/queue/1/issue?limit=50&start=0",
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

			service := &ServiceDeskQueueService{client: mockClient}
			gotResult, gotResponse, err := service.Issues(testCase.context, testCase.serviceDeskID, testCase.queueID, testCase.start, testCase.limit)

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

				for _, issue := range gotResult.Values {
					t.Log(issue)
				}
			}

		})
	}

}
