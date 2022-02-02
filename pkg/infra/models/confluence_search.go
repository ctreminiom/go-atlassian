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
	Result              []*SearchResultScheme `json:"results,omitempty"`
	Start               int                   `json:"start,omitempty"`
	Limit               int                   `json:"limit,omitempty"`
	Size                int                   `json:"size,omitempty"`
	TotalSize           int                   `json:"totalSize,omitempty"`
	CqlQuery            string                `json:"cqlQuery,omitempty"`
	SearchDuration      int                   `json:"searchDuration,omitempty"`
	ArchivedResultCount int                   `json:"archivedResultCount,omitempty"`
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
	EntityType            string                    `json:"entityType"`
	IconCSSClass          string                    `json:"iconCssClass"`
	LastModified          string                    `json:"lastModified"`
	FriendlyLastModified  string                    `json:"friendlyLastModified"`
	Score                 int                       `json:"score"`
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
