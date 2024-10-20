package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewTypeSchemeService creates a new instance of TypeSchemeService.
func NewTypeSchemeService(client service.Connector, version string) (*TypeSchemeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &TypeSchemeService{
		internalClient: &internalTypeSchemeImpl{c: client, version: version},
	}, nil
}

// TypeSchemeService provides methods to manage issue type schemes in Jira Service Management.
type TypeSchemeService struct {
	// internalClient is the connector interface for issue type scheme operations.
	internalClient jira.TypeSchemeConnector
}

// Gets returns a paginated list of issue type schemes.
//
// GET /rest/api/{2-3}/issuetypescheme
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-all-issue-type-schemes
func (t *TypeSchemeService) Gets(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (*model.IssueTypeSchemePageScheme, *model.ResponseScheme, error) {
	return t.internalClient.Gets(ctx, issueTypeSchemeIDs, startAt, maxResults)
}

// Create creates an issue type scheme.
//
// POST /rest/api/{2-3}/issuetypescheme
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#create-issue-type-scheme
func (t *TypeSchemeService) Create(ctx context.Context, payload *model.IssueTypeSchemePayloadScheme) (*model.NewIssueTypeSchemeScheme, *model.ResponseScheme, error) {
	return t.internalClient.Create(ctx, payload)
}

// Items returns a paginated list of issue type scheme items.
//
// GET /rest/api/{2-3}/issuetypescheme/mapping
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-issue-type-scheme-items
func (t *TypeSchemeService) Items(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (*model.IssueTypeSchemeItemPageScheme, *model.ResponseScheme, error) {
	return t.internalClient.Items(ctx, issueTypeSchemeIDs, startAt, maxResults)
}

// Projects returns a paginated list of issue type schemes and, for each issue type scheme, a list of the projects that use it.
//
// GET /rest/api/{2-3}/issuetypescheme/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-issue-type-schemes-for-projects
func (t *TypeSchemeService) Projects(ctx context.Context, projectIDs []int, startAt, maxResults int) (*model.ProjectIssueTypeSchemePageScheme, *model.ResponseScheme, error) {
	return t.internalClient.Projects(ctx, projectIDs, startAt, maxResults)
}

// Assign assigns an issue type scheme to a project.
//
// PUT /rest/api/{2-3}/issuetypescheme/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#assign-issue-type-scheme-to-project
func (t *TypeSchemeService) Assign(ctx context.Context, issueTypeSchemeID, projectID string) (*model.ResponseScheme, error) {
	return t.internalClient.Assign(ctx, issueTypeSchemeID, projectID)
}

// Update updates an issue type scheme.
//
// PUT /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#update-issue-type-scheme
func (t *TypeSchemeService) Update(ctx context.Context, issueTypeSchemeID int, payload *model.IssueTypeSchemePayloadScheme) (*model.ResponseScheme, error) {
	return t.internalClient.Update(ctx, issueTypeSchemeID, payload)
}

// Delete deletes an issue type scheme.
//
// 1.Only issue type schemes used in classic projects can be deleted.
//
// 2.Any projects assigned to the scheme are reassigned to the default issue type scheme.
//
// DELETE /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#delete-issue-type-scheme
func (t *TypeSchemeService) Delete(ctx context.Context, issueTypeSchemeID int) (*model.ResponseScheme, error) {
	return t.internalClient.Delete(ctx, issueTypeSchemeID)
}

// Append adds issue types to an issue type scheme.
//
// 1.The added issue types are appended to the issue types list.
//
// 2.If any of the issue types exist in the issue type scheme, the operation fails and no issue types are added.
//
// PUT /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeID}/issuetype
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#add-issue-types-to-issue-type-scheme
func (t *TypeSchemeService) Append(ctx context.Context, issueTypeSchemeID int, issueTypeIDs []int) (*model.ResponseScheme, error) {
	return t.internalClient.Append(ctx, issueTypeSchemeID, issueTypeIDs)
}

// Remove removes an issue type from an issue type scheme, this operation cannot remove:
//
// 1.any issue type used by issues.
//
// 2.any issue types from the default issue type scheme.
//
// 3.the last standard issue type from an issue type scheme.
//
// DELETE /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeID}/issuetype/{issueTypeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#remove-issue-type-from-issue-type-scheme
func (t *TypeSchemeService) Remove(ctx context.Context, issueTypeSchemeID, issueTypeID int) (*model.ResponseScheme, error) {
	return t.internalClient.Remove(ctx, issueTypeSchemeID, issueTypeID)
}

type internalTypeSchemeImpl struct {
	c       service.Connector
	version string
}

func (i *internalTypeSchemeImpl) Gets(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (*model.IssueTypeSchemePageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeSchemeIDs {
		params.Add("id", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.IssueTypeSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalTypeSchemeImpl) Create(ctx context.Context, payload *model.IssueTypeSchemePayloadScheme) (*model.NewIssueTypeSchemeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	issueType := new(model.NewIssueTypeSchemeScheme)
	response, err := i.c.Call(request, issueType)
	if err != nil {
		return nil, response, err
	}

	return issueType, response, nil
}

func (i *internalTypeSchemeImpl) Items(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (*model.IssueTypeSchemeItemPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeSchemeIDs {
		params.Add("issueTypeSchemeId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/mapping?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.IssueTypeSchemeItemPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalTypeSchemeImpl) Projects(ctx context.Context, projectIDs []int, startAt, maxResults int) (*model.ProjectIssueTypeSchemePageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range projectIDs {
		params.Add("projectId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/project?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ProjectIssueTypeSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalTypeSchemeImpl) Assign(ctx context.Context, issueTypeSchemeID, projectID string) (*model.ResponseScheme, error) {

	if issueTypeSchemeID == "" {
		return nil, model.ErrNoIssueTypeSchemeID
	}

	if projectID == "" {
		return nil, model.ErrNoProjectID
	}

	payload := map[string]interface{}{
		"issueTypeSchemeId": issueTypeSchemeID,
		"projectId":         projectID,
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeSchemeImpl) Update(ctx context.Context, issueTypeSchemeID int, payload *model.IssueTypeSchemePayloadScheme) (*model.ResponseScheme, error) {

	if issueTypeSchemeID == 0 {
		return nil, model.ErrNoIssueTypeSchemeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v", i.version, issueTypeSchemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeSchemeImpl) Delete(ctx context.Context, issueTypeSchemeID int) (*model.ResponseScheme, error) {

	if issueTypeSchemeID == 0 {
		return nil, model.ErrNoIssueTypeSchemeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v", i.version, issueTypeSchemeID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeSchemeImpl) Append(ctx context.Context, issueTypeSchemeID int, issueTypeIDs []int) (*model.ResponseScheme, error) {

	if len(issueTypeIDs) == 0 {
		return nil, model.ErrNoIssueTypes
	}

	var ids []string
	for _, issueTypeID := range issueTypeIDs {
		ids = append(ids, strconv.Itoa(issueTypeID))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v/issuetype", i.version, issueTypeSchemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"issueTypeIds": ids})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalTypeSchemeImpl) Remove(ctx context.Context, issueTypeSchemeID, issueTypeID int) (*model.ResponseScheme, error) {

	if issueTypeSchemeID == 0 {
		return nil, model.ErrNoIssueTypeSchemeID
	}

	if issueTypeID == 0 {
		return nil, model.ErrNoIssueTypeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v/issuetype/%v", i.version, issueTypeSchemeID, issueTypeID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
