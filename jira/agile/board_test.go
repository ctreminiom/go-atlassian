package agile

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestBoardService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		boardID            int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetBoardWhenTheParametersAreCorrect",
			boardID:            1,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/1",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetBoardWhenTheBoardIsNotSet",
			boardID:            0,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardWhenTheRequestMethodIsIncorrect",
			boardID:            1,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardWhenTheContextIsNil",
			boardID:            1,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/1",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetBoardWhenTheResponseStatusCodeIsIncorrect",
			boardID:            1,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardWhenTheResponseBodyIsEmpty",
			boardID:            1,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/1",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.boardID)

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

func TestBoardService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *model.BoardPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateBoardWhenTheParametersAreCorrect",
			payload: &model.BoardPayloadScheme{
				Name:     "DUMMY Board Name",
				Type:     "scrum", //scrum or kanban
				FilterID: 10016,

				// Omit the Location if you want to the board to yourself (location)
				Location: &model.BoardPayloadLocationScheme{
					ProjectKeyOrID: "KP",
					Type:           "project",
				},
			},
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "CreateBoardWhenTheBoardIsNotSet",
			payload:            nil,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board/1",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name: "CreateBoardWhenTheRequestMethodIsIncorrect",
			payload: &model.BoardPayloadScheme{
				Name:     "DUMMY Board Name",
				Type:     "scrum", //scrum or kanban
				FilterID: 10016,

				// Omit the Location if you want to the board to yourself (location)
				Location: &model.BoardPayloadLocationScheme{
					ProjectKeyOrID: "KP",
					Type:           "project",
				},
			},
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name: "CreateBoardWhenTheContextIsNil",
			payload: &model.BoardPayloadScheme{
				Name:     "DUMMY Board Name",
				Type:     "scrum", //scrum or kanban
				FilterID: 10016,

				// Omit the Location if you want to the board to yourself (location)
				Location: &model.BoardPayloadLocationScheme{
					ProjectKeyOrID: "KP",
					Type:           "project",
				},
			},
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board",
			context:            nil,
			wantErr:            true,
		},

		{
			name: "CreateBoardWhenTheResponseStatusCodeIsIncorrect",
			payload: &model.BoardPayloadScheme{
				Name:     "DUMMY Board Name",
				Type:     "scrum", //scrum or kanban
				FilterID: 10016,

				// Omit the Location if you want to the board to yourself (location)
				Location: &model.BoardPayloadLocationScheme{
					ProjectKeyOrID: "KP",
					Type:           "project",
				},
			},
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name: "CreateBoardWhenTheResponseBodyIsEmpty",
			payload: &model.BoardPayloadScheme{
				Name:     "DUMMY Board Name",
				Type:     "scrum", //scrum or kanban
				FilterID: 10016,

				// Omit the Location if you want to the board to yourself (location)
				Location: &model.BoardPayloadLocationScheme{
					ProjectKeyOrID: "KP",
					Type:           "project",
				},
			},
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board",
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

			service := &BoardService{client: mockClient}
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

func TestBoardService_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		opts                *model.GetBoardsOptions
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name: "GetBoardsTheParametersAreCorrect",
			opts: &model.GetBoardsOptions{
				BoardType:               "scrum",
				BoardName:               "board-name",
				ProjectKeyOrID:          "CID",
				AccountIDLocation:       "account-id-sample",
				ProjectIDLocation:       "2345",
				IncludePrivate:          true,
				NegateLocationFiltering: true,
				OrderBy:                 "name",
				Expand:                  "permissions",
				FilterID:                111234,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-boards-by-filter-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board?accountIdLocation=account-id-sample&expand=permissions&filterId=111234&includePrivate=true&maxResults=50&name=board-name&negateLocationFiltering=true&orderBy=name&projectKeyOrId=CID&projectLocation=2345&startAt=0&type=scrum",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name: "GetBoardByFilterWhenTheRequestMethodIsIncorrect",
			opts: &model.GetBoardsOptions{
				BoardType:               "scrum",
				BoardName:               "board-name",
				ProjectKeyOrID:          "CID",
				AccountIDLocation:       "account-id-sample",
				ProjectIDLocation:       "2345",
				IncludePrivate:          true,
				NegateLocationFiltering: true,
				OrderBy:                 "name",
				Expand:                  "permissions",
				FilterID:                111234,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board?accountIdLocation=account-id-sample&expand=permissions&filterId=111234&includePrivate=true&maxResults=50&name=board-name&negateLocationFiltering=true&orderBy=name&projectKeyOrId=CID&projectLocation=2345&startAt=0&type=scrum",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name: "GetBoardByFilterWhenTheContextIsNil",
			opts: &model.GetBoardsOptions{
				BoardType:               "scrum",
				BoardName:               "board-name",
				ProjectKeyOrID:          "CID",
				AccountIDLocation:       "account-id-sample",
				ProjectIDLocation:       "2345",
				IncludePrivate:          true,
				NegateLocationFiltering: true,
				OrderBy:                 "name",
				Expand:                  "permissions",
				FilterID:                111234,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board?accountIdLocation=account-id-sample&expand=permissions&filterId=111234&includePrivate=true&maxResults=50&name=board-name&negateLocationFiltering=true&orderBy=name&projectKeyOrId=CID&projectLocation=2345&startAt=0&type=scrum",
			context:            nil,
			wantErr:            true,
		},

		{
			name: "GetBoardByFilterWhenTheResponseStatusCodeIsIncorrect",
			opts: &model.GetBoardsOptions{
				BoardType:               "scrum",
				BoardName:               "board-name",
				ProjectKeyOrID:          "CID",
				AccountIDLocation:       "account-id-sample",
				ProjectIDLocation:       "2345",
				IncludePrivate:          true,
				NegateLocationFiltering: true,
				OrderBy:                 "name",
				Expand:                  "permissions",
				FilterID:                111234,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board?accountIdLocation=account-id-sample&expand=permissions&filterId=111234&includePrivate=true&maxResults=50&name=board-name&negateLocationFiltering=true&orderBy=name&projectKeyOrId=CID&projectLocation=2345&startAt=0&type=scrum",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name: "GetBoardByFilterWhenTheResponseBodyIsEmpty",
			opts: &model.GetBoardsOptions{
				BoardType:               "scrum",
				BoardName:               "board-name",
				ProjectKeyOrID:          "CID",
				AccountIDLocation:       "account-id-sample",
				ProjectIDLocation:       "2345",
				IncludePrivate:          true,
				NegateLocationFiltering: true,
				OrderBy:                 "name",
				Expand:                  "permissions",
				FilterID:                111234,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board?accountIdLocation=account-id-sample&expand=permissions&filterId=111234&includePrivate=true&maxResults=50&name=board-name&negateLocationFiltering=true&orderBy=name&projectKeyOrId=CID&projectLocation=2345&startAt=0&type=scrum",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.opts,
				testCase.startAt, testCase.maxResults)

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

func TestBoardService_Filter(t *testing.T) {

	testCases := []struct {
		name                          string
		filterID, startAt, maxResults int
		mockFile                      string
		wantHTTPMethod                string
		endpoint                      string
		context                       context.Context
		wantHTTPCodeReturn            int
		wantErr                       bool
	}{
		{
			name:               "GetBoardByFilterWhenTheParametersAreCorrect",
			filterID:           100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-boards-by-filter-id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/filter/100?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetBoardByFilterWhenTheFilterIDIsNotSet",
			filterID:           0,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/filter/100?maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardByFilterWhenTheRequestMethodIsIncorrect",
			filterID:           100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board/filter/100?maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardByFilterWhenTheContextIsNil",
			filterID:           100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/filter/100?maxResults=50&startAt=0",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetBoardByFilterWhenTheResponseStatusCodeIsIncorrect",
			filterID:           100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/filter/100?maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardByFilterWhenTheResponseBodyIsEmpty",
			filterID:           100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/filter/100?maxResults=50&startAt=0",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.Filter(testCase.context, testCase.filterID,
				testCase.startAt, testCase.maxResults)

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

func TestBoardService_Backlog(t *testing.T) {

	testCases := []struct {
		name                         string
		boardID, startAt, maxResults int
		opts                         *model.IssueOptionScheme
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:       "GetBoardBacklogWhenTheParametersAreCorrect",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-backlog.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "GetBoardBacklogWhenTheResponseBodyIsEmpty",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:       "GetBoardBacklogWhenTheValidateQueryOptionIsNotEnabled",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: false,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-backlog.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "GetBoardBacklogWhenTheFilterIDIsNotSet",
			boardID:    0,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog"},
			},
			mockFile:           "../mocks/get-board-backlog.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery=false",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "GetBoardByFilterWhenTheRequestMethodIsIncorrect",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-board-backlog.json",
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "GetBoardBacklogWhenTheContextIsNil",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-board-backlog.json",
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            nil,
			wantErr:            true,
		},

		{
			name:       "GetBoardBacklogWhenTheResponseStatusCodeIsIncorrect",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-backlog.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "GetBoardBacklogWhenTheResponseBodyIsEmpty",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/backlog?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.Backlog(testCase.context, testCase.boardID,
				testCase.startAt, testCase.maxResults, testCase.opts)

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

func TestBoardService_Configuration(t *testing.T) {

	testCases := []struct {
		name               string
		boardID            int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetBoardConfigurationWhenTheParametersAreCorrect",
			boardID:            1,
			mockFile:           "../mocks/get-board-config.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/1/configuration",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:               "GetBoardConfigurationWhenTheBoardIsNotSet",
			boardID:            0,
			mockFile:           "../mocks/get-board-config.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/1/configuration",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardConfigurationWhenTheRequestMethodIsIncorrect",
			boardID:            1,
			mockFile:           "../mocks/get-board-config.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board/1/configuration",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardConfigurationWhenTheContextIsNil",
			boardID:            1,
			mockFile:           "../mocks/get-board-config.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/1/configuration",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetBoardConfigurationWhenTheResponseStatusCodeIsIncorrect",
			boardID:            1,
			mockFile:           "../mocks/get-board-config.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/1/configuration",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardConfigurationWhenTheResponseBodyIsEmpty",
			boardID:            1,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/1/configuration",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.Configuration(testCase.context, testCase.boardID)

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

func TestBoardService_Epics(t *testing.T) {

	testCases := []struct {
		name                         string
		boardID, startAt, maxResults int
		done                         bool
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:               "GetBoardEpicsWhenTheParametersAreCorrect",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-epics.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic?done=false&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetBoardEpicsWhenTheDoneParameterIsSet",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			done:               true,
			mockFile:           "../mocks/get-board-epics.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic?done=true&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetBoardEpicsWhenTheFilterIDIsNotSet",
			boardID:            0,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-epics.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic?done=false&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardEpicsWhenTheRequestMethodIsIncorrect",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-epics.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board/100/epic?done=false&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardEpicsWhenTheContextIsNil",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-epics.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic?done=false&maxResults=50&startAt=0",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetBoardEpicsWhenTheResponseStatusCodeIsIncorrect",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-epics.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic?done=false&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardEpicsWhenTheResponseBodyIsEmpty",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic?done=false&maxResults=50&startAt=0",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.Epics(testCase.context, testCase.boardID,
				testCase.startAt, testCase.maxResults, testCase.done)

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

func TestBoardService_IssuesWithoutEpic(t *testing.T) {

	testCases := []struct {
		name                         string
		boardID, startAt, maxResults int
		opts                         *model.IssueOptionScheme
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:       "GetBoardIssuesWithoutEpicWhenTheParametersAreCorrect",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issue-without-epic.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/none/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "GetBoardIssuesWithoutEpicWhenTheResponseBodyIsEmpty",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/none/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:       "GetBoardIssuesWithoutEpicWhenTheValidateQueryOptionIsNotEnabled",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: false,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile: "../mocks/get-board-issue-without-epic.json",

			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/none/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "GetBoardIssuesWithoutEpicWhenTheFilterIDIsNotSet",
			boardID:    0,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog"},
			},
			mockFile: "../mocks/get-board-issue-without-epic.json",

			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/none/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "GetBoardIssuesWithoutEpicWhenTheRequestMethodIsIncorrect",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-board-issue-without-epic.json",

			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board/100/epic/none/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "GetBoardIssuesWithoutEpicWhenTheContextIsNil",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-board-issue-without-epic.json",

			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/none/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery+=false",
			context:            nil,
			wantErr:            true,
		},

		{
			name:       "GetBoardIssuesWithoutEpicWhenTheResponseStatusCodeIsIncorrect",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile: "../mocks/get-board-issue-without-epic.json",

			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/none/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "GetBoardIssuesWithoutEpicWhenTheResponseBodyIsEmpty",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile: "../mocks/get-board-issue-without-epic.json",

			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/none/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery+=false",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.IssuesWithoutEpic(testCase.context, testCase.boardID,
				testCase.startAt, testCase.maxResults, testCase.opts)

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

func TestBoardService_IssuesByEpic(t *testing.T) {

	testCases := []struct {
		name                                 string
		boardID, epicID, startAt, maxResults int
		opts                                 *model.IssueOptionScheme
		mockFile                             string
		wantHTTPMethod                       string
		endpoint                             string
		context                              context.Context
		wantHTTPCodeReturn                   int
		wantErr                              bool
	}{
		{
			name:       "IssuesByEpicWhenTheParametersAreCorrect",
			boardID:    100,
			epicID:     22,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "IssuesByEpicWhenTheEpicIDIsNotSet",
			boardID:    100,
			epicID:     0,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:       "IssuesByEpicWhenTheValidateQueryOptionIsNotEnabled",
			boardID:    100,
			epicID:     22,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: false,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "IssuesByEpicWhenTheBoardIDIsNotSet",
			boardID:    0,
			epicID:     22,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "IssuesByEpicWhenTheRequestMethodIsIncorrect",
			boardID:    100,
			epicID:     22,
			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-board-issue-by-epic.json",
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board/100/epic/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:    "IssuesByEpicWhenTheContextIsNil",
			boardID: 100,
			epicID:  22,

			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-board-issue-by-epic.json",
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            nil,
			wantErr:            true,
		},

		{
			name:    "IssuesByEpicWhenTheResponseStatusCodeIsIncorrect",
			boardID: 100,
			epicID:  22,

			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "IssuesByEpicWhenTheResponseBodyIsEmpty",
			boardID:    100,
			epicID:     22,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/epic/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.IssuesByEpic(testCase.context, testCase.boardID, testCase.epicID,
				testCase.startAt, testCase.maxResults, testCase.opts)

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

func TestBoardService_Issues(t *testing.T) {

	testCases := []struct {
		name                         string
		boardID, startAt, maxResults int
		opts                         *model.IssueOptionScheme
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:       "GetBoardIssuesWhenTheParametersAreCorrect",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "GetBoardIssuesWhenTheValidateQueryOptionIsNotEnabled",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: false,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issues.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "GetBoardIssuesWhenTheFilterIDIsNotSet",
			boardID:    0,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog"},
			},
			mockFile:           "../mocks/get-board-issues.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "GetBoardIssuesWhenTheRequestMethodIsIncorrect",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-board-issues.json",
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board/100/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "GetBoardIssuesWhenTheContextIsNil",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-board-issues.json",
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            nil,
			wantErr:            true,
		},

		{
			name:       "GetBoardIssuesWhenTheResponseStatusCodeIsIncorrect",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issues.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "GetBoardIssuesWhenTheResponseBodyIsEmpty",
			boardID:    100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.Issues(testCase.context, testCase.boardID,
				testCase.startAt, testCase.maxResults, testCase.opts)

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

func TestBoardService_Projects(t *testing.T) {

	testCases := []struct {
		name                         string
		boardID, startAt, maxResults int
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:               "GetBoardProjectsWhenTheParametersAreCorrect",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-projects.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/project?maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetBoardProjectsWhenTheBoardIDIsNotSet",
			boardID:            0,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-projects.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/project?maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardProjectsWhenTheRequestMethodIsIncorrect",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-projects.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board/100/project?maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardProjectsWhenTheContextIsNil",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-projects.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/project?maxResults=50&startAt=0",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "GetBoardProjectsWhenTheResponseStatusCodeIsIncorrect",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-projects.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/project?maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:               "GetBoardProjectsWhenTheResponseBodyIsEmpty",
			boardID:            100,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/project?maxResults=50&startAt=0",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.Projects(testCase.context, testCase.boardID,
				testCase.startAt, testCase.maxResults)

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

func TestBoardService_IssuesBySprint(t *testing.T) {

	testCases := []struct {
		name                                  string
		sprintID, epicID, startAt, maxResults int
		opts                                  *model.IssueOptionScheme
		mockFile                              string
		wantHTTPMethod                        string
		endpoint                              string
		context                               context.Context
		wantHTTPCodeReturn                    int
		wantErr                               bool
	}{
		{
			name:       "IssuesBySprintWhenTheParametersAreCorrect",
			sprintID:   100,
			epicID:     22,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "IssuesBySprintWhenTheEpicIDIsNotSet",
			sprintID:   100,
			epicID:     0,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:       "IssuesBySprintWhenTheSprintIDIsNotSet",
			sprintID:   0,
			epicID:     100,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:       "IssuesBySprintWhenTheValidateQueryOptionIsNotEnabled",
			sprintID:   100,
			epicID:     22,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: false,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0&validateQuery+=false",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:       "IssuesBySprintWhenTheBoardIDIsNotSet",
			sprintID:   0,
			epicID:     22,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "IssuesBySprintWhenTheRequestMethodIsIncorrect",
			sprintID:   100,
			epicID:     22,
			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-board-issue-by-epic.json",
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodHead,
			endpoint:           "/rest/agile/1.0/board/100/sprint/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:     "IssuesBySprintWhenTheContextIsNil",
			sprintID: 100,
			epicID:   22,

			startAt:    0,
			maxResults: 50,
			mockFile:   "../mocks/get-board-issue-by-epic.json",
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            nil,
			wantErr:            true,
		},

		{
			name:     "IssuesBySprintWhenTheResponseStatusCodeIsIncorrect",
			sprintID: 100,
			epicID:   22,

			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/get-board-issue-by-epic.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:       "IssuesBySprintWhenTheResponseBodyIsEmpty",
			sprintID:   100,
			epicID:     22,
			startAt:    0,
			maxResults: 50,
			opts: &model.IssueOptionScheme{
				JQL:           "project = KP",
				ValidateQuery: true,
				Fields:        []string{"status", "issuetype", "summary"},
				Expand:        []string{"changelog", "metadata"},
			},
			mockFile:           "../mocks/empty-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint/22/issue?expand=changelog%2Cmetadata&fields=status%2Cissuetype%2Csummary&jql=project+%3D+KP&maxResults=50&startAt=0",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.IssuesBySprint(testCase.context, testCase.sprintID, testCase.epicID,
				testCase.startAt, testCase.maxResults, testCase.opts)

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

func TestBoardService_Versions(t *testing.T) {

	testCases := []struct {
		name                         string
		boardID, startAt, maxResults int
		released                     bool
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:               "GetBoardVersionsWhenTheParametersAreCorrect",
			boardID:            100,
			released:           true,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/version?maxResults=50&released=true&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetBoardVersionsWhenTheBoardIDIsNotSet",
			boardID:            0,
			released:           true,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/version?maxResults=50&released=true&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetBoardVersionsWhenTheReleasedIsNotSet",
			boardID:            100,
			released:           false,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/version?maxResults=50&released=false&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetBoardVersionsWhenTheContextIsNil",
			boardID:            100,
			released:           true,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/version?maxResults=50&released=true&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetBoardVersionsWhenTheRequestMethodIsIncorrect",
			boardID:            100,
			released:           true,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-versions.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/agile/1.0/board/100/version?maxResults=50&released=true&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetBoardVersionsWhenTheResponseStatusCodeIsIncorrect",
			boardID:            100,
			released:           true,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-versions.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/version?maxResults=50&released=true&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantErr:            true,
		},

		{
			name:               "GetBoardVersionsWhenTheResponseEmptyIsEmpty",
			boardID:            100,
			released:           true,
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/version?maxResults=50&released=true&startAt=0",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.Versions(testCase.context, testCase.boardID,
				testCase.startAt, testCase.maxResults, testCase.released)

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

func TestBoardService_Move(t *testing.T) {

	testCases := []struct {
		name               string
		boardID            int
		payload            *model.BoardMovementPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:    "MoveBacklogIssueToBoardWhenTheParametersAreCorrect",
			boardID: 100,
			payload: &model.BoardMovementPayloadScheme{
				Issues:          []string{"KP-3"},
				RankBeforeIssue: "",
				RankAfterIssue:  "",
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board/100/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "MoveBacklogIssueToBoardWhenThePayloadIsNotSet",
			boardID:            100,
			payload:            nil,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board/100/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:    "MoveBacklogIssueToBoardWhenTheBoardIDIsNotSet",
			boardID: 0,
			payload: &model.BoardMovementPayloadScheme{
				Issues:          []string{"KP-3"},
				RankBeforeIssue: "",
				RankAfterIssue:  "",
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board/100/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:    "MoveBacklogIssueToBoardWhenTheContextIsNil",
			boardID: 100,
			payload: &model.BoardMovementPayloadScheme{
				Issues:          []string{"KP-3"},
				RankBeforeIssue: "",
				RankAfterIssue:  "",
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board/100/issue",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:    "MoveBacklogIssueToBoardWhenTheRequestMethodIsIncorrect",
			boardID: 100,
			payload: &model.BoardMovementPayloadScheme{
				Issues:          []string{"KP-3"},
				RankBeforeIssue: "",
				RankAfterIssue:  "",
			},
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/agile/1.0/board/100/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:    "MoveBacklogIssueToBoardWhenTheStatusCodeIsIncorrect",
			boardID: 100,
			payload: &model.BoardMovementPayloadScheme{
				Issues:          []string{"KP-3"},
				RankBeforeIssue: "",
				RankAfterIssue:  "",
			},
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board/100/issue",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadGateway,
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

			service := &BoardService{client: mockClient}
			gotResponse, err := service.Move(testCase.context, testCase.boardID, testCase.payload)

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

func TestBoardService_Sprints(t *testing.T) {

	testCases := []struct {
		name                         string
		boardID, startAt, maxResults int
		states                       []string
		mockFile                     string
		wantHTTPMethod               string
		endpoint                     string
		context                      context.Context
		wantHTTPCodeReturn           int
		wantErr                      bool
	}{
		{
			name:               "GetBoardSprintsWhenTheParametersAreCorrect",
			boardID:            100,
			states:             []string{"open", "closed"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-sprints.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint?maxResults=50&startAt=0&state=open%2Cclosed",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetBoardSprintsWhenTheBoardIDIsNotSet",
			boardID:            0,
			states:             []string{"open", "closed"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-sprints.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint?maxResults=50&startAt=0&state=open%2Cclosed",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:           "GetBoardSprintsWhenTheContextIsNil",
			boardID:        100,
			states:         []string{"open", "closed"},
			startAt:        0,
			maxResults:     50,
			mockFile:       "../mocks/get-board-sprints.json",
			wantHTTPMethod: http.MethodGet,
			endpoint:       "/rest/agile/1.0/board/100/sprint?maxResults=50&startAt=0&state=open%2Cclosed",
			context:        nil,
			wantErr:        true,
		},

		{
			name:               "GetBoardSprintsWhenTheRequestMethodIsIncorrect",
			boardID:            100,
			states:             []string{"open", "closed"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-sprints.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/agile/1.0/board/100/sprint?maxResults=50&startAt=0&state=open%2Cclosed",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetBoardSprintsWhenTheStatusCodeIsIncorrect",
			boardID:            100,
			states:             []string{"open", "closed"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/get-board-sprints.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint?maxResults=50&startAt=0&state=open%2Cclosed",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantErr:            true,
		},

		{
			name:               "GetBoardSprintsWhenTheResponseBodyIsEmpty",
			boardID:            100,
			states:             []string{"open", "closed"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "../mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/agile/1.0/board/100/sprint?maxResults=50&startAt=0&state=open%2Cclosed",
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

			service := &BoardService{client: mockClient}
			gotResult, gotResponse, err := service.Sprints(testCase.context, testCase.boardID,
				testCase.startAt, testCase.maxResults, testCase.states)

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

func TestBoardService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		boardID            int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteBoardWhenTheParametersAreCorrect",
			boardID:            1124,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/board/1124",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteBoardWhenTheContextIsNotProvided",
			boardID:            1124,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/board/1124",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteBoardWhenTheResponseStatusIsNotValid",
			boardID:            1124,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/board/1124",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteBoardWhenTheBoardIDIsNotProvided",
			boardID:            0,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/agile/1.0/board/1124",
			context:            context.Background(),
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

			service := &BoardService{client: mockClient}

			gotResponse, err := service.Delete(testCase.context, testCase.boardID)

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
