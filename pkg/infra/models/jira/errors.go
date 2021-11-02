package jira

import (
	"errors"
)

var (
	ErrNoApplicationRoleError = errors.New("jira: no application role key set")
	ErrNoDashboardIDError     = errors.New("jira: no dashboard id set")
	ErrNoGroupNameError       = errors.New("jira: no group name set")
	ErrNoGroupIDError         = errors.New("jira: no group name set")
	ErrNoIssueKeyOrIDError    = errors.New("jira: no issue key/id set")
	ErrNoIssueSchemeError     = errors.New("jira: no jira.IssueScheme set")
	ErrNoTransitionIDError    = errors.New("jira: no transition id set")
)
