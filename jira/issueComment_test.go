package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestCommentService_Comments(t *testing.T) {

	testCases := []struct {
		name                string
		issueKeyOrID        string
		orderBy             string
		expands             []string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetIssueCommentsWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			orderBy:            "-created",
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			startAt:            0,
			maxResults:         100,
			mockFile:           "./mocks/get-issue-comments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&maxResults=100&orderBy=-created&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueCommentsWhenTheExpandParamIsNotSet",
			issueKeyOrID:       "DUMMY-3",
			orderBy:            "-created",
			expands:            nil,
			startAt:            0,
			maxResults:         100,
			mockFile:           "./mocks/get-issue-comments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?maxResults=100&orderBy=-created&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueCommentsWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			orderBy:            "-created",
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			startAt:            0,
			maxResults:         100,
			mockFile:           "./mocks/get-issue-comments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&maxResults=100&orderBy=-created&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentsWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			orderBy:            "-created",
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			startAt:            0,
			maxResults:         100,
			mockFile:           "./mocks/get-issue-comments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&maxResults=100&orderBy=-created&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentsWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			orderBy:            "-created",
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			startAt:            0,
			maxResults:         100,
			mockFile:           "./mocks/get-issue-comments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&maxResults=100&orderBy=-created&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentsWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			orderBy:            "-created",
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			startAt:            0,
			maxResults:         100,
			mockFile:           "./mocks/get-issue-comments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&maxResults=100&orderBy=-created&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentsWhenTheEndpointIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			orderBy:            "-created",
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			startAt:            0,
			maxResults:         100,
			mockFile:           "./mocks/get-issue-comments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/DUMMY-3/comment?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&maxResults=100&orderBy=-created&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentsWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			orderBy:            "-created",
			expands:            []string{"renderedFields", "names", "schema", "transitions", "changelog"},
			startAt:            0,
			maxResults:         100,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedFields%2Cnames%2Cschema%2Ctransitions%2Cchangelog&maxResults=100&orderBy=-created&startAt=0",
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

			i := &CommentService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(
				testCase.context,
				testCase.issueKeyOrID,
				testCase.orderBy,
				testCase.expands,
				testCase.startAt,
				testCase.maxResults,
			)

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

				for _, comment := range gotResult.Comments {

					t.Log("-------------------------")
					t.Logf("Comment ID: %v", comment.ID)
					t.Logf("Comment Created: %v", comment.Created)
					t.Logf("Comment Author EmailAddress: %v", comment.Author.EmailAddress)
					t.Logf("Comment Author AccountID: %v", comment.Author.AccountID)
					t.Log("------------------------- \n")
				}

			}
		})

	}

}

func TestCommentService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		commentID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteIssueCommentWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteIssueCommentWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			commentID:          "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueCommentWhenTheCommentIDIsNotSet",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueCommentWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "10001",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteIssueCommentWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
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

			i := &CommentService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.issueKeyOrID, testCase.commentID)

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

func TestCommentService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		commentID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueCommentWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "10001",
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueCommentWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			commentID:          "10001",
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentWhenTheCommentIDIsNotSet",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "",
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "10001",
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "10001",
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentWhenTheEndpointIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "10001",
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issue/DUMMY-3/comment/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "10001",
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueCommentWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			commentID:          "10001",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment/10001",
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

			i := &CommentService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.issueKeyOrID, testCase.commentID)

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

				t.Log("-------------------------")
				t.Logf("Comment ID: %v", gotResult.ID)
				t.Logf("Comment Created: %v", gotResult.Created)
				t.Logf("Comment Author EmailAddress: %v", gotResult.Author.EmailAddress)
				t.Logf("Comment Author AccountID: %v", gotResult.Author.AccountID)
				t.Log("------------------------- \n")

			}
		})

	}

}

func TestCommentService_Add(t *testing.T) {

	commentBody := CommentNodeScheme{}
	commentBody.Version = 1
	commentBody.Type = "doc"

	commentBody.AppendNode(&CommentNodeScheme{
		Type: "paragraph",
		Content: []*CommentNodeScheme{
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
	})

	testCases := []struct {
		name                                          string
		issueKeyOrID, visibilityType, visibilityValue string
		body                                          *CommentNodeScheme
		expands                                       []string
		mockFile                                      string
		wantHTTPMethod                                string
		endpoint                                      string
		context                                       context.Context
		wantHTTPCodeReturn                            int
		wantErr                                       bool
	}{
		{
			name:               "AddIssueCommentWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			visibilityType:     "role",
			visibilityValue:    "Administrators",
			body:               &commentBody,
			expands:            []string{"renderedBody"},
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "AddIssueCommentWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			visibilityType:     "role",
			visibilityValue:    "Administrators",
			body:               &commentBody,
			expands:            []string{"renderedBody"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddIssueCommentWhenTheIssueKeyIsNotSet",
			issueKeyOrID:       "",
			visibilityType:     "role",
			visibilityValue:    "Administrators",
			body:               &commentBody,
			expands:            []string{"renderedBody"},
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddIssueCommentWhenTheCommentBodyIsNotSet",
			issueKeyOrID:       "DUMMY-3",
			visibilityType:     "role",
			visibilityValue:    "Administrators",
			body:               nil,
			expands:            []string{"renderedBody"},
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddIssueCommentWhenTheVisibilityIsNotSet",
			issueKeyOrID:       "DUMMY-3",
			visibilityType:     "",
			visibilityValue:    "",
			body:               &commentBody,
			expands:            []string{"renderedBody"},
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "AddIssueCommentWhenTheExpandsAreNotSet",
			issueKeyOrID:       "DUMMY-3",
			visibilityType:     "role",
			visibilityValue:    "Administrators",
			body:               &commentBody,
			expands:            nil,
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddIssueCommentWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			visibilityType:     "role",
			visibilityValue:    "Administrators",
			body:               &commentBody,
			expands:            []string{"renderedBody"},
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddIssueCommentWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			visibilityType:     "role",
			visibilityValue:    "Administrators",
			body:               &commentBody,
			expands:            []string{"renderedBody"},
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AddIssueCommentWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			visibilityType:     "role",
			visibilityValue:    "Administrators",
			body:               &commentBody,
			expands:            []string{"renderedBody"},
			mockFile:           "./mocks/get-issue-comment-by-id.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/comment?expand=renderedBody",
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

			i := &CommentService{client: mockClient}

			gotResult, gotResponse, err := i.Add(testCase.context,
				testCase.issueKeyOrID, testCase.visibilityType,
				testCase.visibilityValue, testCase.body,
				testCase.expands)

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

				t.Log("-------------------------")
				t.Logf("Comment ID: %v", gotResult.ID)
				t.Logf("Comment Created: %v", gotResult.Created)
				t.Logf("Comment Author EmailAddress: %v", gotResult.Author.EmailAddress)
				t.Logf("Comment Author AccountID: %v", gotResult.Author.AccountID)
				t.Log("------------------------- \n")

			}
		})

	}

}
