package models

// ObjectSchemaPageScheme represents a paginated list of object schemas.
// StartAt is the starting index of the list.
// MaxResults is the maximum number of results in the list.
// Total is the total number of object schemas in the list.
// Values is a slice of the object schemas in the list.
type ObjectSchemaPageScheme struct {
	StartAt    int                   `json:"startAt,omitempty"`    // The starting index of the list.
	MaxResults int                   `json:"maxResults,omitempty"` // The maximum number of results in the list.
	Total      int                   `json:"total,omitempty"`      // The total number of object schemas in the list.
	Values     []*ObjectSchemaScheme `json:"values,omitempty"`     // The object schemas in the list.
}

// ObjectSchemaScheme represents an object schema.
type ObjectSchemaScheme struct {
	WorkspaceID     string `json:"workspaceId,omitempty"`     // The ID of the workspace.
	GlobalID        string `json:"globalId,omitempty"`        // The global ID of the object schema.
	ID              string `json:"id,omitempty"`              // The ID of the object schema.
	Name            string `json:"name,omitempty"`            // The name of the object schema.
	ObjectSchemaKey string `json:"objectSchemaKey,omitempty"` // The key of the object schema.
	Description     string `json:"description,omitempty"`     // The description of the object schema.
	Status          string `json:"status,omitempty"`          // The status of the object schema.
	Created         string `json:"created,omitempty"`         // The creation time of the object schema.
	Updated         string `json:"updated,omitempty"`         // The update time of the object schema.
	ObjectCount     int    `json:"objectCount,omitempty"`     // The number of objects in the object schema.
	ObjectTypeCount int    `json:"objectTypeCount,omitempty"` // The number of object types in the object schema.
	CanManage       bool   `json:"canManage,omitempty"`       // Indicates if the object schema can be managed.
}

// ObjectSchemaPayloadScheme represents the payload for an object schema.
// Name is the name of the object schema.
// ObjectSchemaKey is the key of the object schema.
// Description is the description of the object schema.
type ObjectSchemaPayloadScheme struct {
	Name            string `json:"name,omitempty"`            // The name of the object schema.
	ObjectSchemaKey string `json:"objectSchemaKey,omitempty"` // The key of the object schema.
	Description     string `json:"description,omitempty"`     // The description of the object schema.
}

// ObjectSchemaAttributesParamsScheme represents the search parameters for an object schema's attributes.
// OnlyValueEditable return only values that are associated with values that can be edited.
// Extended include the object type with each object type attribute.
// Query it's a query that will be used to filter object type attributes by their name.
type ObjectSchemaAttributesParamsScheme struct {
	OnlyValueEditable bool   // Return only values that are associated with values that can be edited.
	Extended          bool   // Include the object type with each object type attribute.
	Query             string // A query that will be used to filter object type attributes by their name.
}

// ObjectSchemaTypePageScheme represents a page of object types.
// Entries is a slice of the object types on the page.
type ObjectSchemaTypePageScheme struct {
	Entries []*ObjectTypeScheme `json:"entries,omitempty"` // The object types on the page.
}
