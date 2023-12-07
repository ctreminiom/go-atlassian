package models

type PropertyPageScheme struct {
	Keys []*PropertyScheme `json:"keys,omitempty"`
}

type PropertyScheme struct {
	Self string `json:"self,omitempty"`
	Key  string `json:"key,omitempty"`
}

type EntityPropertyScheme struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
