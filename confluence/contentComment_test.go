package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestContentCommentService_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		contentID string
		expand, location []string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetContentCommentsWhenTheParametersAreCorrect",
			contentID:          "737377347",
			expand:             []string{"childtypes.all"},
			location:           []string{"location-00", "location-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-comments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/737377347/child/comment?expand=childtypes.all&limit=50&location=location-00%2Clocation-01&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentCommentsWhenTheContentIDIsNotProvided",
			contentID:          "",
			expand:             []string{"childtypes.all"},
			location:           []string{"location-00", "location-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-comments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/737377347/child/comment?expand=childtypes.all&limit=50&location=location-00%2Clocation-01&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentCommentsWhenTheContextIsNotProvided",
			contentID:          "737377347",
			expand:             []string{"childtypes.all"},
			location:           []string{"location-00", "location-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-comments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/737377347/child/comment?expand=childtypes.all&limit=50&location=location-00%2Clocation-01&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentCommentsWhenTheRequestMethodIsIncorrect",
			contentID:          "737377347",
			expand:             []string{"childtypes.all"},
			location:           []string{"location-00", "location-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-comments.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/wiki/rest/api/content/737377347/child/comment?expand=childtypes.all&limit=50&location=location-00%2Clocation-01&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentCommentsWhenTheStatusCodeIsIncorrect",
			contentID:          "737377347",
			expand:             []string{"childtypes.all"},
			location:           []string{"location-00", "location-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-comments.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/737377347/child/comment?expand=childtypes.all&limit=50&location=location-00%2Clocation-01&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetContentCommentsWhenTheResponseBodyIsEmpty",
			contentID:          "737377347",
			expand:             []string{"childtypes.all"},
			location:           []string{"location-00", "location-01"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/737377347/child/comment?expand=childtypes.all&limit=50&location=location-00%2Clocation-01&start=0",
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

			service := &ContentCommentService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(
				testCase.context, testCase.contentID, testCase.expand, testCase.location, testCase.startAt, testCase.maxResults)

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
