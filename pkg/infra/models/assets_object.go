package models

type ObjectReferenceTypeInfoScheme struct {
	ReferenceTypes            []*TypeReferenceScheme `json:"referenceTypes,omitempty"`
	ObjectType                *ObjectTypeScheme      `json:"objectType,omitempty"`
	NumberOfReferencedObjects int                    `json:"numberOfReferencedObjects,omitempty"`
	OpenIssuesExists          bool                   `json:"openIssuesExists,omitempty"`
}

type TypeReferenceScheme struct {
	WorkspaceID    string `json:"workspaceId,omitempty"`
	GlobalID       string `json:"globalId,omitempty"`
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Color          string `json:"color,omitempty"`
	URL16          string `json:"url16,omitempty"`
	Removable      bool   `json:"removable,omitempty"`
	ObjectSchemaID string `json:"objectSchemaId,omitempty"`
}

type ObjectHistoryScheme struct {
	Actor             *ObjectHistoryActorScheme `json:"actor,omitempty"`
	Id                int                       `json:"id,omitempty"`
	AffectedAttribute string                    `json:"affectedAttribute,omitempty"`
	OldValue          string                    `json:"oldValue,omitempty"`
	NewValue          string                    `json:"newValue,omitempty"`
	Type              int                       `json:"type,omitempty"`
	Created           string                    `json:"created,omitempty"`
	ObjectId          int                       `json:"objectId,omitempty"`
}

type ObjectHistoryActorScheme struct {
	AvatarUrl       string `json:"avatarUrl,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`
	Name            string `json:"name,omitempty"`
	Key             string `json:"key,omitempty"`
	EmailAddress    string `json:"emailAddress,omitempty"`
	Html            string `json:"html,omitempty"`
	RenderedLink    string `json:"renderedLink,omitempty"`
	IsDeleted       bool   `json:"isDeleted,omitempty"`
	LastSeenVersion string `json:"lastSeenVersion,omitempty"`
	Self            string `json:"self,omitempty"`
}

type ObjectPayloadScheme struct {
	ObjectTypeID string                          `json:"objectTypeId,omitempty"`
	AvatarUUID   string                          `json:"avatarUUID,omitempty"`
	HasAvatar    bool                            `json:"hasAvatar,omitempty"`
	Attributes   []*ObjectPayloadAttributeScheme `json:"attributes,omitempty"`
}

type ObjectPayloadAttributeScheme struct {
	ObjectTypeAttributeID string                               `json:"objectTypeAttributeId,omitempty"`
	ObjectAttributeValues []*ObjectPayloadAttributeValueScheme `json:"objectAttributeValues,omitempty"`
}

type ObjectPayloadAttributeValueScheme struct {
	Value string `json:"value,omitempty"`
}

type ObjectScheme struct {
	WorkspaceId string                   `json:"workspaceId,omitempty"`
	GlobalId    string                   `json:"globalId,omitempty"`
	Id          string                   `json:"id,omitempty"`
	Label       string                   `json:"label,omitempty"`
	ObjectKey   string                   `json:"objectKey,omitempty"`
	Avatar      *ObjectAvatarScheme      `json:"avatar,omitempty"`
	ObjectType  *ObjectTypeScheme        `json:"objectType,omitempty"`
	Created     string                   `json:"created,omitempty"`
	Updated     string                   `json:"updated,omitempty"`
	HasAvatar   bool                     `json:"hasAvatar,omitempty"`
	Timestamp   int                      `json:"timestamp,omitempty"`
	Attributes  []*ObjectAttributeScheme `json:"attributes"`
	Links       *ObjectLinksScheme       `json:"_links,omitempty"`
}

type ObjectAvatarScheme struct {
	WorkspaceId string `json:"workspaceId,omitempty"`
	GlobalId    string `json:"globalId,omitempty"`
	Id          string `json:"id,omitempty"`
	AvatarUUID  string `json:"avatarUUID,omitempty"`
	Url16       string `json:"url16,omitempty"`
	Url48       string `json:"url48,omitempty"`
	Url72       string `json:"url72,omitempty"`
	Url144      string `json:"url144,omitempty"`
	Url288      string `json:"url288,omitempty"`
	ObjectId    string `json:"objectId,omitempty"`
}

type ObjectLinksScheme struct {
	Self string `json:"self,omitempty"`
}

type ObjectAttributeScheme struct {
	WorkspaceId           string                                 `json:"workspaceId,omitempty"`
	GlobalId              string                                 `json:"globalId,omitempty"`
	Id                    string                                 `json:"id,omitempty"`
	ObjectTypeAttribute   *ObjectTypeAttributeScheme             `json:"objectTypeAttribute,omitempty"`
	ObjectTypeAttributeId string                                 `json:"objectTypeAttributeId,omitempty"`
	ObjectAttributeValues []*ObjectTypeAssetAttributeValueScheme `json:"objectAttributeValues,omitempty"`
}

type ObjectTypeAttributeScheme struct {
	WorkspaceId             string                                       `json:"workspaceId,omitempty"`
	GlobalId                string                                       `json:"globalId,omitempty"`
	Id                      string                                       `json:"id,omitempty"`
	ObjectType              *ObjectTypeScheme                            `json:"objectType,omitempty"`
	Name                    string                                       `json:"name,omitempty"`
	Label                   bool                                         `json:"label,omitempty"`
	Type                    int                                          `json:"type,omitempty"`
	Description             string                                       `json:"description,omitempty"`
	DefaultType             *ObjectTypeAssetAttributeDefaultTypeScheme   `json:"defaultType,omitempty"`
	TypeValue               string                                       `json:"typeValue,omitempty"`
	TypeValueMulti          []string                                     `json:"typeValueMulti,omitempty"`
	AdditionalValue         string                                       `json:"additionalValue,omitempty"`
	ReferenceType           *ObjectTypeAssetAttributeReferenceTypeScheme `json:"referenceType,omitempty"`
	ReferenceObjectTypeId   string                                       `json:"referenceObjectTypeId,omitempty"`
	ReferenceObjectType     *ObjectTypeScheme                            `json:"referenceObjectType,omitempty"`
	Editable                bool                                         `json:"editable,omitempty"`
	System                  bool                                         `json:"system,omitempty"`
	Indexed                 bool                                         `json:"indexed,omitempty"`
	Sortable                bool                                         `json:"sortable,omitempty"`
	Summable                bool                                         `json:"summable,omitempty"`
	MinimumCardinality      int                                          `json:"minimumCardinality,omitempty"`
	MaximumCardinality      int                                          `json:"maximumCardinality,omitempty"`
	Suffix                  string                                       `json:"suffix,omitempty"`
	Removable               bool                                         `json:"removable,omitempty"`
	ObjectAttributeExists   bool                                         `json:"objectAttributeExists,omitempty"`
	Hidden                  bool                                         `json:"hidden,omitempty"`
	IncludeChildObjectTypes bool                                         `json:"includeChildObjectTypes,omitempty"`
	UniqueAttribute         bool                                         `json:"uniqueAttribute,omitempty"`
	RegexValidation         string                                       `json:"regexValidation,omitempty"`
	Iql                     string                                       `json:"iql,omitempty"`
	QlQuery                 string                                       `json:"qlQuery,omitempty"`
	Options                 string                                       `json:"options,omitempty"`
	Position                int                                          `json:"position,omitempty"`
}

type ObjectTypeAssetAttributeDefaultTypeScheme struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ObjectTypeAssetAttributeReferenceTypeScheme struct {
	WorkspaceId string `json:"workspaceId,omitempty"`
	GlobalId    string `json:"globalId,omitempty"`
	Name        string `json:"name,omitempty"`
}

type ObjectTypeAssetAttributeValueScheme struct {
	Value           string                                    `json:"value,omitempty"`
	DisplayValue    string                                    `json:"displayValue,omitempty"`
	SearchValue     string                                    `json:"searchValue,omitempty"`
	Group           *ObjectTypeAssetAttributeValueGroupScheme `json:"group,omitempty"`
	Status          *ObjectTypeAssetAttributeStatusScheme     `json:"status,omitempty"`
	AdditionalValue string                                    `json:"additionalValue,omitempty"`
}

type ObjectTypeAssetAttributeValueGroupScheme struct {
	AvatarUrl string `json:"avatarUrl,omitempty"`
	Name      string `json:"name,omitempty"`
}

type ObjectTypeAssetAttributeStatusScheme struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Category int    `json:"category,omitempty"`
}
