package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestProjectComponentService_Count(t *testing.T) {

	testCases := []struct {
		name               string
		componentID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectComponentCountWhenTheParametersAreCorrect",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001/relatedIssueCounts",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectComponentCountWhenTheComponentIDIsEmpty",
			componentID:        "",
			mockFile:           "./mocks/get-project-component-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001/relatedIssueCounts",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentCountWhenTheRequestMethodIsIncorrect",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component-count.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/component/1001/relatedIssueCounts",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentCountWhenTheStatusCodeIsIncorrect",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001/relatedIssueCounts",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentCountWhenTheContextIsNil",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001/relatedIssueCounts",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentCountWhenTheEndpointIsIncorrect",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/component/1001/relatedIssueCounts",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentCountWhenTheContextIsNil",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001/relatedIssueCounts",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentCountWhenTheResponseBodyHasADifferentFormat",
			componentID:        "1001",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001/relatedIssueCounts",
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

			i := &ProjectComponentService{client: mockClient}

			gotResult, gotResponse, err := i.Count(testCase.context, testCase.componentID)

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

func TestProjectComponentService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *ProjectComponentPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateProjectComponentWhenTheParametersAreCorrect",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "HSP",
				AssigneeType:        "PROJECT_LEAD",
				LeadAccountID:       "5b10a2844c20165700ede21g",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/component",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateProjectComponentWhenThePayloadIsNil",
			payload:            nil,
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/component",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectComponentWhenTheRequestMethodIsIncorrect",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "HSP",
				AssigneeType:        "PROJECT_LEAD",
				LeadAccountID:       "5b10a2844c20165700ede21g",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/component",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectComponentWhenTheStatusCodeIsIncorrect",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "HSP",
				AssigneeType:        "PROJECT_LEAD",
				LeadAccountID:       "5b10a2844c20165700ede21g",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/component",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectComponentWhenTheContextIsNil",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "HSP",
				AssigneeType:        "PROJECT_LEAD",
				LeadAccountID:       "5b10a2844c20165700ede21g",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/component",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectComponentWhenTheEndpointIsIncorrect",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "HSP",
				AssigneeType:        "PROJECT_LEAD",
				LeadAccountID:       "5b10a2844c20165700ede21g",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/component",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectComponentWhenTheResponseBodyHasADifferentFormat",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "HSP",
				AssigneeType:        "PROJECT_LEAD",
				LeadAccountID:       "5b10a2844c20165700ede21g",
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/component",
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

			i := &ProjectComponentService{client: mockClient}

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
			}
		})

	}

}

func TestProjectComponentService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		componentID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteProjectComponentWhenTheParametersAreCorrect",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/component/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteProjectComponentWhenTheComponentIDIsEmpty",
			componentID:        "",
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/component/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectComponentWhenTheRequestMethodIsIncorrect",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/component/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectComponentWhenTheStatusCodeIsIncorrect",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/component/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectComponentWhenTheContextIsNil",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/component/1001",
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

			i := &ProjectComponentService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.componentID)

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

func TestProjectComponentService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		componentID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectComponentWhenTheParametersAreCorrect",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectComponentWhenTheComponentIDIsEmpty",
			componentID:        "",
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentWhenTheRequestMethodIsIncorrect",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/component/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentWhenTheStatusCodeIsIncorrect",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentWhenTheContextIsNil",
			componentID:        "1001",
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentWhenTheResponseBodyHasADifferentFormat",
			componentID:        "1001",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/component/1001",
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

			i := &ProjectComponentService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.componentID)

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

func TestProjectComponentService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectComponentsWhenTheParametersAreCorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-components.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/components",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectComponentsWhenTheProjectKeyOrIDIsEmpty",
			projectKeyOrID:     "",
			mockFile:           "./mocks/get-project-components.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/components",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentsWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-components.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/components",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentsWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-components.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/components",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentsWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-components.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/components",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectComponentsWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/components",
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

			i := &ProjectComponentService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.projectKeyOrID)

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

func TestProjectComponentService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		componentID        string
		payload            *ProjectComponentPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:        "UpdateProjectComponentWhenTheParametersAreCorrect",
			componentID: "1000",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "PROJECT_LEAD",
				AssigneeType:        "",
				LeadAccountID:       "",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/component/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:        "UpdateProjectComponentWhenTheComponentIDIsEmpty",
			componentID: "",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "PROJECT_LEAD",
				AssigneeType:        "",
				LeadAccountID:       "",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/component/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateProjectComponentWhenThePayloadIsNil",
			componentID:        "1000",
			payload:            nil,
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/component/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "UpdateProjectComponentWhenTheRequestMethodIsIncorrect",
			componentID: "1000",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "PROJECT_LEAD",
				AssigneeType:        "",
				LeadAccountID:       "",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/component/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "UpdateProjectComponentWhenTheStatusCodeIsIncorrect",
			componentID: "1000",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "PROJECT_LEAD",
				AssigneeType:        "",
				LeadAccountID:       "",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/component/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:        "UpdateProjectComponentWhenTheContextIsNil",
			componentID: "1000",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "PROJECT_LEAD",
				AssigneeType:        "",
				LeadAccountID:       "",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/component/1000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "UpdateProjectComponentWhenTheEndpointIsIncorrect",
			componentID: "1000",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "PROJECT_LEAD",
				AssigneeType:        "",
				LeadAccountID:       "",
			},
			mockFile:           "./mocks/get-project-component.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/component/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "UpdateProjectComponentWhenTheResponseBodyHasADifferentFormat",
			componentID: "1000",
			payload: &ProjectComponentPayloadScheme{
				IsAssigneeTypeValid: true,
				Name:                "Component 1",
				Description:         "This is a Jira component",
				Project:             "PROJECT_LEAD",
				AssigneeType:        "",
				LeadAccountID:       "",
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/component/1000",
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

			i := &ProjectComponentService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.componentID, testCase.payload)

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
