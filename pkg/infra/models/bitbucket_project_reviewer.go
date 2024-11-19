package models

type ProjectReviewerPageScheme struct {
	Size     int                      `json:"size,omitempty"`
	Page     int                      `json:"page,omitempty"`
	Pagelen  int                      `json:"pagelen,omitempty"`
	Next     string                   `json:"next,omitempty"`
	Previous string                   `json:"previous,omitempty"`
	Values   []*ProjectReviewerScheme `json:"values,omitempty"`
}

type ProjectReviewerScheme struct {
	Type         string                  `json:"type,omitempty"`
	ReviewerType string                  `json:"reviewer_type,omitempty"`
	User         *BitbucketAccountScheme `json:"user,omitempty"`
}
