package models

// SpacePermissionPayloadScheme represents the payload for a space permission in Confluence.
type SpacePermissionPayloadScheme struct {
	Subject   *PermissionSubjectScheme        `json:"subject,omitempty"`   // The subject of the permission.
	Operation *SpacePermissionOperationScheme `json:"operation,omitempty"` // The operation of the permission.
}

// SpacePermissionArrayPayloadScheme represents the payload for an array of space permissions in Confluence.
type SpacePermissionArrayPayloadScheme struct {
	Subject    *PermissionSubjectScheme       `json:"subject,omitempty"`    // The subject of the permissions.
	Operations []*SpaceOperationPayloadScheme `json:"operations,omitempty"` // The operations of the permissions.
}

// SpaceOperationPayloadScheme represents the payload for a space operation in Confluence.
type SpaceOperationPayloadScheme struct {
	Key    string `json:"key,omitempty"`    // The key of the operation.
	Target string `json:"target,omitempty"` // The target of the operation.
	Access bool   `json:"access,omitempty"` // Indicates if the operation has access.
}

// SpacePermissionOperationScheme represents an operation in a space permission in Confluence.
type SpacePermissionOperationScheme struct {
	Operation string `json:"operation,omitempty"` // The operation.
	Target    string `json:"target,omitempty"`    // The target of the operation.
	Key       string `json:"key,omitempty"`       // The key of the operation.
}

// SpacePermissionV2Scheme represents a version 2 space permission in Confluence.
type SpacePermissionV2Scheme struct {
	Subject   *PermissionSubjectScheme        `json:"subject,omitempty"`   // The subject of the permission.
	Operation *SpacePermissionOperationScheme `json:"operation,omitempty"` // The operation of the permission.
}
