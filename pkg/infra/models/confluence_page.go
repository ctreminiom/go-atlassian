package models

type PageChunkScheme struct {
	Results []*PageScheme         `json:"results,omitempty"`
	Links   *PageChunkLinksScheme `json:"_links,omitempty"`
}

type PageChunkLinksScheme struct {
	Next string `json:"next,omitempty"`
}

type PageScheme struct {
	ID        int                `json:"id,omitempty"`
	Status    string             `json:"status,omitempty"`
	Title     string             `json:"title,omitempty"`
	SpaceID   int                `json:"spaceId,omitempty"`
	ParentID  int                `json:"parentId,omitempty"`
	AuthorID  string             `json:"authorId,omitempty"`
	CreatedAt string             `json:"createdAt,omitempty"`
	Version   *PageVersionScheme `json:"version,omitempty"`
	Body      *PageBodyScheme    `json:"body,omitempty"`
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
