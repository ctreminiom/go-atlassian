package models

type SearchContentOptions struct {
	Context                  string
	Cursor                   string
	Next                     bool
	Prev                     bool
	Limit                    int
	Start                    int
	IncludeArchivedSpaces    bool
	ExcludeCurrentSpaces     bool
	SitePermissionTypeFilter string
	Excerpt                  string
	Expand                   []string
}

type SearchPageScheme struct {
	Results             []*SearchResultScheme  `json:"results,omitempty"`
	Start               int                    `json:"start,omitempty"`
	Limit               int                    `json:"limit,omitempty"`
	Size                int                    `json:"size,omitempty"`
	TotalSize           int                    `json:"totalSize,omitempty"`
	CqlQuery            string                 `json:"cqlQuery,omitempty"`
	SearchDuration      int                    `json:"searchDuration,omitempty"`
	ArchivedResultCount int                    `json:"archivedResultCount,omitempty"`
	Links               *SearchPageLinksScheme `json:"_links,omitempty"`
}

type SearchPageLinksScheme struct {
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Self    string `json:"self,omitempty"`
}

type SearchResultScheme struct {
	Content               *ContentScheme            `json:"content,omitempty"`
	User                  *ContentUserScheme        `json:"user,omitempty"`
	Space                 *SpaceScheme              `json:"space,omitempty"`
	Title                 string                    `json:"title,omitempty"`
	Excerpt               string                    `json:"excerpt,omitempty"`
	URL                   string                    `json:"url,omitempty"`
	ResultParentContainer *ContainerSummaryScheme   `json:"resultParentContainer,omitempty"`
	ResultGlobalContainer *ContainerSummaryScheme   `json:"resultGlobalContainer,omitempty"`
	Breadcrumbs           []*SearchBreadcrumbScheme `json:"breadcrumbs,omitempty"`
	EntityType            string                    `json:"entityType,omitempty"`
	IconCSSClass          string                    `json:"iconCssClass,omitempty"`
	LastModified          string                    `json:"lastModified,omitempty"`
	FriendlyLastModified  string                    `json:"friendlyLastModified,omitempty"`
	Score                 float64                   `json:"score,omitempty"`
}

type ContainerSummaryScheme struct {
	Title      string `json:"title,omitempty"`
	DisplayURL string `json:"displayUrl,omitempty"`
}

type SearchBreadcrumbScheme struct {
	Label     string `json:"label,omitempty"`
	URL       string `json:"url,omitempty"`
	Separator string `json:"separator,omitempty"`
}
