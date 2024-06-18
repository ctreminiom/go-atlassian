package models

// ProjectValidationMessageScheme represents a validation message for a project in Jira.
type ProjectValidationMessageScheme struct {
	ErrorMessages []string `json:"errorMessages"` // The error messages of the project validation.
	Errors        struct {
		ProjectKey string `json:"projectKey"` // The key of the project that has errors.
	} `json:"errors"` // The errors of the project validation.
}
