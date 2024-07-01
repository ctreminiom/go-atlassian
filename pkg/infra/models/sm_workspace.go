package models

// WorkSpacePageScheme represents a page of workspaces.
// It contains information about the page and a slice of workspaces.
type WorkSpacePageScheme struct {
	Size       int                       `json:"size,omitempty"`       // The size of the page.
	Start      int                       `json:"start,omitempty"`      // The start index of the page.
	Limit      int                       `json:"limit,omitempty"`      // The limit of the page.
	IsLastPage bool                      `json:"isLastPage,omitempty"` // Indicates if this is the last page.
	Links      *WorkSpaceLinksPageScheme `json:"_links,omitempty"`     // The links related to the page.
	Values     []*WorkSpaceScheme        `json:"values,omitempty"`     // The workspaces in the page.
}

// WorkSpaceLinksPageScheme represents the links related to a page of workspaces.
type WorkSpaceLinksPageScheme struct {
	Self    string `json:"self,omitempty"`    // The self link of the page.
	Base    string `json:"base,omitempty"`    // The base link of the page.
	Context string `json:"context,omitempty"` // The context link of the page.
}

// WorkSpaceScheme represents a workspace.
// It contains the ID of the workspace.
type WorkSpaceScheme struct {
	WorkspaceID string `json:"workspaceId,omitempty"` // The ID of the workspace.
}
