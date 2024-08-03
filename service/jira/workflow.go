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

	// Bulk returns a list of workflows and related statuses by providing workflow names,
	//
	// workflow IDs, or project and issue types.
	//
	// POST /rest/api/{2-3}/workflows
	Bulk(ctx context.Context, options *model.WorkflowBulkOptionsScheme, expand []string) (*model.WorkflowReadResponseScheme, *model.ResponseScheme, error)

	// Capabilities get the list of workflow capabilities for a specific workflow using either the workflow ID, or the project and issue type ID pair.
	//
	// The response includes the scope of the workflow, defined as global/project-based,
	//
	// and a list of project types that the workflow is scoped to.
	//
	// It also includes all rules organised into their broad categories (conditions, validators, actions, triggers, screens)
	//
	// as well as the source location (Atlassian-provided, Connect, Forge).
	//
	// GET /rest/api/{2-3}/workflows/capabilities
	Capabilities(ctx context.Context, workflowID, projectID, issueTypeID string) (*model.WorkflowCapabilitiesScheme, *model.ResponseScheme, error)

	// Creates creates workflows and related statuses.
	//
	// POST /rest/api/{2-3}/workflows/create
	Creates(ctx context.Context, payload *model.WorkflowCreatesPayloadScheme) (*model.WorkflowCreateResponseScheme, *model.ResponseScheme, error)

	// Updates updates workflows and related statuses.
	//
	// POST /rest/api/{2-3}/workflows/update
	Updates(ctx context.Context, payload *model.WorkflowUpdatesPayloadScheme, expand []string) (*model.WorkflowUpdateResponseScheme, *model.ResponseScheme, error)
}

type WorkflowValidatorConnector interface {

	// Creation validates the payload for bulk create workflows.
	//
	// POST /rest/api/{2-3}/workflows/create/validation
	Creation(ctx context.Context, payload *model.WorkflowCreateValidatorPayloadScheme) (*model.WorkflowValidationErrorListScheme, *model.ResponseScheme, error)

	// Modification validates the payload for bulk update workflows.
	//
	// POST /rest/api/{2-3}/workflows/update/validation
	Modification(ctx context.Context, payload *model.WorkflowUpdateValidatorPayloadScheme) (*model.WorkflowValidationErrorListScheme, *model.ResponseScheme, error)
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

// WorkflowSchemeIssueTypeConnector represents the workflows scheme issue type endpoints.
//
// Use it to search, get, create, delete, and change workflow issue types.
type WorkflowSchemeIssueTypeConnector interface {

	// Get returns the issue type-workflow mapping for an issue type in a workflow scheme.
	//
	// GET /rest/api/{2-3}/workflowscheme/{id}/issuetype/{issueType}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme/issue-type#get-workflow-for-issue-type-in-workflow-scheme
	Get(ctx context.Context, schemeID int, issueTypeID string, returnDraft bool) (*model.IssueTypeWorkflowMappingScheme, *model.ResponseScheme, error)

	// Set sets the workflow for an issue type in a workflow scheme.
	//
	// Note that active workflow schemes cannot be edited.
	//
	// If the workflow scheme is active, set updateDraftIfNeeded to true in the request body and a draft workflow scheme
	//
	// is created or updated with the new issue type-workflow mapping.
	//
	// The draft workflow scheme can be published in Jira.
	//
	// PUT /rest/api/{2-3}/workflowscheme/{id}/issuetype/{issueType}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme/issue-type#set-workflow-for-issue-type-in-workflow-scheme
	Set(ctx context.Context, schemeID int, issueTypeID string, payload *model.IssueTypeWorkflowPayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error)

	// Delete deletes the issue type-workflow mapping for an issue type in a workflow scheme.
	//
	// Note that active workflow schemes cannot be edited.
	//
	// If the workflow scheme is active, set updateDraftIfNeeded to true and a draft workflow scheme is created or
	//
	// updated with the issue type-workflow mapping deleted.
	//
	// The draft workflow scheme can be published in Jira.
	//
	// DELETE /rest/api/{2-3}/workflowscheme/{id}/issuetype/{issueType}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme/issue-type#delete-workflow-for-issue-type-in-workflow-scheme
	Delete(ctx context.Context, schemeID int, issueTypeID string, updateDraft bool) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error)

	// Mapping returns the workflow-issue type mappings for a workflow scheme.
	//
	// GET /rest/api/{2-3}/workflowscheme/{id}/workflow
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme/issue-type#get-issue-types-for-workflows-in-workflow-scheme
	Mapping(ctx context.Context, schemeID int, workflowName string, returnDraft bool) ([]*model.IssueTypesWorkflowMappingScheme, *model.ResponseScheme, error)
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

	// Get returns a status.
	//
	// The status must be associated with an active workflow to be returned.
	//
	// If a name is used on more than one status, only the status found first is returned.
	//
	// Therefore, identifying the status by its ID may be preferable.
	//
	// This operation can be accessed anonymously.
	//
	// GET /rest/api/{2-3}/status/{idOrName}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#get-workflow-status
	Get(ctx context.Context, idOrName string) (*model.StatusDetailScheme, *model.ResponseScheme, error)
}
