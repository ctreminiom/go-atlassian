package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
	"net/http"
	"net/url"
	"strconv"
)

func NewKnowledgebaseService(client service.Connector, version string) *KnowledgebaseService {

	return &KnowledgebaseService{
		internalClient: &internalKnowledgebaseImpl{c: client, version: version},
	}
}

type KnowledgebaseService struct {
	internalClient sm.KnowledgeBaseConnector
}

// Search returns articles which match the given query string across all service desks.
//
// GET /rest/servicedeskapi/knowledgebase/article
//
// https://docs.go-atlassian.io/jira-service-management-cloud/knowledgebase#search-articles
func (k *KnowledgebaseService) Search(ctx context.Context, query string, highlight bool, start, limit int) (*model.ArticlePageScheme, *model.ResponseScheme, error) {
	return k.internalClient.Search(ctx, query, highlight, start, limit)
}

// Gets returns articles which match the given query string across all service desks.
//
// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/knowledgebase/article
//
// https://docs.go-atlassian.io/jira-service-management-cloud/knowledgebase#get-articles
func (k *KnowledgebaseService) Gets(ctx context.Context, serviceDeskID int, query string, highlight bool, start, limit int) (*model.ArticlePageScheme, *model.ResponseScheme, error) {
	return k.internalClient.Gets(ctx, serviceDeskID, query, highlight, start, limit)
}

type internalKnowledgebaseImpl struct {
	c       service.Connector
	version string
}

func (i *internalKnowledgebaseImpl) Search(ctx context.Context, query string, highlight bool, start, limit int) (*model.ArticlePageScheme, *model.ResponseScheme, error) {

	if query == "" {
		return nil, nil, model.ErrNoKBQueryError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))
	params.Add("query", query)
	params.Add("highlight", fmt.Sprintf("%v", highlight))

	endpoint := fmt.Sprintf("rest/servicedeskapi/knowledgebase/article?%v", params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ArticlePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalKnowledgebaseImpl) Gets(ctx context.Context, serviceDeskID int, query string, highlight bool, start, limit int) (*model.ArticlePageScheme, *model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskIDError
	}

	if query == "" {
		return nil, nil, model.ErrNoKBQueryError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))
	params.Add("query", query)
	params.Add("highlight", fmt.Sprintf("%v", highlight))

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/knowledgebase/article?%v", serviceDeskID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ArticlePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}
