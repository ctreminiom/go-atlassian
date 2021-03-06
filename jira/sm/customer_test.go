package sm

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestCustomerService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		email, displayName string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CreateCustomerWhenTheParamsAreCorrect",
			email:              "example@gmail.com",
			displayName:        "Example",
			mockFile:           "./mocks/create-customer.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateCustomerWhenTheParamsAreEmailIsNotSet",
			email:              "",
			displayName:        "Example",
			mockFile:           "./mocks/create-customer.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerWhenTheDisplayNameIsNotSet",
			email:              "example@gmail.com",
			displayName:        "",
			mockFile:           "./mocks/create-customer.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerWhenTheEmailIsIncorrect",
			email:              "examplegmail.com",
			displayName:        "Example",
			mockFile:           "./mocks/create-customer.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "CreateCustomerWhenTheRequestMethodIsIncorrect",
			email:              "example@gmail.com",
			displayName:        "Example",
			mockFile:           "./mocks/create-customer.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerWhenTheStatusCodeIsIncorrect",
			email:              "example@gmail.com",
			displayName:        "Example",
			mockFile:           "./mocks/create-customer.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerWhenTheContextIsNil",
			email:              "example@gmail.com",
			displayName:        "Example",
			mockFile:           "./mocks/create-customer.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/customer",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerWhenTheEndpointIsEmpty",
			email:              "example@gmail.com",
			displayName:        "Example",
			mockFile:           "./mocks/create-customer.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerWhenTheResponseBodyHasADifferentFormat",
			email:              "example@gmail.com",
			displayName:        "Example",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/customer",
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

			service := &CustomerService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.email, testCase.displayName)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				t.Log("--------------------------")
				t.Logf("Customer Name: %v", gotResult.Name)
				t.Logf("Customer Active: %v", gotResult.Active)
				t.Logf("Customer EmailAddress: %v", gotResult.EmailAddress)
				t.Logf("Customer Key: %v", gotResult.Key)
				t.Logf("Customer TimeZone: %v", gotResult.TimeZone)
				t.Log("--------------------------")
			}

		})
	}

}

func Test_isEmailValid(t *testing.T) {

	testCases := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{
			name:    "ValidateEmailWhenTheEmailIsCorrect",
			email:   "example@gmail.com",
			wantErr: false,
		},
		{
			name:    "ValidateEmailWhenTheEmailIsIncorrect",
			email:   "exampleeaasc",
			wantErr: true,
		},
		{
			name:    "ValidateEmailWhenTheEmailIsEmpty",
			email:   "",
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			isValid := isEmailValid(testCase.email)

			if testCase.wantErr {
				assert.Equal(t, false, isValid)
			} else {
				assert.Equal(t, true, isValid)
			}

		})
	}

}
