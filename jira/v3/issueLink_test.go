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

func TestIssueLinkService_Create(t *testing.T) {

	payloadMocked := &models.LinkPayloadScheme{

		Comment: &models.CommentPayloadScheme{

			Body: &models.CommentNodeScheme{
				Version: 1,
				Type:    "doc",
				Content: []*models.CommentNodeScheme{
					{
						Type: "paragraph",
						Content: []*models.CommentNodeScheme{
							{
								Type: "text",
								Text: "Carlos Test",
							},
							{
								Type: "emoji",
								Attrs: map[string]interface{}{
									"shortName": ":grin",
									"id":        "1f601",
									"text":      "😁",
								},
							},
							{
								Type: "text",
								Text: " ",
							},
						},
					},
				},
			},
		},

		InwardIssue: &models.LinkedIssueScheme{
			Key: "KP-1",
		},
		OutwardIssue: &models.LinkedIssueScheme{
			Key: "KP-2",
		},
		Type: &models.LinkTypeScheme{
			Name: "Duplicate",
		},
	}

	testCases := []struct {
		name               string
		payload            *models.LinkPayloadScheme
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CreateIssueLinkWhenThePayloadAreCorrect",
			payload:            payloadMocked,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLink",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateIssueLinkWhenThePayloadIsNotProvided",
			payload:            nil,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLink",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssueLinkWhenTheEndpointIsIncorrect",
			payload:            payloadMocked,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinks",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssueLinkWhenTheRequestMethodIsIncorrect",
			payload:            payloadMocked,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issueLink",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssueLinkWhenTheStatusCodeIsIncorrect",
			payload:            payloadMocked,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLink",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "CreateIssueLinkWhenTheContextIsNil",
			payload:            payloadMocked,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLink",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
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

			service := &IssueLinkService{client: mockClient}
			gotResponse, err := service.Create(testCase.context, testCase.payload)

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

func TestIssueLinkService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		linkID             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteIssueLinkWhenTheLinkIDIDIsCorrect",
			linkID:             "10001",
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "DeleteIssueLinkWhenTheLinkIDIsNotProvided",
			linkID:             "",
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkWhenTheLinkIDIDIsIncorrect",
			linkID:             "10002",
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkWhenTheRequestMethodIsIncorrect",
			linkID:             "10001",
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkWhenTheStatusCodeIsIncorrect",
			linkID:             "10001",
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkWhenTheContextIsNil",
			linkID:             "10001",
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            nil,
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

			service := &IssueLinkService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.linkID)

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

func TestIssueLinkService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		linkID             string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueLinkWhenTheIDIsCorrect",
			linkID:             "10001",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueLinkWhenTheLinkIDIsNotProvided",
			linkID:             "",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueLinkWhenTheIDIsIncorrect",
			linkID:             "10002",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueLinkWhenTheEndpointIsIncorrect",
			linkID:             "10001",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLinks/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueLinkWhenTheContextIsNil",
			linkID:             "10001",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueLinkWhenTheRequestMethodIsIncorrect",
			linkID:             "10001",
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueLinkWhenTheStatusCodeIsIncorrect",
			linkID:             "10001",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issueLink/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetIssueLinkWhenTheResponseBodyHasADifferentFormat",
			linkID:             "10001",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/empty_json.json",
			endpoint:           "/rest/api/3/issueLink/10001",
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

			service := &IssueLinkService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.linkID)

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

func TestIssueLinkService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetsIssueLinkWhenTheIssueKeyOrIDIsCorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issue/DUMMY-3?fields=issuelinks",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetsIssueLinkWhenTheIssueKeyOrIDIsNotProvided",
			issueKeyOrID:       "",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issue/DUMMY-3?fields=issuelinks",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueLinkWhenTheIssueKeyOrIDIsIncorrect",
			issueKeyOrID:       "DUMMY-4",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issue/DUMMY-3?fields=issuelinks",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueLinkWhenTheEndpointIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issues/DUMMY-3?fields=issuelinks",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueLinkWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issue/DUMMY-3?fields=issuelinks",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueLinkWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issue/DUMMY-3?fields=issuelinks",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetsIssueLinkWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/get-issue-link-by-id.json",
			endpoint:           "/rest/api/3/issue/DUMMY-3?fields=issuelinks",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetsIssueLinkWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "./mocks/empty_json.json",
			endpoint:           "/rest/api/3/issue/DUMMY-3?fields=issuelinks",
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

			service := &IssueLinkService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.issueKeyOrID)

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
