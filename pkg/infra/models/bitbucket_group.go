package models

type BitbucketProjectGroupPageScheme struct {
	Values  []*BitbucketProjectGroupScheme `json:"values"`
	Pagelen int                            `json:"pagelen"`
	Size    int                            `json:"size"`
	Page    int                            `json:"page"`
}

type BitbucketProjectGroupScheme struct {
	Group      *BitbucketGroupScheme `json:"group"`
	Type       string                `json:"type"`
	Permission string                `json:"permission"`
}

type BitbucketGroupScheme struct {
	Type  string `json:"type"`
	Owner struct {
		DisplayName string `json:"display_name"`
		Links       struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
		} `json:"links"`
		Type     string `json:"type"`
		Uuid     string `json:"uuid"`
		Username string `json:"username"`
	} `json:"owner"`
	Workspace               *WorkspaceScheme `json:"workspace"`
	Slug                    string           `json:"slug"`
	FullSlug                string           `json:"full_slug"`
	Name                    string           `json:"name"`
	DefaultPermission       interface{}      `json:"default_permission"`
	EmailForwardingDisabled bool             `json:"email_forwarding_disabled"`
	Links                   struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Html struct {
			Href string `json:"href"`
		} `json:"html"`
	} `json:"links"`
	AccountPrivilege string `json:"account_privilege"`
}
