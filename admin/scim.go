package admin

type SCIMService struct {
	client   *Client
	User     *SCIMUserService
	Group    *SCIMGroupService
	Scheme   *SCIMSchemeService
	Resource *SCIMResourceService
}
