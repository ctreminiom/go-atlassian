<p align="center"><img width="600" src="./jira/mocks/go-atlassian-logo.svg" alt="Go-Atlassian logo"></p>

<p align="center">
    <a href="https://pkg.go.dev/github.com/ctreminiom/go-atlassian"><img src="https://pkg.go.dev/badge/github.com/ctreminiom/go-atlassian?utm_source=godoc"></a>
    <a href="https://goreportcard.com/report/github.com/ctreminiom/go-atlassian"><img src="https://goreportcard.com/badge/ctreminiom/go-atlassian"></a>
    <a href="https://codecov.io/gh/ctreminiom/go-atlassian"><img src="https://codecov.io/gh/ctreminiom/go-atlassian/branch/main/graph/badge.svg?token=G0KPNMTIRV"></a>
    <a href="https://www.codacy.com/gh/ctreminiom/go-atlassian/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ctreminiom/go-atlassian&amp;utm_campaign=Badge_Grade"><img src="https://app.codacy.com/project/badge/Grade/fe5c1b3c9fd64f84989ae51c42803456"/></a>
    <a href="https://github.com/ctreminiom/go-atlassian/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
    <a href="https://github.com/ctreminiom/go-atlassian/actions?query=workflow%3ATesting"><img src="https://img.shields.io/github/workflow/status/ctreminiom/go-atlassian/Testing?label=%F0%9F%A7%AA%20tests&style=flat&color=75C46B"></a>
    <a href="https://docs.go-atlassian.io/"><img src="https://img.shields.io/badge/%F0%9F%92%A1%20go-documentation-00ACD7.svg?style=flat"></a>
</p>

## Introduction 📖

go-atlassian is a library written in Go programming language that enables the interaction with the Atlassian Cloud API's. It consists of the following services that Atlassian provide us:
* Jira Software Cloud
* Jira Service Management Cloud
* Confluence Cloud
* Atlassian Access 
* Opsgenie
* Trello
* Bitbucket Cloud

The Complete documentation is available at [docs.go-atlassian.io](https://docs.go-atlassian.io/).

## Development
Right now, the library supports the Jira Software Cloud and Jira Service Management Cloud services. This project's still in progress, and the remaining services will be mapped and documented.

## Jira Software Cloud 
Plan, track, and release world-class software with the #1 software development tool used by agile teams.

### Features
* Create/Edit/Delete/View issues 
* Support the Jira issue custom-fields interactions
* Manage the project screen, screen screens, issue type scheme screens, etc.
* Change the issue status, retrieve the issue changelogs, search issues based on the JQL query and more!!.

#### Installation ✒
```sh
$ go get -u -v github.com/ctreminiom/go-atlassian/jira
```

#### Use Cases

<details><summary>Get Issue</summary>

```go
package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	issue, response, err := atlassian.Issue.Get(context.Background(), "KP-12", []string{"status"}, []string{"transitions"})
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
			log.Println(response.StatusCode)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println(issue.Key)
}
```
</details>

<details><summary>Get Project Categories</summary>

```go
package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	/*
		----------- Set an environment variable in git bash -----------
		export HOST="https://ctreminiom.atlassian.net/"
		export MAIL="MAIL_ADDRESS"
		export TOKEN="TOKEN_API"

		Docs: https://stackoverflow.com/questions/34169721/set-an-environment-variable-in-git-bash
	*/

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	categories, response, err := atlassian.Project.Category.Gets(context.Background())
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, category := range *categories {

		log.Println("----------------")
		log.Println(category.Self)
		log.Println(category.ID)
		log.Println(category.Name)
		log.Println(category.Description)
		log.Println("----------------")
	}
}

```
</details>

<details><summary>Create Project Version</summary>

```go
package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	/*
		----------- Set an environment variable in git bash -----------
		export HOST="https://ctreminiom.atlassian.net/"
		export MAIL="MAIL_ADDRESS"
		export TOKEN="TOKEN_API"

		Docs: https://stackoverflow.com/questions/34169721/set-an-environment-variable-in-git-bash
	*/

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		log.Fatal(err)
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	payload := &jira.ProjectVersionPayloadScheme{
		Archived:    false,
		ReleaseDate: "2021-03-06",
		Name:        "Version Sandbox",
		Description: "Version Sandbox description",
		ProjectID:   10000,
		Released:    false,
		StartDate:   "2021-03-02",
	}

	newVersion, response, err := atlassian.Project.Version.Create(context.Background(), payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Printf("The new version has been created with the ID %v", newVersion.ID)
}

```
</details>

## Jira Service Management Cloud
Collaborate at high-velocity, respond to business changes and deliver great customer and employee service experiences fast.

### Features
* Create/Edit/Delete/View Service Desk Organizations
* Create Request Types, Customers.
* Get the Service Desk Articles, Queues, Request Comments, Participants, etc  

#### Installation ✒
```sh
$ go get -u -v github.com/ctreminiom/go-atlassian/jira
```

#### Use Cases

<details><summary>Create Organization</summary>

```go
package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var organizationName = "Organization Name"

	newOrganization, response, err := atlassian.ServiceManagement.Organization.Create(context.Background(), organizationName)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Printf("The organization has been created: %v", newOrganization.ID)
}
```
</details>

<details><summary>Get Request Approvals</summary>

```go
package main

import (
	"context"
	"encoding/json"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var issueKey = "DESK-12"
	approvals, response, err := atlassian.ServiceManagement.Request.Approval.Gets(context.Background(), issueKey, 0, 50)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, customRequest := range approvals.Values {

		dataAsJson, err := json.MarshalIndent(customRequest, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(dataAsJson))
	}

}
```
</details>

<details><summary>Get Service Desk Queues</summary>

```go
package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var (
		serviceDeskID      = 1
		includeCount  bool = true
		start, limit  int  = 0, 50
	)

	queues, response, err := atlassian.ServiceManagement.ServiceDesk.Queue.Gets(context.Background(), serviceDeskID, includeCount, start, limit)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	for pos, queue := range queues.Values {

		log.Println("------------------------------------")
		log.Printf("Queue ID #%v: %v", pos+1, queue.ID)
		log.Printf("Queue Name #%v: %v", pos+1, queue.Name)
		log.Printf("Queue JQL #%v: %v", pos+1, queue.Jql)
		log.Printf("Queue Issue Count #%v: %v", pos+1, queue.IssueCount)
		log.Printf("Queue Fields #%v: %v", pos+1, queue.Fields)
		log.Println("------------------------------------")
	}

}
```
</details>

## Run tests
```sh
go test -v ./...
```

## Author 

👤 **Carlos Treminio**
* Website: https://ctreminiom.gitbook.io/docs/
* Github: [@ctreminiom](https://github.com/ctreminiom)
* LinkedIn: [@ctreminio](https://linkedin.com/in/ctreminio)

## 🤝 Contributing
Contributions, issues and feature requests are welcome!
Feel free to check [issues page](https://github.com/ctreminiom/go-atlassian/issues).
## Show your support
Give a ⭐️ if this project helped you!
[![Buy Me A Coffee](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/ctreminiom)

## 📝 License
Copyright © 2021 [Carlos Treminio](https://github.com/ctreminiom).
This project is [MIT](https://opensource.org/licenses/MIT) licensed.
