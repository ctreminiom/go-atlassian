package internal

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
	"net/http"
)

// NewInfoService creates a new instance of InfoService.
// It takes a service.Connector and a version string as input and returns a pointer to InfoService.
func NewInfoService(client service.Connector, version string) *InfoService {

	return &InfoService{
		internalClient: &internalInfoImpl{c: client, version: version},
	}
}

// InfoService provides methods to interact with information operations in Jira Service Management.
type InfoService struct {
	// internalClient is the connector interface for information operations.
	internalClient sm.InfoConnector
}

// Get retrieves information about the Jira Service Management instance such as software version, builds, and related links.
//
// GET /rest/servicedeskapi/info
//
// https://docs.go-atlassian.io/jira-service-management-cloud/info#get-info
func (i *InfoService) Get(ctx context.Context) (*model.InfoScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*InfoService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	return i.internalClient.Get(ctx)
}

type internalInfoImpl struct {
	c       service.Connector
	version string
}

func (i *internalInfoImpl) Get(ctx context.Context) (*model.InfoScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalInfoImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	endpoint := "rest/servicedeskapi/info"

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	info := new(model.InfoScheme)
	res, err := i.c.Call(req, info)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return info, res, nil
}
