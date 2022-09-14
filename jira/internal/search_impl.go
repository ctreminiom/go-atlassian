package internal

import (
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
)

func NewSearchService(client service.Client, version string) (*SearchADFService, *SearchRichTextService, error) {

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
