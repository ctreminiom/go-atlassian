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
}

type ContentScheme struct {
	ID         string            `json:"id,omitempty"`
	Type       string            `json:"type,omitempty"`
	Status     string            `json:"status,omitempty"`
	Title      string            `json:"title,omitempty"`
	Expandable *ExpandableScheme `json:"_expandable,omitempty"`
	Links      *LinkScheme       `json:"_links,omitempty"`
	ChildTypes *ChildTypesScheme `json:"childTypes,omitempty"`
	Space      *SpaceScheme      `json:"space,omitempty"`
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
