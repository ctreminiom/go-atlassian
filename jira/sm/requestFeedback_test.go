package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestRequestFeedbackService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		requestIDOrKey     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteCustomerFeedbackWhenTheParametersAreCorrect",
			requestIDOrKey:     "DUMMY-4",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteCustomerFeedbackWhenTheRequestIDOrKeyIsNotSet",
			requestIDOrKey:     "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteCustomerFeedbackWhenTheRequestMethodIsIncorrect",
			requestIDOrKey:     "DUMMY-4",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteCustomerFeedbackWhenTheStatusCodeIsIncorrect",
			requestIDOrKey:     "DUMMY-4",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteCustomerFeedbackWhenTheContextIsNil",
			requestIDOrKey:     "DUMMY-4",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
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

			service := &RequestFeedbackService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.requestIDOrKey)

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

func TestRequestFeedbackService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		requestIDOrKey     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetCustomerFeedbackWhenTheParametersAreCorrect",
			requestIDOrKey:     "DUMMY-4",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerFeedbackWhenTheRequestIDOrKeyIsNotSet",
			requestIDOrKey:     "",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerFeedbackWhenTheRequestMethodIsIncorrect",
			requestIDOrKey:     "DUMMY-4",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetCustomerFeedbackWhenTheStatusCodeIsIncorrect",
			requestIDOrKey:     "DUMMY-4",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetCustomerFeedbackWhenTheContextIsNil",
			requestIDOrKey:     "DUMMY-4",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerFeedbackWhenTheResponseBodyHasADifferentFormat",
			requestIDOrKey:     "DUMMY-4",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
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

			service := &RequestFeedbackService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.requestIDOrKey)

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

				t.Logf("Customer Feedback Rating: %v", gotResult.Rating)
				t.Logf("Customer Feddback Comment: %v", gotResult.Comment.Body)
			}

		})
	}

}

func TestRequestFeedbackService_Post(t *testing.T) {

	testCases := []struct {
		name               string
		requestIDOrKey     string
		rating             int
		comment            string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetCustomerFeedbackWhenTheParametersAreCorrect",
			requestIDOrKey:     "DUMMY-4",
			rating:             5,
			comment:            "Sample comment",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "GetCustomerFeedbackWhenTheRequestIDOrKeyIsNotSet",
			requestIDOrKey:     "",
			rating:             5,
			comment:            "Sample comment",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "GetCustomerFeedbackWhenTheCommentIsNotSet",
			requestIDOrKey:     "DUMMY-4",
			rating:             5,
			comment:            "",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "GetCustomerFeedbackWhenTheRequestMethodIsIncorrect",
			requestIDOrKey:     "DUMMY-4",
			rating:             5,
			comment:            "Sample comment",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "GetCustomerFeedbackWhenTheStatusCodeIsIncorrect",
			requestIDOrKey:     "DUMMY-4",
			rating:             5,
			comment:            "Sample comment",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetCustomerFeedbackWhenTheEndpointIsEmpty",
			requestIDOrKey:     "DUMMY-4",
			rating:             5,
			comment:            "Sample comment",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "GetCustomerFeedbackWhenTheContextIsNil",
			requestIDOrKey:     "DUMMY-4",
			rating:             5,
			comment:            "Sample comment",
			mockFile:           "./mocks/get-customer-feedback.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "GetCustomerFeedbackWhenTheResponseBodyHasADifferentFormat",
			requestIDOrKey:     "DUMMY-4",
			rating:             5,
			comment:            "Sample comment",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-4/feedback",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
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

			service := &RequestFeedbackService{client: mockClient}
			gotResult, gotResponse, err := service.Post(testCase.context, testCase.requestIDOrKey, testCase.rating, testCase.comment)

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

				t.Logf("Customer Feedback Rating: %v", gotResult.Rating)
				t.Logf("Customer Feddback Comment: %v", gotResult.Comment.Body)
			}

		})
	}

}
