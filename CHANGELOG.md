
<a name="v1.4.2"></a>
## [v1.4.2](https://github.com/ctreminiom/go-atlassian/compare/v1.3.1...v1.4.2)

> 2022-01-07

### :bug: Bug Fixes

* Fixed the Field Config Scheme Unlink HTTP method
* Fixed the Field Config Scheme Unlink HTTP method

### :memo: Add or update documentation.

* Updated the method documentation url's.

### :sparkles: Features

* Added the Field Configuration Scheme service
* Mapped the Jira.Field.Configuration.Items.Update endpoint
* Mapped the Jira.Field.Configuration.Delete endpoint
* Mapped the Jira.Field.Configuration.Update endpoint
* Mapped the Jira.Field.Configuration.Create endpoint
* Added the Move screen tab field method

### Pull Requests

* Merge pull request [#95](https://github.com/ctreminiom/go-atlassian/issues/95) from ctreminiom/dev-1
* Merge pull request [#93](https://github.com/ctreminiom/go-atlassian/issues/93) from ctreminiom/dependabot/go_modules/github.com/tidwall/gjson-1.12.1
* Merge pull request [#94](https://github.com/ctreminiom/go-atlassian/issues/94) from ctreminiom/feature/field-screen-tab-move


<a name="v1.3.1"></a>
## [v1.3.1](https://github.com/ctreminiom/go-atlassian/compare/v1.4.0...v1.3.1)

> 2021-11-28


<a name="v1.4.0"></a>
## [v1.4.0](https://github.com/ctreminiom/go-atlassian/compare/v1.4.1...v1.4.0)

> 2021-11-28


<a name="v1.4.1"></a>
## [v1.4.1](https://github.com/ctreminiom/go-atlassian/compare/v2.0.0...v1.4.1)

> 2021-11-28


<a name="v2.0.0"></a>
## [v2.0.0](https://github.com/ctreminiom/go-atlassian/compare/v1.3.0...v2.0.0)

> 2021-11-28

### :art: Improve structure / format of the code.

* Moved the jira-worklog structs to the common model package
* Moved the jira-watcher struts to the common model package
* Moved the jira-vote struts to the common model package
* Moved the jira-type-screen-schemes struts to the common model package
* Moved the jira-types struts to the common model package
* Moved the jira-search struts to the common model package
* Moved the jira-resolutions struts to the common model package
* Moved the jira-priorities struts to the common model package
* Moved the jira-link-types struts to the common model package
* Moved the jira-labels struts to the common model package
* Moved the jira-field-context-options struts to the common model package
* Moved the jira-field-configuration struts to the common model package
* Moved the jira-fields struts to the common model package
* Moved the jira-comments struts to the common model package
* Moved the jira-attachments struts to the common model package
* Moved the jira-issue struts to the common model package
* Moved the models to /pkg
* Moved the jira-group struts to the common model package
* Moved the jira-filter struts to the common model package
* Moved the jira-dashboard struts to the common model package
* Moved the jira-dashboard struts to the common model package
* Moved the jira-audit struts to the common model package
* Moved the Application Role models
* Separated the Jira modules by version

### :bug: Bug Fixes

* Fixed the jira_work_log.go v2 payload format
* Fixed the issueLink.go v2 payload format
* Fixed the issue comment field parsing on the jira v2/3
* Fixed the jira.v2.Issue.Search.* methods
* Fixed the linter problems

### :memo: Add or update documentation.

* Updated README.md
* Added the docs.go-atlassian links on the Jira Cloud module
* Added the docs.go-atlassian links on the Admin Cloud module
* Added the docs.go-atlassian links on the Confluence module
* Added the docs.go-atlassian links on the Agile module
* Documented the WorkflowSchemeService under the jira/examples folder
* Documented the WorkflowService under the jira/examples folder
* Documented the Content Properties methods under the confluence/examples/contentProperties folder.
* Documented the Content Properties methods under the confluence/examples/contentProperties folder.
* Added more examples under the jira/examples/ directory
* Updated the README.md
* Updated the README.md

### :package: Dependencies

* Removed the go-querystring library
* Updated go.mod
* Moved the mock .json files to the common folder/
* Updated go.mod

### :recycle: Code Refactoring

* Standardize the Jira v3
* Standardize the Jira v2/v3
* Refactor the model names by application.
* Agile SprintService refactored
* EpicService refactored
* BoardService refactored

### :sparkles: Features

* Added the workflow scheme service.
* Added the jira workflow service
* Added the v2.Project.Version.Gets method
* Added the v3.Project.Version.Gets method
* Added the v2.Project.Gets() method
* Added the jira version 2 implementation
* Enabled to ability to delete an issue with a sub-tasks
* Added the Ancestor field on the ContentScheme
* Added the ability to get the issue create metadata
* Added the ability to get the issue metadata
* Added the ability to delete boards
* Added the WorkflowSchemeService under the Jira module
* Added the WorkflowService under the Jira module
* Added the Content Property Confluence module under the contentService.
* Added the ContentLabelService
* Added the SchemesByProject method under the issueTypeScreenScheme.go sub-module, closes [#58](https://github.com/ctreminiom/go-atlassian/issues/58)

### :white_check_mark: Ad/Update Test Cases

* Added the Workflow.Create test cases
* Added the Unit Test Cases under the WorkflowSchemeService with a 100% of code coverage.
* Added the Unit Test Cases under the WorkflowService with a 100% of code coverage.
* Added the Unit Test Cases under the contentProperties_test.go with a 100% of code coverage.

### :zap: Performance Improvements

* Added the agile.board.gets() method
* Renamed the v3.Project.Version.Gets method

### Construction_worker

* Added the .codecov.yml

### Pull Requests

* Merge pull request [#91](https://github.com/ctreminiom/go-atlassian/issues/91) from ctreminiom/feature/documentation
* Merge pull request [#90](https://github.com/ctreminiom/go-atlassian/issues/90) from ctreminiom/bug/issue-link-v2
* Merge pull request [#89](https://github.com/ctreminiom/go-atlassian/issues/89) from ctreminiom/bug/issue-link-v2
* Merge pull request [#88](https://github.com/ctreminiom/go-atlassian/issues/88) from ctreminiom/bug/jira-dashboard
* Merge pull request [#87](https://github.com/ctreminiom/go-atlassian/issues/87) from ctreminiom/dependabot/go_modules/github.com/tidwall/gjson-1.12.0
* Merge pull request [#86](https://github.com/ctreminiom/go-atlassian/issues/86) from ctreminiom/feature/struct-refactor
* Merge pull request [#85](https://github.com/ctreminiom/go-atlassian/issues/85) from ctreminiom/feature/custom-fields-get
* Merge pull request [#84](https://github.com/ctreminiom/go-atlassian/issues/84) from ctreminiom/feature/admin-refactor
* Merge pull request [#83](https://github.com/ctreminiom/go-atlassian/issues/83) from ctreminiom/feature/confluence-models
* Merge pull request [#82](https://github.com/ctreminiom/go-atlassian/issues/82) from ctreminiom/feature/sm-models-migration
* Merge pull request [#81](https://github.com/ctreminiom/go-atlassian/issues/81) from ctreminiom/feature/jira-fields-mapping
* Merge pull request [#80](https://github.com/ctreminiom/go-atlassian/issues/80) from ctreminiom/feature/agile-epics
* Merge pull request [#79](https://github.com/ctreminiom/go-atlassian/issues/79) from ctreminiom/feature/workflow-scheme
* Merge pull request [#78](https://github.com/ctreminiom/go-atlassian/issues/78) from ctreminiom/feature/jira-workflow
* Merge pull request [#77](https://github.com/ctreminiom/go-atlassian/issues/77) from ctreminiom/feature/agile-refactor
* Merge pull request [#76](https://github.com/ctreminiom/go-atlassian/issues/76) from ctreminiom/feature/search-issues-v2
* Merge pull request [#75](https://github.com/ctreminiom/go-atlassian/issues/75) from ctreminiom/feature/get-all-boards
* Merge pull request [#74](https://github.com/ctreminiom/go-atlassian/issues/74) from ctreminiom/feature/project-versions
* Merge pull request [#73](https://github.com/ctreminiom/go-atlassian/issues/73) from ctreminiom/feature/v2-get-all-projects
* Merge pull request [#71](https://github.com/ctreminiom/go-atlassian/issues/71) from ctreminiom/dependabot/go_modules/github.com/tidwall/gjson-1.11.0
* Merge pull request [#72](https://github.com/ctreminiom/go-atlassian/issues/72) from ctreminiom/feature/version-segmentation
* Merge pull request [#70](https://github.com/ctreminiom/go-atlassian/issues/70) from ctreminiom/feature/delete-with-sub-tasks
* Merge pull request [#67](https://github.com/ctreminiom/go-atlassian/issues/67) from ctreminiom/dependabot/go_modules/github.com/tidwall/gjson-1.10.2
* Merge pull request [#66](https://github.com/ctreminiom/go-atlassian/issues/66) from ctreminiom/feature/62
* Merge pull request [#65](https://github.com/ctreminiom/go-atlassian/issues/65) from ctreminiom/feature/64
* Merge pull request [#63](https://github.com/ctreminiom/go-atlassian/issues/63) from ctreminiom/feature/18
* Merge pull request [#61](https://github.com/ctreminiom/go-atlassian/issues/61) from ctreminiom/feature/board-delete
* Merge pull request [#60](https://github.com/ctreminiom/go-atlassian/issues/60) from ctreminiom/dev
* Merge pull request [#59](https://github.com/ctreminiom/go-atlassian/issues/59) from ctreminiom/dev
* Merge pull request [#57](https://github.com/ctreminiom/go-atlassian/issues/57) from ctreminiom/dependabot/go_modules/github.com/imdario/mergo-0.3.12


<a name="v1.3.0"></a>
## [v1.3.0](https://github.com/ctreminiom/go-atlassian/compare/v1.2.0...v1.3.0)

> 2021-07-17

### :bug: Bug Fixes

* Removed the Zero validation
* Enable default scheme
* Removed the omitempty struct tag on the CustomFieldContextOptionScheme to enable the ability to enable pass the false value.
* Fixed [#43](https://github.com/ctreminiom/go-atlassian/issues/43)

### :memo: Add or update documentation.

* Updated the README.md
* Updated the README.md
* Reduced the image size on the README.md
* Reduced the image size on the README.md
* Reduced the image size on the README.md
* Updated the README.md
*  Updated the README.md
* Updated the README.md
* CHANGELOG.md updated

### :package: Dependencies

* optimized the module dependencies decreasing the third party libraries needed.
* updated the go.mod

### :recycle: Code Refactoring

* Refactor the Jira Software module
* Refactor the Application Role and Audit modules
* Refactor Jira Agile Module

### :sparkles: Features

* Added project templates as constant
* Added the Issue worklog sub-module
* Added the spaceService module
* Added the ContentPermissionService service
* Added the ContentCommentService service
* Added the ContentChildrenDescendantService service
*  Added attachment Create method
* Added attachment Update/Create method
* Added the content.Attachment.Gets
* Added the content.History method
* Added the content.Get
* Added the content.Search
* content.Create method added
* Added Confluence service
* added the MySelf endpoint, close [#26](https://github.com/ctreminiom/go-atlassian/issues/26)
* Closes [#46](https://github.com/ctreminiom/go-atlassian/issues/46)

### :white_check_mark: Ad/Update Test Cases

* Unit Test Cases Added
* Add Test Cases

### Pull Requests

* Merge pull request [#56](https://github.com/ctreminiom/go-atlassian/issues/56) from ctreminiom/dev
* Merge pull request [#49](https://github.com/ctreminiom/go-atlassian/issues/49) from ctreminiom/feature/confluence-cloud
* Merge pull request [#36](https://github.com/ctreminiom/go-atlassian/issues/36) from ctreminiom/dependabot/go_modules/github.com/stretchr/testify-1.7.0
* Merge pull request [#48](https://github.com/ctreminiom/go-atlassian/issues/48) from ctreminiom/dev
* Merge pull request [#47](https://github.com/ctreminiom/go-atlassian/issues/47) from ctreminiom/dev
* Merge pull request [#45](https://github.com/ctreminiom/go-atlassian/issues/45) from ctreminiom/dev
* Merge pull request [#44](https://github.com/ctreminiom/go-atlassian/issues/44) from ctreminiom/dev


<a name="v1.2.0"></a>
## [v1.2.0](https://github.com/ctreminiom/go-atlassian/compare/v1.1.5...v1.2.0)

> 2021-05-11

### :art: Improve structure / format of the code.

* Added more Tags on the IssueScheme struct.

### :bug: Bug Fixes

* Fixed the error: SA4006: this value of `err` is never used (staticcheck)
* Fixed the Lint error: structtag: struct field Description repeats json tag "status" also at issue.go:58 (govet)
* Closes [#19](https://github.com/ctreminiom/go-atlassian/issues/19)
* Closes [#9](https://github.com/ctreminiom/go-atlassian/issues/9)

### :memo: Add or update documentation.

* README.md updated
* Updated the README.md adding more badgets.
* Updated the CHANGELOG.md

### :sparkles: Features

* Added the Check method under the permission.go service
* Added the ProjectContext method under the issueFieldContext.go service.
* Added the IssueTypesContext method under the issueFieldContext.go service.
* closes [#40](https://github.com/ctreminiom/go-atlassian/issues/40)
* closes [#38](https://github.com/ctreminiom/go-atlassian/issues/38), closes [#39](https://github.com/ctreminiom/go-atlassian/issues/39)
* Added the SprintService under the Agile Module
* Added the first Jira Cloud Agile Integration
* Enable the ability to Assign Field Configuration Scheme to a Jira Project, closes [#20](https://github.com/ctreminiom/go-atlassian/issues/20)

### :white_check_mark: Ad/Update Test Cases

* Increased the code coverage on the Jira on the Jira module
* Increased the code coverage on the AgileService on the Atlassian Agile module

### Pull Requests

* Merge pull request [#42](https://github.com/ctreminiom/go-atlassian/issues/42) from ctreminiom/dev
* Merge pull request [#37](https://github.com/ctreminiom/go-atlassian/issues/37) from ctreminiom/dependabot/add-v2-config-file
* Merge pull request [#41](https://github.com/ctreminiom/go-atlassian/issues/41) from ctreminiom/feature/jira-agile
* Merge pull request [#31](https://github.com/ctreminiom/go-atlassian/issues/31) from ctreminiom/dev


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
## [v1.1.1](https://github.com/ctreminiom/go-atlassian/compare/v1.1.0...v1.1.1)

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
