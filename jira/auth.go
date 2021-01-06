package jira

type AuthenticationService struct {
	client      *Client
	mail, token string
}

func (a *AuthenticationService) SetBasicAuth(mail, token string) {
	a.mail = mail
	a.token = token
}
