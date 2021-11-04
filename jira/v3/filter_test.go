package v3

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"strconv"
	"testing"
)

func TestFilterService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		payload            *FilterPayloadScheme
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateFilterWhenThePayloadIsCorrect",
			payload: &FilterPayloadScheme{
				Name:        "Filter #a5fd86b9-4fef-44c1-8ce4-8d1a63e806e1",
				Description: "Filter's description",
				JQL:         "issuetype = Bug",
				Favorite:    false,
			},
			mockFile:           "./mocks/create_filter.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},
		{
			name:               "CreateFilterWhenThePayloadEmpty",
			payload:            nil,
			mockFile:           "./mocks/create_filter.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "CreateFilterWhenTheContextIsNil",
			payload: &FilterPayloadScheme{
				Name:        "Filter #a5fd86b9-4fef-44c1-8ce4-8d1a63e806e1",
				Description: "Filter's description",
				JQL:         "issuetype = Bug",
				Favorite:    false,
			},
			mockFile:           "./mocks/create_filter.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "CreateFilterWhenTheResponseBodyLengthIsZero",
			payload: &FilterPayloadScheme{
				Name:        "Filter #a5fd86b9-4fef-44c1-8ce4-8d1a63e806e1",
				Description: "Filter's description",
				JQL:         "issuetype = Bug",
				Favorite:    false,
			},
			mockFile:           "",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "CreateFilterWhenTheResponseBodyHasADifferentFormat",
			payload: &FilterPayloadScheme{
				Name:        "Filter #a5fd86b9-4fef-44c1-8ce4-8d1a63e806e1",
				Description: "Filter's description",
				JQL:         "issuetype = Bug",
				Favorite:    false,
			},
			mockFile:           "./mocks/invalid-json.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},
		{
			name: "CreateFilterWhenTheStatusResponseCodeIsInvalid",
			payload: &FilterPayloadScheme{
				Name:        "Filter #a5fd86b9-4fef-44c1-8ce4-8d1a63e806e1",
				Description: "Filter's description",
				JQL:         "issuetype = Bug",
				Favorite:    false,
			},
			mockFile:           "./mocks/create_filter.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name: "CreateFilterWhenTheRequestMethodIsInvalid",
			payload: &FilterPayloadScheme{
				Name:        "Filter #a5fd86b9-4fef-44c1-8ce4-8d1a63e806e1",
				Description: "Filter's description",
				JQL:         "issuetype = Bug",
				Favorite:    false,
			},
			mockFile:           "./mocks/create_filter.json",
			wantHTTPMethod:     http.MethodConnect,
			endpoint:           "/rest/api/3/filter",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusMethodNotAllowed,
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

			service := &FilterService{client: mockClient}

			gotResult, gotResponse, err := service.Create(testCase.context, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)

				if gotResponse != nil {
					t.Logf("HTTP Code Wanted: %v, HTTP Code Returned: %v", testCase.wantHTTPCodeReturn, gotResponse.Code)
					assert.Equal(t, testCase.wantHTTPCodeReturn, gotResponse.Code)
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

func TestFilterService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		filterID           int
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteFilterWhenTheIDIsCorrect",
			filterID:           1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/filter/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},
		{
			name:               "DeleteFilterWhenTheIDIsIncorrect",
			filterID:           2,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/filter/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteFilterWhenTheContextIsNil",
			filterID:           1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/filter/1",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteFilterWhenTheHTTPMethodIsIncorrect",
			filterID:           1,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},
		{
			name:               "DeleteFilterWhenTheHTTPStatusCodeIsIncorrect",
			filterID:           1,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/filter/1",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusMethodNotAllowed,
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

			service := &FilterService{client: mockClient}

			gotResponse, err := service.Delete(testCase.context, testCase.filterID)

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

func TestFilterService_Favorite(t *testing.T) {

	testCases := []struct {
		name               string
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetFavoriteFiltersWhenTheIsCorrect",
			mockFile:           "./mocks/get_favorites_filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/favourite",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetFavoriteFiltersWhenTheContextIsNil",
			mockFile:           "./mocks/get_favorites_filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/favourite",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetFavoriteFiltersWhenRequestMethodIsIncorrect",
			mockFile:           "./mocks/get_favorites_filters.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/favourite",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetFavoriteFiltersWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get_favorites_filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/favourite",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetFavoriteFiltersWhenTheResponseBodyLengthIsZero",
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/favourite",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetFavoriteFiltersWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/favourite",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetFavoriteFiltersWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get_favorites_filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/2/filter/favourites",
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

			service := &FilterService{client: mockClient}
			gotResult, gotResponse, err := service.Favorite(testCase.context)

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

func TestFilterService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		filterID           int
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetFilterWhenTheIDCorrect",
			filterID:           10000,
			mockFile:           "./mocks/get_filter_by_id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetFilterWhenTheIDCorrectAndHasExpandValues",
			filterID:           10000,
			expand:             []string{"sharedUsers", "subscriptions"},
			mockFile:           "./mocks/get_filter_by_id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000?expand=sharedUsers%2Csubscriptions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetFilterWhenTheContextIsNil",
			filterID:           10000,
			mockFile:           "./mocks/get_filter_by_id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetFilterWhenTheIDIsIncorrect",
			filterID:           10001,
			mockFile:           "./mocks/get_filter_by_id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetFilterWhenTheRequestMethodIsIncorrect",
			filterID:           10000,
			mockFile:           "./mocks/get_filter_by_id.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetFilterWhenTheStatusCodeIsIncorrect",
			filterID:           10000,
			mockFile:           "./mocks/get_filter_by_id.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetFilterWhenTheResponseBodyLengthIsZero",
			filterID:           10000,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetFilterWhenTheResponseBodyHasADifferentFormat",
			filterID:           10000,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000",
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

			service := &FilterService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.filterID, testCase.expand)

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

				filterIDAsInt, err := strconv.Atoi(gotResult.ID)
				if err != nil {
					t.Fatal(err)
				}

				t.Logf("Filter ID Wanted: %v, Filter ID Returned: %v", testCase.filterID, gotResult.ID)
				assert.Equal(t, testCase.filterID, filterIDAsInt)
			}

		})
	}
}

func TestFilterService_My(t *testing.T) {

	testCases := []struct {
		name               string
		favorites          bool
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetMyFiltersWhenTheIsCorrect",
			mockFile:           "./mocks/get_favorites_filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/my",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetMyFiltersWhenTheIsCorrectAndTheFavoritesIsSelected",
			mockFile:           "./mocks/get_favorites_filters.json",
			favorites:          true,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/my?includeFavourites=true",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetMyFiltersWhenTheIsCorrectAndTheExpandsIsSelected",
			mockFile:           "./mocks/get_favorites_filters.json",
			favorites:          false,
			expand:             []string{"sharedUsers", "subscriptions"},
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/my?expand=sharedUsers%2Csubscriptions",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "GetMyFiltersWhenTheContextIsNil",
			mockFile:           "./mocks/get_favorites_filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/my",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetMyFiltersWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get_favorites_filters.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/my",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetMyFiltersWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get_favorites_filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/my",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "GetMyFiltersWhenTheResponseBodyLengthIsZero",
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/my",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name:               "GetMyFiltersWhenTheResponseBodyHasADifferentFormat",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/my",
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

			service := &FilterService{client: mockClient}
			gotResult, gotResponse, err := service.My(testCase.context, testCase.favorites, testCase.expand)

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

func TestFilterService_Search(t *testing.T) {

	testCases := []struct {
		name               string
		options            *FilterSearchOptionScheme
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
			name: "SearchFiltersWhenTheFilterNameIsSet",
			options: &FilterSearchOptionScheme{
				Name: "Lists all open bugs",
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?filterName=Lists+all+open+bugs&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name: "SearchFiltersWhenTheAccountIDIsSet",
			options: &FilterSearchOptionScheme{
				AccountID: "XXXXXXXXXXXXXXXXXXXX",
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?accountId=XXXXXXXXXXXXXXXXXXXX&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name: "SearchFiltersWhenTheGroupIsSet",
			options: &FilterSearchOptionScheme{
				Group: "jira-users",
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?groupname=jira-users&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name: "SearchFiltersWhenTheProjectIDIsSet",
			options: &FilterSearchOptionScheme{
				ProjectID: 111,
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?maxResults=50&projectId=111&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name: "SearchFiltersWhenTheFiltersIDsIsSet",
			options: &FilterSearchOptionScheme{
				IDs: []int{1, 2, 3, 4},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?id=1&id=2&id=3&id=4&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name: "SearchFiltersWhenTheOrderIsSet",
			options: &FilterSearchOptionScheme{
				OrderBy: "name",
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?maxResults=50&orderBy=name&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name: "SearchFiltersWhenTheExpandIsSet",
			options: &FilterSearchOptionScheme{
				Expand: []string{"description", "favorite", "jql", "owner", "viewUrl"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?expand=description%2Cfavorite%2Cjql%2Cowner%2CviewUrl&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "SearchFiltersWhenThePayloadIsNil",
			options:            nil,
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?expand=description%2Cfavorite%2Cjql%2Cowner%2CviewUrl&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "SearchFiltersWhenTheContextIsNil",
			options: &FilterSearchOptionScheme{
				Group: "jira-users",
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?groupname=jira-users&maxResults=50&startAt=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "SearchFiltersWhenTheRequestMethodIsIncorrect",
			options: &FilterSearchOptionScheme{
				Group: "jira-users",
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/filter/search?groupname=jira-users&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "SearchFiltersWhenTheStatusCodeIsIncorrect",
			options: &FilterSearchOptionScheme{
				Group: "jira-users",
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/search-filters.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?groupname=jira-users&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name: "SearchFiltersWhenTheResponseBodyLengthIsZero",
			options: &FilterSearchOptionScheme{
				Group: "jira-users",
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?groupname=jira-users&maxResults=50&startAt=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "SearchFiltersWhenTheResponseBodyHasADifferentFormat",
			options: &FilterSearchOptionScheme{
				Group: "jira-users",
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/search?groupname=jira-users&maxResults=50&startAt=0",
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

			service := &FilterService{client: mockClient}
			gotResult, gotResponse, err := service.Search(testCase.context, testCase.options, testCase.startAt, testCase.maxResults)

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

func TestFilterService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		filterID           int
		payload            *FilterPayloadScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "UpdateFilterWhenTheIDIsCorrect",
			payload: &FilterPayloadScheme{
				Name: "All Open Bugs",
				JQL:  "type = Bug and resolution is empty",
			},
			filterID:           10000,
			mockFile:           "./mocks/update-filter.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},
		{
			name: "UpdateFilterWhenTheIDIsIncorrect",
			payload: &FilterPayloadScheme{
				Name: "All Open Bugs",
				JQL:  "type = Bug and resolution is empty",
			},
			filterID:           10001,
			mockFile:           "./mocks/update-filter.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name:               "UpdateFilterWhenThePayloadIsNil",
			payload:            nil,
			filterID:           10000,
			mockFile:           "./mocks/update-filter.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "UpdateFilterWhenTheContextIsNil",
			payload: &FilterPayloadScheme{
				Name: "All Open Bugs",
				JQL:  "type = Bug and resolution is empty",
			},
			filterID:           10000,
			mockFile:           "./mocks/update-filter.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/10000",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "UpdateFilterWhenTheRequestMethodIsIncorrect",
			payload: &FilterPayloadScheme{
				Name: "All Open Bugs",
				JQL:  "type = Bug and resolution is empty",
			},
			filterID:           10000,
			mockFile:           "./mocks/update-filter.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "UpdateFilterWhenTheStatusCodeIsIncorrect",
			payload: &FilterPayloadScheme{
				Name: "All Open Bugs",
				JQL:  "type = Bug and resolution is empty",
			},
			filterID:           10000,
			mockFile:           "./mocks/update-filter.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
		{
			name: "UpdateFilterWhenTheResponseBodyLengthIsZero",
			payload: &FilterPayloadScheme{
				Name: "All Open Bugs",
				JQL:  "type = Bug and resolution is empty",
			},
			filterID:           10000,
			mockFile:           "",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/10000",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
		{
			name: "UpdateFilterWhenTheResponseBodyHasADifferentFormat",
			payload: &FilterPayloadScheme{
				Name: "All Open Bugs",
				JQL:  "type = Bug and resolution is empty",
			},
			filterID:           10000,
			mockFile:           "./mocks/empty_json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/filter/10000",
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

			service := &FilterService{client: mockClient}
			gotResult, gotResponse, err := service.Update(testCase.context, testCase.filterID, testCase.payload)

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
