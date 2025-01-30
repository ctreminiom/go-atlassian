// Package models provides the data structures used in the admin package.
package models

import "time"

// AdminOrganizationPageScheme represents a page of organizations.
type AdminOrganizationPageScheme struct {
	Data  []*OrganizationModelScheme `json:"data,omitempty"`  // The organizations on this page.
	Links *LinkPageModelScheme       `json:"links,omitempty"` // Links to other pages.
}

// LinkPageModelScheme represents the links to other pages.
type LinkPageModelScheme struct {
	Self string `json:"self,omitempty"` // Link to this page.
	Prev string `json:"prev,omitempty"` // Link to the previous page.
	Next string `json:"next,omitempty"` // Link to the next page.
}

// OrganizationModelScheme represents an organization.
type OrganizationModelScheme struct {
	ID            string                          `json:"id,omitempty"`            // The ID of the organization.
	Type          string                          `json:"type,omitempty"`          // The type of the organization.
	Attributes    *OrganizationModelAttribute     `json:"attributes,omitempty"`    // The attributes of the organization.
	Relationships *OrganizationModelRelationships `json:"relationships,omitempty"` // The relationships of the organization.
	Links         *LinkSelfModelScheme            `json:"links,omitempty"`         // Links related to the organization.
}

// OrganizationModelAttribute represents the attributes of an organization.
type OrganizationModelAttribute struct {
	Name string `json:"name,omitempty"` // The name of the organization.
}

// OrganizationModelRelationships represents the relationships of an organization.
type OrganizationModelRelationships struct {
	Domains *OrganizationModelSchemes `json:"domains,omitempty"` // The domains of the organization.
	Users   *OrganizationModelSchemes `json:"users,omitempty"`   // The users of the organization.
}

// OrganizationModelSchemes represents the links to related entities.
type OrganizationModelSchemes struct {
	Links struct {
		Related string `json:"related,omitempty"` // Link to the related entity.
	} `json:"links,omitempty"`
}

// LinkSelfModelScheme represents a link to the entity itself.
type LinkSelfModelScheme struct {
	Self string `json:"self,omitempty"` // Link to the entity itself.
}

// AdminOrganizationScheme represents an organization.
type AdminOrganizationScheme struct {
	Data *OrganizationModelScheme `json:"data,omitempty"` // The organization data.
}

// OrganizationUserPageScheme represents a page of users in an organization.
type OrganizationUserPageScheme struct {
	Data  []*AdminOrganizationUserScheme `json:"data,omitempty"`  // The users on this page.
	Links *LinkPageModelScheme           `json:"links,omitempty"` // Links to other pages.
	Meta  struct {
		Total int `json:"total,omitempty"` // The total number of users.
	} `json:"meta,omitempty"`
}

// AdminOrganizationUserScheme represents a user in an organization.
type AdminOrganizationUserScheme struct {
	AccountID      string                           `json:"account_id,omitempty"`      // The account ID of the user.
	AccountType    string                           `json:"account_type,omitempty"`    // The account type of the user.
	AccountStatus  string                           `json:"account_status,omitempty"`  // The account status of the user.
	Name           string                           `json:"name,omitempty"`            // The name of the user.
	Picture        string                           `json:"picture,omitempty"`         // The picture of the user.
	Email          string                           `json:"email,omitempty"`           // The email of the user.
	AccessBillable bool                             `json:"access_billable,omitempty"` // Whether the user is billable.
	LastActive     string                           `json:"last_active,omitempty"`     // The last active time of the user.
	ProductAccess  []*OrganizationUserProductScheme `json:"product_access,omitempty"`  // The products the user has access to.
	Links          *LinkSelfModelScheme             `json:"links,omitempty"`           // Links related to the user.
}

// OrganizationUserProductScheme represents a product a user has access to.
type OrganizationUserProductScheme struct {
	Key        string `json:"key,omitempty"`         // The key of the product.
	Name       string `json:"name,omitempty"`        // The name of the product.
	URL        string `json:"url,omitempty"`         // The URL of the product.
	LastActive string `json:"last_active,omitempty"` // The last active time of the product.
}

// OrganizationDomainPageScheme represents a page of domains in an organization.
type OrganizationDomainPageScheme struct {
	Data  []*OrganizationDomainModelScheme `json:"data,omitempty"`  // The domains on this page.
	Links *LinkPageModelScheme             `json:"links,omitempty"` // Links to other pages.
}

// OrganizationDomainModelScheme represents a domain in an organization.
type OrganizationDomainModelScheme struct {
	ID         string                                   `json:"id,omitempty"`         // The ID of the domain.
	Type       string                                   `json:"type,omitempty"`       // The type of the domain.
	Attributes *OrganizationDomainModelAttributesScheme `json:"attributes,omitempty"` // The attributes of the domain.
	Links      *LinkSelfModelScheme                     `json:"links,omitempty"`      // Links related to the domain.
}

// OrganizationDomainModelAttributesScheme represents the attributes of a domain.
type OrganizationDomainModelAttributesScheme struct {
	Name  string                                       `json:"name,omitempty"`  // The name of the domain.
	Claim *OrganizationDomainModelAttributeClaimScheme `json:"claim,omitempty"` // The claim of the domain.
}

// OrganizationDomainModelAttributeClaimScheme represents the claim of a domain.
type OrganizationDomainModelAttributeClaimScheme struct {
	Type   string `json:"type,omitempty"`   // The type of the claim.
	Status string `json:"status,omitempty"` // The status of the claim.
}

// OrganizationDomainScheme represents a domain.
type OrganizationDomainScheme struct {
	Data *OrganizationDomainDataScheme `json:"data"` // The domain data.
}

// OrganizationDomainDataScheme represents the data of a domain.
type OrganizationDomainDataScheme struct {
	ID         string `json:"id"`   // The ID of the domain.
	Type       string `json:"type"` // The type of the domain.
	Attributes struct {
		Name  string `json:"name"` // The name of the domain.
		Claim struct {
			Type   string `json:"type"`   // The type of the claim.
			Status string `json:"status"` // The status of the claim.
		} `json:"claim"` // The claim of the domain.
	} `json:"attributes"` // The attributes of the domain.
	Links struct {
		Self string `json:"self"` // Link to the domain itself.
	} `json:"links"` // Links related to the domain.
}

// OrganizationEventOptScheme represents the options for getting events.
type OrganizationEventOptScheme struct {
	Q      string    //Single query term for searching events.
	From   time.Time //The earliest date and time of the event represented as a UNIX epoch time.
	To     time.Time //The latest date and time of the event represented as a UNIX epoch time.
	Action string    //A query filter that returns events of a specific action type.
}

// OrganizationEventPageScheme represents a page of events in an organization.
type OrganizationEventPageScheme struct {
	Data  []*OrganizationEventModelScheme `json:"data,omitempty"`  // The events on this page.
	Links *LinkPageModelScheme            `json:"links,omitempty"` // Links to other pages.
	Meta  struct {
		Next     string `json:"next,omitempty"`      // The next page.
		PageSize int    `json:"page_size,omitempty"` // The page size.
	} `json:"meta,omitempty"`
}

// OrganizationEventModelScheme represents an event in an organization.
type OrganizationEventModelScheme struct {
	ID         string                                  `json:"id,omitempty"`         // The ID of the event.
	Type       string                                  `json:"type,omitempty"`       // The type of the event.
	Attributes *OrganizationEventModelAttributesScheme `json:"attributes,omitempty"` // The attributes of the event.
	Links      *LinkSelfModelScheme                    `json:"links,omitempty"`      // Links related to the event.
}

// OrganizationEventModelAttributesScheme represents the attributes of an event.
type OrganizationEventModelAttributesScheme struct {
	Time      string                          `json:"time,omitempty"`      // The time of the event.
	Action    string                          `json:"action,omitempty"`    // The action of the event.
	Actor     *OrganizationEventActorModel    `json:"actor,omitempty"`     // The actor of the event.
	Context   []*OrganizationEventObjectModel `json:"context,omitempty"`   // The context of the event.
	Container []*OrganizationEventObjectModel `json:"container,omitempty"` // The container of the event.
	Location  *OrganizationEventLocationModel `json:"location,omitempty"`  // The location of the event.
}

// OrganizationEventActorModel represents the actor of an event.
type OrganizationEventActorModel struct {
	ID    string               `json:"id,omitempty"`    // The ID of the actor.
	Name  string               `json:"name,omitempty"`  // The name of the actor.
	Links *LinkSelfModelScheme `json:"links,omitempty"` // Links related to the actor.
}

// OrganizationEventObjectModel represents an object in the context or container of an event.
type OrganizationEventObjectModel struct {
	ID    string `json:"id,omitempty"`   // The ID of the object.
	Type  string `json:"type,omitempty"` // The type of the object.
	Links struct {
		Self string `json:"self,omitempty"` // Link to the object itself.
		Alt  string `json:"alt,omitempty"`  // Alternative link to the object.
	} `json:"links,omitempty"` // Links related to the object.
}

// OrganizationEventLocationModel represents the location of an event.
type OrganizationEventLocationModel struct {
	IP  string `json:"ip,omitempty"`  // The IP address of the location.
	Geo string `json:"geo,omitempty"` // The geographical location.
}

// OrganizationEventScheme represents an event.
type OrganizationEventScheme struct {
	Data *OrganizationEventModelScheme `json:"data,omitempty"` // The event data.
}

// OrganizationEventActionScheme represents an action in an event.
type OrganizationEventActionScheme struct {
	Data []*OrganizationEventActionModelScheme `json:"data,omitempty"` // The action data.
}

// OrganizationEventActionModelScheme represents an action in an event.
type OrganizationEventActionModelScheme struct {
	ID         string                                        `json:"id,omitempty"`         // The ID of the action.
	Type       string                                        `json:"type,omitempty"`       // The type of the action.
	Attributes *OrganizationEventActionModelAttributesScheme `json:"attributes,omitempty"` // The attributes of the action.
}

// OrganizationEventActionModelAttributesScheme represents the attributes of an action in an event.
type OrganizationEventActionModelAttributesScheme struct {
	DisplayName      string `json:"displayName,omitempty"`      // The display name of the action.
	GroupDisplayName string `json:"groupDisplayName,omitempty"` // The group display name of the action.
}

// UserProductAccessScheme represents the product access of a user.
type UserProductAccessScheme struct {
	Data *UserProductAccessDataScheme `json:"data,omitempty"` // The product access data.
}

// UserProductAccessDataScheme represents the data of a user's product access.
type UserProductAccessDataScheme struct {
	ProductAccess []*UserProductLastActiveScheme `json:"product_access,omitempty"` // The products the user has access to.
	AddedToOrg    string                         `json:"added_to_org,omitempty"`   // The time the user was added to the organization.
}

// UserProductLastActiveScheme represents a product a user has access to.
type UserProductLastActiveScheme struct {
	ID         string `json:"id,omitempty"`          // The ID of the product.
	Key        string `json:"key,omitempty"`         // The key of the product.
	Name       string `json:"name,omitempty"`        // The name of the product.
	URL        string `json:"url,omitempty"`         // The URL of the product.
	LastActive string `json:"last_active,omitempty"` // The last active time of the product.
}

// GenericActionSuccessScheme represents a successful action.
type GenericActionSuccessScheme struct {
	Message string `json:"message,omitempty"` // The success message.
}

// StringSearchCriteria represents a common search pattern with exact and partial matches
type StringSearchCriteria struct {
	Eq       []string `json:"eq,omitempty"`       // Exact matches
	Contains string   `json:"contains,omitempty"` // Partial match
}

// OrganizationUserSearchParams represents the parameters for searching users in an organization
type OrganizationUserSearchParams struct {
	AccountIds       []string             `json:"accountIds,omitempty"`
	AccountTypes     []string             `json:"accountTypes,omitempty"`
	AccountStatuses  []string             `json:"accountStatuses,omitempty"`
	NamesOrNicknames StringSearchCriteria `json:"namesOrNicknames,omitempty"`
	EmailUsernames   StringSearchCriteria `json:"emailUsernames,omitempty"`
	EmailDomains     StringSearchCriteria `json:"emailDomains,omitempty"`
	IsSuspended      *bool                `json:"isSuspended,omitempty"`
	Cursor           string               `json:"cursor,omitempty"`
	Limit            int                  `json:"limit,omitempty"`
	Expand           []string             `json:"expand,omitempty"`
}

// OrganizationUserSearchPage represents the response from searching users in an organization
type OrganizationUserSearchPage struct {
	Data  []OrganizationUserSearch `json:"data,omitempty"`
	Links struct {
		Next string `json:"next,omitempty"`
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
}

// OrganizationUserSearch represents a user returned from the search endpoint
type OrganizationUserSearch struct {
	AccountId         string              `json:"accountId,omitempty"`
	Name              string              `json:"name,omitempty"`
	Nickname          string              `json:"nickname,omitempty"`
	AccountType       string              `json:"accountType,omitempty"`
	AccountStatus     string              `json:"accountStatus,omitempty"`
	Email             string              `json:"email,omitempty"`
	EmailVerified     bool                `json:"emailVerified,omitempty"`
	StatusInUserbase  bool                `json:"statusInUserbase,omitempty"`
	ProductLastAccess []ProductLastAccess `json:"productLastAccess,omitempty"`
	Groups            []Group             `json:"groups,omitempty"`
}

// ProductLastAccess represents product access information for a user
type ProductLastAccess struct {
	ProductKey          string `json:"productKey,omitempty"`
	LastActiveTimestamp string `json:"lastActiveTimestamp,omitempty"`
	CloudSiteId         string `json:"cloudSiteId,omitempty"`
}

// Group represents a group that a user belongs to
type Group struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// GroupNames represents the search criteria for group names
type GroupNames struct {
	Eq       []string `json:"eq,omitempty"`       // Exact match for group names
	Contains string   `json:"contains,omitempty"` // Partial match for group names
}

// OrganizationGroupSearchParams represents the parameters for searching groups in an organization
type OrganizationGroupSearchParams struct {
	GroupIds   []string             `json:"groupIds,omitempty"`
	GroupNames StringSearchCriteria `json:"groupNames,omitempty"`
	Cursor     string               `json:"cursor,omitempty"`
	Limit      int                  `json:"limit,omitempty"`
	Expand     []string             `json:"expand,omitempty"`
}

// OrganizationGroupSearchPage represents the response from searching groups in an organization
type OrganizationGroupSearchPage struct {
	Data  []OrganizationGroupSearch `json:"data,omitempty"`
	Links struct {
		Next string `json:"next,omitempty"`
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
}

// OrganizationGroupSearch represents a group returned from the search endpoint
type OrganizationGroupSearch struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// Fields included when expand=META
	Meta struct {
		TotalUserCount int `json:"totalUserCount,omitempty"`
	} `json:"meta,omitempty"`
	// Fields included when expand=ROLE_ASSIGNMENTS
	RoleAssignments []RoleAssignment `json:"roleAssignments,omitempty"`
	// Fields included when expand=MANAGEMENT_ACCESS
	ManagementAccess string `json:"managementAccess,omitempty"`
	// Fields included when expand=USERS
	Users []GroupUser `json:"users,omitempty"`
}

// RoleAssignment represents a role assignment for a group
type RoleAssignment struct {
	ResourceID  string `json:"resourceId,omitempty"`
	PrincipalID string `json:"principalId,omitempty"`
	RoleID      string `json:"roleId,omitempty"`
}

// GroupUser represents a user in a group
type GroupUser struct {
	AccountID     string `json:"accountId,omitempty"`
	AccountType   string `json:"accountType,omitempty"`
	AccountStatus string `json:"accountStatus,omitempty"`
	Name          string `json:"name,omitempty"`
	Email         string `json:"email,omitempty"`
}

// WorkspaceSearchParams represents the parameters for searching workspaces in an organization
type WorkspaceSearchParams struct {
	Query  interface{} `json:"query,omitempty"` // Query can be AndOperator, FieldOperand, SearchWorkspacesOperand, FeatureFilter, or PolicyFilter
	Limit  int         `json:"limit,omitempty"`
	Sort   []SortField `json:"sort,omitempty"`
	Cursor string      `json:"cursor,omitempty"`
}

// SortField represents a field to sort by and its direction
type SortField struct {
	Field string `json:"field"`
	Order string `json:"order,omitempty"` // asc or desc
}

// WorkspaceSearchPage represents the response from searching workspaces
type WorkspaceSearchPage struct {
	Data  []WorkspaceSearch `json:"data,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
		Prev string `json:"prev,omitempty"`
		Next string `json:"next,omitempty"`
	} `json:"links,omitempty"`
	Meta struct {
		PageSize   int `json:"pageSize,omitempty"`
		StartIndex int `json:"startIndex,omitempty"`
		EndIndex   int `json:"endIndex,omitempty"`
		Total      int `json:"total,omitempty"`
	} `json:"meta,omitempty"`
}

// WorkspaceSearch represents a workspace returned from the search endpoint
type WorkspaceSearch struct {
	ID            string                 `json:"id,omitempty"`
	Type          string                 `json:"type,omitempty"`
	Attributes    WorkspaceAttributes    `json:"attributes,omitempty"`
	Links         WorkspaceLinks         `json:"links,omitempty"`
	Relationships WorkspaceRelationships `json:"relationships,omitempty"`
}

// WorkspaceLinks represents the links in a workspace
type WorkspaceLinks struct {
	Self string `json:"self,omitempty"`
}

// WorkspaceAttributes represents the attributes of a workspace
type WorkspaceAttributes struct {
	Name          string            `json:"name,omitempty"`
	TypeKey       string            `json:"typeKey,omitempty"`
	Type          string            `json:"type,omitempty"`
	Owner         string            `json:"owner,omitempty"`
	Status        string            `json:"status,omitempty"`
	StatusDetails []string          `json:"statusDetails,omitempty"`
	Icons         map[string]string `json:"icons,omitempty"`
	Avatars       map[string]string `json:"avatars,omitempty"`
	Labels        []string          `json:"labels,omitempty"`
	Sandbox       WorkspaceSandbox  `json:"sandbox,omitempty"`
	Usage         int               `json:"usage,omitempty"`
	Capacity      int               `json:"capacity,omitempty"`
	CreatedAt     string            `json:"createdAt,omitempty"`
	CreatedBy     string            `json:"createdBy,omitempty"`
	UpdatedAt     string            `json:"updatedAt,omitempty"`
	HostURL       string            `json:"hostUrl,omitempty"`
	Realm         string            `json:"realm,omitempty"`
	Regions       []string          `json:"regions,omitempty"`
}

// WorkspaceSandbox represents the sandbox information of a workspace
type WorkspaceSandbox struct {
	Type     string `json:"type,omitempty"`
	ParentID string `json:"parentId,omitempty"`
}

// WorkspaceRelationships represents the relationships of a workspace
type WorkspaceRelationships struct {
	Entitlement []WorkspaceEntitlement `json:"entitlement,omitempty"`
	Policy      []WorkspacePolicy      `json:"policy,omitempty"`
	Feature     []WorkspaceFeature     `json:"feature,omitempty"`
}

// WorkspacePolicy represents a policy in a workspace
type WorkspacePolicy struct {
	ID         string               `json:"id,omitempty"`
	Type       string               `json:"type,omitempty"`
	Links      WorkspaceLinks       `json:"links,omitempty"`
	Attributes WorkspacePolicyAttrs `json:"attributes,omitempty"`
}

// WorkspacePolicyAttrs represents the attributes of a workspace policy
type WorkspacePolicyAttrs struct {
	Type      string `json:"type,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Suspended string `json:"suspended,omitempty"`
}

// WorkspaceFeature represents a feature in a workspace
type WorkspaceFeature struct {
	ID         string                `json:"id,omitempty"`
	Type       string                `json:"type,omitempty"`
	Links      WorkspaceLinks        `json:"links,omitempty"`
	Attributes WorkspaceFeatureAttrs `json:"attributes,omitempty"`
}

// WorkspaceFeatureAttrs represents the attributes of a workspace feature
type WorkspaceFeatureAttrs struct {
	Type                   string   `json:"type,omitempty"`
	AllInclusive           bool     `json:"allInclusive,omitempty"`
	Events                 []string `json:"events,omitempty"`
	Available              bool     `json:"available,omitempty"`
	Limit                  int      `json:"limit,omitempty"`
	EntitledSandbox        string   `json:"entitledSandbox,omitempty"`
	IP                     int      `json:"ip,omitempty"`
	Portal                 *Limit   `json:"portal,omitempty"`
	Parent                 *Limit   `json:"parent,omitempty"`
	Self                   *Limit   `json:"self,omitempty"`
	Realms                 []string `json:"realms,omitempty"`
	IsDataResidencyAllowed bool     `json:"isDataResidencyAllowed,omitempty"`
	Tracks                 []string `json:"tracks,omitempty"`
}

// Limit represents a limit configuration
type Limit struct {
	Limit int `json:"limit,omitempty"`
}

// WorkspaceEntitlement represents an entitlement in a workspace
type WorkspaceEntitlement struct {
	ID         string                    `json:"id,omitempty"`
	Type       string                    `json:"type,omitempty"`
	Links      WorkspaceEntitlementLinks `json:"links,omitempty"`
	Attributes WorkspaceEntitlementAttrs `json:"attributes,omitempty"`
}

// WorkspaceEntitlementLinks represents the links in a workspace entitlement
type WorkspaceEntitlementLinks struct {
	Self string `json:"self,omitempty"`
}

// WorkspaceEntitlementAttrs represents the attributes of a workspace entitlement
type WorkspaceEntitlementAttrs struct {
	PlanKey string `json:"planKey,omitempty"`
	Plan    string `json:"plan,omitempty"`
	Key     string `json:"key,omitempty"`
}
