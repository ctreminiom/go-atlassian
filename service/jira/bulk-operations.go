package jira

import (
	"context"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type BulkOperationsConnector interface {
	Delete(ctx context.Context, issueKeysOrIDs []string, sendNotification bool)
	Edit(ctx context.Context, fields *models.JiraIssueFieldsScheme, selectedActions, selectedIssueKeyOrIDs []string, sendNotification bool)
	Transition(ctx context.Context, inputs []*models.BulkTransitionSubmitInputScheme, sendNotification bool)
	GetStatus(ctx context.Context, taskID string)
	GetTransitions(ctx context.Context, issueKeysOrIDs []string, startAt, cursor string)
}
