package models

import "time"

type GetContentOptionsScheme struct {
	ContextType, SpaceKey string
	Title                 string
	Trigger               string
	OrderBy               string
	Status, Expand        []string
	PostingDay            time.Time
}

type ContentPageScheme struct {
	Results []*ContentScheme `json:"results"`
	Start   int              `json:"start"`
	Limit   int              `json:"limit"`
	Size    int              `json:"size"`
	Links   *LinkScheme      `json:"_links"`
}

type LinkScheme struct {
	Base       string `json:"base,omitempty"`
	Context    string `json:"context,omitempty"`
	Self       string `json:"self,omitempty"`
	Tinyui     string `json:"tinyui,omitempty"`
	Editui     string `json:"editui,omitempty"`
	Webui      string `json:"webui,omitempty"`
	Download   string `json:"download,omitempty"`
	Next       string `json:"next"`
	Collection string `json:"collection"`
}

type ContentScheme struct {
	ID         string                  `json:"id,omitempty"`
	Type       string                  `json:"type,omitempty"`
	Status     string                  `json:"status,omitempty"`
	Title      string                  `json:"title,omitempty"`
	Expandable *ExpandableScheme       `json:"_expandable,omitempty"`
	Links      *LinkScheme             `json:"_links,omitempty"`
	ChildTypes *ChildTypesScheme       `json:"childTypes,omitempty"`
	Space      *SpaceScheme            `json:"space,omitempty"`
	Metadata   *MetadataScheme         `json:"metadata,omitempty"`
	Operations []*OperationScheme      `json:"operations,omitempty"`
	Body       *BodyScheme             `json:"body,omitempty"`
	Version    *ContentVersionScheme   `json:"version,omitempty"`
	Extensions *ContentExtensionScheme `json:"extensions,omitempty"`
	Ancestors  []*ContentScheme        `json:"ancestors,omitempty"`
	History    *ContentHistoryScheme   `json:"history,omitempty"`
}

type ContentExtensionScheme struct {
	MediaType            string `json:"mediaType,omitempty"`
	FileSize             int    `json:"fileSize,omitempty"`
	Comment              string `json:"comment,omitempty"`
	MediaTypeDescription string `json:"mediaTypeDescription,omitempty"`
	FileID               string `json:"fileId,omitempty"`
}

type BodyScheme struct {
	View                *BodyNodeScheme `json:"view,omitempty"`
	ExportView          *BodyNodeScheme `json:"export_view,omitempty"`
	StyledView          *BodyNodeScheme `json:"styled_view,omitempty"`
	Storage             *BodyNodeScheme `json:"storage,omitempty"`
	Editor2             *BodyNodeScheme `json:"editor2,omitempty"`
	AnonymousExportView *BodyNodeScheme `json:"anonymous_export_view,omitempty"`
}

type BodyNodeScheme struct {
	Value          string `json:"value,omitempty"`
	Representation string `json:"representation,omitempty"`
}

type OperationScheme struct {
	Operation  string `json:"operation,omitempty"`
	TargetType string `json:"targetType,omitempty"`
}

type MetadataScheme struct {
	Labels     *LabelsScheme     `json:"labels"`
	Expandable *ExpandableScheme `json:"_expandable,omitempty"`
	MediaType  string            `json:"mediaType,omitempty"`
}

type LabelsScheme struct {
	Results []*LabelValueScheme `json:"results,omitempty"`
	Start   int                 `json:"start,omitempty"`
	Limit   int                 `json:"limit,omitempty"`
	Size    int                 `json:"size,omitempty"`
	Links   *LinkScheme         `json:"_links,omitempty"`
}

type LabelValueScheme struct {
	Prefix string `json:"prefix,omitempty"`
	Name   string `json:"name,omitempty"`
	ID     string `json:"id,omitempty"`
	Label  string `json:"label,omitempty"`
}

type ChildTypesScheme struct {
	Attachment *ChildTypeScheme `json:"attachment,omitempty"`
	Comment    *ChildTypeScheme `json:"comment,omitempty"`
	Page       *ChildTypeScheme `json:"page,omitempty"`
}

type ChildTypeScheme struct {
	Value bool `json:"value,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type ExpandableScheme struct {
	Container           string `json:"container"`
	Metadata            string `json:"metadata"`
	Restrictions        string `json:"restrictions"`
	History             string `json:"history"`
	Body                string `json:"body"`
	Version             string `json:"version"`
	Descendants         string `json:"descendants"`
	Space               string `json:"space"`
	ChildTypes          string `json:"childTypes"`
	Operations          string `json:"operations"`
	SchedulePublishDate string `json:"schedulePublishDate"`
	Children            string `json:"children"`
	Ancestors           string `json:"ancestors"`
	Settings            string `json:"settings"`
	LookAndFeel         string `json:"lookAndFeel"`
	Identifiers         string `json:"identifiers"`
	Permissions         string `json:"permissions"`
	Icon                string `json:"icon"`
	Description         string `json:"description"`
	Theme               string `json:"theme"`
	Homepage            string `json:"homepage"`
	LastUpdated         string `json:"lastUpdated"`
	PreviousVersion     string `json:"previousVersion"`
	Contributors        string `json:"contributors"`
	NextVersion         string `json:"nextVersion"`
}

type ContentHistoryScheme struct {
	Latest          bool                              `json:"latest,omitempty"`
	CreatedBy       *ContentUserScheme                `json:"createdBy,omitempty"`
	CreatedDate     string                            `json:"createdDate,omitempty"`
	LastUpdated     *ContentVersionScheme             `json:"lastUpdated,omitempty"`
	PreviousVersion *ContentVersionScheme             `json:"previousVersion,omitempty"`
	Contributors    *ContentHistoryContributorsScheme `json:"contributors,omitempty"`
	NextVersion     *ContentVersionScheme             `json:"nextVersion,omitempty"`
	Expandable      *ExpandableScheme                 `json:"_expandable,omitempty"`
	Links           *LinkScheme                       `json:"_links,omitempty"`
}

type ContentHistoryContributorsScheme struct {
	Publishers *VersionCollaboratorsScheme `json:"publishers,omitempty"`
}

type ContentUserScheme struct {
	Type           string                   `json:"type,omitempty"`
	Username       string                   `json:"username,omitempty"`
	UserKey        string                   `json:"userKey,omitempty"`
	AccountID      string                   `json:"accountId,omitempty"`
	AccountType    string                   `json:"accountType,omitempty"`
	Email          string                   `json:"email,omitempty"`
	PublicName     string                   `json:"publicName,omitempty"`
	ProfilePicture *ProfilePictureScheme    `json:"profilePicture,omitempty"`
	DisplayName    string                   `json:"displayName,omitempty"`
	Operations     []*OperationScheme       `json:"operations,omitempty"`
	Details        *ContentUserDetailScheme `json:"details,omitempty"`
	PersonalSpace  *SpaceScheme             `json:"personalSpace,omitempty"`
	Expandable     *ExpandableScheme        `json:"_expandable,omitempty"`
	Links          *LinkScheme              `json:"_links,omitempty"`
}

type ProfilePictureScheme struct {
	Path      string `json:"path,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	IsDefault bool   `json:"isDefault,omitempty"`
}

type ContentUserDetailScheme struct {
	Business *UserBusinessDetailScheme `json:"business,omitempty"`
	Personal *UserPersonalDetailScheme `json:"personal,omitempty"`
}

type UserBusinessDetailScheme struct {
	Position   string `json:"position,omitempty"`
	Department string `json:"department,omitempty"`
	Location   string `json:"location,omitempty"`
}

type UserPersonalDetailScheme struct {
	Phone   string `json:"phone,omitempty"`
	Im      string `json:"im,omitempty"`
	Website string `json:"website,omitempty"`
	Email   string `json:"email,omitempty"`
}

type ContentArchivePayloadScheme struct {
	Pages []*ContentArchiveIDPayloadScheme `json:"pages,omitempty"`
}

type ContentArchiveIDPayloadScheme struct {
	ID int `json:"id,omitempty"`
}

type ContentArchiveResultScheme struct {
	ID    string `json:"id"`
	Links struct {
		Status string `json:"status"`
	} `json:"links"`
}
