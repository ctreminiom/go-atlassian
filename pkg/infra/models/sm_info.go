package models

// InfoScheme represents the information about a system.
type InfoScheme struct {
	Version          string               `json:"version,omitempty"`          // The version of the system.
	PlatformVersion  string               `json:"platformVersion,omitempty"`  // The platform version of the system.
	BuildDate        *InfoBuildDataScheme `json:"buildDate,omitempty"`        // The build date of the system.
	BuildChangeSet   string               `json:"buildChangeSet,omitempty"`   // The build change set of the system.
	IsLicensedForUse bool                 `json:"isLicensedForUse,omitempty"` // Indicates if the system is licensed for use.
	Links            *InfoLinkScheme      `json:"_links,omitempty"`           // Links related to the system.
}

// InfoBuildDataScheme represents the build date of a system.
type InfoBuildDataScheme struct {
	ISO8601     DateTimeScheme `json:"iso8601,omitempty"`     // The ISO 8601 format of the build date.
	Jira        string         `json:"jira,omitempty"`        // The Jira format of the build date.
	Friendly    string         `json:"friendly,omitempty"`    // The friendly format of the build date.
	EpochMillis int64          `json:"epochMillis,omitempty"` // The epoch milliseconds of the build date.
}

// InfoLinkScheme represents a link related to a system.
type InfoLinkScheme struct {
	Self string `json:"self,omitempty"` // The URL of the system itself.
}
