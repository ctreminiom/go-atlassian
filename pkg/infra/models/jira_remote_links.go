package models

type RemoteLinkIdentify struct {
	ID   int    `json:"id,omitempty"`
	Self string `json:"self,omitempty"`
}

type RemoteLinkScheme struct {
	Application  *RemoteLinkApplicationScheme `json:"application,omitempty"`
	GlobalID     string                       `json:"globalId,omitempty"`
	ID           int                          `json:"id,omitempty"`
	Object       *RemoteLinkObjectScheme      `json:"object,omitempty"`
	Relationship string                       `json:"relationship,omitempty"`
	Self         string                       `json:"self,omitempty"`
}

type RemoteLinkObjectScheme struct {
	Icon    *RemoteLinkObjectLinkScheme   `json:"icon,omitempty"`
	Status  *RemoteLinkObjectStatusScheme `json:"status,omitempty"`
	Summary string                        `json:"summary,omitempty"`
	Title   string                        `json:"title,omitempty"`
	URL     string                        `json:"url,omitempty"`
}

type RemoteLinkObjectStatusScheme struct {
	Icon     *RemoteLinkObjectLinkScheme `json:"icon,omitempty"`
	Resolved bool                        `json:"resolved,omitempty"`
}

type RemoteLinkObjectLinkScheme struct {
	Link     string `json:"link,omitempty"`
	Title    string `json:"title,omitempty"`
	URL16X16 string `json:"url16x16,omitempty"`
}

type RemoteLinkApplicationScheme struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}
