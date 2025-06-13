package sm

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type KnowledgeBaseConnector interface {

	// Search returns articles which match the given query string across all service desks.
	//
	// GET /rest/servicedeskapi/knowledgebase/article
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/knowledgebase#search-articles
	Search(ctx context.Context, query string, highlight bool, start, limit int) (*model.ArticlePageScheme, *model.ResponseScheme, error)

	// Gets returns articles which match the given query string across all service desks.
	//
	// GET /rest/servicedeskapi/servicedesk/{serviceDeskID}/knowledgebase/article
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/knowledgebase#get-articles
	Gets(ctx context.Context, serviceDeskID int, query string, highlight bool, start, limit int) (*model.ArticlePageScheme, *model.ResponseScheme, error)
}
