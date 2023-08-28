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

func NewParticipantService(client service.Connector, version string) *ParticipantService {

	return &ParticipantService{
		internalClient: &internalServiceRequestParticipantImpl{c: client, version: version},
	}
}

type ParticipantService struct {
	internalClient sm.ParticipantConnector
}

// Gets returns a list of all the participants on a customer request.
//
// GET /rest/servicedeskapi/request/{issueIdOrKey}/participant
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#get-request-participants
func (s *ParticipantService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, issueKeyOrID, start, limit)
}

// Add adds participants to a customer request.
//
// POST /rest/servicedeskapi/request/{issueIdOrKey}/participant
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#add-request-participants
func (s *ParticipantService) Add(ctx context.Context, issueKeyOrID string, accountIDs []string) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Add(ctx, issueKeyOrID, accountIDs)
}

// Remove removes participants from a customer request.
//
// DELETE /rest/servicedeskapi/request/{issueIdOrKey}/participant
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#remove-request-participants
func (s *ParticipantService) Remove(ctx context.Context, issueKeyOrID string, accountIDs []string) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Remove(ctx, issueKeyOrID, accountIDs)
}

type internalServiceRequestParticipantImpl struct {
	c       service.Connector
	version string
}

func (i *internalServiceRequestParticipantImpl) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/participant?%v", issueKeyOrID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RequestParticipantPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalServiceRequestParticipantImpl) Add(ctx context.Context, issueKeyOrID string, accountIDs []string) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if len(accountIDs) == 0 {
		return nil, nil, model.ErrNoAccountSliceError
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/participant", issueKeyOrID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"accountIds": accountIDs})
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RequestParticipantPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalServiceRequestParticipantImpl) Remove(ctx context.Context, issueKeyOrID string, accountIDs []string) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if len(accountIDs) == 0 {
		return nil, nil, model.ErrNoAccountSliceError
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/participant", issueKeyOrID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", map[string]interface{}{"accountIds": accountIDs})
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RequestParticipantPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}
