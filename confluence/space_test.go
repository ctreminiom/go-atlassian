package confluence

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestSpaceService_Gets(t *testing.T) {

	testCases := []struct {
		name                string
		options             *GetSpacesOptionScheme
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name: "GetSpacesWhenTheParametersAreCorrect",
			options: &GetSpacesOptionScheme{
				SpaceKeys:       []string{"DUMMY", "TEST"},
				SpaceIDs:        []int{1111, 2222, 3333},
				SpaceType:       "global",
				Status:          "archived",
				Labels:          []string{"label-09", "label-02"},
				Favorite:        true,
				FavoriteUserKey: "DUMMY",
				Expand:          []string{"operations"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-spaces.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space?expand=operations&favorite=true&favouriteUserKey=DUMMY&label=label-09%2Clabel-02&limit=50&spaceId=1111&spaceId=2222&spaceId=3333&spaceKey=DUMMY%2CTEST&start=0&status=archived&type=global",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name: "GetSpacesWhenTheContextIsNotProvided",
			options: &GetSpacesOptionScheme{
				SpaceKeys:       []string{"DUMMY", "TEST"},
				SpaceIDs:        []int{1111, 2222, 3333},
				SpaceType:       "global",
				Status:          "archived",
				Labels:          []string{"label-09", "label-02"},
				Favorite:        true,
				FavoriteUserKey: "DUMMY",
				Expand:          []string{"operations"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-spaces.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space?expand=operations&favorite=true&favouriteUserKey=DUMMY&label=label-09%2Clabel-02&limit=50&spaceId=1111&spaceId=2222&spaceId=3333&spaceKey=DUMMY%2CTEST&start=0&status=archived&type=global",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetSpacesWhenTheRequestMethodIsIncorrect",
			options: &GetSpacesOptionScheme{
				SpaceKeys:       []string{"DUMMY", "TEST"},
				SpaceIDs:        []int{1111, 2222, 3333},
				SpaceType:       "global",
				Status:          "archived",
				Labels:          []string{"label-09", "label-02"},
				Favorite:        true,
				FavoriteUserKey: "DUMMY",
				Expand:          []string{"operations"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-spaces.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/space?expand=operations&favorite=true&favouriteUserKey=DUMMY&label=label-09%2Clabel-02&limit=50&spaceId=1111&spaceId=2222&spaceId=3333&spaceKey=DUMMY%2CTEST&start=0&status=archived&type=global",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name: "GetSpacesWhenTheStatusCodeIsIncorrect",
			options: &GetSpacesOptionScheme{
				SpaceKeys:       []string{"DUMMY", "TEST"},
				SpaceIDs:        []int{1111, 2222, 3333},
				SpaceType:       "global",
				Status:          "archived",
				Labels:          []string{"label-09", "label-02"},
				Favorite:        true,
				FavoriteUserKey: "DUMMY",
				Expand:          []string{"operations"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-spaces.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space?expand=operations&favorite=true&favouriteUserKey=DUMMY&label=label-09%2Clabel-02&limit=50&spaceId=1111&spaceId=2222&spaceId=3333&spaceKey=DUMMY%2CTEST&start=0&status=archived&type=global",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "GetSpacesWhenTheResponseBodyIsEmpty",
			options: &GetSpacesOptionScheme{
				SpaceKeys:       []string{"DUMMY", "TEST"},
				SpaceIDs:        []int{1111, 2222, 3333},
				SpaceType:       "global",
				Status:          "archived",
				Labels:          []string{"label-09", "label-02"},
				Favorite:        true,
				FavoriteUserKey: "DUMMY",
				Expand:          []string{"operations"},
			},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space?expand=operations&favorite=true&favouriteUserKey=DUMMY&label=label-09%2Clabel-02&limit=50&spaceId=1111&spaceId=2222&spaceId=3333&spaceKey=DUMMY%2CTEST&start=0&status=archived&type=global",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &SpaceService{client: mockClient}

			gotResult, gotResponse, err := service.Gets(
				testCase.context,
				testCase.options,
				testCase.startAt,
				testCase.maxResults,
			)

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

func TestSpaceService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		payload            *CreateSpaceScheme
		private            bool
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateSpaceWhenTheParametersAreCorrect",
			payload: &CreateSpaceScheme{
				Key:  "DUM",
				Name: "Dum Confluence Space",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Confluence Space Description Sample",
						Representation: "plain",
					},
				},
				AnonymousAccess:  true,
				UnlicensedAccess: false,
			},
			private:            false,
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},

		{
			name: "CreateSpaceWhenTheSpaceKeyIsNotProvided",
			payload: &CreateSpaceScheme{
				Key:  "",
				Name: "Dum Confluence Space",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Confluence Space Description Sample",
						Representation: "plain",
					},
				},
				AnonymousAccess:  true,
				UnlicensedAccess: false,
			},
			private:            false,
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateSpaceWhenTheSpaceNameIsNotProvided",
			payload: &CreateSpaceScheme{
				Key:  "DUM",
				Name: "",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Confluence Space Description Sample",
						Representation: "plain",
					},
				},
				AnonymousAccess:  true,
				UnlicensedAccess: false,
			},
			private:            false,
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name:               "CreateSpaceWhenThePayloadIsNotProvided",
			payload:            nil,
			private:            false,
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateSpaceWhenTheContextIsNotProvided",
			payload: &CreateSpaceScheme{
				Key:  "DUM",
				Name: "Dum Confluence Space",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Confluence Space Description Sample",
						Representation: "plain",
					},
				},
				AnonymousAccess:  true,
				UnlicensedAccess: false,
			},
			private:            false,
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space",
			context:            nil,
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateSpaceWhenTheRequestMethodIsIncorrect",
			payload: &CreateSpaceScheme{
				Key:  "DUM",
				Name: "Dum Confluence Space",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Confluence Space Description Sample",
						Representation: "plain",
					},
				},
				AnonymousAccess:  true,
				UnlicensedAccess: false,
			},
			private:            false,
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/space",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            true,
		},

		{
			name: "CreateSpaceWhenTheStatusCodeIsIncorrect",
			payload: &CreateSpaceScheme{
				Key:  "DUM",
				Name: "Dum Confluence Space",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Confluence Space Description Sample",
						Representation: "plain",
					},
				},
				AnonymousAccess:  true,
				UnlicensedAccess: false,
			},
			private:            false,
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name: "CreateSpaceWhenTheSpaceRequestedIsPrivate",
			payload: &CreateSpaceScheme{
				Key:  "DUM",
				Name: "Dum Confluence Space",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Confluence Space Description Sample",
						Representation: "plain",
					},
				},
				AnonymousAccess:  true,
				UnlicensedAccess: false,
			},
			private:            true,
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/_private",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusCreated,
			wantErr:            false,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &SpaceService{client: mockClient}

			gotResult, gotResponse, err := service.Create(testCase.context, testCase.payload, testCase.private)

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

func TestSpaceService_Get(t *testing.T) {

	testCases := []struct {
		name               string
		spaceKey           string
		expand             []string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetSpaceWhenTheParametersAreCorrect",
			spaceKey:           "DUMMY",
			expand:             []string{"childtypes.all", "operations"},
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY?expand=childtypes.all%2Coperations",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetSpaceWhenTheSpaceKeyIsNotProvided",
			spaceKey:           "",
			expand:             []string{"childtypes.all", "operations"},
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY?expand=childtypes.all%2Coperations",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSpaceWhenTheContextIsNotProvided",
			spaceKey:           "DUMMY",
			expand:             []string{"childtypes.all", "operations"},
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY?expand=childtypes.all%2Coperations",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSpaceWhenTheRequestMethodIsIncorrect",
			spaceKey:           "DUMMY",
			expand:             []string{"childtypes.all", "operations"},
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/space/DUMMY?expand=childtypes.all%2Coperations",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSpaceWhenTheStatusCodeIsIncorrect",
			spaceKey:           "DUMMY",
			expand:             []string{"childtypes.all", "operations"},
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY?expand=childtypes.all%2Coperations",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetSpaceWhenTheResponseBodyIsEmpty",
			spaceKey:           "DUMMY",
			expand:             []string{"childtypes.all", "operations"},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY?expand=childtypes.all%2Coperations",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &SpaceService{client: mockClient}

			gotResult, gotResponse, err := service.Get(testCase.context, testCase.spaceKey, testCase.expand)

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

func TestSpaceService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		spaceKey           string
		payload            *UpdateSpaceScheme
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:     "UpdateSpaceWhenTheParametersAreCorrect",
			spaceKey: "DUMMY",
			payload: &UpdateSpaceScheme{
				Name: "DUMMY Space - Updated",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Dummy Space - Description - Updated",
						Representation: "plain",
					},
				},
				Homepage: &UpdateSpaceHomepageScheme{ID: "65798145"},
			},
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:     "UpdateSpaceWhenTheSpaceKeyIsNotProvided",
			spaceKey: "",
			payload: &UpdateSpaceScheme{
				Name: "DUMMY Space - Updated",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Dummy Space - Description - Updated",
						Representation: "plain",
					},
				},
				Homepage: &UpdateSpaceHomepageScheme{ID: "65798145"},
			},
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "UpdateSpaceWhenThePayloadIsNotProvided",
			spaceKey:           "DUMMY",
			payload:            nil,
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:     "UpdateSpaceWhenTheContextIsNotProvided",
			spaceKey: "DUMMY",
			payload: &UpdateSpaceScheme{
				Name: "DUMMY Space - Updated",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Dummy Space - Description - Updated",
						Representation: "plain",
					},
				},
				Homepage: &UpdateSpaceHomepageScheme{ID: "65798145"},
			},
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:     "UpdateSpaceWhenTheRequestMethodIsIncorrect",
			spaceKey: "DUMMY",
			payload: &UpdateSpaceScheme{
				Name: "DUMMY Space - Updated",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Dummy Space - Description - Updated",
						Representation: "plain",
					},
				},
				Homepage: &UpdateSpaceHomepageScheme{ID: "65798145"},
			},
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:     "UpdateSpaceWhenTheStatusCodeIsIncorrect",
			spaceKey: "DUMMY",
			payload: &UpdateSpaceScheme{
				Name: "DUMMY Space - Updated",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Dummy Space - Description - Updated",
						Representation: "plain",
					},
				},
				Homepage: &UpdateSpaceHomepageScheme{ID: "65798145"},
			},
			mockFile:           "./mocks/get-space.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:     "UpdateSpaceWhenTheResponseBodyIsEmpty",
			spaceKey: "DUMMY",
			payload: &UpdateSpaceScheme{
				Name: "DUMMY Space - Updated",
				Description: &CreateSpaceDescriptionScheme{
					Plain: &CreateSpaceDescriptionPlainScheme{
						Value:          "Dummy Space - Description - Updated",
						Representation: "plain",
					},
				},
				Homepage: &UpdateSpaceHomepageScheme{ID: "65798145"},
			},
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &SpaceService{client: mockClient}

			gotResult, gotResponse, err := service.Update(testCase.context, testCase.spaceKey, testCase.payload)

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

func TestSpaceService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		spaceKey           string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteSpaceWhenTheParametersAreCorrect",
			spaceKey:           "DUMMY",
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            false,
		},

		{
			name:               "DeleteSpaceWhenTheSpaceKeyIsNotProvided",
			spaceKey:           "",
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},

		{
			name:               "DeleteSpaceWhenTheRequestMethodIsIncorrect",
			spaceKey:           "DUMMY",
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},

		{
			name:               "DeleteSpaceWhenTheStatusCodeIsIncorrect",
			spaceKey:           "DUMMY",
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeleteSpaceWhenTheContextIsNotProvided",
			spaceKey:           "DUMMY",
			mockFile:           "./mocks/get-long-task.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/wiki/rest/api/space/DUMMY",
			context:            nil,
			wantHTTPCodeReturn: http.StatusAccepted,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &SpaceService{client: mockClient}

			gotResult, gotResponse, err := service.Delete(testCase.context, testCase.spaceKey)

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

func TestSpaceService_Content(t *testing.T) {

	testCases := []struct {
		name                string
		spaceKey, depth     string
		expand              []string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetSpaceContentWhenTheParametersAreCorrect",
			spaceKey:           "DUMMY",
			depth:              "all",
			expand:             []string{"operations"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY/content?depth=all&expand=operations&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetSpaceContentWhenTheSpaceKeyIsNotProvided",
			spaceKey:           "",
			depth:              "all",
			expand:             []string{"operations"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY/content?depth=all&expand=operations&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSpaceContentWhenTheContextIsNotProvided",
			spaceKey:           "DUMMY",
			depth:              "all",
			expand:             []string{"operations"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY/content?depth=all&expand=operations&limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSpaceContentWhenTheRequestMethodIsIncorrect",
			spaceKey:           "DUMMY",
			depth:              "all",
			expand:             []string{"operations"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-content-children.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/wiki/rest/api/space/DUMMY/content?depth=all&expand=operations&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSpaceContentWhenTheResponseBodyIsEmpty",
			spaceKey:           "DUMMY",
			depth:              "all",
			expand:             []string{"operations"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY/content?depth=all&expand=operations&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &SpaceService{client: mockClient}

			gotResult, gotResponse, err := service.Content(testCase.context, testCase.spaceKey, testCase.depth,
				testCase.expand, testCase.startAt, testCase.maxResults)

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

func TestSpaceService_ContentByType(t *testing.T) {

	testCases := []struct {
		name                string
		spaceKey, depth     string
		expand              []string
		contentType         string
		startAt, maxResults int
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "GetSpaceContentByTypeWhenTheParametersAreCorrect",
			spaceKey:           "DUMMY",
			depth:              "all",
			contentType:        "page",
			expand:             []string{"operations", "restrictions"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY/content/page?depth=all&expand=operations%2Crestrictions&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetSpaceContentByTypeWhenTheSpaceKeyIsNotProvided",
			spaceKey:           "",
			depth:              "all",
			contentType:        "page",
			expand:             []string{"operations", "restrictions"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY/content/page?depth=all&expand=operations%2Crestrictions&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSpaceContentByTypeWhenTheRequestMethodIsIncorrect",
			spaceKey:           "DUMMY",
			depth:              "all",
			contentType:        "page",
			expand:             []string{"operations", "restrictions"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/wiki/rest/api/space/DUMMY/content/page?depth=all&expand=operations%2Crestrictions&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSpaceContentByTypeWhenTheStatusCodeIsIncorrect",
			spaceKey:           "DUMMY",
			depth:              "all",
			contentType:        "page",
			expand:             []string{"operations", "restrictions"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY/content/page?depth=all&expand=operations%2Crestrictions&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetSpaceContentByTypeWhenTheContextIsNotProvided",
			spaceKey:           "DUMMY",
			depth:              "all",
			contentType:        "page",
			expand:             []string{"operations", "restrictions"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/get-contents.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY/content/page?depth=all&expand=operations%2Crestrictions&limit=50&start=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSpaceContentByTypeWhenTheResponseBodyIsEmpty",
			spaceKey:           "DUMMY",
			depth:              "all",
			contentType:        "page",
			expand:             []string{"operations", "restrictions"},
			startAt:            0,
			maxResults:         50,
			mockFile:           "./mocks/empty-json.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/wiki/rest/api/space/DUMMY/content/page?depth=all&expand=operations%2Crestrictions&limit=50&start=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

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

			service := &SpaceService{client: mockClient}

			gotResult, gotResponse, err := service.ContentByType(testCase.context, testCase.spaceKey, testCase.contentType, testCase.depth,
				testCase.expand, testCase.startAt, testCase.maxResults)

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
