package internal

import (
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/common"
	"github.com/ctreminiom/go-atlassian/service/mocks"
	"reflect"
	"testing"
)

func TestAuthenticationService_GetBasicAuth(t *testing.T) {

	type fields struct {
		c                 service.Client
		basicAuthProvided bool
		mail              string
		token             string
		userAgentProvided bool
		agent             string
	}
	testCases := []struct {
		name   string
		fields fields
		want   string
		want1  string
	}{
		{
			name: "when the basic auth is already set",
			fields: fields{
				c:     mocks.NewClient(t),
				mail:  "mail",
				token: "token",
			},
			want:  "mail",
			want1: "token",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			a := &AuthenticationService{
				c:                 testCase.fields.c,
				basicAuthProvided: testCase.fields.basicAuthProvided,
				mail:              testCase.fields.mail,
				token:             testCase.fields.token,
				userAgentProvided: testCase.fields.userAgentProvided,
				agent:             testCase.fields.agent,
			}

			got, got1 := a.GetBasicAuth()
			if got != testCase.want {
				t.Errorf("GetBasicAuth() got = %v, want %v", got, testCase.want)
			}

			if got1 != testCase.want1 {
				t.Errorf("GetBasicAuth() got1 = %v, want %v", got1, testCase.want1)
			}
		})
	}
}

func TestAuthenticationService_GetUserAgent(t *testing.T) {
	type fields struct {
		c                 service.Client
		basicAuthProvided bool
		mail              string
		token             string
		userAgentProvided bool
		agent             string
	}
	testCases := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "when the user agent is already set",
			fields: fields{
				c:     mocks.NewClient(t),
				agent: "firefox-09",
			},
			want: "firefox-09",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			a := &AuthenticationService{
				c:                 testCase.fields.c,
				basicAuthProvided: testCase.fields.basicAuthProvided,
				mail:              testCase.fields.mail,
				token:             testCase.fields.token,
				userAgentProvided: testCase.fields.userAgentProvided,
				agent:             testCase.fields.agent,
			}
			if got := a.GetUserAgent(); got != testCase.want {
				t.Errorf("GetUserAgent() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestAuthenticationService_HasBasicAuth(t *testing.T) {
	type fields struct {
		c                 service.Client
		basicAuthProvided bool
		mail              string
		token             string
		userAgentProvided bool
		agent             string
	}
	testCases := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "when the params are correct",
			fields: fields{
				c:                 mocks.NewClient(t),
				basicAuthProvided: true,
			},
			want: true,
		},
	}
	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			a := &AuthenticationService{
				c:                 testCase.fields.c,
				basicAuthProvided: testCase.fields.basicAuthProvided,
				mail:              testCase.fields.mail,
				token:             testCase.fields.token,
				userAgentProvided: testCase.fields.userAgentProvided,
				agent:             testCase.fields.agent,
			}
			if got := a.HasBasicAuth(); got != testCase.want {
				t.Errorf("HasBasicAuth() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestAuthenticationService_HasUserAgent(t *testing.T) {
	type fields struct {
		c                 service.Client
		basicAuthProvided bool
		mail              string
		token             string
		userAgentProvided bool
		agent             string
	}
	testCases := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "when the parameters are correct",
			fields: fields{
				userAgentProvided: true,
			},
			want: true,
		},
	}
	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {
			a := &AuthenticationService{
				c:                 testCase.fields.c,
				basicAuthProvided: testCase.fields.basicAuthProvided,
				mail:              testCase.fields.mail,
				token:             testCase.fields.token,
				userAgentProvided: testCase.fields.userAgentProvided,
				agent:             testCase.fields.agent,
			}

			if got := a.HasUserAgent(); got != testCase.want {
				t.Errorf("HasUserAgent() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestNewAuthenticationService(t *testing.T) {

	clientMocked := mocks.NewClient(t)

	type args struct {
		client service.Client
	}
	testCases := []struct {
		name string
		args args
		want common.Authentication
	}{
		{
			name: "when the parameters are correct",
			args: args{
				client: clientMocked,
			},
			want: NewAuthenticationService(clientMocked),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := NewAuthenticationService(testCase.args.client); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("NewAuthenticationService() = %v, want %v", got, testCase.want)
			}
		})
	}
}
