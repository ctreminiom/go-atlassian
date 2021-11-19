package v2

import (
	"context"
	"fmt"
	models2 "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/imdario/mergo"
	"net/http"
	"net/url"
	"strings"
)

type IssueService struct {
	client     *Client
	Attachment *AttachmentService
	Comment    *CommentService
	Field      *FieldService
	Link       *IssueLinkService
	Priority   *PriorityService
	Resolution *ResolutionService
	Type       *IssueTypeService
	Votes      *VoteService
	Watchers   *WatcherService
	Label      *LabelService
	Search     *IssueSearchService
	Worklog    *IssueWorklogService
	Metadata   *IssueMetadataService
}

// Create creates an issue or, where the option to create subtasks is enabled in Jira, a subtask.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#create-issue
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-post
func (i *IssueService) Create(ctx context.Context, payload *models2.IssueSchemeV2, customFields *models2.CustomFields) (
	result *IssueResponseScheme, response *ResponseScheme, err error) {

	var (
		endpoint = "rest/api/2/issue"
		request  *http.Request
	)

	if customFields != nil {

		payloadWithCustomFields, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, nil, err
		}

		payloadAsReader, _ := transformStructToReader(&payloadWithCustomFields)

		request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
		if err != nil {
			return nil, nil, err
		}
	} else {

		payloadAsReader, err := transformStructToReader(payload)
		if err != nil {
			return nil, nil, err
		}

		request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
		if err != nil {
			return nil, nil, err
		}
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type IssueResponseScheme struct {
	ID   string `json:"id,omitempty"`
	Key  string `json:"key,omitempty"`
	Self string `json:"self,omitempty"`
}

type IssueBulkScheme struct {
	Payload      *models2.IssueSchemeV2
	CustomFields *models2.CustomFields
}

// Creates issues and, where the option to create subtasks is enabled in Jira, subtasks.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#bulk-create-issue
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-bulk-post
func (i *IssueService) Creates(ctx context.Context, payload []*IssueBulkScheme) (result *IssueBulkResponseScheme,
	response *ResponseScheme, err error) {

	if len(payload) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid []*IssueBulkScheme slice of pointers")
	}

	var issuePayloadsNodeAsList []map[string]interface{}
	for pos, newIssue := range payload {

		if newIssue.Payload == nil {
			return nil, nil, fmt.Errorf("error, the IssueSchemeV2 payload #%v is nil, please provide a valid *IssueSchemeV2 pointer", pos)
		}

		//Convert the IssueSchemeV2 struct to map
		newIssueAsMap, err := newIssue.Payload.MergeCustomFields(newIssue.CustomFields)
		if err != nil {
			return nil, nil, err
		}

		issuePayloadsNodeAsList = append(issuePayloadsNodeAsList, newIssueAsMap)
	}

	var issueUpdatesNode = map[string]interface{}{}
	issueUpdatesNode["issueUpdates"] = issuePayloadsNodeAsList

	payloadAsReader, _ := transformStructToReader(&issueUpdatesNode)

	var endpoint = "rest/api/2/issue/bulk"

	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type BulkIssueSchemeV2 struct {
	Issues []*models2.IssueSchemeV2 `json:"issues,omitempty"`
}

type IssueBulkResponseScheme struct {
	Issues []struct {
		ID   string `json:"id,omitempty"`
		Key  string `json:"key,omitempty"`
		Self string `json:"self,omitempty"`
	} `json:"issues,omitempty"`
	Errors []*IssueBulkResponseErrorScheme `json:"errors,omitempty"`
}

type IssueBulkResponseErrorScheme struct {
	Status        int `json:"status"`
	ElementErrors struct {
		ErrorMessages []string `json:"errorMessages"`
		Status        int      `json:"status"`
	} `json:"elementErrors"`
	FailedElementNumber int `json:"failedElementNumber"`
}

// Get returns the details for an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#get-issue
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-get
func (i *IssueService) Get(ctx context.Context, issueKeyOrID string, fields []string, expand []string) (result *models2.IssueSchemeV2,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models2.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if len(fields) != 0 {
		params.Add("fields", strings.Join(fields, ","))
	}

	var endpointBuffer strings.Builder
	endpointBuffer.WriteString(fmt.Sprintf("rest/api/2/issue/%v", issueKeyOrID))

	if params.Encode() != "" {
		endpointBuffer.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.client.newRequest(ctx, http.MethodGet, endpointBuffer.String(), nil)
	if err != nil {
		return
	}

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update edits an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#edit-issue
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-put
func (i *IssueService) Update(ctx context.Context, issueKeyOrID string, notify bool, payload *models2.IssueSchemeV2,
	customFields *models2.CustomFields, operations *models2.UpdateOperations) (response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, models2.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	if !notify {
		params.Add("notifyUsers", "false")
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/2/issue/%v", issueKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	var request *http.Request

	// Executed when customfields or operation are not provided
	if customFields == nil && operations == nil {

		payloadAsReader, err := transformStructToReader(payload)
		if err != nil {
			return nil, err
		}

		request, err = i.client.newRequest(ctx, http.MethodPut, endpoint.String(), payloadAsReader)
		if err != nil {
			return nil, err
		}
	}

	// Executed when customfields and operation are provided
	if customFields != nil && operations != nil {

		payloadWithCustomFields, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, err
		}

		payloadWithOperations, err := payload.MergeOperations(operations)
		if err != nil {
			return nil, err
		}

		//Merge the map[string]interface{} into one
		_ = mergo.Map(&payloadWithCustomFields, &payloadWithOperations, mergo.WithOverride)

		payloadAsReader, _ := transformStructToReader(&payloadWithCustomFields)

		request, err = i.client.newRequest(ctx, http.MethodPut, endpoint.String(), payloadAsReader)
		if err != nil {
			return nil, err
		}

	}

	// Executed when customfields are provided but not the operations
	if customFields != nil && operations == nil {

		payloadWithCustomFields, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, err
		}

		payloadAsReader, _ := transformStructToReader(&payloadWithCustomFields)
		request, err = i.client.newRequest(ctx, http.MethodPut, endpoint.String(), payloadAsReader)
		if err != nil {
			return nil, err
		}

	}

	// Executed when operations are provided but not the customfields
	if customFields == nil && operations != nil {

		payloadWithOperations, err := payload.MergeOperations(operations)
		if err != nil {
			return nil, err
		}

		payloadAsReader, _ := transformStructToReader(&payloadWithOperations)
		request, err = i.client.newRequest(ctx, http.MethodPut, endpoint.String(), payloadAsReader)
		if err != nil {
			return nil, err
		}
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Delete deletes an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#delete-issue
func (i *IssueService) Delete(ctx context.Context, issueKeyOrID string, deleteSubTasks bool) (response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, models2.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	if deleteSubTasks {
		params.Add("deleteSubtasks", "true")
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/2/issue/%v", issueKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return
	}

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Assign assigns an issue to a user.
// Use this operation when the calling user does not have the Edit Issues permission but has the
// Assign issue permission for the project that the issue is in.
// If accountId is set to:
//  1. "-1", the issue is assigned to the default assignee for the project.
//  2. null, the issue is set to unassigned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#assign-issue
func (i *IssueService) Assign(ctx context.Context, issueKeyOrID, accountID string) (response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, models2.ErrNoIssueKeyOrIDError
	}

	payload := struct {
		AccountID string `json:"accountId"`
	}{
		AccountID: accountID,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("/rest/api/2/issue/%v/assignee", issueKeyOrID)

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Notify creates an email notification for an issue and adds it to the mail queue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#send-notification-for-issue
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-notify-post
func (i *IssueService) Notify(ctx context.Context, issueKeyOrID string, options *models2.IssueNotifyOptionsScheme) (
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, models2.ErrNoIssueKeyOrIDError
	}

	payloadAsReader, err := transformStructToReader(options)
	if err != nil {
		return nil, err
	}

	var endpoint = fmt.Sprintf("rest/api/2/issue/%v/notify", issueKeyOrID)

	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Transitions returns either all transitions or a transition that can be performed by the user on an issue, based on the issue's status.
// Note, if a request is made for a transition that does not exist or cannot be performed on the issue,
// given its status, the response will return any empty transitions list.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#get-transitions
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issues/#api-rest-api-2-issue-issueidorkey-transitions-get
func (i *IssueService) Transitions(ctx context.Context, issueKeyOrID string) (result *models2.IssueTransitionsScheme,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models2.ErrNoIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/issue/%v/transitions", issueKeyOrID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type IssueMoveOptions struct {
	Fields       *models2.IssueSchemeV2
	CustomFields *models2.CustomFields
	Operations   *models2.UpdateOperations
}

// Move performs an issue transition and
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#transition-issue
func (i *IssueService) Move(ctx context.Context, issueKeyOrID, transitionID string, options *IssueMoveOptions) (
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, models2.ErrNoIssueKeyOrIDError
	}

	if len(transitionID) == 0 {
		return nil, models2.ErrNoTransitionIDError
	}

	payloadWithTransition := make(map[string]interface{})
	payloadWithTransition["transition"] = map[string]interface{}{"id": transitionID}

	var (
		endpoint = fmt.Sprintf("rest/api/2/issue/%v/transitions", issueKeyOrID)
		request  *http.Request
	)

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

			//Merge the map[string]interface{} into one
			_ = mergo.Map(&payloadWithCustomFields, &payloadWithOperations, mergo.WithOverride)
			_ = mergo.Map(&payloadWithCustomFields, &payloadWithTransition, mergo.WithOverride)

			payloadAsReader, _ := transformStructToReader(&payloadWithCustomFields)
			request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
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

			_ = mergo.Map(&payloadWithCustomFields, &payloadWithTransition, mergo.WithOverride)
			payloadAsReader, _ := transformStructToReader(&payloadWithCustomFields)
			request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
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

			_ = mergo.Map(&payloadWithOperations, &payloadWithTransition, mergo.WithOverride)
			payloadAsReader, _ := transformStructToReader(&payloadWithOperations)
			request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
			if err != nil {
				return nil, err
			}
		}

	} else {
		payloadAsReader, _ := transformStructToReader(&payloadWithTransition)
		request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
		if err != nil {
			return
		}
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
