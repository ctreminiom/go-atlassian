package sm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type KnowledgebaseService struct{ client *Client }

// Search returns articles which match the given query string across all service desks.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/knowledgebase#search-articles
func (k *KnowledgebaseService) Search(ctx context.Context, query string, highlight bool, start, limit int) (
	result *ArticlePageScheme, response *ResponseScheme, err error) {

	if len(query) == 0 {
		return nil, nil, notQueryError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))
	params.Add("query", query)

	if !highlight {
		params.Add("highlight", "false")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/knowledgebase/article?%v", params.Encode())

	request, err := k.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = k.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Gets returns articles which match the given query string across all service desks.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/knowledgebase#get-articles
func (k *KnowledgebaseService) Gets(ctx context.Context, serviceDeskID int, query string, highlight bool, start,
	limit int) (result *ArticlePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if query != "" {
		params.Add("query", query)
	}

	if !highlight {
		params.Add("highlight", "false")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/knowledgebase/article?%v", serviceDeskID, params.Encode())

	request, err := k.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = k.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

type ArticlePageScheme struct {
	Size       int                    `json:"size"`
	Start      int                    `json:"start"`
	Limit      int                    `json:"limit"`
	IsLastPage bool                   `json:"isLastPage"`
	Values     []*ArticleScheme       `json:"values"`
	Expands    []string               `json:"_expands"`
	Links      *ArticlePageLinkScheme `json:"_links"`
}

type ArticlePageLinkScheme struct {
	Self    string `json:"self"`
	Base    string `json:"base"`
	Context string `json:"context"`
	Next    string `json:"next"`
	Prev    string `json:"prev"`
}

type ArticleScheme struct {
	Title   string                `json:"title,omitempty"`
	Excerpt string                `json:"excerpt,omitempty"`
	Source  *ArticleSourceScheme  `json:"source,omitempty"`
	Content *ArticleContentScheme `json:"content,omitempty"`
}

type ArticleSourceScheme struct {
	Type string `json:"type,omitempty"`
}

type ArticleContentScheme struct {
	IframeSrc string `json:"iframeSrc,omitempty"`
}

var (
	notQueryError = fmt.Errorf("error, please provide a valid query value")
)
