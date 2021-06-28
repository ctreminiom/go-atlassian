package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestContentChildrenDescendantService_Children(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		expand             []string
		parentVersion      int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetContentChildrenWhenTheParametersAreCorrect",
			contentID:          "3958383",
			expand:             []string{"page", "comments"},
			parentVersion:      3,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/child?expand=page%2Ccomments&parentVersion=3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentChildrenWhenTheContentIDIsNotProvided",
			contentID:          "",
			expand:             []string{"page", "comments"},
			parentVersion:      3,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/child?expand=page%2Ccomments&parentVersion=3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentChildrenWhenTheExpandsAreNotProvided",
			contentID:          "3958383",
			expand:             nil,
			parentVersion:      3,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/child?parentVersion=3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentChildrenWhenTheContextIsNotProvided",
			contentID:          "3958383",
			expand:             []string{"page", "comments"},
			parentVersion:      3,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/child?expand=page%2Ccomments&parentVersion=3",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentChildrenWhenTheRequestMethodIsIncorrect",
			contentID:          "3958383",
			expand:             []string{"page", "comments"},
			parentVersion:      3,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/wiki/rest/api/content/3958383/child?expand=page%2Ccomments&parentVersion=3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentChildrenWhenTheStatusCodeIsIncorrect",
			contentID:          "3958383",
			expand:             []string{"page", "comments"},
			parentVersion:      3,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/child?expand=page%2Ccomments&parentVersion=3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetContentChildrenWhenTheResponseBodyIsEmpty",
			contentID:          "3958383",
			expand:             []string{"page", "comments"},
			parentVersion:      3,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/child?expand=page%2Ccomments&parentVersion=3",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &ContentChildrenDescendantService{client: mockClient}

			gotResult, gotResponse, err := service.Children(testCase.context, testCase.contentID, testCase.expand, testCase.parentVersion)

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

func TestContentChildrenDescendantService_ChildrenByType(t *testing.T) {

	testCases := []struct {
		name                   string
		contentID, contentType string
		parentVersion          int
		expand                 []string
		startAt, maxResults    int
		mockFile               string
		wantHTTPMethod         string
		endpoint               string
		context                context.Context
		wantHTTPCodeReturn     int
		wantErr                bool
	}{
		{
			name:               "GetContentChildrenByTypeWhenTheParametersAreCorrect",
			contentID:          "33473773",
			contentType:        "page",
			parentVersion:      1,
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentChildrenByTypeWhenTheContentIDIsNotProvided",
			contentID:          "",
			contentType:        "page",
			parentVersion:      1,
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentChildrenByTypeWhenTheContentTypeIsNotProvided",
			contentID:          "33473773",
			contentType:        "",
			parentVersion:      1,
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentChildrenByTypeWhenTheContentTypeIsInvalid",
			contentID:          "33473773",
			contentType:        "pages",
			parentVersion:      1,
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentChildrenByTypeWhenTheContextIsNotProvided",
			contentID:          "33473773",
			contentType:        "page",
			parentVersion:      1,
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentChildrenByTypeWhenTheRequestMethodIsIncorrect",
			contentID:          "33473773",
			contentType:        "page",
			parentVersion:      1,
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentChildrenByTypeWhenTheStatusCodeIsIncorrect",
			contentID:          "33473773",
			contentType:        "page",
			parentVersion:      1,
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantErr:            true,
		},

		{
			name:               "GetContentChildrenByTypeWhenTheResponseBodyIsEmpty",
			contentID:          "33473773",
			contentType:        "page",
			parentVersion:      1,
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &ContentChildrenDescendantService{client: mockClient}

			gotResult, gotResponse, err := service.ChildrenByType(testCase.context, testCase.contentID, testCase.contentType,
				testCase.parentVersion, testCase.expand, testCase.startAt, testCase.maxResults)

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

func TestContentChildrenDescendantService_Descendants(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetContentDescendantsWhenTheParametersAreCorrect",
			contentID:          "3958383",
			expand:             []string{"page", "comments"},
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/descendant?expand=page%2Ccomments",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentDescendantsWhenTheContentIDIsNotProvided",
			contentID:          "",
			expand:             []string{"page", "comments"},
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/descendant?expand=page%2Ccomments",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentDescendantsWhenTheExpandsAreNotProvided",
			contentID:          "3958383",
			expand:             nil,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/descendant",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentDescendantsWhenTheContextIsNotProvided",
			contentID:          "3958383",
			expand:             []string{"page", "comments"},
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/descendant?expand=page%2Ccomments",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentDescendantsWhenTheRequestMethodIsIncorrect",
			contentID:          "3958383",
			expand:             []string{"page", "comments"},
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/wiki/rest/api/content/3958383/descendant?expand=page%2Ccomments",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentDescendantsWhenTheStatusCodeIsIncorrect",
			contentID:          "3958383",
			expand:             []string{"page", "comments"},
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/descendant?expand=page%2Ccomments",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetContentDescendantsWhenTheResponseBodyIsEmpty",
			contentID:          "3958383",
			expand:             []string{"page", "comments"},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3958383/descendant?expand=page%2Ccomments",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &ContentChildrenDescendantService{client: mockClient}

			gotResult, gotResponse, err := service.Descendants(testCase.context, testCase.contentID, testCase.expand)

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

func TestContentChildrenDescendantService_DescendantsByType(t *testing.T) {

	testCases := []struct {
		name                          string
		contentID, contentType, depth string
		expand                        []string
		startAt, maxResults           int
		mockFile                      string
		wantHTTPMethod                string
		endpoint                      string
		context                       context.Context
		wantHTTPCodeReturn            int
		wantErr                       bool
	}{
		{
			name:               "GetContentDescendantsByTypeWhenTheParametersAreCorrect",
			contentID:          "33473773",
			contentType:        "page",
			depth:              "1",
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/descendant/page?depth=1&expand=childTypes.all&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentDescendantsByTypeWhenTheContentIDIsNotProvided",
			contentID:          "",
			contentType:        "page",
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentDescendantsByTypeWhenTheContentTypeIsNotProvided",
			contentID:          "33473773",
			contentType:        "",
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentDescendantsByTypeWhenTheContentTypeIsInvalid",
			contentID:          "33473773",
			contentType:        "pages",
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentDescendantsByTypeWhenTheContextIsNotProvided",
			contentID:          "33473773",
			contentType:        "page",
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentDescendantsByTypeWhenTheRequestMethodIsIncorrect",
			contentID:          "33473773",
			contentType:        "page",
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentDescendantsByTypeWhenTheStatusCodeIsIncorrect",
			contentID:          "33473773",
			contentType:        "page",
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantErr:            true,
		},

		{
			name:               "GetContentDescendantsByTypeWhenTheResponseBodyIsEmpty",
			contentID:          "33473773",
			contentType:        "page",
			expand:             []string{"childTypes.all"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/33473773/child/page?expand=childTypes.all&limit=50&parentVersion=1&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &ContentChildrenDescendantService{client: mockClient}

			gotResult, gotResponse, err := service.DescendantsByType(testCase.context, testCase.contentID, testCase.contentType,
				testCase.depth, testCase.expand, testCase.startAt, testCase.maxResults)

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

func TestContentChildrenDescendantService_CopyHierarchy(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		options            *CopyOptionsScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:      "CopyContentHierarchyWhenTheParametersAreCorrect",
			contentID: "44747372",
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				DestinationPageID:  "80412692",
				TitleOptions: &CopyTitleOptionScheme{
					Prefix: "copy-01-",
				},
			},
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/44747372/pagehierarchy/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            false,
		},

		{
			name:      "CopyContentHierarchyWhenTheContentIDIsNotProvided",
			contentID: "",
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				DestinationPageID:  "80412692",
				TitleOptions: &CopyTitleOptionScheme{
					Prefix: "copy-01-",
				},
			},
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/44747372/pagehierarchy/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},

		{
			name:               "CopyContentHierarchyWhenTheOptionsAreNotProvided",
			contentID:          "44747372",
			options:            nil,
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/44747372/pagehierarchy/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},

		{
			name:      "CopyContentHierarchyWhenTheContextIsNotProvided",
			contentID: "44747372",
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				DestinationPageID:  "80412692",
				TitleOptions: &CopyTitleOptionScheme{
					Prefix: "copy-01-",
				},
			},
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/44747372/pagehierarchy/copy",
			context:            nil,
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},

		{
			name:      "CopyContentHierarchyWhenTheRequestMethodIsIncorrect",
			contentID: "44747372",
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				DestinationPageID:  "80412692",
				TitleOptions: &CopyTitleOptionScheme{
					Prefix: "copy-01-",
				},
			},
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodConnect,
			endpoint:           "/wiki/rest/api/content/44747372/pagehierarchy/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},

		{
			name:      "CopyContentHierarchyWhenTheStatusCodeIsIncorrect",
			contentID: "44747372",
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				DestinationPageID:  "80412692",
				TitleOptions: &CopyTitleOptionScheme{
					Prefix: "copy-01-",
				},
			},
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/44747372/pagehierarchy/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:      "CopyContentHierarchyWhenTheResponseBodyIsEmpty",
			contentID: "44747372",
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				DestinationPageID:  "80412692",
				TitleOptions: &CopyTitleOptionScheme{
					Prefix: "copy-01-",
				},
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/44747372/pagehierarchy/copy",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &ContentChildrenDescendantService{client: mockClient}

			gotResult, gotResponse, err := service.CopyHierarchy(testCase.context, testCase.contentID, testCase.options)

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

func TestContentChildrenDescendantService_CopyPage(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		expand             []string
		options            *CopyOptionsScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:      "CopyPageWhenTheParametersAreCorrect",
			contentID: "747366262",
			expand:    []string{"childTypes.all", "container"},
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				PageTitle:          "new-page-copied-title",
				Destination: &CopyPageDestinationScheme{
					Type:  "parent_page",
					Value: "64290817",
				},
			},
			mockFile:           "./mocks/get-content.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/747366262/copy?expand=childTypes.all%2Ccontainer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:      "CopyPageWhenTheContentIDIsNotProvided",
			contentID: "",
			expand:    []string{"childTypes.all", "container"},
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				PageTitle:          "new-page-copied-title",
				Destination: &CopyPageDestinationScheme{
					Type:  "parent_page",
					Value: "64290817",
				},
			},
			mockFile:           "./mocks/get-content.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/747366262/copy?expand=childTypes.all%2Ccontainer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CopyPageWhenTheOptionsAreNotProvided",
			contentID:          "747366262",
			expand:             []string{"childTypes.all", "container"},
			options:            nil,
			mockFile:           "./mocks/get-content.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/747366262/copy?expand=childTypes.all%2Ccontainer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "CopyPageWhenTheRequestMethodIsIncorrect",
			contentID: "747366262",
			expand:    []string{"childTypes.all", "container"},
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				PageTitle:          "new-page-copied-title",
				Destination: &CopyPageDestinationScheme{
					Type:  "parent_page",
					Value: "64290817",
				},
			},
			mockFile:           "./mocks/get-content.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/747366262/copy?expand=childTypes.all%2Ccontainer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "CopyPageWhenTheContextIsNotProvided",
			contentID: "747366262",
			expand:    []string{"childTypes.all", "container"},
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				PageTitle:          "new-page-copied-title",
				Destination: &CopyPageDestinationScheme{
					Type:  "parent_page",
					Value: "64290817",
				},
			},
			mockFile:           "./mocks/get-content.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/747366262/copy?expand=childTypes.all%2Ccontainer",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "CopyPageWhenTheStatusCodeIsIncorrect",
			contentID: "747366262",
			expand:    []string{"childTypes.all", "container"},
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				PageTitle:          "new-page-copied-title",
				Destination: &CopyPageDestinationScheme{
					Type:  "parent_page",
					Value: "64290817",
				},
			},
			mockFile:           "./mocks/get-content.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/747366262/copy?expand=childTypes.all%2Ccontainer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:      "CopyPageWhenTheResponseBodyIsEmpty",
			contentID: "747366262",
			expand:    []string{"childTypes.all", "container"},
			options: &CopyOptionsScheme{
				CopyAttachments:    true,
				CopyPermissions:    true,
				CopyProperties:     true,
				CopyLabels:         true,
				CopyCustomContents: true,
				PageTitle:          "new-page-copied-title",
				Destination: &CopyPageDestinationScheme{
					Type:  "parent_page",
					Value: "64290817",
				},
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/747366262/copy?expand=childTypes.all%2Ccontainer",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &ContentChildrenDescendantService{client: mockClient}

			gotResult, gotResponse, err := service.CopyPage(testCase.context, testCase.contentID, testCase.expand, testCase.options)

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
