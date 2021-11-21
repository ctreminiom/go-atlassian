package admin

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestSCIMUserService_Create(t *testing.T) {

	testCases := []struct {
		name                           string
		directoryID                    string
		payload                        *model.SCIMUserScheme
		attributes, excludedAttributes []string
		mockFile                       string
		wantHTTPMethod                 string
		endpoint                       string
		context                        context.Context
		wantHTTPCodeReturn             int
		wantErr                        bool
	}{
		{
			name:        "CreateSCIMUserWhenTheParametersAreCorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:        "CreateSCIMUserWhenTheAttributesAndExcludedAttributesAreNotSet",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         nil,
			excludedAttributes: nil,
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "CreateSCIMUserWhenThePayloadIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload:            nil,
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CreateSCIMUserWhenTheDirectoryIDIsNotSet",
			directoryID: "",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CreateSCIMUserWhenTheRequestMethodIsIncorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CreateSCIMUserWhenTheStatusCodeIsIncorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:        "CreateSCIMUserWhenTheContextIsNil",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CreateSCIMUserWhenTheEndpointIsEmpty",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CreateSCIMUserWhenTheResponseBodyIsEmpty",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
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

			service := &SCIMUserService{client: mockClient}
			gotResult, gotResponse, err := service.Create(testCase.context, testCase.directoryID,
				testCase.payload, testCase.attributes, testCase.excludedAttributes)

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

				t.Log("User ID", gotResult.ID)

				for index, email := range gotResult.Emails {
					t.Log("Email #", index, email.Value, email.Type, email.Primary)
				}

				t.Log("User Status: ", gotResult.Active)
				t.Log("User Department: ", gotResult.Department)
				t.Log("User Title: ", gotResult.Title)
			}

		})
	}

}

func TestSCIMUserService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		directoryID        string
		opts               *model.SCIMUserGetsOptionsScheme
		startIndex         int
		count              int
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:        "CreateSCIMUserWhenTheParametersAreCorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			opts: &model.SCIMUserGetsOptionsScheme{
				Attributes:         []string{"userName", "emails.value"},
				ExcludedAttributes: []string{"timezone", "department"},
				Filter:             "users",
			},
			startIndex:         0,
			count:              50,
			mockFile:           "./mocks/scim-get-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&count=50&excludedAttributes=timezone%2Cdepartment&filter=users&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:        "CreateSCIMUserWhenTheDirectoryIDIsNotSet",
			directoryID: "",
			opts: &model.SCIMUserGetsOptionsScheme{
				Attributes:         []string{"userName", "emails.value"},
				ExcludedAttributes: []string{"timezone", "department"},
				Filter:             "users",
			},
			startIndex:         0,
			count:              50,
			mockFile:           "./mocks/scim-get-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&count=50&excludedAttributes=timezone%2Cdepartment&filter=users&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "CreateSCIMUserWhenTheOptionsAreNotSet",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			opts:               nil,
			startIndex:         0,
			count:              50,
			mockFile:           "./mocks/scim-get-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?count=50&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:        "CreateSCIMUserWhenTheRequestMethodIsIncorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			opts: &model.SCIMUserGetsOptionsScheme{
				Attributes:         []string{"userName", "emails.value"},
				ExcludedAttributes: []string{"timezone", "department"},
				Filter:             "users",
			},
			startIndex:         0,
			count:              50,
			mockFile:           "./mocks/scim-get-users.json",
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&count=50&excludedAttributes=timezone%2Cdepartment&filter=users&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CreateSCIMUserWhenTheStatusCodeIsIncorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			opts: &model.SCIMUserGetsOptionsScheme{
				Attributes:         []string{"userName", "emails.value"},
				ExcludedAttributes: []string{"timezone", "department"},
				Filter:             "users",
			},
			startIndex:         0,
			count:              50,
			mockFile:           "./mocks/scim-get-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&count=50&excludedAttributes=timezone%2Cdepartment&filter=users&startIndex=0",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:        "CreateSCIMUserWhenTheContextIsNil",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			opts: &model.SCIMUserGetsOptionsScheme{
				Attributes:         []string{"userName", "emails.value"},
				ExcludedAttributes: []string{"timezone", "department"},
				Filter:             "users",
			},
			startIndex:         0,
			count:              50,
			mockFile:           "./mocks/scim-get-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&count=50&excludedAttributes=timezone%2Cdepartment&filter=users&startIndex=0",
			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CreateSCIMUserWhenTheEndpointIsEmpty",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			opts: &model.SCIMUserGetsOptionsScheme{
				Attributes:         []string{"userName", "emails.value"},
				ExcludedAttributes: []string{"timezone", "department"},
				Filter:             "users",
			},
			startIndex:         0,
			count:              50,
			mockFile:           "./mocks/scim-get-users.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "CreateSCIMUserWhenTheResponseBodyIsEmpty",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			opts: &model.SCIMUserGetsOptionsScheme{
				Attributes:         []string{"userName", "emails.value"},
				ExcludedAttributes: []string{"timezone", "department"},
				Filter:             "users",
			},
			startIndex:         0,
			count:              50,
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users?attributes=userName%2Cemails.value&count=50&excludedAttributes=timezone%2Cdepartment&filter=users&startIndex=0",
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

			service := &SCIMUserService{client: mockClient}
			gotResult, gotResponse, err := service.Gets(testCase.context, testCase.directoryID,
				testCase.opts, testCase.startIndex, testCase.count)

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

				for _, user := range gotResult.Resources {

					t.Log("User ID", user.ID)

					for index, email := range user.Emails {
						t.Log("Email #", index, email.Value, email.Type, email.Primary)
					}

					t.Log("User Status: ", user.Active)
					t.Log("User Department: ", user.Department)
					t.Log("User Title: ", user.Title)

				}
			}

		})
	}

}

func TestSCIMUserService_Get(t *testing.T) {

	testCases := []struct {
		name                           string
		directoryID, userID            string
		attributes, excludedAttributes []string
		mockFile                       string
		wantHTTPMethod                 string
		endpoint                       string
		context                        context.Context
		wantHTTPCodeReturn             int
		wantErr                        bool
	}{
		{
			name:               "GetSCIMUserWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetSCIMUserWhenTheAttributesAndExcludedAttributesAreNotSet",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			attributes:         nil,
			excludedAttributes: nil,
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "GetSCIMUserWhenTheUserIDIsNotSet",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "",
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "GetSCIMUserWhenTheResponseBodyIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
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

			service := &SCIMUserService{client: mockClient}
			gotResult, gotResponse, err := service.Get(testCase.context, testCase.directoryID,
				testCase.userID, testCase.attributes, testCase.excludedAttributes)

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

				t.Log("User ID", gotResult.ID)

				for index, email := range gotResult.Emails {
					t.Log("Email #", index, email.Value, email.Type, email.Primary)
				}

				t.Log("User Status: ", gotResult.Active)
				t.Log("User Department: ", gotResult.Department)
				t.Log("User Title: ", gotResult.Title)
			}

		})
	}

}

func TestSCIMUserService_Deactivate(t *testing.T) {

	testCases := []struct {
		name                string
		directoryID, userID string
		mockFile            string
		wantHTTPMethod      string
		endpoint            string
		context             context.Context
		wantHTTPCodeReturn  int
		wantErr             bool
	}{
		{
			name:               "DeactivateSCIMUserWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeactivateSCIMUserWhenTheDirectoryIDIsNotSet",
			directoryID:        "",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeactivateSCIMUserWhenTheUserIDIsNotSet",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeactivateSCIMUserWhenTheRequestMethodIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeactivateSCIMUserWhenTheStatusCodeIsIncorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:               "DeactivateSCIMUserWhenTheContextIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeactivateSCIMUserWhenTheEndpointIsEmpty",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
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

			service := &SCIMUserService{client: mockClient}
			gotResponse, err := service.Deactivate(testCase.context, testCase.directoryID, testCase.userID)

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

func TestSCIMUserService_Update(t *testing.T) {

	testCases := []struct {
		name                           string
		directoryID                    string
		userID                         string
		payload                        *model.SCIMUserScheme
		attributes, excludedAttributes []string
		mockFile                       string
		wantHTTPMethod                 string
		endpoint                       string
		context                        context.Context
		wantHTTPCodeReturn             int
		wantErr                        bool
	}{
		{
			name:        "OverwriteSCIMUserWhenTheParametersAreCorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:        "OverwriteSCIMUserWhenTheUserIDIsNotSet",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheAttributesAndExcludedAttributesAreNotSet",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         nil,
			excludedAttributes: nil,
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "OverwriteSCIMUserWhenThePayloadIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload:            nil,
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheDirectoryIDIsNotSet",
			directoryID: "",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheRequestMethodIsIncorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheStatusCodeIsIncorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheContextIsNil",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheEndpointIsEmpty",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheResponseBodyIsEmpty",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload: &model.SCIMUserScheme{
				UserName: "Example Username 2",
				Emails: []*model.SCIMUserEmailScheme{
					{
						Value:   "example-2@go-atlassian.io",
						Type:    "work",
						Primary: true,
					},
				},
				Name: &model.SCIMUserNameScheme{
					Formatted:       "Example Full Name with Last Name",
					FamilyName:      "Example Family Name",
					GivenName:       "Example Name",
					MiddleName:      "Name",
					HonorificPrefix: "",
					HonorificSuffix: "",
				},

				DisplayName:       "Example Display Name",
				NickName:          "Example NickName",
				Title:             "Atlassian Administrator",
				PreferredLanguage: "en-US",
				Active:            true,
			},
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

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

			service := &SCIMUserService{client: mockClient}
			gotResult, gotResponse, err := service.Update(testCase.context, testCase.directoryID,
				testCase.userID, testCase.payload, testCase.attributes, testCase.excludedAttributes)

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

				t.Log("User ID", gotResult.ID)

				for index, email := range gotResult.Emails {
					t.Log("Email #", index, email.Value, email.Type, email.Primary)
				}

				t.Log("User Status: ", gotResult.Active)
				t.Log("User Department: ", gotResult.Department)
				t.Log("User Title: ", gotResult.Title)
			}

		})
	}

}

func TestSCIMUserService_Path(t *testing.T) {

	payload := &model.SCIMUserToPathScheme{
		Schemas: []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
	}

	if err := payload.AddStringOperation("replace", "displayName", "Docs Atlassian DisplayName 2"); err != nil {
		t.Fatal(err)
	}

	if err := payload.AddStringOperation("replace", "userName", "user-name-updated2"); err != nil {
		t.Fatal(err)
	}

	if err := payload.AddBoolOperation("replace", "active", false); err != nil {
		t.Fatal(err)
	}

	if err := payload.AddComplexOperation("add", "emails", []*model.SCIMUserComplexOperationScheme{
		{
			Value:     "primary@go-atlassian.io",
			ValueType: "work",
			Primary:   true,
		},
		{
			Value:     "second@go-atlassian.io",
			ValueType: "other",
			Primary:   false,
		},
	}); err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name                           string
		directoryID                    string
		userID                         string
		payload                        *model.SCIMUserToPathScheme
		attributes, excludedAttributes []string
		mockFile                       string
		wantHTTPMethod                 string
		endpoint                       string
		context                        context.Context
		wantHTTPCodeReturn             int
		wantErr                        bool
	}{
		{
			name:               "UpdateSCIMUserWhenTheParametersAreCorrect",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload:            payload,
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:        "OverwriteSCIMUserWhenTheUserIDIsNotSet",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "",
			payload:     payload,

			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:               "OverwriteSCIMUserWhenTheAttributesAndExcludedAttributesAreNotSet",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload:            payload,
			attributes:         nil,
			excludedAttributes: nil,
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            false,
		},

		{
			name:               "OverwriteSCIMUserWhenThePayloadIsNil",
			directoryID:        "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:             "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload:            nil,
			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheDirectoryIDIsNotSet",
			directoryID: "",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload:     payload,

			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheRequestMethodIsIncorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload:     payload,

			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheStatusCodeIsIncorrect",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload:     payload,

			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheContextIsNil",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload:     payload,

			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

			context:            nil,
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheEndpointIsEmpty",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload:     payload,

			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/scim-get-user.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusOK,
			wantErr:            true,
		},

		{
			name:        "OverwriteSCIMUserWhenTheResponseBodyIsEmpty",
			directoryID: "651c2e11-afea-4475-a0c4-422b89683e0f",
			userID:      "ef5ff80e-9ca6-449c-8cca-5b621085c6c9",
			payload:     payload,

			attributes:         []string{"userName", "emails.value"},
			excludedAttributes: []string{"timezone", "department"},
			mockFile:           "./mocks/empty.json",
			wantHTTPMethod:     http.MethodPatch,
			endpoint:           "/scim/directory/651c2e11-afea-4475-a0c4-422b89683e0f/Users/ef5ff80e-9ca6-449c-8cca-5b621085c6c9?attributes=userName%2Cemails.value&excludedAttributes=timezone%2Cdepartment",

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

			service := &SCIMUserService{client: mockClient}
			gotResult, gotResponse, err := service.Path(testCase.context, testCase.directoryID,
				testCase.userID, testCase.payload, testCase.attributes, testCase.excludedAttributes)

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

				t.Log("User ID", gotResult.ID)

				for index, email := range gotResult.Emails {
					t.Log("Email #", index, email.Value, email.Type, email.Primary)
				}

				t.Log("User Status: ", gotResult.Active)
				t.Log("User Department: ", gotResult.Department)
				t.Log("User Title: ", gotResult.Title)
			}

		})
	}

}

func TestSCIMUserToPathScheme_AddBoolOperation(t *testing.T) {
	type fields struct {
		Schemas    []string
		Operations []*model.SCIMUserToPathOperationScheme
	}
	type args struct {
		operation string
		path      string
		value     bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "AddBoolOperationWhenTheOperationIsNotSet",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "",
				path:      "active",
				value:     false,
			},
			wantErr: true,
		},

		{
			name: "AddBoolOperationWhenThePathIsNotSet",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "replace",
				path:      "",
				value:     false,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &model.SCIMUserToPathScheme{
				Schemas:    tt.fields.Schemas,
				Operations: tt.fields.Operations,
			}
			if err := s.AddBoolOperation(tt.args.operation, tt.args.path, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("AddBoolOperation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSCIMUserToPathScheme_AddComplexOperation(t *testing.T) {
	type fields struct {
		Schemas    []string
		Operations []*model.SCIMUserToPathOperationScheme
	}
	type args struct {
		operation string
		path      string
		values    []*model.SCIMUserComplexOperationScheme
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "AddComplexOperationWhenTheParametersAreCorrect",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "replace",
				path:      "emails",
				values: []*model.SCIMUserComplexOperationScheme{
					{
						Value:     "primary@go-atlassian.io",
						ValueType: "work",
						Primary:   true,
					},
					{
						Value:     "second@go-atlassian.io",
						ValueType: "other",
						Primary:   false,
					},
				},
			},
			wantErr: false,
		},

		{
			name: "AddComplexOperationWhenTheOperationIsNotSet",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "",
				path:      "emails",
				values: []*model.SCIMUserComplexOperationScheme{
					{
						Value:     "primary@go-atlassian.io",
						ValueType: "work",
						Primary:   true,
					},
					{
						Value:     "second@go-atlassian.io",
						ValueType: "other",
						Primary:   false,
					},
				},
			},
			wantErr: true,
		},

		{
			name: "AddComplexOperationWhenThePathIsNotSet",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "replace",
				path:      "",
				values: []*model.SCIMUserComplexOperationScheme{
					{
						Value:     "primary@go-atlassian.io",
						ValueType: "work",
						Primary:   true,
					},
					{
						Value:     "second@go-atlassian.io",
						ValueType: "other",
						Primary:   false,
					},
				},
			},
			wantErr: true,
		},

		{
			name: "AddComplexOperationWhenTheValuesAreNil",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "replace",
				path:      "emails",
				values:    nil,
			},
			wantErr: true,
		},

		{
			name: "AddComplexOperationWhenTheValuesDoNotContainsValues",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "replace",
				path:      "emails",
				values:    []*model.SCIMUserComplexOperationScheme{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &model.SCIMUserToPathScheme{
				Schemas:    tt.fields.Schemas,
				Operations: tt.fields.Operations,
			}
			if err := s.AddComplexOperation(tt.args.operation, tt.args.path, tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("AddComplexOperation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSCIMUserToPathScheme_AddStringOperation(t *testing.T) {
	type fields struct {
		Schemas    []string
		Operations []*model.SCIMUserToPathOperationScheme
	}
	type args struct {
		operation string
		path      string
		value     string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "AddStringOperationWhenTheParametersAreCorrect",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "displayName",
				path:      "active",
				value:     "DisplayName sample",
			},
			wantErr: false,
		},

		{
			name: "AddStringOperationWhenTheOperationIsNotSet",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "",
				path:      "active",
				value:     "DisplayName sample",
			},
			wantErr: true,
		},

		{
			name: "AddStringOperationWhenThePathIsNotSet",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "displayName",
				path:      "",
				value:     "DisplayName sample",
			},
			wantErr: true,
		},

		{
			name: "AddStringOperationWhenTheValueIsNotSet",
			fields: fields{
				Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:PatchOp"},
				Operations: nil,
			},
			args: args{
				operation: "displayName",
				path:      "active",
				value:     "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &model.SCIMUserToPathScheme{
				Schemas:    tt.fields.Schemas,
				Operations: tt.fields.Operations,
			}
			if err := s.AddStringOperation(tt.args.operation, tt.args.path, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("AddStringOperation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
