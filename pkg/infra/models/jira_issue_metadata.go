package models

// IssueMetadataCreateOptions represents the options for creating issue metadata in Jira.
type IssueMetadataCreateOptions struct {
	ProjectIDs     []string // The IDs of the projects.
	ProjectKeys    []string // The keys of the projects.
	IssueTypeIDs   []string // The IDs of the issue types.
	IssueTypeNames []string // The names of the issue types.
	Expand         string   // The fields to be expanded in the issue metadata.
}
