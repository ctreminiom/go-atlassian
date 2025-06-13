package common

import (
	"context"
	"io"
	"net/http"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type Client interface {
	NewJSONRequest(ctx context.Context, method, apiEndpoint string, payload io.Reader) (*http.Request, error)
	NewRequest(ctx context.Context, method, apiEndpoint string, payload io.Reader) (*http.Request, error)
	NewFormRequest(ctx context.Context, method, apiEndpoint, formDataContentType string, payload io.Reader) (*http.Request, error)
	Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error)
	TransformTheHTTPResponse(response *http.Response, structure interface{}) (*models.ResponseScheme, error)
	TransformStructToReader(structure interface{}) (io.Reader, error)
}
