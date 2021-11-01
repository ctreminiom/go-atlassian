package v3

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestScreenSchemeService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *ScreenSchemePayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateScreenSchemeWhenTheParamsAreCorrect",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "Screen Scheme Name",
				Description: "Screen Scheme Description",
			},
			mockFile:           "./mocks/create-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name: "CreateScreenSchemeWhenTheContextIsNil",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "Screen Scheme Name",
				Description: "Screen Scheme Description",
			},
			mockFile:           "./mocks/create-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screenscheme",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateScreenSchemeWhenThePayloadIsNil",
			payload:            nil,
			mockFile:           "./mocks/create-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateScreenSchemeWhenTheRequestMethodIsIncorrect",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "Screen Scheme Name",
				Description: "Screen Scheme Description",
			},
			mockFile:           "./mocks/create-screen-scheme.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateScreenSchemeWhenTheStatusCodeIsIncorrect",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "Screen Scheme Name",
				Description: "Screen Scheme Description",
			},
			mockFile:           "./mocks/create-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CreateScreenSchemeWhenTheEndpointIsIncorrect",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "Screen Scheme Name",
				Description: "Screen Scheme Description",
			},
			mockFile:           "./mocks/create-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/screenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateScreenSchemeWhenTheResponseBodyHasADifferentFormat",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "Screen Scheme Name",
				Description: "Screen Scheme Description",
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screenscheme",
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

			i := &ScreenSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.payload)

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
				t.Logf("Screen Screen Name: %v", gotResult.Name)
				t.Logf("Screen Screen Description: %v", gotResult.Description)
				t.Logf("Screen Screen ID: %v", gotResult.ID)
				t.Log("------------------------------ \n")

			}
		})

	}

}

func TestScreenSchemeService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		screenSchemeID     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteScreenSchemeWhenTheParamsAreCorrect",
			screenSchemeID:     "1002",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screenscheme/1002",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteScreenSchemeWhenTheScreenSchemeIDIsEmpty",
			screenSchemeID:     "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screenscheme/1002",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteScreenSchemeWhenTheRequestMethodIsIncorrect",
			screenSchemeID:     "1002",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screenscheme/1002",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteScreenSchemeWhenTheStatusCodeIsIncorrect",
			screenSchemeID:     "1002",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screenscheme/1002",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteScreenSchemeWhenTheContextIsNil",
			screenSchemeID:     "1002",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/screenscheme/1002",
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

			i := &ScreenSchemeService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.screenSchemeID)

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

func TestScreenSchemeService_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		screenSchemeIDs     []int
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetScreenSchemesWhenTheParamsAreCorrect",
			screenSchemeIDs:    []int{1000, 1001, 1002, 1003},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screenscheme?id=1000&id=1001&id=1002&id=1003&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetScreenSchemesWhenTheScreenSchemeIDsIsNil",
			screenSchemeIDs:    nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screenscheme?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetScreenSchemesWhenTheContextIsNil",
			screenSchemeIDs:    []int{1000, 1001, 1002, 1003},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screenscheme?id=1000&id=1001&id=1002&id=1003&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreenSchemesWhenTheRequestMethodIsIncorrect",
			screenSchemeIDs:    []int{1000, 1001, 1002, 1003},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screen-schemes.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/screenscheme?id=1000&id=1001&id=1002&id=1003&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetScreenSchemesWhenTheStatusCodeIsIncorrect",
			screenSchemeIDs:    []int{1000, 1001, 1002, 1003},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screenscheme?id=1000&id=1001&id=1002&id=1003&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetScreenSchemesWhenTheResponseBodyHasADifferentFormat",
			screenSchemeIDs:    []int{1000, 1001, 1002, 1003},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screenscheme?id=1000&id=1001&id=1002&id=1003&maxResults=50&startAt=0",
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

			i := &ScreenSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.screenSchemeIDs, testCase.startAt,
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

				for _, screenScheme := range gotResult.Values {

					t.Log("------------------------------")
					t.Logf("Screen Screen Name: %v", screenScheme.Name)
					t.Logf("Screen Screen Description: %v", screenScheme.Description)
					t.Logf("Screen Screen ID: %v", screenScheme.ID)
					t.Log("------------------------------ \n")
				}

			}
		})

	}

}

func TestScreenSchemeService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		screenSchemeID     string
		payload            *ScreenSchemePayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:           "UpdateScreenSchemeWhenTheParamsAreCorrect",
			screenSchemeID: "2001",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "FX | Epic Screen Scheme",
				Description: "sample description",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screenscheme/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:           "UpdateScreenSchemeWhenTheScreenSchemeIDIsEmpty",
			screenSchemeID: "",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "FX | Epic Screen Scheme",
				Description: "sample description",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screenscheme/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "UpdateScreenSchemeWhenThePayloadIsNil",
			screenSchemeID:     "2001",
			payload:            nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screenscheme/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:           "UpdateScreenSchemeWhenTheRequestMethodIsIncorrect",
			screenSchemeID: "2001",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "FX | Epic Screen Scheme",
				Description: "sample description",
			},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/screenscheme/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:           "UpdateScreenSchemeWhenTheStatusCodeIsIncorrect",
			screenSchemeID: "2001",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "FX | Epic Screen Scheme",
				Description: "sample description",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screenscheme/2001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:           "UpdateScreenSchemeWhenTheContextIsNil",
			screenSchemeID: "2001",
			payload: &ScreenSchemePayloadScheme{
				Screens: &ScreenTypesScheme{
					Default: 10000,
					View:    10000,
					Edit:    10000,
				},
				Name:        "FX | Epic Screen Scheme",
				Description: "sample description",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/screenscheme/2001",
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

			i := &ScreenSchemeService{client: mockClient}

			gotResponse, err := i.Update(testCase.context, testCase.screenSchemeID, testCase.payload)

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
