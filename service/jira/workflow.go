package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type WorkflowConnector interface {

	// Create creates a workflow.
	//
	// You can define transition rules using the shapes detailed in the following sections.
	//
	// If no transitional rules are specified the default system transition rules are used.
	//
	// POST /rest/api/{2-3}/workflow
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow#create-workflow
	Create(ctx context.Context, payload *model.WorkflowPayloadScheme) (*model.WorkflowCreatedResponseScheme, *model.ResponseScheme, error)

	// Gets returns a paginated list of published classic workflows.
	//
	// When workflow names are specified, details of those workflows are returned.
	//
	// Otherwise, all published classic workflows are returned.
	//
	// GET /rest/api/{2-3}/workflow/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow#search-workflows
	Gets(ctx context.Context, options *model.WorkflowSearchOptions, startAt, maxResults int) (*model.WorkflowPageScheme, *model.ResponseScheme, error)

	// Delete deletes a workflow.
	//
	// The workflow cannot be deleted if it is:
	//
	// 1. an active workflow.
	// 2. a system workflow.
	// 3. associated with any workflow scheme.
	// 4. associated with any draft workflow scheme.
	//
	// DELETE /rest/api/{2-3}/workflow/{entityId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow#search-workflows
	Delete(ctx context.Context, workflowId string) (*model.ResponseScheme, error)
}

type WorkflowSchemeConnector interface {

	// Gets returns a paginated list of all workflow schemes, not including draft workflow schemes.
	//
	// GET /rest/api/{2-3}/workflowscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#gets-workflows-schemes
	Gets(ctx context.Context, startAt, maxResults int) (*model.WorkflowSchemePageScheme, *model.ResponseScheme, error)

	// Create creates a workflow scheme.
	//
	// POST /rest/api/{2-3}/workflowscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#create-workflows-scheme
	Create(ctx context.Context, payload *model.WorkflowSchemePayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error)

	// Get returns a workflow scheme
	//
	// GET /rest/api/{2-3}/workflowscheme/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#get-workflow-scheme
	Get(ctx context.Context, schemeId int, returnDraftIfExists bool) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error)

	// Update updates a workflow scheme, including the name, default workflow, issue type to project mappings, and more.
	//
	// If the workflow scheme is active (that is, being used by at least one project), then a draft workflow scheme is created or updated instead,
	//
	// provided that updateDraftIfNeeded is set to true.
	//
	// PUT /rest/api/{2-3}/workflowscheme/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#update-workflow-scheme
	Update(ctx context.Context, schemeId int, payload *model.WorkflowSchemePayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error)

	// Delete deletes a workflow scheme.
	//
	// Note that a workflow scheme cannot be deleted if it is active (that is, being used by at least one project).
	//
	// DELETE /rest/api/{2-3}/workflowscheme/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#delete-workflow-scheme
	Delete(ctx context.Context, schemeId int) (*model.ResponseScheme, error)

	// Associations returns a list of the workflow schemes associated with a list of projects.
	//
	// Each returned workflow scheme includes a list of the requested projects associated with it.
	//
	// Any team-managed or non-existent projects in the request are ignored and no errors are returned.
	//
	// If the project is associated with the Default Workflow Scheme no ID is returned.
	//
	// This is because the way the Default Workflow Scheme is stored means it has no ID.
	//
	// GET /rest/api/{2-3}/workflowscheme/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#get-workflow-schemes-associations
	Associations(ctx context.Context, projectIds []int) (*model.WorkflowSchemeAssociationPageScheme, *model.ResponseScheme, error)

	// Assign assigns a workflow scheme to a project.
	//
	// This operation is performed only when there are no issues in the project.
	//
	// Workflow schemes can only be assigned to classic projects.
	//
	// PUT /rest/api/{2-3}/workflowscheme/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#get-workflow-schemes-associations
	Assign(ctx context.Context, schemeId, projectId string) (*model.ResponseScheme, error)
}

// WorkflowStatusConnector represents the workflows statuses.
//
// Use it to search, get, create, delete, and change statuses.
type WorkflowStatusConnector interface {

	// Gets returns a list of the statuses specified by one or more status IDs.
	//
	// GET /rest/api/{2-3}/statuses
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#gets-workflow-statuses
	Gets(ctx context.Context, ids, expand []string) ([]*model.WorkflowStatusDetailScheme, *model.ResponseScheme, error)

	// Update updates statuses by ID.
	//
	// PUT /rest/api/{2-3}/statuses
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#update-workflow-statuses
	Update(ctx context.Context, payload *model.WorkflowStatusPayloadScheme) (*model.ResponseScheme, error)

	// Create creates statuses for a global or project scope.
	//
	// POST /rest/api/{2-3}/statuses
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#create-workflow-statuses
	Create(ctx context.Context, payload *model.WorkflowStatusPayloadScheme) ([]*model.WorkflowStatusDetailScheme, *model.ResponseScheme, error)

	// Delete deletes statuses by ID.
	//
	// DELETE /rest/api/{2-3}/statuses
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#delete-workflow-statuses
	Delete(ctx context.Context, ids []string) (*model.ResponseScheme, error)

	// Search returns a paginated list of statuses that match a search on name or project.
	//
	// GET /rest/api/{2-3}/statuses/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#search-workflow-statuses
	Search(ctx context.Context, options *model.WorkflowStatusSearchParams, startAt, maxResults int) (*model.WorkflowStatusDetailPageScheme,
		*model.ResponseScheme, error)

	// Bulk returns a list of all statuses associated with active workflows.
	//
	// GET /rest/api/{2-3}/status
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#bulk-workflow-statuses
	Bulk(ctx context.Context) ([]*model.StatusDetailScheme, *model.ResponseScheme, error)
}
