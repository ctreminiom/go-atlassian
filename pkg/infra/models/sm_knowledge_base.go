package models

type ArticlePageScheme struct {
	Size       int                    `json:"size,omitempty"`
	Start      int                    `json:"start,omitempty"`
	Limit      int                    `json:"limit,omitempty"`
	IsLastPage bool                   `json:"isLastPage,omitempty"`
	Values     []*ArticleScheme       `json:"values,omitempty"`
	Expands    []string               `json:"_expands,omitempty"`
	Links      *ArticlePageLinkScheme `json:"_links,omitempty"`
}

type ArticlePageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type ArticleScheme struct {
	Title   string                `json:"title,omitempty"`
	Excerpt string                `json:"excerpt,omitempty"`
	Source  *ArticleSourceScheme  `json:"source,omitempty"`
	Content *ArticleContentScheme `json:"content,omitempty"`
}

type ArticleSourceScheme struct {
	Type string `json:"type,omitempty"`
}

type ArticleContentScheme struct {
	IframeSrc string `json:"iframeSrc,omitempty"`
}
