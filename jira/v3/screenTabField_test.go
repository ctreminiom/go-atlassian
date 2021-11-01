package v3

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestScreenTabFieldService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		screenID, tabID    int
		fieldID            string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddScreenTabFieldWhenTheParamsAreCorrect",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/add-field-to-screen-tab.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "AddScreenTabFieldWhenTheFieldIsIsEmpty",
			fieldID:            "",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/add-field-to-screen-tab.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddScreenTabFieldWhenTheRequestMethodIsIncorrect",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/add-field-to-screen-tab.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddScreenTabFieldWhenTheStatusCodeIsIncorrect",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/add-field-to-screen-tab.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AddScreenTabFieldWhenTheEndpointIsIncorrect",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/add-field-to-screen-tab.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/screens/10001/tabs/12/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddScreenTabFieldWhenTheContextIsNil",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/add-field-to-screen-tab.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddScreenTabFieldWhenTheResponseBodyHasADifferentFormat",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
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

			i := &ScreenTabFieldService{client: mockClient}

			gotResult, gotResponse, err := i.Add(testCase.context, testCase.screenID, testCase.tabID, testCase.fieldID)

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

func TestScreenTabFieldService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		screenID, tabID    int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetScreenTabFieldsWhenTheParamsAreCorrect",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/get-screen-field-tabs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetScreenTabFieldsWhenTheEndpointIsIncorrect",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/get-screen-field-tabs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10002/tabs/12/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreenTabFieldsWhenTheRequestMethodIsIncorrect",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/get-screen-field-tabs.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreenTabFieldsWhenTheStatusCodeIsIncorrect",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/get-screen-field-tabs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetScreenTabFieldsWhenTheContextIsNil",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/get-screen-field-tabs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreenTabFieldsWhenTheResponseBodyHasADifferentFormat",
			screenID:           10001,
			tabID:              12,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields",
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

			i := &ScreenTabFieldService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.screenID, testCase.tabID)

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

func TestScreenTabFieldService_Remove(t *testing.T) {

	testCases := []struct {
		name               string
		screenID, tabID    int
		fieldID            string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "RemoveFieldFromScreenTabWhenTheParamsAreCorrect",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "RemoveFieldFromScreenTabWhenTheFieldIDIsEmpty",
			fieldID:            "",
			screenID:           10001,
			tabID:              12,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveFieldFromScreenTabWhenTheRequestMethodIsIncorrect",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveFieldFromScreenTabWhenTheStatusCodeIsIncorrect",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "RemoveFieldFromScreenTabWhenTheContextIsNil",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001/tabs/12/fields/1000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveFieldFromScreenTabWhenTheEndpointIsIncorrect",
			fieldID:            "1000",
			screenID:           10001,
			tabID:              12,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/screens/10001/tabs/12/fields/1000",
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

			i := &ScreenTabFieldService{client: mockClient}

			gotResponse, err := i.Remove(testCase.context, testCase.screenID, testCase.tabID, testCase.fieldID)

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
