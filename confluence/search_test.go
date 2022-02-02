package confluence

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestSearchService_Content(t *testing.T) {

	testCases := []struct {
		name               string
		cql                string
		options            *models.SearchContentOptions
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name: "when the parameters are correct",
			cql:  "type=page",
			options: &models.SearchContentOptions{
				Context:                  "spaceKey",
				Cursor:                   "raNDoMsTRiNg",
				Next:                     true,
				Prev:                     true,
				Limit:                    20,
				Start:                    10,
				IncludeArchivedSpaces:    true,
				ExcludeCurrentSpaces:     true,
				SitePermissionTypeFilter: "externalCollaborator",
				Excerpt:                  "indexed",
				Expand:                   []string{"space"},
			},
			mockFile:       "./mocks/search-content.json",
			wantHTTPMethod: http.MethodGet,
			endpoint: "/wiki/rest/api/search?cql=type%3Dpage&cqlcontext=spaceKey&cursor=raNDoMsTRiNg&excerpt=" +
				"indexed&excludeCurrentSpaces=true&includeArchivedSpaces=true&limit=20&next=true&prev=true&sitePermission" +
				"TypeFilter=externalCollaborator&start=10",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name: "when the CQL query is not provided",
			cql:  "",
			options: &models.SearchContentOptions{
				Context:                  "spaceKey",
				Cursor:                   "raNDoMsTRiNg",
				Next:                     true,
				Prev:                     true,
				Limit:                    20,
				Start:                    10,
				IncludeArchivedSpaces:    true,
				ExcludeCurrentSpaces:     true,
				SitePermissionTypeFilter: "externalCollaborator",
				Excerpt:                  "indexed",
				Expand:                   []string{"space"},
			},
			mockFile:       "./mocks/search-content.json",
			wantHTTPMethod: http.MethodGet,
			endpoint: "/wiki/rest/api/search?cql=type%3Dpage&cqlcontext=spaceKey&cursor=raNDoMsTRiNg&excerpt=" +
				"indexed&excludeCurrentSpaces=true&includeArchivedSpaces=true&limit=20&next=true&prev=true&sitePermission" +
				"TypeFilter=externalCollaborator&start=10",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no CQL query set",
		},

		{
			name: "when the context is not provided",
			cql:  "type=page",
			options: &models.SearchContentOptions{
				Context:                  "spaceKey",
				Cursor:                   "raNDoMsTRiNg",
				Next:                     true,
				Prev:                     true,
				Limit:                    20,
				Start:                    10,
				IncludeArchivedSpaces:    true,
				ExcludeCurrentSpaces:     true,
				SitePermissionTypeFilter: "externalCollaborator",
				Excerpt:                  "indexed",
				Expand:                   []string{"space"},
			},
			mockFile:       "./mocks/search-content.json",
			wantHTTPMethod: http.MethodGet,
			endpoint: "/wiki/rest/api/search?cql=type%3Dpage&cqlcontext=spaceKey&cursor=raNDoMsTRiNg&excerpt=" +
				"indexed&excludeCurrentSpaces=true&includeArchivedSpaces=true&limit=20&next=true&prev=true&sitePermission" +
				"TypeFilter=externalCollaborator&start=10",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name: "when the response status is invalid",
			cql:  "type=page",
			options: &models.SearchContentOptions{
				Context:                  "spaceKey",
				Cursor:                   "raNDoMsTRiNg",
				Next:                     true,
				Prev:                     true,
				Limit:                    20,
				Start:                    10,
				IncludeArchivedSpaces:    true,
				ExcludeCurrentSpaces:     true,
				SitePermissionTypeFilter: "externalCollaborator",
				Excerpt:                  "indexed",
				Expand:                   []string{"space"},
			},
			mockFile:       "./mocks/search-content.json",
			wantHTTPMethod: http.MethodGet,
			endpoint: "/wiki/rest/api/search?cql=type%3Dpage&cqlcontext=spaceKey&cursor=raNDoMsTRiNg&excerpt=" +
				"indexed&excludeCurrentSpaces=true&includeArchivedSpaces=true&limit=20&next=true&prev=true&sitePermission" +
				"TypeFilter=externalCollaborator&start=10",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name: "when the response body is empty",
			cql:  "type=page",
			options: &models.SearchContentOptions{
				Context:                  "spaceKey",
				Cursor:                   "raNDoMsTRiNg",
				Next:                     true,
				Prev:                     true,
				Limit:                    20,
				Start:                    10,
				IncludeArchivedSpaces:    true,
				ExcludeCurrentSpaces:     true,
				SitePermissionTypeFilter: "externalCollaborator",
				Excerpt:                  "indexed",
				Expand:                   []string{"space"},
			},
			mockFile:       "./mocks/empty-json.json",
			wantHTTPMethod: http.MethodGet,
			endpoint: "/wiki/rest/api/search?cql=type%3Dpage&cqlcontext=spaceKey&cursor=raNDoMsTRiNg&excerpt=" +
				"indexed&excludeCurrentSpaces=true&includeArchivedSpaces=true&limit=20&next=true&prev=true&sitePermission" +
				"TypeFilter=externalCollaborator&start=10",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
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

			service := &SearchService{client: mockClient}

			gotResult, gotResponse, err := service.Content(testCase.context, testCase.cql, testCase.options)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

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

func TestSearchService_Users(t *testing.T) {

	testCases := []struct {
		name               string
		cql                string
		start, limit       int
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the parameters are correct",
			cql:                "type=user",
			start:              10,
			limit:              50,
			expand:             []string{"operations", "personalSpace"},
			mockFile:           "./mocks/search-content.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/search/user?cql=type%3Duser&expand=operations%2CpersonalSpace&limit=50&start=10",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:               "when the CQL query is not provided",
			cql:                "",
			start:              10,
			limit:              50,
			expand:             []string{"operations", "personalSpace"},
			mockFile:           "./mocks/search-content.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/search/user?cql=type%3Duser&expand=operations%2CpersonalSpace&limit=50&start=10",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no CQL query set",
		},

		{
			name:               "when the context is not provided",
			cql:                "type=user",
			start:              10,
			limit:              50,
			expand:             []string{"operations", "personalSpace"},
			mockFile:           "./mocks/search-content.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/search/user?cql=type%3Duser&expand=operations%2CpersonalSpace&limit=50&start=10",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is not provided",
			cql:                "type=user",
			start:              10,
			limit:              50,
			expand:             []string{"operations", "personalSpace"},
			mockFile:           "./mocks/search-content.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/search/user?cql=type%3Duser&expand=operations%2CpersonalSpace&limit=50&start=10",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the response body is empty",
			cql:                "type=user",
			start:              10,
			limit:              50,
			expand:             []string{"operations", "personalSpace"},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/search/user?cql=type%3Duser&expand=operations%2CpersonalSpace&limit=50&start=10",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
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

			service := &SearchService{client: mockClient}

			gotResult, gotResponse, err := service.Users(testCase.context, testCase.cql, testCase.start, testCase.limit, testCase.expand)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

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
