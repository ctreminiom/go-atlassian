package models

type ContentVersionPageScheme struct {
	Results []*ContentVersionScheme `json:"results,omitempty"`
	Start   int                     `json:"start,omitempty"`
	Limit   int                     `json:"limit,omitempty"`
	Size    int                     `json:"size,omitempty"`
}

type ContentVersionScheme struct {
	By            *ContentUserScheme          `json:"by,omitempty"`
	Number        int                         `json:"number,omitempty"`
	When          string                      `json:"when,omitempty"`
	FriendlyWhen  string                      `json:"friendlyWhen,omitempty"`
	Message       string                      `json:"message,omitempty"`
	MinorEdit     bool                        `json:"minorEdit,omitempty"`
	Content       *ContentScheme              `json:"content,omitempty"`
	Collaborators *VersionCollaboratorsScheme `json:"collaborators,omitempty"`
	Expandable    struct {
		Content       string `json:"content"`
		Collaborators string `json:"collaborators"`
	} `json:"_expandable"`
	ContentTypeModified bool `json:"contentTypeModified"`
}

type VersionCollaboratorsScheme struct {
	Users    []*ContentUserScheme `json:"users,omitempty"`
	UserKeys []string             `json:"userKeys,omitempty"`
}

type ContentRestorePayloadScheme struct {
	OperationKey string                             `json:"operationKey,omitempty"`
	Params       *ContentRestoreParamsPayloadScheme `json:"params,omitempty"`
}

type ContentRestoreParamsPayloadScheme struct {
	VersionNumber int    `json:"versionNumber,omitempty"`
	Message       string `json:"message,omitempty"`
	RestoreTitle  bool   `json:"restoreTitle,omitempty"`
}
