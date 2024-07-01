package models

// TicketPageScheme represents a paginated list of tickets.
// Tickets is a slice of the tickets on the page.
// AllTicketsQuery is a query that fetches all tickets.
type TicketPageScheme struct {
	Tickets         []*TicketScheme `json:"tickets,omitempty"`         // The tickets on the page.
	AllTicketsQuery string          `json:"allTicketsQuery,omitempty"` // A query that fetches all tickets.
}

// TicketScheme represents a ticket.
type TicketScheme struct {
	WorkspaceID string                `json:"workspaceId,omitempty"` // The ID of the workspace.
	GlobalID    string                `json:"globalId,omitempty"`    // The global ID of the ticket.
	Key         string                `json:"key,omitempty"`         // The key of the ticket.
	ID          string                `json:"id,omitempty"`          // The ID of the ticket.
	Reporter    string                `json:"reporter,omitempty"`    // The reporter of the ticket.
	Created     string                `json:"created,omitempty"`     // The creation time of the ticket.
	Updated     string                `json:"updated,omitempty"`     // The update time of the ticket.
	Title       string                `json:"title,omitempty"`       // The title of the ticket.
	Status      *TicketStatusScheme   `json:"status,omitempty"`      // The status of the ticket.
	Type        *TicketTypeScheme     `json:"type,omitempty"`        // The type of the ticket.
	Priority    *TicketPriorityScheme `json:"priority,omitempty"`    // The priority of the ticket.
}

// TicketStatusScheme represents the status of a ticket.
// Name is the name of the status.
// Description is the description of the status.
// ColorName is the name of the color associated with the status.
type TicketStatusScheme struct {
	Name        string `json:"name,omitempty"`        // The name of the status.
	Description string `json:"description,omitempty"` // The description of the status.
	ColorName   string `json:"colorName,omitempty"`   // The name of the color associated with the status.
}

// TicketTypeScheme represents the type of a ticket.
// Name is the name of the type.
// Description is the description of the type.
// IconUrl is the URL of the icon associated with the type.
type TicketTypeScheme struct {
	Name        string `json:"name,omitempty"`        // The name of the type.
	Description string `json:"description,omitempty"` // The description of the type.
	IconUrl     string `json:"iconUrl,omitempty"`     // The URL of the icon associated with the type.
}

// TicketPriorityScheme represents the priority of a ticket.
// Name is the name of the priority.
// IconUrl is the URL of the icon associated with the priority.
type TicketPriorityScheme struct {
	Name    string `json:"name,omitempty"`    // The name of the priority.
	IconUrl string `json:"iconUrl,omitempty"` // The URL of the icon associated with the priority.
}
