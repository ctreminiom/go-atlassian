package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
	"net/url"
	"strconv"
)

// NewLabelService creates a new instance of LabelService.
func NewLabelService(client service.Connector, version string) (*LabelService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &LabelService{
		internalClient: &internalLabelServiceImpl{c: client, version: version},
	}, nil
}

// LabelService provides methods to manage labels in Jira Service Management.
type LabelService struct {
	// internalClient is the connector interface for label operations.
	internalClient jira.LabelConnector
}

// Gets returns a paginated list of labels.
//
// GET /rest/api/{2-3}/label
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/labels#get-all-labels
func (i *LabelService) Gets(ctx context.Context, startAt, maxResults int) (*model.IssueLabelsScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, startAt, maxResults)
}

type internalLabelServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalLabelServiceImpl) Gets(ctx context.Context, startAt, maxResults int) (*model.IssueLabelsScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/label?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	labels := new(model.IssueLabelsScheme)
	response, err := i.c.Call(request, labels)
	if err != nil {
		return nil, response, err
	}

	return labels, response, nil
}
