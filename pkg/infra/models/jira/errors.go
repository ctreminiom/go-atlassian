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
	ErrNoAttachmentIDError    = errors.New("jira: no attachment id set")
	ErrNoAttachmentNameError  = errors.New("jira: no attachment filename set")
	ErrNoReaderError          = errors.New("jira: no reader set")
	ErrNoCommentIDError       = errors.New("jira: no comment id set")
	ErrNoProjectIDError       = errors.New("jira: no project id set")
	ErrNoFieldIDError         = errors.New("jira: no field id set")
	ErrNoFieldContextIDError  = errors.New("jira: no field context id set")
	ErrNoIssueTypesError      = errors.New("jira: no issue types id's set")
	ErrNoProjectsError        = errors.New("jira: no projects set")
	ErrNoContextOptionIDError = errors.New("jira: no field context option id set")
	ErrNoTypeIDError          = errors.New("jira: no link id set")
	ErrNoLinkTypeIDError      = errors.New("jira: no link type id set")
	ErrPriorityIDError        = errors.New("jira: no priority id set")
)
