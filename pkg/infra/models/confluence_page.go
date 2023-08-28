package models

type ChildPageChunkLinksScheme struct {
	Next string `json:"next,omitempty"`
}

type ChildPageChunkScheme struct {
	Results []*ChildPageScheme         `json:"results,omitempty"`
	Links   *ChildPageChunkLinksScheme `json:"_links,omitempty"`
}

type ChildPageScheme struct {
	ID            string `json:"id,omitempty"`
	Status        string `json:"status,omitempty"`
	Title         string `json:"title,omitempty"`
	SpaceID       string `json:"spaceId,omitempty"`
	ChildPosition int    `json:"childPosition,omitempty"`
}

type PageOptionsScheme struct {
	PageIDs    []int
	SpaceIDs   []int
	Sort       string
	Status     []string
	Title      string
	BodyFormat string
}

type PageChunkScheme struct {
	Results []*PageScheme         `json:"results,omitempty"`
	Links   *PageChunkLinksScheme `json:"_links,omitempty"`
}

type PageChunkLinksScheme struct {
	Next string `json:"next,omitempty"`
}

type PageScheme struct {
	ID         string             `json:"id,omitempty"`
	Status     string             `json:"status,omitempty"`
	Title      string             `json:"title,omitempty"`
	SpaceID    string             `json:"spaceId,omitempty"`
	ParentID   string             `json:"parentId,omitempty"`
	AuthorID   string             `json:"authorId,omitempty"`
	CreatedAt  string             `json:"createdAt,omitempty"`
	ParentType string             `json:"parentType,omitempty"`
	Position   int                `json:"position,omitempty"`
	Version    *PageVersionScheme `json:"version,omitempty"`
	Body       *PageBodyScheme    `json:"body,omitempty"`
}

type PageVersionScheme struct {
	CreatedAt string `json:"createdAt,omitempty"`
	Message   string `json:"message,omitempty"`
	Number    int    `json:"number,omitempty"`
	MinorEdit bool   `json:"minorEdit,omitempty"`
	AuthorID  string `json:"authorId,omitempty"`
}

type PageBodyScheme struct {
	Storage        *PageBodyRepresentationScheme `json:"storage,omitempty"`
	AtlasDocFormat *PageBodyRepresentationScheme `json:"atlas_doc_format,omitempty"`
}

type PageCreatePayloadScheme struct {
	SpaceID int                           `json:"spaceId,omitempty"`
	Status  string                        `json:"status,omitempty"`
	Title   string                        `json:"title,omitempty"`
	Body    *PageBodyRepresentationScheme `json:"body,omitempty"`
}

type PageBodyRepresentationScheme struct {
	Representation string `json:"representation,omitempty"`
	Value          string `json:"value,omitempty"`
}

type PageUpdatePayloadScheme struct {
	ID      int                             `json:"id,omitempty"`
	Status  string                          `json:"status,omitempty"`
	Title   string                          `json:"title,omitempty"`
	SpaceID int                             `json:"spaceId,omitempty"`
	Body    *PageBodyRepresentationScheme   `json:"body,omitempty"`
	Version *PageUpdatePayloadVersionScheme `json:"version,omitempty"`
}

type PageUpdatePayloadVersionScheme struct {
	Number  int    `json:"number,omitempty"`
	Message string `json:"message,omitempty"`
}
