package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestWorkflowService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		workflowID         string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteWorkflowWhenTheParametersAreCorrect",
			workflowID:         "as49949ja-asjdasjd94-jasdjasjd",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/workflow/as49949ja-asjdasjd94-jasdjasjd",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteWorkflowWhenTheWorkflowIDIsNotSet",
			workflowID:         "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/workflow/as49949ja-asjdasjd94-jasdjasjd",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteWorkflowWhenTheContextIsNotSet",
			workflowID:         "as49949ja-asjdasjd94-jasdjasjd",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/workflow/as49949ja-asjdasjd94-jasdjasjd",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteWorkflowWhenTheRequestMethodIsIncorrect",
			workflowID:         "as49949ja-asjdasjd94-jasdjasjd",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/3/workflow/as49949ja-asjdasjd94-jasdjasjd",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteWorkflowWhenTheStatusCodeIsIncorrect",
			workflowID:         "as49949ja-asjdasjd94-jasdjasjd",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/workflow/as49949ja-asjdasjd94-jasdjasjd",
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

			i := &WorkflowService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.workflowID)

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

func TestWorkflowService_Gets(t *testing.T) {

	testCases := []struct {
		name                  string
		startAt, maxResults   int
		mockFile              string
		workflowNames, expand []string
		wantHTTPMethod        string
		endpoint              string
		context               context.Context
		wantHTTPCodeReturn    int
		wantErr               bool
	}{
		{
			name:               "GetsWorkflowsWheTheParametersAreCorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-workflows.json",
			workflowNames:      []string{"workflow name 1", "workflow name 2"},
			expand:             []string{"transitions", "transitions.rules", "default"},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflow/search?expand=transitions%2Ctransitions.rules%2Cdefault&maxResults=50&startAt=0&workflowName=workflow+name+1&workflowName=workflow+name+2",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsWorkflowsWheTheContextIsNotSet",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-workflows.json",
			workflowNames:      []string{"workflow name 1", "workflow name 2"},
			expand:             []string{"transitions", "transitions.rules", "default"},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflow/search?expand=transitions%2Ctransitions.rules%2Cdefault&maxResults=50&startAt=0&workflowName=workflow+name+1&workflowName=workflow+name+2",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsWorkflowsWheTheRequestMethodIsIncorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-workflows.json",
			workflowNames:      []string{"workflow name 1", "workflow name 2"},
			expand:             []string{"transitions", "transitions.rules", "default"},
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/3/workflow/search?expand=transitions%2Ctransitions.rules%2Cdefault&maxResults=50&startAt=0&workflowName=workflow+name+1&workflowName=workflow+name+2",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsWorkflowsWheTheStatusCodeIsIncorrect",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-workflows.json",
			workflowNames:      []string{"workflow name 1", "workflow name 2"},
			expand:             []string{"transitions", "transitions.rules", "default"},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflow/search?expand=transitions%2Ctransitions.rules%2Cdefault&maxResults=50&startAt=0&workflowName=workflow+name+1&workflowName=workflow+name+2",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetsWorkflowsWheTheExpandsAreNotSet",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-workflows.json",
			workflowNames:      []string{"workflow name 1", "workflow name 2"},
			expand:             nil,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflow/search?maxResults=50&startAt=0&workflowName=workflow+name+1&workflowName=workflow+name+2",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsWorkflowsWheTheWorkflowNamesAreNotSet",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-workflows.json",
			workflowNames:      nil,
			expand:             []string{"transitions", "transitions.rules", "default"},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflow/search?expand=transitions%2Ctransitions.rules%2Cdefault&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsWorkflowsWheTheResponseBodyIsEmpty",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			workflowNames:      []string{"workflow name 1", "workflow name 2"},
			expand:             []string{"transitions", "transitions.rules", "default"},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflow/search?expand=transitions%2Ctransitions.rules%2Cdefault&maxResults=50&startAt=0&workflowName=workflow+name+1&workflowName=workflow+name+2",
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

			i := &WorkflowService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.workflowNames, testCase.expand,
				testCase.startAt,
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
			}
		})

	}

}

func TestWorkflowService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.WorkflowPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateWorkflowsWhenTheParametersAreCorrect",
			payload: &models.WorkflowPayloadScheme{
				Name:        "Workflow 1",
				Description: "This is a workflow used for Stories and Tasks",
				Statuses: []*models.WorkflowTransitionScreenScheme{
					{
						ID: "10002",
					},
					{
						ID: "3",
					},
				},
				Transitions: []*models.WorkflowTransitionPayloadScheme{
					{
						Name:   "Create",
						To:     "10012",
						Type:   "initial",
						Screen: nil,
					},
					{
						Name: "Start Progress",
						To:   "3",
						Type: "directed",
						From: []string{"10012"},
					},
					{
						Name: "Done",
						To:   "10002",
						Type: "directed",
						From: []string{"10012"},
					},
					{
						Name: "Stop Progress",
						To:   "10002",
						Type: "directed",
						From: []string{"3"},
					},
				},
			},
			mockFile:           "../mocks/workflow-created.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/workflow",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "CreateWorkflowsWhenThePayloadIsNotProvided",
			payload:            nil,
			mockFile:           "../mocks/workflow-created.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/workflow",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "CreateWorkflowsWhenTheContextIsNotProvided",
			payload: &models.WorkflowPayloadScheme{
				Name:        "Workflow 1",
				Description: "This is a workflow used for Stories and Tasks",
				Statuses: []*models.WorkflowTransitionScreenScheme{
					{
						ID: "10002",
					},
					{
						ID: "3",
					},
				},
				Transitions: []*models.WorkflowTransitionPayloadScheme{
					{
						Name:   "Create",
						To:     "10012",
						Type:   "initial",
						Screen: nil,
					},
					{
						Name: "Start Progress",
						To:   "3",
						Type: "directed",
						From: []string{"10012"},
					},
					{
						Name: "Done",
						To:   "10002",
						Type: "directed",
						From: []string{"10012"},
					},
					{
						Name: "Stop Progress",
						To:   "10002",
						Type: "directed",
						From: []string{"3"},
					},
				},
			},
			mockFile:           "../mocks/workflow-created.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/workflow",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "CreateWorkflowsWhenTheRequestMethodIsIncorrect",
			payload: &models.WorkflowPayloadScheme{
				Name:        "Workflow 1",
				Description: "This is a workflow used for Stories and Tasks",
				Statuses: []*models.WorkflowTransitionScreenScheme{
					{
						ID: "10002",
					},
					{
						ID: "3",
					},
				},
				Transitions: []*models.WorkflowTransitionPayloadScheme{
					{
						Name:   "Create",
						To:     "10012",
						Type:   "initial",
						Screen: nil,
					},
					{
						Name: "Start Progress",
						To:   "3",
						Type: "directed",
						From: []string{"10012"},
					},
					{
						Name: "Done",
						To:   "10002",
						Type: "directed",
						From: []string{"10012"},
					},
					{
						Name: "Stop Progress",
						To:   "10002",
						Type: "directed",
						From: []string{"3"},
					},
				},
			},
			mockFile:           "../mocks/workflow-created.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/workflow",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "CreateWorkflowsWhenTheStatusCodeIsIncorrect",
			payload: &models.WorkflowPayloadScheme{
				Name:        "Workflow 1",
				Description: "This is a workflow used for Stories and Tasks",
				Statuses: []*models.WorkflowTransitionScreenScheme{
					{
						ID: "10002",
					},
					{
						ID: "3",
					},
				},
				Transitions: []*models.WorkflowTransitionPayloadScheme{
					{
						Name:   "Create",
						To:     "10012",
						Type:   "initial",
						Screen: nil,
					},
					{
						Name: "Start Progress",
						To:   "3",
						Type: "directed",
						From: []string{"10012"},
					},
					{
						Name: "Done",
						To:   "10002",
						Type: "directed",
						From: []string{"10012"},
					},
					{
						Name: "Stop Progress",
						To:   "10002",
						Type: "directed",
						From: []string{"3"},
					},
				},
			},
			mockFile:           "../mocks/workflow-created.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/workflow",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CreateWorkflowsWhenTheResponseBodyIsEmpty",
			payload: &models.WorkflowPayloadScheme{
				Name:        "Workflow 1",
				Description: "This is a workflow used for Stories and Tasks",
				Statuses: []*models.WorkflowTransitionScreenScheme{
					{
						ID: "10002",
					},
					{
						ID: "3",
					},
				},
				Transitions: []*models.WorkflowTransitionPayloadScheme{
					{
						Name:   "Create",
						To:     "10012",
						Type:   "initial",
						Screen: nil,
					},
					{
						Name: "Start Progress",
						To:   "3",
						Type: "directed",
						From: []string{"10012"},
					},
					{
						Name: "Done",
						To:   "10002",
						Type: "directed",
						From: []string{"10012"},
					},
					{
						Name: "Stop Progress",
						To:   "10002",
						Type: "directed",
						From: []string{"3"},
					},
				},
			},
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/workflow",
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

			i := &WorkflowService{client: mockClient}

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
