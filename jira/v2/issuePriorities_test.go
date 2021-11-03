package v2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPriorityService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssuePrioritiesWhenIsCorrect",
			mockFile:           "../v3/mocks/get_priorities.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            false,
		},
		{
			name:               "GetIssuePrioritiesWhenTheContextIsNil",
			mockFile:           "../v3/mocks/get_priorities.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority",
			context:            nil,
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},

		{
			name:               "GetIssuePrioritiesWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},

		{
			name:               "GetIssuePrioritiesWhenTheRequestMethodIsIncorrect",
			mockFile:           "../v3/mocks/get_priorities.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/priority",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},

		{
			name:               "GetIssuePrioritiesWhenTheStatusCodeIsIncorrect",
			mockFile:           "../v3/mocks/get_priorities.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
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

			i := &PriorityService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context)

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

func TestPriorityService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		priorityID         string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssuePriorityByIDWhenTheIDIsCorrect",
			mockFile:           "../v3/mocks/get_priorities_1.json",
			priorityID:         "1",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority/1",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            false,
		},

		{
			name:               "GetIssuePriorityByIDWhenTheContextIsNil",
			mockFile:           "../v3/mocks/get_priorities_1.json",
			priorityID:         "1",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority/1",
			context:            nil,
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},

		{
			name:               "GetIssuePriorityByIDWhenTheIDIsNotSet",
			mockFile:           "../v3/mocks/get_priorities_1.json",
			priorityID:         "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority/1",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},

		{
			name:               "GetIssuePriorityByIDWhenTheIDIsIncorrect",
			mockFile:           "../v3/mocks/get_priorities_1.json",
			priorityID:         "12",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority/1",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssuePriorityByIDWhenTheIDIsEmpty",
			mockFile:           "../v3/mocks/get_priorities_1.json",
			priorityID:         "",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority/1",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssuePriorityByIDWhenEndpointURLIsIncorrect",
			mockFile:           "../v3/mocks/get_priorities_1.json",
			priorityID:         "1",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority/1asd",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},

		{
			name:               "GetIssuePriorityByIDWhenTheRequestMethodIsIncorrect",
			mockFile:           "../v3/mocks/get_priorities_1.json",
			priorityID:         "1",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/priority/1",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},

		{
			name:               "GetIssuePriorityByIDWhenTheStatusCodeIsIncorrect",
			mockFile:           "../v3/mocks/get_priorities_1.json",
			priorityID:         "1",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority/1",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},

		{
			name:               "GetIssuePriorityByIDWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "../v3/mocks/empty_json.json",
			priorityID:         "1",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/priority/1",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
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

			i := &PriorityService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.priorityID)

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
