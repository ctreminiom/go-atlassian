package agile

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira/agile/mocksV2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestProductService_Insert(t *testing.T) {

	repo := mocksV2.NewBoard(t)
	repo.On("Get", context.Background(), 1).Return(nil, nil, nil).Once()

	board, _, err := repo.Get(context.Background(), 1)

	assert.Nil(t, err)
	assert.Nil(t, board)

	/*

		repo := &mocksV2.NewBoardT(t)
		repo.On("Add", mock.AnythingOfType("models.Product")).
			Return(nil).
			Once()

		service := services.NewProductService(repo)

		err := service.Insert("2f1afe98-63c4-4f59-bcaf-1df835602bdb", models.InsertProductDTO{
			Name:  "Macbook",
			Price: 20500,
			Stock: 10,
		})

		assert.Nil(t, err)
	*/
}

func Test_BoardService_Get(t *testing.T) {

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
			name:               "when the parameters are correct",
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

			service := &BoardService{c: mockClient, version: "1.0"}
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
