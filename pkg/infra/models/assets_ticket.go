package models

type TicketPageScheme struct {
	Tickets         []*TicketScheme `json:"tickets,omitempty"`
	AllTicketsQuery string          `json:"allTicketsQuery,omitempty"`
}

type TicketScheme struct {
	WorkspaceId string                `json:"workspaceId,omitempty"`
	GlobalId    string                `json:"globalId,omitempty"`
	Key         string                `json:"key,omitempty"`
	Id          string                `json:"id,omitempty"`
	Reporter    string                `json:"reporter,omitempty"`
	Created     string                `json:"created,omitempty"`
	Updated     string                `json:"updated,omitempty"`
	Title       string                `json:"title,omitempty"`
	Status      *TicketStatusScheme   `json:"status,omitempty"`
	Type        *TicketTypeScheme     `json:"type,omitempty"`
	Priority    *TicketPriorityScheme `json:"priority,omitempty"`
}

type TicketStatusScheme struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ColorName   string `json:"colorName,omitempty"`
}

type TicketTypeScheme struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IconUrl     string `json:"iconUrl,omitempty"`
}

type TicketPriorityScheme struct {
	Name    string `json:"name,omitempty"`
	IconUrl string `json:"iconUrl,omitempty"`
}
