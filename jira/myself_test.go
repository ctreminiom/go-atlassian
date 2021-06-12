package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestMySelfService_Details(t *testing.T) {

	testCases := []struct {
		name               string
		expands            []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetMySelfDetailsWhenTheParamsAreCorrect",
			expands:            []string{"applicationRoles", "groups"},
			mockFile:           "./mocks/get-myselft.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/myself?expand=applicationRoles%2Cgroups",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetMySelfDetailsWhenTheContextIsNil",
			expands:            []string{"applicationRoles", "groups"},
			mockFile:           "./mocks/get-myselft.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/myself?expand=applicationRoles%2Cgroups",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetMySelfDetailsWhenTheExpandsAreNotSet",
			expands:            nil,
			mockFile:           "./mocks/get-myselft.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/myself",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetMySelfDetailsWhenTheResponseBodyIsEmpty",
			expands:            []string{"applicationRoles", "groups"},
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/myself?expand=applicationRoles%2Cgroups",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetMySelfDetailsWhenTheRequestMethodIsIncorrect",
			expands:            []string{"applicationRoles", "groups"},
			mockFile:           "./mocks/get-myselft.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/api/3/myself?expand=applicationRoles%2Cgroups",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetMySelfDetailsWhenTheStatusCodeIsIncorrect",
			expands:            []string{"applicationRoles", "groups"},
			mockFile:           "./mocks/get-myselft.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/myself?expand=applicationRoles%2Cgroups",
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

			i := &MySelfService{client: mockClient}

			gotResult, gotResponse, err := i.Details(testCase.context, testCase.expands)

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

				for _, group := range gotResult.Groups.Items {
					t.Log(group.Self, group.Name)
				}

				for _, item := range gotResult.ApplicationRoles.Items {
					t.Log(item.Key, item.Name)
				}
			}
		})

	}

}
