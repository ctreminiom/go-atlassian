package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strings"
)

func NewWorkflowSchemeIssueTypeService(client service.Client, version string) *WorkflowSchemeIssueTypeService {

	return &WorkflowSchemeIssueTypeService{
		internalClient: &internalWorkflowSchemeIssueTypeImpl{c: client, version: version},
	}
}

type WorkflowSchemeIssueTypeService struct {
	internalClient jira.WorkflowSchemeIssueTypeConnector
}

// Get returns the issue type-workflow mapping for an issue type in a workflow scheme.
//
// GET /rest/api/{2-3}/workflowscheme/{id}/issuetype/{issueType}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme/issue-type#get-workflow-for-issue-type-in-workflow-scheme
func (w *WorkflowSchemeIssueTypeService) Get(ctx context.Context, schemeID int, issueTypeID string, returnDraft bool) (*model.IssueTypeWorkflowMappingScheme, *model.ResponseScheme, error) {
	return w.internalClient.Get(ctx, schemeID, issueTypeID, returnDraft)
}

// Set sets the workflow for an issue type in a workflow scheme.
//
// Note that active workflow schemes cannot be edited.
//
// # If the workflow scheme is active, set updateDraftIfNeeded to true in the request body and a draft workflow scheme
//
// is created or updated with the new issue type-workflow mapping.
//
// The draft workflow scheme can be published in Jira.
//
// PUT /rest/api/{2-3}/workflowscheme/{id}/issuetype/{issueType}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme/issue-type#set-workflow-for-issue-type-in-workflow-scheme
func (w *WorkflowSchemeIssueTypeService) Set(ctx context.Context, schemeID int, issueTypeID string, payload *model.IssueTypeWorkflowPayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	return w.internalClient.Set(ctx, schemeID, issueTypeID, payload)
}

// Delete deletes the issue type-workflow mapping for an issue type in a workflow scheme.
//
// Note that active workflow schemes cannot be edited.
//
// # If the workflow scheme is active, set updateDraftIfNeeded to true and a draft workflow scheme is created or
//
// updated with the issue type-workflow mapping deleted.
//
// The draft workflow scheme can be published in Jira.
//
// DELETE /rest/api/{2-3}/workflowscheme/{id}/issuetype/{issueType}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme/issue-type#delete-workflow-for-issue-type-in-workflow-scheme
func (w *WorkflowSchemeIssueTypeService) Delete(ctx context.Context, schemeID int, issueTypeID string, updateDraft bool) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	return w.internalClient.Delete(ctx, schemeID, issueTypeID, updateDraft)
}

// Mapping returns the workflow-issue type mappings for a workflow scheme.
//
// GET /rest/api/{2-3}/workflowscheme/{id}/workflow
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme/issue-type#get-issue-types-for-workflows-in-workflow-scheme
func (w *WorkflowSchemeIssueTypeService) Mapping(ctx context.Context, schemeID int, workflowName string, returnDraft bool) ([]*model.IssueTypesWorkflowMappingScheme, *model.ResponseScheme, error) {
	return w.internalClient.Mapping(ctx, schemeID, workflowName, returnDraft)
}

type internalWorkflowSchemeIssueTypeImpl struct {
	c       service.Client
	version string
}

func (i *internalWorkflowSchemeIssueTypeImpl) Get(ctx context.Context, schemeID int, issueTypeID string, returnDraft bool) (*model.IssueTypeWorkflowMappingScheme, *model.ResponseScheme, error) {

	if schemeID == 0 {
		return nil, nil, model.ErrNoWorkflowSchemeIDError
	}

	if issueTypeID == "" {
		return nil, nil, model.ErrNoIssueTypeIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflowscheme/%v/issuetype/%v", i.version, schemeID, issueTypeID))

	if returnDraft {
		params := url.Values{}
		params.Add("returnDraftIfExists", "true")

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	mapping := new(model.IssueTypeWorkflowMappingScheme)
	response, err := i.c.Call(request, mapping)
	if err != nil {
		return nil, response, err
	}

	return mapping, response, nil
}

func (i *internalWorkflowSchemeIssueTypeImpl) Set(ctx context.Context, schemeID int, issueTypeID string, payload *model.IssueTypeWorkflowPayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {

	if schemeID == 0 {
		return nil, nil, model.ErrNoWorkflowSchemeIDError
	}

	if issueTypeID == "" {
		return nil, nil, model.ErrNoIssueTypeIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme/%v/issuetype/%v", i.version, schemeID, issueTypeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	scheme := new(model.WorkflowSchemeScheme)
	response, err := i.c.Call(request, scheme)
	if err != nil {
		return nil, response, err
	}

	return scheme, response, nil
}

func (i *internalWorkflowSchemeIssueTypeImpl) Delete(ctx context.Context, schemeID int, issueTypeID string, updateDraft bool) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {

	if schemeID == 0 {
		return nil, nil, model.ErrNoWorkflowSchemeIDError
	}

	if issueTypeID == "" {
		return nil, nil, model.ErrNoIssueTypeIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflowscheme/%v/issuetype/%v", i.version, schemeID, issueTypeID))

	if updateDraft {
		params := url.Values{}
		params.Add("updateDraftIfNeeded", "true")

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	scheme := new(model.WorkflowSchemeScheme)
	response, err := i.c.Call(request, scheme)
	if err != nil {
		return nil, response, err
	}

	return scheme, response, nil
}

func (i *internalWorkflowSchemeIssueTypeImpl) Mapping(ctx context.Context, schemeID int, workflowName string, returnDraft bool) ([]*model.IssueTypesWorkflowMappingScheme, *model.ResponseScheme, error) {

	if schemeID == 0 {
		return nil, nil, model.ErrNoWorkflowSchemeIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflowscheme/%v/workflow", i.version, schemeID))

	params := url.Values{}

	if workflowName != "" {
		params.Add("workflowName", workflowName)
	}

	if returnDraft {
		params.Add("returnDraftIfExists", "true")
	}

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	var mapping []*model.IssueTypesWorkflowMappingScheme
	response, err := i.c.Call(request, mapping)
	if err != nil {
		return nil, response, err
	}

	return mapping, response, nil
}
