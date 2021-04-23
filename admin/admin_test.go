package admin

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestClient_Do(t *testing.T) {

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
			name:               "DoHTTPRequestWhenTheParametersAreCorrect",
			mockFile:           "./mocks/get-organization-policy.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/9a1jj823-jac8-123d-jj01-63315k059cb2/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            false,
		},

		{
			name:               "DoHTTPRequestWhenTheContextIsNil",
			mockFile:           "./mocks/get-organization-policy.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/admin/v1/orgs/9a1jj823-jac8-123d-jj01-63315k059cb2/policies/60f0f660-be3e-4d70-bd34-9c2858ec040f",
			context:            nil,
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},

		{
			name:               "DoHTTPRequestWhenTheEndpointIsIncorrect",
			mockFile:           "./mocks/get-organization-policy.json",
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "postgres://user:abc{DEf1=ghi@example.com:5432/db?sslmode=require",
			context:            context.Background(),
			wantHTTPCodeReturn: 200,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

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

			//Init the mocked HTTP request
			requestMocked, err := mockClient.newRequest(testCase.context, testCase.wantHTTPMethod, testCase.endpoint, nil)
			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
				return
			}

			gotResponse, err := mockClient.Do(requestMocked)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}
				assert.Error(t, err)
				return
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
			}

		})

	}

}

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
