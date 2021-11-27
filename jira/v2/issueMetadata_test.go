package v2

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_IssueMetadataService_Get_Success(t *testing.T) {

	expectedJSONAsBytes, err := ioutil.ReadFile("../v3/mocks/get-issue-metadata.json")
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name                   string
		overrideScreenSecurity bool
		overrideEditableFlag   bool
		issueKeyOrID           string
		wantHTTPMethod         string
		mockFile               string
		endpoint               string
		context                context.Context
		wantHTTPCodeReturn     int
		wantResult             gjson.Result
	}{
		{
			name:                   "when_the_parameters_are_correct",
			overrideScreenSecurity: true,
			overrideEditableFlag:   true,
			issueKeyOrID:           "KP-19",
			wantHTTPMethod:         http.MethodGet,
			mockFile:               "../v3/mocks/get-issue-metadata.json",
			endpoint:               "/rest/api/2/issue/KP-19/editmeta?overrideEditableFlag=true&overrideScreenSecurity=true",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			wantResult:             gjson.ParseBytes(expectedJSONAsBytes),
		},
		{
			name:                   "when_the_overrideEditableFlag_param_is_not_set",
			overrideScreenSecurity: true,
			overrideEditableFlag:   false,
			issueKeyOrID:           "KP-19",
			wantHTTPMethod:         http.MethodGet,
			mockFile:               "../v3/mocks/get-issue-metadata.json",
			endpoint:               "/rest/api/2/issue/KP-19/editmeta?overrideScreenSecurity=true",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			wantResult:             gjson.ParseBytes(expectedJSONAsBytes),
		},
		{
			name:                   "when_the_overrideScreenSecurity_param_is_not_set",
			overrideScreenSecurity: false,
			overrideEditableFlag:   true,
			issueKeyOrID:           "KP-19",
			wantHTTPMethod:         http.MethodGet,
			mockFile:               "../v3/mocks/get-issue-metadata.json",
			endpoint:               "/rest/api/2/issue/KP-19/editmeta?overrideEditableFlag=true",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			wantResult:             gjson.ParseBytes(expectedJSONAsBytes),
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

			service := &IssueMetadataService{client: mockClient}

			gotResult, gotResponse, err := service.Get(
				testCase.context,
				testCase.issueKeyOrID,
				testCase.overrideScreenSecurity,
				testCase.overrideEditableFlag,
			)

			assert.NoError(t, err)
			assert.NotEqual(t, gotResponse, nil)
			assert.NotEqual(t, gotResult, nil)
			assert.Equal(t, testCase.wantResult, gotResult)

			endpointToAssert, err := extractEndpotintToAssert(gotResponse.Endpoint)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, testCase.endpoint, endpointToAssert)
		})
	}
}

func Test_IssueMetadataService_Get_Failed(t *testing.T) {

	testCases := []struct {
		name                   string
		overrideScreenSecurity bool
		overrideEditableFlag   bool
		issueKeyOrID           string
		wantHTTPMethod         string
		mockFile               string
		endpoint               string
		context                context.Context
		wantHTTPCodeReturn     int
		expectedErrorMessage   string
	}{
		{
			name:                   "when_the_http_request_method_is_incorrect",
			overrideScreenSecurity: true,
			overrideEditableFlag:   true,
			issueKeyOrID:           "KP-19",
			wantHTTPMethod:         http.MethodPost,
			mockFile:               "../v3/mocks/get-issue-metadata.json",
			endpoint:               "/rest/api/2/issue/KP-19/editmeta?overrideEditableFlag=true&overrideScreenSecurity=true",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			expectedErrorMessage:   "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name:                   "when_the_context_provided_is_nil",
			overrideScreenSecurity: true,
			overrideEditableFlag:   true,
			issueKeyOrID:           "KP-19",
			wantHTTPMethod:         http.MethodGet,
			mockFile:               "../v3/mocks/get-issue-metadata.json",
			endpoint:               "/rest/api/2/issue/KP-19/editmeta?overrideEditableFlag=true&overrideScreenSecurity=true",
			context:                nil,
			wantHTTPCodeReturn:     http.StatusOK,
			expectedErrorMessage:   "request creation failed: net/http: nil Context",
		},

		{
			name:                   "when_the_issue_key_or_id_is_not_provided",
			overrideScreenSecurity: true,
			overrideEditableFlag:   true,
			issueKeyOrID:           "",
			wantHTTPMethod:         http.MethodPost,
			mockFile:               "../v3/mocks/get-issue-metadata.json",
			endpoint:               "/rest/api/2/issue/KP-19/editmeta?overrideEditableFlag=true&overrideScreenSecurity=true",
			context:                context.Background(),
			wantHTTPCodeReturn:     http.StatusOK,
			expectedErrorMessage:   "jira: no issue key/id set",
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

			service := &IssueMetadataService{client: mockClient}

			_, _, err = service.Get(
				testCase.context,
				testCase.issueKeyOrID,
				testCase.overrideScreenSecurity,
				testCase.overrideEditableFlag,
			)

			assert.EqualError(t, err, testCase.expectedErrorMessage)
		})
	}
}

func Test_IssueMetadataService_Create_Success(t *testing.T) {

	expectedJSONAsBytes, err := ioutil.ReadFile("../v3/mocks/get-issue-create-metadata.json")
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name               string
		opts               *models.IssueMetadataCreateOptions
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantResult         gjson.Result
	}{
		{
			name: "when_the_parameters_are_correct",
			opts: &models.IssueMetadataCreateOptions{
				ProjectIDs:     []string{"11101"},
				ProjectKeys:    []string{"KP"},
				IssueTypeIDs:   []string{"12"},
				IssueTypeNames: []string{"issue-type"},
				Expand:         "expand",
			},
			wantHTTPMethod:     http.MethodGet,
			mockFile:           "../v3/mocks/get-issue-create-metadata.json",
			endpoint:           "/rest/api/2/issue/createmeta?expand=expand&issuetypeIds=12&issuetypeNames=issue-type&projectIds=11101&projectKeys=KP",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantResult:         gjson.ParseBytes(expectedJSONAsBytes),
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

			service := &IssueMetadataService{client: mockClient}

			gotResult, gotResponse, err := service.Create(
				testCase.context,
				testCase.opts,
			)

			assert.NoError(t, err)
			assert.NotEqual(t, gotResponse, nil)
			assert.NotEqual(t, gotResult, nil)
			assert.Equal(t, testCase.wantResult, gotResult)

			endpointToAssert, err := extractEndpotintToAssert(gotResponse.Endpoint)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, testCase.endpoint, endpointToAssert)

		})
	}
}

func Test_IssueMetadataService_Create_Failed(t *testing.T) {

	testCases := []struct {
		name                 string
		opts                 *models.IssueMetadataCreateOptions
		wantHTTPMethod       string
		mockFile             string
		endpoint             string
		context              context.Context
		wantHTTPCodeReturn   int
		wantResult           gjson.Result
		expectedErrorMessage string
	}{
		{
			name: "when_the_http_request_method_is_incorrect",
			opts: &models.IssueMetadataCreateOptions{
				ProjectIDs:   []string{"11101"},
				ProjectKeys:  []string{"KP"},
				IssueTypeIDs: []string{"12"},
			},
			wantHTTPMethod:       http.MethodPost,
			mockFile:             "../v3/mocks/get-issue-create-metadata.json",
			endpoint:             "/rest/api/2/issue/createmeta?issuetypeIds=12&projectIds=11101&projectKeys=KP",
			context:              context.Background(),
			wantHTTPCodeReturn:   http.StatusOK,
			expectedErrorMessage: "request failed. Please analyze the request body for more details. Status Code: 405",
		},

		{
			name: "when_the_context_provided_is_nil",
			opts: &models.IssueMetadataCreateOptions{
				ProjectIDs:   []string{"11101"},
				ProjectKeys:  []string{"KP"},
				IssueTypeIDs: []string{"12"},
			},
			wantHTTPMethod:       http.MethodGet,
			mockFile:             "../v3/mocks/get-issue-create-metadata.json",
			endpoint:             "/rest/api/2/issue/createmeta?issuetypeIds=12&projectIds=11101&projectKeys=KP",
			context:              nil,
			wantHTTPCodeReturn:   http.StatusOK,
			expectedErrorMessage: "request creation failed: net/http: nil Context",
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

			service := &IssueMetadataService{client: mockClient}

			_, _, err = service.Create(
				testCase.context,
				testCase.opts,
			)

			assert.EqualError(t, err, testCase.expectedErrorMessage)

		})
	}
}
