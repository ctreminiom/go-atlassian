package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestIssueTypeScreenSchemeService_Append(t *testing.T) {

	testCases := []struct {
		name                    string
		issueTypeScreenSchemeID string
		payload                 *models.IssueTypeScreenSchemePayloadScheme
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheParametersAreCorrect",
			issueTypeScreenSchemeID: "10000",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "10000", // Epic Issue Type
						ScreenSchemeID: "10002",
					},
					{
						IssueTypeID:    "10002", // Task Issue Type
						ScreenSchemeID: "10002",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheIssueTypeScreenSchemeIDParamIsEmpty",
			issueTypeScreenSchemeID: "",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "10000", // Epic Issue Type
						ScreenSchemeID: "10002",
					},
					{
						IssueTypeID:    "10002", // Task Issue Type
						ScreenSchemeID: "10002",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheMappingsParamIsNil",
			issueTypeScreenSchemeID: "10000",
			payload:                 nil,
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheEndpointIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "10000", // Epic Issue Type
						ScreenSchemeID: "10002",
					},
					{
						IssueTypeID:    "10002", // Task Issue Type
						ScreenSchemeID: "10002",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenschemes/10000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "10000", // Epic Issue Type
						ScreenSchemeID: "10002",
					},
					{
						IssueTypeID:    "10002", // Task Issue Type
						ScreenSchemeID: "10002",
					},
				},
			},
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "10000", // Epic Issue Type
						ScreenSchemeID: "10002",
					},
					{
						IssueTypeID:    "10002", // Task Issue Type
						ScreenSchemeID: "10002",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:                    "AppendMappingsToIssueTypeScreenSchemeWhenTheContextIsNil",
			issueTypeScreenSchemeID: "10000",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "10000", // Epic Issue Type
						ScreenSchemeID: "10002",
					},
					{
						IssueTypeID:    "10002", // Task Issue Type
						ScreenSchemeID: "10002",
					},
				},
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenscheme/10000/mapping",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResponse, err := i.Append(testCase.context, testCase.issueTypeScreenSchemeID, testCase.payload)

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

func TestIssueTypeScreenSchemeService_Assign(t *testing.T) {

	testCases := []struct {
		name                    string
		issueTypeScreenSchemeID string
		projectID               string
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheParametersAreCorrect",
			issueTypeScreenSchemeID: "10000",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 false,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheProjectIDIsNotSet",
			issueTypeScreenSchemeID: "10000",
			projectID:               "",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheIssueTypeScreenSchemeIDIsNotSet",
			issueTypeScreenSchemeID: "",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheEndpointIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/projects",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusBadRequest,
			wantErr:                 true,
		},

		{
			name:                    "AssignIssueTypeSchemeToProjectWhenTheContextIsNil",
			issueTypeScreenSchemeID: "10000",
			projectID:               "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/project",
			context:                 nil,
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResponse, err := i.Assign(testCase.context, testCase.issueTypeScreenSchemeID, testCase.projectID)

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

func TestIssueTypeScreenSchemeService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.IssueTypeScreenSchemePayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateIssueTypeSchemeWhenTheParametersAreCorrect",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "default",
						ScreenSchemeID: "10000",
					},
					{
						IssueTypeID:    "10004", // Bug
						ScreenSchemeID: "10002",
					},
				},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheResponseBodyIsNotAStringValue",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "default",
						ScreenSchemeID: "10000",
					},
					{
						IssueTypeID:    "10004", // Bug
						ScreenSchemeID: "10002",
					},
				},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme-not-string.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateIssueTypeSchemeWhenTheIssueTypeScreenSchemePayloadSchemeParamIsNil",
			payload:            nil,
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheEndpointIsIncorrect",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "default",
						ScreenSchemeID: "10000",
					},
					{
						IssueTypeID:    "10004", // Bug
						ScreenSchemeID: "10002",
					},
				},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/apis/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheRequestMethodIsIncorrect",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "default",
						ScreenSchemeID: "10000",
					},
					{
						IssueTypeID:    "10004", // Bug
						ScreenSchemeID: "10002",
					},
				},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheStatusCodeIsIncorrect",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "default",
						ScreenSchemeID: "10000",
					},
					{
						IssueTypeID:    "10004", // Bug
						ScreenSchemeID: "10002",
					},
				},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheContextIsNil",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "default",
						ScreenSchemeID: "10000",
					},
					{
						IssueTypeID:    "10004", // Bug
						ScreenSchemeID: "10002",
					},
				},
			},
			mockFile:           "./mocks/create-issue-type-screen-scheme.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeSchemeWhenTheResponseBodyHasADifferentFormat",
			payload: &models.IssueTypeScreenSchemePayloadScheme{
				Name: "Scrum issue type screen scheme",
				IssueTypeMappings: []*models.IssueTypeScreenSchemeMappingPayloadScheme{
					{
						IssueTypeID:    "default",
						ScreenSchemeID: "10000",
					},
					{
						IssueTypeID:    "10004", // Bug
						ScreenSchemeID: "10002",
					},
				},
			},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme",
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.payload)

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

func TestIssueTypeScreenSchemeService_Delete(t *testing.T) {

	testCases := []struct {
		name                    string
		issueTypeScreenSchemeID string
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:                    "DeleteIssueTypeSchemeWhenTheIssueTypeScreenSchemeIDIsValid",
			issueTypeScreenSchemeID: "10000",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 false,
		},

		{
			name:                    "DeleteIssueTypeSchemeWhenTheIssueTypeScreenSchemeIDIsNotSet",
			issueTypeScreenSchemeID: "",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "DeleteIssueTypeSchemeWhenTheEndpointIsIncorrect",
			issueTypeScreenSchemeID: "10001",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "DeleteIssueTypeSchemeWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "DeleteIssueTypeSchemeWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusBadRequest,
			wantErr:                 true,
		},

		{
			name:                    "DeleteIssueTypeSchemeWhenTheContextIsNil",
			issueTypeScreenSchemeID: "10000",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000",
			context:                 nil,
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.issueTypeScreenSchemeID)

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

func TestIssueTypeScreenSchemeService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		ids                []int
		startAt            int
		maxResults         int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueTypeSchemesWhenTheParametersAreCorrect",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheIdsAreNotSet",
			ids:                nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheEndpointIsIncorrect",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenschemes?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheRequestMethodIsIncorrect",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheStatusCodeIsIncorrect",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheContextIsNil",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeSchemesWhenTheResponseBodyHasADifferentFormat",
			ids:                []int{1000, 1001, 1002},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme?id=1000&id=1001&id=1002&maxResults=50&startAt=0",
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.ids, testCase.startAt, testCase.maxResults)

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

func TestIssueTypeScreenSchemeService_Update(t *testing.T) {

	testCases := []struct {
		name                      string
		issueTypeScreenSchemeID   string
		issueTypeScreenSchemeName string
		description               string
		mockFile                  string
		wantHTTPMethod            string
		endpoint                  string
		context                   context.Context
		wantHTTPCodeReturn        int
		wantErr                   bool
	}{
		{
			name:                      "UpdateIssueTypeSchemeWhenTheParametersAreCorrect",
			issueTypeScreenSchemeID:   "10000",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPut,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10000",
			context:                   context.Background(),
			wantHTTPCodeReturn:        http.StatusNoContent,
			wantErr:                   false,
		},

		{
			name:                      "UpdateIssueTypeSchemeWhenTheIssueTypeScreenSchemeIDIsNotSet",
			issueTypeScreenSchemeID:   "",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPut,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10000",
			context:                   context.Background(),
			wantHTTPCodeReturn:        http.StatusNoContent,
			wantErr:                   true,
		},

		{
			name:                      "UpdateIssueTypeSchemeWhenTheEndpointIsIncorrect",
			issueTypeScreenSchemeID:   "10000",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPut,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10001",
			context:                   context.Background(),
			wantHTTPCodeReturn:        http.StatusNoContent,
			wantErr:                   true,
		},

		{
			name:                      "UpdateIssueTypeSchemeWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID:   "10000",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPost,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10000",
			context:                   context.Background(),
			wantHTTPCodeReturn:        http.StatusNoContent,
			wantErr:                   true,
		},

		{
			name:                      "UpdateIssueTypeSchemeWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID:   "10000",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPut,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10000",
			context:                   context.Background(),
			wantHTTPCodeReturn:        http.StatusBadRequest,
			wantErr:                   true,
		},

		{
			name:                      "UpdateIssueTypeSchemeWhenTheContextIsNil",
			issueTypeScreenSchemeID:   "10000",
			issueTypeScreenSchemeName: "New issue type scheme name",
			description:               "New issue type scheme description",
			wantHTTPMethod:            http.MethodPut,
			endpoint:                  "/rest/api/3/issuetypescreenscheme/10000",
			context:                   nil,
			wantHTTPCodeReturn:        http.StatusNoContent,
			wantErr:                   true,
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResponse, err := i.Update(testCase.context, testCase.issueTypeScreenSchemeID,
				testCase.issueTypeScreenSchemeName, testCase.description)

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

func TestIssueTypeScreenSchemeService_UpdateDefault(t *testing.T) {

	testCases := []struct {
		name                    string
		issueTypeScreenSchemeID string
		screenSchemeID          string
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:                    "UpdateIssueTypeScreenSchemeDefaultScreenWhenTheParametersAreCorrect",
			issueTypeScreenSchemeID: "10000",
			screenSchemeID:          "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/default",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 false,
		},

		{
			name:                    "UpdateIssueTypeScreenSchemeDefaultScreenWhenTheIssueTypeScreenSchemeIDIsNotSet",
			issueTypeScreenSchemeID: "",
			screenSchemeID:          "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/default",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "UpdateIssueTypeScreenSchemeDefaultScreenWhenTheScreenSchemeIDIsNotSet",
			issueTypeScreenSchemeID: "10000",
			screenSchemeID:          "",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/default",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "UpdateIssueTypeScreenSchemeDefaultScreenWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			screenSchemeID:          "10001",
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/default",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "UpdateIssueTypeScreenSchemeDefaultScreenWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			screenSchemeID:          "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/default",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusBadRequest,
			wantErr:                 true,
		},

		{
			name:                    "UpdateIssueTypeScreenSchemeDefaultScreenWhenTheContextIsNil",
			issueTypeScreenSchemeID: "10000",
			screenSchemeID:          "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/default",
			context:                 nil,
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "UpdateIssueTypeScreenSchemeDefaultScreenWhenTheEndpointIsEmpty",
			issueTypeScreenSchemeID: "10000",
			screenSchemeID:          "10001",
			wantHTTPMethod:          http.MethodPut,
			endpoint:                "",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResponse, err := i.UpdateDefault(testCase.context, testCase.issueTypeScreenSchemeID, testCase.screenSchemeID)

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

func TestIssueTypeScreenSchemeService_Remove(t *testing.T) {

	testCases := []struct {
		name                    string
		issueTypeScreenSchemeID string
		issueTypeIDs            []string
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:                    "RemoveMappingsFromIssueTypeScreenSchemeWhenTheParametersAreCorrect",
			issueTypeScreenSchemeID: "10000",
			issueTypeIDs:            []string{"10001", "10002"},
			wantHTTPMethod:          http.MethodPost,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/remove",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 false,
		},

		{
			name:                    "RemoveMappingsFromIssueTypeScreenSchemeWhenTheIssueTypeScreenSchemeIDIsNotSet",
			issueTypeScreenSchemeID: "",
			issueTypeIDs:            []string{"10001", "10002"},
			wantHTTPMethod:          http.MethodPost,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/remove",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "RemoveMappingsFromIssueTypeScreenSchemeWhenTheIssueTypeIDsAreNotSet",
			issueTypeScreenSchemeID: "10000",
			issueTypeIDs:            nil,
			wantHTTPMethod:          http.MethodPost,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/remove",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "RemoveMappingsFromIssueTypeScreenSchemeWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			issueTypeIDs:            []string{"10001", "10002"},
			wantHTTPMethod:          http.MethodDelete,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/remove",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "RemoveMappingsFromIssueTypeScreenSchemeWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID: "10000",
			issueTypeIDs:            []string{"10001", "10002"},
			wantHTTPMethod:          http.MethodPost,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/remove",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusBadRequest,
			wantErr:                 true,
		},

		{
			name:                    "RemoveMappingsFromIssueTypeScreenSchemeWhenTheContextIsNil",
			issueTypeScreenSchemeID: "10000",
			issueTypeIDs:            []string{"10001", "10002"},
			wantHTTPMethod:          http.MethodPost,
			endpoint:                "/rest/api/3/issuetypescreenscheme/10000/mapping/remove",
			context:                 nil,
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
		},

		{
			name:                    "RemoveMappingsFromIssueTypeScreenSchemeWhenTheEndpointIsEmpty",
			issueTypeScreenSchemeID: "10000",
			issueTypeIDs:            []string{"10001", "10002"},
			wantHTTPMethod:          http.MethodPost,
			endpoint:                "",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusNoContent,
			wantErr:                 true,
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResponse, err := i.Remove(testCase.context, testCase.issueTypeScreenSchemeID, testCase.issueTypeIDs)

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

func TestIssueTypeScreenSchemeService_Mapping(t *testing.T) {

	testCases := []struct {
		name                     string
		issueTypeScreenSchemeIDs []int
		startAt, maxResults      int
		mockFile                 string
		wantHTTPMethod           string
		endpoint                 string
		context                  context.Context
		wantHTTPCodeReturn       int
		wantErr                  bool
	}{
		{
			name:                     "GetIssueTypeScreenSchemeMappingsWhenTheParametersAreCorrect",
			issueTypeScreenSchemeIDs: []int{1000, 1001},
			startAt:                  0,
			maxResults:               50,
			mockFile:                 "./mocks/get-issue-type-screen-scheme-items.json",
			wantHTTPMethod:           http.MethodGet,
			endpoint:                 "/rest/api/3/issuetypescreenscheme/mapping?issueTypeScreenSchemeId=1000&issueTypeScreenSchemeId=1001&maxResults=50&startAt=0",
			context:                  context.Background(),
			wantHTTPCodeReturn:       http.StatusOK,
			wantErr:                  false,
		},

		{
			name:                     "GetIssueTypeScreenSchemeMappingsWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeIDs: []int{1000, 1001},
			startAt:                  0,
			maxResults:               50,
			mockFile:                 "./mocks/get-issue-type-screen-scheme-items.json",
			wantHTTPMethod:           http.MethodPost,
			endpoint:                 "/rest/api/3/issuetypescreenscheme/mapping?issueTypeScreenSchemeId=1000&issueTypeScreenSchemeId=1001&maxResults=50&startAt=0",
			context:                  context.Background(),
			wantHTTPCodeReturn:       http.StatusOK,
			wantErr:                  true,
		},

		{
			name:                     "GetIssueTypeScreenSchemeMappingsWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeIDs: []int{1000, 1001},
			startAt:                  0,
			maxResults:               50,
			mockFile:                 "./mocks/get-issue-type-screen-scheme-items.json",
			wantHTTPMethod:           http.MethodGet,
			endpoint:                 "/rest/api/3/issuetypescreenscheme/mapping?issueTypeScreenSchemeId=1000&issueTypeScreenSchemeId=1001&maxResults=50&startAt=0",
			context:                  context.Background(),
			wantHTTPCodeReturn:       http.StatusBadRequest,
			wantErr:                  true,
		},

		{
			name:                     "GetIssueTypeScreenSchemeMappingsWhenTheContextIsNil",
			issueTypeScreenSchemeIDs: []int{1000, 1001},
			startAt:                  0,
			maxResults:               50,
			mockFile:                 "./mocks/get-issue-type-screen-scheme-items.json",
			wantHTTPMethod:           http.MethodGet,
			endpoint:                 "/rest/api/3/issuetypescreenscheme/mapping?issueTypeScreenSchemeId=1000&issueTypeScreenSchemeId=1001&maxResults=50&startAt=0",
			context:                  nil,
			wantHTTPCodeReturn:       http.StatusOK,
			wantErr:                  true,
		},

		{
			name:                     "GetIssueTypeScreenSchemeMappingsWhenTheEndpointIsIncorrect",
			issueTypeScreenSchemeIDs: []int{1000, 1001},
			startAt:                  0,
			maxResults:               50,
			mockFile:                 "./mocks/get-issue-type-screen-scheme-items.json",
			wantHTTPMethod:           http.MethodGet,
			endpoint:                 "/rest/chemeId=1000&issueTypeScreenSchemeId=1001&maxResults=50&startAt=0",
			context:                  context.Background(),
			wantHTTPCodeReturn:       http.StatusOK,
			wantErr:                  true,
		},

		{
			name:                     "GetIssueTypeScreenSchemeMappingsWhenTheResponseBodyHasADifferentFormat",
			issueTypeScreenSchemeIDs: []int{1000, 1001},
			startAt:                  0,
			maxResults:               50,
			mockFile:                 "./mocks/empty_json.json",
			wantHTTPMethod:           http.MethodGet,
			endpoint:                 "/rest/api/3/issuetypescreenscheme/mapping?issueTypeScreenSchemeId=1000&issueTypeScreenSchemeId=1001&maxResults=50&startAt=0",
			context:                  context.Background(),
			wantHTTPCodeReturn:       http.StatusOK,
			wantErr:                  true,
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Mapping(testCase.context, testCase.issueTypeScreenSchemeIDs, testCase.startAt, testCase.maxResults)

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

				//issue type screen scheme items
				for _, item := range gotResult.Values {

					t.Log("------------------------------")
					t.Logf("Screen Scheme Item IssueTypeID: %v", item.IssueTypeID)
					t.Logf("Screen Scheme Item ScreenSchemeID: %v", item.ScreenSchemeID)
					t.Logf("Screen Scheme Item IssueTypeScreenSchemeID: %v", item.IssueTypeScreenSchemeID)
					t.Log("------------------------------")

				}

			}
		})

	}

}

func TestIssueTypeScreenSchemeService_Projects(t *testing.T) {

	testCases := []struct {
		name                string
		projectIDs          []int
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetIssueTypeScreenSchemeProjectsWhenTheParametersAreCorrect",
			projectIDs:         []int{1000, 1001},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes-for-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme/project?maxResults=50&projectId=1000&projectId=1001&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTypeScreenSchemeProjectsWhenTheProjectIDsAreNotSet",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes-for-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme/project?maxResults=50&projectId=1000&projectId=1001&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeScreenSchemeProjectsWhenTheRequestMethodIsIncorrect",
			projectIDs:         []int{1000, 1001},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes-for-projects.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issuetypescreenscheme/project?maxResults=50&projectId=1000&projectId=1001&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeScreenSchemeProjectsWhenTheStatusCodeIsIncorrect",
			projectIDs:         []int{1000, 1001},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes-for-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme/project?maxResults=50&projectId=1000&projectId=1001&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeScreenSchemeProjectsWhenTheContextIsNil",
			projectIDs:         []int{1000, 1001},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes-for-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme/project?maxResults=50&projectId=1000&projectId=1001&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeScreenSchemeProjectsWhenTheEndpointIsIncorrect",
			projectIDs:         []int{1000, 1001},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-issue-type-screen-schemes-for-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/chemeId=1000&issueTypeScreenSchemeId=1001&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeScreenSchemeProjectsWhenTheResponseBodyHasADifferentFormat",
			projectIDs:         []int{1000, 1001},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issuetypescreenscheme/project?maxResults=50&projectId=1000&projectId=1001&startAt=0",
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.Projects(testCase.context, testCase.projectIDs, testCase.startAt, testCase.maxResults)

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

				//issue type screen scheme items
				for _, item := range gotResult.Values {

					t.Log("------------------------------")
					t.Log(item.IssueTypeScreenScheme.ID, item.IssueTypeScreenScheme.Name, item.IssueTypeScreenScheme.Description)
					t.Log(item.ProjectIds)
					t.Log("------------------------------")
				}

			}
		})

	}

}

func TestIssueTypeScreenSchemeService_SchemesByProject(t *testing.T) {

	testCases := []struct {
		name                    string
		issueTypeScreenSchemeID int
		startAt                 int
		maxResults              int
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:                    "GetSchemesByProjectWhenTheParametersAreCorrect",
			issueTypeScreenSchemeID: 1000,
			startAt:                 0,
			maxResults:              50,
			mockFile:                "./mocks/get-issue-type-screen-schemes-by-project.json",
			wantHTTPMethod:          http.MethodGet,
			endpoint:                "/rest/api/3/issuetypescreenscheme/1000/project?maxResults=50&startAt=0",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusOK,
			wantErr:                 false,
		},

		{
			name:                    "GetSchemesByProjectWhenTheRequestMethodIsIncorrect",
			issueTypeScreenSchemeID: 1000,
			startAt:                 0,
			maxResults:              50,
			mockFile:                "./mocks/get-issue-type-screen-schemes-by-project.json",
			wantHTTPMethod:          http.MethodHead,
			endpoint:                "/rest/api/3/issuetypescreenscheme/1000/project?maxResults=50&startAt=0",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusOK,
			wantErr:                 true,
		},

		{
			name:                    "GetSchemesByProjectWhenTheStatusCodeIsIncorrect",
			issueTypeScreenSchemeID: 1000,
			startAt:                 0,
			maxResults:              50,
			mockFile:                "./mocks/get-issue-type-screen-schemes-by-project.json",
			wantHTTPMethod:          http.MethodGet,
			endpoint:                "/rest/api/3/issuetypescreenscheme/1000/project?maxResults=50&startAt=0",
			context:                 context.Background(),
			wantHTTPCodeReturn:      http.StatusBadRequest,
			wantErr:                 true,
		},

		{
			name:                    "GetSchemesByProjectWhenTheContextIsNotProvided",
			issueTypeScreenSchemeID: 1000,
			startAt:                 0,
			maxResults:              50,
			mockFile:                "./mocks/get-issue-type-screen-schemes-by-project.json",
			wantHTTPMethod:          http.MethodGet,
			endpoint:                "/rest/api/3/issuetypescreenscheme/1000/project?maxResults=50&startAt=0",
			context:                 nil,
			wantHTTPCodeReturn:      http.StatusOK,
			wantErr:                 true,
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

			i := &IssueTypeScreenSchemeService{client: mockClient}

			gotResult, gotResponse, err := i.SchemesByProject(testCase.context, testCase.issueTypeScreenSchemeID, testCase.startAt, testCase.maxResults)

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
