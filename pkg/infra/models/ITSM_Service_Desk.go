package models

type ServiceDeskTemporaryFileScheme struct {
	TemporaryAttachments []*TemporaryAttachmentScheme `json:"temporaryAttachments,omitempty"`
}

type TemporaryAttachmentScheme struct {
	TemporaryAttachmentID string `json:"temporaryAttachmentId,omitempty"`
	FileName              string `json:"fileName,omitempty"`
}

type ServiceDeskPageScheme struct {
	Expands    []string                   `json:"_expands,omitempty"`
	Size       int                        `json:"size,omitempty"`
	Start      int                        `json:"start,omitempty"`
	Limit      int                        `json:"limit,omitempty"`
	IsLastPage bool                       `json:"isLastPage,omitempty"`
	Links      *ServiceDeskPageLinkScheme `json:"_links,omitempty"`
	Values     []*ServiceDeskScheme       `json:"values,omitempty"`
}

type ServiceDeskPageLinkScheme struct {
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type ServiceDeskScheme struct {
	ID          string `json:"id,omitempty"`
	ProjectID   string `json:"projectId,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	ProjectKey  string `json:"projectKey,omitempty"`
	Links       struct {
		Self string `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}
