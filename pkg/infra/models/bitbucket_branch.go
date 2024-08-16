package models

// BranchScheme represents a branch in a repository.
type BranchScheme struct {
	MergeStrategies      []string `json:"merge_strategies"`       // The merge strategies available for the branch.
	DefaultMergeStrategy string   `json:"default_merge_strategy"` // The default merge strategy used for the branch.
}
