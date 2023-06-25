package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/assets"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewAQLService(client service.Client) *AQLService {

	return &AQLService{
		internalClient: &internalAQLImpl{c: client},
	}
}

type AQLService struct {
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
	return a.internalClient.Filter(ctx, workspaceID, parameters)
}

type internalAQLImpl struct {
	c service.Client
}

func (i *internalAQLImpl) Filter(ctx context.Context, workspaceID string, parameters *model.AQLSearchParamsScheme) (*model.ObjectListScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	object := new(model.ObjectListScheme)
	response, err := i.c.Call(request, object)
	if err != nil {
		return nil, response, err
	}

	return object, response, nil
}
