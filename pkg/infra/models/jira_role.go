package models

type ProjectRolesScheme struct {
	AtlassianAddonsProjectAccess string `json:"atlassian-addons-project-access,omitempty"`
	ServiceDeskTeam              string `json:"Service Desk Team,omitempty"`
	ServiceDeskCustomers         string `json:"Service Desk Customers,omitempty"`
	Administrators               string `json:"Administrators,omitempty"`
}

type ProjectRoleScheme struct {
	Self             string                         `json:"self,omitempty"`
	Name             string                         `json:"name,omitempty"`
	ID               int                            `json:"id,omitempty"`
	Description      string                         `json:"description,omitempty"`
	Actors           []*RoleActorScheme             `json:"actors,omitempty"`
	Scope            *TeamManagedProjectScopeScheme `json:"scope,omitempty"`
	TranslatedName   string                         `json:"translatedName,omitempty"`
	CurrentUserRole  bool                           `json:"currentUserRole,omitempty"`
	Admin            bool                           `json:"admin,omitempty"`
	RoleConfigurable bool                           `json:"roleConfigurable,omitempty"`
	Default          bool                           `json:"default,omitempty"`
}

type RoleActorScheme struct {
	ID          int                  `json:"id,omitempty"`
	DisplayName string               `json:"displayName,omitempty"`
	Type        string               `json:"type,omitempty"`
	Name        string               `json:"name,omitempty"`
	AvatarURL   string               `json:"avatarUrl,omitempty"`
	ActorGroup  *GroupScheme         `json:"actorGroup,omitempty"`
	ActorUser   *RoleActorUserScheme `json:"actorUser,omitempty"`
}

type RoleActorUserScheme struct {
	AccountID string `json:"accountId,omitempty"`
}
