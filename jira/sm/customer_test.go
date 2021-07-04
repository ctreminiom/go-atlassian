package sm

import (
	"context"
	"fmt"
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
			name:               "CreateCustomerWhenTheEmailIsToShort",
			email:              "exa",
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
			name:    "ValidateEmailWhenTheEmailDoesNotHaveFormat",
			email:   "ex",
			wantErr: true,
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

func TestCustomerService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		serviceDeskID      int
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
			name:               "GetProjectCustomersWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			query:              "Charlie",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-project-customers.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer?limit=50&query=Charlie&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectCustomersWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			query:              "Charlie",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-project-customers.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer?limit=50&query=Charlie&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectCustomersWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			query:              "Charlie",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-project-customers.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer?limit=50&query=Charlie&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectCustomersWhenTheContextIsNil",
			serviceDeskID:      1,
			query:              "Charlie",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-project-customers.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer?limit=50&query=Charlie&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetProjectCustomersWhenTheEndpointIsEmpty",
			serviceDeskID:      1,
			query:              "Charlie",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-project-customers.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectCustomersWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			query:              "Charlie",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer?limit=50&query=Charlie&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectCustomersWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			query:              "Charlie",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-project-customers.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer?limit=50&query=Charlie&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
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
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.serviceDeskID, testCase.query, testCase.start, testCase.limit)

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

				for _, customer := range gotResult.Values {

					t.Log("--------------------------")
					t.Logf("Customer Name: %v", customer.Name)
					t.Logf("Customer Active: %v", customer.Active)
					t.Logf("Customer EmailAddress: %v", customer.EmailAddress)
					t.Logf("Customer Key: %v", customer.Key)
					t.Logf("Customer TimeZone: %v", customer.TimeZone)
					t.Log("--------------------------")
				}

			}

		})
	}

}

func TestCustomerService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		serviceDeskID      int
		accountIDs         []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddCustomerToProjectWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			accountIDs:         []string{"f81a5efd-992a-4f3c-b079-95ea536c1205", "ed7e6547-9677-497f-8f7d-89021fdd8fdb"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "AddCustomerToProjectWhenTheContextIsNil",
			serviceDeskID:      1,
			accountIDs:         []string{"f81a5efd-992a-4f3c-b079-95ea536c1205", "ed7e6547-9677-497f-8f7d-89021fdd8fdb"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddCustomerToProjectWhenTheAccountIDsAreNotSet",
			serviceDeskID:      1,
			accountIDs:         nil,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddCustomerToProjectWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			accountIDs:         []string{"f81a5efd-992a-4f3c-b079-95ea536c1205", "ed7e6547-9677-497f-8f7d-89021fdd8fdb"},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddCustomerToProjectWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			accountIDs:         []string{"f81a5efd-992a-4f3c-b079-95ea536c1205", "ed7e6547-9677-497f-8f7d-89021fdd8fdb"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AddCustomerToProjectWhenTheContextIsNil",
			serviceDeskID:      1,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "AddCustomerToProjectWhenTheEndpointIsEmpty",
			serviceDeskID:      1,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AddCustomerToProjectWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            context.Background(),
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

			service := &CustomerService{client: mockClient}
			gotResponse, err := service.Add(testCase.context, testCase.serviceDeskID, testCase.accountIDs)

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

func TestCustomerService_Remove(t *testing.T) {

	testCases := []struct {
		name               string
		serviceDeskID      int
		accountIDs         []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "RemoveCustomerToProjectWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			accountIDs:         []string{"f81a5efd-992a-4f3c-b079-95ea536c1205", "ed7e6547-9677-497f-8f7d-89021fdd8fdb"},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "RemoveCustomerToProjectWhenTheContextIsNil",
			serviceDeskID:      1,
			accountIDs:         []string{"f81a5efd-992a-4f3c-b079-95ea536c1205", "ed7e6547-9677-497f-8f7d-89021fdd8fdb"},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerToProjectWhenTheEndpointIsNotValid",
			serviceDeskID:      1,
			accountIDs:         []string{"f81a5efd-992a-4f3c-b079-95ea536c1205", "ed7e6547-9677-497f-8f7d-89021fdd8fdb"},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           `ssadasf81a5efd-992a-4f3c-b079-95ea536c1205///asd33>>>sas`,
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerToProjectWhenTheAccountIDsAreNotSet",
			serviceDeskID:      1,
			accountIDs:         nil,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerToProjectWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			accountIDs:         []string{"f81a5efd-992a-4f3c-b079-95ea536c1205", "ed7e6547-9677-497f-8f7d-89021fdd8fdb"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerToProjectWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			accountIDs:         []string{"f81a5efd-992a-4f3c-b079-95ea536c1205", "ed7e6547-9677-497f-8f7d-89021fdd8fdb"},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerToProjectWhenTheContextIsNil",
			serviceDeskID:      1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "RemoveCustomerToProjectWhenTheEndpointIsEmpty",
			serviceDeskID:      1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveCustomerToProjectWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/customer",
			context:            context.Background(),
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

			service := &CustomerService{client: mockClient}
			gotResponse, err := service.Remove(testCase.context, testCase.serviceDeskID, testCase.accountIDs)

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
