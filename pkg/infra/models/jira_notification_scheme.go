package models

type NotificationSchemeSearchOptions struct {
	NotificationSchemeIDs []string
	ProjectIDs            []string
	OnlyDefault           bool
	Expand                []string
}

type NotificationSchemePayloadScheme struct {
	Description string                                  `json:"description,omitempty"`
	Name        string                                  `json:"name,omitempty"`
	Events      []*NotificationSchemePayloadEventScheme `json:"notificationSchemeEvents,omitempty"`
}

type NotificationSchemePayloadEventScheme struct {
	Event         *NotificationSchemeEventTypeScheme           `json:"event,omitempty"`
	Notifications []*NotificationSchemeEventNotificationScheme `json:"notifications,omitempty"`
}

type NotificationSchemeEventTypeScheme struct {
	ID string `json:"id,omitempty"`
}

type NotificationSchemeEventNotificationScheme struct {
	NotificationType string `json:"notificationType,omitempty"`
	Parameter        string `json:"parameter,omitempty"`
}

type NotificationSchemePageScheme struct {
	MaxResults int                         `json:"maxResults,omitempty"`
	StartAt    int                         `json:"startAt,omitempty"`
	Total      int                         `json:"total,omitempty"`
	IsLast     bool                        `json:"isLast,omitempty"`
	Values     []*NotificationSchemeScheme `json:"values,omitempty"`
}

type NotificationSchemeCreatedPayload struct {
	Id string `json:"id"`
}

type NotificationSchemeProjectPageScheme struct {
	MaxResults int                                `json:"maxResults,omitempty"`
	StartAt    int                                `json:"startAt,omitempty"`
	Total      int                                `json:"total,omitempty"`
	IsLast     bool                               `json:"isLast,omitempty"`
	Values     []*NotificationSchemeProjectScheme `json:"values,omitempty"`
}

type NotificationSchemeProjectScheme struct {
	NotificationSchemeId string `json:"notificationSchemeId,omitempty"`
	ProjectId            string `json:"projectId,omitempty"`
}

type NotificationSchemeEventsPayloadScheme struct {
	NotificationSchemeEvents []*NotificationSchemePayloadEventScheme `json:"notificationSchemeEvents,omitempty"`
}
