package jira

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestIssuePropertyService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		issueKeyORID       string
		propertyKey        string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssuePropertyWhenTheKeyIsCorrect",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "issue.support",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.support",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetIssuePropertyWhenTheKeyIsIncorrect",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "issue.support.test",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.support",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssuePropertyWhenTheKeyIsEmpty",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.support",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssuePropertyWhenTheKeyIsANumber",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "1",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.support",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssuePropertyWhenTheIssueIsIncorrect",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "issue.support",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.supports",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssuePropertyWhenTheIssueIsEmpty",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "",
			propertyKey:        "issue.support",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.support",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssuePropertyWhenTheHTTPMethodIsDifferent",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "issue.support",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.supports",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetIssuePropertyWhenTheHTTPCodeIsDifferent",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "issue.support",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.supports",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
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
				Headers:            testCase.wantHTTPHeaders,
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

			service := &IssuePropertyService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.issueKeyORID, testCase.propertyKey)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func TestIssuePropertyService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		issueKeyORID       string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetsIssuePropertiesWhenTheKeyIsCorrect",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "DUMMY-3",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetsIssuePropertiesWhenTheKeyIsIncorrect",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "DUMMY-1",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsIssuePropertiesWhenTheKeyIsEmpty",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsIssuePropertiesWhenTheKeyIsANumber",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "1",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsIssuePropertiesWhenTheKeyIsASpecialCharacters",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "$#@##",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetsIssuePropertiesWhenTheKeyIsASpecialCharacters",
			mockFile:           "./mocks/get-issue-properties.json",
			issueKeyORID:       "$#@##",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
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
				Headers:            testCase.wantHTTPHeaders,
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

			service := &IssuePropertyService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.issueKeyORID)

			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)
			}

		})
	}
}

func TestIssuePropertyService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		issueKeyORID       string
		propertyKey        string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
		errMessage         error
	}{
		{
			name:               "DeleteIssuePropertiesWhenTheKeyIsCorrect",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "issue.support",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.support",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},
		{
			name:               "DeleteIssuePropertiesWhenTheKeyIsIncorrect",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "issue.supports",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.support",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			errMessage:         errors.New("request failed. Please analyze the request body for more details. Status Code: 400"),
		},
		{
			name:               "DeleteIssuePropertiesWhenTheKeyIsEmpty",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.support",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			errMessage:         errors.New("request failed. Please analyze the request body for more details. Status Code: 400"),
		},
		{
			name:               "DeleteIssuePropertiesWhenTheKeyIsANumber",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "2222",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.support",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			errMessage:         errors.New("request failed. Please analyze the request body for more details. Status Code: 400"),
		},
		{
			name:               "DeleteIssuePropertiesWhenTheKeyIsANumber",
			issueKeyORID:       "DUMMY-3",
			propertyKey:        "1",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issue/DUMMY-3/properties/issue.support",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
			errMessage:         errors.New("request failed. Please analyze the request body for more details. Status Code: 400"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MethodAccepted:     testCase.wantHTTPMethod,
				Headers:            testCase.wantHTTPHeaders,
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

			service := &IssuePropertyService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.issueKeyORID, testCase.propertyKey)

			if testCase.wantErr {
				if assert.Error(t, err) {
					assert.Equal(t, testCase.errMessage, err)
				}
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}

		})
	}
}
