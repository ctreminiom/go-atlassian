package agile

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models/agile"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestEpicService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		epicIDOrKey        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetEpicWhenTheParametersAreCorrect",
			epicIDOrKey:        "KP-16",
			mockFile:           "../mocks/get-epic.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/epic/KP-16",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetEpicWhenTheEpicKeyOrIDIsNotSet",
			epicIDOrKey:        "",
			mockFile:           "../mocks/get-epic.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/epic/KP-16",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetEpicWhenTheRequestMethodIsIncorrect",
			epicIDOrKey:        "KP-16",
			mockFile:           "../mocks/get-epic.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/epic/KP-16",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetEpicWhenTheContextIsNil",
			epicIDOrKey:        "KP-16",
			mockFile:           "../mocks/get-epic.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/epic/KP-16",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetEpicWhenTheResponseStatusCodeIsIncorrect",
			epicIDOrKey:        "KP-16",
			mockFile:           "../mocks/get-epic.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/epic/KP-16",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetEpicWhenTheResponseBodyIsEmpty",
			epicIDOrKey:        "KP-16",
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/epic/KP-16",
			context:            context.Background(),
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

			service := &EpicService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.epicIDOrKey)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}
		})
	}

}

func TestEpicService_Issues(t *testing.T) {

	testCases := []struct {
		name                string
		epicIDOrKey         string
		startAt, maxResults int
		opts                *model.IssueOptionScheme
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:        "GetEpicIssuesBacklogWhenTheParametersAreCorrect",
			epicIDOrKey: "KP-16",
			startAt:     0,
			maxResults:  50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-epic-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/epic/KP-16/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:        "GetEpicIssuesBacklogWhenTheEpicIDIsKeyIsNotSet",
			epicIDOrKey: "",
			startAt:     0,
			maxResults:  50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-epic-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/epic/KP-16/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "GetEpicIssuesBacklogWhenTheResponseBodyIsEmpty",
			epicIDOrKey: "KP-16",
			startAt:     0,
			maxResults:  50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/epic/KP-16/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "GetEpicIssuesWhenTheValidateQueryOptionIsNotEnabled",
			epicIDOrKey: "KP-16",
			startAt:     0,
			maxResults:  50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: false,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile: "../mocks/get-epic-issues.json",

			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/epic/KP-16/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:        "GetEpicIssuesWhenTheFilterIDIsNotSet",
			epicIDOrKey: "KP-16",

			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog"},
			},
			mockFile: "../mocks/get-epic-issues.json",

			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=true",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:        "GetEpicIssuesWhenTheRequestMethodIsIncorrect",
			epicIDOrKey: "KP-16",

			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-epic-issues.json",

			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=true",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:        "GetEpicIssuesWhenTheContextIsNil",
			epicIDOrKey: "KP-16",

			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-epic-issues.json",

			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=true",
			context:            nil,
			wantErr:            true,
		},

		{
			name:        "GetEpicIssuesWhenTheResponseStatusCodeIsIncorrect",
			epicIDOrKey: "KP-16",

			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile: "../mocks/get-epic-issues.json",

			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=true",
			context:            context.Background(),
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

			service := &EpicService{client: mockClient}
			gotResult, gotResponse, err := service.Issues(testCase.context, testCase.epicIDOrKey,
				testCase.startAt, testCase.maxResults, testCase.opts)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}
		})
	}

}

func TestEpicService_Move(t *testing.T) {

	testCases := []struct {
		name               string
		epicIDOrKey        string
		issues             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "MoveIssuesToEpicWhenTheParametersAreCorrect",
			epicIDOrKey:        "EPIC-1",
			issues:             []string{"STORY-1"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/epic/EPIC-1/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},
		{
			name:               "MoveIssuesToEpicWhenTheEpicKeyIsNotProvided",
			epicIDOrKey:        "",
			issues:             []string{"STORY-1"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/epic/EPIC-1/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "MoveIssuesToEpicWhenTheContextIsNotProvided",
			epicIDOrKey:        "EPIC-1",
			issues:             []string{"STORY-1"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/epic/EPIC-1/issue",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "MoveIssuesToEpicWhenTheRequestMethodIsIncorrect",
			epicIDOrKey:        "EPIC-1",
			issues:             []string{"STORY-1"},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/agile/1.0/epic/EPIC-1/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "MoveIssuesToEpicWhenTheResponseStatusCodeIsIncorrect",
			epicIDOrKey:        "EPIC-1",
			issues:             []string{"STORY-1"},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/epic/EPIC-1/issue",
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

			service := &EpicService{client: mockClient}
			gotResponse, err := service.Move(testCase.context, testCase.epicIDOrKey, testCase.issues)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}

			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}
		})
	}
}
