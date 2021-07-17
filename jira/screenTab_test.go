package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestScreenTabService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		screenID           int
		tabName            string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CreateScreenTabWhenTheParamsAreCorrect",
			screenID:           10001,
			tabName:            "Tab Name example",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateScreenTabWhenTheTabNameIsEmpty",
			screenID:           10001,
			tabName:            "",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateScreenTabWhenTheRequestMethodIsIncorrect",
			screenID:           10001,
			tabName:            "Tab Name example",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001/tabs",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateScreenTabWhenTheStatusCodeIsIncorrect",
			screenID:           10001,
			tabName:            "Tab Name example",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "CreateScreenTabWhenTheContextIsNil",
			screenID:           10001,
			tabName:            "Tab Name example",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateScreenTabWhenTheResponseBodyHasADifferentFormat",
			screenID:           10001,
			tabName:            "Tab Name example",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs",
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

			i := &ScreenTabService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.screenID, testCase.tabName)

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

				t.Log("------------------------------")
				t.Logf("Screen Tab Name: %v", gotResult.Name)
				t.Logf("Screen Tab ID: %v", gotResult.ID)
				t.Log("------------------------------ \n")
			}
		})

	}

}

func TestScreenTabService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		screenID           int
		tabID              int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteScreenTabWhenTheParamsAreCorrect",
			screenID:           10001,
			tabID:              1001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "DeleteScreenTabWhenTheContextIsNil",
			screenID:           10001,
			tabID:              1001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "DeleteScreenTabWhenTheEndpointIsIncorrect",
			screenID:           10001,
			tabID:              1001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/screens/10001/tabs/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "DeleteScreenTabWhenTheRequestMethodIsIncorrect",
			screenID:           10001,
			tabID:              1001,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "DeleteScreenTabWhenTheStatusCodeIsIncorrect",
			screenID:           10001,
			tabID:              1001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
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

			i := &ScreenTabService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.screenID, testCase.tabID)

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

func TestScreenTabService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		screenID           int
		projectKey         string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetScreenTabsWhenTheParamsAreCorrect",
			screenID:           10001,
			projectKey:         "DUMMY",
			mockFile:           "./mocks/get-screen-tabs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/tabs?projectKey=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetScreenTabsWhenTheProjectKeyIsEmpty",
			screenID:           10001,
			projectKey:         "",
			mockFile:           "./mocks/get-screen-tabs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/tabs",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetScreenTabsWhenTheRequestMethodIsIncorrect",
			screenID:           10001,
			projectKey:         "DUMMY",
			mockFile:           "./mocks/get-screen-tabs.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001/tabs?projectKey=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreenTabsWhenTheStatusCodeIsIncorrect",
			screenID:           10001,
			projectKey:         "DUMMY",
			mockFile:           "./mocks/get-screen-tabs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/tabs?projectKey=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetScreenTabsWhenTheContextIsNil",
			screenID:           10001,
			projectKey:         "DUMMY",
			mockFile:           "./mocks/get-screen-tabs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/tabs?projectKey=DUMMY",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreenTabsWhenTheEndpointIsIncorrect",
			screenID:           10001,
			projectKey:         "DUMMY",
			mockFile:           "./mocks/get-screen-tabs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/screens/10001/tabs?projectKey=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreenTabsWhenTheResponseBodyHasADifferentFormat",
			screenID:           10001,
			projectKey:         "DUMMY",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/tabs?projectKey=DUMMY",
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

			i := &ScreenTabService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.screenID, testCase.projectKey)

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

				for _, tab := range gotResult {
					t.Log("------------------------------")
					t.Logf("Screen Tab Name: %v", tab.Name)
					t.Logf("Screen Tab ID: %v", tab.ID)
					t.Log("------------------------------ \n")
				}

			}
		})

	}

}

func TestScreenTabService_Move(t *testing.T) {

	testCases := []struct {
		name                         string
		screenID, tabID, tabPosition int
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:               "MoveScreenTabWhenTheParamsAreCorrect",
			screenID:           10001,
			tabID:              1001,
			tabPosition:        21,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001/move/21",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "MoveScreenTabWhenTheContextIsNil",
			screenID:           10001,
			tabID:              1001,
			tabPosition:        21,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001/move/21",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "MoveScreenTabWhenTheStatusCodeIsIncorrect",
			screenID:           10001,
			tabID:              1001,
			tabPosition:        21,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001/move/21",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "MoveScreenTabWhenTheRequestMethodIsIncorrect",
			screenID:           10001,
			tabID:              1001,
			tabPosition:        21,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001/move/21",
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

			i := &ScreenTabService{client: mockClient}

			gotResponse, err := i.Move(testCase.context, testCase.screenID, testCase.tabID, testCase.tabPosition)

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

func TestScreenTabService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		screenID, tabID    int
		newTabName         string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "UpdateScreenTabWhenTheParamsAreCorrect",
			screenID:           10001,
			tabID:              1001,
			newTabName:         "Date Tracking",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "UpdateScreenTabWhenTheNewTabNameIsEmpty",
			screenID:           10001,
			tabID:              1001,
			newTabName:         "",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateScreenTabWhenTheRequestMethodIsIncorrect",
			screenID:           10001,
			tabID:              1001,
			newTabName:         "Date Tracking",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateScreenTabWhenTheStatusCodeIsIncorrect",
			screenID:           10001,
			tabID:              1001,
			newTabName:         "Date Tracking",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "UpdateScreenTabWhenTheEndpointIsIncorrect",
			screenID:           10001,
			tabID:              1001,
			newTabName:         "Date Tracking",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/screens/10001/tabs/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateScreenTabWhenTheContextIsNil",
			screenID:           10001,
			tabID:              1001,
			newTabName:         "Date Tracking",
			mockFile:           "./mocks/create-screen-tab.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateScreenTabWhenTheResponseBodyHasADifferentFormat",
			screenID:           10001,
			tabID:              1001,
			newTabName:         "Date Tracking",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001/tabs/1001",
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

			i := &ScreenTabService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.screenID, testCase.tabID, testCase.newTabName)

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

				t.Log("------------------------------")
				t.Logf("Screen Tab Name: %v", gotResult.Name)
				t.Logf("Screen Tab ID: %v", gotResult.ID)
				t.Log("------------------------------ \n")

			}
		})

	}

}
