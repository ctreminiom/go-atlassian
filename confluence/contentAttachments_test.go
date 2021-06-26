package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestContentAttachmentService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		startAt, maxResults int
		options *GetContentAttachmentsOptionsScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetContentAttachmentsWhenTheParametersAreCorrect",
			contentID:          "5949392",
			startAt:            0,
			maxResults:         50,
			options:            &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentAttachmentsWhenTheContentIDIsNotProvided",
			contentID:          "",
			startAt:            0,
			maxResults:         50,
			options:            &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentAttachmentsWhenTheContextIsNotProvided",
			contentID:          "5949392",
			startAt:            0,
			maxResults:         50,
			options:            &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentAttachmentsWhenTheRequestMethodIsIncorrect",
			contentID:          "5949392",
			startAt:            0,
			maxResults:         50,
			options:            &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentAttachmentsWhenTheOptionsAreNotProvided",
			contentID:          "5949392",
			startAt:            0,
			maxResults:         50,
			options:            nil,
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentAttachmentsWhenTheStatusCodeIsIncorrect",
			contentID:          "5949392",
			startAt:            0,
			maxResults:         50,
			options:            &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/get-content-attachments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetContentAttachmentsWhenTheResponseBodyIsEmpty",
			contentID:          "5949392",
			startAt:            0,
			maxResults:         50,
			options:            &GetContentAttachmentsOptionsScheme{
				Expand:    []string{"operations", "file"},
				FileName:  "filename.png",
				MediaType: ".png",
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/5949392/child/attachment?expand=operations%2Cfile&filename=filename.png&limit=50&mediaType=.png&start=0",
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

			service := &ContentAttachmentService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.contentID, testCase.startAt, testCase.maxResults, testCase.options)

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
