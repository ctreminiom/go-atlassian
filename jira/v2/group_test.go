package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestGroupService_Create_V2(t *testing.T) {

	testCases := []struct {
		name               string
		groupName          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CreateGroupWhenTheNameIsCorrect",
			groupName:          "power-users",
			mockFile:           "../v3/mocks/create-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateGroupWhenTheGroupNameIsNotSet",
			groupName:          "",
			mockFile:           "../v3/mocks/create-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateGroupWhenTheContextIsNil",
			groupName:          "power-users",
			mockFile:           "../v3/mocks/create-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "CreateGroupWhenTheRequestMethodIsIncorrect",
			groupName:          "power-users",
			mockFile:           "../v3/mocks/create-group.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "CreateGroupWhenTheStatusCodeIsIncorrect",
			groupName:          "power-users",
			mockFile:           "../v3/mocks/create-group.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "CreateGroupWhenTheResponseBodyLengthIsZero",
			groupName:          "power-users",
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "CreateGroupWhenTheResponseBodyHasADifferentFormat",
			groupName:          "power-users",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group",
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

			service := &GroupService{client: mockClient}

			gotResult, gotResponse, err := service.Create(testCase.context, testCase.groupName)

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
			}

		})
	}
}

func TestGroupService_Bulk_V2(t *testing.T) {

	testCases := []struct {
		name               string
		options            *models.GroupBulkOptionsScheme
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
			name: "BulkGroupsWhenTheGroupIDHasValues",
			options: &models.GroupBulkOptionsScheme{
				GroupNames: []string{"dog-developers", "jira-users"},
			},
			mockFile:           "../v3/mocks/bulk-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/bulk?groupName=dog-developers&groupName=jira-users&maxResults=0&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name: "BulkGroupsWhenTheGroupNameHasValues",
			options: &models.GroupBulkOptionsScheme{
				GroupIDs: []string{"5b10ac8d82e05b22cc7d4ef5", "5b10ac8d82e05b22cc4jas21409"},
			},
			mockFile:           "../v3/mocks/bulk-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/bulk?groupId=5b10ac8d82e05b22cc7d4ef5&groupId=5b10ac8d82e05b22cc4jas21409&maxResults=0&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "BulkGroupsWhenTheOptionIsNil",
			options:            nil,
			mockFile:           "../v3/mocks/bulk-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/bulk?groupName=dog-developers&groupName=jira-users&maxResults=0&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "BulkGroupsWhenTheRequestMethodIsIncorrect",
			options: &models.GroupBulkOptionsScheme{
				GroupNames: []string{"dog-developers", "jira-users"},
			},
			mockFile:           "../v3/mocks/bulk-groups.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/bulk?groupName=dog-developers&groupName=jira-users&maxResults=0&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "BulkGroupsWhenTheStatusCodeIsIncorrect",
			options: &models.GroupBulkOptionsScheme{
				GroupNames: []string{"dog-developers", "jira-users"},
			},
			mockFile:           "../v3/mocks/bulk-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/bulk?groupName=dog-developers&groupName=jira-users&maxResults=0&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name: "BulkGroupsWhenTheContextIsNil",
			options: &models.GroupBulkOptionsScheme{
				GroupNames: []string{"dog-developers", "jira-users"},
			},
			mockFile:           "../v3/mocks/bulk-groups.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/bulk?groupName=dog-developers&groupName=jira-users&maxResults=0&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "BulkGroupsWhenTheResponseBodyLengthIsZero",
			options: &models.GroupBulkOptionsScheme{
				GroupNames: []string{"dog-developers", "jira-users"},
			},
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/bulk?groupName=dog-developers&groupName=jira-users&maxResults=0&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "BulkGroupsWhenTheResponseBodyHasADifferentFormat",
			options: &models.GroupBulkOptionsScheme{
				GroupNames: []string{"dog-developers", "jira-users"},
			},
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/bulk?groupName=dog-developers&groupName=jira-users&maxResults=0&startAt=0",
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

			service := &GroupService{client: mockClient}

			gotResult, gotResponse, err := service.Bulk(testCase.context, testCase.options, 0, 0)

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
			}

		})
	}
}

func TestGroupService_Delete_V2(t *testing.T) {

	testCases := []struct {
		name               string
		groupName          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteGroupWhenTheNameIsCorrect",
			groupName:          "power-users",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteGroupWhenTheGroupNameIsNotSet",
			groupName:          "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteGroupWhenTheNameIsIncorrect",
			groupName:          "power-users-uat",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteGroupWhenTheContextIsNil",
			groupName:          "power-users",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group?groupname=power-users",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteGroupWhenTheRequestMethodIsIncorrect",
			groupName:          "power-users",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteGroupWhenTheStatusCodeIsIncorrect",
			groupName:          "power-users",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
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

			service := &GroupService{client: mockClient}

			gotResponse, err := service.Delete(testCase.context, testCase.groupName)

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

func TestGroupService_Members_V2(t *testing.T) {

	testCases := []struct {
		name               string
		groupName          string
		inactive           bool
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
			name:               "GetMembersGroupWhenTheGroupNameIsCorrect",
			groupName:          "power-users",
			inactive:           false,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/member?groupname=power-users&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetMembersGroupWhenTheGroupNameIsNotSet",
			groupName:          "",
			inactive:           false,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/member?groupname=power-users&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetMembersGroupWhenTheInactiveParameterIsSelected",
			groupName:          "power-users",
			inactive:           true,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/member?groupname=power-users&includeInactiveUsers=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetMembersGroupWhenTheGroupNameIsIncorrect",
			groupName:          "power-users-uat",
			inactive:           false,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/member?groupname=power-users&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetMembersGroupWhenTheRequestMethodIsIncorrect",
			groupName:          "power-users",
			inactive:           false,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/member?groupname=power-users&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetMembersGroupWhenTheStatusCodeIsIncorrect",
			groupName:          "power-users",
			inactive:           false,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/member?groupname=power-users&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetMembersGroupWhenTheContextIsNil",
			groupName:          "power-users",
			inactive:           false,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/member?groupname=power-users&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetMembersGroupWhenTheResponseBodyLengthIsZero",
			groupName:          "power-users",
			inactive:           false,
			startAt:            0,
			maxResults:         50,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/member?groupname=power-users&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetMembersGroupWhenTheResponseBodyHasADifferentFormat",
			groupName:          "power-users",
			inactive:           false,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/member?groupname=power-users&maxResults=50&startAt=0",
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

			service := &GroupService{client: mockClient}

			gotResult, gotResponse, err := service.Members(testCase.context,
				testCase.groupName,
				testCase.inactive,
				testCase.startAt,
				testCase.maxResults)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func TestGroupService_Add_V2(t *testing.T) {

	testCases := []struct {
		name               string
		groupName          string
		accountID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddGroupMemberWhenTheGroupNameIsCorrect",
			groupName:          "power-users",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/user?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "AddGroupMemberWhenTheGroupNameIsNotSet",
			groupName:          "",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/user?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddGroupMemberWhenTheAccountIDIsNotSet",
			groupName:          "power-users",
			accountID:          "",
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/user?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddGroupMemberWhenTheGroupNameIsIncorrect",
			groupName:          "power-users-uat",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/user?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "AddGroupMemberWhenTheRequestMethodIsIncorrect",
			groupName:          "power-users",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/group/user?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "AddGroupMemberWhenTheStatusCodeIsIncorrect",
			groupName:          "power-users",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/user?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "AddGroupMemberWhenTheContextIsNil",
			groupName:          "power-users",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			mockFile:           "../v3/mocks/group-members.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/user?groupname=power-users",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "AddGroupMemberWhenTheResponseBodyLengthIsZero",
			groupName:          "power-users",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/user?groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name:               "AddGroupMemberWhenTheResponseBodyHasADifferentFormat",
			groupName:          "power-users",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/user?groupname=power-users",
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

			service := &GroupService{client: mockClient}

			gotResult, gotResponse, err := service.Add(testCase.context, testCase.groupName, testCase.accountID)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}

func TestGroupService_Remove_V2(t *testing.T) {

	testCases := []struct {
		name               string
		groupName          string
		accountID          string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "RemoveGroupMemberWhenTheGroupNameIsCorrect",
			groupName:          "power-users",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group/user?accountId=b78f2e47-f267-48a2-b91d-682992f9a0b0&groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "RemoveGroupMemberWhenTheGroupNameIsNotSet",
			groupName:          "",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group/user?accountId=b78f2e47-f267-48a2-b91d-682992f9a0b0&groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveGroupMemberWhenTheAccountIDIsNotSet",
			groupName:          "power-users",
			accountID:          "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group/user?accountId=b78f2e47-f267-48a2-b91d-682992f9a0b0&groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "RemoveGroupMemberWhenTheGroupNameIsIncorrect",
			groupName:          "power-users-uat",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group/user?accountId=b78f2e47-f267-48a2-b91d-682992f9a0b0&groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "RemoveGroupMemberWhenTheRequestMethodIsIncorrect",
			groupName:          "power-users",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/group/user?accountId=b78f2e47-f267-48a2-b91d-682992f9a0b0&groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "RemoveGroupMemberWhenTheStatusCodeIsIncorrect",
			groupName:          "power-users",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group/user?accountId=b78f2e47-f267-48a2-b91d-682992f9a0b0&groupname=power-users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "RemoveGroupMemberWhenTheContextIsNil",
			groupName:          "power-users",
			accountID:          "b78f2e47-f267-48a2-b91d-682992f9a0b0",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/group/user?accountId=b78f2e47-f267-48a2-b91d-682992f9a0b0&groupname=power-users",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       "",
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

			service := &GroupService{client: mockClient}

			gotResponse, err := service.Remove(testCase.context, testCase.groupName, testCase.accountID)

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

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}
}
