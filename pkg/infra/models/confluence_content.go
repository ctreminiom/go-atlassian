package models

import "time"

// GetContentOptionsScheme represents the options for getting content.
// ContextType is the type of the context.
// SpaceKey is the key of the space.
// Title is the title of the content.
// Trigger is the trigger for getting the content.
// OrderBy is the field by which to order the content.
// Status is the status of the content.
// Expand is the fields to expand in the content.
// PostingDay is the day the content was posted.
type GetContentOptionsScheme struct {
	ContextType, SpaceKey string
	Title                 string
	Trigger               string
	OrderBy               string
	Status, Expand        []string
	PostingDay            time.Time
}

// ContentPageScheme represents a page of content.
// Results is the content in the current page.
// Start is the start index of the content in the current page.
// Limit is the limit on the number of content in the current page.
// Size is the size of the content in the current page.
// Links is the links related to the current page.
type ContentPageScheme struct {
	Results []*ContentScheme `json:"results"`
	Start   int              `json:"start"`
	Limit   int              `json:"limit"`
	Size    int              `json:"size"`
	Links   *LinkScheme      `json:"_links"`
}

// LinkScheme represents a link.
// Base is the base URL of the link.
// Context is the context of the link.
// Self is the URL to the link itself.
// Tinyui is the tiny UI URL of the link.
// Editui is the edit UI URL of the link.
// Webui is the web UI URL of the link.
// Download is the download URL of the link.
// Next is the URL to the next link.
// Collection is the collection of the link.
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

// ContentScheme represents content.
// ID is the unique identifier of the content.
// Type is the type of the content.
// Status is the status of the content.
// Title is the title of the content.
// Expandable is the fields that can be expanded in the content.
// Links is the links related to the content.
// ChildTypes is the types of the children of the content.
// Space is the space of the content.
// Metadata is the metadata of the content.
// Operations is the operations on the content.
// Body is the body of the content.
// Version is the version of the content.
// Extensions is the extensions of the content.
// Ancestors is the ancestors of the content.
// History is the history of the content.
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

// ContentExtensionScheme represents an extension of content.
// MediaType is the media type of the extension.
// FileSize is the size of the file of the extension.
// Comment is the comment on the extension.
// MediaTypeDescription is the description of the media type of the extension.
// FileID is the ID of the file of the extension.
type ContentExtensionScheme struct {
	MediaType            string `json:"mediaType,omitempty"`
	FileSize             int    `json:"fileSize,omitempty"`
	Comment              string `json:"comment,omitempty"`
	MediaTypeDescription string `json:"mediaTypeDescription,omitempty"`
	FileID               string `json:"fileId,omitempty"`
}

// BodyScheme represents the body of content.
// View is the view of the body.
// ExportView is the export view of the body.
// StyledView is the styled view of the body.
// Storage is the storage of the body.
// Editor2 is the editor2 of the body.
// AnonymousExportView is the anonymous export view of the body.
type BodyScheme struct {
	View                *BodyNodeScheme `json:"view,omitempty"`
	ExportView          *BodyNodeScheme `json:"export_view,omitempty"`
	StyledView          *BodyNodeScheme `json:"styled_view,omitempty"`
	Storage             *BodyNodeScheme `json:"storage,omitempty"`
	Editor2             *BodyNodeScheme `json:"editor2,omitempty"`
	AnonymousExportView *BodyNodeScheme `json:"anonymous_export_view,omitempty"`
}

// BodyNodeScheme represents a node in the body of content.
// Value is the value of the node.
// Representation is the representation of the node.
type BodyNodeScheme struct {
	Value          string `json:"value,omitempty"`
	Representation string `json:"representation,omitempty"`
}

// OperationScheme represents an operation.
// Operation is the operation.
// TargetType is the target type of the operation.
type OperationScheme struct {
	Operation  string `json:"operation,omitempty"`
	TargetType string `json:"targetType,omitempty"`
}

// MetadataScheme represents the metadata of content.
// Labels is the labels of the metadata.
// Expandable is the fields that can be expanded in the metadata.
// MediaType is the media type of the metadata.
type MetadataScheme struct {
	Labels     *LabelsScheme     `json:"labels"`
	Expandable *ExpandableScheme `json:"_expandable,omitempty"`
	MediaType  string            `json:"mediaType,omitempty"`
}

// LabelsScheme represents labels.
// Results is the results of the labels.
// Start is the start index of the labels.
// Limit is the limit on the number of labels.
// Size is the size of the labels.
// Links is the links related to the labels.
type LabelsScheme struct {
	Results []*LabelValueScheme `json:"results,omitempty"`
	Start   int                 `json:"start,omitempty"`
	Limit   int                 `json:"limit,omitempty"`
	Size    int                 `json:"size,omitempty"`
	Links   *LinkScheme         `json:"_links,omitempty"`
}

// LabelValueScheme represents a value of a label.
// Prefix is the prefix of the label value.
// Name is the name of the label value.
// ID is the ID of the label value.
// Label is the label of the label value.
type LabelValueScheme struct {
	Prefix string `json:"prefix,omitempty"`
	Name   string `json:"name,omitempty"`
	ID     string `json:"id,omitempty"`
	Label  string `json:"label,omitempty"`
}

// ChildTypesScheme represents the types of children of content.
// Attachment is the attachment type of the child.
// Comment is the comment type of the child.
// Page is the page type of the child.
type ChildTypesScheme struct {
	Attachment *ChildTypeScheme `json:"attachment,omitempty"`
	Comment    *ChildTypeScheme `json:"comment,omitempty"`
	Page       *ChildTypeScheme `json:"page,omitempty"`
}

// ChildTypeScheme represents a type of a child of content.
// Value is the value of the child type.
// Links is the links related to the child type.
type ChildTypeScheme struct {
	Value bool `json:"value,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

// ExpandableScheme represents the fields that can be expanded in the content.
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

// ContentHistoryScheme represents the history of the content.
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

// ContentHistoryContributorsScheme represents the contributors of the content history.
// Publishers are the publishers of the content.
type ContentHistoryContributorsScheme struct {
	Publishers *VersionCollaboratorsScheme `json:"publishers,omitempty"`
}

// ContentUserScheme represents a user of the content.
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

// ProfilePictureScheme represents a profile picture.
type ProfilePictureScheme struct {
	Path      string `json:"path,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	IsDefault bool   `json:"isDefault,omitempty"`
}

// ContentUserDetailScheme represents the detailed information of a user in the content.
// Business is a pointer to a UserBusinessDetailScheme which represents the business details of the user.
// Personal is a pointer to a UserPersonalDetailScheme which represents the personal details of the user.
type ContentUserDetailScheme struct {
	Business *UserBusinessDetailScheme `json:"business,omitempty"` // The business details of the user.
	Personal *UserPersonalDetailScheme `json:"personal,omitempty"` // The personal details of the user.
}

// UserBusinessDetailScheme represents the business details of a user.
// Position is the position of the user in the business.
// Department is the department of the user in the business.
// Location is the location of the user in the business.
type UserBusinessDetailScheme struct {
	Position   string `json:"position,omitempty"`   // The position of the user in the business.
	Department string `json:"department,omitempty"` // The department of the user in the business.
	Location   string `json:"location,omitempty"`   // The location of the user in the business.
}

// UserPersonalDetailScheme represents the personal details of a user.
// Phone is the phone number of the user.
// Im is the instant messaging handle of the user.
// Website is the website of the user.
// Email is the email of the user.
type UserPersonalDetailScheme struct {
	Phone   string `json:"phone,omitempty"`   // The phone number of the user.
	Im      string `json:"im,omitempty"`      // The instant messaging handle of the user.
	Website string `json:"website,omitempty"` // The website of the user.
	Email   string `json:"email,omitempty"`   // The email of the user.
}

// ContentArchivePayloadScheme represents the payload for archiving content.
// Pages is a slice of pointers to ContentArchiveIDPayloadScheme which represents the pages to be archived.
type ContentArchivePayloadScheme struct {
	Pages []*ContentArchiveIDPayloadScheme `json:"pages,omitempty"` // The pages to be archived.
}

// ContentArchiveIDPayloadScheme represents the ID payload for archiving content.
// ID is the ID of the content to be archived.
type ContentArchiveIDPayloadScheme struct {
	ID int `json:"id,omitempty"` // The ID of the content to be archived.
}

// ContentArchiveResultScheme represents the result of archiving content.
// ID is the ID of the archived content.
// Links is a struct containing the status of the archived content.
type ContentArchiveResultScheme struct {
	ID    string `json:"id"` // The ID of the archived content.
	Links struct {
		Status string `json:"status"` // The status of the archived content.
	}
}

// ContentMoveScheme represents the scheme for moving content.
// ID is the ID of the content to be moved.
type ContentMoveScheme struct {
	ID string `json:"pageId"` // The ID of the content to be moved.
}
