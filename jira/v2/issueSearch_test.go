package v2

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestIssueSearchService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		jql                string
		fields             []string
		expand             []string
		startAt            int
		maxResult          int
		validate           string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "SearchIssuesUsingGetWhenTheAllParametersAreCorrect",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/search?expand=changelogs%2Coperations&fields=summary%2Cassign%2Cresolution&jql=project+%3D+KP+and+issuetype+%3D+Story&maxResults=50&startAt=0&validateQuery=strict",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "SearchIssuesUsingGetWhenTheJQLIsNotProvided",
			jql:                "",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/search?expand=changelogs%2Coperations&fields=summary%2Cassign%2Cresolution&jql=project+%3D+KP+and+issuetype+%3D+Story&maxResults=50&startAt=0&validateQuery=strict",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchIssuesUsingGetWhenTheValidateParamIsEmpty",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/search?expand=changelogs%2Coperations&fields=summary%2Cassign%2Cresolution&jql=project+%3D+KP+and+issuetype+%3D+Story&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "SearchIssuesUsingGetWhenTheValidateParamIsIncorrect",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "stricts",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/search?expand=changelogs%2Coperations&fields=summary%2Cassign%2Cresolution&jql=project+%2D+KP+and+issuetype+%2D+Story&maxResults=50&startAt=0&validateQuery=strict",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "SearchIssuesUsingGetWhenTheEndpointIsIncorrect",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/searchs?expand=changelogs%2Coperations&fields=summary%2Cassign%2Cresolution&jql=project+%2D+KP+and+issuetype+%2D+Story&maxResults=50&startAt=0&validateQuery=strict",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchIssuesUsingGetWhenTheRequestMethodIsIncorrect",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/search?expand=changelogs%2Coperations&fields=summary%2Cassign%2Cresolution&jql=project+%2D+KP+and+issuetype+%2D+Story&maxResults=50&startAt=0&validateQuery=strict",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchIssuesUsingGetWhenTheStatusCodeIsIncorrect",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/search?expand=changelogs%2Coperations&fields=summary%2Cassign%2Cresolution&jql=project+%2D+KP+and+issuetype+%2D+Story&maxResults=50&startAt=0&validateQuery=strict",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "SearchIssuesUsingGetWhenTheContextIsNil",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/search?expand=changelogs%2Coperations&fields=summary%2Cassign%2Cresolution&jql=project+%2D+KP+and+issuetype+%2D+Story&maxResults=50&startAt=0&validateQuery=strict",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchIssuesUsingGetWhenTheResponseBodyHasADifferentFormat",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/search?expand=changelogs%2Coperations&fields=summary%2Cassign%2Cresolution&jql=project+%2D+KP+and+issuetype+%2D+Story&maxResults=50&startAt=0&validateQuery=strict",
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

			i := &IssueSearchService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.jql, testCase.fields, testCase.expand, testCase.startAt, testCase.maxResult, testCase.validate)

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

func TestIssueSearchService_Post(t *testing.T) {

	testCases := []struct {
		name               string
		jql                string
		fields             []string
		expand             []string
		startAt            int
		maxResult          int
		validate           string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "SearchIssuesUsingPostWhenTheAllParametersAreCorrect",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/search",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "SearchIssuesUsingPostWhenTheValidateParamIsEmpty",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/search",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "SearchIssuesUsingPostWhenTheExpandParameterIsEmpty",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             nil,
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/search",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "SearchIssuesUsingPostWhenTheEndpointIsIncorrect",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/search?expand=changelogs%2Coperations&fields=summary%2Cassign%2Cresolution&jql=project+%2D+KP+and+issuetype+%2D+Story&maxResults=50&startAt=0&validateQuery=strict",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchIssuesUsingPostWhenTheRequestMethodIsIncorrect",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/search",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchIssuesUsingPostWhenTheStatusCodeIsIncorrect",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/search",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "SearchIssuesUsingPostWhenTheContextIsNil",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/search-issues.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/search",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "SearchIssuesUsingPostWhenTheResponseBodyHasADifferentFormat",
			jql:                "project = KP and issuetype = Story",
			fields:             []string{"summary", "assign", "resolution"},
			expand:             []string{"changelogs", "operations"},
			startAt:            0,
			maxResult:          50,
			validate:           "strict",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/search",
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

			i := &IssueSearchService{client: mockClient}

			gotResult, gotResponse, err := i.Post(testCase.context, testCase.jql, testCase.fields, testCase.expand, testCase.startAt, testCase.maxResult, testCase.validate)

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
