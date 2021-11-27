package models

type IssueMetadataCreateOptions struct {
	ProjectIDs     []string
	ProjectKeys    []string
	IssueTypeIDs   []string
	IssueTypeNames []string
	Expand         string
}
