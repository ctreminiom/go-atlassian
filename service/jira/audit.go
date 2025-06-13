package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// AuditRecordConnector  audits that record activities undertaken in Jira.
// Use it to get a list of audit records.
type AuditRecordConnector interface {

	// Get allows you to retrieve the audit records for specific activities that have occurred within Jira.
	//
	// GET /rest/api/{2-3}/auditing/record
	//
	// https://docs.go-atlassian.io/jira-software-cloud/audit-records#get-audit-records
	Get(ctx context.Context, options *model.AuditRecordGetOptions, offSet, limit int) (*model.AuditRecordPageScheme, *model.ResponseScheme, error)
}
