package internal

import (
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
	"github.com/ctreminiom/go-atlassian/v2/service/mocks"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestAuthenticationService_GetBasicAuth(t *testing.T) {

	type fields struct {
		c                 service.Connector
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
				c:     mocks.NewConnector(t),
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
		c                 service.Connector
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
				c:     mocks.NewConnector(t),
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
		c                 service.Connector
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
				c:                 mocks.NewConnector(t),
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
		c                 service.Connector
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

	clientMocked := mocks.NewConnector(t)

	type args struct {
		client service.Connector
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

func TestAuthenticationService_SetUserAgent(t *testing.T) {

	type fields struct {
		c                 service.Connector
		basicAuthProvided bool
		mail              string
		token             string
		userAgentProvided bool
		agent             string
	}
	type args struct {
		agent string
	}
	testCases := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "when the parameters are correct",
			args: args{agent: "mozilla-9"},
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

			a.SetUserAgent(testCase.args.agent)
		})
	}
}

func TestAuthenticationService_SetBasicAuth(t *testing.T) {

	type fields struct {
		c                 service.Connector
		basicAuthProvided bool
		mail              string
		token             string
		userAgentProvided bool
		agent             string
	}
	type args struct {
		mail  string
		token string
	}
	testCases := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "when the parameters are correct",
			args: args{
				mail:  "example@example.com",
				token: "token-sample",
			},
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

			a.SetBasicAuth(testCase.args.mail, testCase.args.token)
		})
	}
}

func TestAuthenticationService_HasSetExperimentalFlag(t *testing.T) {

	type fields struct {
		c                 service.Connector
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
			name:   "when the parameters are correct",
			fields: fields{},
			want:   false,
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
			if got := a.HasSetExperimentalFlag(); got != testCase.want {
				t.Errorf("HasSetExperimentalFlag() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestAuthenticationService_GetBearerToken(t *testing.T) {
	type fields struct {
		c                 service.Connector
		basicAuthProvided bool
		mail              string
		token             string
		userAgentProvided bool
		agent             string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "when the parameters are correct",
			fields: fields{
				c:     mocks.NewConnector(t),
				token: "token-sample",
			},
			want: "token-sample",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthenticationService{
				c:                 tt.fields.c,
				basicAuthProvided: tt.fields.basicAuthProvided,
				mail:              tt.fields.mail,
				token:             tt.fields.token,
				userAgentProvided: tt.fields.userAgentProvided,
				agent:             tt.fields.agent,
			}
			assert.Equalf(t, tt.want, a.GetBearerToken(), "GetBearerToken()")
		})
	}
}

func TestAuthenticationService_SetBearerToken(t *testing.T) {
	type fields struct {
		c                 service.Connector
		basicAuthProvided bool
		mail              string
		token             string
		userAgentProvided bool
		agent             string
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "when the parameters are correct",
			fields: fields{
				token: "token-sample",
			},
			args: args{
				token: "token-sample",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthenticationService{
				c:                 tt.fields.c,
				basicAuthProvided: tt.fields.basicAuthProvided,
				mail:              tt.fields.mail,
				token:             tt.fields.token,
				userAgentProvided: tt.fields.userAgentProvided,
				agent:             tt.fields.agent,
			}
			a.SetBearerToken(tt.args.token)
		})
	}
}

func TestAuthenticationService_SetExperimentalFlag(t *testing.T) {
	type fields struct {
		c                 service.Connector
		basicAuthProvided bool
		mail              string
		token             string
		userAgentProvided bool
		agent             string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "when the parameters are correct",
			fields: fields{
				c:                 nil,
				basicAuthProvided: false,
				mail:              "",
				token:             "",
				userAgentProvided: false,
				agent:             "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthenticationService{
				c:                 tt.fields.c,
				basicAuthProvided: tt.fields.basicAuthProvided,
				mail:              tt.fields.mail,
				token:             tt.fields.token,
				userAgentProvided: tt.fields.userAgentProvided,
				agent:             tt.fields.agent,
			}
			a.SetExperimentalFlag()
		})
	}
}
