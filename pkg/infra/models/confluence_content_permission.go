package models

type CheckPermissionScheme struct {
	Subject   *PermissionSubjectScheme `json:"subject,omitempty"`
	Operation string                   `json:"operation,omitempty"`
}

type PermissionSubjectScheme struct {
	Identifier string `json:"identifier,omitempty"`
	Type       string `json:"type,omitempty"`
}

type PermissionCheckResponseScheme struct {
	HasPermission bool                            `json:"hasPermission"`
	Errors        []*PermissionCheckMessageScheme `json:"errors,omitempty"`
}

type PermissionCheckMessageScheme struct {
	Translation string `json:"translation"`
	Args        []struct {
	} `json:"args"`
}
