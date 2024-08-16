package internal

import (
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
)

// NewLinkService creates new instances of LinkADFService and LinkRichTextService.
func NewLinkService(client service.Connector, version string, type_ *LinkTypeService, remote *RemoteLinkService) (*LinkADFService, *LinkRichTextService, error) {

	if version == "" {
		return nil, nil, model.ErrNoVersionProvided
	}

	adfService := &LinkADFService{
		internalClient: &internalLinkADFServiceImpl{
			c:       client,
			version: version,
		},
		Type:   type_,
		Remote: remote,
	}

	richTextService := &LinkRichTextService{
		internalClient: &internalLinkRichTextServiceImpl{
			c:       client,
			version: version,
		},
		Type:   type_,
		Remote: remote,
	}

	return adfService, richTextService, nil
}
