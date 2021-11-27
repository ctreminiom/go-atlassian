package models

import "time"

type AuditRecordPageScheme struct {
	Offset  int                  `json:"offset,omitempty"`
	Limit   int                  `json:"limit,omitempty"`
	Total   int                  `json:"total,omitempty"`
	Records []*AuditRecordScheme `json:"records,omitempty"`
}

type AuditRecordScheme struct {
	ID              int                                `json:"id,omitempty"`
	Summary         string                             `json:"summary,omitempty"`
	RemoteAddress   string                             `json:"remoteAddress,omitempty"`
	AuthorKey       string                             `json:"authorKey,omitempty"`
	Created         string                             `json:"created,omitempty"`
	Category        string                             `json:"category,omitempty"`
	EventSource     string                             `json:"eventSource,omitempty"`
	Description     string                             `json:"description,omitempty"`
	ObjectItem      *AuditRecordObjectItemScheme       `json:"objectItem,omitempty"`
	ChangedValues   []*AuditRecordChangedValueScheme   `json:"changedValues,omitempty"`
	AssociatedItems []*AuditRecordAssociatedItemScheme `json:"associatedItems,omitempty"`
}

type AuditRecordObjectItemScheme struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	TypeName   string `json:"typeName,omitempty"`
	ParentID   string `json:"parentId,omitempty"`
	ParentName string `json:"parentName,omitempty"`
}

type AuditRecordChangedValueScheme struct {
	FieldName   string `json:"fieldName,omitempty"`
	ChangedFrom string `json:"changedFrom,omitempty"`
	ChangedTo   string `json:"changedTo,omitempty"`
}

type AuditRecordAssociatedItemScheme struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	TypeName   string `json:"typeName,omitempty"`
	ParentID   string `json:"parentId,omitempty"`
	ParentName string `json:"parentName,omitempty"`
}

type AuditRecordGetOptions struct {
	Filter string
	From   time.Time
	To     time.Time
}
