package models

// JiraTeamPageScheme represents a page of teams in Jira.
type JiraTeamPageScheme struct {
	MoreResultsAvailable bool                    `json:"moreResultsAvailable,omitempty"` // Indicates if more results are available.
	Teams                []*JiraTeamScheme       `json:"teams,omitempty"`                // The teams on the page.
	Persons              []*JiraTeamPersonScheme `json:"persons,omitempty"`              // The persons on the page.
	ZeroMemberTeamsCount int                     `json:"zeroMemberTeamsCount,omitempty"` // The count of teams with zero members.
}

// JiraTeamScheme represents a team in Jira.
type JiraTeamScheme struct {
	ID         int                       `json:"id,omitempty"`         // The ID of the team.
	ExternalID string                    `json:"externalId,omitempty"` // The external ID of the team.
	Title      string                    `json:"title,omitempty"`      // The title of the team.
	Shareable  bool                      `json:"shareable,omitempty"`  // Indicates if the team is shareable.
	Resources  []*JiraTeamResourceScheme `json:"resources,omitempty"`  // The resources of the team.
}

// JiraTeamResourceScheme represents a resource in a Jira team.
type JiraTeamResourceScheme struct {
	ID       int `json:"id,omitempty"`       // The ID of the resource.
	PersonID int `json:"personId,omitempty"` // The ID of the person associated with the resource.
}

// JiraTeamPersonScheme represents a person in a Jira team.
type JiraTeamPersonScheme struct {
	PersonID int                 `json:"personId,omitempty"` // The ID of the person.
	JiraUser *JiraTeamUserScheme `json:"jiraUser,omitempty"` // The Jira user associated with the person.
}

// JiraTeamUserScheme represents a user in a Jira team.
type JiraTeamUserScheme struct {
	Title     string `json:"title,omitempty"`     // The title of the user.
	Email     string `json:"email,omitempty"`     // The email of the user.
	AvatarURL string `json:"avatarUrl,omitempty"` // The avatar URL of the user.
}

// JiraTeamCreatePayloadScheme represents the payload for creating a Jira team.
type JiraTeamCreatePayloadScheme struct {
	Title     string                    `json:"title,omitempty"`     // The title of the team.
	Shareable bool                      `json:"shareable,omitempty"` // Indicates if the team is shareable.
	Resources []*JiraTeamResourceScheme `json:"resources,omitempty"` // The resources of the team.
}

// JiraTeamCreateResponseScheme represents the response from creating a Jira team.
type JiraTeamCreateResponseScheme struct {
	ID      int                     `json:"id,omitempty"`      // The ID of the created team.
	Team    *JiraTeamScheme         `json:"team,omitempty"`    // The created team.
	Persons []*JiraTeamPersonScheme `json:"persons,omitempty"` // The persons associated with the created team.
}
