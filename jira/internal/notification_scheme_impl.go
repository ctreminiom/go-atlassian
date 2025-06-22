package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewNotificationSchemeService creates a new instance of NotificationSchemeService.
func NewNotificationSchemeService(client service.Connector, version string) (*NotificationSchemeService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
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
	ctx, span := tracer().Start(ctx, "(*NotificationSchemeService).Search", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "search_notification_schemes"),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := n.internalClient.Search(ctx, options, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Create creates a notification scheme with notifications. You can create up to 1000 notifications per request.
//
// POST /rest/api/{2-3}/notificationscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#create-notification-scheme
func (n *NotificationSchemeService) Create(ctx context.Context, payload *model.NotificationSchemePayloadScheme) (*model.NotificationSchemeCreatedPayload, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*NotificationSchemeService).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create_notification_scheme"),
	)

	result, response, err := n.internalClient.Create(ctx, payload)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
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
	ctx, span := tracer().Start(ctx, "(*NotificationSchemeService).Projects", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_notification_scheme_projects"),
		attribute.StringSlice("jira.scheme_ids", schemeIDs),
		attribute.StringSlice("jira.project_ids", projectIDs),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := n.internalClient.Projects(ctx, schemeIDs, projectIDs, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Get returns a notification scheme, including the list of events and the recipients who will
//
// receive notifications for those events.
//
// GET /rest/api/{2-3}/notificationscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#get-notification-scheme
func (n *NotificationSchemeService) Get(ctx context.Context, schemeID string, expand []string) (*model.NotificationSchemeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*NotificationSchemeService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_notification_scheme"),
		attribute.String("jira.scheme.id", schemeID),
		attribute.StringSlice("jira.expand", expand),
	)

	result, response, err := n.internalClient.Get(ctx, schemeID, expand)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Update updates a notification scheme.
//
// PUT /rest/api/{2-3}/notificationscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#update-notification-scheme
func (n *NotificationSchemeService) Update(ctx context.Context, schemeID string, payload *model.NotificationSchemePayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*NotificationSchemeService).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update_notification_scheme"),
		attribute.String("jira.scheme.id", schemeID),
	)

	response, err := n.internalClient.Update(ctx, schemeID, payload)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// Append adds notifications to a notification scheme.
//
// You can add up to 1000 notifications per request.
//
// PUT /rest/api/{2-3}/notificationscheme/{id}/notification
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#append-notifications-to-scheme
func (n *NotificationSchemeService) Append(ctx context.Context, schemeID string, payload *model.NotificationSchemeEventsPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*NotificationSchemeService).Append", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "append_notification_scheme"),
		attribute.String("jira.scheme.id", schemeID),
	)

	response, err := n.internalClient.Append(ctx, schemeID, payload)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// Delete deletes a notification scheme.
//
// DELETE /rest/api/{2-3}/notificationscheme/{notificationSchemeId}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#delete-notification-scheme
func (n *NotificationSchemeService) Delete(ctx context.Context, schemeID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*NotificationSchemeService).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_notification_scheme"),
		attribute.String("jira.scheme.id", schemeID),
	)

	response, err := n.internalClient.Delete(ctx, schemeID)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// Remove removes a notification from a notification scheme.
//
// DELETE /rest/api/{2-3}/notificationscheme/{notificationSchemeId}/notification/{notificationId}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#remove-notifications-to-scheme
func (n *NotificationSchemeService) Remove(ctx context.Context, schemeID, notificationID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*NotificationSchemeService).Remove", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "remove_notification_scheme"),
		attribute.String("jira.scheme.id", schemeID),
		attribute.String("jira.notification.id", notificationID),
	)

	response, err := n.internalClient.Remove(ctx, schemeID, notificationID)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

type internalNotificationSchemeImpl struct {
	c       service.Connector
	version string
}

func (i *internalNotificationSchemeImpl) Search(ctx context.Context, options *model.NotificationSchemeSearchOptions, startAt, maxResults int) (*model.NotificationSchemePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalNotificationSchemeImpl).Search", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "search_notification_schemes"),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

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
		recordError(span, err)
		return nil, nil, err
	}

	notificationSchemes := new(model.NotificationSchemePageScheme)
	response, err := i.c.Call(request, notificationSchemes)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return notificationSchemes, response, nil
}

func (i *internalNotificationSchemeImpl) Create(ctx context.Context, payload *model.NotificationSchemePayloadScheme) (*model.NotificationSchemeCreatedPayload, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalNotificationSchemeImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create_notification_scheme"),
	)

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	notificationScheme := new(model.NotificationSchemeCreatedPayload)
	response, err := i.c.Call(request, notificationScheme)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return notificationScheme, response, nil
}

func (i *internalNotificationSchemeImpl) Projects(ctx context.Context, schemeIDs, projectIDs []string, startAt, maxResults int) (*model.NotificationSchemeProjectPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalNotificationSchemeImpl).Projects", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_notification_scheme_projects"),
		attribute.StringSlice("jira.scheme_ids", schemeIDs),
		attribute.StringSlice("jira.project_ids", projectIDs),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

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
		recordError(span, err)
		return nil, nil, err
	}

	notificationProjectSchemes := new(model.NotificationSchemeProjectPageScheme)
	response, err := i.c.Call(request, notificationProjectSchemes)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return notificationProjectSchemes, response, nil
}

func (i *internalNotificationSchemeImpl) Get(ctx context.Context, schemeID string, expand []string) (*model.NotificationSchemeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalNotificationSchemeImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_notification_scheme"),
		attribute.String("jira.scheme.id", schemeID),
		attribute.StringSlice("jira.expand", expand),
	)

	if schemeID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoNotificationSchemeID)
		recordError(span, err)
		return nil, nil, err
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
		recordError(span, err)
		return nil, nil, err
	}

	notificationScheme := new(model.NotificationSchemeScheme)
	response, err := i.c.Call(request, notificationScheme)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return notificationScheme, response, nil
}

func (i *internalNotificationSchemeImpl) Update(ctx context.Context, schemeID string, payload *model.NotificationSchemePayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalNotificationSchemeImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update_notification_scheme"),
		attribute.String("jira.scheme.id", schemeID),
	)

	if schemeID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoNotificationSchemeID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme/%v", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func (i *internalNotificationSchemeImpl) Append(ctx context.Context, schemeID string, payload *model.NotificationSchemeEventsPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalNotificationSchemeImpl).Append", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "append_notification_scheme"),
		attribute.String("jira.scheme.id", schemeID),
	)

	if schemeID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoNotificationSchemeID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme/%v/notification", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func (i *internalNotificationSchemeImpl) Delete(ctx context.Context, schemeID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalNotificationSchemeImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_notification_scheme"),
		attribute.String("jira.scheme.id", schemeID),
	)

	if schemeID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoNotificationSchemeID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme/%v", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func (i *internalNotificationSchemeImpl) Remove(ctx context.Context, schemeID, notificationID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalNotificationSchemeImpl).Remove", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "remove_notification_scheme"),
		attribute.String("jira.scheme.id", schemeID),
		attribute.String("jira.notification.id", notificationID),
	)

	if schemeID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoNotificationSchemeID)
		recordError(span, err)
		return nil, err
	}

	if notificationID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoNotificationID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/notificationscheme/%v/notification/%v", i.version, schemeID, notificationID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}
