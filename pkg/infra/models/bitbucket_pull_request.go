package models

type BitbucketPullRequestPageScheme struct {
	Size     int                           `json:"size"`
	Page     int                           `json:"page"`
	Pagelen  int                           `json:"pagelen"`
	Next     string                        `json:"next"`
	Previous string                        `json:"previous"`
	Values   []*BitbucketPullRequestScheme `json:"values"`
}

type BitbucketPullRequestScheme struct {
	Links             *BitbucketPullRequestLinksScheme      `json:"links,omitempty"`
	Id                int                                   `json:"id,omitempty"`
	Title             string                                `json:"title,omitempty"`
	Rendered          *BitbucketPullRequestRenderedScheme   `json:"rendered,omitempty"`
	Summary           *BitbucketPullRequestRenderItemScheme `json:"summary,omitempty"`
	State             string                                `json:"state,omitempty"`
	Author            *BitbucketAccountScheme               `json:"author,omitempty"`
	Source            *BitbucketPullRequestEndpointScheme   `json:"source,omitempty"`
	Destination       *BitbucketPullRequestEndpointScheme   `json:"destination,omitempty"`
	MergeCommit       *BitbucketPullRequestCommitScheme     `json:"merge_commit,omitempty"`
	CommentCount      int                                   `json:"comment_count,omitempty"`
	TaskCount         int                                   `json:"task_count,omitempty"`
	CloseSourceBranch bool                                  `json:"close_source_branch,omitempty"`
	ClosedBy          *BitbucketAccountScheme               `json:"closed_by,omitempty"`
	Reason            string                                `json:"reason,omitempty"`
	CreatedOn         string                                `json:"created_on,omitempty"`
	UpdatedOn         string                                `json:"updated_on,omitempty"`
	Reviewers         []*BitbucketAccountScheme             `json:"reviewers,omitempty"`
	Participants      []*BitbucketParticipantScheme         `json:"participants,omitempty"`
}

type BitbucketPullRequestLinksScheme struct {
	Self     *BitbucketLinkScheme `json:"self,omitempty"`
	Html     *BitbucketLinkScheme `json:"html,omitempty"`
	Commits  *BitbucketLinkScheme `json:"commits,omitempty"`
	Approve  *BitbucketLinkScheme `json:"approve,omitempty"`
	Diff     *BitbucketLinkScheme `json:"diff,omitempty"`
	Diffstat *BitbucketLinkScheme `json:"diffstat,omitempty"`
	Comments *BitbucketLinkScheme `json:"comments,omitempty"`
	Activity *BitbucketLinkScheme `json:"activity,omitempty"`
	Merge    *BitbucketLinkScheme `json:"merge,omitempty"`
	Decline  *BitbucketLinkScheme `json:"decline,omitempty"`
}

type BitbucketPullRequestRenderedScheme struct {
	Title       *BitbucketPullRequestRenderItemScheme `json:"title,omitempty"`
	Description *BitbucketPullRequestRenderItemScheme `json:"description,omitempty"`
	Reason      *BitbucketPullRequestRenderItemScheme `json:"reason,omitempty"`
}

type BitbucketPullRequestRenderItemScheme struct {
	Raw    string `json:"raw,omitempty"`
	Markup string `json:"markup,omitempty"`
	Html   string `json:"html,omitempty"`
}

type BitbucketPullRequestEndpointScheme struct {
	Repository *RepositoryScheme                 `json:"repository,omitempty"`
	Branch     *BitbucketPullRequestBranchScheme `json:"branch,omitempty"`
	Commit     *BitbucketPullRequestCommitScheme `json:"commit,omitempty"`
}

type BitbucketPullRequestBranchScheme struct {
	Name                 string   `json:"name,omitempty"`
	MergeStrategies      []string `json:"merge_strategies,omitempty"`
	DefaultMergeStrategy string   `json:"default_merge_strategy,omitempty"`
}

type BitbucketPullRequestCommitScheme struct {
	Hash string `json:"hash,omitempty"`
}

type BitbucketParticipantScheme struct {
	User           *BitbucketAccountScheme `json:"user"`
	Role           string                  `json:"role"`
	Approved       bool                    `json:"approved"`
	State          string                  `json:"state"`
	ParticipatedOn string                  `json:"participated_on"`
}
