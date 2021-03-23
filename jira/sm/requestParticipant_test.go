package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestRequestParticipantService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		accountIDs         []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddCustomerRequestParticipantsWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "AddCustomerRequestParticipantsWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddCustomerRequestParticipantsWhenTheAccountIDsAreNotSet",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         nil,
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddCustomerRequestParticipantsWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddCustomerRequestParticipantsWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AddCustomerRequestParticipantsWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddCustomerRequestParticipantsWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddCustomerRequestParticipantsWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
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

			service := &RequestParticipantService{client: mockClient}
			gotResult, gotResponse, err := service.Add(testCase.context, testCase.issueKeyOrID, testCase.accountIDs)

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

				for _, participant := range gotResult.Values {

					t.Log("----------------------------------")
					t.Logf("Comment, ID: %v", participant.AccountID)
					t.Logf("Comment, Creator Name: %v", participant.DisplayName)
					t.Logf("Comment, Created Date: %v", participant.EmailAddress)
					t.Log("----------------------------------")

				}

			}

		})
	}

}

func TestRequestParticipantService_Gets(t *testing.T) {

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
			name:               "GetCustomerRequestParticipantsWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-4",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerRequestParticipantsWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestParticipantsWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-4",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestParticipantsWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-4",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestParticipantsWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-4",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant?limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerRequestParticipantsWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-4",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant?limit=50&start=0",
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

			service := &RequestParticipantService{client: mockClient}
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

				for _, participant := range gotResult.Values {

					t.Log("----------------------------------")
					t.Logf("Comment, ID: %v", participant.AccountID)
					t.Logf("Comment, Creator Name: %v", participant.DisplayName)
					t.Logf("Comment, Created Date: %v", participant.EmailAddress)
					t.Log("----------------------------------")

				}

			}

		})
	}

}

func TestRequestParticipantService_Remove(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		accountIDs         []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "RemoveCustomerRequestParticipantsWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "RemoveCustomerRequestParticipantsWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerRequestParticipantsWhenTheAccountIDsAreNotSet",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         nil,
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerRequestParticipantsWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerRequestParticipantsWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerRequestParticipantsWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerRequestParticipantsWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/get-request-participants.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerRequestParticipantsWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-4",
			accountIDs:         []string{"62bd1d15-aeb7-4975-b7d9-40214e226e18", "86981eef-188d-435b-80cf-c29a62b23959"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/participant",
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

			service := &RequestParticipantService{client: mockClient}
			gotResult, gotResponse, err := service.Remove(testCase.context, testCase.issueKeyOrID, testCase.accountIDs)

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

				for _, participant := range gotResult.Values {

					t.Log("----------------------------------")
					t.Logf("Comment, ID: %v", participant.AccountID)
					t.Logf("Comment, Creator Name: %v", participant.DisplayName)
					t.Logf("Comment, Created Date: %v", participant.EmailAddress)
					t.Log("----------------------------------")

				}

			}

		})
	}

}
