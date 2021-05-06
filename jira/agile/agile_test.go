package agile

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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

func startMockClient(instance string) (*Client, error) {

	mockClient, err := New(nil, instance)
	if err != nil {
		return nil, err
	}

	return mockClient, nil
}

func TestNew(t *testing.T) {

	mockClient, err := New(nil, "http://%3Fam:pa%3Fsword@google.com/")
	if err != nil {
		t.Fatal(err)
	}

	mockClient.Auth.SetBasicAuth("test", "test")
	mockClient.Auth.SetUserAgent("aaa")

	mockClient2, _ := New(nil, " https://zhidao.baidu.com/special/view?id=49105a24626975510000&preview=1")

	type args struct {
		httpClient *http.Client
		site       string
	}
	tests := []struct {
		name       string
		args       args
		wantClient *Client
		wantErr    bool
	}{
		{
			name: "NewWhenTHeParametersAreCorrect",
			args: args{
				httpClient: nil,
				site:       "http://%3Fam:pa%3Fsword@google.com",
			},
			wantClient: mockClient,
			wantErr:    false,
		},

		{
			name: "NewWhenTheURLIsNotCorrect",
			args: args{
				httpClient: nil,
				site:       " https://zhidao.baidu.com/special/view?id=49105a24626975510000&preview=1",
			},
			wantClient: mockClient2,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotClient, err := New(tt.args.httpClient, tt.args.site)

			if tt.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotClient, nil)
			}

		})
	}
}
