package internal

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
)

// NewParticipantService creates a new instance of ParticipantService.
// It takes a service.Connector and a version string as input and returns a pointer to ParticipantService.
func NewParticipantService(client service.Connector, version string) *ParticipantService {
	return &ParticipantService{
		internalClient: &internalServiceRequestParticipantImpl{c: client, version: version},
	}
}

// ParticipantService provides methods to interact with participant operations in Jira Service Management.
type ParticipantService struct {
	// internalClient is the connector interface for participant operations.
	internalClient sm.ParticipantConnector
}

// Gets returns a list of all the participants on a customer request.
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}/participant
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#get-request-participants
func (s *ParticipantService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ParticipantService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	return s.internalClient.Gets(ctx, issueKeyOrID, start, limit)
}

// Add adds participants to a customer request.
//
// POST /rest/servicedeskapi/request/{issueKeyOrID}/participant
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#add-request-participants
func (s *ParticipantService) Add(ctx context.Context, issueKeyOrID string, accountIDs []string) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ParticipantService).Add", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "add"))

	return s.internalClient.Add(ctx, issueKeyOrID, accountIDs)
}

// Remove removes participants from a customer request.
//
// DELETE /rest/servicedeskapi/request/{issueKeyOrID}/participant
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#remove-request-participants
func (s *ParticipantService) Remove(ctx context.Context, issueKeyOrID string, accountIDs []string) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ParticipantService).Remove", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "remove"))

	return s.internalClient.Remove(ctx, issueKeyOrID, accountIDs)
}

type internalServiceRequestParticipantImpl struct {
	c       service.Connector
	version string
}

func (i *internalServiceRequestParticipantImpl) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestParticipantImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	if issueKeyOrID == "" {

			return nil, nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/participant?%v", issueKeyOrID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.RequestParticipantPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return page, res, nil
}

func (i *internalServiceRequestParticipantImpl) Add(ctx context.Context, issueKeyOrID string, accountIDs []string) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestParticipantImpl).Add", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "add"))

	if issueKeyOrID == "" {

			return nil, nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	if len(accountIDs) == 0 {

			return nil, nil, fmt.Errorf("sm: %w", model.ErrNoAccountSlice)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/participant", issueKeyOrID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"accountIds": accountIDs})
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.RequestParticipantPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return page, res, nil
}

func (i *internalServiceRequestParticipantImpl) Remove(ctx context.Context, issueKeyOrID string, accountIDs []string) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestParticipantImpl).Remove", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "remove"))

	if issueKeyOrID == "" {

			return nil, nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	if len(accountIDs) == 0 {

			return nil, nil, fmt.Errorf("sm: %w", model.ErrNoAccountSlice)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/participant", issueKeyOrID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", map[string]interface{}{"accountIds": accountIDs})
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.RequestParticipantPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return page, res, nil
}
