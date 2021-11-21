package models

type NotificationSchemeScheme struct {
	Expand                   string                                  `json:"expand,omitempty"`
	ID                       int                                     `json:"id,omitempty"`
	Self                     string                                  `json:"self,omitempty"`
	Name                     string                                  `json:"name,omitempty"`
	Description              string                                  `json:"description,omitempty"`
	NotificationSchemeEvents []*ProjectNotificationSchemeEventScheme `json:"notificationSchemeEvents,omitempty"`
	Scope                    *TeamManagedProjectScopeScheme          `json:"scope,omitempty"`
}

type ProjectNotificationSchemeEventScheme struct {
	Event         *NotificationEventScheme   `json:"event,omitempty"`
	Notifications []*EventNotificationScheme `json:"notifications,omitempty"`
}

type NotificationEventScheme struct {
	ID            int                      `json:"id,omitempty"`
	Name          string                   `json:"name,omitempty"`
	Description   string                   `json:"description,omitempty"`
	TemplateEvent *NotificationEventScheme `json:"templateEvent,omitempty"`
}

type EventNotificationScheme struct {
	Expand           string             `json:"expand,omitempty"`
	ID               int                `json:"id,omitempty"`
	NotificationType string             `json:"notificationType,omitempty"`
	Parameter        string             `json:"parameter,omitempty"`
	EmailAddress     string             `json:"emailAddress,omitempty"`
	Group            *GroupScheme       `json:"group,omitempty"`
	Field            *IssueFieldScheme  `json:"field,omitempty"`
	ProjectRole      *ProjectRoleScheme `json:"projectRole,omitempty"`
	User             *UserScheme        `json:"user,omitempty"`
}
