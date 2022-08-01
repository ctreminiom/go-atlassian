package internal

import (
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
)

func NewLinkService(client service.Client, version string) (*LinkADFService, *LinkRichTextService, error) {

	if version == "" {
		return nil, nil, model.ErrNoVersionProvided
	}

	adfService := &LinkADFService{
		internalClient: &internalLinkADFServiceImpl{
			c:       client,
			version: version,
		},
	}

	richTextService := &LinkRichTextService{
		internalClient: &internalLinkRichTextServiceImpl{
			c:       client,
			version: version,
		},
	}

	return adfService, richTextService, nil
}
