// Package models provides the data structures used in the admin package.
package models

// SCIMGroupPathScheme represents the path scheme for a SCIM group.
type SCIMGroupPathScheme struct {
	Schemas    []string                    `json:"schemas,omitempty"`    // The schemas for the SCIM group.
	Operations []*SCIMGroupOperationScheme `json:"Operations,omitempty"` // The operations for the SCIM group.
}

// SCIMGroupOperationScheme represents the operation scheme for a SCIM group.
type SCIMGroupOperationScheme struct {
	Op    string                           `json:"op,omitempty"`    // The operation type.
	Path  string                           `json:"path,omitempty"`  // The path for the operation.
	Value []*SCIMGroupOperationValueScheme `json:"value,omitempty"` // The values for the operation.
}

// SCIMGroupOperationValueScheme represents the value scheme for a SCIM group operation.
type SCIMGroupOperationValueScheme struct {
	Value   string `json:"value,omitempty"`   // The value for the operation.
	Display string `json:"display,omitempty"` // The display for the operation.
}

// ScimGroupPageScheme represents a page of SCIM groups.
type ScimGroupPageScheme struct {
	Schemas      []string           `json:"schemas,omitempty"`      // The schemas for the SCIM groups.
	TotalResults int                `json:"totalResults,omitempty"` // The total number of SCIM groups.
	StartIndex   int                `json:"startIndex,omitempty"`   // The start index for the SCIM groups.
	ItemsPerPage int                `json:"itemsPerPage,omitempty"` // The number of SCIM groups per page.
	Resources    []*ScimGroupScheme `json:"Resources,omitempty"`    // The SCIM groups on the page.
}

// ScimGroupScheme represents a SCIM group.
type ScimGroupScheme struct {
	Schemas     []string                 `json:"schemas,omitempty"`     // The schemas for the SCIM group.
	ID          string                   `json:"id,omitempty"`          // The ID of the SCIM group.
	ExternalID  string                   `json:"externalId,omitempty"`  // The external ID of the SCIM group.
	DisplayName string                   `json:"displayName,omitempty"` // The display name of the SCIM group.
	Members     []*ScimGroupMemberScheme `json:"members,omitempty"`     // The members of the SCIM group.
	Meta        *ScimMetadata            `json:"meta,omitempty"`        // The metadata for the SCIM group.
}

// ScimGroupMemberScheme represents a member of a SCIM group.
type ScimGroupMemberScheme struct {
	Type    string `json:"type,omitempty"`    // The type of the member.
	Value   string `json:"value,omitempty"`   // The value of the member.
	Display string `json:"display,omitempty"` // The display of the member.
	Ref     string `json:"$ref,omitempty"`    // The reference of the member.
}

// ScimMetadata represents the metadata for a SCIM group.
type ScimMetadata struct {
	ResourceType string `json:"resourceType,omitempty"` // The resource type of the SCIM group.
	Location     string `json:"location,omitempty"`     // The location of the SCIM group.
	LastModified string `json:"lastModified,omitempty"` // The last modified time of the SCIM group.
	Created      string `json:"created,omitempty"`      // The creation time of the SCIM group.
}
