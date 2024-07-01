package models

// NotificationSchemeSearchOptions represents the search options for notification schemes in Jira.
type NotificationSchemeSearchOptions struct {
	NotificationSchemeIDs []string // The IDs of the notification schemes to search for.
	ProjectIDs            []string // The IDs of the projects to search for.
	OnlyDefault           bool     // Indicates if only default notification schemes should be searched for.
	Expand                []string // The fields to expand in the search results.
}

// NotificationSchemePayloadScheme represents the payload for a notification scheme in Jira.
type NotificationSchemePayloadScheme struct {
	Description string                                  `json:"description,omitempty"`              // The description of the notification scheme.
	Name        string                                  `json:"name,omitempty"`                     // The name of the notification scheme.
	Events      []*NotificationSchemePayloadEventScheme `json:"notificationSchemeEvents,omitempty"` // The events of the notification scheme.
}

// NotificationSchemePayloadEventScheme represents an event in the payload for a notification scheme in Jira.
type NotificationSchemePayloadEventScheme struct {
	Event         *NotificationSchemeEventTypeScheme           `json:"event,omitempty"`         // The type of the event.
	Notifications []*NotificationSchemeEventNotificationScheme `json:"notifications,omitempty"` // The notifications for the event.
}

// NotificationSchemeEventTypeScheme represents the type of an event in the payload for a notification scheme in Jira.
type NotificationSchemeEventTypeScheme struct {
	ID string `json:"id,omitempty"` // The ID of the event type.
}

// NotificationSchemeEventNotificationScheme represents a notification for an event in the payload for a notification scheme in Jira.
type NotificationSchemeEventNotificationScheme struct {
	NotificationType string `json:"notificationType,omitempty"` // The type of the notification.
	Parameter        string `json:"parameter,omitempty"`        // The parameter for the notification.
}

// NotificationSchemePageScheme represents a page of notification schemes in Jira.
type NotificationSchemePageScheme struct {
	MaxResults int                         `json:"maxResults,omitempty"` // The maximum number of results per page.
	StartAt    int                         `json:"startAt,omitempty"`    // The starting index of the results.
	Total      int                         `json:"total,omitempty"`      // The total number of results.
	IsLast     bool                        `json:"isLast,omitempty"`     // Indicates if this is the last page of results.
	Values     []*NotificationSchemeScheme `json:"values,omitempty"`     // The notification schemes in the page.
}

// NotificationSchemeCreatedPayload represents the payload for a created notification scheme in Jira.
type NotificationSchemeCreatedPayload struct {
	ID string `json:"id"` // The ID of the created notification scheme.
}

// NotificationSchemeProjectPageScheme represents a page of projects for a notification scheme in Jira.
type NotificationSchemeProjectPageScheme struct {
	MaxResults int                                `json:"maxResults,omitempty"` // The maximum number of results per page.
	StartAt    int                                `json:"startAt,omitempty"`    // The starting index of the results.
	Total      int                                `json:"total,omitempty"`      // The total number of results.
	IsLast     bool                               `json:"isLast,omitempty"`     // Indicates if this is the last page of results.
	Values     []*NotificationSchemeProjectScheme `json:"values,omitempty"`     // The projects for the notification scheme in the page.
}

// NotificationSchemeProjectScheme represents a project for a notification scheme in Jira.
type NotificationSchemeProjectScheme struct {
	NotificationSchemeID string `json:"notificationSchemeId,omitempty"` // The ID of the notification scheme.
	ProjectID            string `json:"projectId,omitempty"`            // The ID of the project.
}

// NotificationSchemeEventsPayloadScheme represents the payload for the events of a notification scheme in Jira.
type NotificationSchemeEventsPayloadScheme struct {
	NotificationSchemeEvents []*NotificationSchemePayloadEventScheme `json:"notificationSchemeEvents,omitempty"` // The events of the notification scheme.
}
