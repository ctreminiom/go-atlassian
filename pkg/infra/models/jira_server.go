package models

type ServerInformationScheme struct {
	BaseURL        string                     `json:"baseUrl,omitempty"`
	Version        string                     `json:"version,omitempty"`
	VersionNumbers []int                      `json:"versionNumbers,omitempty"`
	DeploymentType string                     `json:"deploymentType,omitempty"`
	BuildNumber    int                        `json:"buildNumber,omitempty"`
	BuildDate      string                     `json:"buildDate,omitempty"`
	ServerTime     string                     `json:"serverTime,omitempty"`
	ScmInfo        string                     `json:"scmInfo,omitempty"`
	ServerTitle    string                     `json:"serverTitle,omitempty"`
	HealthChecks   []*ServerHealthCheckScheme `json:"healthChecks,omitempty"`
}

type ServerHealthCheckScheme struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Passed      bool   `json:"passed,omitempty"`
}
