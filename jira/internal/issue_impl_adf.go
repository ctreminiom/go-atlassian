package internal

import (
	"context"
	"errors"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"github.com/imdario/mergo"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type IssueADFService struct {
	internalClient jira.IssueADFConnector
	Attachment     *IssueAttachmentService
	Comment        *CommentADFService
	Field          *IssueFieldService
	Label          *LabelService
	Link           *LinkADFService
	Metadata       *MetadataService
	Priority       *PriorityService
	Resolution     *ResolutionService
	Search         *SearchADFService
	Type           *TypeService
	Vote           *VoteService
	Watcher        *WatcherService
	Worklog        *WorklogADFService
}

// Delete deletes an issue.
//
// 1.An issue cannot be deleted if it has one or more subtasks.
//
// 2.To delete an issue with subtasks, set deleteSubtasks.
//
// 3.This causes the issue's subtasks to be deleted with the issue.
//
// DELETE /rest/api/{2-3}/issue/{issueIdOrKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#delete-issue
func (i *IssueADFService) Delete(ctx context.Context, issueKeyOrId string, deleteSubTasks bool) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, issueKeyOrId, deleteSubTasks)
}

// Assign assigns an issue to a user.
//
// Use this operation when the calling user does not have the Edit Issues permission but has the
//
// Assign issue permission for the project that the issue is in.
//
// If accountId is set to:
//
//  1. "-1", the issue is assigned to the default assignee for the project.
//  2. null, the issue is set to unassigned.
//
// PUT /rest/api/{2-3}/issue/{issueIdOrKey}/assignee
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#assign-issue
func (i *IssueADFService) Assign(ctx context.Context, issueKeyOrId, accountId string) (*model.ResponseScheme, error) {
	return i.internalClient.Assign(ctx, issueKeyOrId, accountId)
}

// Notify creates an email notification for an issue and adds it to the mail queue.
//
// POST /rest/api/{2-3}/issue/{issueIdOrKey}/notify
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#send-notification-for-issue
func (i *IssueADFService) Notify(ctx context.Context, issueKeyOrId string, options *model.IssueNotifyOptionsScheme) (*model.ResponseScheme, error) {
	return i.internalClient.Notify(ctx, issueKeyOrId, options)
}

// Transitions returns either all transitions or a transition that can be performed by the user on an issue, based on the issue's status.
//
// Note, if a request is made for a transition that does not exist or cannot be performed on the issue,
//
// given its status, the response will return any empty transitions list.
//
// GET /rest/api/{2-3}/issue/{issueIdOrKey}/transitions
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#get-transitions
func (i *IssueADFService) Transitions(ctx context.Context, issueKeyOrId string) (*model.IssueTransitionsScheme, *model.ResponseScheme, error) {
	return i.internalClient.Transitions(ctx, issueKeyOrId)
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
// GET /rest/api/{2-3}/issue/{issueIdOrKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#get-issue
func (i *IssueADFService) Get(ctx context.Context, issueKeyOrId string, fields, expand []string) (*model.IssueScheme, *model.ResponseScheme, error) {
	return i.internalClient.Get(ctx, issueKeyOrId, fields, expand)
}

// Update edits an issue.
//
// Edits an issue. A transition may be applied and issue properties updated as part of the edit.
//
// The edits to the issue's fields are defined using update and fields
//
// PUT /rest/api/{2-3}/issue/{issueIdOrKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#edit-issue
func (i *IssueADFService) Update(ctx context.Context, issueKeyOrId string, notify bool, payload *model.IssueScheme, customFields *model.CustomFields, operations *model.UpdateOperations) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, issueKeyOrId, notify, payload, customFields, operations)
}

// Move performs an issue transition and, if the transition has a screen, updates the fields from the transition screen.
//
// sortByCategory To update the fields on the transition screen, specify the fields in the fields or update parameters in the request body. Get details about the fields using Get transitions with the transitions.fields expand.
//
// POST /rest/api/{2-3}/issue/{issueIdOrKey}/transitions
//
// https://docs.go-atlassian.io/jira-software-cloud/issues#transition-issue
func (i *IssueADFService) Move(ctx context.Context, issueKeyOrId, transitionId string, options *model.IssueMoveOptionsV3) (*model.ResponseScheme, error) {
	return i.internalClient.Move(ctx, issueKeyOrId, transitionId, options)
}

type internalIssueADFServiceImpl struct {
	c       service.Client
	version string
}

func (i *internalIssueADFServiceImpl) Delete(ctx context.Context, issueKeyOrId string, deleteSubTasks bool) (*model.ResponseScheme, error) {
	return deleteIssue(ctx, i.c, i.version, issueKeyOrId, deleteSubTasks)
}

func (i *internalIssueADFServiceImpl) Assign(ctx context.Context, issueKeyOrId, accountId string) (*model.ResponseScheme, error) {
	return assignIssue(ctx, i.c, i.version, issueKeyOrId, accountId)
}

func (i *internalIssueADFServiceImpl) Notify(ctx context.Context, issueKeyOrId string, options *model.IssueNotifyOptionsScheme) (*model.ResponseScheme, error) {
	return sendNotification(ctx, i.c, i.version, issueKeyOrId, options)
}

func (i *internalIssueADFServiceImpl) Transitions(ctx context.Context, issueKeyOrId string) (*model.IssueTransitionsScheme, *model.ResponseScheme, error) {
	return getTransitions(ctx, i.c, i.version, issueKeyOrId)
}

func (i *internalIssueADFServiceImpl) Create(ctx context.Context, payload *model.IssueScheme, customFields *model.CustomFields) (*model.IssueResponseScheme, *model.ResponseScheme, error) {

	var reader io.Reader
	var err error

	if customFields != nil {

		payloadUpdated, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, nil, err
		}

		reader, err = i.c.TransformStructToReader(payloadUpdated)
		if err != nil {
			return nil, nil, err
		}

	} else {

		reader, err = i.c.TransformStructToReader(payload)
		if err != nil {
			return nil, nil, err
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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
		return nil, nil, errors.New("error, please provide a valid []*IssueBulkScheme slice of pointers")
		// TODO: The errors when the bulk creates does not contains values needs to be parsed and moved to the model package
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

	var bulkPayload = map[string]interface{}{}
	bulkPayload["issueUpdates"] = issuePayloads

	reader, err := i.c.TransformStructToReader(&bulkPayload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/bulk", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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

func (i *internalIssueADFServiceImpl) Get(ctx context.Context, issueKeyOrId string, fields, expand []string) (*model.IssueScheme, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if len(fields) != 0 {
		params.Add("fields", strings.Join(fields, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v", i.version, issueKeyOrId))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
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

func (i *internalIssueADFServiceImpl) Update(ctx context.Context, issueKeyOrId string, notify bool, payload *model.IssueScheme, customFields *model.CustomFields, operations *model.UpdateOperations) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("notifyUsers", fmt.Sprintf("%v", notify))
	endpoint := fmt.Sprintf("rest/api/%v/issue/%v?%v", i.version, issueKeyOrId, params.Encode())

	var reader io.Reader
	var err error

	// Executed when customfields and operations are not provided
	if customFields == nil && operations == nil {

		reader, err = i.c.TransformStructToReader(payload)
		if err != nil {
			return nil, err
		}
	}

	// Executed when customfields and operation are provided
	if customFields != nil && operations != nil {

		payloadUpdated, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, err
		}

		payloadWithOperations, err := payload.MergeOperations(operations)
		if err != nil {
			return nil, err
		}

		if err := mergo.Map(&payloadUpdated, &payloadWithOperations, mergo.WithOverride); err != nil {
			return nil, err
		}

		reader, err = i.c.TransformStructToReader(&payloadUpdated)
		if err != nil {
			return nil, err
		}
	}

	// Executed when customfields are provided but not the operations
	if customFields != nil && operations == nil {

		payloadUpdated, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, err
		}

		reader, err = i.c.TransformStructToReader(&payloadUpdated)
		if err != nil {
			return nil, err
		}
	}

	// Executed when operations are provided but not the customfields
	if customFields == nil && operations != nil {

		payloadUpdated, err := payload.MergeOperations(operations)
		if err != nil {
			return nil, err
		}

		reader, err = i.c.TransformStructToReader(&payloadUpdated)
		if err != nil {
			return nil, err
		}
	}

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueADFServiceImpl) Move(ctx context.Context, issueKeyOrId, transitionId string, options *model.IssueMoveOptionsV3) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if transitionId == "" {
		return nil, model.ErrNoTransitionIDError
	}

	payloadUpdated := make(map[string]interface{})
	payloadUpdated["transition"] = map[string]interface{}{"id": transitionId}

	var reader io.Reader
	var err error

	if options != nil && options.Fields != nil {

		// Executed when customfields and operation are provided
		if options.CustomFields != nil && options.Operations != nil {

			payloadWithCustomFields, err := options.Fields.MergeCustomFields(options.CustomFields)
			if err != nil {
				return nil, err
			}

			payloadWithOperations, err := options.Fields.MergeOperations(options.Operations)
			if err != nil {
				return nil, err
			}

			if err := mergo.Map(&payloadWithCustomFields, &payloadWithOperations, mergo.WithOverride); err != nil {
				return nil, err
			}

			if err := mergo.Map(&payloadWithCustomFields, &payloadUpdated, mergo.WithOverride); err != nil {
				return nil, err
			}

			reader, err = i.c.TransformStructToReader(&payloadWithCustomFields)
			if err != nil {
				return nil, err
			}
		}

		// Executed when customfields are provided but not the operations
		if options.CustomFields != nil && options.Operations == nil {

			payloadWithCustomFields, err := options.Fields.MergeCustomFields(options.CustomFields)
			if err != nil {
				return nil, err
			}

			if err := mergo.Map(&payloadWithCustomFields, &payloadUpdated, mergo.WithOverride); err != nil {
				return nil, err
			}

			reader, err = i.c.TransformStructToReader(&payloadWithCustomFields)
			if err != nil {
				return nil, err
			}
		}

		// Executed when operations are provided but not the customfields
		if options.CustomFields == nil && options.Operations != nil {

			payloadWithOperations, err := options.Fields.MergeOperations(options.Operations)
			if err != nil {
				return nil, err
			}

			if err := mergo.Map(&payloadWithOperations, &payloadUpdated, mergo.WithOverride); err != nil {
				return nil, err
			}

			reader, err = i.c.TransformStructToReader(&payloadWithOperations)
			if err != nil {
				return nil, err
			}
		}
	} else {
		reader, err = i.c.TransformStructToReader(&payloadUpdated)
		if err != nil {
			return nil, err
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/transitions", i.version, issueKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
