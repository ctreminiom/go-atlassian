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

func TestSpacePermissionService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		spaceKey           string
		payload            *models.SpacePermissionPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:     "when the parameters are correct",
			spaceKey: "DUMMY",
			payload: &models.SpacePermissionPayloadScheme{
				Subject: &models.PermissionSubjectScheme{
					Type:       "user",
					Identifier: "account-id-sample",
				},
				Operation: &models.SpacePermissionOperationScheme{
					Operation: "administer",
					Target:    "page",
				},
			},
			mockFile:           "./mocks/space-permission-v2.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:     "when space key is not provided",
			spaceKey: "",
			payload: &models.SpacePermissionPayloadScheme{
				Subject: &models.PermissionSubjectScheme{
					Type:       "user",
					Identifier: "account-id-sample",
				},
				Operation: &models.SpacePermissionOperationScheme{
					Operation: "administer",
					Target:    "page",
				},
			},
			mockFile:           "./mocks/space-permission-v2.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no space key set",
		},

		{
			name:               "when the payload is not provided",
			spaceKey:           "DUMMY",
			payload:            nil,
			mockFile:           "./mocks/space-permission-v2.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "failed to parse the interface pointer, please provide a valid one",
		},

		{
			name:     "when the context is not provided",
			spaceKey: "DUMMY",
			payload: &models.SpacePermissionPayloadScheme{
				Subject: &models.PermissionSubjectScheme{
					Type:       "user",
					Identifier: "account-id-sample",
				},
				Operation: &models.SpacePermissionOperationScheme{
					Operation: "administer",
					Target:    "page",
				},
			},
			mockFile:           "./mocks/space-permission-v2.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:     "when response body is empty",
			spaceKey: "DUMMY",
			payload: &models.SpacePermissionPayloadScheme{
				Subject: &models.PermissionSubjectScheme{
					Type:       "user",
					Identifier: "account-id-sample",
				},
				Operation: &models.SpacePermissionOperationScheme{
					Operation: "administer",
					Target:    "page",
				},
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission",
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

			implementation := &SpacePermissionService{client: mockClient}

			gotResult, gotResponse, err := implementation.Add(testCase.context, testCase.spaceKey, testCase.payload)

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
			}
		})
	}
}

func TestSpacePermissionService_Bulk(t *testing.T) {

	testCases := []struct {
		name               string
		spaceKey           string
		payload            *models.SpacePermissionArrayPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
		expectedError      string
	}{
		{
			name:     "when the parameters are correct",
			spaceKey: "DUMMY",
			payload: &models.SpacePermissionArrayPayloadScheme{
				Subject: &models.PermissionSubjectScheme{
					Type:       "user",
					Identifier: "account-id-sample",
				},
				Operations: []*models.SpaceOperationPayloadScheme{
					{
						Key:    "read",
						Target: "space",
						Access: true,
					},
					{
						Key:    "delete",
						Target: "space",
						Access: false,
					},
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission/custom-content",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:     "when space key is not provided",
			spaceKey: "",
			payload: &models.SpacePermissionArrayPayloadScheme{
				Subject: &models.PermissionSubjectScheme{
					Type:       "user",
					Identifier: "account-id-sample",
				},
				Operations: []*models.SpaceOperationPayloadScheme{
					{
						Key:    "read",
						Target: "space",
						Access: true,
					},
					{
						Key:    "delete",
						Target: "space",
						Access: false,
					},
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission/custom-content",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no space key set",
		},

		{
			name:               "when the payload is not provided",
			spaceKey:           "DUMMY",
			payload:            nil,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission/custom-content",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "failed to parse the interface pointer, please provide a valid one",
		},

		{
			name:     "when the context is not provided",
			spaceKey: "DUMMY",
			payload: &models.SpacePermissionArrayPayloadScheme{
				Subject: &models.PermissionSubjectScheme{
					Type:       "user",
					Identifier: "account-id-sample",
				},
				Operations: []*models.SpaceOperationPayloadScheme{
					{
						Key:    "read",
						Target: "space",
						Access: true,
					},
					{
						Key:    "delete",
						Target: "space",
						Access: false,
					},
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission/custom-content",
			wantHTTPCodeReturn: http.StatusOK,
			context:            nil,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:     "when the response code is invalid",
			spaceKey: "DUMMY",
			payload: &models.SpacePermissionArrayPayloadScheme{
				Subject: &models.PermissionSubjectScheme{
					Type:       "user",
					Identifier: "account-id-sample",
				},
				Operations: []*models.SpaceOperationPayloadScheme{
					{
						Key:    "read",
						Target: "space",
						Access: true,
					},
					{
						Key:    "delete",
						Target: "space",
						Access: false,
					},
				},
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission/custom-content",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
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

			implementation := &SpacePermissionService{client: mockClient}

			gotResponse, err := implementation.Bulk(testCase.context, testCase.spaceKey, testCase.payload)

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
			}
		})
	}
}

func TestSpacePermissionService_Remove(t *testing.T) {

	testCases := []struct {
		name               string
		spaceKey           string
		permissionId       int
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
			spaceKey:           "DUMMY",
			permissionId:       100001,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission/100001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "when space key is not provided",
			spaceKey:           "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission/100001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
			expectedError:      "confluence: no space key set",
		},

		{
			name:               "when the context is not provided",
			spaceKey:           "DUMMY",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission/100001",
			wantHTTPCodeReturn: http.StatusOK,
			context:            nil,
			wantErr:            true,
			expectedError:      "request creation failed: net/http: nil Context",
		},

		{
			name:               "when the response code is invalid",
			spaceKey:           "DUMMY",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/space/DUMMY/permission/100001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
			expectedError:      "invalid character 'R' looking for beginning of value",
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

			implementation := &SpacePermissionService{client: mockClient}

			gotResponse, err := implementation.Remove(testCase.context, testCase.spaceKey, testCase.permissionId)

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
			}
		})
	}
}
