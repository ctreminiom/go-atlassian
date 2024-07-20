package models

// SpaceChunkV2Scheme represents a chunk of spaces in Confluence.
type SpaceChunkV2Scheme struct {
	Results []*SpaceSchemeV2 `json:"results,omitempty"` // The spaces in the chunk.
	Links   struct {
		Next string `json:"next"` // The link to the next chunk of spaces.
	} `json:"_links"`
}

// SpacePageLinkSchemeV2 represents the links of a page of spaces in Confluence.
type SpacePageLinkSchemeV2 struct {
	Next string `json:"next,omitempty"` // The link to the next page of spaces.
}

// GetSpacesOptionSchemeV2 represents the options for getting spaces in Confluence.
type GetSpacesOptionSchemeV2 struct {
	IDs               []string // The IDs of the spaces.
	Keys              []string // The keys of the spaces.
	Type              string   // The type of the spaces.
	Status            string   // The status of the spaces.
	Labels            []string // The labels of the spaces.
	Sort              string   // The sort order of the spaces.
	DescriptionFormat string   // The format of the description of the spaces.
	SerializeIDs      bool     // Indicates if the IDs of the spaces should be serialized.
}

// SpaceSchemeV2 represents a space in Confluence.
type SpaceSchemeV2 struct {
	ID          string                    `json:"id,omitempty"`          // The ID of the space.
	Key         string                    `json:"key,omitempty"`         // The key of the space.
	Name        string                    `json:"name,omitempty"`        // The name of the space.
	Type        string                    `json:"type,omitempty"`        // The type of the space.
	Status      string                    `json:"status,omitempty"`      // The status of the space.
	HomepageID  string                    `json:"homepageId,omitempty"`  // The ID of the home page of the space.
	Description *SpaceDescriptionSchemeV2 `json:"description,omitempty"` // The description of the space.
}

// SpaceDescriptionSchemeV2 represents the description of a space in Confluence.
type SpaceDescriptionSchemeV2 struct {
	Plain *PageBodyRepresentationScheme `json:"plain,omitempty"` // The plain text description of the space.
	View  *PageBodyRepresentationScheme `json:"view,omitempty"`  // The view description of the space.
}

// SpacePermissionPageScheme represents a page of space permissions in Confluence.
type SpacePermissionPageScheme struct {
	Results []*SpacePermissionsV2Scheme    `json:"results,omitempty"` // The space permissions in the page.
	Links   *SpacePermissionPageLinkScheme `json:"_links,omitempty"`  // The links of the page.
}

// SpacePermissionsV2Scheme represents a version 2 space permission in Confluence.
type SpacePermissionsV2Scheme struct {
	ID        string                           `json:"id"`                  // The ID of the space permission.
	Principal *SpacePermissionsPrincipalScheme `json:"principal,omitempty"` // The principal of the space permission.
	Operation *SpacePermissionsOperationScheme `json:"operation,omitempty"` // The operation of the space permission.
}

// SpacePermissionsPrincipalScheme represents a principal in a space permission in Confluence.
type SpacePermissionsPrincipalScheme struct {
	Type string `json:"type,omitempty"` // The type of the principal.
	ID   string `json:"id,omitempty"`   // The ID of the principal.
}

// SpacePermissionsOperationScheme represents an operation in a space permission in Confluence.
type SpacePermissionsOperationScheme struct {
	Key        string `json:"key,omitempty"`        // The key of the operation.
	TargetType string `json:"targetType,omitempty"` // The target type of the operation.
}

// SpacePermissionPageLinkScheme represents the links of a page of space permissions in Confluence.
type SpacePermissionPageLinkScheme struct {
	Next string `json:"next,omitempty"` // The link to the next page of space permissions.
}
