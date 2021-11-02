package jira

type VersionScheme struct {
	Self                      string                                  `json:"self,omitempty"`
	ID                        string                                  `json:"id,omitempty"`
	Description               string                                  `json:"description,omitempty"`
	Name                      string                                  `json:"name,omitempty"`
	Archived                  bool                                    `json:"archived,omitempty"`
	Released                  bool                                    `json:"released,omitempty"`
	ReleaseDate               string                                  `json:"releaseDate,omitempty"`
	Overdue                   bool                                    `json:"overdue,omitempty"`
	UserReleaseDate           string                                  `json:"userReleaseDate,omitempty"`
	ProjectID                 int                                     `json:"projectId,omitempty"`
	Operations                []*VersionOperation                     `json:"operations,omitempty"`
	IssuesStatusForFixVersion *VersionIssuesStatusForFixVersionScheme `json:"issuesStatusForFixVersion,omitempty"`
}

type VersionOperation struct {
	ID         string `json:"id,omitempty"`
	StyleClass string `json:"styleClass,omitempty"`
	Label      string `json:"label,omitempty"`
	Href       string `json:"href,omitempty"`
	Weight     int    `json:"weight,omitempty"`
}

type VersionIssuesStatusForFixVersionScheme struct {
	Unmapped   int `json:"unmapped,omitempty"`
	ToDo       int `json:"toDo,omitempty"`
	InProgress int `json:"inProgress,omitempty"`
	Done       int `json:"done,omitempty"`
}
