package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestRequestApprovalService_Answer(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		approvalID         int
		approve            bool
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AnswerRequestApprovalWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			approve:            true,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "AnswerRequestApprovalWhenTheParametersAreCorrectAndTheApproveIsDeclined",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			approve:            false,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "AnswerRequestApprovalWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			approvalID:         1,
			approve:            true,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AnswerRequestApprovalWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			approve:            true,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AnswerRequestApprovalWhenTheEndpointsIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			approve:            true,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AnswerRequestApprovalWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			approve:            true,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AnswerRequestApprovalWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			approve:            true,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AnswerRequestApprovalWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			approve:            true,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
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

			service := &RequestApprovalService{client: mockClient}
			gotResult, gotResponse, err := service.Answer(testCase.context, testCase.issueKeyOrID, testCase.approvalID, testCase.approve)

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

				t.Log(gotResult)
			}

		})
	}

}

func TestRequestApprovalService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		approvalID         int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetRequestApprovalWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetRequestApprovalWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			approvalID:         1,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalWhenTheEndpointsIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			mockFile:           "./mocks/get-approval.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			approvalID:         1,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval/1",
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

			service := &RequestApprovalService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.issueKeyOrID, testCase.approvalID)

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

				t.Log(gotResult)
			}

		})
	}

}

func TestRequestApprovalService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		start, limit       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetRequestApprovalsWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-approvals.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetRequestApprovalsWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-approvals.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalsWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-approvals.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalsWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-approvals.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalsWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-approvals.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalsWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-approvals.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval?limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestApprovalsWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/approval?limit=50&start=0",
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

			service := &RequestApprovalService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.issueKeyOrID, testCase.start, testCase.limit)

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

				for _, node := range gotResult.Values {
					t.Log(node)
				}
			}

		})
	}

}
