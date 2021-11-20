package models

type InfoScheme struct {
	Version          string               `json:"version,omitempty"`
	PlatformVersion  string               `json:"platformVersion,omitempty"`
	BuildDate        *InfoBuildDataScheme `json:"buildDate,omitempty"`
	BuildChangeSet   string               `json:"buildChangeSet,omitempty"`
	IsLicensedForUse bool                 `json:"isLicensedForUse,omitempty"`
	Links            *InfoLinkScheme      `json:"_links,omitempty"`
}

type InfoBuildDataScheme struct {
	Iso8601     string `json:"iso8601,omitempty"`
	Jira        string `json:"jira,omitempty"`
	Friendly    string `json:"friendly,omitempty"`
	EpochMillis int64  `json:"epochMillis,omitempty"`
}

type InfoLinkScheme struct {
	Self string `json:"self,omitempty"`
}
