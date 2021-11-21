package models

type SCIMGroupPathScheme struct {
	Schemas    []string                    `json:"schemas,omitempty"`
	Operations []*SCIMGroupOperationScheme `json:"Operations,omitempty"`
}

type SCIMGroupOperationScheme struct {
	Op    string                           `json:"op,omitempty"`
	Path  string                           `json:"path,omitempty"`
	Value []*SCIMGroupOperationValueScheme `json:"value,omitempty"`
}

type SCIMGroupOperationValueScheme struct {
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
}

type ScimGroupPageScheme struct {
	Schemas      []string           `json:"schemas,omitempty"`
	TotalResults int                `json:"totalResults,omitempty"`
	StartIndex   int                `json:"startIndex,omitempty"`
	ItemsPerPage int                `json:"itemsPerPage,omitempty"`
	Resources    []*ScimGroupScheme `json:"Resources,omitempty"`
}

type ScimGroupScheme struct {
	Schemas     []string                 `json:"schemas,omitempty"`
	ID          string                   `json:"id,omitempty"`
	ExternalID  string                   `json:"externalId,omitempty"`
	DisplayName string                   `json:"displayName,omitempty"`
	Members     []*ScimGroupMemberScheme `json:"members,omitempty"`
	Meta        *ScimMetadata            `json:"meta,omitempty"`
}

type ScimGroupMemberScheme struct {
	Type    string `json:"type,omitempty"`
	Value   string `json:"value,omitempty"`
	Display string `json:"display,omitempty"`
	Ref     string `json:"$ref,omitempty"`
}

type ScimMetadata struct {
	ResourceType string `json:"resourceType,omitempty"`
	Location     string `json:"location,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
	Created      string `json:"created,omitempty"`
}
