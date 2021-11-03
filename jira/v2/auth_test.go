package v2

import "testing"

func Test_AuthenticationService_SetBasicAuth_V2(t *testing.T) {

	mockedClient, err := startMockClient("")
	if err != nil {
		t.Log(err)
	}

	type fields struct {
		client      *Client
		token, mail string
		agent       string
	}
	type args struct {
		mail, token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "SetAccessTokenWhenTheParamsAreCorrect",
			fields: fields{
				client: mockedClient,
			},
			args: args{
				token: "$TOKEN_ID",
				mail:  "$MAIL_ID",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthenticationService{
				client: tt.fields.client,
				mail:   tt.fields.mail,
				token:  tt.fields.token,
				agent:  tt.fields.agent,
			}

			a.SetBasicAuth(tt.args.mail, tt.args.token)
		})
	}
}

func TestAuthenticationService_SetUserAgent_V2(t *testing.T) {

	mockedClient, err := startMockClient("")
	if err != nil {
		t.Log(err)
	}

	type fields struct {
		client *Client
		agent  string
	}
	type args struct {
		userAgent string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "SetUserAgentWhenTheParamsAreCorrect",
			fields: fields{
				client: mockedClient,
			},
			args: args{
				userAgent: "$USER_AGENT_SAMPLE",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthenticationService{
				client: tt.fields.client,
				agent:  tt.fields.agent,
			}

			a.SetUserAgent(tt.args.userAgent)
		})
	}

}
