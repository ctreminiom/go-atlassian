// Package models provides the data structures used in the admin package.
package models

import "time"

// SCIMSchemasScheme represents a SCIM schema.
type SCIMSchemasScheme struct {
	TotalResults int               `json:"totalResults,omitempty"` // The total number of results.
	ItemsPerPage int               `json:"itemsPerPage,omitempty"` // The number of items per page.
	StartIndex   int               `json:"startIndex,omitempty"`   // The start index of the results.
	Schemas      []string          `json:"schemas,omitempty"`      // The schemas.
	Resources    []*ResourceScheme `json:"Resources,omitempty"`    // The resources.
}

// ResourceScheme represents a resource.
type ResourceScheme struct {
	ID          string              `json:"id,omitempty"`          // The ID of the resource.
	Name        string              `json:"name,omitempty"`        // The name of the resource.
	Description string              `json:"description,omitempty"` // The description of the resource.
	Attributes  []*AttributeScheme  `json:"attributes,omitempty"`  // The attributes of the resource.
	Meta        *ResourceMetaScheme `json:"meta,omitempty"`        // The metadata of the resource.
}

// ResourceMetaScheme represents the metadata of a resource.
type ResourceMetaScheme struct {
	ResourceType string `json:"resourceType,omitempty"` // The type of the resource.
	Location     string `json:"location,omitempty"`     // The location of the resource.
}

// AttributeScheme represents an attribute.
type AttributeScheme struct {
	Name          string                `json:"name,omitempty"`          // The name of the attribute.
	Type          string                `json:"type,omitempty"`          // The type of the attribute.
	MultiValued   bool                  `json:"multiValued,omitempty"`   // Whether the attribute is multi-valued.
	Description   string                `json:"description,omitempty"`   // The description of the attribute.
	Required      bool                  `json:"required,omitempty"`      // Whether the attribute is required.
	CaseExact     bool                  `json:"caseExact,omitempty"`     // Whether the attribute is case exact.
	Mutability    string                `json:"mutability,omitempty"`    // The mutability of the attribute.
	Returned      string                `json:"returned,omitempty"`      // When the attribute is returned.
	Uniqueness    string                `json:"uniqueness,omitempty"`    // The uniqueness of the attribute.
	SubAttributes []*SubAttributeScheme `json:"subAttributes,omitempty"` // The sub-attributes of the attribute.
}

// SubAttributeScheme represents a sub-attribute.
type SubAttributeScheme struct {
	Name        string `json:"name,omitempty"`        // The name of the sub-attribute.
	Type        string `json:"type,omitempty"`        // The type of the sub-attribute.
	MultiValued bool   `json:"multiValued,omitempty"` // Whether the sub-attribute is multi-valued.
	Description string `json:"description,omitempty"` // The description of the sub-attribute.
	Required    bool   `json:"required,omitempty"`    // Whether the sub-attribute is required.
	CaseExact   bool   `json:"caseExact,omitempty"`   // Whether the sub-attribute is case exact.
	Mutability  string `json:"mutability,omitempty"`  // The mutability of the sub-attribute.
	Returned    string `json:"returned,omitempty"`    // When the sub-attribute is returned.
	Uniqueness  string `json:"uniqueness,omitempty"`  // The uniqueness of the sub-attribute.
}

// SCIMSchemaScheme represents a SCIM schema.
type SCIMSchemaScheme struct {
	ID          string              `json:"id,omitempty"`          // The ID of the schema.
	Name        string              `json:"name,omitempty"`        // The name of the schema.
	Description string              `json:"description,omitempty"` // The description of the schema.
	Attributes  []*AttributeScheme  `json:"attributes,omitempty"`  // The attributes of the schema.
	Meta        *ResourceMetaScheme `json:"meta,omitempty"`        // The metadata of the schema.
}

// ServiceProviderConfigScheme represents a service provider configuration.
type ServiceProviderConfigScheme struct {
	Schemas []string `json:"schemas"` // The schemas.
	Patch   struct {
		Supported bool `json:"supported"` // Whether patching is supported.
	} `json:"patch"`
	Bulk struct {
		Supported      bool `json:"supported"`      // Whether bulk operations are supported.
		MaxOperations  int  `json:"maxOperations"`  // The maximum number of operations.
		MaxPayloadSize int  `json:"maxPayloadSize"` // The maximum payload size.
	} `json:"bulk"`
	Filter struct {
		MaxResults int  `json:"maxResults"` // The maximum number of results.
		Supported  bool `json:"supported"`  // Whether filtering is supported.
	} `json:"filter"`
	ChangePassword struct {
		Supported bool `json:"supported"` // Whether password change is supported.
	} `json:"changePassword"`
	Sort struct {
		Supported bool `json:"supported"` // Whether sorting is supported.
	} `json:"sort"`
	Etag struct {
		Supported bool `json:"supported"` // Whether ETag is supported.
	} `json:"etag"`
	AuthenticationSchemes []struct {
		Type        string `json:"type"`        // The type of the authentication scheme.
		Name        string `json:"name"`        // The name of the authentication scheme.
		Description string `json:"description"` // The description of the authentication scheme.
	} `json:"authenticationSchemes"`
	Meta struct {
		Location     string    `json:"location"`     // The location of the metadata.
		ResourceType string    `json:"resourceType"` // The type of the resource.
		LastModified time.Time `json:"lastModified"` // The last modified time.
		Created      time.Time `json:"created"`      // The creation time.
	} `json:"meta"`
}
