package models

import "time"

// AuditRecordPageScheme represents a page of audit records in Jira.
type AuditRecordPageScheme struct {
	Offset  int                  `json:"offset,omitempty"`  // The offset of the page.
	Limit   int                  `json:"limit,omitempty"`   // The limit of the page.
	Total   int                  `json:"total,omitempty"`   // The total number of audit records.
	Records []*AuditRecordScheme `json:"records,omitempty"` // The audit records in the page.
}

// AuditRecordScheme represents an audit record in Jira.
type AuditRecordScheme struct {
	ID              int                                `json:"id,omitempty"`              // The ID of the audit record.
	Summary         string                             `json:"summary,omitempty"`         // The summary of the audit record.
	RemoteAddress   string                             `json:"remoteAddress,omitempty"`   // The remote address of the audit record.
	AuthorKey       string                             `json:"authorKey,omitempty"`       // The author key of the audit record.
	Created         string                             `json:"created,omitempty"`         // The creation time of the audit record.
	Category        string                             `json:"category,omitempty"`        // The category of the audit record.
	EventSource     string                             `json:"eventSource,omitempty"`     // The event source of the audit record.
	Description     string                             `json:"description,omitempty"`     // The description of the audit record.
	ObjectItem      *AuditRecordObjectItemScheme       `json:"objectItem,omitempty"`      // The object item of the audit record.
	ChangedValues   []*AuditRecordChangedValueScheme   `json:"changedValues,omitempty"`   // The changed values of the audit record.
	AssociatedItems []*AuditRecordAssociatedItemScheme `json:"associatedItems,omitempty"` // The associated items of the audit record.
}

// AuditRecordObjectItemScheme represents an object item in an audit record in Jira.
type AuditRecordObjectItemScheme struct {
	ID         string `json:"id,omitempty"`         // The ID of the object item.
	Name       string `json:"name,omitempty"`       // The name of the object item.
	TypeName   string `json:"typeName,omitempty"`   // The type name of the object item.
	ParentID   string `json:"parentId,omitempty"`   // The parent ID of the object item.
	ParentName string `json:"parentName,omitempty"` // The parent name of the object item.
}

// AuditRecordChangedValueScheme represents a changed value in an audit record in Jira.
type AuditRecordChangedValueScheme struct {
	FieldName   string `json:"fieldName,omitempty"`   // The field name of the changed value.
	ChangedFrom string `json:"changedFrom,omitempty"` // The previous value of the changed field.
	ChangedTo   string `json:"changedTo,omitempty"`   // The new value of the changed field.
}

// AuditRecordAssociatedItemScheme represents an associated item in an audit record in Jira.
type AuditRecordAssociatedItemScheme struct {
	ID         string `json:"id,omitempty"`         // The ID of the associated item.
	Name       string `json:"name,omitempty"`       // The name of the associated item.
	TypeName   string `json:"typeName,omitempty"`   // The type name of the associated item.
	ParentID   string `json:"parentId,omitempty"`   // The parent ID of the associated item.
	ParentName string `json:"parentName,omitempty"` // The parent name of the associated item.
}

// AuditRecordGetOptions represents the options for getting audit records in Jira.
type AuditRecordGetOptions struct {
	Filter string    // The filter for the audit records.
	From   time.Time // The start time for the audit records.
	To     time.Time // The end time for the audit records.
}
