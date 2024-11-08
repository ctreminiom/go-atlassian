package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"dario.cat/mergo"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// IssueADFService provides methods to manage issues in Jira Service Management using the ADF format.
type IssueADFService struct {
	// internalClient is the connector interface for ADF issue operations.
	internalClient jira.IssueADFConnector
	// Attachment is the service for managing issue attachments.
	Attachment *IssueAttachmentService
	// Comment is the service for managing ADF comments.
	Comment *CommentADFService
	// Field is the service for managing issue fields.
	Field *IssueFieldService
	// Label is the service for managing issue labels.
	Label *LabelService
	// Link is the service for managing ADF issue links.
	Link *LinkADFService
	// Metadata is the service for managing issue metadata.
	Metadata *MetadataService
	// Priority is the service for managing issue priorities.
	Priority *PriorityService
	// Resolution is the service for managing issue resolutions.
	Resolution *ResolutionService
	// Search is the service for managing ADF issue searches.
	Search *SearchADFService
	// Type is the service for managing issue types.
	Type *TypeService
	// Vote is the service for managing issue votes.
	Vote *VoteService
	// Watcher is the service for managing issue watchers.
	Watcher *WatcherService
	// Worklog is the service for managing ADF worklogs.
	Worklog *WorklogADFService
	// Property is the service for managing issue properties.
	Property *IssuePropertyService
}

// Delete deletes an issue.
//
// 1.An issue cannot be deleted if it has one or more subtasks.
//
// 2.To delete an issue with subtasks, set deleteSubtasks.
//
// 3.This causes the issue's subtasks to be deleted with the issue.
//
// DELETE /rest/api/{2-3}/issue/{issueKeyOrID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#delete-issue
func (i *IssueADFService) Delete(ctx context.Context, issueKeyOrID string, deleteSubTasks bool) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, issueKeyOrID, deleteSubTasks)
}

// Assign assigns an issue to a user.
//
// # Use this operation when the calling user does not have the Edit Issues permission but has the
//
// Assign issue permission for the project that the issue is in.
//
// If accountID is set to:
//
//  1. "-1", the issue is assigned to the default assignee for the project.
//  2. null, the issue is set to unassigned.
//
// PUT /rest/api/{2-3}/issue/{issueKeyOrID}/assignee
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#assign-issue
func (i *IssueADFService) Assign(ctx context.Context, issueKeyOrID, accountID string) (*model.ResponseScheme, error) {
	return i.internalClient.Assign(ctx, issueKeyOrID, accountID)
}

// Notify creates an email notification for an issue and adds it to the mail queue.
//
// POST /rest/api/{2-3}/issue/{issueKeyOrID}/notify
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#send-notification-for-issue
func (i *IssueADFService) Notify(ctx context.Context, issueKeyOrID string, options *model.IssueNotifyOptionsScheme) (*model.ResponseScheme, error) {
	return i.internalClient.Notify(ctx, issueKeyOrID, options)
}

// Transitions returns either all transitions or a transition that can be performed by the user on an issue, based on the issue's status.
//
// Note, if a request is made for a transition that does not exist or cannot be performed on the issue,
//
// given its status, the response will return any empty transitions list.
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/transitions
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#get-transitions
func (i *IssueADFService) Transitions(ctx context.Context, issueKeyOrID string) (*model.IssueTransitionsScheme, *model.ResponseScheme, error) {
	return i.internalClient.Transitions(ctx, issueKeyOrID)
}

// Create creates an issue or, where the option to create subtasks is enabled in Jira, a subtask.
//
// POST /rest/api/{2-3}/issue
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#create-issue
func (i *IssueADFService) Create(ctx context.Context, payload *model.IssueScheme, customFields *model.CustomFields) (*model.IssueResponseScheme, *model.ResponseScheme, error) {
	return i.internalClient.Create(ctx, payload, customFields)
}

// Creates issues and, where the option to create subtasks is enabled in Jira, subtasks.
//
// 1.Creates upto 50 issues and, where the option to create subtasks is enabled in Jira, subtasks.
//
// 2.Transitions may be applied, to move the issues or subtasks to a workflow step other than the default start step, and issue properties set.
//
// POST /rest/api/{2-3}/issue/bulk
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#bulk-create-issue
func (i *IssueADFService) Creates(ctx context.Context, payload []*model.IssueBulkSchemeV3) (*model.IssueBulkResponseScheme, *model.ResponseScheme, error) {
	return i.internalClient.Creates(ctx, payload)
}

// Get returns the details for an issue.
//
// The issue is identified by its ID or key, however, if the identifier doesn't match an issue, a case-insensitive search
//
// and check for moved issues is performed. If a matching issue is found its details are returned, a 302 or other redirect is not returned.
//
// The issue key returned to the response is the key of the issue found.
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#get-issue
func (i *IssueADFService) Get(ctx context.Context, issueKeyOrID string, fields, expand []string) (*model.IssueScheme, *model.ResponseScheme, error) {
	return i.internalClient.Get(ctx, issueKeyOrID, fields, expand)
}

// Update edits an issue.
//
// Edits an issue. A transition may be applied and issue properties updated as part of the edit.
//
// # The edits to the issue's fields are defined using update and fields
//
// PUT /rest/api/{2-3}/issue/{issueKeyOrID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#edit-issue
func (i *IssueADFService) Update(ctx context.Context, issueKeyOrID string, notify bool, payload *model.IssueScheme, customFields *model.CustomFields, operations *model.UpdateOperations) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, issueKeyOrID, notify, payload, customFields, operations)
}

// Move performs an issue transition and, if the transition has a screen, updates the fields from the transition screen.
//
// sortByCategory To update the fields on the transition screen, specify the fields in the fields or update parameters in the request body. Get details about the fields using Get transitions with the transitions.fields expand.
//
// POST /rest/api/{2-3}/issue/{issueKeyOrID}/transitions
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#transition-issue
func (i *IssueADFService) Move(ctx context.Context, issueKeyOrID, transitionID string, options *model.IssueMoveOptionsV3) (*model.ResponseScheme, error) {
	return i.internalClient.Move(ctx, issueKeyOrID, transitionID, options)
}

type internalIssueADFServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueADFServiceImpl) Delete(ctx context.Context, issueKeyOrID string, deleteSubTasks bool) (*model.ResponseScheme, error) {
	return deleteIssue(ctx, i.c, i.version, issueKeyOrID, deleteSubTasks)
}

func (i *internalIssueADFServiceImpl) Assign(ctx context.Context, issueKeyOrID, accountID string) (*model.ResponseScheme, error) {
	return assignIssue(ctx, i.c, i.version, issueKeyOrID, accountID)
}

func (i *internalIssueADFServiceImpl) Notify(ctx context.Context, issueKeyOrID string, options *model.IssueNotifyOptionsScheme) (*model.ResponseScheme, error) {
	return sendNotification(ctx, i.c, i.version, issueKeyOrID, options)
}

func (i *internalIssueADFServiceImpl) Transitions(ctx context.Context, issueKeyOrID string) (*model.IssueTransitionsScheme, *model.ResponseScheme, error) {
	return getTransitions(ctx, i.c, i.version, issueKeyOrID)
}

func (i *internalIssueADFServiceImpl) Create(ctx context.Context, payload *model.IssueScheme, customFields *model.CustomFields) (*model.IssueResponseScheme, *model.ResponseScheme, error) {
	var body interface{} = payload
	var err error

	if customFields != nil && len(customFields.Fields) != 0 {

		payloadWithFields, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, nil, err
		}

		body = payloadWithFields
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue", i.version)
	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", body)
	if err != nil {
		return nil, nil, err
	}

	issue := new(model.IssueResponseScheme)
	response, err := i.c.Call(request, issue)
	if err != nil {
		return nil, response, err
	}

	return issue, response, nil
}

func (i *internalIssueADFServiceImpl) Creates(ctx context.Context, payload []*model.IssueBulkSchemeV3) (*model.IssueBulkResponseScheme, *model.ResponseScheme, error) {

	if len(payload) == 0 {
		return nil, nil, model.ErrNoCreateIssues
	}

	var issuePayloads []map[string]interface{}
	for _, newIssue := range payload {

		if newIssue.Payload == nil {
			continue
		}

		issuePayload, err := newIssue.Payload.MergeCustomFields(newIssue.CustomFields)
		if err != nil {
			return nil, nil, err
		}

		issuePayloads = append(issuePayloads, issuePayload)
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/bulk", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"issueUpdates": issuePayloads})
	if err != nil {
		return nil, nil, err
	}

	issues := new(model.IssueBulkResponseScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}

func (i *internalIssueADFServiceImpl) Get(ctx context.Context, issueKeyOrID string, fields, expand []string) (*model.IssueScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	params := url.Values{}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if len(fields) != 0 {
		params.Add("fields", strings.Join(fields, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v", i.version, issueKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	issue := new(model.IssueScheme)
	response, err := i.c.Call(request, issue)
	if err != nil {
		return nil, response, err
	}

	return issue, response, nil
}

func (i *internalIssueADFServiceImpl) Update(ctx context.Context, issueKeyOrID string, notify bool, payload *model.IssueScheme, customFields *model.CustomFields, operations *model.UpdateOperations) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	params := url.Values{}
	params.Add("notifyUsers", fmt.Sprintf("%v", notify))
	endpoint := fmt.Sprintf("rest/api/%v/issue/%v?%v", i.version, issueKeyOrID, params.Encode())

	withCustomFields := customFields != nil
	withOperations := operations != nil

	var err error
	payloadUpdated := make(map[string]interface{})

	// Executed when the customfields and operations are provided
	if withCustomFields && withOperations {

		payloadUpdated, err = payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, err
		}

		payloadWithOperations, err := payload.MergeOperations(operations)
		if err != nil {
			return nil, err
		}

		if err = mergo.Map(&payloadUpdated, &payloadWithOperations, mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	// Executed when only the customfields are provided, but not the operations
	if withCustomFields && !withOperations {

		payloadUpdated, err = payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, err
		}
	}

	// Executed when only the operations are provided, but not the customfields
	if withOperations && !withCustomFields {

		payloadUpdated, err = payload.MergeOperations(operations)
		if err != nil {
			return nil, err
		}
	}

	// After the payload transformation, validate if the shadowed payloadUpdated variable contains data
	var request *http.Request

	if len(payloadUpdated) != 0 {

		request, err = i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payloadUpdated)
		if err != nil {
			return nil, err
		}
	} else {
		request, err = i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
		if err != nil {
			return nil, err
		}
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueADFServiceImpl) Move(ctx context.Context, issueKeyOrID, transitionID string, options *model.IssueMoveOptionsV3) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	if transitionID == "" {
		return nil, model.ErrNoTransitionID
	}

	payload := map[string]interface{}{"transition": map[string]interface{}{"id": transitionID}}

	if options != nil {
		if options.Fields == nil {
			return nil, model.ErrNoIssueScheme
		}

		// Merge the customfields and operations
		payloadWithFields, err := options.Fields.MergeCustomFields(options.CustomFields)
		if err != nil {
			return nil, err
		}
		if err = mergo.Map(&payload, &payloadWithFields, mergo.WithOverride); err != nil {
			return nil, err
		}

		payloadWithOperation, err := options.Fields.MergeOperations(options.Operations)
		if err != nil {
			return nil, err
		}
		if err = mergo.Map(&payload, &payloadWithOperation, mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/transitions", i.version, issueKeyOrID)
	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
