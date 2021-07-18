package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestRequestTypeService_Search(t *testing.T) {

	testCases := []struct {
		name               string
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
			name:               "SearchRequestApprovalWhenTheParametersAreCorrect",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/requesttype?limit=50&searchQuery=test&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "SearchRequestApprovalWhenTheRequestMethodIsIncorrect",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/requesttype?limit=50&searchQuery=test&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchRequestApprovalWhenTheStatusCodeIsIncorrect",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/requesttype?limit=50&searchQuery=test&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "SearchRequestApprovalWhenTheContextIsNil",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/requesttype?limit=50&searchQuery=test&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchRequestApprovalWhenTheEndpointIsIncorrect",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchRequestApprovalWhenTheResponseBodyHasADifferentFormat",
			query:              "test",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/requesttype?limit=50&searchQuery=test&start=0",
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

			service := &RequestTypeService{client: mockClient}
			gotResult, gotResponse, err := service.Search(testCase.context, testCase.query, testCase.start, testCase.limit)

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

				for _, requestType := range gotResult.Values {
					t.Log(requestType.ID, requestType.Name, requestType.Description)
				}
			}

		})
	}

}

func TestRequestTypeService_Create(t *testing.T) {

	testCases := []struct {
		name                                                string
		serviceDeskID                                       int
		issueTypeID, requestTypeName, description, helpText string
		mockFile                                            string
		wantHTTPMethod                                      string
		endpoint                                            string
		context                                             context.Context
		wantHTTPCodeReturn                                  int
		wantErr                                             bool
	}{
		{
			name:               "CreateRequestTypeWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			issueTypeID:        "10005",
			requestTypeName:    "Request Type Sample Name",
			description:        "Request Type Sample Description",
			helpText:           "Request Type Sample HelpText",
			mockFile:           "./mocks/create-project-request-type.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateRequestTypeWhenTheServiceDeskIDIsNotSet",
			issueTypeID:        "10005",
			requestTypeName:    "Request Type Sample Name",
			description:        "Request Type Sample Description",
			helpText:           "Request Type Sample HelpText",
			mockFile:           "./mocks/create-project-request-type.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateRequestTypeWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			issueTypeID:        "10005",
			requestTypeName:    "Request Type Sample Name",
			description:        "Request Type Sample Description",
			helpText:           "Request Type Sample HelpText",
			mockFile:           "./mocks/create-project-request-type.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateRequestTypeWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			issueTypeID:        "10005",
			requestTypeName:    "Request Type Sample Name",
			description:        "Request Type Sample Description",
			helpText:           "Request Type Sample HelpText",
			mockFile:           "./mocks/create-project-request-type.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "CreateRequestTypeWhenTheContextIsNil",
			serviceDeskID:      1,
			issueTypeID:        "10005",
			requestTypeName:    "Request Type Sample Name",
			description:        "Request Type Sample Description",
			helpText:           "Request Type Sample HelpText",
			mockFile:           "./mocks/create-project-request-type.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateRequestTypeWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			issueTypeID:        "10005",
			requestTypeName:    "Request Type Sample Name",
			description:        "Request Type Sample Description",
			helpText:           "Request Type Sample HelpText",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype",
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

			service := &RequestTypeService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.serviceDeskID, testCase.issueTypeID,
				testCase.requestTypeName,
				testCase.description,
				testCase.helpText)

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

				t.Log("newRequestType", gotResult.ID, gotResult.Name)
			}

		})
	}

}

func TestRequestTypeService_Delete(t *testing.T) {

	testCases := []struct {
		name                         string
		serviceDeskID, requestTypeID int
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{

		{
			name:               "DeleteRequestTypeWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteRequestTypeWhenTheServiceDeskIDIsNotSet",
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteRequestTypeWhenTheRequestTypeIDIsNotSet",
			serviceDeskID:      1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteRequestTypeWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteRequestTypeWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteRequestTypeWhenTheContextIsNil",
			serviceDeskID:      1,
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteRequestTypeWhenTheEndpointIsEmpty",
			serviceDeskID:      1,
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
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

			service := &RequestTypeService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.serviceDeskID, testCase.requestTypeID)

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

func TestRequestTypeService_Fields(t *testing.T) {
	testCases := []struct {
		name                         string
		serviceDeskID, requestTypeID int
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:               "GetRequestTypeFieldsWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			requestTypeID:      1,
			mockFile:           "./mocks/get-request-type-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetRequestTypeFieldsWhenTheServiceDeskIDIsNotSet",
			requestTypeID:      1,
			mockFile:           "./mocks/get-request-type-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypeFieldsWhenTheRequestTypeIDIsNotSet",
			serviceDeskID:      1,
			mockFile:           "./mocks/get-request-type-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypeFieldsWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			requestTypeID:      1,
			mockFile:           "./mocks/get-request-type-fields.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypeFieldsWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			requestTypeID:      1,
			mockFile:           "./mocks/get-request-type-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1/field",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypeFieldsWhenTheContextIsNil",
			serviceDeskID:      1,
			requestTypeID:      1,
			mockFile:           "./mocks/get-request-type-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1/field",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypeFieldsWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			requestTypeID:      1,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1/field",
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

			service := &RequestTypeService{client: mockClient}
			gotResult, gotResponse, err := service.Fields(testCase.context, testCase.serviceDeskID, testCase.requestTypeID)

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

			}

		})
	}
}

func TestRequestTypeService_Get(t *testing.T) {
	testCases := []struct {
		name                         string
		serviceDeskID, requestTypeID int
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:               "GetRequestTypeWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			mockFile:           "./mocks/get-project-request-type.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetRequestTypeWhenTheServiceDeskIDIsNotSet",
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			mockFile:           "./mocks/get-project-request-type.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypeWhenTheRequesTypeIDIsNotSet",
			serviceDeskID:      1,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			mockFile:           "./mocks/get-project-request-type.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypeWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			mockFile:           "./mocks/get-project-request-type.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypeWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			mockFile:           "./mocks/get-project-request-type.json",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypeWhenTheContextIsNil",
			serviceDeskID:      1,
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			mockFile:           "./mocks/get-project-request-type.json",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypeWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			requestTypeID:      1,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype/1",
			mockFile:           "./mocks/empty_json.json",
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

			service := &RequestTypeService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.serviceDeskID, testCase.requestTypeID)

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

				t.Log("newRequestType", gotResult.ID, gotResult.Name)
			}

		})
	}
}

func TestRequestTypeService_Gets(t *testing.T) {
	testCases := []struct {
		name                                 string
		serviceDeskID, groupID, start, limit int
		mockFile                             string
		wantHTTPMethod                       string
		endpoint                             string
		context                              context.Context
		wantHTTPCodeReturn                   int
		wantErr                              bool
	}{
		{
			name:               "GetRequestTypesWhenTheParametersAreCorrect",
			serviceDeskID:      1,
			groupID:            10001,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype?groupId=10001&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetRequestTypesWhenTheServiceDeskIDIsNotSet",
			groupID:            10001,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype?groupId=10001&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypesWhenTheRequestMethodIsIncorrect",
			serviceDeskID:      1,
			groupID:            10001,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype?groupId=10001&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetRequestTypesWhenTheStatusCodeIsIncorrect",
			serviceDeskID:      1,
			groupID:            10001,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype?groupId=10001&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypesWhenTheContextIsNil",
			serviceDeskID:      1,
			groupID:            10001,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype?groupId=10001&limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetRequestTypesWhenTheEndpointIsEmpty",
			serviceDeskID:      1,
			groupID:            10001,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-request-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestTypesWhenTheResponseBodyHasADifferentFormat",
			serviceDeskID:      1,
			groupID:            10001,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/servicedesk/1/requesttype?groupId=10001&limit=50&start=0",
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

			service := &RequestTypeService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.serviceDeskID, testCase.groupID,
				testCase.start, testCase.limit)

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

				for _, requestType := range gotResult.Values {
					t.Log(requestType.Name, requestType.Name)
				}
			}

		})
	}
}
