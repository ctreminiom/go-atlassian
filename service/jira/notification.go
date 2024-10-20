package jira

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// NotificationSchemeConnector represents notification schemes,
// lists of events and the recipients who will receive notifications for those events.
// Use it to get details of a notification scheme and a list of notification schemes.
type NotificationSchemeConnector interface {

	// Search returns a paginated list of notification schemes ordered by the display name.
	//
	// GET /rest/api/{2-3}/notificationscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#get-notification-schemes
	Search(ctx context.Context, options *models.NotificationSchemeSearchOptions, startAt, maxResults int) (*models.NotificationSchemePageScheme, *models.ResponseScheme, error)

	// Create creates a notification scheme with notifications. You can create up to 1000 notifications per request.
	//
	// POST /rest/api/{2-3}/notificationscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#create-notification-scheme
	Create(ctx context.Context, payload *models.NotificationSchemePayloadScheme) (*models.NotificationSchemeCreatedPayload, *models.ResponseScheme, error)

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
	Projects(ctx context.Context, schemeIDs, projectIDs []string, startAt, maxResults int) (*models.NotificationSchemeProjectPageScheme, *models.ResponseScheme, error)

	// Get returns a notification scheme, including the list of events and the recipients who will
	//
	// receive notifications for those events.
	//
	// GET /rest/api/{2-3}/notificationscheme/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#get-notification-scheme
	Get(ctx context.Context, schemeID string, expand []string) (*models.NotificationSchemeScheme, *models.ResponseScheme, error)

	// Update updates a notification scheme.
	//
	// PUT /rest/api/{2-3}/notificationscheme/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#update-notification-scheme
	Update(ctx context.Context, schemeID string, payload *models.NotificationSchemePayloadScheme) (*models.ResponseScheme, error)

	// Append adds notifications to a notification scheme.
	//
	// You can add up to 1000 notifications per request.
	//
	// PUT /rest/api/{2-3}/notificationscheme/{id}/notification
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#append-notifications-to-scheme
	Append(ctx context.Context, schemeID string, payload *models.NotificationSchemeEventsPayloadScheme) (*models.ResponseScheme, error)

	// Delete deletes a notification scheme.
	//
	// DELETE /rest/api/{2-3}/notificationscheme/{schemeID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#delete-notification-scheme
	Delete(ctx context.Context, schemeID string) (*models.ResponseScheme, error)

	// Remove removes a notification from a notification scheme.
	//
	// DELETE /rest/api/{2-3}/notificationscheme/{schemeID}/notification/{notificationID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/notification-schemes#remove-notifications-to-scheme
	Remove(ctx context.Context, schemeID, notificationID string) (*models.ResponseScheme, error)
}
