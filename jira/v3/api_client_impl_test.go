package v3

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service/common"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestClient_Call(t *testing.T) {

	expectedResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	nonExpectedResponse := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader("Hello, world!")),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	type fields struct {
		HTTP            common.HttpClient
		Site            *url.URL
		Authentication  common.Authentication
		ApplicationRole jira.AppRoleConnector
	}

	type args struct {
		request   *http.Request
		structure interface{}
	}

	testCases := []struct {
		name    string
		fields  fields
		on      func(*fields)
		args    args
		want    *models.ResponseScheme
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			on: func(fields *fields) {

				client := mocks.NewHttpClient(t)

				client.On("Do", (*http.Request)(nil)).
					Return(expectedResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   nil,
				structure: nil,
			},
			want: &models.ResponseScheme{
				Response: expectedResponse,
				Code:     http.StatusOK,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: false,
		},

		{
			name: "when the http callback cannot be executed",
			on: func(fields *fields) {

				client := mocks.NewHttpClient(t)

				client.On("Do", (*http.Request)(nil)).
					Return(nil, errors.New("error, unable to execute the http call"))

				fields.HTTP = client
			},
			wantErr: true,
			Err:     errors.New("error, unable to execute the http call"),
		},

		{
			name: "when the response status is not valid",
			on: func(fields *fields) {

				client := mocks.NewHttpClient(t)

				client.On("Do", (*http.Request)(nil)).
					Return(nonExpectedResponse, nil)

				fields.HTTP = client
			},
			args: args{
				request:   nil,
				structure: nil,
			},
			want: &models.ResponseScheme{
				Response: nonExpectedResponse,
				Code:     http.StatusBadRequest,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString("Hello, world!"),
			},
			wantErr: true,
			Err:     models.ErrInvalidStatusCodeError,
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
				Auth: testCase.fields.Authentication,
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

func TestNew(t *testing.T) {

	mockClient, err := New(http.DefaultClient, "https://ctreminiom.atlassian.net")
	if err != nil {
		t.Fatal(err)
	}

	mockClient.Auth.SetBasicAuth("test", "test")
	mockClient.Auth.SetUserAgent("aaa")

	mockClient2, _ := New(nil, " https://zhidao.baidu.com/special/view?id=sd&preview=1")

	type args struct {
		httpClient common.HttpClient
		site       string
	}

	testCases := []struct {
		name    string
		args    args
		on      func(*args)
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
			name: "when the site url is not valid",
			args: args{
				httpClient: http.DefaultClient,
				site:       " https://zhidao.baidu.com/special/view?id=sd&preview=1",
			},
			want:    mockClient2,
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

func TestClient_TransformTheHTTPResponse(t *testing.T) {

	expectedJsonResponse := `
	{
	  "id": 4,
	  "self": "https://ctreminiom.atlassian.net/rest/agile/1.0/board/4",
	  "name": "KP - Scrum",
	  "type": "scrum"
	}`

	expectedResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(expectedJsonResponse)),
		Request: &http.Request{
			Method: http.MethodGet,
			URL:    &url.URL{},
		},
	}

	type fields struct {
		HTTP            common.HttpClient
		Site            *url.URL
		Authentication  common.Authentication
		ApplicationRole jira.AppRoleConnector
	}

	type args struct {
		response  *http.Response
		structure interface{}
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		want    *models.ResponseScheme
		wantErr bool
		Err     error
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				response:  expectedResponse,
				structure: models.BoardScheme{},
			},
			want: &models.ResponseScheme{
				Response: expectedResponse,
				Code:     http.StatusOK,
				Method:   http.MethodGet,
				Bytes:    *bytes.NewBufferString(expectedJsonResponse),
			},
			wantErr: false,
		},

		{
			name:   "when the payload is not provided",
			fields: fields{},
			args: args{
				response:  nil,
				structure: models.BoardScheme{},
			},
			wantErr: true,
			Err:     errors.New("validation failed, please provide a http.Response pointer"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := &Client{
				HTTP: testCase.fields.HTTP,
				Site: testCase.fields.Site,
				Auth: testCase.fields.Authentication,
			}

			got, err := c.TransformTheHTTPResponse(testCase.args.response, testCase.args.structure)

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

func TestClient_TransformStructToReader(t *testing.T) {

	expectedBytes, err := json.Marshal(&models.BoardScheme{
		Name: "Board Sample",
		Type: "Scrum",
	})

	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		HTTP            common.HttpClient
		Site            *url.URL
		Authentication  common.Authentication
		ApplicationRole jira.AppRoleConnector
	}

	type args struct {
		structure interface{}
	}

	testCases := []struct {
		name    string
		fields  fields
		args    args
		want    io.Reader
		wantErr bool
		Err     error
	}{
		{
			name: "when the parameters are correct",
			args: args{
				structure: &models.BoardScheme{
					Name: "Board Sample",
					Type: "Scrum",
				},
			},
			want:    bytes.NewReader(expectedBytes),
			wantErr: false,
		},

		{
			name: "when the payload provided is not a pointer",
			args: args{
				structure: models.BoardScheme{
					Name: "Board Sample",
					Type: "Scrum",
				},
			},
			want:    bytes.NewReader(expectedBytes),
			wantErr: true,
			Err:     models.ErrNonPayloadPointerError,
		},

		{
			name: "when the payload is not provided",
			args: args{
				structure: nil,
			},
			want:    bytes.NewReader(expectedBytes),
			wantErr: true,
			Err:     models.ErrNilPayloadError,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			c := &Client{
				HTTP: testCase.fields.HTTP,
				Site: testCase.fields.Site,
				Auth: testCase.fields.Authentication,
			}

			got, err := c.TransformStructToReader(testCase.args.structure)

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
