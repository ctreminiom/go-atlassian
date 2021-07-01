package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestContentPermissionService_Check(t *testing.T) {

	testCases := []struct {
		name               string
		contentID string
		payload *CheckPermissionScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CheckContentPermissionsWhenTheParametersAreCorrect",
			contentID:          "847373747",
			payload:            &CheckPermissionScheme{
				Subject:   &PermissionSubjectScheme{
					Type:       "user",
					Identifier: "5b86be50b8e3cb5895860d6d",
				},
				Operation: "read",
			},
			mockFile:           "./mocks/get-content-permission-check.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/847373747/permission/check",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "CheckContentPermissionsWhenTheContentIDIsNotProvided",
			contentID:          "",
			payload:            &CheckPermissionScheme{
				Subject:   &PermissionSubjectScheme{
					Type:       "user",
					Identifier: "5b86be50b8e3cb5895860d6d",
				},
				Operation: "read",
			},
			mockFile:           "./mocks/get-content-permission-check.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/847373747/permission/check",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CheckContentPermissionsWhenThePayloadIsNotProvided",
			contentID:          "847373747",
			payload:            nil,
			mockFile:           "./mocks/get-content-permission-check.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/847373747/permission/check",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CheckContentPermissionsWhenTheRequestMethodIsIncorrect",
			contentID:          "847373747",
			payload:            &CheckPermissionScheme{
				Subject:   &PermissionSubjectScheme{
					Type:       "user",
					Identifier: "5b86be50b8e3cb5895860d6d",
				},
				Operation: "read",
			},
			mockFile:           "./mocks/get-content-permission-check.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/wiki/rest/api/content/847373747/permission/check",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CheckContentPermissionsWhenTheStatusCodeIsIncorrect",
			contentID:          "847373747",
			payload:            &CheckPermissionScheme{
				Subject:   &PermissionSubjectScheme{
					Type:       "user",
					Identifier: "5b86be50b8e3cb5895860d6d",
				},
				Operation: "read",
			},
			mockFile:           "./mocks/get-content-permission-check.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/847373747/permission/check",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "CheckContentPermissionsWhenTheContextIsNotProvided",
			contentID:          "847373747",
			payload:            &CheckPermissionScheme{
				Subject:   &PermissionSubjectScheme{
					Type:       "user",
					Identifier: "5b86be50b8e3cb5895860d6d",
				},
				Operation: "read",
			},
			mockFile:           "./mocks/get-content-permission-check.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/847373747/permission/check",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CheckContentPermissionsWhenTheResponseBodyIsEmpty",
			contentID:          "847373747",
			payload:            &CheckPermissionScheme{
				Subject:   &PermissionSubjectScheme{
					Type:       "user",
					Identifier: "5b86be50b8e3cb5895860d6d",
				},
				Operation: "read",
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/content/847373747/permission/check",
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

			service := &ContentPermissionService{client: mockClient}

			gotResult, gotResponse, err := service.Check(testCase.context, testCase.contentID, testCase.payload)

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
