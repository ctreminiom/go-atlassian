package models

type ContentRestrictionPageScheme struct {
	Start            int                         `json:"start,omitempty"`
	Limit            int                         `json:"limit,omitempty"`
	Size             int                         `json:"size,omitempty"`
	RestrictionsHash string                      `json:"restrictionsHash,omitempty"`
	Results          []*ContentRestrictionScheme `json:"results,omitempty"`
}

type ContentRestrictionScheme struct {
	Operation    string                              `json:"operation,omitempty"`
	Restrictions *ContentRestrictionDetailScheme     `json:"restrictions,omitempty"`
	Content      *ContentScheme                      `json:"content,omitempty"`
	Expandable   *ContentRestrictionExpandableScheme `json:"_expandable,omitempty"`
}

type ContentRestrictionExpandableScheme struct {
	Restrictions string `json:"restrictions,omitempty"`
	Content      string `json:"content,omitempty"`
}

type ContentRestrictionDetailScheme struct {
	User  *UserPermissionScheme  `json:"user,omitempty"`
	Group *GroupPermissionScheme `json:"group,omitempty"`
}

type ContentRestrictionUpdatePayloadScheme struct {
	Results []*ContentRestrictionUpdateScheme `json:"results,omitempty"`
}

type ContentRestrictionUpdateScheme struct {
	Operation    string                                     `json:"operation,omitempty"`
	Restrictions *ContentRestrictionRestrictionUpdateScheme `json:"restrictions,omitempty"`
	Content      *ContentScheme                             `json:"content,omitempty"`
}

type ContentRestrictionRestrictionUpdateScheme struct {
	Group []*SpaceGroupScheme  `json:"group,omitempty"`
	User  []*ContentUserScheme `json:"user,omitempty"`
}

type ContentRestrictionByOperationScheme struct {
	OperationType *ContentRestrictionScheme `json:"operationType,omitempty"`
}
