package models

// ServiceDeskTemporaryFileScheme represents a temporary file in a service desk.
// It contains a slice of temporary attachments.
type ServiceDeskTemporaryFileScheme struct {
	TemporaryAttachments []*TemporaryAttachmentScheme `json:"temporaryAttachments,omitempty"` // The temporary attachments of the file.
}

// TemporaryAttachmentScheme represents a temporary attachment in a service desk.
// It contains the ID and the name of the temporary attachment.
type TemporaryAttachmentScheme struct {
	TemporaryAttachmentID string `json:"temporaryAttachmentId,omitempty"` // The ID of the temporary attachment.
	FileName              string `json:"fileName,omitempty"`              // The name of the temporary attachment.
}

// ServiceDeskPageScheme represents a page of service desks.
// It contains information about the page and a slice of service desks.
type ServiceDeskPageScheme struct {
	Expands    []string                   `json:"_expands,omitempty"`   // The fields to expand in the page.
	Size       int                        `json:"size,omitempty"`       // The size of the page.
	Start      int                        `json:"start,omitempty"`      // The start index of the page.
	Limit      int                        `json:"limit,omitempty"`      // The limit of the page.
	IsLastPage bool                       `json:"isLastPage,omitempty"` // Indicates if this is the last page.
	Links      *ServiceDeskPageLinkScheme `json:"_links,omitempty"`     // The links related to the page.
	Values     []*ServiceDeskScheme       `json:"values,omitempty"`     // The service desks in the page.
}

// ServiceDeskPageLinkScheme represents the links related to a page of service desks.
type ServiceDeskPageLinkScheme struct {
	Base    string `json:"base,omitempty"`    // The base link of the page.
	Context string `json:"context,omitempty"` // The context link of the page.
	Next    string `json:"next,omitempty"`    // The next link of the page.
	Prev    string `json:"prev,omitempty"`    // The previous link of the page.
}

// ServiceDeskScheme represents a service desk.
// It contains information about the service desk and its related project.
type ServiceDeskScheme struct {
	ID          string `json:"id,omitempty"`          // The ID of the service desk.
	ProjectID   string `json:"projectId,omitempty"`   // The ID of the related project.
	ProjectName string `json:"projectName,omitempty"` // The name of the related project.
	ProjectKey  string `json:"projectKey,omitempty"`  // The key of the related project.
	Links       struct {
		Self string `json:"self,omitempty"` // The self link of the service desk.
	} `json:"_links,omitempty"` // The links related to the service desk.
}
