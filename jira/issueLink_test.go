package jira

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIssueLinkService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		issueTypeLinkID    string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueLinksWhenTheIDIsCorrect",
			issueTypeLinkID:    "10001",
			mockFile:           "./mocks/get_issue_link_id_10001.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            false,
		},
		{
			name:               "GetIssueLinksWhenTheIDIsIncorrect",
			issueTypeLinkID:    "10002",
			mockFile:           "./mocks/get_issue_link_id_10001.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksWhenTheIDEmpty",
			issueTypeLinkID:    "",
			mockFile:           "./mocks/get_issue_link_id_10001.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksWhenTheHTTPMethodIsInvalid",
			issueTypeLinkID:    "10001",
			mockFile:           "./mocks/get_issue_link_id_10001.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
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
				Headers:            testCase.wantHTTPHeaders,
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

			i := &IssueLinkService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.issueTypeLinkID)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				assert.Equal(t, gotResult.ID, testCase.issueTypeLinkID)
			}
		})

	}

}

func TestIssueLinkService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		wantHTTPMethod     string
		issueTypeLinkID    string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteIssueLinksWhenTheIDCorrect",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			issueTypeLinkID:    "10000",
			endpoint:           "/rest/api/3/issueLink/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            false,
		},
		{
			name:               "DeleteIssueLinksWhenTheIDIncorrect",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			issueTypeLinkID:    "10001",
			endpoint:           "/rest/api/3/issueLink/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "DeleteIssueLinksWhenTheIDEmpty",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			issueTypeLinkID:    "",
			endpoint:           "/rest/api/3/issueLink/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "DeleteIssueLinksWhenTheURLIsIncorrect",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			issueTypeLinkID:    "",
			endpoint:           "/rest/api/3/issueLassdsadink/10000",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "DeleteIssueLinksWhenTheURLIsEmpty",
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			issueTypeLinkID:    "",
			endpoint:           "",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MethodAccepted:     testCase.wantHTTPMethod,
				Headers:            testCase.wantHTTPHeaders,
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

			i := &IssueLinkService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.issueTypeLinkID)

			if testCase.wantErr {
				assert.Error(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResponse.StatusCode, 204)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.Equal(t, gotResponse.StatusCode, 204)
			}
		})

	}

}
