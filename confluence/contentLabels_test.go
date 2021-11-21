package confluence

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestContentLabelService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		payload            []*model.ContentLabelPayloadScheme
		want400Response    bool
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:      "AddContentLabelWhenTheParametersAreCorrect",
			contentID: "48483737372",
			payload: []*model.ContentLabelPayloadScheme{
				{
					Prefix: "global",
					Name:   "label-02",
				},
			},
			want400Response:    true,
			mockFile:           "./mocks/get-content-labels.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/48483737372/label?use-400-error-response=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:      "AddContentLabelWhenTheContentIDIsNotProvided",
			contentID: "",
			payload: []*model.ContentLabelPayloadScheme{
				{
					Prefix: "global",
					Name:   "label-02",
				},
			},
			want400Response:    true,
			mockFile:           "./mocks/get-content-labels.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/48483737372/label?use-400-error-response=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "AddContentLabelWhenTheContextIsNotProvided",
			contentID: "48483737372",
			payload: []*model.ContentLabelPayloadScheme{
				{
					Prefix: "global",
					Name:   "label-02",
				},
			},
			want400Response:    true,
			mockFile:           "./mocks/get-content-labels.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/48483737372/label?use-400-error-response=true",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddContentLabelWhenThePayloadIsNotProvided",
			contentID:          "48483737372",
			payload:            nil,
			want400Response:    true,
			mockFile:           "./mocks/get-content-labels.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/48483737372/label?use-400-error-response=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:      "AddContentLabelWhenTheStatusCodeIsIncorrect",
			contentID: "48483737372",
			payload: []*model.ContentLabelPayloadScheme{
				{
					Prefix: "global",
					Name:   "label-02",
				},
			},
			want400Response:    true,
			mockFile:           "./mocks/get-content-labels.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/48483737372/label?use-400-error-response=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:      "AddContentLabelWhenTheResponseBodyIsEmpty",
			contentID: "48483737372",
			payload: []*model.ContentLabelPayloadScheme{
				{
					Prefix: "global",
					Name:   "label-02",
				},
			},
			want400Response:    true,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/48483737372/label?use-400-error-response=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
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

			service := &ContentLabelService{client: mockClient}

			gotResult, gotResponse, err := service.Add(testCase.context, testCase.contentID, testCase.payload, testCase.want400Response)

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

func TestContentLabelService_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		contentID           string
		prefix              string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetContentLabelsWhenTheParametersAreCorrect",
			contentID:          "3483373",
			prefix:             "global",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-labels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3483373/label?limit=50&prefix=global&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentLabelsWhenTheParametersAreCorrect",
			contentID:          "3483373",
			prefix:             "global",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-labels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3483373/label?limit=50&prefix=global&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetContentLabelsWhenTheContextIsNotProvided",
			contentID:          "",
			prefix:             "global",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-labels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3483373/label?limit=50&prefix=global&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentLabelsWhenTheContextIsNotProvided",
			contentID:          "3483373",
			prefix:             "global",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-labels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3483373/label?limit=50&prefix=global&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetContentLabelsWhenTheStatusCodeIsIncorrect",
			contentID:          "3483373",
			prefix:             "global",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-labels.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3483373/label?limit=50&prefix=global&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetContentLabelsWhenTheResponseBodyIsEmpty",
			contentID:          "3483373",
			prefix:             "global",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/3483373/label?limit=50&prefix=global&start=0",
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

			service := &ContentLabelService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.contentID, testCase.prefix, testCase.startAt, testCase.maxResults)

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

func TestContentLabelService_Remove(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		labelName          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "RemoveContentLabelsWhenTheParametersAreCorrect",
			contentID:          "3483373",
			labelName:          "global",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/3483373/label/global",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "RemoveContentLabelsWhenTheContextIDIsNotProvided",
			contentID:          "",
			labelName:          "global",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/3483373/label/global",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveContentLabelsWhenTheContextIsNotProvided",
			contentID:          "3483373",
			labelName:          "global",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/3483373/label/global",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveContentLabelsWhenTheRequestMethodIsIncorrect",
			contentID:          "3483373",
			labelName:          "global",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content/3483373/label/global",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveContentLabelsWhenTheLabelNameIsNotProvided",
			contentID:          "3483373",
			labelName:          "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/3483373/label/global",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "RemoveContentLabelsWhenTheStatusCodeIsIncorrect",
			contentID:          "3483373",
			labelName:          "global",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/3483373/label/global",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
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

			service := &ContentLabelService{client: mockClient}

			gotResponse, err := service.Remove(testCase.context, testCase.contentID, testCase.labelName)

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
