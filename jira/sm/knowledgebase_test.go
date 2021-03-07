package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestKnowledgebaseService_Articles(t *testing.T) {

	testCases := []struct {
		name               string
		query              string
		highlight          bool
		start, limit       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetKnowledgebaseArticlesWhenTheParametersAreCorrect",
			query:              "title",
			highlight:          true,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-knowledgebase-articles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/knowledgebase/article?limit=50&query=title&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetKnowledgebaseArticlesWhenTheHighlightIsFalse",
			query:              "title",
			highlight:          false,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-knowledgebase-articles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/knowledgebase/article?limit=50&query=title&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetKnowledgebaseArticlesWhenTheContextIsNil",
			query:              "title",
			highlight:          true,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-knowledgebase-articles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/knowledgebase/article?limit=50&query=title&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetKnowledgebaseArticlesWhenTheRequestMethodIsIncorrect",
			query:              "title",
			highlight:          true,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-knowledgebase-articles.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/servicedeskapi/knowledgebase/article?limit=50&query=title&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetKnowledgebaseArticlesWhenTheQueryIsNotSet",
			query:              "",
			highlight:          true,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-knowledgebase-articles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/knowledgebase/article?limit=50&query=title&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetKnowledgebaseArticlesWhenTheStatusCodeIsIncorrect",
			query:              "title",
			highlight:          true,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-knowledgebase-articles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/knowledgebase/article?limit=50&query=title&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetKnowledgebaseArticlesWhenTheEndpointIsEmpty",
			query:              "title",
			highlight:          true,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-knowledgebase-articles.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetKnowledgebaseArticlesWhenTheResponseBodyHasADifferentFormat",
			query:              "title",
			highlight:          true,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/knowledgebase/article?limit=50&query=title&start=0",
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

			service := &KnowledgebaseService{client: mockClient}
			gotResult, gotResponse, err := service.Articles(testCase.context, testCase.query, testCase.highlight, testCase.start, testCase.limit)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.StatusCode)
				assert.Equal(t, gotResponse.StatusCode, testCase.wantHTTPCodeReturn)

				for _, article := range gotResult.Values {
					t.Log("-------------------------------------------")
					t.Logf("Article Title: %v", article.Title)
					t.Logf("Article Source Type: %v", article.Source.Type)
					t.Log("-------------------------------------------")
				}
			}

		})
	}

}
