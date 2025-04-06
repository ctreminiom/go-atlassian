package models

// FolderScheme represents a folder in Confluence.
type FolderScheme struct {
	ID         string               `json:"id,omitempty"`         // The ID of the folder.
	Type       string               `json:"type,omitempty"`       // The type of the folder.
	Status     string               `json:"status,omitempty"`     // The status of the folder.
	Title      string               `json:"title,omitempty"`      // The title of the folder.
	ParentID   string               `json:"parentId,omitempty"`   // The ID of the parent of the folder.
	ParentType string               `json:"parentType,omitempty"` // The type of the parent of the folder.
	Position   int                  `json:"position,omitempty"`   // The position of the folder.
	AuthorID   string               `json:"authorId,omitempty"`   // The ID of the author of the folder.
	OwnerID    string               `json:"ownerId,omitempty"`    // The ID of the owner of the folder.
	CreatedAt  int                  `json:"createdAt,omitempty"`  // The timestamp of the creation of the folder.
	Version    *FolderVersionScheme `json:"version,omitempty"`    // The version of the folder.
}

// FolderVersionScheme represents the version of a folder in Confluence.
type FolderVersionScheme struct {
	CreatedAt string `json:"createdAt,omitempty"` // The timestamp of the creation of the version.
	Message   string `json:"message,omitempty"`   // The message of the version.
	Number    int    `json:"number,omitempty"`    // The number of the version.
	MinorEdit bool   `json:"minorEdit,omitempty"` // Indicates if the version is a minor edit.
	AuthorID  string `json:"authorId,omitempty"`  // The ID of the author of the version.
}

// FolderCreatePayloadScheme represents the payload for creating a folder in Confluence.
type FolderCreatePayloadScheme struct {
	SpaceID  string `json:"spaceId,omitempty"`  // The ID of the space of the folder.
	Title    string `json:"title,omitempty"`    // The title of the folder.
	ParentID string `json:"parentId,omitempty"` // The ID of the parent of the folder.
}
