package models

type WorkSpacePageScheme struct {
	Size       int                       `json:"size,omitempty"`
	Start      int                       `json:"start,omitempty"`
	Limit      int                       `json:"limit,omitempty"`
	IsLastPage bool                      `json:"isLastPage,omitempty"`
	Links      *WorkSpaceLinksPageScheme `json:"_links,omitempty"`
	Values     []*WorkSpaceScheme        `json:"values,omitempty"`
}

type WorkSpaceLinksPageScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
}

type WorkSpaceScheme struct {
	WorkspaceId string `json:"workspaceId,omitempty"`
}
