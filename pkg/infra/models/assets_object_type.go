package models

type ObjectTypePositionPayloadScheme struct {
	ToObjectTypeId string `json:"toObjectTypeId,omitempty"`
	Position       int    `json:"position,omitempty"`
}

type ObjectTypeScheme struct {
	WorkspaceId               string      `json:"workspaceId,omitempty"`
	GlobalId                  string      `json:"globalId,omitempty"`
	Id                        string      `json:"id,omitempty"`
	Name                      string      `json:"name,omitempty"`
	Description               string      `json:"description,omitempty"`
	Icon                      *IconScheme `json:"icon,omitempty"`
	Position                  int         `json:"position,omitempty"`
	Created                   string      `json:"created,omitempty"`
	Updated                   string      `json:"updated,omitempty"`
	ObjectCount               int         `json:"objectCount,omitempty"`
	ParentObjectTypeId        int         `json:"parentObjectTypeId,omitempty"`
	ObjectSchemaId            string      `json:"objectSchemaId,omitempty"`
	Inherited                 bool        `json:"inherited,omitempty"`
	AbstractObjectType        bool        `json:"abstractObjectType,omitempty"`
	ParentObjectTypeInherited bool        `json:"parentObjectTypeInherited,omitempty"`
}

type ObjectTypePayloadScheme struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	IconId             string `json:"iconId,omitempty"`
	ObjectSchemaId     string `json:"objectSchemaId,omitempty"`
	ParentObjectTypeId string `json:"parentObjectTypeId,omitempty"`
	Inherited          bool   `json:"inherited,omitempty"`
	AbstractObjectType bool   `json:"abstractObjectType,omitempty"`
}

type ObjectTypeAttributesParamsScheme struct {
	OnlyValueEditable       bool
	OrderByName             bool
	Query                   string
	IncludeValuesExist      bool
	ExcludeParentAttributes bool
	IncludeChildren         bool
	OrderByRequired         bool
}
