package models

// ServerInformationScheme represents the server information in Jira.
type ServerInformationScheme struct {
	BaseURL        string                     `json:"baseUrl,omitempty"`        // The base URL of the Jira server.
	Version        string                     `json:"version,omitempty"`        // The version of the Jira server.
	VersionNumbers []int                      `json:"versionNumbers,omitempty"` // The version numbers of the Jira server.
	DeploymentType string                     `json:"deploymentType,omitempty"` // The deployment type of the Jira server.
	BuildNumber    int                        `json:"buildNumber,omitempty"`    // The build number of the Jira server.
	BuildDate      string                     `json:"buildDate,omitempty"`      // The build date of the Jira server.
	ServerTime     string                     `json:"serverTime,omitempty"`     // The server time of the Jira server.
	ScmInfo        string                     `json:"scmInfo,omitempty"`        // The SCM information of the Jira server.
	ServerTitle    string                     `json:"serverTitle,omitempty"`    // The server title of the Jira server.
	HealthChecks   []*ServerHealthCheckScheme `json:"healthChecks,omitempty"`   // The health checks of the Jira server.
}

// ServerHealthCheckScheme represents a health check of a server in Jira.
type ServerHealthCheckScheme struct {
	Name        string `json:"name,omitempty"`        // The name of the health check.
	Description string `json:"description,omitempty"` // The description of the health check.
	Passed      bool   `json:"passed,omitempty"`      // Indicates if the health check passed.
}
