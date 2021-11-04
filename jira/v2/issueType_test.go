package v2

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestIssueTypeService_Alternatives(t *testing.T) {

	testCases := []struct {
		name               string
		issueTypeID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueTypeAlternativesWhenTheIssueTypeIDIsCorrect",
			issueTypeID:        "10000",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000/alternatives",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTypeAlternativesWhenTheIssueTypeIDIsNotProvided",
			issueTypeID:        "",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000/alternatives",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeAlternativesWhenTheIssueTypeIDIsIncorrect",
			issueTypeID:        "10001",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000/alternatives",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeAlternativesWhenTheContextIsNil",
			issueTypeID:        "10000",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000/alternatives",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeAlternativesWhenTheRequestMethodIsIncorrect",
			issueTypeID:        "10000",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetype/10000/alternatives",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeAlternativesWhenTheStatusCodeIsIncorrect",
			issueTypeID:        "10000",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000/alternatives",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeAlternativesWhenTheResponseBodyHasADifferentFormat",
			issueTypeID:        "10000",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000/alternatives",
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

			i := &IssueTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Alternatives(testCase.context, testCase.issueTypeID)

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

func TestIssueTypeService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *models.IssueTypePayloadScheme
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateIssueTypeWhenThePayloadIsCorrect",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name:               "CreateIssueTypeWhenThePayloadIsNil",
			payload:            nil,
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeWhenTheRequestMethodIsIncorrect",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeWhenTheStatusCodeIsIncorrect",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeWhenTheContextIsNil",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateIssueTypeWhenTheResponseBodyHasADifferentFormat",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodPost,
			mockFile:           "../v3/mocks/empty_json.json",
			endpoint:           "/rest/api/2/issuetype",
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

			i := &IssueTypeService{client: mockClient}

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

func TestIssueTypeService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		issueTypeID        string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteIssueTypeWhenTheIssueTypeIDIsCorrect",
			issueTypeID:        "10000",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteIssueTypeWhenTheIssueTypeIDIsNotProvided",
			issueTypeID:        "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueTypeWhenTheIssueTypeIDIsIncorrect",
			issueTypeID:        "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueTypeWhenTheRequestMethodIsIncorrect",
			issueTypeID:        "10000",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueTypeWhenTheStatusCodeIsIncorrect",
			issueTypeID:        "10000",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueTypeWhenTheContextIsNil",
			issueTypeID:        "10000",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/2/issuetype/10000",
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

			i := &IssueTypeService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.issueTypeID)

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

func TestIssueTypeService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueTypesWhenTheParamsAreCorrect",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTypesWhenTheEndpointIsIncorrect",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetypes",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypesWhenTheContextIsNil",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypesWhenTheRequestMethodIsNil",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/2/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypesWhenTheStatusCodeIsIncorrect",
			mockFile:           "../v3/mocks/get-issue-types.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypesWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype",
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

			i := &IssueTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context)

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

func TestIssueTypeService_Get(t *testing.T) {
	testCases := []struct {
		name               string
		issueTypeID        string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueTypeWhenTheIssueTypeIDIsCorrect",
			issueTypeID:        "10000",
			mockFile:           "../v3/mocks/get-issue-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetIssueTypeWhenTheIssueTypeIDIsNotProvided",
			issueTypeID:        "",
			mockFile:           "../v3/mocks/get-issue-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeWhenTheIssueTypeIDIsIncorrect",
			issueTypeID:        "10001",
			mockFile:           "../v3/mocks/get-issue-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeWhenTheContextIsNil",
			issueTypeID:        "10000",
			mockFile:           "../v3/mocks/get-issue-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeWhenTheRequestMethodIsIncorrect",
			issueTypeID:        "10000",
			mockFile:           "../v3/mocks/get-issue-type.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeWhenTheStatusCodeIsIncorrect",
			issueTypeID:        "10000",
			mockFile:           "../v3/mocks/get-issue-type.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetIssueTypeWhenTheResponseBodyHasADifferentFormat",
			issueTypeID:        "10000",
			mockFile:           "../v3/mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/issuetype/10000",
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

			i := &IssueTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.issueTypeID)

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

func TestIssueTypeService_Update(t *testing.T) {
	testCases := []struct {
		name               string
		issueTypeID        string
		payload            *models.IssueTypePayloadScheme
		wantHTTPMethod     string
		mockFile           string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:        "CreateIssueTypeWhenThePayloadIsCorrect",
			issueTypeID: "10001",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:        "CreateIssueTypeWhenTheIssueTypeIDIsNotProvided",
			issueTypeID: "",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateIssueTypeWhenThePayloadIsNil",
			issueTypeID:        "10001",
			payload:            nil,
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CreateIssueTypeWhenTheRequestMethodIsIncorrect",
			issueTypeID: "10001",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodDelete,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CreateIssueTypeWhenTheStatusCodeIsIncorrect",
			issueTypeID: "10001",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:        "CreateIssueTypeWhenTheContextIsNil",
			issueTypeID: "10001",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "../v3/mocks/get-issue-type.json",
			endpoint:           "/rest/api/2/issuetype/10001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:        "CreateIssueTypeWhenTheResponseBodyHasADifferentFormat",
			issueTypeID: "10001",
			payload: &models.IssueTypePayloadScheme{
				Name:        "Risk",
				Description: "Risk description",
				Type:        "standard",
			},
			wantHTTPMethod:     http.MethodPut,
			mockFile:           "../v3/mocks/empty_json.json",
			endpoint:           "/rest/api/2/issuetype/10001",
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

			i := &IssueTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.issueTypeID, testCase.payload)

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
