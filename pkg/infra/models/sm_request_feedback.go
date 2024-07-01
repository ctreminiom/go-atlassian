package models

// CustomerFeedbackScheme represents the feedback provided by a customer.
type CustomerFeedbackScheme struct {
	Type    string                         `json:"type,omitempty"`    // The type of feedback.
	Rating  int                            `json:"rating,omitempty"`  // The rating provided in the feedback.
	Comment *CustomerFeedbackCommentScheme `json:"comment,omitempty"` // The comment provided in the feedback.
}

// CustomerFeedbackCommentScheme represents the comment provided in a customer's feedback.
type CustomerFeedbackCommentScheme struct {
	Body string `json:"body,omitempty"` // The body of the comment.
}
