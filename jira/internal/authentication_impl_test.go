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
		oauth2ConfigProvided bool
		clientID, clientSecret, redirectURI string
		oauth2AccessTokenProvided bool
		oauth2AccessToken string
		oauth2RefreshTokenProvided bool
		oauth2RefreshToken string
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

func TestAuthenticationService_OAuth2Config(t *testing.T) {
	testCases := []struct {
		name                string
		clientID            string
		clientSecret        string
		redirectURI         string
		wantClientID        string
		wantClientSecret    string
		wantRedirectURI     string
		wantHasOAuth2Config bool
	}{
		{
			name:                "sets and gets OAuth2 config correctly",
			clientID:            "test-client-id",
			clientSecret:        "test-client-secret",
			redirectURI:         "https://example.com/callback",
			wantClientID:        "test-client-id",
			wantClientSecret:    "test-client-secret",
			wantRedirectURI:     "https://example.com/callback",
			wantHasOAuth2Config: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := &AuthenticationService{}
			
			// Test SetOAuth2Config
			a.SetOAuth2Config(tc.clientID, tc.clientSecret, tc.redirectURI)
			
			// Test GetOAuth2Config
			gotClientID, gotClientSecret, gotRedirectURI := a.GetOAuth2Config()
			assert.Equal(t, tc.wantClientID, gotClientID)
			assert.Equal(t, tc.wantClientSecret, gotClientSecret)
			assert.Equal(t, tc.wantRedirectURI, gotRedirectURI)
			
			// Test HasOAuth2Config
			assert.Equal(t, tc.wantHasOAuth2Config, a.HasOAuth2Config())
		})
	}
}

func TestAuthenticationService_OAuth2AccessToken(t *testing.T) {
	testCases := []struct {
		name                      string
		accessToken               string
		wantAccessToken           string
		wantHasOAuth2AccessToken  bool
	}{
		{
			name:                      "sets and gets OAuth2 access token correctly",
			accessToken:               "test-access-token",
			wantAccessToken:           "test-access-token",
			wantHasOAuth2AccessToken:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := &AuthenticationService{}
			
			// Test SetOAuth2AccessToken
			a.SetOAuth2AccessToken(tc.accessToken)
			
			// Test GetOAuth2AccessToken
			gotAccessToken := a.GetOAuth2AccessToken()
			assert.Equal(t, tc.wantAccessToken, gotAccessToken)
			
			// Test HasOAuth2AccessToken
			assert.Equal(t, tc.wantHasOAuth2AccessToken, a.HasOAuth2AccessToken())
		})
	}
}

func TestAuthenticationService_OAuth2RefreshToken(t *testing.T) {
	testCases := []struct {
		name                       string
		refreshToken               string
		wantRefreshToken           string
		wantHasOAuth2RefreshToken  bool
	}{
		{
			name:                       "sets and gets OAuth2 refresh token correctly",
			refreshToken:               "test-refresh-token",
			wantRefreshToken:           "test-refresh-token",
			wantHasOAuth2RefreshToken:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := &AuthenticationService{}
			
			// Test SetOAuth2RefreshToken
			a.SetOAuth2RefreshToken(tc.refreshToken)
			
			// Test GetOAuth2RefreshToken
			gotRefreshToken := a.GetOAuth2RefreshToken()
			assert.Equal(t, tc.wantRefreshToken, gotRefreshToken)
			
			// Test HasOAuth2RefreshToken
			assert.Equal(t, tc.wantHasOAuth2RefreshToken, a.HasOAuth2RefreshToken())
		})
	}
}
