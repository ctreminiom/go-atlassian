package models

// NotificationSchemeScheme represents a notification scheme in Jira.
type NotificationSchemeScheme struct {
	Expand                   string                                  `json:"expand,omitempty"`                   // The expand field for the notification scheme.
	ID                       int                                     `json:"id,omitempty"`                       // The ID of the notification scheme.
	Self                     string                                  `json:"self,omitempty"`                     // The URL of the notification scheme.
	Name                     string                                  `json:"name,omitempty"`                     // The name of the notification scheme.
	Description              string                                  `json:"description,omitempty"`              // The description of the notification scheme.
	NotificationSchemeEvents []*ProjectNotificationSchemeEventScheme `json:"notificationSchemeEvents,omitempty"` // The events of the notification scheme.
	Scope                    *TeamManagedProjectScopeScheme          `json:"scope,omitempty"`                    // The scope of the notification scheme.
	Projects                 []int                                   `json:"projects,omitempty"`                 // The projects associated with the notification scheme.
}

// ProjectNotificationSchemeEventScheme represents an event in a project notification scheme in Jira.
type ProjectNotificationSchemeEventScheme struct {
	Event         *NotificationEventScheme   `json:"event,omitempty"`         // The event of the project notification scheme.
	Notifications []*EventNotificationScheme `json:"notifications,omitempty"` // The notifications of the project notification scheme.
}

// NotificationEventScheme represents a notification event in Jira.
type NotificationEventScheme struct {
	ID            int                      `json:"id,omitempty"`            // The ID of the notification event.
	Name          string                   `json:"name,omitempty"`          // The name of the notification event.
	Description   string                   `json:"description,omitempty"`   // The description of the notification event.
	TemplateEvent *NotificationEventScheme `json:"templateEvent,omitempty"` // The template event of the notification event.
}

// EventNotificationScheme represents an event notification in Jira.
type EventNotificationScheme struct {
	Expand           string             `json:"expand,omitempty"`           // The expand field for the event notification.
	ID               int                `json:"id,omitempty"`               // The ID of the event notification.
	NotificationType string             `json:"notificationType,omitempty"` // The type of the event notification.
	Parameter        string             `json:"parameter,omitempty"`        // The parameter of the event notification.
	EmailAddress     string             `json:"emailAddress,omitempty"`     // The email address associated with the event notification.
	Group            *GroupScheme       `json:"group,omitempty"`            // The group associated with the event notification.
	Field            *IssueFieldScheme  `json:"field,omitempty"`            // The field associated with the event notification.
	ProjectRole      *ProjectRoleScheme `json:"projectRole,omitempty"`      // The project role associated with the event notification.
	User             *UserScheme        `json:"user,omitempty"`             // The user associated with the event notification.
}
