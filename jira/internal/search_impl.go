package internal

import (
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
)

// NewSearchService creates a new instance of SearchADFService and SearchRichTextService.
func NewSearchService(client service.Connector, version string) (*SearchADFService, *SearchRichTextService, error) {

	if version == "" {
		return nil, nil, model.ErrNoVersionProvided
	}

	rtService := &SearchRichTextService{
		internalClient: &internalSearchRichTextImpl{
			c:       client,
			version: version,
		},
	}

	adfService := &SearchADFService{
		internalClient: &internalSearchADFImpl{
			c:       client,
			version: version,
		},
	}

	return adfService, rtService, nil
}
