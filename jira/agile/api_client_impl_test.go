package agile

import (
	"bytes"
	"errors"
	"github.com/ctreminiom/go-atlassian/jira/agile/internal/mocks"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service/agile"
	"github.com/ctreminiom/go-atlassian/service/common"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
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

	type fields struct {
		HTTP           HttpClient
		Site           *url.URL
		Authentication common.Authentication
		Board          agile.Board
		Epic           agile.Epic
		Sprint         agile.Sprint
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
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			if testCase.on != nil {
				testCase.on(&testCase.fields)
			}

			c := &Client{
				HTTP:           testCase.fields.HTTP,
				Site:           testCase.fields.Site,
				Authentication: testCase.fields.Authentication,
				Board:          testCase.fields.Board,
				Epic:           testCase.fields.Epic,
				Sprint:         testCase.fields.Sprint,
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

			/*

				if (err != nil) != testCase.wantErr {
					assert.EqualError(t, err, testCase.Err.Error())
					t.Errorf("Call() error = %v, wantErr %v", err, testCase.wantErr)
					return
				}
			*/

		})
	}
}
