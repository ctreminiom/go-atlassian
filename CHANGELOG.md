
<a name="v1.1.5"></a>
## [v1.1.5](https://github.com/ctreminiom/go-atlassian/compare/v1.1.4...v1.1.5)

> 2021-04-23

### :bug: Bug Fixes

* replaced the Overwrite method with the Update method in the User SCIM service.

### :memo: Add or update documentation.

* Updated the CHANGELOG.md configurations
* Updated the CHANGELOG.md
* Added the CHANGELOG.md
* Updated README.md
* Updated README.md and Issue Get example
* Updated README.md
* Document the SCIM methods linking those with the Official Documentation.
* Added the codacy.com badge on the README.md

### :package: Dependencies

* removed dependencies unused on the library itself and used on the advance examples
* removed the /vendor folder and update go.mod dependencies file

### :recycle: Code Refactoring

* refactor OrganizationService and OrganizationPolicyService with the correct struct types and documented examples
* refactor Service Management Module linking the methods with the official documentation
* refactor UserService with the correct struct types and documented examples
* refactor ScreenSchemeService with the correct struct types
* refactor ScreenService with the correct struct types
* refactor ProjectRoleService with the correct struct types, and documented the examples.
* refactor ProjectPermissionSchemeService with the correct struct types, and documented the examples.
* refactor ProjectComponentService with the correct struct types, and documented the examples.
* refactor ProjectService with the correct struct types, and documented the examples.
* refactor PermissionSchemeService and PermissionGrantService with the correct struct types, increased the code coverage and documented the examples.
* refactor IssueWatcherScheme with the correct struct types and documented the examples
* refactor IssueVoteScheme with the correct struct types and documented the examples
* refactor IssueTypeScreenSchemeService with the correct struct types and documented the examples
* refactor IssueTypeSchemeService with the correct struct types and documented the examples
* refactor IssueTypeService with the correct struct types and documented the examples
* refactor IssueSearchService with the correct struct types and documented the examples
* refactor IssueSearchService with the correct struct types and documented the examples
* refactor IssueLinkTypeService with the correct struct types and documented the examples
* refactor IssueLinkService with the correct struct types and documented the examples
* refactor IssueFieldContextOptionService with the correct struct types and documented the examples
* refactor IssueFieldContextService with the correct struct types and documented the examples
* refactor IssueFieldConfigurationService with the correct struct types and documented the examples
* refactor IssueFieldService with the correct struct types and examples
* refactor IssueCommentService with the correct struct types and examples
* refactor IssueService with the correct struct types and examples
* refactor GroupService with the correct struct types
* refactor FilterShareService with the correct struct types, added more struct tags and format examples
* refactor FilterService with the correct struct types, added more struct tags and format examples
* refactor ApplicationRoleService with the correct struct types and examples
* refactor DashboardService with the correct struct types and examples

### :sparkles: Features

* Added the SCIMGroupService on the Atlassian User provisioning API

### :white_check_mark: Ad/Update Test Cases

* Increased the code coverage on the OrganizationService on the Atlassian Admin module
* Increased the code coverage on the CustomerService on the Service Management module.
* Increased the code coverage on the IssueResolutionService
* Increased the code coverage on the IssuePriorityService
* Increased the code coverage on the issueService

### Pull Requests

* Merge pull request [#29](https://github.com/ctreminiom/go-atlassian/issues/29) from ctreminiom/dev
* Merge pull request [#28](https://github.com/ctreminiom/go-atlassian/issues/28) from fossabot/add-license-scan-badge
* Merge pull request [#27](https://github.com/ctreminiom/go-atlassian/issues/27) from ctreminiom/dev
* Merge pull request [#22](https://github.com/ctreminiom/go-atlassian/issues/22) from ctreminiom/dev


<a name="v1.1.4"></a>
## [v1.1.4](https://github.com/ctreminiom/go-atlassian/compare/v1.1.3...v1.1.4)

> 2021-04-07

### :bug: Bug Fixes

* Fixed Issue Test Cases coverage issues
* Types the Lint problems

### :memo: Add or update documentation.

* updated the time-in-status example

### :sparkles: Features

* Implemented the Issue Update using operations

### :white_check_mark: Ad/Update Test Cases

* Added the Test Cases for the AuthService
* Increased the code coverage on the issueService
* Added test cases on the Operations Update

### Pull Requests

* Merge pull request [#17](https://github.com/ctreminiom/go-atlassian/issues/17) from ctreminiom/dev


<a name="v1.1.3"></a>
## [v1.1.3](https://github.com/ctreminiom/go-atlassian/compare/v1.1.2...v1.1.3)

> 2021-04-07

### :memo: Add or update documentation.

* added the issue labels examples
* added the issue fields examples
* added the time-in-status example and updated the dependencies,
* added advanced examples like "add field to project" or extract issue changelogs

### :sparkles: Features

* Added the SCIM User Service
* Added the SCIM Scheme Service
* Added the user SCIM service with the User CRUD endpoints mapped.
* added the Atlassian User Admin service with examples and test cases
* mapped the /rest/api/3/issuetypescreenscheme/project and /rest/api/3/issuetypescreenscheme/mapping endpoints

### :white_check_mark: Ad/Update Test Cases

* added more test cases on the admin and auth services
* added the test cases on the OrganizationPolicyService
* added more test cases on the OrganizationService

### Pull Requests

* Merge pull request [#16](https://github.com/ctreminiom/go-atlassian/issues/16) from ctreminiom/feature/cloud-admin
* Merge pull request [#15](https://github.com/ctreminiom/go-atlassian/issues/15) from ctreminiom/dev
* Merge pull request [#13](https://github.com/ctreminiom/go-atlassian/issues/13) from ctreminiom/dev


<a name="v1.1.2"></a>
## [v1.1.2](https://github.com/ctreminiom/go-atlassian/compare/v1.1.1...v1.1.2)

> 2021-03-24

### :memo: Add or update documentation.

* added examples

### Pull Requests

* Merge pull request [#11](https://github.com/ctreminiom/go-atlassian/issues/11) from ctreminiom/feature/jira-software-example


<a name="v1.1.1"></a>
## [v1.1.1](https://github.com/ctreminiom/go-atlassian/compare/v1.0.2...v1.1.1)

> 2021-03-24

### :art: Improve structure / format of the code.

* Updated README.md
* Updated README.md
* Format the badges
* Added the new Library Logo

### :white_check_mark: Ad/Update Test Cases

* Increased the coverage adding more Test on the issueComment.go

### Pull Requests

* Merge pull request [#8](https://github.com/ctreminiom/go-atlassian/issues/8) from ctreminiom/feature-adf-comments
* Merge pull request [#6](https://github.com/ctreminiom/go-atlassian/issues/6) from ctreminiom/feat/updated-readme.md
* Merge pull request [#5](https://github.com/ctreminiom/go-atlassian/issues/5) from ctreminiom/feature/jira-service-management


<a name="v1.0.2"></a>
## [v1.0.2](https://github.com/ctreminiom/go-atlassian/compare/v1.1.0...v1.0.2)

> 2021-03-23


<a name="v1.1.0"></a>
## [v1.1.0](https://github.com/ctreminiom/go-atlassian/compare/v1.0.1...v1.1.0)

> 2021-03-23

### :art: Improve structure / format of the code.

* Improved the code samples folder

### :sparkles: Features

* Added the RequestTypeService on the Service Management Module
* Added the ServiceDeskQueueService on the Service Management Module
* Added the Project Get method on the KnowledgebaseService module on the Service Management Module
* Added the Add/Remove methods on the CustomerService module on the Service Management Module
* Added the ServiceDeskProjectService on the Service Management Module
* Added the RequestFeedbackService on the Service Management Module
* Added the RequestSLAService on the Service Management Module
* Added the RequestParticipantService on the Service Management Module
* Added the Subscribe and Unsubscribe methods on the RequestService module.
* Added the RequestCommentService on the Service Management Module
* Added the RequestAttachmentService on the Service Management Module
* Added the RequestService, RequestApprovalService and RequestTypeService

### :white_check_mark: Ad/Update Test Cases

* Added the Unit Test cases on the KnowledgebaseService and OrganizationService
* Added the Unit Test cases on the customerService

### Ambulance

* fixed the GoLint non-used warning on the Service Management code samples

### Construction

* Created the first Jira Service Management services and linked it into the JiraService struct.

### Pull Requests

* Merge pull request [#5](https://github.com/ctreminiom/go-atlassian/issues/5) from ctreminiom/feature/jira-service-management


<a name="v1.0.1"></a>
## [v1.0.1](https://github.com/ctreminiom/go-atlassian/compare/v1.0.0...v1.0.1)

> 2021-03-04

### :bug: Bug Fixes

* Closes [#2](https://github.com/ctreminiom/go-atlassian/issues/2)

### Pull Requests

* Merge pull request [#3](https://github.com/ctreminiom/go-atlassian/issues/3) from ctreminiom/issue-2-unmarshal_array_error_on_IssueSearchService


<a name="v1.0.0"></a>
## v1.0.0

> 2021-03-03

### :art: Improve structure / format of the code.

* added the Jira Date Format constant

### :bug: Bug Fixes

* Fixed the missing request headers blocking the HTTP callback returning a 415 and updated the comments documentation.

### :memo: Add or update documentation.

* updated license source link
* Added the PULL_REQUEST_TEMPLATE.md
* updated the bug_report,md
*  Created CONTRIBUTING.md
*  Update issue templates
* Create CODE_OF_CONDUCT.md
* updated the comment documentation.
* updated the comment documentation.
* updated the comment documentation.
* updated the comment documentation.
* updated the comment documentation.
* updated the comment documentation validate non empty parameters values
* linked the documentation docs.go-atlassian.io with the GroupService
* linked the documentation docs.go-atlassian.io with the FilterShareService
* linked the documentation docs.go-atlassian.io with the FilterService.
* added the mergo golang module and updated the dependencies.

### :recycle: Code Refactoring

* Added the Unit Test cases with a 100% of coverage on the issueComment.go file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the audit.go file and the .json mock files needed to run the tests.

### :sparkles: Features

* updated Readme.md
* Added more test case on the issueTypeScreenScheme.go
* handled the non resolutionID or priorityID params.
* added more method in the dashboardService

### :white_check_mark: Ad/Update Test Cases

* Added the Unit Test cases with a 100% of coverage on the userSearch.go file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the projectVersion.go file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the screenSchemes.go file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the screenTab.go  file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the screen.go  file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the user.go file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the projectTypes.go file and the .json mock files needed to run the tests.

### Ambulance

* Fixed empty key param on the applicationRole.go and updated the code examples

### Construction

* updated the /vendor folder
* updated the /vendor folder

### Pull Requests

* Merge pull request [#1](https://github.com/ctreminiom/go-atlassian/issues/1) from ctreminiom/dev

