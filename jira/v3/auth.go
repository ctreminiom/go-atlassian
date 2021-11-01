package v3

type AuthenticationService struct {
	client *Client

	basicAuthProvided bool
	mail, token       string

	userAgentProvided bool
	agent             string
}

func (a *AuthenticationService) SetBasicAuth(mail, token string) {

	//Inject the auth credentials into the child modules
	if a.client.ServiceManagement != nil {
		a.client.ServiceManagement.Auth.SetBasicAuth(mail, token)
	}

	if a.client.Agile != nil {
		a.client.Agile.Auth.SetBasicAuth(mail, token)
	}

	a.mail = mail
	a.token = token

	a.basicAuthProvided = true
}

func (a *AuthenticationService) SetUserAgent(agent string) {

	if a.client.ServiceManagement != nil {
		a.client.ServiceManagement.Auth.SetUserAgent(agent)
	}

	a.agent = agent

	a.userAgentProvided = true
}
