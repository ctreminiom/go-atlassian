package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
	"net/http"
	"net/url"
	"strconv"
)

// NewKnowledgebaseService creates a new instance of KnowledgebaseService.
// It takes a service.Connector and a version string as input and returns a pointer to KnowledgebaseService.
func NewKnowledgebaseService(client service.Connector, version string) *KnowledgebaseService {
	return &KnowledgebaseService{
		internalClient: &internalKnowledgebaseImpl{c: client, version: version},
	}
}

// KnowledgebaseService provides methods to interact with knowledge base operations in Jira Service Management.
type KnowledgebaseService struct {
	// internalClient is the connector interface for knowledge base operations.
	internalClient sm.KnowledgeBaseConnector
}

// Search returns articles which match the given query string across all service desks.
//
// GET /rest/servicedeskapi/knowledgebase/article
//
// https://docs.go-atlassian.io/jira-service-management-cloud/knowledgebase#search-articles
func (k *KnowledgebaseService) Search(ctx context.Context, query string, highlight bool, start, limit int) (*model.ArticlePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*KnowledgebaseService).Search")
	defer span.End()

	return k.internalClient.Search(ctx, query, highlight, start, limit)
}

// Gets returns articles which match the given query string across all service desks.
//
// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/knowledgebase/article
//
// https://docs.go-atlassian.io/jira-service-management-cloud/knowledgebase#get-articles
func (k *KnowledgebaseService) Gets(ctx context.Context, serviceDeskID int, query string, highlight bool, start, limit int) (*model.ArticlePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*KnowledgebaseService).Gets")
	defer span.End()

	return k.internalClient.Gets(ctx, serviceDeskID, query, highlight, start, limit)
}

type internalKnowledgebaseImpl struct {
	c       service.Connector
	version string
}

func (i *internalKnowledgebaseImpl) Search(ctx context.Context, query string, highlight bool, start, limit int) (*model.ArticlePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalKnowledgebaseImpl).Search")
	defer span.End()

	if query == "" {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoKBQuery)
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
	ctx, span := tracer().Start(ctx, "(*internalKnowledgebaseImpl).Gets")
	defer span.End()

	if serviceDeskID == 0 {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoServiceDeskID)
	}

	if query == "" {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoKBQuery)
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
