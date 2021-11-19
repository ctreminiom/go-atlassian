package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestProjectCategoryService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.ProjectCategoryPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateProjectCategoryWhenTheParametersAreCorrect",
			payload: &models.ProjectCategoryPayloadScheme{
				Name:        "CREATED",
				Description: "Created Project Category",
			},
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/projectCategory",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateProjectCategoryWhenThePayloadIsNotProvided",
			payload:            nil,
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/projectCategory",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectCategoryWhenTheRequestMethodIsIncorrect",
			payload: &models.ProjectCategoryPayloadScheme{
				Name:        "CREATED",
				Description: "Created Project Category",
			},
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/projectCategory",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectCategoryWhenTheStatusCodeIsIncorrect",
			payload: &models.ProjectCategoryPayloadScheme{
				Name:        "CREATED",
				Description: "Created Project Category",
			},
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/projectCategory",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CreateProjectCategoryWhenTheContextIsNil",
			payload: &models.ProjectCategoryPayloadScheme{
				Name:        "CREATED",
				Description: "Created Project Category",
			},
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/projectCategory",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectCategoryWhenTheResponseBodyHasADifferentFormat",
			payload: &models.ProjectCategoryPayloadScheme{
				Name:        "CREATED",
				Description: "Created Project Category",
			},
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/projectCategory",
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

			i := &ProjectCategoryService{client: mockClient}

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

func TestProjectCategoryService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		projectCategoryID  int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteProjectCategoryWhenTheParametersAreCorrect",
			projectCategoryID:  1000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteProjectCategoryWhenTheRequestMethodIsIncorrect",
			projectCategoryID:  1000,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectCategoryWhenTheStatusCodeIsIncorrect",
			projectCategoryID:  1000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectCategoryWhenTheContextIsNil",
			projectCategoryID:  1000,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/projectCategory/1000",
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

			i := &ProjectCategoryService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.projectCategoryID)

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

func TestProjectCategoryService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		projectCategoryID  int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectCategoryWhenTheParametersAreCorrect",
			projectCategoryID:  1000,
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectCategoryWhenTheRequestMethodIsIncorrect",
			projectCategoryID:  1000,
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectCategoryWhenTheStatusCodeIsIncorrect",
			projectCategoryID:  1000,
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectCategoryWhenTheContextIsNil",
			projectCategoryID:  1000,
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectCategoryWhenTheResponseBodyHasADifferentFormat",
			projectCategoryID:  1000,
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectCategory/1000",
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

			i := &ProjectCategoryService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.projectCategoryID)

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

func TestProjectCategoryService_Gets(t *testing.T) {

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
			name:               "GetProjectCategoriesWhenTheParametersAreCorrect",
			mockFile:           "../v3/mocks/get-project-categories.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectCategory",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectCategoriesWhenTheContextIsNil",
			mockFile:           "../v3/mocks/get-project-categories.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectCategory",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectCategoriesWhenTheRequestMethodIsIncorrect",
			mockFile:           "../v3/mocks/get-project-categories.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/projectCategory",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectCategoriesWhenTheStatusCodeIsIncorrect",
			mockFile:           "../v3/mocks/get-project-categories.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectCategory",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectCategoriesWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/projectCategory",
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

			i := &ProjectCategoryService{client: mockClient}

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
			}
		})

	}

}

func TestProjectCategoryService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		projectCategoryID  int
		payload            *models.ProjectCategoryPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:              "UpdateProjectCategoryWhenTheParametersAreCorrect",
			projectCategoryID: 1000,
			payload: &models.ProjectCategoryPayloadScheme{
				Name:        "UPDATED",
				Description: "Updated Project Category",
			},
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "UpdateProjectCategoryWhenThePayloadIsNotProvided",
			projectCategoryID:  1000,
			payload:            nil,
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:              "UpdateProjectCategoryWhenTheContextIsNil",
			projectCategoryID: 1000,
			payload: &models.ProjectCategoryPayloadScheme{
				Name:        "UPDATED",
				Description: "Updated Project Category",
			},
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:              "UpdateProjectCategoryWhenTheRequestMethodIsIncorrect",
			projectCategoryID: 1000,
			payload: &models.ProjectCategoryPayloadScheme{
				Name:        "UPDATED",
				Description: "Updated Project Category",
			},
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:              "UpdateProjectCategoryWhenTheStatusCodeIsIncorrect",
			projectCategoryID: 1000,
			payload: &models.ProjectCategoryPayloadScheme{
				Name:        "UPDATED",
				Description: "Updated Project Category",
			},
			mockFile:           "../v3/mocks/get-project-category.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/projectCategory/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:              "UpdateProjectCategoryWhenTheResponseBodyHasADifferentFormat",
			projectCategoryID: 1000,
			payload: &models.ProjectCategoryPayloadScheme{
				Name:        "UPDATED",
				Description: "Updated Project Category",
			},
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/projectCategory/1000",
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

			i := &ProjectCategoryService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.projectCategoryID, testCase.payload)

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
