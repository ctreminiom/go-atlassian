package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestProjectRoleActorService_Add(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		projectRoleID      int
		accountIDs         []string
		groups             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "AddActorsToProjectRoleWhenTheParamsAreCorrect",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountIDs:         []string{"17bbc4fc-88ba-40a8-b38d-8a5afecaf830", "f019b08b-5faf-4691-b0c3-ce057313001a"},
			groups:             []string{"jira-users", "jira-developers"},
			mockFile:           "./mocks/create-project-role-actor.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "AddActorsToProjectRoleWhenTheProjectKeyOrIDIsEmpty",
			projectKeyOrID:     "",
			projectRoleID:      1001,
			accountIDs:         []string{"17bbc4fc-88ba-40a8-b38d-8a5afecaf830", "f019b08b-5faf-4691-b0c3-ce057313001a"},
			groups:             []string{"jira-users", "jira-developers"},
			mockFile:           "./mocks/create-project-role-actor.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddActorsToProjectRoleWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountIDs:         []string{"17bbc4fc-88ba-40a8-b38d-8a5afecaf830", "f019b08b-5faf-4691-b0c3-ce057313001a"},
			groups:             []string{"jira-users", "jira-developers"},
			mockFile:           "./mocks/create-project-role-actor.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddActorsToProjectRoleWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountIDs:         []string{"17bbc4fc-88ba-40a8-b38d-8a5afecaf830", "f019b08b-5faf-4691-b0c3-ce057313001a"},
			groups:             []string{"jira-users", "jira-developers"},
			mockFile:           "./mocks/create-project-role-actor.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "AddActorsToProjectRoleWhenTheEndpointIsIncorrect",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountIDs:         []string{"17bbc4fc-88ba-40a8-b38d-8a5afecaf830", "f019b08b-5faf-4691-b0c3-ce057313001a"},
			groups:             []string{"jira-users", "jira-developers"},
			mockFile:           "./mocks/create-project-role-actor.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/project/DUMMY/role/1001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddActorsToProjectRoleWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountIDs:         []string{"17bbc4fc-88ba-40a8-b38d-8a5afecaf830", "f019b08b-5faf-4691-b0c3-ce057313001a"},
			groups:             []string{"jira-users", "jira-developers"},
			mockFile:           "./mocks/create-project-role-actor.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "AddActorsToProjectRoleWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountIDs:         []string{"17bbc4fc-88ba-40a8-b38d-8a5afecaf830", "f019b08b-5faf-4691-b0c3-ce057313001a"},
			groups:             []string{"jira-users", "jira-developers"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001",
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

			i := &ProjectRoleActorService{client: mockClient}

			gotResult, gotResponse, err := i.Add(testCase.context, testCase.projectKeyOrID, testCase.projectRoleID,
				testCase.accountIDs, testCase.groups)

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

				t.Log("---------------------------------")
				t.Logf("Role Name: %v", gotResult.Name)
				t.Logf("Role ID: %v", gotResult.ID)
				t.Logf("Role Self: %v", gotResult.Self)
				t.Logf("Role Description: %v", gotResult.Description)
				t.Logf("Role Actors: %v", len(gotResult.Actors))
				t.Log("---------------------------------")

			}
		})

	}

}

func TestProjectRoleActorService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		projectRoleID      int
		accountID, group   string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteActorsToProjectRoleWhenTheParamsAreCorrect",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountID:          "17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			group:              "jira-developers",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001?group=jira-developers&user=17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteActorsToProjectRoleWhenTheProjectKeyOrIDIsEmpty",
			projectKeyOrID:     "",
			projectRoleID:      1001,
			accountID:          "17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			group:              "jira-developers",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001?group=jira-developers&user=17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteActorsToProjectRoleWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountID:          "17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			group:              "jira-developers",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001?group=jira-developers&user=17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteActorsToProjectRoleWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountID:          "17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			group:              "jira-developers",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001?group=jira-developers&user=17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteActorsToProjectRoleWhenTheEndpointsIsIncorrect",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountID:          "17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			group:              "jira-developers",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/project/DUMMY/role/1001?group=jira-developers&user=17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteActorsToProjectRoleWhenTheContextIsNil",
			projectKeyOrID:     "DUMMY",
			projectRoleID:      1001,
			accountID:          "17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
			group:              "jira-developers",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/project/DUMMY/role/1001?group=jira-developers&user=17bbc4fc-88ba-40a8-b38d-8a5afecaf830",
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

			i := &ProjectRoleActorService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.projectKeyOrID, testCase.projectRoleID,
				testCase.accountID, testCase.group)

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
