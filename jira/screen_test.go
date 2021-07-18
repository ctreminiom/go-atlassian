package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestScreenService_AddToDefault(t *testing.T) {

	testCases := []struct {
		name               string
		fieldID            string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddFieldToDefaultScreenWhenTheParamsAreCorrect",
			fieldID:            "customfield_10032",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/addToDefault/customfield_10032",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "AddFieldToDefaultScreenWhenTheFieldIDIsEmpty",
			fieldID:            "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/addToDefault/customfield_10032",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddFieldToDefaultScreenWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10032",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/addToDefault/customfield_10032",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "AddFieldToDefaultScreenWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10032",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/addToDefault/customfield_10032",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AddFieldToDefaultScreenWhenTheContextIsNil",
			fieldID:            "customfield_10032",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/addToDefault/customfield_10032",
			context:            nil,
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

			i := &ScreenService{client: mockClient}

			gotResponse, err := i.AddToDefault(testCase.context, testCase.fieldID)

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

func TestScreenService_Available(t *testing.T) {

	testCases := []struct {
		name               string
		screenID           int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetFieldsAvailableOnScreenWhenTheParamsAreCorrect",
			screenID:           10001,
			mockFile:           "./mocks/get-available-screen-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/availableFields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetFieldsAvailableOnScreenWhenTheRequestMethodIsIncorrect",
			screenID:           10001,
			mockFile:           "./mocks/get-available-screen-fields.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens/10001/availableFields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetFieldsAvailableOnScreenWhenTheStatusCodeIsIncorrect",
			screenID:           10001,
			mockFile:           "./mocks/get-available-screen-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/availableFields",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetFieldsAvailableOnScreenWhenTheContextIsNil",
			screenID:           10001,
			mockFile:           "./mocks/get-available-screen-fields.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/availableFields",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetFieldsAvailableOnScreenWhenTheResponseBodyHasADifferentFormat",
			screenID:           10001,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001/availableFields",
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

			i := &ScreenService{client: mockClient}

			gotResult, gotResponse, err := i.Available(testCase.context, testCase.screenID)

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

				for _, field := range gotResult {

					t.Log("------------------------------")
					t.Logf("Field Name: %v", field.Name)
					t.Logf("Field ID: %v", field.ID)
					t.Log("------------------------------ \n")

				}

			}
		})

	}

}

func TestScreenService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		screenName         string
		screenDescription  string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CreateScreenWhenTheParamsAreCorrect",
			screenName:         "DUMMY Bug screen",
			screenDescription:  "This's the bug screen description",
			mockFile:           "./mocks/create-screen.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateScreenWhenTheScreenNameIsEmpty",
			screenName:         "",
			screenDescription:  "This's the bug screen description",
			mockFile:           "./mocks/create-screen.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateScreenWhenTheRequestMethodIsIncorrect",
			screenName:         "DUMMY Bug screen",
			screenDescription:  "This's the bug screen description",
			mockFile:           "./mocks/create-screen.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateScreenWhenTheStatusCodeIsIncorrect",
			screenName:         "DUMMY Bug screen",
			screenDescription:  "This's the bug screen description",
			mockFile:           "./mocks/create-screen.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "CreateScreenWhenTheEndpointIsIncorrect",
			screenName:         "DUMMY Bug screen",
			screenDescription:  "This's the bug screen description",
			mockFile:           "./mocks/create-screen.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/screens",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateScreenWhenTheContextIsNil",
			screenName:         "DUMMY Bug screen",
			screenDescription:  "This's the bug screen description",
			mockFile:           "./mocks/create-screen.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateScreenWhenTheResponseBodyHasADifferentFormat",
			screenName:         "DUMMY Bug screen",
			screenDescription:  "This's the bug screen description",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screens",
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

			i := &ScreenService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.screenName, testCase.screenDescription)

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
				t.Logf("Screen Name: %v", gotResult.Name)
				t.Logf("Screen ID: %v", gotResult.ID)
				t.Logf("Screen Description: %v", gotResult.Description)
				t.Log("------------------------------ \n")

			}
		})

	}

}

func TestScreenService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		screenID           int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteScreenWhenTheParamsAreCorrect",
			screenID:           10001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteScreenWhenTheContextIsNil",
			screenID:           10001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteScreenWhenTheRequestMethodIsIncorrect",
			screenID:           10001,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteScreenWhenTheStatusCodeIsIncorrect",
			screenID:           10001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteScreenWhenTheEndpointIsIncorrect",
			screenID:           10001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/screens/10001",
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

			i := &ScreenService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.screenID)

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

func TestScreenService_Get(t *testing.T) {

	testCases := []struct {
		name                string
		fieldID             string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetScreensForFieldWhenTheParamsAreCorrect",
			fieldID:            "customfield_10032",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens-for-field.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10032/screens?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetScreensForFieldWhenTheFieldIsIsEmpty",
			fieldID:            "",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens-for-field.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10032/screens?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreensForFieldWhenTheRequestMethodIsIncorrect",
			fieldID:            "customfield_10032",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens-for-field.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/field/customfield_10032/screens?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreensForFieldWhenTheStatusCodeIsIncorrect",
			fieldID:            "customfield_10032",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens-for-field.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10032/screens?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetScreensForFieldWhenTheContextIsNil",
			fieldID:            "customfield_10032",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens-for-field.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10032/screens?maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreensForFieldWhenTheEndpointIsIncorrect",
			fieldID:            "customfield_10032",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens-for-field.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/field/customfield_10032/screens?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreensForFieldWhenTheResponseBodyHasADifferentFormat",
			fieldID:            "customfield_10032",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/field/customfield_10032/screens?maxResults=50&startAt=0",
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

			i := &ScreenService{client: mockClient}

			gotResult, gotResponse, err := i.Fields(testCase.context, testCase.fieldID, testCase.startAt,
				testCase.maxResults)

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

				for _, screen := range gotResult.Values {

					t.Log("------------------------------")
					t.Logf("Screen Name: %v", screen.Name)
					t.Logf("Screen ID: %v", screen.ID)
					t.Log("------------------------------ \n")
				}

			}
		})

	}

}

func TestScreenService_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		screenIDs           []int
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetScreensWhenTheParamsAreCorrect",
			screenIDs:          []int{10001, 10002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens?id=10001&id=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetScreensWhenTheScreenIDsIsNil",
			screenIDs:          nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetScreensWhenTheRequestMethodIsIncorrect",
			screenIDs:          []int{10001, 10002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screens?id=10001&id=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreensWhenTheStatusCodeIsIncorrect",
			screenIDs:          []int{10001, 10002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens?id=10001&id=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetScreensWhenTheContextIsNil",
			screenIDs:          []int{10001, 10002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens?id=10001&id=10002&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreensWhenTheEndpointIsIncorrect",
			screenIDs:          []int{10001, 10002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screens.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/screens?id=10001&id=10002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreensWhenTheResponseBodyHasADifferentFormat",
			screenIDs:          []int{10001, 10002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens?id=10001&id=10002&maxResults=50&startAt=0",
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

			i := &ScreenService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.screenIDs, testCase.startAt, testCase.maxResults)

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

				for _, screen := range gotResult.Values {

					t.Log("------------------------------")
					t.Logf("Screen Name: %v", screen.Name)
					t.Logf("Screen ID: %v", screen.ID)
					t.Log("------------------------------ \n")
				}

			}
		})

	}

}

func TestScreenService_Update(t *testing.T) {

	testCases := []struct {
		name                          string
		screenID                      int
		screenName, screenDescription string
		mockFile                      string
		wantHTTPMethod                string
		endpoint                      string
		context                       context.Context
		wantHTTPCodeReturn            int
		wantErr                       bool
	}{
		{
			name:               "GetScreensForFieldWhenTheParamsAreCorrect",
			screenID:           10001,
			screenName:         "New screen name",
			screenDescription:  "New screen description",
			mockFile:           "./mocks/update-screen.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetScreensForFieldWhenTheRequestMethodIsIncorrect",
			screenID:           10001,
			screenName:         "New screen name",
			screenDescription:  "New screen description",
			mockFile:           "./mocks/update-screen.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screens/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreensForFieldWhenTheStatusCodeIsIncorrect",
			screenID:           10001,
			screenName:         "New screen name",
			screenDescription:  "New screen description",
			mockFile:           "./mocks/update-screen.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetScreensForFieldWhenTheEndpointIsIncorrect",
			screenID:           10001,
			screenName:         "New screen name",
			screenDescription:  "New screen description",
			mockFile:           "./mocks/update-screen.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/screens/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreensForFieldWhenTheContextIsNil",
			screenID:           10001,
			screenName:         "New screen name",
			screenDescription:  "New screen description",
			mockFile:           "./mocks/update-screen.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreensForFieldWhenTheResponseBodyHasADifferentFormat",
			screenID:           10001,
			screenName:         "New screen name",
			screenDescription:  "New screen description",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screens/10001",
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

			i := &ScreenService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.screenID, testCase.screenName,
				testCase.screenDescription)

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
				t.Logf("Screen Name: %v", gotResult.Name)
				t.Logf("Screen ID: %v", gotResult.ID)
				t.Log("------------------------------ \n")

			}
		})

	}

}
