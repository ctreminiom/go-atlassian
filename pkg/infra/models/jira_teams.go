package models

type JiraTeamPageScheme struct {
	MoreResultsAvailable bool                    `json:"moreResultsAvailable,omitempty"`
	Teams                []*JiraTeamScheme       `json:"teams,omitempty"`
	Persons              []*JiraTeamPersonScheme `json:"persons,omitempty"`
	ZeroMemberTeamsCount int                     `json:"zeroMemberTeamsCount,omitempty"`
}

type JiraTeamScheme struct {
	Id         int                       `json:"id,omitempty"`
	ExternalId string                    `json:"externalId,omitempty"`
	Title      string                    `json:"title,omitempty"`
	Shareable  bool                      `json:"shareable,omitempty"`
	Resources  []*JiraTeamResourceScheme `json:"resources,omitempty"`
}

type JiraTeamResourceScheme struct {
	Id       int `json:"id,omitempty"`
	PersonId int `json:"personId,omitempty"`
}

type JiraTeamPersonScheme struct {
	PersonId int                 `json:"personId,omitempty"`
	JiraUser *JiraTeamUserScheme `json:"jiraUser,omitempty"`
}

type JiraTeamUserScheme struct {
	Title     string `json:"title,omitempty"`
	Email     string `json:"email,omitempty"`
	AvatarUrl string `json:"avatarUrl,omitempty"`
}

type JiraTeamCreatePayloadScheme struct {
	Title     string                    `json:"title,omitempty"`
	Shareable bool                      `json:"shareable,omitempty"`
	Resources []*JiraTeamResourceScheme `json:"resources,omitempty"`
}

type JiraTeamCreateResponseScheme struct {
	Id      int                     `json:"id,omitempty"`
	Team    *JiraTeamScheme         `json:"team,omitempty"`
	Persons []*JiraTeamPersonScheme `json:"persons,omitempty"`
}
