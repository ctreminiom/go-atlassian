package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/confluence"
	"net/http"
	"net/url"
	"strconv"
)

func NewLabelService(client service.Client) *LabelService {

	return &LabelService{
		internalClient: &internalLabelImpl{c: client},
	}
}

type LabelService struct {
	internalClient confluence.LabelConnector
}

// Get returns label information and a list of contents associated with the label.
//
// GET /wiki/rest/api/label
//
// https://docs.go-atlassian.io/confluence-cloud/label#get-label-information
func (l *LabelService) Get(ctx context.Context, labelName, labelType string, start, limit int) (*model.LabelDetailsScheme, *model.ResponseScheme, error) {
	return l.internalClient.Get(ctx, labelName, labelType, start, limit)
}

type internalLabelImpl struct {
	c service.Client
}

func (i *internalLabelImpl) Get(ctx context.Context, labelName, labelType string, start, limit int) (*model.LabelDetailsScheme, *model.ResponseScheme, error) {

	if labelName == "" {
		return nil, nil, model.ErrNoLabelNameError
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(start))
	query.Add("limit", strconv.Itoa(limit))
	query.Add("name", labelName)
	query.Add("type", labelType)

	endpoint := fmt.Sprintf("wiki/rest/api/label?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	details := new(model.LabelDetailsScheme)
	response, err := i.c.Call(request, details)
	if err != nil {
		return nil, response, err
	}

	return details, response, nil
}
