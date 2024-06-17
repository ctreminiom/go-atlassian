package models

// CustomContentOptionsScheme represents the options for custom content in Confluence.
type CustomContentOptionsScheme struct {
	IDs        []int  `json:"ids,omitempty"`        // The IDs of the custom content.
	SpaceIDs   []int  `json:"spaceIDs,omitempty"`   // The IDs of the spaces of the custom content.
	Sort       string `json:"sort,omitempty"`       // The sort order of the custom content.
	BodyFormat string `json:"bodyFormat,omitempty"` // The body format of the custom content.
}

// CustomContentPageScheme represents a page of custom content in Confluence.
type CustomContentPageScheme struct {
	Results []*CustomContentScheme        `json:"results,omitempty"` // The custom content in the page.
	Links   *CustomContentPageLinksScheme `json:"_links,omitempty"`  // The links of the page.
}

// CustomContentPageLinksScheme represents the links of a page of custom content in Confluence.
type CustomContentPageLinksScheme struct {
	Next string `json:"next,omitempty"` // The link to the next page.
}

// CustomContentScheme represents custom content in Confluence.
type CustomContentScheme struct {
	ID              string                      `json:"id,omitempty"`              // The ID of the custom content.
	Type            string                      `json:"type,omitempty"`            // The type of the custom content.
	Status          string                      `json:"status,omitempty"`          // The status of the custom content.
	Title           string                      `json:"title,omitempty"`           // The title of the custom content.
	SpaceID         string                      `json:"spaceId,omitempty"`         // The ID of the space of the custom content.
	PageID          string                      `json:"pageId,omitempty"`          // The ID of the page of the custom content.
	BlogPostID      string                      `json:"blogPostId,omitempty"`      // The ID of the blog post of the custom content.
	CustomContentID string                      `json:"customContentId,omitempty"` // The ID of the custom content.
	AuthorID        string                      `json:"authorId,omitempty"`        // The ID of the author of the custom content.
	CreatedAt       string                      `json:"createdAt,omitempty"`       // The timestamp of the creation of the custom content.
	Version         *CustomContentVersionScheme `json:"version,omitempty"`         // The version of the custom content.
	Body            *CustomContentBodyScheme    `json:"body,omitempty"`            // The body of the custom content.
	Links           *CustomContentLinksScheme   `json:"_links,omitempty"`          // The links of the custom content.
}

// CustomContentVersionScheme represents the version of custom content in Confluence.
type CustomContentVersionScheme struct {
	CreatedAt string `json:"createdAt,omitempty"` // The timestamp of the creation of the version.
	Message   string `json:"message,omitempty"`   // The message of the version.
	Number    int    `json:"number,omitempty"`    // The number of the version.
	MinorEdit bool   `json:"minorEdit,omitempty"` // Indicates if the version is a minor edit.
	AuthorID  string `json:"authorId,omitempty"`  // The ID of the author of the version.
}

// CustomContentBodyScheme represents the body of custom content in Confluence.
type CustomContentBodyScheme struct {
	Raw            *BodyTypeScheme `json:"raw,omitempty"`              // The raw body.
	Storage        *BodyTypeScheme `json:"storage,omitempty"`          // The storage body.
	AtlasDocFormat *BodyTypeScheme `json:"atlas_doc_format,omitempty"` // The Atlas doc format body.
}

// CustomContentLinksScheme represents the links of custom content in Confluence.
type CustomContentLinksScheme struct {
	Webui string `json:"webui,omitempty"` // The web UI link.
}

// BodyTypeScheme represents a type of body in Confluence.
type BodyTypeScheme struct {
	Representation string `json:"representation,omitempty"` // The representation of the body.
	Value          string `json:"value,omitempty"`          // The value of the body.
}

// CustomContentPayloadScheme represents the payload for custom content in Confluence.
type CustomContentPayloadScheme struct {
	ID              string                             `json:"id,omitempty"`              // The ID of the custom content.
	Type            string                             `json:"type,omitempty"`            // The type of the custom content.
	Status          string                             `json:"status,omitempty"`          // The status of the custom content.
	SpaceID         string                             `json:"spaceId,omitempty"`         // The ID of the space of the custom content.
	PageID          string                             `json:"pageId,omitempty"`          // The ID of the page of the custom content.
	BlogPostID      string                             `json:"blogPostId,omitempty"`      // The ID of the blog post of the custom content.
	CustomContentID string                             `json:"customContentId,omitempty"` // The ID of the custom content.
	Title           string                             `json:"title,omitempty"`           // The title of the custom content.
	Body            *BodyTypeScheme                    `json:"body,omitempty"`            // The body of the custom content.
	Version         *CustomContentPayloadVersionScheme `json:"version,omitempty"`         // The version of the custom content.
}

// CustomContentPayloadVersionScheme represents the version of the payload for custom content in Confluence.
type CustomContentPayloadVersionScheme struct {
	Number  int    `json:"number,omitempty"`  // The number of the version.
	Message string `json:"message,omitempty"` // The message of the version.
}
