package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestRequestCommentService_Attachments(t *testing.T) {

	testCases := []struct {
		name                    string
		issueKeyOrID            string
		commentID, start, limit int
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:               "GetCustomerCommentAttachmentsWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          11111,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comment-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/11111/attachment?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerCommentAttachmentsWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			commentID:          11111,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comment-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/11111/attachment?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentAttachmentsWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          11111,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comment-attachments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/11111/attachment?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetCustomerCommentAttachmentsWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          11111,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comment-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/11111/attachment?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentAttachmentsWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			commentID:          11111,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comment-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/11111/attachment?limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentAttachmentsWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			commentID:          11111,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comment-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentAttachmentsWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			commentID:          11111,
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/11111/attachment?limit=50&start=0",
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

			service := &RequestCommentService{client: mockClient}
			gotResult, gotResponse, err := service.Attachments(testCase.context, testCase.issueKeyOrID,
				testCase.commentID, testCase.start, testCase.limit)

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

				for _, attachment := range gotResult.Values {
					t.Log(attachment.Filename, attachment.MimeType, attachment.Size)
				}

			}

		})
	}

}

func TestRequestCommentService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID, body string
		public             bool
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CreateCustomerCommentWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			body:               "Hello There",
			public:             true,
			mockFile:           "./mocks/create-request-comment.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateCustomerCommentWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			body:               "Hello There",
			public:             true,
			mockFile:           "./mocks/create-request-comment.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerCommentWhenTheBodyIsNotSet",
			issueKeyOrID:       "DUMMY-3",
			body:               "",
			public:             true,
			mockFile:           "./mocks/create-request-comment.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerCommentWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			body:               "Hello There",
			public:             true,
			mockFile:           "./mocks/create-request-comment.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerCommentWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			body:               "Hello There",
			public:             true,
			mockFile:           "./mocks/create-request-comment.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerCommentWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			body:               "Hello There",
			public:             true,
			mockFile:           "./mocks/create-request-comment.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "CreateCustomerCommentWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			body:               "Hello There",
			public:             true,
			mockFile:           "./mocks/create-request-comment.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateCustomerCommentWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			body:               "Hello There",
			public:             true,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment",
			context:            context.Background(),
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

			service := &RequestCommentService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.issueKeyOrID, testCase.body, testCase.public)

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

				t.Log("----------------------------------")
				t.Logf("Comment, ID: %v", gotResult.ID)
				t.Logf("Comment, Creator Name: %v", gotResult.Author.DisplayName)
				t.Logf("Comment, Created Date: %v", gotResult.Created.Friendly)
				t.Log("----------------------------------")

			}

		})
	}

}

func TestRequestCommentService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		commentID          int
		expands            []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetCustomerCommentWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          101,
			expands:            []string{"attachment", "renderedBody"},
			mockFile:           "./mocks/get-comment-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/101?expand=attachment%2CrenderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerCommentWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			commentID:          101,
			expands:            []string{"attachment", "renderedBody"},
			mockFile:           "./mocks/get-comment-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/101?expand=attachment%2CrenderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentWhenTheExpandsAreNil",
			issueKeyOrID:       "DUMMY-3",
			commentID:          101,
			expands:            nil,
			mockFile:           "./mocks/get-comment-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/101",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerCommentWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          101,
			expands:            []string{"attachment", "renderedBody"},
			mockFile:           "./mocks/get-comment-request.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/101?expand=attachment%2CrenderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			commentID:          101,
			expands:            []string{"attachment", "renderedBody"},
			mockFile:           "./mocks/get-comment-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/101?expand=attachment%2CrenderedBody",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			commentID:          101,
			expands:            []string{"attachment", "renderedBody"},
			mockFile:           "./mocks/get-comment-request.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/101?expand=attachment%2CrenderedBody",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			commentID:          101,
			expands:            []string{"attachment", "renderedBody"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment/101?expand=attachment%2CrenderedBody",
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

			service := &RequestCommentService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.issueKeyOrID, testCase.commentID, testCase.expands)

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

				t.Log("----------------------------------")
				t.Logf("Comment, ID: %v", gotResult.ID)
				t.Logf("Comment, Creator Name: %v", gotResult.Author.DisplayName)
				t.Logf("Comment, Created Date: %v", gotResult.Created.Friendly)
				t.Log("----------------------------------")

			}

		})
	}

}

func TestRequestCommentService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		public             bool
		expands            []string
		start, limit       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetCustomerCommentsWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			public:             true,
			expands:            []string{"attachment", "renderedBody"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comments-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment?expand=attachment%2CrenderedBody&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerCommentsWhenThePublicParamIsNotTrue",
			issueKeyOrID:       "DUMMY-3",
			public:             false,
			expands:            []string{"attachment", "renderedBody"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comments-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment?expand=attachment%2CrenderedBody&limit=50&public=false&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetCustomerCommentsWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			public:             true,
			expands:            []string{"attachment", "renderedBody"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comments-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment?expand=attachment%2CrenderedBody&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentsWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			public:             true,
			expands:            []string{"attachment", "renderedBody"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comments-requests.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment?expand=attachment%2CrenderedBody&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentsWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			public:             true,
			expands:            []string{"attachment", "renderedBody"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comments-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment?expand=attachment%2CrenderedBody&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentsWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			public:             true,
			expands:            []string{"attachment", "renderedBody"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comments-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment?expand=attachment%2CrenderedBody&limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentsWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			public:             true,
			expands:            []string{"attachment", "renderedBody"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-comments-requests.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetCustomerCommentsWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			public:             true,
			expands:            []string{"attachment", "renderedBody"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/comment?expand=attachment%2CrenderedBody&limit=50&start=0",
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

			service := &RequestCommentService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.issueKeyOrID, testCase.public,
				testCase.expands, testCase.start, testCase.limit)

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

				for _, comment := range gotResult.Values {

					t.Log("----------------------------------")
					t.Logf("Comment, ID: %v", comment.ID)
					t.Logf("Comment, Creator Name: %v", comment.Author.DisplayName)
					t.Logf("Comment, Created Date: %v", comment.Created.Friendly)
					t.Log("----------------------------------")

				}

			}

		})
	}

}
