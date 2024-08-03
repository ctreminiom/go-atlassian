package models

// ObjectTypePositionPayloadScheme represents the payload for the position of an object type.
// ToObjectTypeID is the ID of the object type to which the position is related.
// Position is the position of the object type.
type ObjectTypePositionPayloadScheme struct {
	ToObjectTypeID string `json:"toObjectTypeId,omitempty"` // The ID of the object type to which the position is related.
	Position       int    `json:"position,omitempty"`       // The position of the object type.
}

// ObjectTypeScheme represents an object type.
// WorkspaceID is the ID of the workspace.
// GlobalID is the global ID of the object type.
// ID is the unique identifier of the object type.
// Name is the name of the object type.
// Description is the description of the object type.
// Icon is the icon of the object type.
// Position is the position of the object type.
// Created is the creation time of the object type.
// Updated is the update time of the object type.
// ObjectCount is the number of objects of the object type.
// ParentObjectTypeID is the ID of the parent object type.
// ObjectSchemaID is the ID of the object schema.
// Inherited indicates if the object type is inherited.
// AbstractObjectType indicates if the object type is abstract.
// ParentObjectTypeInherited indicates if the parent object type is inherited.
type ObjectTypeScheme struct {
	WorkspaceID               string      `json:"workspaceId,omitempty"`               // The ID of the workspace.
	GlobalID                  string      `json:"globalId,omitempty"`                  // The global ID of the object type.
	ID                        string      `json:"id,omitempty"`                        // The ID of the object type.
	Name                      string      `json:"name,omitempty"`                      // The name of the object type.
	Description               string      `json:"description,omitempty"`               // The description of the object type.
	Icon                      *IconScheme `json:"icon,omitempty"`                      // The icon of the object type.
	Position                  int         `json:"position,omitempty"`                  // The position of the object type.
	Created                   string      `json:"created,omitempty"`                   // The creation time of the object type.
	Updated                   string      `json:"updated,omitempty"`                   // The update time of the object type.
	ObjectCount               int         `json:"objectCount,omitempty"`               // The number of objects of the object type.
	ParentObjectTypeID        string      `json:"parentObjectTypeId,omitempty"`        // The ID of the parent object type.
	ObjectSchemaID            string      `json:"objectSchemaId,omitempty"`            // The ID of the object schema.
	Inherited                 bool        `json:"inherited,omitempty"`                 // Indicates if the object type is inherited.
	AbstractObjectType        bool        `json:"abstractObjectType,omitempty"`        // Indicates if the object type is abstract.
	ParentObjectTypeInherited bool        `json:"parentObjectTypeInherited,omitempty"` // Indicates if the parent object type is inherited.
}

// ObjectTypePayloadScheme represents the payload for an object type.
// Name is the name of the object type.
// Description is the description of the object type.
// IconID is the ID of the icon of the object type.
// ObjectSchemaID is the ID of the object schema.
// ParentObjectTypeID is the ID of the parent object type.
// Inherited indicates if the object type is inherited.
// AbstractObjectType indicates if the object type is abstract.
type ObjectTypePayloadScheme struct {
	Name               string `json:"name,omitempty"`               // The name of the object type.
	Description        string `json:"description,omitempty"`        // The description of the object type.
	IconID             string `json:"iconId,omitempty"`             // The ID of the icon of the object type.
	ObjectSchemaID     string `json:"objectSchemaId,omitempty"`     // The ID of the object schema.
	ParentObjectTypeID string `json:"parentObjectTypeId,omitempty"` // The ID of the parent object type.
	Inherited          bool   `json:"inherited,omitempty"`          // Indicates if the object type is inherited.
	AbstractObjectType bool   `json:"abstractObjectType,omitempty"` // Indicates if the object type is abstract.
}

// ObjectTypeAttributesParamsScheme represents the search parameters for an object type's attributes.
// OnlyValueEditable return only values that are associated with values that can be edited.
// OrderByName order by the name of the attribute.
// Query is a query that will be used to filter object type attributes by their name.
// IncludeValuesExist include values that exist.
// ExcludeParentAttributes exclude parent attributes.
// IncludeChildren include children.
// OrderByRequired order by required.
type ObjectTypeAttributesParamsScheme struct {
	OnlyValueEditable       bool   // Return only values that are associated with values that can be edited.
	OrderByName             bool   // Order by the name of the attribute.
	Query                   string // A query that will be used to filter object type attributes by their name.
	IncludeValuesExist      bool   // Include values that exist.
	ExcludeParentAttributes bool   // Exclude parent attributes.
	IncludeChildren         bool   // Include children.
	OrderByRequired         bool   // Order by required.
}
