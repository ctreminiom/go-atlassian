package models

// CheckPermissionScheme represents the scheme for checking permissions in Confluence.
// Subject is a pointer to a PermissionSubjectScheme which represents the subject of the permission.
// Operation is a string representing the operation of the permission.
type CheckPermissionScheme struct {
	Subject   *PermissionSubjectScheme `json:"subject,omitempty"`   // The subject of the permission.
	Operation string                   `json:"operation,omitempty"` // The operation of the permission.
}

// PermissionSubjectScheme represents the subject of a permission in Confluence.
// Identifier is a string representing the identifier of the subject.
// Type is a string representing the type of the subject.
type PermissionSubjectScheme struct {
	Identifier string `json:"identifier,omitempty"` // The identifier of the subject.
	Type       string `json:"type,omitempty"`       // The type of the subject.
}

// PermissionCheckResponseScheme represents the response scheme for checking permissions in Confluence.
// HasPermission is a boolean indicating if the permission is granted.
// Errors is a slice of pointers to PermissionCheckMessageScheme which represents the errors occurred during the permission check.
type PermissionCheckResponseScheme struct {
	HasPermission bool                            `json:"hasPermission"`    // Indicates if the permission is granted.
	Errors        []*PermissionCheckMessageScheme `json:"errors,omitempty"` // The errors occurred during the permission check.
}

// PermissionCheckMessageScheme represents a message scheme for checking permissions in Confluence.
// Translation is a string representing the translation of the message.
// Args is a slice of anonymous structs representing the arguments of the message.
type PermissionCheckMessageScheme struct {
	Translation string `json:"translation"` // The translation of the message.
	Args        []struct {
	} `json:"args"` // The arguments of the message.
}
