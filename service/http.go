package service

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"io"
	"net/http"
)

type Client interface {
	NewRequest(ctx context.Context, method, apiEndpoint string, payload io.Reader) (*http.Request, error)
	NewFormRequest(ctx context.Context, method, apiEndpoint, contentType string, payload io.Reader) (*http.Request, error)
	Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error)
	TransformTheHTTPResponse(response *http.Response, structure interface{}) (*models.ResponseScheme, error)
	TransformStructToReader(structure interface{}) (io.Reader, error)
}
