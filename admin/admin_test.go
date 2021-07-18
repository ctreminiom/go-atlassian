package admin

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

type mockServerOptions struct {
	Endpoint           string
	MockFilePath       string
	MethodAccepted     string
	Headers            map[string]string
	ResponseCodeWanted int
}

func startMockServer(opts *mockServerOptions) (*httptest.Server, error) {

	mockServer := httptest.NewServer(

		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != opts.MethodAccepted {
				http.Error(w, fmt.Sprintf("Request method: %v, want %v", r.Method, opts.MethodAccepted), http.StatusMethodNotAllowed)
				return
			}

			if r.URL.Query().Encode() != "" {

				var pathWithQueries = fmt.Sprintf("%v?%v", r.URL.Path, r.URL.Query().Encode())

				if pathWithQueries != opts.Endpoint {
					http.Error(w, fmt.Sprintf("Request URL: %v, want %v", r.URL.Path, opts.Endpoint), 400)
					return
				}

			} else {
				if r.URL.Path != opts.Endpoint {
					http.Error(w, fmt.Sprintf("Request URL: %v, want %v", r.URL.Path, opts.Endpoint), 400)
					return
				}
			}

			//Append the custom headers
			for key, value := range opts.Headers {
				w.Header().Add(key, value)
			}

			//Append the Method
			w.WriteHeader(opts.ResponseCodeWanted)

			//Append the JSON Mock file if it's provided
			if len(opts.MockFilePath) != 0 {
				mockResponse, err := ioutil.ReadFile(opts.MockFilePath)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
				_, err = w.Write(mockResponse)
				if err != nil {
					http.Error(w, err.Error(), 500)
					return
				}
			}

		}),
	)

	return mockServer, nil
}

func startMockClient(site string) (*Client, error) {

	siteAsURL, err := url.Parse(site)
	if err != nil {
		return nil, err
	}

	client := &Client{}
	client.HTTP = http.DefaultClient
	client.Site = siteAsURL

	client.Auth = &AuthenticationService{client: client}
	client.Organization = &OrganizationService{
		client: client,
		Policy: &OrganizationPolicyService{
			client: client,
		},
	}

	client.User = &UserService{
		client: client,
		Token:  &UserTokenService{client: client},
	}

	client.SCIM = &SCIMService{
		client: client,
		User:   &SCIMUserService{client: client},
		Group:  &SCIMGroupService{client: client},
		Scheme: &SCIMSchemeService{client: client},
	}

	return client, nil
}

func TestNew(t *testing.T) {

	mockedClient, err := startMockClient(ApiEndpoint)
	if err != nil {
		t.Log(err)
	}

	type args struct {
		httpClient *http.Client
	}
	tests := []struct {
		name       string
		args       args
		wantClient *Client
		wantErr    bool
	}{
		{
			name: "NewClientWhenTheParametersAreCorrect",
			args: args{
				httpClient: http.DefaultClient,
			},
			wantClient: mockedClient,
			wantErr:    false,
		},

		{
			name: "NewClientWhenTheHTTPClientIsNil",
			args: args{
				httpClient: nil,
			},
			wantClient: mockedClient,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClient, err := New(tt.args.httpClient)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotClient, tt.wantClient) {
				t.Errorf("New() gotClient = %v, want %v", gotClient, tt.wantClient)
			}
		})
	}
}

func TestClient_newRequest(t *testing.T) {

	mockClient, err := New(nil)
	if err != nil {
		t.Fatal(err)
	}

	mockClient.Auth.SetBearerToken("user")
	mockClient.Auth.SetUserAgent("bot-1.0.0")

	type args struct {
		ctx         context.Context
		method      string
		apiEndpoint string
		payload     io.Reader
	}
	tests := []struct {
		name    string
		client  *Client
		args    args
		wantErr bool
	}{
		{
			name:   "CreateNewRequestWhenTheParametersAreCorrect",
			client: mockClient,
			args: args{
				ctx:         context.Background(),
				method:      http.MethodGet,
				apiEndpoint: "/example",
				payload:     nil,
			},
			wantErr: false,
		},

		{
			name:   "CreateNewRequestWhenTheContextIsNotProvided",
			client: mockClient,
			args: args{
				ctx:         nil,
				method:      http.MethodGet,
				apiEndpoint: "/example",
				payload:     nil,
			},
			wantErr: true,
		},

		{
			name:   "CreateNewRequestWhenTheEndpointIsNotAvailableToFormat",
			client: mockClient,
			args: args{
				ctx:         context.Background(),
				method:      http.MethodGet,
				apiEndpoint: " https://zhidao.baidu.com/special/view?id=49105a24626975510000&preview=1",
				payload:     nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotRequest, err := tt.client.newRequest(tt.args.ctx, tt.args.method, tt.args.apiEndpoint, tt.args.payload)
			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotRequest, nil)
			}
		})
	}
}

func Test_transformStructToReader(t *testing.T) {
	type args struct {
		structure interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantReader io.Reader
		wantErr    bool
	}{

		{
			name: "TransformStructToReaderWhenTheStructIsNotSerialize",
			args: args{
				structure: make(chan int),
			},
			wantReader: nil,
			wantErr:    true,
		},

		{
			name: "TransformStructToReaderWhenTheParametersAreCorrect",
			args: args{
				structure: &SCIMUserScheme{
					UserName: "Example Username 3",
					Emails: []*SCIMUserEmailScheme{
						{
							Value:   "example-2@go-atlassian.io",
							Type:    "work",
							Primary: true,
						},
					},
					Name: &SCIMUserNameScheme{
						Formatted:       "Example Full Name with Last Name",
						FamilyName:      "Example Family Name",
						GivenName:       "Example Name",
						MiddleName:      "Name",
						HonorificPrefix: "",
						HonorificSuffix: "",
					},

					DisplayName:       "Example Display Name 3",
					NickName:          "Example NickName",
					Title:             "Atlassian Administrator",
					PreferredLanguage: "en-US",
					Active:            true,
				},
			},
			wantReader: nil,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReader, err := transformStructToReader(tt.args.structure)
			if (err != nil) != tt.wantErr {
				t.Errorf("transformStructToReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.wantReader = gotReader
			if !reflect.DeepEqual(gotReader, tt.wantReader) {
				t.Errorf("transformStructToReader() gotReader = %v, want %v", gotReader, tt.wantReader)
			}
		})
	}
}

func Test_transformTheHTTPResponse(t *testing.T) {

	var (
		responseScenarios      = make(map[string]*http.Response)
		responseConfigurations = make(map[string]map[string]interface{})
	)

	//Add the scenarios
	responseConfigurations["badRequestResponseWithIncorrectFormat"] = map[string]interface{}{
		"endpoint":      "/example",
		"mock-response": "./mocks/get-contents.json",
		"method":        http.MethodGet,
		"status":        http.StatusBadRequest,
		"closed?": false,
	}

	responseConfigurations["badRequestResponseWithNotResponseBody"] = map[string]interface{}{
		"endpoint":      "/example",
		"mock-response": "",
		"method":        http.MethodGet,
		"status":        http.StatusBadRequest,
		"closed?": true,
	}

	responseConfigurations["OkRequestResponseWithNotResponseBody"] = map[string]interface{}{
		"endpoint":      "/",
		"mock-response": "./mocks/get-contents.json",
		"method":        http.MethodGet,
		"status":        http.StatusOK,
		"closed?": true,
	}

	for scenario, configuration := range responseConfigurations {

		mockOptions := mockServerOptions{
			Endpoint:           configuration["endpoint"].(string),
			MockFilePath:       configuration["mock-response"].(string),
			MethodAccepted:     configuration["method"].(string),
			ResponseCodeWanted: configuration["status"].(int),
		}

		mockServer, err := startMockServer(&mockOptions)
		if err != nil {
			t.Fatal(err)
		}

		mockRequest, err := http.NewRequest(http.MethodGet, mockServer.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		mockResponse, err := http.DefaultClient.Do(mockRequest)
		if err != nil {
			t.Fatal(err)
		}

		if configuration["closed?"].(bool) {
			mockResponse.Body.Close()
		}

		responseScenarios[scenario] = mockResponse
	}

	type args struct {
		response  *http.Response
		structure interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult *ResponseScheme
		wantErr    bool
	}{
		{
			name: "TransformHTTPResponseWhenResponseIsNil",
			args: args{
				response:  nil,
				structure: nil,
			},
			wantResult: nil,
			wantErr:    true,
		},

		{
			name: "TransformHTTPResponseWhenResponseReturnsABadRequestAndIncorrectFormat",
			args: args{
				response:  responseScenarios["badRequestResponseWithIncorrectFormat"],
				structure: nil,
			},
			wantResult: nil,
			wantErr:    true,
		},

		{
			name: "TransformHTTPResponseWhenResponseReturnsABadRequestAndWithNotResponseBody",
			args: args{
				response:  responseScenarios["badRequestResponseWithNotResponseBody"],
				structure: nil,
			},
			wantResult: nil,
			wantErr:    true,
		},

		{
			name: "TransformHTTPResponseWhenResponseReturnsAOkRequestAndWithNotResponseBody",
			args: args{
				response:  responseScenarios["OkRequestResponseWithNotResponseBody"],
				structure: nil,
			},
			wantResult: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := transformTheHTTPResponse(tt.args.response, tt.args.structure)
			if (err != nil) != tt.wantErr {
				t.Errorf("transformTheHTTPResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.wantResult = gotResult

			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("transformTheHTTPResponse() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
