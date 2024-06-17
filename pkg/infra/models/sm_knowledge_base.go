package models

// ArticlePageScheme represents a page of articles in a system.
type ArticlePageScheme struct {
	Size       int                    `json:"size,omitempty"`       // The number of articles on the page.
	Start      int                    `json:"start,omitempty"`      // The index of the first article on the page.
	Limit      int                    `json:"limit,omitempty"`      // The maximum number of articles that can be on the page.
	IsLastPage bool                   `json:"isLastPage,omitempty"` // Indicates if this is the last page of articles.
	Values     []*ArticleScheme       `json:"values,omitempty"`     // The articles on the page.
	Expands    []string               `json:"_expands,omitempty"`   // Additional data related to the articles.
	Links      *ArticlePageLinkScheme `json:"_links,omitempty"`     // Links related to the page of articles.
}

// ArticlePageLinkScheme represents links related to a page of articles.
type ArticlePageLinkScheme struct {
	Self    string `json:"self,omitempty"`    // The URL of the page itself.
	Base    string `json:"base,omitempty"`    // The base URL for the links.
	Context string `json:"context,omitempty"` // The context for the links.
	Next    string `json:"next,omitempty"`    // The URL for the next page of articles.
	Prev    string `json:"prev,omitempty"`    // The URL for the previous page of articles.
}

// ArticleScheme represents an article in a system.
type ArticleScheme struct {
	Title   string                `json:"title,omitempty"`   // The title of the article.
	Excerpt string                `json:"excerpt,omitempty"` // An excerpt from the article.
	Source  *ArticleSourceScheme  `json:"source,omitempty"`  // The source of the article.
	Content *ArticleContentScheme `json:"content,omitempty"` // The content of the article.
}

// ArticleSourceScheme represents the source of an article.
type ArticleSourceScheme struct {
	Type string `json:"type,omitempty"` // The type of the source.
}

// ArticleContentScheme represents the content of an article.
type ArticleContentScheme struct {
	IframeSrc string `json:"iframeSrc,omitempty"` // The source of the iframe for the article content.
}
