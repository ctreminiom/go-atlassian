package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestProjectValidationService_Key(t *testing.T) {

	testCases := []struct {
		name               string
		projectKey         string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "ValidateProjectKeyWhenTheParamsAreCorrect",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/validate-project-key.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/validProjectKey?key=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "ValidateProjectKeyWhenTheProjectKeyIsEmpty",
			projectKey:         "",
			mockFile:           "./mocks/validate-project-key.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/validProjectKey?key=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheRequestMethodIsIncorrect",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/validate-project-key.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/projectvalidate/validProjectKey?key=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheStatusCodeIsIncorrect",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/validate-project-key.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/validProjectKey?key=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheEndpointsIsIncorrect",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/validate-project-key.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectvalidate/validProjectKey?key=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheContextIsNil",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/validate-project-key.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/validProjectKey?key=DUMMY",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheResponseBodyHasADifferentFormat",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/validProjectKey?key=DUMMY",
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

			i := &ProjectValidationService{client: mockClient}

			gotResult, gotResponse, err := i.Key(testCase.context, testCase.projectKey)

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

				t.Log("-------------------------------------")
				t.Logf("Project key: %v", gotResult)
				t.Log("-------------------------------------")

			}
		})

	}

}

func TestProjectValidationService_Name(t *testing.T) {

	testCases := []struct {
		name               string
		projectName        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "ValidateProjectKeyWhenTheParamsAreCorrect",
			projectName:        "Project name example",
			mockFile:           "./mocks/validate-project-name.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/validProjectName?name=Project+name+example",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "ValidateProjectKeyWhenTheProjectNameIsEmpty",
			projectName:        "",
			mockFile:           "./mocks/validate-project-name.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/validProjectName?name=Project+name+example",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheRequestMethodIsIncorrect",
			projectName:        "Project name example",
			mockFile:           "./mocks/validate-project-name.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/projectvalidate/validProjectName?name=Project+name+example",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheStatusCodeIsIncorrect",
			projectName:        "Project name example",
			mockFile:           "./mocks/validate-project-name.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/validProjectName?name=Project+name+example",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheContextIsNil",
			projectName:        "Project name example",
			mockFile:           "./mocks/validate-project-name.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/validProjectName?name=Project+name+example",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheEndpointIsIncorrect",
			projectName:        "Project name example",
			mockFile:           "./mocks/validate-project-name.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectvalidate/validProjectName?name=Project+name+example",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheResponseBodyHasADifferentFormat",
			projectName:        "Project name example",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/validProjectName?name=Project+name+example",
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

			i := &ProjectValidationService{client: mockClient}

			gotResult, gotResponse, err := i.Name(testCase.context, testCase.projectName)

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

				t.Log("-------------------------------------")
				t.Logf("Project name: %v", gotResult)
				t.Log("-------------------------------------")

			}
		})

	}

}

func TestProjectValidationService_Validate(t *testing.T) {

	testCases := []struct {
		name               string
		projectKey         string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "ValidateProjectKeyWhenTheParamsAreCorrect",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/validate-project-error-message.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/key?key=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "ValidateProjectKeyWhenTheProjectKeyIsEmpty",
			projectKey:         "",
			mockFile:           "./mocks/validate-project-error-message.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/key?key=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheRequestMethodIsIncorrect",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/validate-project-error-message.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/projectvalidate/key?key=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheStatusCodeIsIncorrect",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/validate-project-error-message.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/key?key=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheContextIsNil",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/validate-project-error-message.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/key?key=DUMMY",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheEndpointIsIncorrect",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/validate-project-error-message.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectvalidate/key?key=DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "ValidateProjectKeyWhenTheResponseBodyHasADifferentFormat",
			projectKey:         "DUMMY",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projectvalidate/key?key=DUMMY",
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

			i := &ProjectValidationService{client: mockClient}

			gotResult, gotResponse, err := i.Validate(testCase.context, testCase.projectKey)

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

				t.Log("-------------------------------------")
				t.Logf("Project key: %v", gotResult)
				t.Log("-------------------------------------")

			}
		})

	}

}
