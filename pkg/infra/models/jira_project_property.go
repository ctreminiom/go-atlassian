package models

type ProjectPropertyPageScheme struct {
	Keys []*ProjectPropertyScheme `json:"keys,omitempty"`
}

type ProjectPropertyScheme struct {
	Self string `json:"self,omitempty"`
	Key  string `json:"key,omitempty"`
}

type EntityPropertyScheme struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
