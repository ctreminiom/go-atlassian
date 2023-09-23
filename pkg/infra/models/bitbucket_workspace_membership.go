package models

import "time"

type WorkspaceMembershipPageScheme struct {
	Size     int                          `json:"size,omitempty"`
	Page     int                          `json:"page,omitempty"`
	Pagelen  int                          `json:"pagelen,omitempty"`
	Next     string                       `json:"next,omitempty"`
	Previous string                       `json:"previous,omitempty"`
	Values   []*WorkspaceMembershipScheme `json:"values,omitempty"`
}

type WorkspaceMembershipScheme struct {
	Links        *WorkspaceMembershipLinksScheme `json:"links,omitempty"`
	User         *BitbucketAccountScheme         `json:"user,omitempty"`
	Workspace    *WorkspaceScheme                `json:"workspace,omitempty"`
	AddedOn      time.Time                       `json:"added_on,omitempty"`
	Permission   string                          `json:"permission,omitempty"`
	LastAccessed time.Time                       `json:"last_accessed,omitempty"`
}

type WorkspaceMembershipLinksScheme struct {
	Self *BitbucketLinkScheme `json:"self,omitempty"`
}

type BitbucketAccountScheme struct {
	Links       *BitbucketAccountLinksScheme `json:"links,omitempty"`
	CreatedOn   string                       `json:"created_on,omitempty"`
	DisplayName string                       `json:"display_name,omitempty"`
	Username    string                       `json:"username,omitempty"`
	Uuid        string                       `json:"uuid,omitempty"`
	Type        string                       `json:"type,omitempty"`
	AccountId   string                       `json:"account_id,omitempty"`
	Nickname    string                       `json:"nickname,omitempty"`
}

type BitbucketAccountLinksScheme struct {
	Avatar *BitbucketLinkScheme `json:"avatar,omitempty"`
	Self   *BitbucketLinkScheme `json:"self,omitempty"`
	Html   *BitbucketLinkScheme `json:"html,omitempty"`
}

/*

 */

type T struct {
	Values []struct {
		Repository struct {
			Type     string `json:"type"`
			FullName string `json:"full_name"`
			Links    struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Html struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Name string `json:"name"`
			Uuid string `json:"uuid"`
		} `json:"repository"`
		Type       string `json:"type"`
		Permission string `json:"permission"`
	} `json:"values"`
	Pagelen int `json:"pagelen"`
	Size    int `json:"size"`
	Page    int `json:"page"`
}
