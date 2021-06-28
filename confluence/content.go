package confluence

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type ContentService struct {
	client *Client
	Attachment *ContentAttachmentService
	ChildrenDescendant *ContentChildrenDescendantService
	Comment *ContentCommentService
}

type GetContentOptionsScheme struct {
	ContextType, SpaceKey   string
	Title, Trigger, OrderBy string
	Status, Expand          []string
	PostingDay              time.Time
}

type ContentPageScheme struct {
	Results []*ContentScheme `json:"results"`
	Start   int              `json:"start"`
	Limit   int              `json:"limit"`
	Size    int              `json:"size"`
	Links   *LinkScheme      `json:"_links"`
}

type LinkScheme struct {
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Self    string `json:"self,omitempty"`
	Tinyui  string `json:"tinyui,omitempty"`
	Editui  string `json:"editui,omitempty"`
	Webui   string `json:"webui,omitempty"`
	Download string `json:"download,omitempty"`
	Next    string `json:"next"`
}

type ContentScheme struct {
	ID         string             `json:"id,omitempty"`
	Type       string             `json:"type,omitempty"`
	Status     string             `json:"status,omitempty"`
	Title      string             `json:"title,omitempty"`
	Expandable *ExpandableScheme  `json:"_expandable,omitempty"`
	Links      *LinkScheme        `json:"_links,omitempty"`
	ChildTypes *ChildTypesScheme  `json:"childTypes,omitempty"`
	Space      *SpaceScheme       `json:"space,omitempty"`
	Metadata   *MetadataScheme    `json:"metadata,omitempty"`
	Operations []*OperationScheme `json:"operations,omitempty"`
	Body       *BodyScheme        `json:"body,omitempty"`
	Version    *VersionScheme     `json:"version,omitempty"`
	Extensions *ContentExtensionScheme `json:"extensions,omitempty"`
}

type ContentExtensionScheme struct {
	MediaType            string `json:"mediaType,omitempty"`
	FileSize             int    `json:"fileSize,omitempty"`
	Comment              string `json:"comment,omitempty"`
	MediaTypeDescription string `json:"mediaTypeDescription,omitempty"`
	FileID               string `json:"fileId,omitempty"`
}

type VersionScheme struct {
	By            *UserScheme                 `json:"by,omitempty"`
	Number        int                         `json:"number,omitempty"`
	When          string                      `json:"when,omitempty"`
	FriendlyWhen  string                      `json:"friendlyWhen,omitempty"`
	Message       string                      `json:"message,omitempty"`
	MinorEdit     bool                        `json:"minorEdit,omitempty"`
	Content       *ContentScheme              `json:"content,omitempty"`
	Collaborators *VersionCollaboratorsScheme `json:"collaborators,omitempty"`
	Expandable    struct {
		Content       string `json:"content"`
		Collaborators string `json:"collaborators"`
	} `json:"_expandable"`
	ContentTypeModified bool `json:"contentTypeModified"`
}

type VersionCollaboratorsScheme struct {
	Users    []*UserScheme `json:"users,omitempty"`
	UserKeys []string      `json:"userKeys,omitempty"`
}

type BodyScheme struct {
	View                *BodyNodeScheme `json:"view"`
	ExportView          *BodyNodeScheme `json:"export_view"`
	StyledView          *BodyNodeScheme `json:"styled_view"`
	Storage             *BodyNodeScheme `json:"storage"`
	Editor2             *BodyNodeScheme `json:"editor2"`
	AnonymousExportView *BodyNodeScheme `json:"anonymous_export_view"`
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
	MediaType string `json:"mediaType,omitempty"`
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

type SpaceScheme struct {
	ID         int               `json:"id"`
	Key        string            `json:"key"`
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	Status     string            `json:"status"`
	Expandable *ExpandableScheme `json:"_expandable"`
	Links      *LinkScheme       `json:"_links"`
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

// Gets returns all content in a Confluence instance.
func (c *ContentService) Gets(ctx context.Context, options *GetContentOptionsScheme, startAt, maxResults int) (result *ContentPageScheme, response *ResponseScheme, err error) {

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if options != nil {

		if options.ContextType != "" {
			query.Add("type", options.ContextType)
		}

		if options.SpaceKey != "" {
			query.Add("spaceKey", options.SpaceKey)
		}

		if options.Title != "" {
			query.Add("title", options.Title)
		}

		if options.Trigger != "" {
			query.Add("trigger", options.Trigger)
		}

		if options.OrderBy != "" {
			query.Add("orderby", options.OrderBy)
		}

		if !options.PostingDay.IsZero() {
			query.Add("postingDay", options.PostingDay.Format("2006-01-02"))
		}

		if len(options.Status) != 0 {
			query.Add("status", strings.Join(options.Status, ","))
		}

		if len(options.Expand) != 0 {
			query.Add("expand", strings.Join(options.Expand, ","))
		}

	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content?%v", query.Encode())

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Create creates a new piece of content or publishes an existing draft
// To publish a draft, add the id and status properties to the body of the request.
// Set the id to the ID of the draft and set the status to 'current'.
// When the request is sent, a new piece of content will be created and the metadata from the draft will be transferred into it.
func (c *ContentService) Create(ctx context.Context, payload *ContentScheme) (result *ContentScheme, response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = "/wiki/rest/api/content"

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Search returns the list of content that matches a Confluence Query Language (CQL) query
func (c *ContentService) Search(ctx context.Context, cql, cqlContext string, expand []string, cursor string, maxResults int) (result *ContentPageScheme, response *ResponseScheme, err error) {

	if cql == "" {
		return nil, nil, notCqlProvidedError
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(maxResults))
	query.Add("cql", cql)

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if cqlContext != "" {
		query.Add("cqlcontext", cqlContext)
	}

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/search?%v", query.Encode())

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Get returns a single piece of content, like a page or a blog post.
// By default, the following objects are expanded: space, history, version.
func (c *ContentService) Get(ctx context.Context, contentID string, expand []string, version int) (result *ContentScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	query := url.Values{}

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if version != 0 {
		query.Add("version", strconv.Itoa(version))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/content/%v", contentID))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Update updates a piece of content.
// Use this method to update the title or body of a piece of content, change the status, change the parent page, and more.
func (c *ContentService) Update(ctx context.Context, contentID string, payload *ContentScheme) (result *ContentScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v", contentID)

	request, err := c.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Delete moves a piece of content to the space's trash or purges it from the trash,
// depending on the content's type and status:
// If the content's type is page or blogpost and its status is current, it will be trashed.
// If the content's type is page or blogpost and its status is trashed, the content will be purged from the trash and deleted permanently.
// === Note, you must also set the status query parameter to trashed in your request. ===
// If the content's type is comment or attachment, it will be deleted permanently without being trashed.
func (c *ContentService) Delete(ctx context.Context, contentID, status string) (response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, notContentIDError
	}

	query := url.Values{}
	if status != "" {
		query.Add("status", status)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/content/%v", contentID))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err = c.client.Call(request, nil)
	if err != nil {
		return response, err
	}

	return
}

type ContentHistoryScheme struct {
	Latest          bool                              `json:"latest,omitempty"`
	CreatedBy       *UserScheme                       `json:"createdBy,omitempty"`
	CreatedDate     string                            `json:"createdDate,omitempty"`
	LastUpdated     *VersionScheme                    `json:"lastUpdated,omitempty"`
	PreviousVersion *VersionScheme                    `json:"previousVersion,omitempty"`
	Contributors    *ContentHistoryContributorsScheme `json:"contributors,omitempty"`
	NextVersion     *VersionScheme                    `json:"nextVersion,omitempty"`
	Expandable      *ExpandableScheme                 `json:"_expandable,omitempty"`
	Links           *LinkScheme                       `json:"_links,omitempty"`
}

type ContentHistoryContributorsScheme struct {
	Publishers *VersionCollaboratorsScheme `json:"publishers,omitempty"`
}

type UserScheme struct {
	Type           string                `json:"type,omitempty"`
	Username       string                `json:"username,omitempty"`
	UserKey        string                `json:"userKey,omitempty"`
	AccountID      string                `json:"accountId,omitempty"`
	AccountType    string                `json:"accountType,omitempty"`
	Email          string                `json:"email,omitempty"`
	PublicName     string                `json:"publicName,omitempty"`
	ProfilePicture *ProfilePictureScheme `json:"profilePicture,omitempty"`
	DisplayName    string                `json:"displayName,omitempty"`
	Operations     []*OperationScheme    `json:"operations,omitempty"`
	Details        *UserDetailScheme     `json:"details,omitempty"`
	PersonalSpace  *SpaceScheme          `json:"personalSpace,omitempty"`
	Expandable     *ExpandableScheme     `json:"_expandable,omitempty"`
	Links          *LinkScheme           `json:"_links,omitempty"`
}

type ProfilePictureScheme struct {
	Path      string `json:"path,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	IsDefault bool   `json:"isDefault,omitempty"`
}

type UserDetailScheme struct {
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

// History returns the most recent update for a piece of content.
func (c *ContentService) History(ctx context.Context, contentID string, expand []string) (result *ContentHistoryScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	query := url.Values{}
	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/content/%v/history", contentID))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

var (
	notCqlProvidedError = fmt.Errorf("error, cql parameter is required, please provide a valid value")
	notContentIDError   = fmt.Errorf("error, the content ID is required, please provide a valid value")
)
