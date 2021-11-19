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

func TestIssueWorklogService_Issue(t *testing.T) {

	testCases := []struct {
		name                       string
		ctx                        context.Context
		issueKeyOrID               string
		startAt, maxResults, after int
		expand                     []string
		mockFile                   string
		wantHTTPMethod             string
		endpoint                   string
		wantHTTPCodeReturn         int
		wantErr                    bool
	}{
		{
			name:               "GetIssueWorklogWhenTheParametersAreCorrect",
			ctx:                context.Background(),
			issueKeyOrID:       "KP-1",
			startAt:            0,
			maxResults:         50,
			after:              11112,
			expand:             []string{"expand", "now"},
			mockFile:           "../v3/mocks/get-issue-worklogs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?expand=expand%2Cnow&maxResults=50&startAt=0&startedAfter=11112",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueWorklogWhenTheIssueKeyOrIDIsNotProvided",
			ctx:                context.Background(),
			issueKeyOrID:       "",
			startAt:            0,
			maxResults:         50,
			after:              11112,
			expand:             []string{"expand", "now"},
			mockFile:           "../v3/mocks/get-issue-worklogs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?expand=expand%2Cnow&maxResults=50&startAt=0&startedAfter=11112",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogWhenTheContextIsNotProvided",
			ctx:                nil,
			issueKeyOrID:       "KP-1",
			startAt:            0,
			maxResults:         50,
			after:              11112,
			expand:             []string{"expand", "now"},
			mockFile:           "../v3/mocks/get-issue-worklogs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?expand=expand%2Cnow&maxResults=50&startAt=0&startedAfter=11112",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogWhenTheRequestMethodIsIncorrect",
			ctx:                context.Background(),
			issueKeyOrID:       "KP-1",
			startAt:            0,
			maxResults:         50,
			after:              11112,
			expand:             []string{"expand", "now"},
			mockFile:           "../v3/mocks/get-issue-worklogs.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?expand=expand%2Cnow&maxResults=50&startAt=0&startedAfter=11112",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogWhenTheStatusCodeIsIncorrect",
			ctx:                context.Background(),
			issueKeyOrID:       "KP-1",
			startAt:            0,
			maxResults:         50,
			after:              11112,
			expand:             []string{"expand", "now"},
			mockFile:           "../v3/mocks/get-issue-worklogs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?expand=expand%2Cnow&maxResults=50&startAt=0&startedAfter=11112",
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

			i := &IssueWorklogService{client: mockClient}

			gotResult, gotResponse, err := i.Issue(testCase.ctx, testCase.issueKeyOrID, testCase.startAt,
				testCase.maxResults, testCase.after, testCase.expand)

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

func TestIssueWorklogService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		ctx                context.Context
		issueKeyOrID       string
		worklogID          string
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueWorklogWhenTheParametersAreCorrect",
			ctx:                context.Background(),
			issueKeyOrID:       "KP-1",
			worklogID:          "10000",
			expand:             []string{"all"},
			mockFile:           "../v3/mocks/get-issue-worklog.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?expand=all",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueWorklogWhenTheContextIsNotProvided",
			ctx:                nil,
			issueKeyOrID:       "KP-1",
			worklogID:          "10000",
			expand:             []string{"all"},
			mockFile:           "../v3/mocks/get-issue-worklog.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?expand=all",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogWhenTheIssueKeyOrIDIsNotProvided",
			ctx:                context.Background(),
			issueKeyOrID:       "",
			worklogID:          "10000",
			expand:             []string{"all"},
			mockFile:           "../v3/mocks/get-issue-worklog.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?expand=all",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogWhenTheWorklogIDIsNotProvided",
			ctx:                context.Background(),
			issueKeyOrID:       "KP-1",
			worklogID:          "",
			expand:             []string{"all"},
			mockFile:           "../v3/mocks/get-issue-worklog.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?expand=all",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogWhenTheRequestMethodIsIncorrect",
			ctx:                context.Background(),
			issueKeyOrID:       "KP-1",
			worklogID:          "10000",
			expand:             []string{"all"},
			mockFile:           "../v3/mocks/get-issue-worklog.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?expand=all",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogWhenTheStatusCodeIsIncorrect",
			ctx:                context.Background(),
			issueKeyOrID:       "KP-1",
			worklogID:          "10000",
			expand:             []string{"all"},
			mockFile:           "../v3/mocks/get-issue-worklog.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?expand=all",
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

			i := &IssueWorklogService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.ctx, testCase.issueKeyOrID, testCase.worklogID,
				testCase.expand)

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

func TestIssueWorklogService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		ctx                context.Context
		issueKeyOrID       string
		options            *models.WorklogOptionsScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:         "AddIssueWorklogWhenTheParametersAreCorrect",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:         "AddIssueWorklogWhenTheIssueKeyOrIDIsNotProvided",
			ctx:          context.Background(),
			issueKeyOrID: "",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:         true,
				AdjustEstimate: "auto",
				ReduceBy:       "2h",
				//OverrideEditableFlag: true,
				Expand: []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?adjustEstimate=auto&expand=expand%2Cproperties&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:         "AddIssueWorklogWhenTheContextIsNotProvided",
			ctx:          nil,
			issueKeyOrID: "KP-1",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:         true,
				AdjustEstimate: "auto",
				ReduceBy:       "2h",
				//OverrideEditableFlag: true,
				Expand: []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?adjustEstimate=auto&expand=expand%2Cproperties&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:         "AddIssueWorklogWhenThePayloadOptionIsNotProvided",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:         true,
				AdjustEstimate: "auto",
				ReduceBy:       "2h",
				//OverrideEditableFlag: true,
				Expand:  []string{"expand", "properties"},
				Payload: nil,
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?adjustEstimate=auto&expand=expand%2Cproperties&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:         "AddIssueWorklogWhenTheRequestMethodIsIncorrect",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:         true,
				AdjustEstimate: "auto",
				ReduceBy:       "2h",
				//OverrideEditableFlag: true,
				Expand: []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?adjustEstimate=auto&expand=expand%2Cproperties&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:         "AddIssueWorklogWhenTheStatusCodeIsIncorrect",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:         true,
				AdjustEstimate: "auto",
				ReduceBy:       "2h",
				//OverrideEditableFlag: true,
				Expand: []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issue/KP-1/worklog?adjustEstimate=auto&expand=expand%2Cproperties&reduceBy=2h",
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

			i := &IssueWorklogService{client: mockClient}

			gotResult, gotResponse, err := i.Add(testCase.ctx, testCase.issueKeyOrID, testCase.options)

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

func TestIssueWorklogService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		ctx                context.Context
		issueKeyOrID       string
		worklogID          string
		options            *models.WorklogOptionsScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:         "UpdateIssueWorklogWhenTheParametersAreCorrect",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:         "UpdateIssueWorklogWhenTheContextIsNotProvided",
			ctx:          nil,
			issueKeyOrID: "KP-1",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:         "UpdateIssueWorklogWhenTheIssueKeyOrIDIsNotProvided",
			ctx:          context.Background(),
			issueKeyOrID: "",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:         "UpdateIssueWorklogWhenTheWorklogIDIsNotProvided",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			worklogID:    "",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:         "UpdateIssueWorklogWhenThePayloadIsNotProvided",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload:              nil,
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:         "UpdateIssueWorklogWhenTheRequestMethodIsIncorrect",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:         "UpdateIssueWorklogWhenTheStatusCodeIsIncorrect",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
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

			i := &IssueWorklogService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.ctx, testCase.issueKeyOrID, testCase.worklogID, testCase.options)

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

func TestIssueWorklogService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		ctx                context.Context
		issueKeyOrID       string
		worklogID          string
		options            *models.WorklogOptionsScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:         "DeleteIssueWorklogWhenTheParametersAreCorrect",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:         "DeleteIssueWorklogWhenTheContextIsNotProvided",
			ctx:          nil,
			issueKeyOrID: "KP-1",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "DeleteIssueWorklogWhenTheIssueKeyOrIDIsNotProvided",
			ctx:          context.Background(),
			issueKeyOrID: "",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "DeleteIssueWorklogWhenTheWorklogIDIsNotProvided",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			worklogID:    "",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueWorklogWhenTheOptionsAreNotProvided",
			ctx:                context.Background(),
			issueKeyOrID:       "KP-1",
			worklogID:          "10000",
			mockFile:           "../v3/mocks/get-issue-worklog.json",
			options:            nil,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "DeleteIssueWorklogWhenTheRequestMethodIsIncorrect",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
			},
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:         "DeleteIssueWorklogWhenTheStatusCodeIsIncorrect",
			ctx:          context.Background(),
			issueKeyOrID: "KP-1",
			worklogID:    "10000",
			mockFile:     "../v3/mocks/get-issue-worklog.json",
			options: &models.WorklogOptionsScheme{
				Notify:               false,
				AdjustEstimate:       "auto",
				ReduceBy:             "2h",
				OverrideEditableFlag: true,
				NewEstimate:          "2h",
				Expand:               []string{"expand", "properties"},
				Payload: &models.WorklogPayloadScheme{
					/*
						Visibility:       &jira.IssueWorklogVisibilityScheme{
							Type:  "group",
							Value: "jira-users",
						},
					*/
					Started:          "2021-07-16T07:01:10.774+0000",
					TimeSpentSeconds: 12000,
				},
			},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issue/KP-1/worklog/10000?adjustEstimate=auto&expand=expand%2Cproperties&newEstimate=2h&notifyUsers=false&overrideEditableFlag=true&reduceBy=2h",
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

			i := &IssueWorklogService{client: mockClient}

			gotResponse, err := i.Delete(testCase.ctx, testCase.issueKeyOrID, testCase.worklogID, testCase.options)

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

func TestIssueWorklogService_Deleted(t *testing.T) {

	testCases := []struct {
		name               string
		ctx                context.Context
		since              int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueWorklogsDeletedWhenTheParametersAreCorrect",
			ctx:                context.Background(),
			since:              1000202,
			mockFile:           "../v3/mocks/get-issue-worklogs-changelogs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/worklog/deleted?since=1000202",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueWorklogsDeletedWhenTheContextIsNotProvided",
			ctx:                nil,
			since:              1000202,
			mockFile:           "../v3/mocks/get-issue-worklogs-changelogs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/worklog/deleted?since=1000202",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogsDeletedWhenTheRequestMethodIsIncorrect",
			ctx:                context.Background(),
			since:              1000202,
			mockFile:           "../v3/mocks/get-issue-worklogs-changelogs.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/worklog/deleted?since=1000202",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogsDeletedWhenTheStatusCodeIsIncorrect",
			ctx:                context.Background(),
			since:              1000202,
			mockFile:           "../v3/mocks/get-issue-worklogs-changelogs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/worklog/deleted?since=1000202",
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

			i := &IssueWorklogService{client: mockClient}

			gotResult, gotResponse, err := i.Deleted(testCase.ctx, testCase.since)

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

func TestIssueWorklogService_Updated(t *testing.T) {

	testCases := []struct {
		name               string
		ctx                context.Context
		since              int
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueWorklogsUpdatedWhenTheParametersAreCorrect",
			ctx:                context.Background(),
			since:              1000202,
			expand:             []string{"all"},
			mockFile:           "../v3/mocks/get-issue-worklogs-changelogs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/worklog/updated?expand=all&since=1000202",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueWorklogsUpdatedWhenTheContextIsNotProvided",
			ctx:                nil,
			since:              1000202,
			expand:             []string{"all"},
			mockFile:           "../v3/mocks/get-issue-worklogs-changelogs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/worklog/updated?expand=all&since=1000202",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogsUpdatedWhenTheRequestMethodIsIncorrect",
			ctx:                context.Background(),
			since:              1000202,
			expand:             []string{"all"},
			mockFile:           "../v3/mocks/get-issue-worklogs-changelogs.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/2/worklog/updated?expand=all&since=1000202",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogsUpdatedWhenTheStatusCodeIsIncorrect",
			ctx:                context.Background(),
			since:              1000202,
			expand:             []string{"all"},
			mockFile:           "../v3/mocks/get-issue-worklogs-changelogs.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/worklog/updated?expand=all&since=1000202",
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

			i := &IssueWorklogService{client: mockClient}

			gotResult, gotResponse, err := i.Updated(testCase.ctx, testCase.since, testCase.expand)

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

func TestIssueWorklogService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		ctx                context.Context
		worklogIDs         []int
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueWorklogsWhenTheParametersAreCorrect",
			ctx:                context.Background(),
			worklogIDs:         []int{10000, 10001, 10002},
			expand:             []string{"expand", "now"},
			mockFile:           "../v3/mocks/get-worklogs.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/worklog/list?expand=expand%2Cnow",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueWorklogsWhenTheContextIsNotProvided",
			ctx:                nil,
			worklogIDs:         []int{10000, 10001, 10002},
			expand:             []string{"expand", "now"},
			mockFile:           "../v3/mocks/get-worklogs.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/worklog/list?expand=expand%2Cnow",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogsWhenTheWorklogsIDsAreNotProvided",
			ctx:                context.Background(),
			worklogIDs:         nil,
			expand:             []string{"expand", "now"},
			mockFile:           "../v3/mocks/get-worklogs.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/worklog/list?expand=expand%2Cnow",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogsWhenTheRequestMethodIsIncorrect",
			ctx:                context.Background(),
			worklogIDs:         []int{10000, 10001, 10002},
			expand:             []string{"expand", "now"},
			mockFile:           "../v3/mocks/get-worklogs.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/2/worklog/list?expand=expand%2Cnow",
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueWorklogsWhenTheStatusCodeIsIncorrect",
			ctx:                context.Background(),
			worklogIDs:         []int{10000, 10001, 10002},
			expand:             []string{"expand", "now"},
			mockFile:           "../v3/mocks/get-worklogs.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/worklog/list?expand=expand%2Cnow",
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

			i := &IssueWorklogService{client: mockClient}
			gotResult, gotResponse, err := i.Gets(testCase.ctx, testCase.worklogIDs, testCase.expand)

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
