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
}

// Get returns all content in a Confluence instance.
func (c *ContentService) Get(ctx context.Context, options *GetContentOptionsScheme, startAt, maxResults int) (result *ContentPageScheme, response *ResponseScheme, err error) {

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
		return nil, nil, err
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
		return nil, nil, err
	}

	return
}
