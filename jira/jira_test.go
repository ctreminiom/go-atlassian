package jira

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type mockServerOptions struct {
	Endpoint           string
	MockFilePath       string
	MethodAccepted     string
	Headers            map[string]string
	ResponseCodeWanted int
}

func startMockServer(opts *mockServerOptions) (*httptest.Server, error) {

	//Load the mock json file
	mockResponse, err := ioutil.ReadFile(opts.MockFilePath)
	if err != nil {
		return nil, err
	}

	mockServer := httptest.NewServer(

		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != opts.MethodAccepted {
				http.Error(w, fmt.Sprintf("Request method: %v, want %v", r.Method, opts.MethodAccepted), 415)
				return
			}

			if r.URL.Path != opts.Endpoint {
				http.Error(w, fmt.Sprintf("Request URL: %v, want %v", r.URL.Path, opts.Endpoint), 400)
				return
			}

			//Append the custom headers
			for key, value := range opts.Headers {
				w.Header().Add(key, value)
			}

			//Append the Method
			w.WriteHeader(opts.ResponseCodeWanted)
			_, _ = w.Write(mockResponse)

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
