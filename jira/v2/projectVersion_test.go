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

func TestProjectVersionService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.VersionPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateProjectVersionWhenTheParamsAreCorrect",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
				Name:        "New Version 1",
				Description: "An excellent version",
				ProjectID:   0,
				Released:    true,
			},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/version",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateProjectVersionWhenThePayloadIsNil",
			payload:            nil,
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/version",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectVersionWhenTheRequestMethodIsIncorrect",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
				Name:        "New Version 1",
				Description: "An excellent version",
				ProjectID:   0,
				Released:    true,
			},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectVersionWhenTheStatusCodeIsIncorrect",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
				Name:        "New Version 1",
				Description: "An excellent version",
				ProjectID:   0,
				Released:    true,
			},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/version",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CreateProjectVersionWhenTheContextIsNil",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
				Name:        "New Version 1",
				Description: "An excellent version",
				ProjectID:   0,
				Released:    true,
			},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/version",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateProjectVersionWhenTheResponseBodyHasADifferentFormat",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
				Name:        "New Version 1",
				Description: "An excellent version",
				ProjectID:   0,
				Released:    true,
			},
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/version",
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

			i := &ProjectVersionService{client: mockClient}

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

				t.Log("------------------------------")
				t.Logf("Project Component Name: %v", gotResult.Name)
				t.Logf("Project Component ID: %v", gotResult.ID)
				t.Logf("Project Component Description: %v", gotResult.Description)
				t.Logf("Project Component Archived?: %v", gotResult.Archived)
				t.Logf("Project Component Released?: %v", gotResult.Released)
				t.Logf("Project Component ProjectID: %v", gotResult.ProjectID)
				t.Log("------------------------------")

			}
		})

	}

}

func TestProjectVersionService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		versionID          string
		expands            []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectVersionWhenTheParamsAreCorrect",
			versionID:          "10001",
			expands:            []string{"operations", "issuesstatus"},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/10001?expand=operations%2Cissuesstatus",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectVersionWhenTheVersionIDIsNotProvided",
			versionID:          "",
			expands:            []string{"operations", "issuesstatus"},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/10001?expand=operations%2Cissuesstatus",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionWhenTheExpandsIsNil",
			versionID:          "10001",
			expands:            nil,
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectVersionWhenTheRequestMethodIsIncorrect",
			versionID:          "10001",
			expands:            []string{"operations", "issuesstatus"},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/version/10001?expand=operations%2Cissuesstatus",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionWhenTheStatusCodeIsIncorrect",
			versionID:          "10001",
			expands:            []string{"operations", "issuesstatus"},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/10001?expand=operations%2Cissuesstatus",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionWhenTheContextIsNil",
			versionID:          "10001",
			expands:            []string{"operations", "issuesstatus"},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/10001?expand=operations%2Cissuesstatus",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionWhenTheResponseBodyHasADifferentFormat",
			versionID:          "10001",
			expands:            []string{"operations", "issuesstatus"},
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/10001?expand=operations%2Cissuesstatus",
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

			i := &ProjectVersionService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.versionID, testCase.expands)

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

				t.Log("------------------------------")
				t.Logf("Project Component Name: %v", gotResult.Name)
				t.Logf("Project Component ID: %v", gotResult.ID)
				t.Logf("Project Component Description: %v", gotResult.Description)
				t.Logf("Project Component Archived?: %v", gotResult.Archived)
				t.Logf("Project Component Released?: %v", gotResult.Released)
				t.Logf("Project Component ProjectID: %v", gotResult.ProjectID)
				t.Log("------------------------------")

			}
		})

	}

}

func TestProjectVersionService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		projectKeyOrID     string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectVersionsWhenTheParamsAreCorrect",
			projectKeyOrID:     "10001",
			mockFile:           "../mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/10001/versions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectVersionsWhenTheProjectKeyIsNotProvided",
			projectKeyOrID:     "",
			mockFile:           "../mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/10001/versions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionsWhenTheExpandsIsNil",
			projectKeyOrID:     "10001",
			mockFile:           "../mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/10001/versions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectVersionsWhenTheRequestMethodIsIncorrect",
			projectKeyOrID:     "10001",
			mockFile:           "../mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/project/10001/versions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionsWhenTheStatusCodeIsIncorrect",
			projectKeyOrID:     "10001",
			mockFile:           "../mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/10001/versions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionsWhenTheContextIsNil",
			projectKeyOrID:     "10001",
			mockFile:           "../mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/10001/versions",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionsWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID:     "10001",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/10001/versions",
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

			i := &ProjectVersionService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context, testCase.projectKeyOrID)

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

func TestProjectVersionService_Search(t *testing.T) {

	testCases := []struct {
		name                string
		projectKeyOrID      string
		options             *models.VersionGetsOptions
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:           "SearchProjectVersionsWhenTheParamsAreCorrect",
			projectKeyOrID: "DUMMY",
			options: &models.VersionGetsOptions{
				OrderBy: "description",
				Query:   "name",
				Status:  "unreleased",
				Expand:  []string{"issuesstatus", "operations"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/version?expand=issuesstatus%2Coperations&maxResults=50&orderBy=description&query=name&startAt=0&status=unreleased",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:           "SearchProjectVersionsWhenTheOptionsAreNotProvided",
			projectKeyOrID: "DUMMY",
			options: &models.VersionGetsOptions{
				OrderBy: "",
				Query:   "",
				Status:  "",
				Expand:  nil,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/version?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "SearchProjectVersionsWhenTheOptionsIsNil",
			projectKeyOrID:     "DUMMY",
			options:            nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/version?expand=issuesstatus%2Coperations&maxResults=50&orderBy=description&query=name&startAt=0&status=unreleased",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:           "SearchProjectVersionsWhenTheRequestMethodIsIncorrect",
			projectKeyOrID: "",
			options: &models.VersionGetsOptions{
				OrderBy: "description",
				Query:   "name",
				Status:  "unreleased",
				Expand:  []string{"issuesstatus", "operations"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/project/DUMMY/version?expand=issuesstatus%2Coperations&maxResults=50&orderBy=description&query=name&startAt=0&status=unreleased",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:           "SearchProjectVersionsWhenTheStatusCodeIsIncorrect",
			projectKeyOrID: "DUMMY",
			options: &models.VersionGetsOptions{
				OrderBy: "description",
				Query:   "name",
				Status:  "unreleased",
				Expand:  []string{"issuesstatus", "operations"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/version?expand=issuesstatus%2Coperations&maxResults=50&orderBy=description&query=name&startAt=0&status=unreleased",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:           "SearchProjectVersionsWhenTheContextIsNil",
			projectKeyOrID: "DUMMY",
			options: &models.VersionGetsOptions{
				OrderBy: "description",
				Query:   "name",
				Status:  "unreleased",
				Expand:  []string{"issuesstatus", "operations"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/get-project-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/version?expand=issuesstatus%2Coperations&maxResults=50&orderBy=description&query=name&startAt=0&status=unreleased",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:           "SearchProjectVersionsWhenTheResponseBodyHasADifferentFormat",
			projectKeyOrID: "DUMMY",
			options: &models.VersionGetsOptions{
				OrderBy: "description",
				Query:   "name",
				Status:  "unreleased",
				Expand:  []string{"issuesstatus", "operations"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/project/DUMMY/version?expand=issuesstatus%2Coperations&maxResults=50&orderBy=description&query=name&startAt=0&status=unreleased",
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

			i := &ProjectVersionService{client: mockClient}

			gotResult, gotResponse, err := i.Search(testCase.context, testCase.projectKeyOrID, testCase.options,
				testCase.startAt, testCase.maxResults)

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

				for _, version := range gotResult.Values {

					t.Log("------------------------------")
					t.Logf("Project Component Name: %v", version.Name)
					t.Logf("Project Component ID: %v", version.ID)
					t.Logf("Project Component Description: %v", version.Description)
					t.Logf("Project Component Archived?: %v", version.Archived)
					t.Logf("Project Component Released?: %v", version.Released)
					t.Logf("Project Component ProjectID: %v", version.ProjectID)
					t.Log("------------------------------ \n")
				}

			}
		})

	}

}

func TestProjectVersionService_Merge(t *testing.T) {

	testCases := []struct {
		name                    string
		versionID, moveIssuesTo string
		mockFile                string
		wantHTTPMethod          string
		endpoint                string
		context                 context.Context
		wantHTTPCodeReturn      int
		wantErr                 bool
	}{
		{
			name:               "GetProjectVersionsWhenTheParamsAreCorrect",
			versionID:          "1001",
			moveIssuesTo:       "2224422",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1001/mergeto/2224422",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "GetProjectVersionsWhenTheVersionIDIsEmpty",
			versionID:          "",
			moveIssuesTo:       "2224422",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1001/mergeto/2224422",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionsWhenTheMoveIssueToIsEmpty",
			versionID:          "1001",
			moveIssuesTo:       "",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1001/mergeto/2224422",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionsWhenTheRequestMethodIsIncorrect",
			versionID:          "1001",
			moveIssuesTo:       "2224422",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/version/1001/mergeto/2224422",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionsWhenTheStatusCodeIsIncorrect",
			versionID:          "1001",
			moveIssuesTo:       "2224422",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1001/mergeto/2224422",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectVersionsWhenTheContextIsNil",
			versionID:          "1001",
			moveIssuesTo:       "2224422",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1001/mergeto/2224422",
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

			i := &ProjectVersionService{client: mockClient}

			gotResponse, err := i.Merge(testCase.context, testCase.versionID, testCase.moveIssuesTo)

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

func TestProjectVersionService_RelatedIssueCounts(t *testing.T) {

	testCases := []struct {
		name               string
		versionID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectRelatedIssueCountVersionWhenTheParamsAreCorrect",
			versionID:          "1001",
			mockFile:           "../v3/mocks/get-related-issue-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/1001/relatedIssueCounts",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectRelatedIssueCountVersionWhenTheVersionIDIsIncorrect",
			versionID:          "",
			mockFile:           "../v3/mocks/get-related-issue-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/1001/relatedIssueCounts",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRelatedIssueCountVersionWhenTheRequestMethodIsIncorrect",
			versionID:          "1001",
			mockFile:           "../v3/mocks/get-related-issue-count.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/version/1001/relatedIssueCounts",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRelatedIssueCountVersionWhenTheStatusCodeIsIncorrect",
			versionID:          "1001",
			mockFile:           "../v3/mocks/get-related-issue-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/1001/relatedIssueCounts",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectRelatedIssueCountVersionWhenTheContextIsNil",
			versionID:          "1001",
			mockFile:           "../v3/mocks/get-related-issue-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/1001/relatedIssueCounts",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectRelatedIssueCountVersionWhenTheResponseBodyHasADifferentFormat",
			versionID:          "1001",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/1001/relatedIssueCounts",
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

			i := &ProjectVersionService{client: mockClient}

			gotResult, gotResponse, err := i.RelatedIssueCounts(testCase.context, testCase.versionID)

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

				for _, usage := range gotResult.CustomFieldUsage {

					t.Logf("Custom Field Usage Name: %v", usage.FieldName)
					t.Logf("Custom Field Usage ID: %v", usage.CustomFieldID)
					t.Logf("Custom Field Usage Count: %v", usage.IssueCountWithVersionInCustomField)
				}

			}
		})

	}

}

func TestProjectVersionService_UnresolvedIssueCount(t *testing.T) {

	testCases := []struct {
		name               string
		versionID          string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetProjectUnresolvedIssueCountVersionWhenTheParamsAreCorrect",
			versionID:          "1001",
			mockFile:           "../v3/mocks/get-project-version-unresolved-issue-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/1001/unresolvedIssueCount",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetProjectUnresolvedIssueCountVersionWhenTheVersionIDIsEmpty",
			versionID:          "",
			mockFile:           "../v3/mocks/get-project-version-unresolved-issue-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/1001/unresolvedIssueCount",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectUnresolvedIssueCountVersionWhenTheRequestMethodIsIncorrect",
			versionID:          "1001",
			mockFile:           "../v3/mocks/get-project-version-unresolved-issue-count.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/version/1001/unresolvedIssueCount",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectUnresolvedIssueCountVersionWhenTheStatusCodeIsIncorrect",
			versionID:          "1001",
			mockFile:           "../v3/mocks/get-project-version-unresolved-issue-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/1001/unresolvedIssueCount",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetProjectUnresolvedIssueCountVersionWhenTheContextIsNil",
			versionID:          "1001",
			mockFile:           "../v3/mocks/get-project-version-unresolved-issue-count.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/1001/unresolvedIssueCount",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetProjectUnresolvedIssueCountVersionWhenTheResponseBodyHasADifferentFormat",
			versionID:          "1001",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/version/1001/unresolvedIssueCount",
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

			i := &ProjectVersionService{client: mockClient}

			gotResult, gotResponse, err := i.UnresolvedIssueCount(testCase.context, testCase.versionID)

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

				t.Log("-------------------------------")
				t.Logf("Self: %v", gotResult.Self)
				t.Logf("Issues Count: %v", gotResult.IssuesCount)
				t.Logf("Issues Unresolved Count: %v", gotResult.IssuesUnresolvedCount)
				t.Log("-------------------------------")

			}
		})

	}

}

func TestProjectVersionService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		versionID          string
		payload            *models.VersionPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:      "UpdateProjectVersionWhenTheParamsAreCorrect",
			versionID: "1000",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
			},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:      "UpdateProjectVersionWhenTheVersionIDIsIncorrect",
			versionID: "",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
			},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateProjectVersionWhenThePayloadIsNil",
			versionID:          "1000",
			payload:            nil,
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateProjectVersionWhenTheRequestMethodIsIncorrect",
			versionID: "1000",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
			},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/version/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateProjectVersionWhenTheStatusCodeIsIncorrect",
			versionID: "1000",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
			},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:      "UpdateProjectVersionWhenTheContextIsNil",
			versionID: "1000",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
			},
			mockFile:           "../v3/mocks/create-project-version.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:      "UpdateProjectVersionWhenTheResponseBodyHasADifferentFormat",
			versionID: "1000",
			payload: &models.VersionPayloadScheme{
				Archived:    false,
				ReleaseDate: "6/Jul/2020",
			},
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/version/1000",
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

			i := &ProjectVersionService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.versionID, testCase.payload)

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

				t.Log("------------------------------")
				t.Logf("Project Component Name: %v", gotResult.Name)
				t.Logf("Project Component ID: %v", gotResult.ID)
				t.Logf("Project Component Description: %v", gotResult.Description)
				t.Logf("Project Component Archived?: %v", gotResult.Archived)
				t.Logf("Project Component Released?: %v", gotResult.Released)
				t.Logf("Project Component ProjectID: %v", gotResult.ProjectID)
				t.Log("------------------------------")

			}
		})

	}

}
