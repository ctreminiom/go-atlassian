package models

// CheckPermissionScheme represents the scheme for checking permissions in Confluence.
type CheckPermissionScheme struct {
	Subject   *PermissionSubjectScheme `json:"subject,omitempty"`   // The subject of the permission.
	Operation string                   `json:"operation,omitempty"` // The operation of the permission.
}

// PermissionSubjectScheme represents the subject of a permission in Confluence.
type PermissionSubjectScheme struct {
	Identifier string `json:"identifier,omitempty"` // The identifier of the subject.
	Type       string `json:"type,omitempty"`       // The type of the subject.
}

// PermissionCheckResponseScheme represents the response scheme for checking permissions in Confluence.
type PermissionCheckResponseScheme struct {
	HasPermission bool                            `json:"hasPermission"`    // Indicates if the permission is granted.
	Errors        []*PermissionCheckMessageScheme `json:"errors,omitempty"` // The errors occurred during the permission check.
}

// PermissionCheckMessageScheme represents a message scheme for checking permissions in Confluence.
type PermissionCheckMessageScheme struct {
	Translation string `json:"translation"` // The translation of the message.
	Args        []struct {
	} `json:"args"` // The arguments of the message.
}
