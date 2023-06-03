package models

type ObjectSchemaPageScheme struct {
	StartAt    int                   `json:"startAt,omitempty"`
	MaxResults int                   `json:"maxResults,omitempty"`
	Total      int                   `json:"total,omitempty"`
	Values     []*ObjectSchemaScheme `json:"values,omitempty"`
}

type ObjectSchemaScheme struct {
	WorkspaceId     string `json:"workspaceId,omitempty"`
	GlobalId        string `json:"globalId,omitempty"`
	Id              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	ObjectSchemaKey string `json:"objectSchemaKey,omitempty"`
	Description     string `json:"description,omitempty"`
	Status          string `json:"status,omitempty"`
	Created         string `json:"created,omitempty"`
	Updated         string `json:"updated,omitempty"`
	ObjectCount     int    `json:"objectCount,omitempty"`
	ObjectTypeCount int    `json:"objectTypeCount,omitempty"`
	CanManage       bool   `json:"canManage,omitempty"`
}

type ObjectSchemaPayloadScheme struct {
	Name            string `json:"name,omitempty"`
	ObjectSchemaKey string `json:"objectSchemaKey,omitempty"`
	Description     string `json:"description,omitempty"`
}

type ObjectSchemaAttributesParamsScheme struct {

	// OnlyValueEditable return only values that are associated with values that can be edited
	OnlyValueEditable bool

	// Extended include the object type with each object type attribute
	Extended bool

	// Query it's a query that will be used to filter object type attributes by their name
	Query string
}

type ObjectSchemaTypePageScheme struct {
	Entries []*ObjectAssetTypeScheme `json:"entries,omitempty"`
}
