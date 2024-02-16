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
	ID                string                    `json:"id,omitempty"`
	AffectedAttribute string                    `json:"affectedAttribute,omitempty"`
	OldValue          string                    `json:"oldValue,omitempty"`
	NewValue          string                    `json:"newValue,omitempty"`
	Type              int                       `json:"type,omitempty"`
	Created           string                    `json:"created,omitempty"`
	ObjectID          string                    `json:"objectId,omitempty"`
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
	ID          string                   `json:"id,omitempty"`
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
	ID          string `json:"id,omitempty"`
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
	ID                    string                                 `json:"id,omitempty"`
	ObjectTypeAttribute   *ObjectTypeAttributeScheme             `json:"objectTypeAttribute,omitempty"`
	ObjectTypeAttributeId string                                 `json:"objectTypeAttributeId,omitempty"`
	ObjectAttributeValues []*ObjectTypeAssetAttributeValueScheme `json:"objectAttributeValues,omitempty"`
}

type ObjectTypeAttributePayloadScheme struct {
	Name                    string   `json:"name,omitempty"`
	Label                   bool     `json:"label,omitempty"`
	Description             string   `json:"description,omitempty"`
	Type                    *int      `json:"type,omitempty"`
	DefaultTypeId		*int      `json:"defaultTypeId,omitempty"`
 	TypeValue               string   `json:"typeValue,omitempty"`
	TypeValueMulti          []string `json:"typeValueMulti,omitempty"`
	AdditionalValue         string   `json:"additionalValue,omitempty"`
	MinimumCardinality      *int      `json:"minimumCardinality,omitempty"`
	MaximumCardinality      *int      `json:"maximumCardinality,omitempty"`
	Suffix                  string   `json:"suffix,omitempty"`
	IncludeChildObjectTypes bool     `json:"includeChildObjectTypes,omitempty"`
	Hidden                  bool     `json:"hidden,omitempty"`
	UniqueAttribute         bool     `json:"uniqueAttribute,omitempty"`
	Summable                bool     `json:"summable,omitempty"`
	RegexValidation         string   `json:"regexValidation,omitempty"`
	QlQuery                 string   `json:"qlQuery,omitempty"`
	Iql                     string   `json:"iql,omitempty"`
	Options                 string   `json:"options,omitempty"`
}

type ObjectTypeAttributeScheme struct {
	WorkspaceId             string                                       `json:"workspaceId,omitempty"`
	GlobalId                string                                       `json:"globalId,omitempty"`
	ID                      string                                       `json:"id,omitempty"`
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
	ID   int    `json:"id,omitempty"`
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
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Category int    `json:"category,omitempty"`
}

type ObjectListScheme struct {
	ObjectEntries         []*ObjectScheme              `json:"objectEntries,omitempty"`
	ObjectTypeAttributes  []*ObjectTypeAttributeScheme `json:"objectTypeAttributes,omitempty"`
	ObjectTypeId          string                       `json:"objectTypeId,omitempty"`
	ObjectTypeIsInherited bool                         `json:"objectTypeIsInherited,omitempty"`
	AbstractObjectType    bool                         `json:"abstractObjectType,omitempty"`
	TotalFilterCount      int                          `json:"totalFilterCount,omitempty"`
	StartIndex            int                          `json:"startIndex,omitempty"`
	ToIndex               int                          `json:"toIndex,omitempty"`
	PageObjectSize        int                          `json:"pageObjectSize,omitempty"`
	PageNumber            int                          `json:"pageNumber,omitempty"`
	OrderByTypeAttrId     int                          `json:"orderByTypeAttrId,omitempty"`
	OrderWay              string                       `json:"orderWay,omitempty"`
	QlQuery               string                       `json:"qlQuery,omitempty"`
	QlQuerySearchResult   bool                         `json:"qlQuerySearchResult,omitempty"`
	Iql                   string                       `json:"iql,omitempty"`
	IqlSearchResult       bool                         `json:"iqlSearchResult,omitempty"`
	ConversionPossible    bool                         `json:"conversionPossible,omitempty"`
}

type ObjectListResultScheme struct {
	StartAt              int                          `json:"startAt,omitempty"`
	MaxResults           int                          `json:"maxResults,omitempty"`
	Total                int                          `json:"total,omitempty"`
	IsLast               bool                         `json:"isLast,omitempty"`
	Values               []*ObjectScheme              `json:"values,omitempty"`
	ObjectTypeAttributes []*ObjectTypeAttributeScheme `json:"objectTypeAttributes,omitempty"`
}

type ObjectSearchParamsScheme struct {

	// The AQL that will fetch the objects.
	//
	// The object type parameter will be appended implicitly to this AQL
	Query string `json:"qlQuery,omitempty"`

	// Required if qlQuery is not set.
	//
	// Deprecated. Use Query instead.
	Iql          string `json:"iql,omitempty"`
	ObjectTypeID string `json:"objectTypeId,omitempty"`

	// The requested page to be loaded for a paginated result.
	//
	// The default value is page = 1
	Page int `json:"page,omitempty"`

	// How many objects should be returned in the request.
	//
	// It is used with page attribute for pagination.
	ResultPerPage int `json:"resultsPerPage,omitempty"`

	// Which attribute should be used to order by.
	//
	// The preferred way is to use an order by in qlQuery and not pass this argument.
	OrderByTypeAttributeID int `json:"orderByTypeAttrId,omitempty"`

	// Sort objects in ascending order or descending order based on the attribute identified by orderByTypeAttrId.
	//
	// 1 mean ascending all other values mean descending.
	//
	// The preferred way is to not supply the asc parameter and use an order by in qlQuery instead.
	Asc int `json:"asc,omitempty"`

	// Identifies an object that should be included in the result.
	//
	// The page will be calculated accordingly to include the object specified in the result set
	ObjectID string `json:"objectId,omitempty"`

	ObjectSchemaID string `json:"objectSchemaId,omitempty"`

	// Should attribute values be included in the response.
	IncludeAttributes bool `json:"includeAttributes,omitempty"`

	// Identifies attributes to be displayed
	AttributesToDisplay *AttributesToDisplayScheme `json:"attributesToDisplay,omitempty"`
}

type AttributesToDisplayScheme struct {
	AttributesToDisplayIds []int `json:"attributesToDisplayIds,omitempty"`
}
