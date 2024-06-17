package models

// ProjectRolesScheme represents the roles of a project in Jira.
type ProjectRolesScheme struct {
	AtlassianAddonsProjectAccess string `json:"atlassian-addons-project-access,omitempty"` // The access of Atlassian addons in the project.
	ServiceDeskTeam              string `json:"Service Desk Team,omitempty"`               // The team of the service desk in the project.
	ServiceDeskCustomers         string `json:"Service Desk Customers,omitempty"`          // The customers of the service desk in the project.
	Administrators               string `json:"Administrators,omitempty"`                  // The administrators of the project.
}

// ProjectRoleScheme represents a role in a project in Jira.
type ProjectRoleScheme struct {
	Self             string                         `json:"self,omitempty"`             // The URL of the project role.
	Name             string                         `json:"name,omitempty"`             // The name of the project role.
	ID               int                            `json:"id,omitempty"`               // The ID of the project role.
	Description      string                         `json:"description,omitempty"`      // The description of the project role.
	Actors           []*RoleActorScheme             `json:"actors,omitempty"`           // The actors of the project role.
	Scope            *TeamManagedProjectScopeScheme `json:"scope,omitempty"`            // The scope of the project role.
	TranslatedName   string                         `json:"translatedName,omitempty"`   // The translated name of the project role.
	CurrentUserRole  bool                           `json:"currentUserRole,omitempty"`  // Indicates if the current user has this role.
	Admin            bool                           `json:"admin,omitempty"`            // Indicates if the project role has admin privileges.
	RoleConfigurable bool                           `json:"roleConfigurable,omitempty"` // Indicates if the project role is configurable.
	Default          bool                           `json:"default,omitempty"`          // Indicates if the project role is the default role.
}

// RoleActorScheme represents an actor in a role in a project in Jira.
type RoleActorScheme struct {
	ID          int                  `json:"id,omitempty"`          // The ID of the role actor.
	DisplayName string               `json:"displayName,omitempty"` // The display name of the role actor.
	Type        string               `json:"type,omitempty"`        // The type of the role actor.
	Name        string               `json:"name,omitempty"`        // The name of the role actor.
	AvatarURL   string               `json:"avatarUrl,omitempty"`   // The avatar URL of the role actor.
	ActorGroup  *GroupScheme         `json:"actorGroup,omitempty"`  // The group of the role actor.
	ActorUser   *RoleActorUserScheme `json:"actorUser,omitempty"`   // The user of the role actor.
}

// RoleActorUserScheme represents a user in a role actor in a project in Jira.
type RoleActorUserScheme struct {
	AccountID string `json:"accountId,omitempty"` // The account ID of the role actor user.
}
