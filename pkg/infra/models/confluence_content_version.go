package models

// ContentVersionPageScheme represents a page of content versions in Confluence.
type ContentVersionPageScheme struct {
	Results []*ContentVersionScheme `json:"results,omitempty"` // The content versions in the page.
	Start   int                     `json:"start,omitempty"`   // The start index of the content versions in the page.
	Limit   int                     `json:"limit,omitempty"`   // The limit of the content versions in the page.
	Size    int                     `json:"size,omitempty"`    // The size of the content versions in the page.
}

// ContentVersionScheme represents a content version in Confluence.
type ContentVersionScheme struct {
	By            *ContentUserScheme          `json:"by,omitempty"`            // The user who made the version.
	Number        int                         `json:"number,omitempty"`        // The number of the version.
	When          string                      `json:"when,omitempty"`          // The timestamp of the version.
	FriendlyWhen  string                      `json:"friendlyWhen,omitempty"`  // The friendly timestamp of the version.
	Message       string                      `json:"message,omitempty"`       // The message of the version.
	MinorEdit     bool                        `json:"minorEdit,omitempty"`     // Indicates if the version is a minor edit.
	Content       *ContentScheme              `json:"content,omitempty"`       // The content of the version.
	Collaborators *VersionCollaboratorsScheme `json:"collaborators,omitempty"` // The collaborators of the version.
	Expandable    struct {
		Content       string `json:"content"`       // The content of the version.
		Collaborators string `json:"collaborators"` // The collaborators of the version.
	} `json:"_expandable"` // The expandable fields of the version.
	ContentTypeModified bool `json:"contentTypeModified"` // Indicates if the content type is modified in the version.
}

// VersionCollaboratorsScheme represents the collaborators of a version in Confluence.
type VersionCollaboratorsScheme struct {
	Users    []*ContentUserScheme `json:"users,omitempty"`    // The users who are collaborators.
	UserKeys []string             `json:"userKeys,omitempty"` // The keys of the users who are collaborators.
}

// ContentRestorePayloadScheme represents the payload for restoring a content in Confluence.
type ContentRestorePayloadScheme struct {
	OperationKey string                             `json:"operationKey,omitempty"` // The key of the operation for restoring the content.
	Params       *ContentRestoreParamsPayloadScheme `json:"params,omitempty"`       // The parameters for restoring the content.
}

// ContentRestoreParamsPayloadScheme represents the parameters for restoring a content in Confluence.
type ContentRestoreParamsPayloadScheme struct {
	VersionNumber int    `json:"versionNumber,omitempty"` // The number of the version to be restored.
	Message       string `json:"message,omitempty"`       // The message for restoring the content.
	RestoreTitle  bool   `json:"restoreTitle,omitempty"`  // Indicates if the title should be restored.
}
