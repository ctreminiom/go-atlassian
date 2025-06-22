package internal

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
	"net/url"
	"strconv"
)

// NewTaskService creates a new instance of TaskService.
// It takes a service.Connector as input and returns a pointer to TaskService.
func NewTaskService(client service.Connector) *TaskService {
	return &TaskService{
		internalClient: &internalTaskImpl{c: client},
	}
}

// TaskService provides methods to interact with long-running task operations in Confluence.
type TaskService struct {
	// internalClient is the connector interface for task operations.
	internalClient confluence.TaskConnector
}

// Gets returns information about all active long-running tasks (e.g. space export),
//
// such as how long each task has been running and the percentage of each task that has completed.
//
// GET /wiki/rest/api/longtask
//
// https://docs.go-atlassian.io/confluence-cloud/long-task#get-long-running-tasks
func (t *TaskService) Gets(ctx context.Context, start, limit int) (*model.LongTaskPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TaskService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	return t.internalClient.Gets(ctx, start, limit)
}

// Get returns information about an active long-running task (e.g. space export), such as how long it has been running
//
// and the percentage of the task that has completed.
//
// GET /wiki/rest/api/longtask/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/long-task#get-long-running-task
func (t *TaskService) Get(ctx context.Context, taskID string) (*model.LongTaskScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TaskService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	return t.internalClient.Get(ctx, taskID)
}

type internalTaskImpl struct {
	c service.Connector
}

func (i *internalTaskImpl) Gets(ctx context.Context, start, limit int) (*model.LongTaskPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTaskImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	query := url.Values{}
	query.Add("start", strconv.Itoa(start))
	query.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("wiki/rest/api/longtask?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.LongTaskPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalTaskImpl) Get(ctx context.Context, taskID string) (*model.LongTaskScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTaskImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	endpoint := fmt.Sprintf("wiki/rest/api/longtask/%v", taskID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	task := new(model.LongTaskScheme)
	response, err := i.c.Call(request, task)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return task, response, nil
}
