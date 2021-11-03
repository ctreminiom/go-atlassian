package v3

type AuthenticationService struct {
	client *Client

	basicAuthProvided bool
	mail, token       string

	userAgentProvided bool
	agent             string
}

func (a *AuthenticationService) SetBasicAuth(mail, token string) {
	a.mail = mail
	a.token = token

	a.basicAuthProvided = true
}

func (a *AuthenticationService) SetUserAgent(agent string) {

	a.agent = agent
	a.userAgentProvided = true
}
