package models

type BranchScheme struct {
	MergeStrategies      []string `json:"merge_strategies"`
	DefaultMergeStrategy string   `json:"default_merge_strategy"`
}
