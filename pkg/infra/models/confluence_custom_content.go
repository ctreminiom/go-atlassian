package models

type CustomContentOptionsScheme struct {
	IDs        []int
	SpaceIDs   []int
	Sort       string
	BodyFormat string
}

type CustomContentPageScheme struct {
	Results []*CustomContentScheme        `json:"results,omitempty"`
	Links   *CustomContentPageLinksScheme `json:"_links,omitempty"`
}

type CustomContentPageLinksScheme struct {
	Next string `json:"next,omitempty"`
}

type CustomContentScheme struct {
	ID              string                      `json:"id,omitempty"`
	Type            string                      `json:"type,omitempty"`
	Status          string                      `json:"status,omitempty"`
	Title           string                      `json:"title,omitempty"`
	SpaceID         string                      `json:"spaceId,omitempty"`
	PageID          string                      `json:"pageId,omitempty"`
	BlogPostID      string                      `json:"blogPostId,omitempty"`
	CustomContentID string                      `json:"customContentId,omitempty"`
	AuthorID        string                      `json:"authorId,omitempty"`
	CreatedAt       string                      `json:"createdAt,omitempty"`
	Version         *CustomContentVersionScheme `json:"version,omitempty"`
	Body            *CustomContentBodyScheme    `json:"body,omitempty"`
	Links           *CustomContentLinksScheme   `json:"_links,omitempty"`
}

type CustomContentVersionScheme struct {
	CreatedAt string `json:"createdAt,omitempty"`
	Message   string `json:"message,omitempty"`
	Number    int    `json:"number,omitempty"`
	MinorEdit bool   `json:"minorEdit,omitempty"`
	AuthorID  string `json:"authorId,omitempty"`
}

type CustomContentBodyScheme struct {
	Raw            *BodyTypeScheme `json:"raw,omitempty"`
	Storage        *BodyTypeScheme `json:"storage,omitempty"`
	AtlasDocFormat *BodyTypeScheme `json:"atlas_doc_format,omitempty"`
}

type CustomContentLinksScheme struct {
	Webui string `json:"webui,omitempty"`
}

type BodyTypeScheme struct {
	Representation string `json:"representation,omitempty"`
	Value          string `json:"value,omitempty"`
}

type CustomContentPayloadScheme struct {
	ID              string                             `json:"id,omitempty"`
	Type            string                             `json:"type,omitempty"`
	Status          string                             `json:"status,omitempty"`
	SpaceID         string                             `json:"spaceId,omitempty"`
	PageID          string                             `json:"pageId,omitempty"`
	BlogPostID      string                             `json:"blogPostId,omitempty"`
	CustomContentID string                             `json:"customContentId,omitempty"`
	Title           string                             `json:"title,omitempty"`
	Body            *BodyTypeScheme                    `json:"body,omitempty"`
	Version         *CustomContentPayloadVersionScheme `json:"version,omitempty"`
}

type CustomContentPayloadVersionScheme struct {
	Number  int    `json:"number,omitempty"`
	Message string `json:"message,omitempty"`
}
