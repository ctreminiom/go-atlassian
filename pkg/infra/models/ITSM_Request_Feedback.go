package models

type CustomerFeedbackScheme struct {
	Type    string                         `json:"type,omitempty"`
	Rating  int                            `json:"rating,omitempty"`
	Comment *CustomerFeedbackCommentScheme `json:"comment,omitempty"`
}

type CustomerFeedbackCommentScheme struct {
	Body string `json:"body,omitempty"`
}
