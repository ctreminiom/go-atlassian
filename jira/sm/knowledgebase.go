package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type KnowledgebaseService struct{ client *Client }

// Returns articles which match the given query string across all service desks.
func (k *KnowledgebaseService) Articles(ctx context.Context, query string, highlight bool, start, limit int) (result *ArticlePageScheme, response *Response, err error) {

	if len(query) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid query value")
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

	response, err = k.client.Do(request)
	if err != nil {
		return
	}

	result = new(ArticlePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ArticlePageScheme struct {
	Size       int  `json:"size"`
	Start      int  `json:"start"`
	Limit      int  `json:"limit"`
	IsLastPage bool `json:"isLastPage"`
	Values     []struct {
		Title   string `json:"title"`
		Excerpt string `json:"excerpt"`
		Source  struct {
			Type string `json:"type"`
		} `json:"source"`
		Content struct {
			IframeSrc string `json:"iframeSrc"`
		} `json:"content"`
	} `json:"values"`
	Expands []string `json:"_expands"`
	Links   struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}
