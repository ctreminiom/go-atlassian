package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/assets"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewAQLService creates a new instance of AQLService.
// It takes a service.Connector as input and returns a pointer to AQLService.
func NewAQLService(client service.Connector) *AQLService {
	return &AQLService{
		internalClient: &internalAQLImpl{c: client},
	}
}

// AQLService provides methods to interact with the Assets Query Language (AQL) in Jira.
type AQLService struct {
	// internalClient is the connector interface for AQL operations.
	internalClient assets.AQLAssetConnector
}

// Filter search objects based on Assets Query Language (AQL)
//
// POST /jsm/assets/workspace/{workspaceId}/v1/aql/objects
//
// Deprecated. Please use Object.Filter() instead.
//
// https://docs.go-atlassian.io/jira-assets/aql#filter-objects
func (a *AQLService) Filter(ctx context.Context, workspaceID string, parameters *model.AQLSearchParamsScheme) (*model.ObjectListScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*AQLService).Filter")
	defer span.End()

	return a.internalClient.Filter(ctx, workspaceID, parameters)
}

type internalAQLImpl struct {
	c service.Connector
}

func (i *internalAQLImpl) Filter(ctx context.Context, workspaceID string, parameters *model.AQLSearchParamsScheme) (*model.ObjectListScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalAQLImpl).Filter")
	defer span.End()

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("jsm/assets/workspace/%v/v1/aql/objects", workspaceID))

	if parameters != nil {

		query := url.Values{}
		query.Add("qlQuery", parameters.Query)
		query.Add("page", strconv.Itoa(parameters.Page))
		query.Add("resultPerPage", strconv.Itoa(parameters.ResultPerPage))
		query.Add("includeAttributes", fmt.Sprintf("%v", parameters.IncludeAttributes))
		query.Add("includeAttributesDeep", fmt.Sprintf("%v", parameters.IncludeAttributesDeep))
		query.Add("includeTypeAttributes", fmt.Sprintf("%v", parameters.IncludeTypeAttributes))
		query.Add("includeExtendedInfo", fmt.Sprintf("%v", parameters.IncludeExtendedInfo))

		if query.Encode() != "" {
			endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
		}
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	object := new(model.ObjectListScheme)
	res, err := i.c.Call(req, object)
	if err != nil {
		return nil, res, err
	}

	return object, res, nil
}
