package models

// RemoteLinkIdentify represents the identification of a remote link in Jira.
type RemoteLinkIdentify struct {
	ID   int    `json:"id,omitempty"`   // The ID of the remote link.
	Self string `json:"self,omitempty"` // The URL of the remote link.
}

// RemoteLinkScheme represents a remote link in Jira.
type RemoteLinkScheme struct {
	Application  *RemoteLinkApplicationScheme `json:"application,omitempty"`  // The application associated with the remote link.
	GlobalID     string                       `json:"globalId,omitempty"`     // The global ID of the remote link.
	ID           int                          `json:"id,omitempty"`           // The ID of the remote link.
	Object       *RemoteLinkObjectScheme      `json:"object,omitempty"`       // The object associated with the remote link.
	Relationship string                       `json:"relationship,omitempty"` // The relationship of the remote link.
	Self         string                       `json:"self,omitempty"`         // The URL of the remote link.
}

// RemoteLinkObjectScheme represents an object in a remote link in Jira.
type RemoteLinkObjectScheme struct {
	Icon    *RemoteLinkObjectLinkScheme   `json:"icon,omitempty"`    // The icon of the remote link object.
	Status  *RemoteLinkObjectStatusScheme `json:"status,omitempty"`  // The status of the remote link object.
	Summary string                        `json:"summary,omitempty"` // The summary of the remote link object.
	Title   string                        `json:"title,omitempty"`   // The title of the remote link object.
	URL     string                        `json:"url,omitempty"`     // The URL of the remote link object.
}

// RemoteLinkObjectStatusScheme represents the status of an object in a remote link in Jira.
type RemoteLinkObjectStatusScheme struct {
	Icon     *RemoteLinkObjectLinkScheme `json:"icon,omitempty"`     // The icon of the remote link object status.
	Resolved bool                        `json:"resolved,omitempty"` // Indicates if the remote link object status is resolved.
}

// RemoteLinkObjectLinkScheme represents a link of an object in a remote link in Jira.
type RemoteLinkObjectLinkScheme struct {
	Link     string `json:"link,omitempty"`     // The link of the remote link object.
	Title    string `json:"title,omitempty"`    // The title of the remote link object link.
	URL16X16 string `json:"url16x16,omitempty"` // The 16x16 URL of the remote link object link.
}

// RemoteLinkApplicationScheme represents an application in a remote link in Jira.
type RemoteLinkApplicationScheme struct {
	Name string `json:"name,omitempty"` // The name of the remote link application.
	Type string `json:"type,omitempty"` // The type of the remote link application.
}
