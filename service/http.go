package service

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type Connector interface {
	NewRequest(ctx context.Context, method, urlStr, type_ string, body interface{}) (*http.Request, error)
	Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error)
}
