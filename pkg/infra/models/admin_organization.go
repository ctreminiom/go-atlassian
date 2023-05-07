package models

import "time"

type AdminOrganizationPageScheme struct {
	Data  []*OrganizationModelScheme `json:"data,omitempty"`
	Links *LinkPageModelScheme       `json:"links,omitempty"`
}

type LinkPageModelScheme struct {
	Self string `json:"self,omitempty"`
	Prev string `json:"prev,omitempty"`
	Next string `json:"next,omitempty"`
}

type OrganizationModelScheme struct {
	ID            string                          `json:"id,omitempty"`
	Type          string                          `json:"type,omitempty"`
	Attributes    *OrganizationModelAttribute     `json:"attributes,omitempty"`
	Relationships *OrganizationModelRelationships `json:"relationships,omitempty"`
	Links         *LinkSelfModelScheme            `json:"links,omitempty"`
}

type OrganizationModelAttribute struct {
	Name string `json:"name,omitempty"`
}

type OrganizationModelRelationships struct {
	Domains *OrganizationModelSchemes `json:"domains,omitempty"`
	Users   *OrganizationModelSchemes `json:"users,omitempty"`
}

type OrganizationModelSchemes struct {
	Links struct {
		Related string `json:"related,omitempty"`
	} `json:"links,omitempty"`
}

type LinkSelfModelScheme struct {
	Self string `json:"self,omitempty"`
}

type AdminOrganizationScheme struct {
	Data *OrganizationModelScheme `json:"data,omitempty"`
}

type OrganizationUserPageScheme struct {
	Data  []*AdminOrganizationUserScheme `json:"data,omitempty"`
	Links *LinkPageModelScheme           `json:"links,omitempty"`
	Meta  struct {
		Total int `json:"total,omitempty"`
	} `json:"meta,omitempty"`
}

type AdminOrganizationUserScheme struct {
	AccountID      string                           `json:"account_id,omitempty"`
	AccountType    string                           `json:"account_type,omitempty"`
	AccountStatus  string                           `json:"account_status,omitempty"`
	Name           string                           `json:"name,omitempty"`
	Picture        string                           `json:"picture,omitempty"`
	Email          string                           `json:"email,omitempty"`
	AccessBillable bool                             `json:"access_billable,omitempty"`
	LastActive     string                           `json:"last_active,omitempty"`
	ProductAccess  []*OrganizationUserProductScheme `json:"product_access,omitempty"`
	Links          *LinkSelfModelScheme             `json:"links,omitempty"`
}

type OrganizationUserProductScheme struct {
	Key        string `json:"key,omitempty"`
	Name       string `json:"name,omitempty"`
	URL        string `json:"url,omitempty"`
	LastActive string `json:"last_active,omitempty"`
}

type OrganizationDomainPageScheme struct {
	Data  []*OrganizationDomainModelScheme `json:"data,omitempty"`
	Links *LinkPageModelScheme             `json:"links,omitempty"`
}

type OrganizationDomainModelScheme struct {
	ID         string                                   `json:"id,omitempty"`
	Type       string                                   `json:"type,omitempty"`
	Attributes *OrganizationDomainModelAttributesScheme `json:"attributes,omitempty"`
	Links      *LinkSelfModelScheme                     `json:"links,omitempty"`
}

type OrganizationDomainModelAttributesScheme struct {
	Name  string                                       `json:"name,omitempty"`
	Claim *OrganizationDomainModelAttributeClaimScheme `json:"claim,omitempty"`
}

type OrganizationDomainModelAttributeClaimScheme struct {
	Type   string `json:"type,omitempty"`
	Status string `json:"status,omitempty"`
}

type OrganizationDomainScheme struct {
	Data *OrganizationDomainDataScheme `json:"data"`
}

type OrganizationDomainDataScheme struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name  string `json:"name"`
		Claim struct {
			Type   string `json:"type"`
			Status string `json:"status"`
		} `json:"claim"`
	} `json:"attributes"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}

type OrganizationEventOptScheme struct {
	Q      string    //Single query term for searching events.
	From   time.Time //The earliest date and time of the event represented as a UNIX epoch time.
	To     time.Time //The latest date and time of the event represented as a UNIX epoch time.
	Action string    //A query filter that returns events of a specific action type.
}

type OrganizationEventPageScheme struct {
	Data  []*OrganizationEventModelScheme `json:"data,omitempty"`
	Links *LinkPageModelScheme            `json:"links,omitempty"`
	Meta  struct {
		Next     string `json:"next,omitempty"`
		PageSize int    `json:"page_size,omitempty"`
	} `json:"meta,omitempty"`
}

type OrganizationEventModelScheme struct {
	ID         string                                  `json:"id,omitempty"`
	Type       string                                  `json:"type,omitempty"`
	Attributes *OrganizationEventModelAttributesScheme `json:"attributes,omitempty"`
	Links      *LinkSelfModelScheme                    `json:"links,omitempty"`
}

type OrganizationEventModelAttributesScheme struct {
	Time      string                          `json:"time,omitempty"`
	Action    string                          `json:"action,omitempty"`
	Actor     *OrganizationEventActorModel    `json:"actor,omitempty"`
	Context   []*OrganizationEventObjectModel `json:"context,omitempty"`
	Container []*OrganizationEventObjectModel `json:"container,omitempty"`
	Location  *OrganizationEventLocationModel `json:"location,omitempty"`
}

type OrganizationEventActorModel struct {
	ID    string               `json:"id,omitempty"`
	Name  string               `json:"name,omitempty"`
	Links *LinkSelfModelScheme `json:"links,omitempty"`
}

type OrganizationEventObjectModel struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
		Alt  string `json:"alt,omitempty"`
	} `json:"links,omitempty"`
}

type OrganizationEventLocationModel struct {
	IP  string `json:"ip,omitempty"`
	Geo string `json:"geo,omitempty"`
}

type OrganizationEventScheme struct {
	Data *OrganizationEventModelScheme `json:"data,omitempty"`
}

type OrganizationEventActionScheme struct {
	Data []*OrganizationEventActionModelScheme `json:"data,omitempty"`
}

type OrganizationEventActionModelScheme struct {
	ID         string                                        `json:"id,omitempty"`
	Type       string                                        `json:"type,omitempty"`
	Attributes *OrganizationEventActionModelAttributesScheme `json:"attributes,omitempty"`
}

type OrganizationEventActionModelAttributesScheme struct {
	DisplayName      string `json:"displayName,omitempty"`
	GroupDisplayName string `json:"groupDisplayName,omitempty"`
}

type UserProductAccessScheme struct {
	Data *UserProductAccessDataScheme `json:"data,omitempty"`
}

type UserProductAccessDataScheme struct {
	ProductAccess []*UserProductLastActiveScheme `json:"product_access,omitempty"`
	AddedToOrg    string                         `json:"added_to_org,omitempty"`
}

type UserProductLastActiveScheme struct {
	Id         string `json:"id,omitempty"`
	Key        string `json:"key,omitempty"`
	Name       string `json:"name,omitempty"`
	Url        string `json:"url,omitempty"`
	LastActive string `json:"last_active,omitempty"`
}

type GenericActionSuccessScheme struct {
	Message string `json:"message,omitempty"`
}
