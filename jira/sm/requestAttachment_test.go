package sm

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestRequestAttachmentService_Create(t *testing.T) {

	testCases := []struct {
		name                   string
		issueKeyOrID           string
		temporaryAttachmentIDs []string
		public                 bool
		mockFile               string
		wantHTTPMethod         string
		endpoint               string
		context                context.Context
		wantHTTPCodeReturn     int
		wantErr                bool
	}{
		{
			name:                   "CreateRequestAttachmentWhenTheParametersAreCorrect",
			issueKeyOrID:           "DUMMY-3",
			temporaryAttachmentIDs: []string{"temp910441317820424274", "temp3600755449679003114"},
			public:                 true,
			mockFile:               "./mocks/create-attachment.json",
			wantHTTPMethod:         http.MethodPost,
			endpoint:               "/rest/servicedeskapi/request/DUMMY-3/attachment",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			wantErr:                false,
		},

		{
			name:                   "CreateRequestAttachmentWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:           "",
			temporaryAttachmentIDs: []string{"temp910441317820424274", "temp3600755449679003114"},
			public:                 true,
			mockFile:               "./mocks/create-attachment.json",
			wantHTTPMethod:         http.MethodPost,
			endpoint:               "/rest/servicedeskapi/request/DUMMY-3/attachment",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			wantErr:                true,
		},

		{
			name:                   "CreateRequestAttachmentWhenTheTemporaryAttachmentsIDsAreNotSet",
			issueKeyOrID:           "DUMMY-3",
			temporaryAttachmentIDs: nil,
			public:                 true,
			mockFile:               "./mocks/create-attachment.json",
			wantHTTPMethod:         http.MethodPost,
			endpoint:               "/rest/servicedeskapi/request/DUMMY-3/attachment",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			wantErr:                true,
		},

		{
			name:                   "CreateRequestAttachmentWhenTheAttachmentIsNotPublic",
			issueKeyOrID:           "DUMMY-3",
			temporaryAttachmentIDs: []string{"temp910441317820424274", "temp3600755449679003114"},
			public:                 false,
			mockFile:               "./mocks/create-attachment.json",
			wantHTTPMethod:         http.MethodPost,
			endpoint:               "/rest/servicedeskapi/request/DUMMY-3/attachment",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			wantErr:                false,
		},

		{
			name:                   "CreateRequestAttachmentWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:           "DUMMY-3",
			temporaryAttachmentIDs: []string{"temp910441317820424274", "temp3600755449679003114"},
			public:                 true,
			mockFile:               "./mocks/create-attachment.json",
			wantHTTPMethod:         http.MethodGet,
			endpoint:               "/rest/servicedeskapi/request/DUMMY-3/attachment",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			wantErr:                true,
		},

		{
			name:                   "CreateRequestAttachmentWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:           "DUMMY-3",
			temporaryAttachmentIDs: []string{"temp910441317820424274", "temp3600755449679003114"},
			public:                 true,
			mockFile:               "./mocks/create-attachment.json",
			wantHTTPMethod:         http.MethodPost,
			endpoint:               "/rest/servicedeskapi/request/DUMMY-3/attachment",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusBadRequest,
			wantErr:                true,
		},

		{
			name:                   "CreateRequestAttachmentWhenTheContextIsNil",
			issueKeyOrID:           "DUMMY-3",
			temporaryAttachmentIDs: []string{"temp910441317820424274", "temp3600755449679003114"},
			public:                 true,
			mockFile:               "./mocks/create-attachment.json",
			wantHTTPMethod:         http.MethodPost,
			endpoint:               "/rest/servicedeskapi/request/DUMMY-3/attachment",
			context:                nil,
			wantHTTPCodeReturn:     http.StatusOK,
			wantErr:                true,
		},

		{
			name:                   "CreateRequestAttachmentWhenTheEndpointIsEmpty",
			issueKeyOrID:           "DUMMY-3",
			temporaryAttachmentIDs: []string{"temp910441317820424274", "temp3600755449679003114"},
			public:                 true,
			mockFile:               "./mocks/create-attachment.json",
			wantHTTPMethod:         http.MethodPost,
			endpoint:               "",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			wantErr:                true,
		},

		{
			name:                   "CreateRequestAttachmentWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:           "DUMMY-3",
			temporaryAttachmentIDs: []string{"temp910441317820424274", "temp3600755449679003114"},
			public:                 true,
			mockFile:               "./mocks/empty_json.json",
			wantHTTPMethod:         http.MethodPost,
			endpoint:               "/rest/servicedeskapi/request/DUMMY-3/attachment",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			wantErr:                true,
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

			service := &RequestAttachmentService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.issueKeyOrID, testCase.temporaryAttachmentIDs, testCase.public)

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

				t.Log(gotResult)
			}

		})
	}

}

func TestRequestAttachmentService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyOrID       string
		start, limit       int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetRequestAttachmentsWhenTheParametersAreCorrect",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/attachment?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetRequestAttachmentsWhenTheIssueKeyOrIDIsNotSet",
			issueKeyOrID:       "",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/attachment?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestAttachmentsWhenTheRequestMethodIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-attachments.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/attachment?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestAttachmentsWhenTheStatusCodeIsIncorrect",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/attachment?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetRequestAttachmentsWhenTheContextIsNil",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/attachment?limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestAttachmentsWhenTheEndpointIsEmpty",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetRequestAttachmentsWhenTheResponseBodyHasADifferentFormat",
			issueKeyOrID:       "DUMMY-3",
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/servicedeskapi/request/DUMMY-3/attachment?limit=50&start=0",
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

			service := &RequestAttachmentService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.issueKeyOrID, testCase.start, testCase.limit)

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

				for _, file := range gotResult.Values {
					t.Log(file.Filename, file.Size, file.Author.EmailAddress, file.MimeType, file.Created.Iso8601)
				}
			}

		})
	}

}
