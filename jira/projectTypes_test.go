package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestProjectTypeService_Accessible(t *testing.T) {

	testCases := []struct {
		name               string
		projectTypeKey     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetAccessibleProjectTypeWhenTheParamsAreCorrect",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/DUMMY/accessible",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetAccessibleProjectTypeWhenTheProjectTypeKeyIsEmpty",
			projectTypeKey:     "",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/DUMMY/accessible",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetAccessibleProjectTypeWhenTheRequestMethodIsIncorrect",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/type/DUMMY/accessible",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetAccessibleProjectTypeWhenTheStatusCodeIsIncorrect",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/DUMMY/accessible",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetAccessibleProjectTypeWhenTheContextIsNil",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/DUMMY/accessible",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetAccessibleProjectTypeWhenTheEndpointsIsIncorrect",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/type/DUMMY/accessible",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetAccessibleProjectTypeWhenTheResponseBodyHasADifferentFormat",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/DUMMY/accessible",
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

			i := &ProjectTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Accessible(testCase.context, testCase.projectTypeKey)

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

				t.Log("---------------------------------")
				t.Logf("Project Type key: %v", gotResult.Key)
				t.Logf("Project Type color: %v", gotResult.Color)
				t.Logf("Project Type Description: %v", gotResult.DescriptionI18NKey)
				t.Logf("Project Type Icon: %v", gotResult.Icon)
				t.Log("---------------------------------")

			}
		})

	}

}

func TestProjectTypeService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		projectTypeKey     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectTypeWhenTheParamsAreCorrect",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectTypeWhenTheProjectTypeKeyIsEmpty",
			projectTypeKey:     "",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypeWhenTheRequestMethodIsIncorrect",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/type/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypeWhenTheStatusCodeIsIncorrect",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypeWhenTheEndpointIsIncorrect",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/type/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypeWhenTheContextIsNil",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/get-project-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/DUMMY",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypeWhenTheResponseBodyHasADifferentFormat",
			projectTypeKey:     "DUMMY",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/DUMMY",
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

			i := &ProjectTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.projectTypeKey)

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

				t.Log("---------------------------------")
				t.Logf("Project Type key: %v", gotResult.Key)
				t.Logf("Project Type color: %v", gotResult.Color)
				t.Logf("Project Type Description: %v", gotResult.DescriptionI18NKey)
				t.Logf("Project Type Icon: %v", gotResult.Icon)
				t.Log("---------------------------------")

			}
		})

	}

}

func TestProjectTypeService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectTypesWhenTheParamsAreCorrect",
			mockFile:           "./mocks/get-project-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectTypesWhenTheContextIsNil",
			mockFile:           "./mocks/get-project-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypesWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-project-types.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/type",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypesWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-project-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypesWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-project-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/type",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypesWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type",
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

			i := &ProjectTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context)

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

				for _, projectType := range *gotResult {

					t.Log("---------------------------------")
					t.Logf("Project Type key: %v", projectType.Key)
					t.Logf("Project Type color: %v", projectType.Color)
					t.Logf("Project Type Description: %v", projectType.DescriptionI18NKey)
					t.Logf("Project Type Icon: %v", projectType.Icon)
					t.Log("---------------------------------")
				}

			}
		})

	}

}

func TestProjectTypeService_Licensed(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectTypesLicensedWhenTheParamsAreCorrect",
			mockFile:           "./mocks/get-project-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/accessible",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectTypesLicensedWhenTheContextIsNil",
			mockFile:           "./mocks/get-project-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/accessible",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypesLicensedWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get-project-types.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/type/accessible",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypesLicensedWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get-project-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/accessible",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypesLicensedWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-project-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/type/accessible",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectTypesLicensedTheResponseBodyHasADifferentFormat",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/type/accessible",
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

			i := &ProjectTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Licensed(testCase.context)

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

				for _, projectType := range *gotResult {

					t.Log("---------------------------------")
					t.Logf("Project Type key: %v", projectType.Key)
					t.Logf("Project Type color: %v", projectType.Color)
					t.Logf("Project Type Description: %v", projectType.DescriptionI18NKey)
					t.Logf("Project Type Icon: %v", projectType.Icon)
					t.Log("---------------------------------")
				}

			}
		})

	}

}
