package internal

// SCIMService encapsulates various SCIM-related services within a single structure.
// It provides a convenient way to access and manage different SCIM functionalities.
type SCIMService struct {
	// User is a pointer to an instance of SCIMUserService.
	// It handles operations related to SCIM users.
	User *SCIMUserService

	// Group is a pointer to an instance of SCIMGroupService.
	// It manages SCIM group operations.
	Group *SCIMGroupService

	// Schema is a pointer to an instance of SCIMSchemaService.
	// It deals with SCIM schema operations.
	Schema *SCIMSchemaService
}
