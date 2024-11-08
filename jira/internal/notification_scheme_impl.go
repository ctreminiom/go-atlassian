package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewNotificationSchemeService creates a new instance of NotificationSchemeService.
func NewNotificationSchemeService(client service.Connector, version string) (*NotificationSchemeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &NotificationSchemeService{
		internalClient: &internalNotificationSchemeImpl{c: client, version: version},
	}, nil
}

// NotificationSchemeService provides methods to manage notification schemes in Jira Service Management.
type NotificationSchemeService struct {
	// internalClient is the connector interface for notification scheme operations.
	internalClient jira.NotificationSchemeConnector
}

// Search returns a paginated list of notification schemes ordered by the display name.
//
// GET /rest/api/{2-3}/notificationscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#get-notification-schemes
func (n *NotificationSchemeService) Search(ctx context.Context, options *model.NotificationSchemeSearchOptions, startAt, maxResults int) (*model.NotificationSchemePageScheme, *model.ResponseScheme, error) {
	return n.internalClient.Search(ctx, options, startAt, maxResults)
}

// Create creates a notification scheme with notifications. You can create up to 1000 notifications per request.
//
// POST /rest/api/{2-3}/notificationscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#create-notification-scheme
func (n *NotificationSchemeService) Create(ctx context.Context, payload *model.NotificationSchemePayloadScheme) (*model.NotificationSchemeCreatedPayload, *model.ResponseScheme, error) {
	return n.internalClient.Create(ctx, payload)
}

// Projects returns a paginated mapping of project that have notification scheme assigned.
//
// You can provide either one or multiple notification scheme IDs or project IDs to filter by.
//
// If you don't provide any, this will return a list of all mappings.
//
// Note that only company-managed (classic) projects are supported.
//
// This is because team-managed projects don't have a concept of a default notification scheme.
//
// The mappings are ordered by projectID.
//
// GET /rest/api/{2-3}/notificationscheme/project
func (n *NotificationSchemeService) Projects(ctx context.Context, schemeIDs, projectIDs []string, startAt, maxResults int) (*model.NotificationSchemeProjectPageScheme, *model.ResponseScheme, error) {
	return n.internalClient.Projects(ctx, schemeIDs, projectIDs, startAt, maxResults)
}

// Get returns a notification scheme, including the list of events and the recipients who will
//
// receive notifications for those events.
//
// GET /rest/api/{2-3}/notificationscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#get-notification-scheme
func (n *NotificationSchemeService) Get(ctx context.Context, schemeID string, expand []string) (*model.NotificationSchemeScheme, *model.ResponseScheme, error) {
	return n.internalClient.Get(ctx, schemeID, expand)
}

// Update updates a notification scheme.
//
// PUT /rest/api/{2-3}/notificationscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#update-notification-scheme
func (n *NotificationSchemeService) Update(ctx context.Context, schemeID string, payload *model.NotificationSchemePayloadScheme) (*model.ResponseScheme, error) {
	return n.internalClient.Update(ctx, schemeID, payload)
}

// Append adds notifications to a notification scheme.
//
// You can add up to 1000 notifications per request.
//
// PUT /rest/api/{2-3}/notificationscheme/{id}/notification
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#append-notifications-to-scheme
func (n *NotificationSchemeService) Append(ctx context.Context, schemeID string, payload *model.NotificationSchemeEventsPayloadScheme) (*model.ResponseScheme, error) {
	return n.internalClient.Append(ctx, schemeID, payload)
}

// Delete deletes a notification scheme.
//
// DELETE /rest/api/{2-3}/notificationscheme/{notificationSchemeId}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#delete-notification-scheme
func (n *NotificationSchemeService) Delete(ctx context.Context, schemeID string) (*model.ResponseScheme, error) {
	return n.internalClient.Delete(ctx, schemeID)
}

// Remove removes a notification from a notification scheme.
//
// DELETE /rest/api/{2-3}/notificationscheme/{notificationSchemeId}/notification/{notificationId}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#remove-notifications-to-scheme
func (n *NotificationSchemeService) Remove(ctx context.Context, schemeID, notificationID string) (*model.ResponseScheme, error) {
	return n.internalClient.Remove(ctx, schemeID, notificationID)
}

type internalNotificationSchemeImpl struct {
	c       service.Connector
	version string
}

func (i *internalNotificationSchemeImpl) Search(ctx context.Context, options *model.NotificationSchemeSearchOptions, startAt, maxResults int) (*model.NotificationSchemePageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		for _, scheme := range options.NotificationSchemeIDs {
			params.Add("id", scheme)
		}

		for _, id := range options.ProjectIDs {
			params.Add("projectId", id)
		}

		if options.OnlyDefault {
			params.Add("onlyDefault", "true")
		}

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	notificationSchemes := new(model.NotificationSchemePageScheme)
	response, err := i.c.Call(request, notificationSchemes)
	if err != nil {
		return nil, response, err
	}

	return notificationSchemes, response, nil
}

func (i *internalNotificationSchemeImpl) Create(ctx context.Context, payload *model.NotificationSchemePayloadScheme) (*model.NotificationSchemeCreatedPayload, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	notificationScheme := new(model.NotificationSchemeCreatedPayload)
	response, err := i.c.Call(request, notificationScheme)
	if err != nil {
		return nil, response, err
	}

	return notificationScheme, response, nil
}

func (i *internalNotificationSchemeImpl) Projects(ctx context.Context, schemeIDs, projectIDs []string, startAt, maxResults int) (*model.NotificationSchemeProjectPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, scheme := range schemeIDs {
		params.Add("notificationSchemeId", scheme)
	}

	for _, id := range projectIDs {
		params.Add("projectId", id)
	}

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme/project?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	notificationProjectSchemes := new(model.NotificationSchemeProjectPageScheme)
	response, err := i.c.Call(request, notificationProjectSchemes)
	if err != nil {
		return nil, response, err
	}

	return notificationProjectSchemes, response, nil
}

func (i *internalNotificationSchemeImpl) Get(ctx context.Context, schemeID string, expand []string) (*model.NotificationSchemeScheme, *model.ResponseScheme, error) {

	if schemeID == "" {
		return nil, nil, model.ErrNoNotificationSchemeID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/notificationscheme/%v", i.version, schemeID))

	if len(expand) != 0 {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	notificationScheme := new(model.NotificationSchemeScheme)
	response, err := i.c.Call(request, notificationScheme)
	if err != nil {
		return nil, response, err
	}

	return notificationScheme, response, nil
}

func (i *internalNotificationSchemeImpl) Update(ctx context.Context, schemeID string, payload *model.NotificationSchemePayloadScheme) (*model.ResponseScheme, error) {

	if schemeID == "" {
		return nil, model.ErrNoNotificationSchemeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme/%v", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalNotificationSchemeImpl) Append(ctx context.Context, schemeID string, payload *model.NotificationSchemeEventsPayloadScheme) (*model.ResponseScheme, error) {

	if schemeID == "" {
		return nil, model.ErrNoNotificationSchemeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme/%v/notification", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalNotificationSchemeImpl) Delete(ctx context.Context, schemeID string) (*model.ResponseScheme, error) {

	if schemeID == "" {
		return nil, model.ErrNoNotificationSchemeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme/%v", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalNotificationSchemeImpl) Remove(ctx context.Context, schemeID, notificationID string) (*model.ResponseScheme, error) {

	if schemeID == "" {
		return nil, model.ErrNoNotificationSchemeID
	}

	if notificationID == "" {
		return nil, model.ErrNoNotificationID
	}

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme/%v/notification/%v", i.version, schemeID, notificationID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
