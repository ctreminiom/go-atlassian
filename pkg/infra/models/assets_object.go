package models

// ObjectReferenceTypeInfoScheme represents the information about an object reference type.
// ReferenceTypes is a slice of the reference types.
// ObjectType is the type of the object.
// NumberOfReferencedObjects is the number of objects that reference this type.
// OpenIssuesExists is a boolean indicating if there are open issues related to this type.
type ObjectReferenceTypeInfoScheme struct {
	ReferenceTypes            []*TypeReferenceScheme `json:"referenceTypes,omitempty"`            // The reference types of the object.
	ObjectType                *ObjectTypeScheme      `json:"objectType,omitempty"`                // The type of the object.
	NumberOfReferencedObjects int                    `json:"numberOfReferencedObjects,omitempty"` // The number of referenced objects.
	OpenIssuesExists          bool                   `json:"openIssuesExists,omitempty"`          // Indicates if there are open issues related to this type.
}

// TypeReferenceScheme represents a type reference.
// WorkspaceID is the ID of the workspace.
// GlobalID is the global ID of the type reference.
// ID is the unique identifier of the type reference.
// Name is the name of the type reference.
// Description is the description of the type reference.
// Color is the color of the type reference.
// URL16 is the URL for the 16x16 version of the type reference.
// Removable indicates if the type reference is removable.
// ObjectSchemaID is the ID of the object schema.
type TypeReferenceScheme struct {
	WorkspaceID    string `json:"workspaceId,omitempty"`    // The ID of the workspace.
	GlobalID       string `json:"globalId,omitempty"`       // The global ID of the type reference.
	ID             string `json:"id,omitempty"`             // The ID of the type reference.
	Name           string `json:"name,omitempty"`           // The name of the type reference.
	Description    string `json:"description,omitempty"`    // The description of the type reference.
	Color          string `json:"color,omitempty"`          // The color of the type reference.
	URL16          string `json:"url16,omitempty"`          // The URL for the 16x16 version of the type reference.
	Removable      bool   `json:"removable,omitempty"`      // Indicates if the type reference is removable.
	ObjectSchemaID string `json:"objectSchemaId,omitempty"` // The ID of the object schema.
}

// ObjectHistoryScheme represents the history of an object.
// Actor is the actor who made the change.
// ID is the unique identifier of the history entry.
// AffectedAttribute is the attribute that was affected.
// OldValue is the old value of the attribute.
// NewValue is the new value of the attribute.
// Type is the type of the history entry.
// Created is the creation time of the history entry.
// ObjectID is the ID of the object.
type ObjectHistoryScheme struct {
	Actor             *ObjectHistoryActorScheme `json:"actor,omitempty"`             // The actor who made the change.
	ID                string                    `json:"id,omitempty"`                // The ID of the history entry.
	AffectedAttribute string                    `json:"affectedAttribute,omitempty"` // The affected attribute.
	OldValue          string                    `json:"oldValue,omitempty"`          // The old value of the attribute.
	NewValue          string                    `json:"newValue,omitempty"`          // The new value of the attribute.
	Type              int                       `json:"type,omitempty"`              // The type of the history entry.
	Created           string                    `json:"created,omitempty"`           // The creation time of the history entry.
	ObjectID          string                    `json:"objectId,omitempty"`          // The ID of the object.
}

// ObjectHistoryActorScheme represents the actor who made a change in the object history.
// AvatarURL is the URL of the actor's avatar.
// DisplayName is the display name of the actor.
// Name is the name of the actor.
// Key is the key of the actor.
// EmailAddress is the email address of the actor.
// HTML is the HTML representation of the actor.
// RenderedLink is the rendered link of the actor.
// IsDeleted indicates if the actor is deleted.
// LastSeenVersion is the last seen version of the actor.
// Self is the self URL of the actor.
type ObjectHistoryActorScheme struct {
	AvatarURL       string `json:"avatarUrl,omitempty"`       // The URL of the actor's avatar.
	DisplayName     string `json:"displayName,omitempty"`     // The display name of the actor.
	Name            string `json:"name,omitempty"`            // The name of the actor.
	Key             string `json:"key,omitempty"`             // The key of the actor.
	EmailAddress    string `json:"emailAddress,omitempty"`    // The email address of the actor.
	HTML            string `json:"html,omitempty"`            // The HTML representation of the actor.
	RenderedLink    string `json:"renderedLink,omitempty"`    // The rendered link of the actor.
	IsDeleted       bool   `json:"isDeleted,omitempty"`       // Indicates if the actor is deleted.
	LastSeenVersion string `json:"lastSeenVersion,omitempty"` // The last seen version of the actor.
	Self            string `json:"self,omitempty"`            // The self URL of the actor.
}

// ObjectPayloadScheme represents the payload for an object.
// ObjectTypeID is the ID of the object type.
// AvatarUUID is the UUID of the avatar.
// HasAvatar indicates if the object has an avatar.
// Attributes is a slice of the attributes of the object.
type ObjectPayloadScheme struct {
	ObjectTypeID string                          `json:"objectTypeId,omitempty"` // The ID of the object type.
	AvatarUUID   string                          `json:"avatarUUID,omitempty"`   // The UUID of the avatar.
	HasAvatar    bool                            `json:"hasAvatar,omitempty"`    // Indicates if the object has an avatar.
	Attributes   []*ObjectPayloadAttributeScheme `json:"attributes,omitempty"`   // The attributes of the object.
}

// ObjectPayloadAttributeScheme represents an attribute in the payload for an object.
// ObjectTypeAttributeID is the ID of the object type attribute.
// ObjectAttributeValues is a slice of the values of the object attribute.
type ObjectPayloadAttributeScheme struct {
	ObjectTypeAttributeID string                               `json:"objectTypeAttributeId,omitempty"` // The ID of the object type attribute.
	ObjectAttributeValues []*ObjectPayloadAttributeValueScheme `json:"objectAttributeValues,omitempty"` // The values of the object attribute.
}

// ObjectPayloadAttributeValueScheme represents the value of an attribute in the payload for an object.
// Value is the value of the attribute.
type ObjectPayloadAttributeValueScheme struct {
	Value string `json:"value,omitempty"` // The value of the attribute.
}

// ObjectScheme represents an object.
// WorkspaceID is the ID of the workspace.
// GlobalID is the global ID of the object.
// ID is the unique identifier of the object.
// Label is the label of the object.
// ObjectKey is the key of the object.
// Avatar is the avatar of the object.
// ObjectType is the type of the object.
// Created is the creation time of the object.
// Updated is the update time of the object.
// HasAvatar indicates if the object has an avatar.
// Timestamp is the timestamp of the object.
// Attributes is a slice of the attributes of the object.
// Links is the links of the object.
type ObjectScheme struct {
	WorkspaceID string                   `json:"workspaceId,omitempty"` // The ID of the workspace.
	GlobalID    string                   `json:"globalId,omitempty"`    // The global ID of the object.
	ID          string                   `json:"id,omitempty"`          // The ID of the object.
	Label       string                   `json:"label,omitempty"`       // The label of the object.
	ObjectKey   string                   `json:"objectKey,omitempty"`   // The key of the object.
	Avatar      *ObjectAvatarScheme      `json:"avatar,omitempty"`      // The avatar of the object.
	ObjectType  *ObjectTypeScheme        `json:"objectType,omitempty"`  // The type of the object.
	Created     string                   `json:"created,omitempty"`     // The creation time of the object.
	Updated     string                   `json:"updated,omitempty"`     // The update time of the object.
	HasAvatar   bool                     `json:"hasAvatar,omitempty"`   // Indicates if the object has an avatar.
	Timestamp   int                      `json:"timestamp,omitempty"`   // The timestamp of the object.
	Attributes  []*ObjectAttributeScheme `json:"attributes"`            // The attributes of the object.
	Links       *ObjectLinksScheme       `json:"_links,omitempty"`      // The links of the object.
}

// ObjectAvatarScheme represents an avatar of an object.
// WorkspaceID is the ID of the workspace.
// GlobalID is the global ID of the avatar.
// ID is the unique identifier of the avatar.
// AvatarUUID is the UUID of the avatar.
// URL16 is the URL for the 16x16 version of the avatar.
// URL48 is the URL for the 48x48 version of the avatar.
// URL72 is the URL for the 72x72 version of the avatar.
// URL144 is the URL for the 144x144 version of the avatar.
// URL288 is the URL for the 288x288 version of the avatar.
// ObjectID is the ID of the object.
type ObjectAvatarScheme struct {
	WorkspaceID string `json:"workspaceId,omitempty"` // The ID of the workspace.
	GlobalID    string `json:"globalId,omitempty"`    // The global ID of the avatar.
	ID          string `json:"id,omitempty"`          // The ID of the avatar.
	AvatarUUID  string `json:"avatarUUID,omitempty"`  // The UUID of the avatar.
	URL16       string `json:"url16,omitempty"`       // The URL for the 16x16 version of the avatar.
	URL48       string `json:"url48,omitempty"`       // The URL for the 48x48 version of the avatar.
	URL72       string `json:"url72,omitempty"`       // The URL for the 72x72 version of the avatar.
	URL144      string `json:"url144,omitempty"`      // The URL for the 144x144 version of the avatar.
	URL288      string `json:"url288,omitempty"`      // The URL for the 288x288 version of the avatar.
	ObjectID    string `json:"objectId,omitempty"`    // The ID of the object.
}

// ObjectLinksScheme represents the links of an object.
// Self is the self URL of the object.
type ObjectLinksScheme struct {
	Self string `json:"self,omitempty"` // The self URL of the object.
}

// ObjectAttributeScheme represents an attribute of an object.
// WorkspaceID is the ID of the workspace.
// GlobalID is the global ID of the attribute.
// ID is the unique identifier of the attribute.
// ObjectTypeAttribute is the type of the attribute.
// ObjectTypeAttributeID is the ID of the attribute type.
// ObjectAttributeValues is a slice of the values of the attribute.
type ObjectAttributeScheme struct {
	WorkspaceID           string                                 `json:"workspaceId,omitempty"`           // The ID of the workspace.
	GlobalID              string                                 `json:"globalId,omitempty"`              // The global ID of the attribute.
	ID                    string                                 `json:"id,omitempty"`                    // The ID of the attribute.
	ObjectTypeAttribute   *ObjectTypeAttributeScheme             `json:"objectTypeAttribute,omitempty"`   // The type of the attribute.
	ObjectTypeAttributeID string                                 `json:"objectTypeAttributeId,omitempty"` // The ID of the attribute type.
	ObjectAttributeValues []*ObjectTypeAssetAttributeValueScheme `json:"objectAttributeValues,omitempty"` // The values of the attribute.
}

// ObjectTypeAttributePayloadScheme represents the payload for an attribute type.
// It includes various properties of the attribute type like name, label, description, type, default type ID, type value, etc.
type ObjectTypeAttributePayloadScheme struct {
	Name                    string   `json:"name,omitempty"`                    // The name of the attribute type.
	Label                   bool     `json:"label,omitempty"`                   // Indicates if the attribute type is a label.
	Description             string   `json:"description,omitempty"`             // The description of the attribute type.
	Type                    *int     `json:"type,omitempty"`                    // The type of the attribute type.
	DefaultTypeID           *int     `json:"defaultTypeId,omitempty"`           // The default type ID of the attribute type.
	TypeValue               string   `json:"typeValue,omitempty"`               // The type value of the attribute type.
	TypeValueMulti          []string `json:"typeValueMulti,omitempty"`          // The multiple type values of the attribute type.
	AdditionalValue         string   `json:"additionalValue,omitempty"`         // The additional value of the attribute type.
	MinimumCardinality      *int     `json:"minimumCardinality,omitempty"`      // The minimum cardinality of the attribute type.
	MaximumCardinality      *int     `json:"maximumCardinality,omitempty"`      // The maximum cardinality of the attribute type.
	Suffix                  string   `json:"suffix,omitempty"`                  // The suffix of the attribute type.
	IncludeChildObjectTypes bool     `json:"includeChildObjectTypes,omitempty"` // Indicates if child object types are included.
	Hidden                  bool     `json:"hidden,omitempty"`                  // Indicates if the attribute type is hidden.
	UniqueAttribute         bool     `json:"uniqueAttribute,omitempty"`         // Indicates if the attribute type is unique.
	Summable                bool     `json:"summable,omitempty"`                // Indicates if the attribute type is summable.
	RegexValidation         string   `json:"regexValidation,omitempty"`         // The regex validation of the attribute type.
	QlQuery                 string   `json:"qlQuery,omitempty"`                 // The QL query of the attribute type.
	Iql                     string   `json:"iql,omitempty"`                     // The IQL of the attribute type.
	Options                 string   `json:"options,omitempty"`                 // The options of the attribute type.
}

// ObjectTypeAttributeScheme represents an attribute type of an object.
// It includes various properties of the attribute type like workspace ID, global ID, ID, object type, name, label, type, description, etc.
type ObjectTypeAttributeScheme struct {
	WorkspaceID             string                                       `json:"workspaceId,omitempty"`             // The ID of the workspace.
	GlobalID                string                                       `json:"globalId,omitempty"`                // The global ID of the attribute type.
	ID                      string                                       `json:"id,omitempty"`                      // The ID of the attribute type.
	ObjectType              *ObjectTypeScheme                            `json:"objectType,omitempty"`              // The type of the object.
	Name                    string                                       `json:"name,omitempty"`                    // The name of the attribute type.
	Label                   bool                                         `json:"label,omitempty"`                   // Indicates if the attribute type is a label.
	Type                    int                                          `json:"type,omitempty"`                    // The type of the attribute type.
	Description             string                                       `json:"description,omitempty"`             // The description of the attribute type.
	DefaultType             *ObjectTypeAssetAttributeDefaultTypeScheme   `json:"defaultType,omitempty"`             // The default type of the attribute type.
	TypeValue               string                                       `json:"typeValue,omitempty"`               // The type value of the attribute type.
	TypeValueMulti          []string                                     `json:"typeValueMulti,omitempty"`          // The multiple type values of the attribute type.
	AdditionalValue         string                                       `json:"additionalValue,omitempty"`         // The additional value of the attribute type.
	ReferenceType           *ObjectTypeAssetAttributeReferenceTypeScheme `json:"referenceType,omitempty"`           // The reference type of the attribute type.
	ReferenceObjectTypeID   string                                       `json:"referenceObjectTypeId,omitempty"`   // The ID of the reference object type.
	ReferenceObjectType     *ObjectTypeScheme                            `json:"referenceObjectType,omitempty"`     // The reference object type.
	Editable                bool                                         `json:"editable,omitempty"`                // Indicates if the attribute type is editable.
	System                  bool                                         `json:"system,omitempty"`                  // Indicates if the attribute type is a system attribute.
	Indexed                 bool                                         `json:"indexed,omitempty"`                 // Indicates if the attribute type is indexed.
	Sortable                bool                                         `json:"sortable,omitempty"`                // Indicates if the attribute type is sortable.
	Summable                bool                                         `json:"summable,omitempty"`                // Indicates if the attribute type is summable.
	MinimumCardinality      int                                          `json:"minimumCardinality,omitempty"`      // The minimum cardinality of the attribute type.
	MaximumCardinality      int                                          `json:"maximumCardinality,omitempty"`      // The maximum cardinality of the attribute type.
	Suffix                  string                                       `json:"suffix,omitempty"`                  // The suffix of the attribute type.
	Removable               bool                                         `json:"removable,omitempty"`               // Indicates if the attribute type is removable.
	ObjectAttributeExists   bool                                         `json:"objectAttributeExists,omitempty"`   // Indicates if the attribute exists in the object.
	Hidden                  bool                                         `json:"hidden,omitempty"`                  // Indicates if the attribute type is hidden.
	IncludeChildObjectTypes bool                                         `json:"includeChildObjectTypes,omitempty"` // Indicates if child object types are included.
	UniqueAttribute         bool                                         `json:"uniqueAttribute,omitempty"`         // Indicates if the attribute type is unique.
	RegexValidation         string                                       `json:"regexValidation,omitempty"`         // The regex validation of the attribute type.
	Iql                     string                                       `json:"iql,omitempty"`                     // The IQL of the attribute type.
	QlQuery                 string                                       `json:"qlQuery,omitempty"`                 // The QL query of the attribute type.
	Options                 string                                       `json:"options,omitempty"`                 // The options of the attribute type.
	Position                int                                          `json:"position,omitempty"`                // The position of the attribute type.
}

// ObjectTypeAssetAttributeDefaultTypeScheme represents the default type of an attribute in an asset.
// ID is the unique identifier of the default type.
// Name is the name of the default type.
type ObjectTypeAssetAttributeDefaultTypeScheme struct {
	ID   int    `json:"id,omitempty"`   // The ID of the default type.
	Name string `json:"name,omitempty"` // The name of the default type.
}

// ObjectTypeAssetAttributeReferenceTypeScheme represents a reference type of an attribute in an asset.
// WorkspaceID is the ID of the workspace.
// GlobalID is the global ID of the reference type.
// Name is the name of the reference type.
type ObjectTypeAssetAttributeReferenceTypeScheme struct {
	WorkspaceID string `json:"workspaceId,omitempty"` // The ID of the workspace.
	GlobalID    string `json:"globalId,omitempty"`    // The global ID of the reference type.
	Name        string `json:"name,omitempty"`        // The name of the reference type.
}

// ObjectTypeAssetAttributeValueScheme represents the value of an attribute in an asset.
// Value is the value of the attribute.
// DisplayValue is the display value of the attribute.
// SearchValue is the search value of the attribute.
// Group is the group of the attribute value.
// Status is the status of the attribute value.
// AdditionalValue is the additional value of the attribute.
type ObjectTypeAssetAttributeValueScheme struct {
	Value           string                                    `json:"value,omitempty"`           // The value of the attribute.
	DisplayValue    string                                    `json:"displayValue,omitempty"`    // The display value of the attribute.
	SearchValue     string                                    `json:"searchValue,omitempty"`     // The search value of the attribute.
	Group           *ObjectTypeAssetAttributeValueGroupScheme `json:"group,omitempty"`           // The group of the attribute value.
	Status          *ObjectTypeAssetAttributeStatusScheme     `json:"status,omitempty"`          // The status of the attribute value.
	AdditionalValue string                                    `json:"additionalValue,omitempty"` // The additional value of the attribute.
}

// ObjectTypeAssetAttributeValueGroupScheme represents a group of attribute values in an asset.
// AvatarURL is the URL of the avatar for the group.
// Name is the name of the group.
type ObjectTypeAssetAttributeValueGroupScheme struct {
	AvatarURL string `json:"avatarUrl,omitempty"` // The URL of the avatar for the group.
	Name      string `json:"name,omitempty"`      // The name of the group.
}

// ObjectTypeAssetAttributeStatusScheme represents the status of an attribute value in an asset.
// ID is the unique identifier of the status.
// Name is the name of the status.
// Category is the category of the status.
type ObjectTypeAssetAttributeStatusScheme struct {
	ID       string `json:"id,omitempty"`       // The ID of the status.
	Name     string `json:"name,omitempty"`     // The name of the status.
	Category int    `json:"category,omitempty"` // The category of the status.
}

// ObjectListScheme represents a list of objects.
type ObjectListScheme struct {
	ObjectEntries         []*ObjectScheme              `json:"objectEntries,omitempty"`         // The objects in the list.
	ObjectTypeAttributes  []*ObjectTypeAttributeScheme `json:"objectTypeAttributes,omitempty"`  // The attributes of the object type.
	ObjectTypeID          string                       `json:"objectTypeId,omitempty"`          // The ID of the object type.
	ObjectTypeIsInherited bool                         `json:"objectTypeIsInherited,omitempty"` // Indicates if the object type is inherited.
	AbstractObjectType    bool                         `json:"abstractObjectType,omitempty"`    // Indicates if the object type is abstract.
	TotalFilterCount      int                          `json:"totalFilterCount,omitempty"`      // The total count of filters.
	StartIndex            int                          `json:"startIndex,omitempty"`            // The start index of the list.
	ToIndex               int                          `json:"toIndex,omitempty"`               // The end index of the list.
	PageObjectSize        int                          `json:"pageObjectSize,omitempty"`        // The size of the page of objects.
	PageNumber            int                          `json:"pageNumber,omitempty"`            // The number of the page of objects.
	OrderByTypeAttrID     int                          `json:"orderByTypeAttrId,omitempty"`     // The ID of the attribute type to order by.
	OrderWay              string                       `json:"orderWay,omitempty"`              // The way to order the list.
	QlQuery               string                       `json:"qlQuery,omitempty"`               // The QL query to filter the list.
	QlQuerySearchResult   bool                         `json:"qlQuerySearchResult,omitempty"`   // Indicates if the QL query search result is included.
	Iql                   string                       `json:"iql,omitempty"`                   // The IQL to filter the list.
	IqlSearchResult       bool                         `json:"iqlSearchResult,omitempty"`       // Indicates if the IQL search result is included.
	ConversionPossible    bool                         `json:"conversionPossible,omitempty"`    // Indicates if a conversion is possible.
}

// ObjectListResultScheme represents the result of a list of objects.
// StartAt is the starting index of the list.
// MaxResults is the maximum number of results in the list.
// Total is the total number of objects in the list.
// IsLast indicates if this is the last page of the list.
// Values is a slice of the objects in the list.
// ObjectTypeAttributes is a slice of the attributes of the object type.
type ObjectListResultScheme struct {
	StartAt              int                          `json:"startAt,omitempty"`              // The starting index of the list.
	MaxResults           int                          `json:"maxResults,omitempty"`           // The maximum number of results in the list.
	Total                int                          `json:"total,omitempty"`                // The total number of objects in the list.
	IsLast               bool                         `json:"isLast,omitempty"`               // Indicates if this is the last page of the list.
	Values               []*ObjectScheme              `json:"values,omitempty"`               // The objects in the list.
	ObjectTypeAttributes []*ObjectTypeAttributeScheme `json:"objectTypeAttributes,omitempty"` // The attributes of the object type.
}

// ObjectSearchParamsScheme represents the search parameters for an object.
// Query is the AQL that will fetch the objects.
// Iql is deprecated and Query should be used instead.
// ObjectTypeID is the ID of the object type.
// Page is the requested page to be loaded for a paginated result.
// ResultPerPage is how many objects should be returned in the request.
// OrderByTypeAttributeID is which attribute should be used to order by.
// Asc is to sort objects in ascending order or descending order based on the attribute identified by OrderByTypeAttributeID.
// ObjectID identifies an object that should be included in the result.
// ObjectSchemaID is the ID of the object schema.
// IncludeAttributes indicates if attribute values should be included in the response.
// AttributesToDisplay identifies attributes to be displayed.
type ObjectSearchParamsScheme struct {
	Query                  string                     `json:"qlQuery,omitempty"`             // The AQL that will fetch the objects.
	Iql                    string                     `json:"iql,omitempty"`                 // Deprecated. Use Query instead.
	ObjectTypeID           string                     `json:"objectTypeId,omitempty"`        // The ID of the object type.
	Page                   int                        `json:"page,omitempty"`                // The requested page to be loaded for a paginated result.
	ResultPerPage          int                        `json:"resultsPerPage,omitempty"`      // How many objects should be returned in the request.
	OrderByTypeAttributeID int                        `json:"orderByTypeAttrId,omitempty"`   // Which attribute should be used to order by.
	Asc                    int                        `json:"asc,omitempty"`                 // Sort objects in ascending order or descending order based on the attribute identified by OrderByTypeAttributeID.
	ObjectID               string                     `json:"objectId,omitempty"`            // Identifies an object that should be included in the result.
	ObjectSchemaID         string                     `json:"objectSchemaId,omitempty"`      // The ID of the object schema.
	IncludeAttributes      bool                       `json:"includeAttributes,omitempty"`   // Should attribute values be included in the response.
	AttributesToDisplay    *AttributesToDisplayScheme `json:"attributesToDisplay,omitempty"` // Identifies attributes to be displayed.
}

// AttributesToDisplayScheme represents a scheme for attributes to be displayed.
// AttributesToDisplayIDs is a slice of the IDs of the attributes to be displayed.
type AttributesToDisplayScheme struct {
	AttributesToDisplayIDs []int `json:"attributesToDisplayIds,omitempty"` // The IDs of the attributes to be displayed.
}
