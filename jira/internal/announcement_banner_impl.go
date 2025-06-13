package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewAnnouncementBannerService creates a new instance of AnnouncementBannerService.
// It takes a service.Connector and a version string as input and returns a pointer to AnnouncementBannerService.
func NewAnnouncementBannerService(client service.Connector, version string) *AnnouncementBannerService {

	return &AnnouncementBannerService{
		internalClient: &internalAnnouncementBannerImpl{c: client, version: version},
	}
}

// AnnouncementBannerService provides methods to interact with announcement banner operations in Jira Service Management.
type AnnouncementBannerService struct {
	// internalClient is the connector interface for announcement banner operations.
	internalClient jira.AnnouncementBannerConnector
}

// Get returns the current announcement banner configuration.
//
// GET /rest/api/{2-3}/announcementBanner
//
// https://docs.go-atlassian.io/jira-software-cloud/announcement-banner#get-announcement-banner-configuration
func (a *AnnouncementBannerService) Get(ctx context.Context) (*model.AnnouncementBannerScheme, *model.ResponseScheme, error) {
	return a.internalClient.Get(ctx)
}

// Update updates the announcement banner configuration.
//
// PUT /rest/api/{2-3}/announcementBanner
//
// https://docs.go-atlassian.io/jira-software-cloud/announcement-banner#get-announcement-banner-configuration
func (a *AnnouncementBannerService) Update(ctx context.Context, payload *model.AnnouncementBannerPayloadScheme) (*model.ResponseScheme, error) {
	return a.internalClient.Update(ctx, payload)
}

type internalAnnouncementBannerImpl struct {
	c       service.Connector
	version string
}

func (i *internalAnnouncementBannerImpl) Get(ctx context.Context) (*model.AnnouncementBannerScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/announcementBanner", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	banner := new(model.AnnouncementBannerScheme)
	response, err := i.c.Call(request, banner)
	if err != nil {
		return nil, response, err
	}

	return banner, response, nil
}

func (i *internalAnnouncementBannerImpl) Update(ctx context.Context, payload *model.AnnouncementBannerPayloadScheme) (*model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/announcementBanner", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
