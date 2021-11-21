package models

import "time"

type SCIMSchemasScheme struct {
	TotalResults int               `json:"totalResults,omitempty"`
	ItemsPerPage int               `json:"itemsPerPage,omitempty"`
	StartIndex   int               `json:"startIndex,omitempty"`
	Schemas      []string          `json:"schemas,omitempty"`
	Resources    []*ResourceScheme `json:"Resources,omitempty"`
}

type ResourceScheme struct {
	ID          string              `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Attributes  []*AttributeScheme  `json:"attributes,omitempty"`
	Meta        *ResourceMetaScheme `json:"meta,omitempty"`
}

type ResourceMetaScheme struct {
	ResourceType string `json:"resourceType,omitempty"`
	Location     string `json:"location,omitempty"`
}

type AttributeScheme struct {
	Name          string                `json:"name,omitempty"`
	Type          string                `json:"type,omitempty"`
	MultiValued   bool                  `json:"multiValued,omitempty"`
	Description   string                `json:"description,omitempty"`
	Required      bool                  `json:"required,omitempty"`
	CaseExact     bool                  `json:"caseExact,omitempty"`
	Mutability    string                `json:"mutability,omitempty"`
	Returned      string                `json:"returned,omitempty"`
	Uniqueness    string                `json:"uniqueness,omitempty"`
	SubAttributes []*SubAttributeScheme `json:"subAttributes,omitempty"`
}

type SubAttributeScheme struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	MultiValued bool   `json:"multiValued,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	CaseExact   bool   `json:"caseExact,omitempty"`
	Mutability  string `json:"mutability,omitempty"`
	Returned    string `json:"returned,omitempty"`
	Uniqueness  string `json:"uniqueness,omitempty"`
}

type SCIMSchemaScheme struct {
	ID          string              `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Attributes  []*AttributeScheme  `json:"attributes,omitempty"`
	Meta        *ResourceMetaScheme `json:"meta,omitempty"`
}

type ServiceProviderConfigScheme struct {
	Schemas []string `json:"schemas"`
	Patch   struct {
		Supported bool `json:"supported"`
	} `json:"patch"`
	Bulk struct {
		Supported      bool `json:"supported"`
		MaxOperations  int  `json:"maxOperations"`
		MaxPayloadSize int  `json:"maxPayloadSize"`
	} `json:"bulk"`
	Filter struct {
		MaxResults int  `json:"maxResults"`
		Supported  bool `json:"supported"`
	} `json:"filter"`
	ChangePassword struct {
		Supported bool `json:"supported"`
	} `json:"changePassword"`
	Sort struct {
		Supported bool `json:"supported"`
	} `json:"sort"`
	Etag struct {
		Supported bool `json:"supported"`
	} `json:"etag"`
	AuthenticationSchemes []struct {
		Type        string `json:"type"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"authenticationSchemes"`
	Meta struct {
		Location     string    `json:"location"`
		ResourceType string    `json:"resourceType"`
		LastModified time.Time `json:"lastModified"`
		Created      time.Time `json:"created"`
	} `json:"meta"`
}
