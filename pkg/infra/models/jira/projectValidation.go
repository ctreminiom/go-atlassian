package jira

type ProjectValidationMessageScheme struct {
	ErrorMessages []string `json:"errorMessages"`
	Errors        struct {
		ProjectKey string `json:"projectKey"`
	} `json:"errors"`
}
