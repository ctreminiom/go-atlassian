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
