package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/imdario/mergo"
	"net/http"
	"net/url"
	"time"
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
}

type IssueScheme struct {
	ID          string                   `json:"id,omitempty"`
	Key         string                   `json:"key,omitempty"`
	Self        string                   `json:"self,omitempty"`
	Transitions []*IssueTransitionScheme `json:"transitions,omitempty"`
	Changelog   *IssueChangelogScheme    `json:"changelog,omitempty"`
	Fields      *IssueFieldsScheme       `json:"fields,omitempty"`
}

type IssueFieldsScheme struct {
	IssueType                *IssueTypeScheme          `json:"issuetype,omitempty"`
	IssueLinks               []*IssueLinkScheme        `json:"issuelinks,omitempty"`
	Watcher                  *IssueWatcherScheme       `json:"watches,omitempty"`
	Votes                    *IssueVoteScheme          `json:"votes,omitempty"`
	Versions                 []*ProjectVersionScheme   `json:"versions,omitempty"`
	Project                  *ProjectScheme            `json:"project,omitempty"`
	FixVersions              []*ProjectVersionScheme   `json:"fixVersions,omitempty"`
	Priority                 *PriorityScheme           `json:"priority,omitempty"`
	Components               *[]ProjectComponentScheme `json:"components,omitempty"`
	Creator                  *UserScheme               `json:"creator,omitempty"`
	Reporter                 *UserScheme               `json:"reporter,omitempty"`
	Resolution               *IssueResolutionScheme    `json:"resolution,omitempty"`
	Resolutiondate           string                    `json:"resolutiondate,omitempty"`
	Workratio                int                       `json:"workratio,omitempty"`
	StatuscategoryChangeDate string                    `json:"statuscategorychangedate,omitempty"`
	LastViewed               string                    `json:"lastViewed,omitempty"`
	Summary                  string                    `json:"summary,omitempty"`
	Created                  string                    `json:"created,omitempty"`
	Updated                  string                    `json:"updated,omitempty"`
	Labels                   []string                  `json:"labels,omitempty"`
	Status                   *StatusScheme             `json:"status,omitempty"`
	Description              *CommentNodeScheme        `json:"description,omitempty"`
	Comments                 []*IssueCommentPageScheme `json:"comments,omitempty"`
	Subtasks                 *[]IssueScheme            `json:"subtasks,omitempty"`
}

func (i *IssueScheme) MergeCustomFields(fields *CustomFields) (result map[string]interface{}, err error) {

	if fields == nil {
		return nil, fmt.Errorf("error, please provide a value *CustomFields pointer")
	}

	if len(fields.Fields) == 0 {
		return nil, fmt.Errorf("error!, the Fields tag does not contains custom fields")
	}

	//Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, _ := json.Marshal(i)

	issueSchemeAsMap := make(map[string]interface{})
	_ = json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap)

	//For each customField created, merge it into the eAsMap
	for _, customField := range fields.Fields {
		_ = mergo.Merge(&issueSchemeAsMap, customField, mergo.WithOverride)
	}

	return issueSchemeAsMap, nil
}

func (i *IssueScheme) MergeOperations(operations *UpdateOperations) (result map[string]interface{}, err error) {

	if operations == nil {
		return nil, fmt.Errorf("error, please provide a value *UpdateOperations pointer")
	}

	if len(operations.Fields) == 0 {
		return nil, fmt.Errorf("error!, the Fields tag does not contains custom fields")
	}

	//Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, _ := json.Marshal(i)

	issueSchemeAsMap := make(map[string]interface{})
	_ = json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap)

	//For each customField created, merge it into the eAsMap
	for _, customField := range operations.Fields {
		_ = mergo.Merge(&issueSchemeAsMap, customField, mergo.WithOverride)
	}

	return issueSchemeAsMap, nil
}

func (i *IssueScheme) ToMap() (result map[string]interface{}, err error) {

	//Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, _ := json.Marshal(i)

	issueSchemeAsMap := make(map[string]interface{})
	_ = json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap)

	return issueSchemeAsMap, err
}

type CustomFields struct{ Fields []map[string]interface{} }

func (c *CustomFields) Groups(customFieldID string, groups []string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(groups) == 0 {
		return fmt.Errorf("error, please provide a valid groups value")
	}

	var groupsNode []map[string]interface{}
	for _, group := range groups {

		var groupNode = map[string]interface{}{}
		groupNode["name"] = group

		groupsNode = append(groupsNode, groupNode)
	}

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = groupsNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) Group(customFieldID, group string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(group) == 0 {
		return fmt.Errorf("error, please provide a valid group value")
	}

	var groupNode = map[string]interface{}{}
	groupNode["name"] = group

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = groupNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) URL(customFieldID, URL string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(URL) == 0 {
		return fmt.Errorf("error, please provide a valid URL value")
	}

	var urlNode = map[string]interface{}{}
	urlNode[customFieldID] = URL

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = urlNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) Text(customFieldID, textValue string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(textValue) == 0 {
		return fmt.Errorf("error, please provide a valid textValue value")
	}

	var urlNode = map[string]interface{}{}
	urlNode[customFieldID] = textValue

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = urlNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) DateTime(customFieldID string, dateValue time.Time) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if dateValue.IsZero() {
		return fmt.Errorf("error, please provide a valid dateValue value")
	}

	var dateNode = map[string]interface{}{}
	dateNode[customFieldID] = dateValue.Format(time.RFC3339)

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = dateNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) Date(customFieldID string, dateTimeValue time.Time) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if dateTimeValue.IsZero() {
		return fmt.Errorf("error, please provide a valid dateValue value")
	}

	var dateTimeNode = map[string]interface{}{}
	dateTimeNode[customFieldID] = dateTimeValue.Format("2006-01-02")

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = dateTimeNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) MultiSelect(customFieldID string, options []string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(options) == 0 {
		return fmt.Errorf("error, please provide a valid options value")
	}

	var groupsNode []map[string]interface{}
	for _, group := range options {

		var groupNode = map[string]interface{}{}
		groupNode["value"] = group

		groupsNode = append(groupsNode, groupNode)
	}

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = groupsNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) Select(customFieldID string, option string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(option) == 0 {
		return fmt.Errorf("error, please provide a valid option value")
	}

	var selectNode = map[string]interface{}{}
	selectNode["value"] = option

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = selectNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) RadioButton(customFieldID, button string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(button) == 0 {
		return fmt.Errorf("error, please provide a button option value")
	}

	var selectNode = map[string]interface{}{}
	selectNode["value"] = button

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = selectNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) User(customFieldID string, accountID string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(accountID) == 0 {
		return fmt.Errorf("error, please provide a accountID option value")
	}

	var userNode = map[string]interface{}{}
	userNode["accountId"] = accountID

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = userNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) Users(customFieldID string, accountIDs []string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(accountIDs) == 0 {
		return fmt.Errorf("error, please provide a accountIDs value")
	}

	var accountsNode []map[string]interface{}
	for _, accountID := range accountIDs {

		var groupNode = map[string]interface{}{}
		groupNode["accountId"] = accountID

		accountsNode = append(accountsNode, groupNode)
	}

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = accountsNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) Number(customFieldID string, numberValue float64) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	var urlNode = map[string]interface{}{}
	urlNode[customFieldID] = numberValue

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = urlNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) CheckBox(customFieldID string, options []string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(options) == 0 {
		return fmt.Errorf("error, please provide a valid options value")
	}

	var groupsNode []map[string]interface{}
	for _, group := range options {

		var groupNode = map[string]interface{}{}
		groupNode["value"] = group

		groupsNode = append(groupsNode, groupNode)
	}

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = groupsNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

func (c *CustomFields) Cascading(customFieldID, parent, child string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(parent) == 0 {
		return fmt.Errorf("error, please provide a parent value")
	}

	if len(child) == 0 {
		return fmt.Errorf("error, please provide a child value")
	}

	var childNode = map[string]interface{}{}
	childNode["value"] = child

	var parentNode = map[string]interface{}{}
	parentNode["value"] = parent
	parentNode["child"] = childNode

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = parentNode

	var fieldsNode = map[string]interface{}{}
	fieldsNode["fields"] = fieldNode

	c.Fields = append(c.Fields, fieldsNode)
	return
}

// Creates an issue or, where the option to create subtasks is enabled in Jira, a subtask.
// https://docs.go-atlassian.io/jira-software-cloud/issues#create-issue
func (i *IssueService) Create(ctx context.Context, payload *IssueScheme, customFields *CustomFields) (result *IssueResponseScheme, response *Response, err error) {

	var (
		endpoint = "rest/api/3/issue"
		request  *http.Request
	)

	if customFields != nil {

		payloadWithCustomFields, err := payload.MergeCustomFields(customFields)
		if err != nil {
			return nil, nil, err
		}

		request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payloadWithCustomFields)
		if err != nil {
			return nil, nil, err
		}
	} else {

		request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payload)
		if err != nil {
			return nil, nil, err
		}
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueResponseScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type IssueResponseScheme struct {
	ID         string `json:"id"`
	Key        string `json:"key"`
	Self       string `json:"self"`
	Transition struct {
		Status          int `json:"status"`
		ErrorCollection struct {
			ErrorMessages []string `json:"errorMessages"`
			Errors        struct {
			} `json:"errors"`
			Status int `json:"status"`
		} `json:"errorCollection"`
	} `json:"transition"`
}

type IssueBulkScheme struct {
	Payload      *IssueScheme
	CustomFields *CustomFields
}

// Creates issues and, where the option to create subtasks is enabled in Jira, subtasks.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#bulk-create-issue
func (i *IssueService) Creates(ctx context.Context, payload []*IssueBulkScheme) (result *IssueBulkResponseScheme, response *Response, err error) {

	if len(payload) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid []*IssueBulkScheme slice of pointers")
	}

	var issuePayloadsNodeAsList []map[string]interface{}
	for pos, newIssue := range payload {

		if newIssue.Payload == nil {
			return nil, nil, fmt.Errorf("error, the issueScheme payload #%v is nil, please provide a valid *IssueScheme pointer", pos)
		}

		//Convert the issueScheme struct to map
		newIssueAsMap, err := newIssue.Payload.MergeCustomFields(newIssue.CustomFields)
		if err != nil {
			return nil, nil, err
		}

		issuePayloadsNodeAsList = append(issuePayloadsNodeAsList, newIssueAsMap)
	}

	var issueUpdatesNode = map[string]interface{}{}
	issueUpdatesNode["issueUpdates"] = issuePayloadsNodeAsList

	var endpoint = "rest/api/3/issue/bulk"

	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, issueUpdatesNode)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueBulkResponseScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type BulkIssueScheme struct {
	Issues []IssueScheme `json:"issues"`
}

type IssueBulkResponseScheme struct {
	Issues []struct {
		ID         string `json:"id"`
		Key        string `json:"key"`
		Self       string `json:"self"`
		Transition struct {
			Status          int `json:"status"`
			ErrorCollection struct {
			} `json:"errorCollection"`
		} `json:"transition"`
	} `json:"issues"`
	Errors []struct {
		Status        int `json:"status"`
		ElementErrors struct {
			ErrorMessages []string `json:"errorMessages"`
			Errors        struct {
			} `json:"errors"`
			Status int `json:"status"`
		} `json:"elementErrors"`
		FailedElementNumber int `json:"failedElementNumber"`
	} `json:"errors"`
}

// Returns the details for an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#get-issue
func (i *IssueService) Get(ctx context.Context, issueKeyOrID string, fields []string, expands []string) (result *IssueScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	params := url.Values{}

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var fieldsNames string
	for index, value := range fields {

		if index == 0 {
			fieldsNames = value
			continue
		}

		fieldsNames += "," + value
	}

	if len(fieldsNames) != 0 {
		params.Add("fields", fieldsNames)
	}

	var endpoint string
	if params.Encode() != "" {
		endpoint = fmt.Sprintf("rest/api/3/issue/%v?%v", issueKeyOrID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/issue/%v", issueKeyOrID)
	}

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type UpdateOperations struct{ Fields []map[string]interface{} }

func (u *UpdateOperations) AddArrayOperation(customFieldID string, mapping map[string]string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	var operations []map[string]interface{}
	for value, operation := range mapping {

		var operationNode = map[string]interface{}{}
		operationNode[operation] = value

		operations = append(operations, operationNode)
	}

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = operations

	var updateNode = map[string]interface{}{}
	updateNode["update"] = fieldNode

	u.Fields = append(u.Fields, updateNode)
	return
}

func (u *UpdateOperations) AddStringOperation(customFieldID, operation, value string) (err error) {

	if len(customFieldID) == 0 {
		return fmt.Errorf("error, please provide a valid customFieldID value")
	}

	if len(operation) == 0 {
		return fmt.Errorf("error, please provide a valid operation value")
	}

	if len(value) == 0 {
		return fmt.Errorf("error, please provide a valid value value")
	}

	var operations []map[string]interface{}

	var operationNode = map[string]interface{}{}
	operationNode[operation] = value

	operations = append(operations, operationNode)

	var fieldNode = map[string]interface{}{}
	fieldNode[customFieldID] = operations

	var updateNode = map[string]interface{}{}
	updateNode["update"] = fieldNode

	u.Fields = append(u.Fields, updateNode)

	return
}

// Edits an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#edit-issue
func (i *IssueService) Update(ctx context.Context, issueKeyOrID string, notify bool, payload *IssueScheme, customFields *CustomFields, operations *UpdateOperations) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	if payload == nil {
		return nil, fmt.Errorf("error, please provide a valid *IssueScheme pointer")
	}

	params := url.Values{}
	if !notify {
		params.Add("notifyUsers", "false")
	}

	var endpoint string
	if params.Encode() != "" {
		endpoint = fmt.Sprintf("rest/api/3/issue/%v?%v", issueKeyOrID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/issue/%v", issueKeyOrID)
	}

	var request *http.Request

	// Executed when customfields or operation are not provided
	if customFields == nil && operations == nil {

		request, err = i.client.newRequest(ctx, http.MethodPut, endpoint, payload)
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

		request, err = i.client.newRequest(ctx, http.MethodPut, endpoint, payloadWithCustomFields)
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

		request, err = i.client.newRequest(ctx, http.MethodPut, endpoint, payloadWithCustomFields)
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

		request, err = i.client.newRequest(ctx, http.MethodPut, endpoint, payloadWithOperations)
		if err != nil {
			return nil, err
		}
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#delete-issue
func (i *IssueService) Delete(ctx context.Context, issueKeyOrID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v", issueKeyOrID)
	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Assigns an issue to a user.
// Use this operation when the calling user does not have the Edit Issues permission but has the
// Assign issue permission for the project that the issue is in.
// If accountId is set to:
//  1. "-1", the issue is assigned to the default assignee for the project.
//  2. null, the issue is set to unassigned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#assign-issue
func (i *IssueService) Assign(ctx context.Context, issueKeyOrID, accountID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	payload := struct {
		AccountID string `json:"accountId"`
	}{AccountID: accountID}

	var endpoint = fmt.Sprintf("/rest/api/3/issue/%v/assignee", issueKeyOrID)

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

type IssueNotifyOptionsScheme struct {
	HTMLBody string                     `json:"htmlBody,omitempty"`
	Subject  string                     `json:"subject,omitempty"`
	TextBody string                     `json:"textBody,omitempty"`
	To       *IssueNotifyToScheme       `json:"to,omitempty"`
	Restrict *IssueNotifyRestrictScheme `json:"restrict,omitempty"`
}

type IssueNotifyRestrictScheme struct {
	Groups      []*IssueNotifyGroupScheme      `json:"groups,omitempty"`
	Permissions []*IssueNotifyPermissionScheme `json:"permissions,omitempty"`
}

type IssueNotifyToScheme struct {
	Reporter bool                      `json:"reporter,omitempty"`
	Assignee bool                      `json:"assignee,omitempty"`
	Watchers bool                      `json:"watchers,omitempty"`
	Voters   bool                      `json:"voters,omitempty"`
	Users    []*IssueNotifyUserScheme  `json:"users,omitempty"`
	Groups   []*IssueNotifyGroupScheme `json:"groups,omitempty"`
}

type IssueNotifyPermissionScheme struct {
	ID  string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
}

type IssueNotifyUserScheme struct {
	AccountID string `json:"accountId,omitempty"`
}

type IssueNotifyGroupScheme struct {
	Name string `json:"name,omitempty"`
}

// Creates an email notification for an issue and adds it to the mail queue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#send-notification-for-issue
func (i *IssueService) Notify(ctx context.Context, issueKeyOrID string, options *IssueNotifyOptionsScheme) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid issueKeyOrID string value")
	}

	if options == nil {
		return nil, fmt.Errorf("error, please provide a valid *IssueNotifyOptionsScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/notify", issueKeyOrID)
	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, &options)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Returns either all transitions or a transition that can be performed by the user on an issue, based on the issue's status.
// Note, if a request is made for a transition that does not exist or cannot be performed on the issue,
// given its status, the response will return any empty transitions list.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#get-transitions
func (i *IssueService) Transitions(ctx context.Context, issueKeyOrID string) (result *IssueTransitionsScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID string value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/transitions", issueKeyOrID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueTransitionsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type IssueMoveOptions struct {
	Fields       *IssueScheme
	CustomFields *CustomFields
	Operations   *UpdateOperations
}

// Performs an issue transition and
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues#transition-issue
func (i *IssueService) Move(ctx context.Context, issueKeyOrID, transitionID string, options *IssueMoveOptions) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid issueKeyOrID string value")
	}

	if len(transitionID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid transitionID string value")
	}

	payloadWithTransition := make(map[string]interface{})
	payloadWithTransition["transition"] = map[string]interface{}{"id": transitionID}

	var (
		endpoint = fmt.Sprintf("rest/api/3/issue/%v/transitions", issueKeyOrID)
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

			request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payloadWithCustomFields)
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
			request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payloadWithCustomFields)
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
			request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, payloadWithOperations)
			if err != nil {
				return nil, err
			}
		}

	} else {
		request, err = i.client.newRequest(ctx, http.MethodPost, endpoint, &payloadWithTransition)
		if err != nil {
			return
		}
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}
