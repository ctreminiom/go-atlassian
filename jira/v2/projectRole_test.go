package v2

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestProjectRoleService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.ProjectRolePayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "GetProjectRoleWhenTheParamsAreCorrect",
			payload: &models.ProjectRolePayloadScheme{
				Name:        "ISOs",
				Description: "lorem",
			},
			mockFile:           "../v3/mocks/get-project-role.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "GetProjectRoleWhenThePayloadIsNotProvided",
			payload:            nil,
			mockFile:           "../v3/mocks/get-project-role.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "GetProjectRoleWhenTheRequestMethodIsIncorrect",
			payload: &models.ProjectRolePayloadScheme{
				Name:        "ISOs",
				Description: "lorem",
			},
			mockFile:           "../v3/mocks/get-project-role.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "GetProjectRoleWhenTheStatusCodeIsIncorrect",
			payload: &models.ProjectRolePayloadScheme{
				Name:        "ISOs",
				Description: "lorem",
			},
			mockFile:           "../v3/mocks/get-project-role.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "GetProjectRoleWhenTheContextIsNil",
			payload: &models.ProjectRolePayloadScheme{
				Name:        "ISOs",
				Description: "lorem",
			},
			mockFile:           "../v3/mocks/get-project-role.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/role",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "GetProjectRoleWhenTheResponseBodyHasADifferentFormat",
			payload: &models.ProjectRolePayloadScheme{
				Name:        "ISOs",
				Description: "lorem",
			},
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/role",
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

			i := &ProjectRoleService{client: mockClient}

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

				t.Logf("Role Name: %v", gotResult.Name)
				t.Logf("Role ID: %v", gotResult.ID)
				t.Logf("Role Self: %v", gotResult.Self)
				t.Logf("Role Description: %v", gotResult.Description)
				t.Logf("Role Actors: %v", len(gotResult.Actors))

			}
		})

	}

}

func TestProjectRoleService_Details(t *testing.T) {
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
			name:               "GetProjectRoleWhenTheParamsAreCorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/get-project-role-details.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/roledetails",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectRoleWhenTheProjectKeyOrIDIsEmpty",
			projectKeyOrID:     "",
			mockFile:           "../v3/mocks/get-project-role-details.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/roledetails",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRoleWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/get-project-role-details.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/DUMMY/roledetails",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRoleWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/get-project-role-details.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/roledetails",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectRoleWhenTheContextIsNIl",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/get-project-role-details.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/roledetails",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRoleWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/roledetails",
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

			i := &ProjectRoleService{client: mockClient}

			gotResult, gotResponse, err := i.Details(testCase.context, testCase.projectKeyOrID)

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

				for _, role := range gotResult {

					t.Log("------------------------------------------")
					t.Logf("Role Name: %v", role.Name)
					t.Logf("Role ID: %v", role.ID)
					t.Logf("Role Self: %v", role.Self)
					t.Logf("Role Description: %v", role.Description)
					t.Log("------------------------------------------")

				}

			}
		})

	}

}

func TestProjectRoleService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		roleID             int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectRoleWhenTheParamsAreCorrect",
			projectKeyOrID:     "DUMMY",
			roleID:             1000,
			mockFile:           "../v3/mocks/get-project-role.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectRoleWhenTheProjectKeyOrIDIsEmpty",
			projectKeyOrID:     "",
			roleID:             1000,
			mockFile:           "../v3/mocks/get-project-role.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRoleWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			roleID:             1000,
			mockFile:           "../v3/mocks/get-project-role.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/DUMMY/role/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRoleWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			roleID:             1000,
			mockFile:           "../v3/mocks/get-project-role.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectRoleWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			roleID:             1000,
			mockFile:           "../v3/mocks/get-project-role.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role/1000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetProjectRoleWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			roleID:             1000,
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role/1000",
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

			i := &ProjectRoleService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.projectKeyOrID, testCase.roleID)

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

				t.Logf("Role Name: %v", gotResult.Name)
				t.Logf("Role ID: %v", gotResult.ID)
				t.Logf("Role Self: %v", gotResult.Self)
				t.Logf("Role Description: %v", gotResult.Description)
				t.Logf("Role Actors: %v", len(gotResult.Actors))

			}
		})

	}

}

func TestProjectRoleService_Gets(t *testing.T) {

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
			name:               "GetProjectRolesWhenTheParamsAreCorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/get-project-roles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectRolesWhenTheRoleIDsAreStrings",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/get-project-roles-when-the-ids-are-string.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRolesWhenTheRoleURLsAreIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/get-project-roles-when-the-urls-are-incorrect.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRolesWhenTheProjectKeyOrIDIsEmpty",
			projectKeyOrID:     "",
			mockFile:           "../v3/mocks/get-project-roles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRolesWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/get-project-roles.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/DUMMY/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRolesWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/get-project-roles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetProjectRolesWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/get-project-roles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRolesWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/role",
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

			i := &ProjectRoleService{client: mockClient}

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

				for name, key := range *gotResult {
					t.Logf("Project Role Name: %v | Project Role ID: %v", name, key)
				}

			}
		})

	}

}

func TestProjectRoleService_Global(t *testing.T) {

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
			name:               "GetGlobalProjectRolesWhenTheParamsAreCorrect",
			mockFile:           "../v3/mocks/get-project-global-roles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetGlobalProjectRolesWhenTheContextIsNil",
			mockFile:           "../v3/mocks/get-project-global-roles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/role",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetGlobalProjectRolesWhenTheRequestMethodIsIncorrect",
			mockFile:           "../v3/mocks/get-project-global-roles.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetGlobalProjectRolesWhenTheStatusCodeIsIncorrect",
			mockFile:           "../v3/mocks/get-project-global-roles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/role",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetGlobalProjectRolesWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/role",
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

			i := &ProjectRoleService{client: mockClient}

			gotResult, gotResponse, err := i.Global(testCase.context)

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

				for _, role := range gotResult {

					t.Log("------------------------------------------")
					t.Logf("Role Name: %v", role.Name)
					t.Logf("Role ID: %v", role.ID)
					t.Logf("Role Self: %v", role.Self)
					t.Logf("Role Description: %v", role.Description)
					t.Log("------------------------------------------")

				}

			}
		})

	}

}
