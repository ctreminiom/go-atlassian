package sm

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type KnowledgeBaseService struct{ client *Client }

// Search returns articles which match the given query string across all service desks.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/knowledgebase#search-articles
func (k *KnowledgeBaseService) Search(ctx context.Context, query string, highlight bool, start, limit int) (
	result *models.ArticlePageScheme, response *ResponseScheme, err error) {

	if len(query) == 0 {
		return nil, nil, models.ErrNoKBQueryError
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
func (k *KnowledgeBaseService) Gets(ctx context.Context, serviceDeskID int, query string, highlight bool, start,
	limit int) (result *models.ArticlePageScheme, response *ResponseScheme, err error) {

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
