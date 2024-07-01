package models

// PropertyPageScheme represents a page of properties in Jira.
type PropertyPageScheme struct {
	Keys []*PropertyScheme `json:"keys,omitempty"` // The keys of the properties on the page.
}

// PropertyScheme represents a property in Jira.
type PropertyScheme struct {
	Self string `json:"self,omitempty"` // The URL of the property.
	Key  string `json:"key,omitempty"`  // The key of the property.
}

// EntityPropertyScheme represents an entity property in Jira.
type EntityPropertyScheme struct {
	Key   string      `json:"key"`   // The key of the entity property.
	Value interface{} `json:"value"` // The value of the entity property.
}
