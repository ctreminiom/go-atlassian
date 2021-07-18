package agile

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestSprintService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		sprintID           int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetSprintWhenTheParametersAreCorrect",
			sprintID:           1,
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetSprintWhenTheSprintIDIsZero",
			sprintID:           0,
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetSprintWhenTheRequestMethodIsIncorrect",
			sprintID:           1,
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetSprintWhenTheStatusCodeIsIncorrect",
			sprintID:           1,
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetSprintWhenTheContextIsNil",
			sprintID:           1,
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetSprintWhenTheResponseBodyIsEmpty",
			sprintID:           1,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
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

			service := &SprintService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.sprintID)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestSprintService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *SprintPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateSprintWhenTheParametersAreCorrect",
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "CreateSprintWhenThePayloadIsNotProvided",
			payload:            nil,
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name: "CreateSprintWhenTheRequestMethodIsIncorrect",
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/sprint",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name: "CreateSprintWhenTheStatusCodeIsIncorrect",
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name: "CreateSprintWhenTheContextIsNil",
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint",
			context:            nil,
			wantErr:            true,
		},

		{
			name: "CreateSprintWhenTheResponseBodyIsEmpty",
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint",
			context:            context.Background(),
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

			service := &SprintService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.payload)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestSprintService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		sprintID           int
		payload            *SprintPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:     "UpdateSprintWhenTheParametersAreCorrect",
			sprintID: 1,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:     "UpdateSprintWhenTheSprintIDIsZero",
			sprintID: 0,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "UpdateSprintWhenThePayloadIsNotProvided",
			sprintID:           1,
			payload:            nil,
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:     "UpdateSprintWhenTheRequestMethodIsIncorrect",
			sprintID: 1,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:     "UpdateSprintWhenTheStatusCodeIsIncorrect",
			sprintID: 1,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:     "UpdateSprintWhenTheContextIsNil",
			sprintID: 1,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            nil,
			wantErr:            true,
		},

		{
			name:     "UpdateSprintWhenTheResponseBodyIsEmpty",
			sprintID: 1,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
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

			service := &SprintService{client: mockClient}
			gotResult, gotResponse, err := service.Update(testCase.context, testCase.sprintID, testCase.payload)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestSprintService_Path(t *testing.T) {

	testCases := []struct {
		name               string
		sprintID           int
		payload            *SprintPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:     "PathSprintWhenTheParametersAreCorrect",
			sprintID: 1,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:     "PathSprintWhenTheSprintIDIsZero",
			sprintID: 0,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "PathSprintWhenThePayloadIsNotProvided",
			sprintID:           1,
			payload:            nil,
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:     "PathSprintWhenTheRequestMethodIsIncorrect",
			sprintID: 1,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:     "PathSprintWhenTheStatusCodeIsIncorrect",
			sprintID: 1,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:     "PathSprintWhenTheContextIsNil",
			sprintID: 1,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/get-sprint-by-id.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            nil,
			wantErr:            true,
		},

		{
			name:     "PathSprintWhenTheResponseBodyIsEmpty",
			sprintID: 1,
			payload: &SprintPayloadScheme{
				Name:          "Sprint XX",
				StartDate:     "2015-04-11T15:22:00.000+10:00",
				EndDate:       "2015-04-20T01:22:00.000+10:00",
				OriginBoardID: 4,
				Goal:          "Sprint XX goal",
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
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

			service := &SprintService{client: mockClient}
			gotResult, gotResponse, err := service.Path(testCase.context, testCase.sprintID, testCase.payload)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestSprintService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		sprintID           int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteSprintWhenTheParametersAreCorrect",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "DeleteSprintWhenTheSprintIDIsZero",
			sprintID:           0,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "DeleteSprintWhenTheRequestMethodIsIncorrect",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "DeleteSprintWhenTheStatusCodeIsIncorrect",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "DeleteSprintWhenTheContextIsNil",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            nil,
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

			service := &SprintService{client: mockClient}
			gotResponse, err := service.Delete(testCase.context, testCase.sprintID)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestSprintService_Start(t *testing.T) {

	testCases := []struct {
		name               string
		sprintID           int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "StartSprintWhenTheParametersAreCorrect",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "StartSprintWhenTheSprintIDIsZero",
			sprintID:           0,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "StartSprintWhenTheRequestMethodIsIncorrect",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "StartSprintWhenTheStatusCodeIsIncorrect",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "StartSprintWhenTheContextIsNil",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            nil,
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

			service := &SprintService{client: mockClient}
			gotResponse, err := service.Start(testCase.context, testCase.sprintID)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestSprintService_Close(t *testing.T) {

	testCases := []struct {
		name               string
		sprintID           int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "CloseSprintWhenTheParametersAreCorrect",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "CloseSprintWhenTheSprintIDIsZero",
			sprintID:           0,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "CloseSprintWhenTheRequestMethodIsIncorrect",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "CloseSprintWhenTheStatusCodeIsIncorrect",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "CloseSprintWhenTheContextIsNil",
			sprintID:           1,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/sprint/1",
			context:            nil,
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

			service := &SprintService{client: mockClient}
			gotResponse, err := service.Close(testCase.context, testCase.sprintID)

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

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, apiEndpoint.Path)
				assert.Equal(t, testCase.endpoint, apiEndpoint.Path)

				t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
				assert.Equal(t, gotResponse.Code, testCase.wantHTTPCodeReturn)
			}

		})
	}

}

func TestSprintService_Issues(t *testing.T) {

	testCases := []struct {
		name                string
		sprintID            int
		opts                *IssueOptionScheme
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:     "GetSprintIssuesWhenTheParametersAreCorrect",
			sprintID: 1,
			opts: &IssueOptionScheme{
				JQL:           "project = DUMMY",
				Fields:        []string{"summary", "status"},
				Expand:        []string{"changelogs", "transitions"},
				ValidateQuery: false,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-sprint-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1/issue?expand=changelogs%2Ctransitions&fields=summary%2Cstatus&jql=project+%3D+DUMMY&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:     "GetSprintIssuesWhenTheSprintIDIsNotProvided",
			sprintID: 0,
			opts: &IssueOptionScheme{
				JQL:           "project = DUMMY",
				Fields:        []string{"summary", "status"},
				Expand:        []string{"changelogs"},
				ValidateQuery: false,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-sprint-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1/issue?expand=changelogs&fields=summary%2Cstatus&jql=project+%3D+DUMMY&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:     "GetSprintIssuesWhenTheRequestMethodIsIncorrect",
			sprintID: 1,
			opts: &IssueOptionScheme{
				JQL:           "project = DUMMY",
				Fields:        []string{"summary", "status"},
				Expand:        []string{"changelogs"},
				ValidateQuery: false,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-sprint-issues.json",
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/sprint/1/issue?expand=changelogs&fields=summary%2Cstatus&jql=project+%3D+DUMMY&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:     "GetSprintIssuesWhenTheStatusCodeIsIncorrect",
			sprintID: 1,
			opts: &IssueOptionScheme{
				JQL:           "project = DUMMY",
				Fields:        []string{"summary", "status"},
				Expand:        []string{"changelogs"},
				ValidateQuery: false,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-sprint-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1/issue?expand=changelogs&fields=summary%2Cstatus&jql=project+%3D+DUMMY&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantErr:            true,
		},

		{
			name:     "GetSprintIssuesWhenTheContextIsNil",
			sprintID: 1,
			opts: &IssueOptionScheme{
				JQL:           "project = DUMMY",
				Fields:        []string{"summary", "status"},
				Expand:        []string{"changelogs"},
				ValidateQuery: false,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-sprint-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1/issue?expand=changelogs&fields=summary%2Cstatus&jql=project+%3D+DUMMY&maxResults=50&startAt=0&validateQuery+=false",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:     "GetSprintIssuesWhenTheResponseBodyIsEmpty",
			sprintID: 1,
			opts: &IssueOptionScheme{
				JQL:           "project = DUMMY",
				Fields:        []string{"summary", "status"},
				Expand:        []string{"changelogs"},
				ValidateQuery: false,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/sprint/1/issue?expand=changelogs&fields=summary%2Cstatus&jql=project+%3D+DUMMY&maxResults=50&startAt=0&validateQuery+=false",
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

			service := &SprintService{client: mockClient}
			gotResult, gotResponse, err := service.Issues(testCase.context, testCase.sprintID,
				testCase.opts, testCase.startAt, testCase.maxResults)

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
