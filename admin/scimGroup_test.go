package admin

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestSCIMGroupService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		filter             string
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
			name:               "GetsSCIMGroupWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			filter:             "jira",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/scim-get-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups?count=50&filter=jira&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "GetsSCIMGroupWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			filter:             "jira",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/scim-get-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups?count=50&filter=jira&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "GetsSCIMGroupWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			filter:             "jira",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/scim-get-groups.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups?count=50&filter=jira&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "GetsSCIMGroupWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			filter:             "jira",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/scim-get-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups?count=50&filter=jira&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: 400,
			wantErr:            true,
		},
		{
			name:               "GetsSCIMGroupWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4",
			filter:             "jira",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/scim-get-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups?count=50&filter=jira&startIndex=0",
			context:            nil,
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "GetsSCIMGroupWhenTheFilterIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			filter:             "",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/scim-get-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups?count=50&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "GetsSCIMGroupWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			filter:             "jira",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/scim-get-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "GetsSCIMGroupWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			filter:             "jira",
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups?count=50&filter=jira&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
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

			service := &SCIMGroupService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.directoryID,
				testCase.filter, testCase.startAt, testCase.maxResults)

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

				for _, group := range gotResult.Resources {
					t.Log(group.DisplayName, group.ID, len(group.Members))
				}

			}

		})
	}

}

func TestSCIMGroupService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		groupID            string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetSCIMGroupWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "GetSCIMGroupWhenTheGroupIDIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "GetSCIMGroupWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			groupID:            "4475-a0c4",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "GetSCIMGroupWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "GetSCIMGroupWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 400,
			wantErr:            true,
		},
		{
			name:               "GetSCIMGroupWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            nil,
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "GetSCIMGroupWhenTheFilterIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "GetSCIMGroupWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "GetSCIMGroupWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
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

			service := &SCIMGroupService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.directoryID,
				testCase.groupID)

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

				t.Log(gotResult.DisplayName, gotResult.ID, len(gotResult.Members))

			}

		})
	}

}

func TestSCIMGroupService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		groupID            string
		newGroupName       string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "UpdateSCIMGroupWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			newGroupName:       "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "UpdateSCIMGroupWhenTheNewGroupNameIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			newGroupName:       "",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "UpdateSCIMGroupWhenTheGroupIDIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "",
			newGroupName:       "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "UpdateSCIMGroupWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			groupID:            "4475-a0c4",
			newGroupName:       "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "UpdateSCIMGroupWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			newGroupName:       "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "UpdateSCIMGroupWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			newGroupName:       "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 400,
			wantErr:            true,
		},
		{
			name:               "UpdateSCIMGroupWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			newGroupName:       "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            nil,
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "UpdateSCIMGroupWhenTheFilterIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			newGroupName:       "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "UpdateSCIMGroupWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			newGroupName:       "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "UpdateSCIMGroupWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			newGroupName:       "jira-users-updated",
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
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

			service := &SCIMGroupService{client: mockClient}
			gotResult, gotResponse, err := service.Update(testCase.context, testCase.directoryID,
				testCase.groupID, testCase.newGroupName)

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

				t.Log(gotResult.DisplayName, gotResult.ID, len(gotResult.Members))

			}

		})
	}

}

func TestSCIMGroupService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		group              string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CreateSCIMGroupWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			group:              "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "CreateSCIMGroupWhenTheNewGroupNameIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			group:              "",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "CreateSCIMGroupWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			group:              "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "CreateSCIMGroupWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			group:              "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "CreateSCIMGroupWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			group:              "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups",
			context:            context.Background(),
			wantHTTPCodeReturn: 400,
			wantErr:            true,
		},
		{
			name:               "CreateSCIMGroupWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4",
			group:              "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups",
			context:            nil,
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "CreateSCIMGroupWhenTheFilterIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			group:              "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "CreateSCIMGroupWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			group:              "jira-users-updated",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "CreateSCIMGroupWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			group:              "jira-users-updated",
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
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

			service := &SCIMGroupService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.directoryID, testCase.group)

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

				t.Log(gotResult.DisplayName, gotResult.ID, len(gotResult.Members))

			}

		})
	}

}

func TestSCIMGroupService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		groupID            string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteSCIMGroupWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "DeleteSCIMGroupWhenTheGroupIDIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "DeleteSCIMGroupWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			groupID:            "4475-a0c4",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "DeleteSCIMGroupWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "DeleteSCIMGroupWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 400,
			wantErr:            true,
		},
		{
			name:               "DeleteSCIMGroupWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            nil,
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "DeleteSCIMGroupWhenTheFilterIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "DeleteSCIMGroupWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
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

			service := &SCIMGroupService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.directoryID, testCase.groupID)

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

func TestSCIMGroupService_Path(t *testing.T) {

	payloadWithOperationsMocked := &model.SCIMGroupPathScheme{
		Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
		Operations: []*model.SCIMGroupOperationScheme{
			{
				Op:   "add",
				Path: "members",
				Value: []*model.SCIMGroupOperationValueScheme{
					{
						Value:   "635cdb2f-e72c-4122-bfd3-3aa",
						Display: "Example Display Name",
					},
				},
			},
		},
	}

	payloadWithOutOperationsMocked := &model.SCIMGroupPathScheme{
		Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
	}

	testCases := []struct {
		name               string
		directoryID        string
		groupID            string
		payload            *model.SCIMGroupPathScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "PatchSCIMGroupWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			payload:            payloadWithOperationsMocked,
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},

		{
			name:               "PatchSCIMGroupWhenThePayloadDoNotContainsOperations",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			payload:            payloadWithOutOperationsMocked,
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},

		{
			name:               "PatchSCIMGroupWhenThePayloadIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			payload:            nil,
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},

		{
			name:               "PatchSCIMGroupWhenTheGroupIDIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "",
			payload:            payloadWithOperationsMocked,
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "PatchSCIMGroupWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			payload:            payloadWithOperationsMocked,
			groupID:            "4475-a0c4",
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "PatchSCIMGroupWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			payload:            payloadWithOperationsMocked,
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "PatchSCIMGroupWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			payload:            payloadWithOperationsMocked,
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 400,
			wantErr:            true,
		},
		{
			name:               "PatchSCIMGroupWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			payload:            payloadWithOperationsMocked,
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            nil,
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "PatchSCIMGroupWhenTheFilterIsNotProvided",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			payload:            payloadWithOperationsMocked,
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},
		{
			name:               "PatchSCIMGroupWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			payload:            payloadWithOperationsMocked,
			mockFile:           "./mocks/scim-get-group.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
		{
			name:               "PatchSCIMGroupWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4",
			groupID:            "4475-a0c4",
			payload:            payloadWithOperationsMocked,
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4/Groups/4475-a0c4",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
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

			service := &SCIMGroupService{client: mockClient}
			gotResult, gotResponse, err := service.Path(testCase.context, testCase.directoryID,
				testCase.groupID, testCase.payload)

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

				t.Log(gotResult.DisplayName, gotResult.ID, len(gotResult.Members))

			}

		})
	}

}
