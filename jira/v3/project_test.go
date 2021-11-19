package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestProjectService_Archive(t *testing.T) {

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
			name:               "ArchiveProjectWhenTheProjectKeyOrIDIsCorrect",
			projectKeyOrID:     "DUMMY",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/archive",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "ArchiveProjectWhenTheProjectKeyOrIDIsNotProvided",
			projectKeyOrID:     "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/archive",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "ArchiveProjectWhenTheProjectKeyOrIDIsIncorrect",
			projectKeyOrID:     "DUMMYS",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/archive",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "ArchiveProjectWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/archive",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "ArchiveProjectWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/archive",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "ArchiveProjectWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/archive",
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

			i := &ProjectService{client: mockClient}

			gotResponse, err := i.Archive(testCase.context, testCase.projectKeyOrID)

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

func TestProjectService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.ProjectPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "ArchiveProjectWhenThePayloadIsCorrect",
			payload: &models.ProjectPayloadScheme{
				NotificationScheme:  10021,
				Description:         "Example Project description",
				LeadAccountID:       "396c4bf1-361b-4754-ae47-91fe6aabbc40",
				URL:                 "https://www.example.com",
				ProjectTemplateKey:  "com.atlassian.jira-core-project-templates:jira-core-simplified-process-control",
				AvatarID:            10200,
				IssueSecurityScheme: 10001,
				Name:                "Project Example",
				PermissionScheme:    10011,
				AssigneeType:        "PROJECT_LEAD",
				ProjectTypeKey:      "business",
				Key:                 "DUMMY",
				CategoryID:          10120,
			},
			mockFile:           "./mocks/create-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "ArchiveProjectWhenThePayloadIsNil",
			payload:            nil,
			mockFile:           "./mocks/create-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "ArchiveProjectWhenTheEndpointIsIncorrect",
			payload: &models.ProjectPayloadScheme{
				NotificationScheme:  10021,
				Description:         "Example Project description",
				LeadAccountID:       "396c4bf1-361b-4754-ae47-91fe6aabbc40",
				URL:                 "https://www.example.com",
				ProjectTemplateKey:  "com.atlassian.jira-core-project-templates:jira-core-simplified-process-control",
				AvatarID:            10200,
				IssueSecurityScheme: 10001,
				Name:                "Project Example",
				PermissionScheme:    10011,
				AssigneeType:        "PROJECT_LEAD",
				ProjectTypeKey:      "business",
				Key:                 "DUMMY",
				CategoryID:          10120,
			},
			mockFile:           "./mocks/create-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "ArchiveProjectWhenTheRequestMethodIsIncorrect",
			payload: &models.ProjectPayloadScheme{
				NotificationScheme:  10021,
				Description:         "Example Project description",
				LeadAccountID:       "396c4bf1-361b-4754-ae47-91fe6aabbc40",
				URL:                 "https://www.example.com",
				ProjectTemplateKey:  "com.atlassian.jira-core-project-templates:jira-core-simplified-process-control",
				AvatarID:            10200,
				IssueSecurityScheme: 10001,
				Name:                "Project Example",
				PermissionScheme:    10011,
				AssigneeType:        "PROJECT_LEAD",
				ProjectTypeKey:      "business",
				Key:                 "DUMMY",
				CategoryID:          10120,
			},
			mockFile:           "./mocks/create-project.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "ArchiveProjectWhenTheStatusCodeIsIncorrect",
			payload: &models.ProjectPayloadScheme{
				NotificationScheme:  10021,
				Description:         "Example Project description",
				LeadAccountID:       "396c4bf1-361b-4754-ae47-91fe6aabbc40",
				URL:                 "https://www.example.com",
				ProjectTemplateKey:  "com.atlassian.jira-core-project-templates:jira-core-simplified-process-control",
				AvatarID:            10200,
				IssueSecurityScheme: 10001,
				Name:                "Project Example",
				PermissionScheme:    10011,
				AssigneeType:        "PROJECT_LEAD",
				ProjectTypeKey:      "business",
				Key:                 "DUMMY",
				CategoryID:          10120,
			},
			mockFile:           "./mocks/create-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "ArchiveProjectWhenTheContextIsNil",
			payload: &models.ProjectPayloadScheme{
				NotificationScheme:  10021,
				Description:         "Example Project description",
				LeadAccountID:       "396c4bf1-361b-4754-ae47-91fe6aabbc40",
				URL:                 "https://www.example.com",
				ProjectTemplateKey:  "com.atlassian.jira-core-project-templates:jira-core-simplified-process-control",
				AvatarID:            10200,
				IssueSecurityScheme: 10001,
				Name:                "Project Example",
				PermissionScheme:    10011,
				AssigneeType:        "PROJECT_LEAD",
				ProjectTypeKey:      "business",
				Key:                 "DUMMY",
				CategoryID:          10120,
			},
			mockFile:           "./mocks/create-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "ArchiveProjectWhenTheResponseBodyHasADifferentFormat",
			payload: &models.ProjectPayloadScheme{
				NotificationScheme:  10021,
				Description:         "Example Project description",
				LeadAccountID:       "396c4bf1-361b-4754-ae47-91fe6aabbc40",
				URL:                 "https://www.example.com",
				ProjectTemplateKey:  "com.atlassian.jira-core-project-templates:jira-core-simplified-process-control",
				AvatarID:            10200,
				IssueSecurityScheme: 10001,
				Name:                "Project Example",
				PermissionScheme:    10011,
				AssigneeType:        "PROJECT_LEAD",
				ProjectTypeKey:      "business",
				Key:                 "DUMMY",
				CategoryID:          10120,
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project",
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

			i := &ProjectService{client: mockClient}

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

func TestProjectService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		enableUndo         bool
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteProjectWhenTheProjectKeyOrIDIsCorrect",
			projectKeyOrID:     "DUMMY",
			enableUndo:         true,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY?enableUndo=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteProjectWhenTheProjectKeyOrIDIsNotProvided",
			projectKeyOrID:     "",
			enableUndo:         true,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY?enableUndo=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectWhenTheEnableUndoIsFalse",
			projectKeyOrID:     "DUMMY",
			enableUndo:         false,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteProjectWhenTheProjectKeyOrIDIsIncorrect",
			projectKeyOrID:     "DUMMY",
			enableUndo:         true,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/project/DUMMY?enableUndo=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			enableUndo:         true,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY?enableUndo=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			enableUndo:         true,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY?enableUndo=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			enableUndo:         true,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY?enableUndo=true",
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

			i := &ProjectService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.projectKeyOrID, testCase.enableUndo)

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

func TestProjectService_DeleteAsynchronously(t *testing.T) {

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
			name:               "DeleteProjectAsynchronouslyWhenTheProjectKeyOrIDIsCorrect",
			projectKeyOrID:     "DUMMY",
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/delete-project-asynchronously.json",
			endpoint:           "/rest/api/3/project/DUMMY/delete",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "DeleteProjectAsynchronouslyWhenTheProjectKeyOrIDIsNotProvided",
			projectKeyOrID:     "",
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/delete-project-asynchronously.json",
			endpoint:           "/rest/api/3/project/DUMMY/delete",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectAsynchronouslyWhenTheProjectKeyOrIDIsIncorrect",
			projectKeyOrID:     "DUMMY",
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/delete-project-asynchronously.json",
			endpoint:           "/rest/api/2/project/DUMMY/delete",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectAsynchronouslyWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "./mocks/delete-project-asynchronously.json",
			endpoint:           "/rest/api/3/project/DUMMY/delete",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectAsynchronouslyWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/delete-project-asynchronously.json",
			endpoint:           "/rest/api/3/project/DUMMY/delete",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectAsynchronouslyWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/delete-project-asynchronously.json",
			endpoint:           "/rest/api/3/project/DUMMY/delete",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "DeleteProjectAsynchronouslyWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/empty_json.json",
			endpoint:           "/rest/api/3/project/DUMMY/delete",
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

			i := &ProjectService{client: mockClient}

			gotResult, gotResponse, err := i.DeleteAsynchronously(testCase.context, testCase.projectKeyOrID)

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

func TestProjectService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		expands            []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectWhenTheProjectKeyOrIDIsCorrect",
			projectKeyOrID:     "DUMMY",
			expands:            []string{"issueTypes", "lead", "description"},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY?expand=issueTypes%2Clead%2Cdescription",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectWhenTheProjectKeyOrIDIsEmpty",
			projectKeyOrID:     "",
			expands:            []string{"issueTypes", "lead", "description"},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY?expand=issueTypes%2Clead%2Cdescription",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectWhenTheProjectExpandIsNil",
			projectKeyOrID:     "DUMMY",
			expands:            nil,
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectWhenTheProjectEndpointIsIncorrect",
			projectKeyOrID:     "DUMMY",
			expands:            []string{"issueTypes", "lead", "description"},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY?expand=issueTypes%2Clead%2Cdescription",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			expands:            []string{"issueTypes", "lead", "description"},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY?expand=issueTypes%2Clead%2Cdescription",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			expands:            []string{"issueTypes", "lead", "description"},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY?expand=issueTypes%2Clead%2Cdescription",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			expands:            []string{"issueTypes", "lead", "description"},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY?expand=issueTypes%2Clead%2Cdescription",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			expands:            []string{"issueTypes", "lead", "description"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY?expand=issueTypes%2Clead%2Cdescription",
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

			i := &ProjectService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.projectKeyOrID, testCase.expands)

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

func TestProjectService_Hierarchy(t *testing.T) {

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
			name:               "GetProjectIssueTypeHierarchyWhenTheProjectKeyOrIDIsCorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-issue-type-hierarchy.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/hierarchy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectIssueTypeHierarchyWhenTheProjectKeyOrIDIsNotProvided",
			projectKeyOrID:     "",
			mockFile:           "./mocks/get-project-issue-type-hierarchy.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/hierarchy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectIssueTypeHierarchyWhenTheProjectKeyOrIDIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-issue-type-hierarchy.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/hierarchy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectIssueTypeHierarchyWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-issue-type-hierarchy.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/hierarchy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectIssueTypeHierarchyWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-issue-type-hierarchy.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/hierarchy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectIssueTypeHierarchyWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-issue-type-hierarchy.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/hierarchy",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectIssueTypeHierarchyWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/hierarchy",
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

			i := &ProjectService{client: mockClient}

			gotResult, gotResponse, err := i.Hierarchy(testCase.context, testCase.projectKeyOrID)

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

func TestProjectService_NotificationScheme(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectNotificationSchemeWhenTheProjectKeyOrIDIsCorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-notification-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/notificationscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectNotificationSchemeWhenTheProjectKeyOrIDIsNotProvided",
			projectKeyOrID:     "",
			mockFile:           "./mocks/get-project-notification-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/notificationscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectNotificationSchemeWhenTheExpandsAreSet",
			projectKeyOrID:     "DUMMY",
			expand:             []string{"field", "group"},
			mockFile:           "./mocks/get-project-notification-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/notificationscheme?expand=field%2Cgroup",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectNotificationSchemeWhenTheProjectKeyOrIDIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-notification-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/notificationscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectNotificationSchemeWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-notification-scheme.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/notificationscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectNotificationSchemeWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-notification-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/notificationscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectNotificationSchemeWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-notification-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/notificationscheme",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectNotificationSchemeWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/notificationscheme",
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

			i := &ProjectService{client: mockClient}

			gotResult, gotResponse, err := i.NotificationScheme(testCase.context, testCase.projectKeyOrID, testCase.expand)

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

func TestProjectService_Restore(t *testing.T) {

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
			name:               "RestoreDeletedProjectWhenTheProjectKeyOrIDIsCorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/restore",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "RestoreDeletedProjectWhenTheProjectKeyOrIDIsNotProvided",
			projectKeyOrID:     "",
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/restore",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RestoreDeletedProjectWhenTheProjectKeyOrIDIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/project/DUMMY/restore",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RestoreDeletedProjectWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/restore",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RestoreDeletedProjectWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/restore",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "RestoreDeletedProjectWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/restore",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RestoreDeletedProjectWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/restore",
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

			i := &ProjectService{client: mockClient}

			gotResult, gotResponse, err := i.Restore(testCase.context, testCase.projectKeyOrID)

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

func TestProjectService_Search(t *testing.T) {

	testCases := []struct {
		name               string
		opts               *models.ProjectSearchOptionsScheme
		startAt            int
		maxResults         int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "SearchProjectsWhenTheOptionsAreCorrect",
			opts: &models.ProjectSearchOptionsScheme{
				OrderBy:        "category",
				Query:          "key",
				Action:         "view",
				ProjectKeyType: "business",
				CategoryID:     1000,
				Expand:         []string{"description", "insight"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/search?action=view&categoryId=1000&expand=description%2Cinsight&maxResults=50&orderBy=category&query=key&startAt=0&typeKey=business",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "SearchProjectsWhenTheOptionsIsNil",
			opts:               nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/search?action=view&categoryId=1000&expand=description%2Cinsight&maxResults=50&orderBy=category&query=key&startAt=0&typeKey=business",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "SearchProjectsWhenTheEndpointIsIncorrect",
			opts: &models.ProjectSearchOptionsScheme{
				OrderBy:        "category",
				Query:          "key",
				Action:         "view",
				ProjectKeyType: "business",
				CategoryID:     1000,
				Expand:         []string{"description", "insight"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/projects/search?action=view&categoryId=1000&expand=description%2Cinsight&maxResults=50&orderBy=category&query=key&startAt=0&typeKey=business",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "SearchProjectsWhenTheRequestMethodIsIncorrect",
			opts: &models.ProjectSearchOptionsScheme{
				OrderBy:        "category",
				Query:          "key",
				Action:         "view",
				ProjectKeyType: "business",
				CategoryID:     1000,
				Expand:         []string{"description", "insight"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-projects.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/search?action=view&categoryId=1000&expand=description%2Cinsight&maxResults=50&orderBy=category&query=key&startAt=0&typeKey=business",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "SearchProjectsWhenTheStatusCodeIsIncorrect",
			opts: &models.ProjectSearchOptionsScheme{
				OrderBy:        "category",
				Query:          "key",
				Action:         "view",
				ProjectKeyType: "business",
				CategoryID:     1000,
				Expand:         []string{"description", "insight"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/search?action=view&categoryId=1000&expand=description%2Cinsight&maxResults=50&orderBy=category&query=key&startAt=0&typeKey=business",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "SearchProjectsWhenTheContextIsNil",
			opts: &models.ProjectSearchOptionsScheme{
				OrderBy:        "category",
				Query:          "key",
				Action:         "view",
				ProjectKeyType: "business",
				CategoryID:     1000,
				Expand:         []string{"description", "insight"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/search?action=view&categoryId=1000&expand=description%2Cinsight&maxResults=50&orderBy=category&query=key&startAt=0&typeKey=business",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "SearchProjectsWhenTheTheResponseBodyHasADifferentFormat",
			opts: &models.ProjectSearchOptionsScheme{
				OrderBy:        "category",
				Query:          "key",
				Action:         "view",
				ProjectKeyType: "business",
				CategoryID:     1000,
				Expand:         []string{"description", "insight"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/search?action=view&categoryId=1000&expand=description%2Cinsight&maxResults=50&orderBy=category&query=key&startAt=0&typeKey=business",
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

			i := &ProjectService{client: mockClient}

			gotResult, gotResponse, err := i.Search(testCase.context, testCase.opts, testCase.startAt, testCase.maxResults)

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

func TestProjectService_Statuses(t *testing.T) {

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
			name:               "GetProjectStatusesWhenTheProjectKeyIsCorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-statuses.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/statuses",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectStatusesWhenTheProjectKeyIsNotProvided",
			projectKeyOrID:     "",
			mockFile:           "./mocks/get-project-statuses.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/statuses",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectStatusesWhenTheProjectKeyIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-statuses.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/statuses",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectStatusesWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-statuses.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/statuses",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectStatusesWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-statuses.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/statuses",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectStatusesWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/get-project-statuses.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/statuses",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectStatusesWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/project/DUMMY/statuses",
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

			i := &ProjectService{client: mockClient}

			gotResult, gotResponse, err := i.Statuses(testCase.context, testCase.projectKeyOrID)

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

func TestProjectService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		payload            *models.ProjectUpdateScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:           "UpdateProjectWhenThePayloadIsCorrect",
			projectKeyOrID: "DUMMY",
			payload: &models.ProjectUpdateScheme{
				NotificationScheme: 10000,
				Name:               "New project",
				PermissionScheme:   10001,
			},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:           "UpdateProjectWhenTheProjectKeyOrIDIsNotProvided",
			projectKeyOrID: "",
			payload: &models.ProjectUpdateScheme{
				NotificationScheme: 10000,
				Name:               "New project",
				PermissionScheme:   10001,
			},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateProjectWhenThePayloadIsNil",
			projectKeyOrID:     "DUMMY",
			payload:            nil,
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:           "UpdateProjectWhenTheEndpointIsIncorrect",
			projectKeyOrID: "DUMMY",
			payload: &models.ProjectUpdateScheme{
				NotificationScheme: 10000,
				Name:               "New project",
				PermissionScheme:   10001,
			},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/project/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:           "UpdateProjectWhenTheRequestMethodIsIncorrect",
			projectKeyOrID: "DUMMY",
			payload: &models.ProjectUpdateScheme{
				NotificationScheme: 10000,
				Name:               "New project",
				PermissionScheme:   10001,
			},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:           "UpdateProjectWhenTheStatusCodeIsIncorrect",
			projectKeyOrID: "DUMMY",
			payload: &models.ProjectUpdateScheme{
				NotificationScheme: 10000,
				Name:               "New project",
				PermissionScheme:   10001,
			},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:           "UpdateProjectWhenTheContextIsNil",
			projectKeyOrID: "DUMMY",
			payload: &models.ProjectUpdateScheme{
				NotificationScheme: 10000,
				Name:               "New project",
				PermissionScheme:   10001,
			},
			mockFile:           "./mocks/get-project.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:           "UpdateProjectWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID: "DUMMY",
			payload: &models.ProjectUpdateScheme{
				NotificationScheme: 10000,
				Name:               "New project",
				PermissionScheme:   10001,
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/project/DUMMY",
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

			i := &ProjectService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.projectKeyOrID, testCase.payload)

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
