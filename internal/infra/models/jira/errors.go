package jira

import (
	"errors"
)

var (
	ErrNoApplicationRoleError = errors.New("jira-cloud: no application role key set")
	ErrNoDashboardIDError     = errors.New("jira-cloud: no dashboard id set")
)
