package models

// ContentRestrictionPageScheme represents a page of content restrictions in Confluence.
type ContentRestrictionPageScheme struct {
	Start            int                         `json:"start,omitempty"`            // The start index of the content restrictions in the page.
	Limit            int                         `json:"limit,omitempty"`            // The limit of the content restrictions in the page.
	Size             int                         `json:"size,omitempty"`             // The size of the content restrictions in the page.
	RestrictionsHash string                      `json:"restrictionsHash,omitempty"` // The hash of the restrictions.
	Results          []*ContentRestrictionScheme `json:"results,omitempty"`          // The content restrictions in the page.
}

// ContentRestrictionScheme represents a content restriction in Confluence.
type ContentRestrictionScheme struct {
	Operation    string                              `json:"operation,omitempty"`    // The operation of the restriction.
	Restrictions *ContentRestrictionDetailScheme     `json:"restrictions,omitempty"` // The details of the restriction.
	Content      *ContentScheme                      `json:"content,omitempty"`      // The content of the restriction.
	Expandable   *ContentRestrictionExpandableScheme `json:"_expandable,omitempty"`  // The expandable fields of the restriction.
}

// ContentRestrictionExpandableScheme represents the expandable fields of a content restriction in Confluence.
type ContentRestrictionExpandableScheme struct {
	Restrictions string `json:"restrictions,omitempty"` // The restrictions of the content restriction.
	Content      string `json:"content,omitempty"`      // The content of the content restriction.
}

// ContentRestrictionDetailScheme represents the details of a content restriction in Confluence.
type ContentRestrictionDetailScheme struct {
	User  *UserPermissionScheme  `json:"user,omitempty"`  // The user permissions of the restriction.
	Group *GroupPermissionScheme `json:"group,omitempty"` // The group permissions of the restriction.
}

// ContentRestrictionUpdatePayloadScheme represents the payload for updating a content restriction in Confluence.
type ContentRestrictionUpdatePayloadScheme struct {
	Results []*ContentRestrictionUpdateScheme `json:"results,omitempty"` // The content restrictions to be updated.
}

// ContentRestrictionUpdateScheme represents a content restriction to be updated in Confluence.
type ContentRestrictionUpdateScheme struct {
	Operation    string                                     `json:"operation,omitempty"` // The operation of the restriction.
	Restrictions *ContentRestrictionRestrictionUpdateScheme `json:"restrictions"`        // The details of the restriction to be updated.
	Content      *ContentScheme                             `json:"content,omitempty"`   // The content of the restriction to be updated.
}

// ContentRestrictionRestrictionUpdateScheme represents the details of a content restriction to be updated in Confluence.
type ContentRestrictionRestrictionUpdateScheme struct {
	Group []*SpaceGroupScheme  `json:"group"` // The group permissions of the restriction to be updated.
	User  []*ContentUserScheme `json:"user"`  // The user permissions of the restriction to be updated.
}

// ContentRestrictionByOperationScheme represents a content restriction by operation in Confluence.
type ContentRestrictionByOperationScheme struct {
	OperationType *ContentRestrictionScheme `json:"operationType,omitempty"` // The type of the operation of the restriction.
}
