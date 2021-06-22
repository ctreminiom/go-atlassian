package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestContentService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *ContentScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateContentWhenTheParametersAreCorrect",
			payload: &ContentScheme{
				Type:  "page", // Valid values: page, blogpost, comment
				Title: "Confluence Page Title",
				Space: &SpaceScheme{Key: "DUMMY"},
				Body: &BodyScheme{
					Storage: &BodyNodeScheme{
						Value:          "<p>This is <br/> a new page</p>",
						Representation: "storage",
					},
				},
			},
			mockFile:           "./mocks/create-content.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},


		{
			name: "CreateContentWhenThePayloadIsNotProvided",
			payload: nil,
			mockFile:           "./mocks/create-content.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateContentWhenTheRequestMethodIsIncorrect",
			payload: &ContentScheme{
				Type:  "page", // Valid values: page, blogpost, comment
				Title: "Confluence Page Title",
				Space: &SpaceScheme{Key: "DUMMY"},
				Body: &BodyScheme{
					Storage: &BodyNodeScheme{
						Value:          "<p>This is <br/> a new page</p>",
						Representation: "storage",
					},
				},
			},
			mockFile:           "./mocks/create-content.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateContentWhenTheContextIsNotProvided",
			payload: &ContentScheme{
				Type:  "page", // Valid values: page, blogpost, comment
				Title: "Confluence Page Title",
				Space: &SpaceScheme{Key: "DUMMY"},
				Body: &BodyScheme{
					Storage: &BodyNodeScheme{
						Value:          "<p>This is <br/> a new page</p>",
						Representation: "storage",
					},
				},
			},
			mockFile:           "./mocks/create-content.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateContentWhenTheRequestBodyIsEmpty",
			payload: &ContentScheme{
				Type:  "page", // Valid values: page, blogpost, comment
				Title: "Confluence Page Title",
				Space: &SpaceScheme{Key: "DUMMY"},
				Body: &BodyScheme{
					Storage: &BodyNodeScheme{
						Value:          "<p>This is <br/> a new page</p>",
						Representation: "storage",
					},
				},
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content",
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

			service := &ContentService{client: mockClient}

			gotResult, gotResponse, err := service.Create(testCase.context, testCase.payload)

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

func TestContentService_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		options             *GetContentOptionsScheme
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name: "GetsContentWhenTheAllOptionsAreProvided",
			options: &GetContentOptionsScheme{
				ContextType: "page",
				SpaceKey:    "DUMMY",
				Title:       "*page*",
				Trigger:     "trigger-sample",
				OrderBy:     "id",
				Status:      []string{"status-00", "status-01"},
				Expand:      []string{"all"},
				PostingDay:  time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC),
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content?expand=all&limit=50&orderby=id&postingDay=2019-11-17&spaceKey=DUMMY&start=0&status=status-00%2Cstatus-01&title=%2Apage%2A&trigger=trigger-sample&type=page",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name: "GetsContentWhenTheContextIsNotSet",
			options: &GetContentOptionsScheme{
				ContextType: "page",
				SpaceKey:    "DUMMY",
				Title:       "*page*",
				Trigger:     "trigger-sample",
				OrderBy:     "id",
				Status:      []string{"status-00", "status-01"},
				Expand:      []string{"all"},
				PostingDay:  time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC),
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content?expand=all&limit=50&orderby=id&postingDay=2019-11-17&spaceKey=DUMMY&start=0&status=status-00%2Cstatus-01&title=%2Apage%2A&trigger=trigger-sample&type=page",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetsContentWhenTheRequestMethodIsIncorrect",
			options: &GetContentOptionsScheme{
				ContextType: "page",
				SpaceKey:    "DUMMY",
				Title:       "*page*",
				Trigger:     "trigger-sample",
				OrderBy:     "id",
				Status:      []string{"status-00", "status-01"},
				Expand:      []string{"all"},
				PostingDay:  time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC),
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/content?expand=all&limit=50&orderby=id&postingDay=2019-11-17&spaceKey=DUMMY&start=0&status=status-00%2Cstatus-01&title=%2Apage%2A&trigger=trigger-sample&type=page",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetsContentWhenTheResponseStatusIsIncorrect",
			options: &GetContentOptionsScheme{
				ContextType: "page",
				SpaceKey:    "DUMMY",
				Title:       "*page*",
				Trigger:     "trigger-sample",
				OrderBy:     "id",
				Status:      []string{"status-00", "status-01"},
				Expand:      []string{"all"},
				PostingDay:  time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC),
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content?expand=all&limit=50&orderby=id&postingDay=2019-11-17&spaceKey=DUMMY&start=0&status=status-00%2Cstatus-01&title=%2Apage%2A&trigger=trigger-sample&type=page",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "GetsContentWhenTheResponseBodyIsEmpty",
			options: &GetContentOptionsScheme{
				ContextType: "page",
				SpaceKey:    "DUMMY",
				Title:       "*page*",
				Trigger:     "trigger-sample",
				OrderBy:     "id",
				Status:      []string{"status-00", "status-01"},
				Expand:      []string{"all"},
				PostingDay:  time.Date(2019, 11, 17, 20, 34, 58, 651387237, time.UTC),
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content?expand=all&limit=50&orderby=id&postingDay=2019-11-17&spaceKey=DUMMY&start=0&status=status-00%2Cstatus-01&title=%2Apage%2A&trigger=trigger-sample&type=page",
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

			service := &ContentService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(
				testCase.context,
				testCase.options,
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
			}

		})

	}

}
