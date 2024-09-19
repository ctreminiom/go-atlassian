package models

// ChildPageChunkLinksScheme represents the links of a chunk of child pages in Confluence.
type ChildPageChunkLinksScheme struct {
	Next string `json:"next,omitempty"` // The link to the next chunk of child pages.
}

// ChildPageChunkScheme represents a chunk of child pages in Confluence.
type ChildPageChunkScheme struct {
	Results []*ChildPageScheme         `json:"results,omitempty"` // The child pages in the chunk.
	Links   *ChildPageChunkLinksScheme `json:"_links,omitempty"`  // The links of the chunk.
}

// ChildPageScheme represents a child page in Confluence.
type ChildPageScheme struct {
	ID            string `json:"id,omitempty"`            // The ID of the child page.
	Status        string `json:"status,omitempty"`        // The status of the child page.
	Title         string `json:"title,omitempty"`         // The title of the child page.
	SpaceID       string `json:"spaceId,omitempty"`       // The ID of the space of the child page.
	ChildPosition int    `json:"childPosition,omitempty"` // The position of the child page.
}

// PageOptionsScheme represents the options for a page in Confluence.
type PageOptionsScheme struct {
	PageIDs    []int    `json:"pageIDs,omitempty"`    // The IDs of the pages.
	SpaceIDs   []int    `json:"spaceIDs,omitempty"`   // The IDs of the spaces of the pages.
	Sort       string   `json:"sort,omitempty"`       // The sort order of the pages.
	Status     []string `json:"status,omitempty"`     // The statuses of the pages.
	Title      string   `json:"title,omitempty"`      // The title of the pages.
	BodyFormat string   `json:"bodyFormat,omitempty"` // The body format of the pages.
}

// PageChunkScheme represents a chunk of pages in Confluence.
type PageChunkScheme struct {
	Results []*PageScheme         `json:"results,omitempty"` // The pages in the chunk.
	Links   *PageChunkLinksScheme `json:"_links,omitempty"`  // The links of the chunk.
}

// PageChunkLinksScheme represents the links of a chunk of pages in Confluence.
type PageChunkLinksScheme struct {
	Next string `json:"next,omitempty"` // The link to the next chunk of pages.
}

// PageScheme represents a page in Confluence.
type PageScheme struct {
	ID         string             `json:"id,omitempty"`         // The ID of the page.
	Status     string             `json:"status,omitempty"`     // The status of the page.
	Title      string             `json:"title,omitempty"`      // The title of the page.
	SpaceID    string             `json:"spaceId,omitempty"`    // The ID of the space of the page.
	ParentID   string             `json:"parentId,omitempty"`   // The ID of the parent of the page.
	AuthorID   string             `json:"authorId,omitempty"`   // The ID of the author of the page.
	CreatedAt  string             `json:"createdAt,omitempty"`  // The timestamp of the creation of the page.
	ParentType string             `json:"parentType,omitempty"` // The type of the parent of the page.
	Position   int                `json:"position,omitempty"`   // The position of the page.
	Version    *PageVersionScheme `json:"version,omitempty"`    // The version of the page.
	Body       *PageBodyScheme    `json:"body,omitempty"`       // The body of the page.
}

// PageVersionScheme represents the version of a page in Confluence.
type PageVersionScheme struct {
	CreatedAt string `json:"createdAt,omitempty"` // The timestamp of the creation of the version.
	Message   string `json:"message,omitempty"`   // The message of the version.
	Number    int    `json:"number,omitempty"`    // The number of the version.
	MinorEdit bool   `json:"minorEdit,omitempty"` // Indicates if the version is a minor edit.
	AuthorID  string `json:"authorId,omitempty"`  // The ID of the author of the version.
}

// PageBodyScheme represents the body of a page in Confluence.
type PageBodyScheme struct {
	Storage        *PageBodyRepresentationScheme `json:"storage,omitempty"`          // The storage body.
	AtlasDocFormat *PageBodyRepresentationScheme `json:"atlas_doc_format,omitempty"` // The Atlas doc format body.
}

// PageCreatePayloadScheme represents the payload for creating a page in Confluence.
type PageCreatePayloadScheme struct {
	SpaceID  string                        `json:"spaceId,omitempty"`  // The ID of the space of the page.
	Status   string                        `json:"status,omitempty"`   // The status of the page.
	Title    string                        `json:"title,omitempty"`    // The title of the page.
	ParentID string                        `json:"parentId,omitempty"` // The ID of the parent of the page.
	Body     *PageBodyRepresentationScheme `json:"body,omitempty"`     // The body of the page.
}

// PageBodyRepresentationScheme represents a representation of a body in Confluence.
type PageBodyRepresentationScheme struct {
	Representation string `json:"representation,omitempty"` // The representation of the body.
	Value          string `json:"value,omitempty"`          // The value of the body.
}

// PageUpdatePayloadScheme represents the payload for updating a page in Confluence.
type PageUpdatePayloadScheme struct {
	ID       string                          `json:"id,omitempty"`       // The ID of the page.
	Status   string                          `json:"status,omitempty"`   // The status of the page.
	Title    string                          `json:"title,omitempty"`    // The title of the page.
	SpaceID  string                          `json:"spaceId,omitempty"`  // The ID of the space of the page.
	ParentID string                          `json:"parentId,omitempty"` // The ID of the parent of the page.
	OwnerID  string                          `json:"ownerId,omitempty"`  // The ID of the owner of the page.
	Body     *PageBodyRepresentationScheme   `json:"body,omitempty"`     // The body of the page.
	Version  *PageUpdatePayloadVersionScheme `json:"version,omitempty"`  // The version of the page.
}

// PageUpdatePayloadVersionScheme represents the version of the payload for updating a page in Confluence.
type PageUpdatePayloadVersionScheme struct {
	Number  int    `json:"number,omitempty"`  // The number of the version.
	Message string `json:"message,omitempty"` // The message of the version.
}
