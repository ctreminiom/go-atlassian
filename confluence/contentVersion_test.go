package confluence

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func Test_Content_Version_Service_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		expand             []string
		start              int
		limit              int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the parameters are correct",
			contentID:          "233838383",
			expand:             []string{"restrictions.user", "restrictions.group"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-content-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			expand:             []string{"restrictions.user", "restrictions.group"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-content-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			expand:             []string{"restrictions.user", "restrictions.group"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-content-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is not correct",
			contentID:          "233838383",
			expand:             []string{"restrictions.user", "restrictions.group"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/get-content-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the response body is not empty",
			contentID:          "233838383",
			expand:             []string{"restrictions.user", "restrictions.group"},
			start:              0,
			limit:              50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=restrictions.user%2Crestrictions.group&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
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

			service := &ContentVersionService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.contentID, testCase.expand,
				testCase.start, testCase.limit)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

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
			}
		})
	}
}

func Test_Content_Version_Service_Get(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		expand             []string
		versionNumber      int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the parameters are correct",
			contentID:          "233838383",
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/get-content-version.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/version/0?expand=collaborators%2Ccontent",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/get-content-version.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/version/0?expand=collaborators%2Ccontent",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/get-content-version.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/version/0?expand=collaborators%2Ccontent",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is not correct",
			contentID:          "233838383",
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/get-content-version.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/version/0?expand=collaborators%2Ccontent",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:               "when the response body is not empty",
			contentID:          "233838383",
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/content/233838383/version/0?expand=collaborators%2Ccontent",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
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

			service := &ContentVersionService{client: mockClient}

			gotResult, gotResponse, err := service.Get(testCase.context, testCase.contentID, testCase.versionNumber, testCase.expand)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

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
			}
		})
	}
}

func Test_Content_Version_Service_Restore(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		payload            *models.ContentRestorePayloadScheme
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:      "when the parameters are correct",
			contentID: "233838383",
			payload: &models.ContentRestorePayloadScheme{
				OperationKey: "restore",
				Params: &models.ContentRestoreParamsPayloadScheme{
					VersionNumber: 034,
					Message:       "message sample :)",
					RestoreTitle:  true,
				},
			},
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/get-content-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=collaborators%2Ccontent",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
		},

		{
			name:               "when the paylod is not provided",
			contentID:          "233838383",
			payload:            nil,
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/get-content-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=collaborators%2Ccontent",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "failed to parse the interface pointer, please provide a valid one",
		},

		{
			name:      "when the content id is not provided",
			contentID: "",
			payload: &models.ContentRestorePayloadScheme{
				OperationKey: "restore",
				Params: &models.ContentRestoreParamsPayloadScheme{
					VersionNumber: 034,
					Message:       "message sample :)",
					RestoreTitle:  true,
				},
			},
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/get-content-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=collaborators%2Ccontent",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:      "when the context is not provided",
			contentID: "233838383",
			payload: &models.ContentRestorePayloadScheme{
				OperationKey: "restore",
				Params: &models.ContentRestoreParamsPayloadScheme{
					VersionNumber: 034,
					Message:       "message sample :)",
					RestoreTitle:  true,
				},
			},
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/get-content-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=collaborators%2Ccontent",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:      "when the response status is not correct",
			contentID: "233838383",
			payload: &models.ContentRestorePayloadScheme{
				OperationKey: "restore",
				Params: &models.ContentRestoreParamsPayloadScheme{
					VersionNumber: 034,
					Message:       "message sample :)",
					RestoreTitle:  true,
				},
			},
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/get-content-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=collaborators%2Ccontent",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 400",
		},

		{
			name:      "when the response body is not empty",
			contentID: "233838383",
			payload: &models.ContentRestorePayloadScheme{
				OperationKey: "restore",
				Params: &models.ContentRestoreParamsPayloadScheme{
					VersionNumber: 034,
					Message:       "message sample :)",
					RestoreTitle:  true,
				},
			},
			expand:             []string{"collaborators", "content"},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/233838383/version?expand=collaborators%2Ccontent",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "unexpected end of JSON input",
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

			service := &ContentVersionService{client: mockClient}

			gotResult, gotResponse, err := service.Restore(testCase.context, testCase.contentID, testCase.payload, testCase.expand)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

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
			}
		})
	}
}

func Test_Content_Version_Service_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		contentID          string
		versionNumber      int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:               "when the parameters are correct",
			contentID:          "233838383",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/version/0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
		},

		{
			name:               "when the content id is not provided",
			contentID:          "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/version/0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "confluence: no content id set",
		},

		{
			name:               "when the context is not provided",
			contentID:          "233838383",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/version/0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response status is not correct",
			contentID:          "233838383",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/content/233838383/version/0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusInternalServerError,
			wantErr:            true,
			expectedError:      "request failed. Please analyze the request body for more details. Status Code: 500",
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

			service := &ContentVersionService{client: mockClient}

			gotResponse, err := service.Delete(testCase.context, testCase.contentID, testCase.versionNumber)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.expectedError)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				}

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}
		})
	}
}
