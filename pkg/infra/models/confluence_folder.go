package models

import "time"

// FolderScheme represents a folder in Confluence.
type FolderScheme struct {
	ID         string                 `json:"id,omitempty"`         // The ID of the folder.
	Type       string                 `json:"type,omitempty"`       // The type of the folder.
	Status     string                 `json:"status,omitempty"`     // The status of the folder.
	Title      string                 `json:"title,omitempty"`      // The title of the folder.
	ParentID   string                 `json:"parentId,omitempty"`   // The ID of the parent of the folder.
	ParentType string                 `json:"parentType,omitempty"` // The type of the parent of the folder.
	Position   int                    `json:"position,omitempty"`   // The position of the folder.
	AuthorID   string                 `json:"authorId,omitempty"`   // The ID of the author of the folder.
	OwnerID    string                 `json:"ownerId,omitempty"`    // The ID of the owner of the folder.
	CreatedAt  *time.Time             `json:"createdAt,omitempty"`  // The creation time of the folder.
	Version    *FolderVersionScheme   `json:"version,omitempty"`    // The version information of the folder.
	Links      *FolderLinksScheme     `json:"_links,omitempty"`     // The links of the folder.
}

// FolderVersionScheme represents the version information of a folder.
type FolderVersionScheme struct {
	Number    int        `json:"number,omitempty"`    // The version number.
	CreatedAt *time.Time `json:"createdAt,omitempty"` // The creation time of the version.
}

// FolderLinksScheme represents the links of a folder.
type FolderLinksScheme struct {
	Self string `json:"self,omitempty"` // The self link of the folder.
}

// FolderChunkScheme represents a chunk of folders in Confluence.
type FolderChunkScheme struct {
	Results []*FolderScheme         `json:"results,omitempty"` // The folders in the chunk.
	Links   *FolderChunkLinksScheme `json:"_links,omitempty"`  // The links of the chunk.
}

// FolderChunkLinksScheme represents the links of a chunk of folders.
type FolderChunkLinksScheme struct {
	Next string `json:"next,omitempty"` // The link to the next chunk of folders.
}

// FolderCreatePayloadScheme represents the payload for creating a folder.
type FolderCreatePayloadScheme struct {
	SpaceID  string `json:"spaceId"`           // The ID of the space where the folder will be created.
	Title    string `json:"title"`             // The title of the folder.
	ParentID string `json:"parentId,omitempty"` // The ID of the parent folder (optional).
}

// FolderUpdatePayloadScheme represents the payload for updating a folder.
type FolderUpdatePayloadScheme struct {
	Title    string                           `json:"title,omitempty"`    // The updated title of the folder.
	ParentID string                           `json:"parentId,omitempty"` // The updated parent ID of the folder.
	Version  *FolderUpdatePayloadVersionScheme `json:"version"`            // The version information for the update.
}

// FolderUpdatePayloadVersionScheme represents the version information for updating a folder.
type FolderUpdatePayloadVersionScheme struct {
	Number int `json:"number"` // The current version number that is being updated.
}

// FolderOptionsScheme represents the options for filtering folders.
type FolderOptionsScheme struct {
	SpaceIDs []int  `json:"spaceIDs,omitempty"` // The IDs of the spaces to filter folders by.
	Sort     string `json:"sort,omitempty"`     // The sort order of the folders.
	ParentID string `json:"parentId,omitempty"` // The parent ID to filter folders by.
}