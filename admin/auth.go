package admin

type AuthenticationService struct {
	client      *Client
	beaverToken string
	agent       string
}

func (a *AuthenticationService) SetBearerToken(token string) {
	a.beaverToken = token
}

func (a *AuthenticationService) SetUserAgent(agent string) {
	a.agent = agent
}
