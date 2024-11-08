package service

import (
	"context"
	"net/http"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type Connector interface {
	NewRequest(ctx context.Context, method, urlStr, contentType string, body interface{}) (*http.Request, error)
	Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error)
}
