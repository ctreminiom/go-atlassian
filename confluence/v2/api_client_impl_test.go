package v2

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ctreminiom/go-atlassian/v2/confluence/internal"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
)

func TestClient_Call(t *testing.T) {

	expectedResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	badRequestResponse := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	internalServerResponse := &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body:       io.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	unauthorizedResponse := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       io.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	notFoundResponse := &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       io.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	type fields struct {
		HTTP common.HTTPClient
		Site *url.URL
		Auth common.Authentication
	}
	type args struct {
		request   *http.Request
		structure interface{}
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		on      func(*fields)
		want    *model.ResponseScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			on: func(fields *fields) {

				client := mocks.NewHTTPClient(t)

				client.On("Do", (*http.Request)(nil)).
					Return(expectedResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   nil,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: expectedResponse,
				Code:     http.StatusOK,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: false,
		},

		{
			name: "when the response status is a bad request",
			on: func(fields *fields) {

				client := mocks.NewHTTPClient(t)

				client.On("Do", (*http.Request)(nil)).
					Return(badRequestResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   nil,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: badRequestResponse,
				Code:     http.StatusBadRequest,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: true,
			Err:     model.ErrBadRequest,
		},

		{
			name: "when the response status is an internal service error",
			on: func(fields *fields) {

				client := mocks.NewHTTPClient(t)

				client.On("Do", (*http.Request)(nil)).
					Return(internalServerResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   nil,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: internalServerResponse,
				Code:     http.StatusInternalServerError,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: true,
			Err:     model.ErrInternal,
		},

		{
			name: "when the response status is a not found",
			on: func(fields *fields) {

				client := mocks.NewHTTPClient(t)

				client.On("Do", (*http.Request)(nil)).
					Return(notFoundResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   nil,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: notFoundResponse,
				Code:     http.StatusNotFound,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: true,
			Err:     model.ErrNotFound,
		},

		{
			name: "when the response status is unauthorized",
			on: func(fields *fields) {

				client := mocks.NewHTTPClient(t)

				client.On("Do", (*http.Request)(nil)).
					Return(unauthorizedResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   nil,
				structure: nil,
			},
			want: &model.ResponseScheme{
				Response: unauthorizedResponse,
				Code:     http.StatusUnauthorized,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: true,
			Err:     model.ErrUnauthorized,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			c := &Client{
				HTTP: testCase.fields.HTTP,
				Site: testCase.fields.Site,
			}

			got, err := c.Call(testCase.args.request, testCase.args.structure)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.EqualError(t, err, testCase.Err.Error())

			} else {
				assert.NoError(t, err)
				assert.Equal(t, got, testCase.want)
			}

		})
	}
}

func TestClient_NewRequest(t *testing.T) {

	authMocked := internal.NewAuthenticationService(nil)
	authMocked.SetBasicAuth("mail", "token")
	authMocked.SetUserAgent("firefox")
	authMocked.SetBearerToken("token_sample")

	siteAsURL, err := url.Parse("https://ctreminiom.atlassian.net")
	if err != nil {
		t.Fatal(err)
	}

	requestMocked, err := http.NewRequestWithContext(context.TODO(),
		http.MethodGet,
		"https://ctreminiom.atlassian.net/rest/2/issue/attachment",
		bytes.NewReader([]byte("Hello World")),
	)

	if err != nil {
		t.Fatal(err)
	}

	requestMocked.Header.Set("Accept", "application/json")
	requestMocked.Header.Set("Content-Type", "application/json")

	type fields struct {
		HTTP common.HTTPClient
		Auth common.Authentication
		Site *url.URL
	}

	type args struct {
		ctx         context.Context
		method      string
		urlStr      string
		contentType string
		body        interface{}
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "when the parameters are correct",
			fields: fields{
				HTTP: http.DefaultClient,
				Auth: authMocked,
				Site: siteAsURL,
			},
			args: args{
				ctx:         context.Background(),
				method:      http.MethodGet,
				urlStr:      "rest/2/issue/attachment",
				contentType: "",
				body:        bytes.NewReader([]byte("Hello World")),
			},
			want:    requestMocked,
			wantErr: false,
		},

		{
			name: "when the url cannot be parsed",
			fields: fields{
				HTTP: http.DefaultClient,
				Auth: internal.NewAuthenticationService(nil),
				Site: siteAsURL,
			},
			args: args{
				ctx:    context.Background(),
				method: http.MethodGet,
				urlStr: " https://zhidao.baidu.com/special/view?id=49105a24626975510000&preview=1",
				body:   bytes.NewReader([]byte("Hello World")),
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "when the content type is provided",
			fields: fields{
				HTTP: http.DefaultClient,
				Auth: authMocked,
				Site: siteAsURL,
			},
			args: args{
				ctx:         context.Background(),
				method:      http.MethodGet,
				urlStr:      "rest/2/issue/attachment",
				contentType: "type_sample",
				body:        bytes.NewReader([]byte("Hello World")),
			},
			want:    requestMocked,
			wantErr: false,
		},

		{
			name: "when the request cannot be created",
			fields: fields{
				HTTP: http.DefaultClient,
				Auth: internal.NewAuthenticationService(nil),
				Site: siteAsURL,
			},
			args: args{
				ctx:    nil,
				method: http.MethodGet,
				urlStr: "rest/2/issue/attachment",
				body:   bytes.NewReader([]byte("Hello World")),
			},
			want:    requestMocked,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			c := &Client{
				HTTP: testCase.fields.HTTP,
				Auth: testCase.fields.Auth,
				Site: testCase.fields.Site,
			}

			got, err := c.NewRequest(
				testCase.args.ctx,
				testCase.args.method,
				testCase.args.urlStr,
				testCase.args.contentType,
				testCase.args.body,
			)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, got, nil)
			}

		})
	}
}

func TestClient_processResponse(t *testing.T) {

	expectedJSONResponse := `
	{
	  "id": 4,
	  "self": "https://ctreminiom.atlassian.net/rest/agile/1.0/board/4",
	  "name": "KP - Scrum",
	  "type": "scrum"
	}`

	expectedResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(expectedJSONResponse)),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	type fields struct {
		HTTP           common.HTTPClient
		Site           *url.URL
		Authentication common.Authentication
	}
	type args struct {
		response  *http.Response
		structure interface{}
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ResponseScheme
		wantErr bool
		Err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				response:  expectedResponse,
				structure: model.BoardScheme{},
			},
			want: &model.ResponseScheme{
				Response: expectedResponse,
				Code:     http.StatusOK,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString(expectedJSONResponse),
			},
			wantErr: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			c := &Client{
				HTTP: testCase.fields.HTTP,
				Site: testCase.fields.Site,
				Auth: testCase.fields.Authentication,
			}

			got, err := c.processResponse(testCase.args.response, testCase.args.structure)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
				assert.EqualError(t, err, testCase.Err.Error())

			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, got, nil)
			}

		})
	}
}

func TestNew(t *testing.T) {

	mockClient, err := New(http.DefaultClient, "https://ctreminiom.atlassian.net")
	if err != nil {
		t.Fatal(err)
	}

	mockClient.Auth.SetBasicAuth("test", "test")
	mockClient.Auth.SetUserAgent("aaa")

	invalidURLClientMocked, _ := New(nil, " https://zhidao.baidu.com/special/view?id=sd&preview=1")

	noURLClientMocked, _ := New(nil, "")

	type args struct {
		httpClient common.HTTPClient
		site       string
	}

	testCases := []struct {
		name    string
		args    args
		want    *Client
		wantErr bool
		Err     error
	}{

		{
			name: "when the parameters are correct",
			args: args{
				httpClient: http.DefaultClient,
				site:       "https://ctreminiom.atlassian.net",
			},
			want:    mockClient,
			wantErr: false,
		},

		{
			name: "when the site url are not provided",
			args: args{
				httpClient: http.DefaultClient,
				site:       "",
			},
			want:    noURLClientMocked,
			wantErr: true,
			Err:     model.ErrNoSite,
		},
		{
			name: "when the site url is not valid",
			args: args{
				httpClient: http.DefaultClient,
				site:       " https://zhidao.baidu.com/special/view?id=sd&preview=1",
			},
			want:    invalidURLClientMocked,
			wantErr: true,
			Err:     errors.New("parse \" https://zhidao.baidu.com/special/view?id=sd&preview=1/\": first path segment in URL cannot contain colon"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			gotClient, err := New(testCase.args.httpClient, testCase.args.site)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
				assert.EqualError(t, err, testCase.Err.Error())

			} else {
				assert.NoError(t, err)
				assert.NotEqual(t, gotClient, nil)
			}
		})
	}
}
