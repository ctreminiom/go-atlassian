package internal

type SCIMService struct {
	User   *SCIMUserService
	Group  *SCIMGroupService
	Schema *SCIMSchemaService
}
