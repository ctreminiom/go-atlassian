package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/url"
	"testing"
)

func TestWorkflowSchemeService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.WorkflowSchemePayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:     "CreateWorkflowSchemeWhenTheParametersAreCorrect",
			mockFile: "./mocks/get-worflow-scheme.json",
			payload: &models.WorkflowSchemePayloadScheme{
				DefaultWorkflow: "jira",
				Name:            "Example workflow scheme",
				Description:     "The description of the example workflow scheme.",
				IssueTypeMappings: map[string]string{
					"10000": "scrum workflow",
					"10001": "builds workflow",
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/workflowscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "CreateWorkflowSchemeWhenThePayloadIsNotSet",
			mockFile:           "./mocks/get-worflow-scheme.json",
			payload:            nil,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/workflowscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:     "CreateWorkflowSchemeWhenTheContextIsNotSet",
			mockFile: "./mocks/get-worflow-scheme.json",
			payload: &models.WorkflowSchemePayloadScheme{
				DefaultWorkflow: "jira",
				Name:            "Example workflow scheme",
				Description:     "The description of the example workflow scheme.",
				IssueTypeMappings: map[string]string{
					"10000": "scrum workflow",
					"10001": "builds workflow",
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/workflowscheme",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:     "CreateWorkflowSchemeWhenTheRequestMethodIsIncorrect",
			mockFile: "./mocks/get-worflow-scheme.json",
			payload: &models.WorkflowSchemePayloadScheme{
				DefaultWorkflow: "jira",
				Name:            "Example workflow scheme",
				Description:     "The description of the example workflow scheme.",
				IssueTypeMappings: map[string]string{
					"10000": "scrum workflow",
					"10001": "builds workflow",
				},
			},
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/3/workflowscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:     "CreateWorkflowSchemeWhenTheStatusCodeIsIncorrect",
			mockFile: "./mocks/get-worflow-scheme.json",
			payload: &models.WorkflowSchemePayloadScheme{
				DefaultWorkflow: "jira",
				Name:            "Example workflow scheme",
				Description:     "The description of the example workflow scheme.",
				IssueTypeMappings: map[string]string{
					"10000": "scrum workflow",
					"10001": "builds workflow",
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/workflowscheme",
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

			i := &WorkflowSchemeService{client: mockClient}

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

func TestWorkflowSchemeService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		workflowSchemeID   int
		isExits            bool
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetWorkflowWorkflowSchemeWhenTheParametersAreCorrect",
			workflowSchemeID:   1006,
			isExits:            true,
			mockFile:           "./mocks/get-worflow-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme/1006?returnDraftIfExists=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetWorkflowWorkflowSchemeWhenTheContextIsNotSet",
			workflowSchemeID:   1006,
			isExits:            true,
			mockFile:           "./mocks/get-worflow-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme/1006?returnDraftIfExists=true",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetWorkflowWorkflowSchemeWhenTheRequestMethodIsIncorrect",
			workflowSchemeID:   1006,
			isExits:            true,
			mockFile:           "./mocks/get-worflow-scheme.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/3/workflowscheme/1006?returnDraftIfExists=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetWorkflowWorkflowSchemeWhenTheStatusCodeIsIncorrect",
			workflowSchemeID:   1006,
			isExits:            true,
			mockFile:           "./mocks/get-worflow-scheme.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme/1006?returnDraftIfExists=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetWorkflowWorkflowSchemeWhenTheResponseBodyIsEmpty",
			workflowSchemeID:   1006,
			isExits:            true,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme/1006?returnDraftIfExists=true",
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

			i := &WorkflowSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.workflowSchemeID, testCase.isExits)

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

func TestWorkflowSchemeService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		workflowSchemeID   int
		payload            *models.WorkflowSchemePayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:             "UpdateWorkflowSchemeWhenTheParametersAreCorrect",
			workflowSchemeID: 1006,
			mockFile:         "./mocks/get-worflow-scheme.json",
			payload: &models.WorkflowSchemePayloadScheme{
				DefaultWorkflow: "jira",
				Name:            "Example workflow scheme",
				Description:     "The description of the example workflow scheme.",
				IssueTypeMappings: map[string]string{
					"10000": "scrum workflow",
					"10001": "builds workflow",
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/workflowscheme/1006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:             "UpdateWorkflowSchemeWhenTheContextIsNotSet",
			workflowSchemeID: 1006,
			mockFile:         "./mocks/get-worflow-scheme.json",
			payload: &models.WorkflowSchemePayloadScheme{
				DefaultWorkflow: "jira",
				Name:            "Example workflow scheme",
				Description:     "The description of the example workflow scheme.",
				IssueTypeMappings: map[string]string{
					"10000": "scrum workflow",
					"10001": "builds workflow",
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/workflowscheme/1006",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:             "UpdateWorkflowSchemeWhenTheResponseBodyIsEmpty",
			workflowSchemeID: 1006,
			mockFile:         "./mocks/empty_json.json",
			payload: &models.WorkflowSchemePayloadScheme{
				DefaultWorkflow: "jira",
				Name:            "Example workflow scheme",
				Description:     "The description of the example workflow scheme.",
				IssueTypeMappings: map[string]string{
					"10000": "scrum workflow",
					"10001": "builds workflow",
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/workflowscheme/1006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:             "UpdateWorkflowSchemeWhenTheRequestMethodIsIncorrect",
			workflowSchemeID: 1006,
			mockFile:         "./mocks/get-worflow-scheme.json",
			payload: &models.WorkflowSchemePayloadScheme{
				DefaultWorkflow: "jira",
				Name:            "Example workflow scheme",
				Description:     "The description of the example workflow scheme.",
				IssueTypeMappings: map[string]string{
					"10000": "scrum workflow",
					"10001": "builds workflow",
				},
			},
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/3/workflowscheme/1006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:             "UpdateWorkflowSchemeWhenTheStatusCodeIsIncorrect",
			workflowSchemeID: 1006,
			mockFile:         "./mocks/get-worflow-scheme.json",
			payload: &models.WorkflowSchemePayloadScheme{
				DefaultWorkflow: "jira",
				Name:            "Example workflow scheme",
				Description:     "The description of the example workflow scheme.",
				IssueTypeMappings: map[string]string{
					"10000": "scrum workflow",
					"10001": "builds workflow",
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/workflowscheme/1006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "UpdateWorkflowSchemeWhenThePayloadIsNotSet",
			workflowSchemeID:   1006,
			mockFile:           "./mocks/get-worflow-scheme.json",
			payload:            nil,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/workflowscheme/1006",
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

			i := &WorkflowSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.workflowSchemeID, testCase.payload)

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

func TestWorkflowSchemeService_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetWorkflowSchemesWhenTheParametersAreCorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-workflow-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetWorkflowSchemesWhenTheRequestMethodIsIncorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-workflow-schemes.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/3/workflowscheme?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetWorkflowSchemesWhenTheStatusCodeIsIncorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-workflow-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetWorkflowSchemesWhenTheContextIsNotSet",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-workflow-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme?maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetWorkflowSchemesWhenTheResponseBodyIsEmpty",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme?maxResults=50&startAt=0",
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

			i := &WorkflowSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.startAt, testCase.maxResults)

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

func TestWorkflowSchemeService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		workflowSchemeID   int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteWorkflowSchemeWhenTheParametersAreCorrect",
			workflowSchemeID:   1006,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/workflowscheme/1006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteWorkflowSchemeWhenTheRequestMethodIsIncorrect",
			workflowSchemeID:   1006,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme/1006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteWorkflowSchemeWhenTheStatusCodeIsIncorrect",
			workflowSchemeID:   1006,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/workflowscheme/1006",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteWorkflowSchemeWhenTheContextIsNotSet",
			workflowSchemeID:   1006,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/workflowscheme/1006",
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

			i := &WorkflowSchemeService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.workflowSchemeID)

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

func TestWorkflowSchemeService_Associations(t *testing.T) {

	testCases := []struct {
		name               string
		projectIDs         []int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetWorkflowAssociationsWhenTheParametersAreCorrect",
			projectIDs:         []int{10001, 10002, 10003},
			mockFile:           "../mocks/get-worflow-scheme-associations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme/project?projectId=10001&projectId=10002&projectId=10003",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetWorkflowAssociationsWhenTheProjectIDsAreNotSet",
			projectIDs:         nil,
			mockFile:           "../mocks/get-worflow-scheme-associations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme/project?projectId=10001&projectId=10002&projectId=10003",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetWorkflowAssociationsWhenTheContextIsNotProvided",
			projectIDs:         []int{10001, 10002, 10003},
			mockFile:           "../mocks/get-worflow-scheme-associations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme/project?projectId=10001&projectId=10002&projectId=10003",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetWorkflowAssociationsWhenTheRequestMethodIsIncorrect",
			projectIDs:         []int{10001, 10002, 10003},
			mockFile:           "../mocks/get-worflow-scheme-associations.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/workflowscheme/project?projectId=10001&projectId=10002&projectId=10003",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetWorkflowAssociationsWhenTheStatusCodeIsIncorrect",
			projectIDs:         []int{10001, 10002, 10003},
			mockFile:           "../mocks/get-worflow-scheme-associations.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme/project?projectId=10001&projectId=10002&projectId=10003",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetWorkflowAssociationsWhenTheResponseBodyIsEmpty",
			projectIDs:         []int{10001, 10002, 10003},
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflowscheme/project?projectId=10001&projectId=10002&projectId=10003",
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

			i := &WorkflowSchemeService{client: mockClient}

			getResult, gotResponse, err := i.Associations(testCase.context, testCase.projectIDs)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, getResult, nil)

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

				for _, value := range getResult.Values {
					log.Println(value.ProjectIds, value.WorkflowScheme.Name)
				}
			}

		})
	}
}

func TestWorkflowSchemeService_Assign(t *testing.T) {

	testCases := []struct {
		name               string
		workflowSchemeID   string
		projectID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AssignWorkflowSchemeToProjectWhenTheParametersAreCorrect",
			workflowSchemeID:   "10001",
			projectID:          "10002",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/workflowscheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "AssignWorkflowSchemeToProjectWhenTheWorkflowSchemeIDIsNotSet",
			workflowSchemeID:   "",
			projectID:          "10002",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/workflowscheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignWorkflowSchemeToProjectWhenTheProjectIDIsNotSet",
			workflowSchemeID:   "10001",
			projectID:          "",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/workflowscheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignWorkflowSchemeToProjectWhenTheRequestMethodIsIncorrect",
			workflowSchemeID:   "10001",
			projectID:          "10002",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/3/workflowscheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "AssignWorkflowSchemeToProjectWhenTheStatusCodeIsIncorrect",
			workflowSchemeID:   "10001",
			projectID:          "10002",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/workflowscheme/project",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AssignWorkflowSchemeToProjectWhenTheContextIsNotProvided",
			workflowSchemeID:   "10001",
			projectID:          "10002",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/workflowscheme/project",
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

			i := &WorkflowSchemeService{client: mockClient}

			gotResponse, err := i.Assign(testCase.context, testCase.workflowSchemeID, testCase.projectID)

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
