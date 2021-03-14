package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestRequestService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		opts               *RequestGetOptionsScheme
		start, limit       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "GetCustomerRequestsWhenTheParametersAreCorrect",
			opts: &RequestGetOptionsScheme{
				SearchTerm:        "IT Login",
				RequestOwnerships: []string{"OWNED_REQUESTS"},
				RequestStatus:     "ALL_REQUESTS",
				ApprovalStatus:    "MY_PENDING_APPROVAL",
				OrganizationId:    2,
				ServiceDeskID:     1,
				RequestTypeID:     33,
				Expand:            []string{"serviceDesk", "requestType", "status", "action"},
			},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-customer-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request?approvalStatus=MY_PENDING_APPROVAL&expand=serviceDesk%2CrequestType%2Cstatus%2Caction&limit=50&organizationId=2&requestOwnership=OWNED_REQUESTS&requestStatus=ALL_REQUESTS&requestTypeId=33&searchTerm=IT+Login&serviceDeskId=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerRequestsWhenTheOptionsAreNil",
			opts:               nil,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-customer-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerRequestsWhenTheRequestMethodIsIncorrect",
			opts:               nil,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-customer-requests.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/servicedeskapi/request?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestsWhenTheStatusCodeIsIncorrect",
			opts:               nil,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-customer-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestsWhenTheContextIsNil",
			opts:               nil,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-customer-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request?limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestsWhenTheEndpointIsEmpty",
			opts:               nil,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-customer-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestsWhenTheResponseBodyHasADifferentFormat",
			opts:               nil,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request?limit=50&start=0",
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

			service := &RequestService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.opts, testCase.start, testCase.limit)

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

				for _, request := range gotResult.Values {
					t.Log("-------------------------------------------")
					t.Log(request)
					t.Log("-------------------------------------------")
				}
			}

		})
	}

}

func TestRequestService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		expands            []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetCustomerRequestWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			expands:            []string{"serviceDesk", "requestType", "participant", "sla", "status", "attachment", "action"},
			mockFile:           "./mocks/get-customer-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3?expand=serviceDesk%2CrequestType%2Cparticipant%2Csla%2Cstatus%2Cattachment%2Caction",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerRequestWhenTheExpandsAreNotSet",
			issueKeyOrID:       "DUMMY-3",
			expands:            nil,
			mockFile:           "./mocks/get-customer-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerRequestWhenTheIssueKeyIsNotSet",
			issueKeyOrID:       "",
			expands:            []string{"serviceDesk", "requestType", "participant", "sla", "status", "attachment", "action"},
			mockFile:           "./mocks/get-customer-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3?expand=serviceDesk%2CrequestType%2Cparticipant%2Csla%2Cstatus%2Cattachment%2Caction",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			expands:            []string{"serviceDesk", "requestType", "participant", "sla", "status", "attachment", "action"},
			mockFile:           "./mocks/get-customer-request.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3?expand=serviceDesk%2CrequestType%2Cparticipant%2Csla%2Cstatus%2Cattachment%2Caction",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			expands:            []string{"serviceDesk", "requestType", "participant", "sla", "status", "attachment", "action"},
			mockFile:           "./mocks/get-customer-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3?expand=serviceDesk%2CrequestType%2Cparticipant%2Csla%2Cstatus%2Cattachment%2Caction",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			expands:            []string{"serviceDesk", "requestType", "participant", "sla", "status", "attachment", "action"},
			mockFile:           "./mocks/get-customer-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3?expand=serviceDesk%2CrequestType%2Cparticipant%2Csla%2Cstatus%2Cattachment%2Caction",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetCustomerRequestWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			expands:            []string{"serviceDesk", "requestType", "participant", "sla", "status", "attachment", "action"},
			mockFile:           "./mocks/get-customer-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			expands:            []string{"serviceDesk", "requestType", "participant", "sla", "status", "attachment", "action"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3?expand=serviceDesk%2CrequestType%2Cparticipant%2Csla%2Cstatus%2Cattachment%2Caction",
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

			service := &RequestService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.issueKeyOrID, testCase.expands)

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

				t.Log("-------------------------------------------")
				t.Logf("Custom Request Issue Key: %v", gotResult.IssueKey)
				t.Logf("Custom Request Type Name: %v", gotResult.RequestType.Name)
				t.Log("-------------------------------------------")
			}

		})
	}

}
