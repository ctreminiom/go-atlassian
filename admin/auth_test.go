package admin

import "testing"

func TestAuthenticationService_SetAccessToken(t *testing.T) {

	mockedClient, err := startMockClient("")
	if err != nil {
		t.Log(err)
	}

	type fields struct {
		client      *Client
		beaverToken string
		agent       string
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
			name: "SetAccessTokenWhenTheParamsAreCorrect",
			fields: fields{
				client: mockedClient,
			},
			args: args{
				token: "$TOKEN_ID",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthenticationService{
				client:      tt.fields.client,
				beaverToken: tt.fields.beaverToken,
				agent:       tt.fields.agent,
			}

			a.SetBearerToken(tt.args.token)
		})
	}
}

func TestAuthenticationService_SetUserAgent(t *testing.T) {

	mockedClient, err := startMockClient("")
	if err != nil {
		t.Log(err)
	}

	type fields struct {
		client      *Client
		beaverToken string
		agent       string
	}
	type args struct {
		agent string
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
				agent: "$USER_AGENT_SAMPLE",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AuthenticationService{
				client:      tt.fields.client,
				beaverToken: tt.fields.beaverToken,
				agent:       tt.fields.agent,
			}

			a.SetUserAgent(tt.args.agent)
		})
	}
}
