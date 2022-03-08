package models

type SpacePermissionPayloadScheme struct {
	Subject   *PermissionSubjectScheme        `json:"subject,omitempty"`
	Operation *SpacePermissionOperationScheme `json:"operation,omitempty"`
}

type SpacePermissionArrayPayloadScheme struct {
	Subject    *PermissionSubjectScheme       `json:"subject,omitempty"`
	Operations []*SpaceOperationPayloadScheme `json:"operations,omitempty"`
}

type SpaceOperationPayloadScheme struct {
	Key    string `json:"key,omitempty"`
	Target string `json:"target,omitempty"`
	Access bool   `json:"access,omitempty"`
}

type SpacePermissionOperationScheme struct {
	Operation string `json:"operation,omitempty"`
	Target    string `json:"target,omitempty"`
	Key       string `json:"key,omitempty"`
}

type SpacePermissionV2Scheme struct {
	Subject   *PermissionSubjectScheme        `json:"subject,omitempty"`
	Operation *SpacePermissionOperationScheme `json:"operation,omitempty"`
}
