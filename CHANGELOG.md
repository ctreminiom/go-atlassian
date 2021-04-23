
<a name="v1.1.4"></a>
## [v1.1.4](https://github.com/ctreminiom/go-atlassian/compare/v1.1.3...v1.1.4) (2021-04-07)

### Bug

* Fixed Issue Test Cases coverage issues
* Types the Lint problems

### Memo

* updated the time-in-status example

### Sparkles

* Implemented the Issue Update using operations

### White_check_mark

* Added the Test Cases for the AuthService
* Increased the code coverage on the issueService
* Added test cases on the Operations Update

### Pull Requests

* Merge pull request [#17](https://github.com/ctreminiom/go-atlassian/issues/17) from ctreminiom/dev


<a name="v1.1.3"></a>
## [v1.1.3](https://github.com/ctreminiom/go-atlassian/compare/v1.1.2...v1.1.3) (2021-04-07)

### Beers

* added the Atlassian Cloud Admin Organization Module with examples

### Children_crossing

* Updated the dev branch with the most recent changes.

### Memo

* added the issue labels examples
* added the issue fields examples
* added the time-in-status example and updated the dependencies,
* added advanced examples like "add field to project" or extract issue changelogs

### Pencil2

* Fixed the policy Update example

### Sparkles

* Added the SCIM User Service
* Added the SCIM Scheme Service
* Added the user SCIM service with the User CRUD endpoints mapped.
* added the Atlassian User Admin service with examples and test cases
* mapped the /rest/api/3/issuetypescreenscheme/project and /rest/api/3/issuetypescreenscheme/mapping endpoints

### White_check_mark

* added more test cases on the admin and auth services
* added the test cases on the OrganizationPolicyService
* added more test cases on the OrganizationService

### Pull Requests

* Merge pull request [#16](https://github.com/ctreminiom/go-atlassian/issues/16) from ctreminiom/feature/cloud-admin
* Merge pull request [#15](https://github.com/ctreminiom/go-atlassian/issues/15) from ctreminiom/dev
* Merge pull request [#13](https://github.com/ctreminiom/go-atlassian/issues/13) from ctreminiom/dev


<a name="v1.1.2"></a>
## [v1.1.2](https://github.com/ctreminiom/go-atlassian/compare/v1.1.1...v1.1.2) (2021-03-24)

### Memo

* added examples

### Pull Requests

* Merge pull request [#11](https://github.com/ctreminiom/go-atlassian/issues/11) from ctreminiom/feature/jira-software-example


<a name="v1.1.1"></a>
## [v1.1.1](https://github.com/ctreminiom/go-atlassian/compare/v1.0.2...v1.1.1) (2021-03-24)

### Alembic

* Added the Issue.Comment.Add method

### Art

* Updated README.md
* Updated README.md
* Format the badges
* Added the new Library Logo

### Lipstick

* Added the go-atlassian-logo in the .svg format

### Page_facing_up

* changed the logo image

### White_check_mark

* Increased the coverage adding more Test on the issueComment.go

### Pull Requests

* Merge pull request [#8](https://github.com/ctreminiom/go-atlassian/issues/8) from ctreminiom/feature-adf-comments
* Merge pull request [#6](https://github.com/ctreminiom/go-atlassian/issues/6) from ctreminiom/feat/updated-readme.md
* Merge pull request [#5](https://github.com/ctreminiom/go-atlassian/issues/5) from ctreminiom/feature/jira-service-management


<a name="v1.0.2"></a>
## [v1.0.2](https://github.com/ctreminiom/go-atlassian/compare/v1.1.0...v1.0.2) (2021-03-23)


<a name="v1.1.0"></a>
## [v1.1.0](https://github.com/ctreminiom/go-atlassian/compare/v1.0.1...v1.1.0) (2021-03-23)

### Ambulance

* fixed the GoLint non-used warning on the Service Management code samples

### Art

* Improved the code samples folder

### Construction

* Created the first Jira Service Management services and linked it into the JiraService struct.

### Rocket

*  Added the ability to retrieves the customer transitions and move issues on the Service Management Module
* Fixed the Lint error S1039: unnecessary use of fmt.Sprintf (gosimple)
* added the feature/* branches on the CI/CD Actions

### Sparkles

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

### White_check_mark

* Added the Unit Test cases on the KnowledgebaseService and OrganizationService
* Added the Unit Test cases on the customerService

### Pull Requests

* Merge pull request [#5](https://github.com/ctreminiom/go-atlassian/issues/5) from ctreminiom/feature/jira-service-management


<a name="v1.0.1"></a>
## [v1.0.1](https://github.com/ctreminiom/go-atlassian/compare/v1.0.0...v1.0.1) (2021-03-04)

### Bug

* Closes [#2](https://github.com/ctreminiom/go-atlassian/issues/2)

### Pull Requests

* Merge pull request [#3](https://github.com/ctreminiom/go-atlassian/issues/3) from ctreminiom/issue-2-unmarshal_array_error_on_IssueSearchService


<a name="v1.0.0"></a>
## v1.0.0 (2021-03-03)

### Ambulance

* Fixed empty key param on the applicationRole.go and updated the code examples

### Art

* added the Jira Date Format constant

### Bug

* Fixed the missing request headers blocking the HTTP callback returning a 415 and updated the comments documentation.

### Construction

* updated the /vendor folder
* updated the /vendor folder

### Fire

* Implemented the dynamic customfield parsing on the issue create/creates/update method in the issue.go module

### Lipstick

* Increased the code coverage on the issuePriorities_test.go

### Memo

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

### Passport_control

* Added the SetUserAgent method.

### Pushpin

* formatted the dashboards examples

### Recycle

* Added the Unit Test cases with a 100% of coverage on the issueComment.go file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the audit.go file and the .json mock files needed to run the tests.

### Rocket

* changed the nested struct under the IssueScheme struct to pointers with the purpose to use the omitempty tag on custom issue payload creations.

### Seedling

* Added more Test cases on the issueAttachment_test.go
* Added more Test cases on the GroupService
* Increased the code coverage on the applicationRole_test.go
* added the .json mock field needed to the issue.go test cases.

### Sparkles

* updated Readme.md
* Added more test case on the issueTypeScreenScheme.go
* handled the non resolutionID or priorityID params.
* added more method in the dashboardService

### Speech_balloon

* ignored the .idea/ folder created by GoLand

### Truck

* migrated the code samples into https://docs.go-atlassian.io/

### White_check_mark

* Added the Unit Test cases with a 100% of coverage on the userSearch.go file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the projectVersion.go file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the screenSchemes.go file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the screenTab.go  file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the screen.go  file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the user.go file and the .json mock files needed to run the tests.
* Added the Unit Test cases with a 100% of coverage on the projectTypes.go file and the .json mock files needed to run the tests.

### Wrench

* mapped more endpoint on the issueFieldContext service.

### Pull Requests

* Merge pull request [#1](https://github.com/ctreminiom/go-atlassian/issues/1) from ctreminiom/dev

