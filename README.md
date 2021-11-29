<p align="center"><img width="600" src="https://user-images.githubusercontent.com/16035390/131232958-022b0382-e6bc-42db-97b6-82fbd190e19a.png"></p>

<p align="center">
    <a href="https://github.com/ctreminiom/go-atlassian/releases/latest"><img src="https://img.shields.io/github/v/release/ctreminiom/go-atlassian"></a>
    <a href="https://pkg.go.dev/github.com/ctreminiom/go-atlassian"><img src="https://pkg.go.dev/badge/github.com/ctreminiom/go-atlassian?utm_source=godoc"></a>
    <a href="https://goreportcard.com/report/github.com/ctreminiom/go-atlassian"><img src="https://goreportcard.com/badge/ctreminiom/go-atlassian"></a>
    <a href="https://app.fossa.com/projects/git%2Bgithub.com%2Fctreminiom%2Fgo-atlassian?ref=badge_shield" alt="FOSSA Status"><img src="https://app.fossa.com/api/projects/git%2Bgithub.com%2Fctreminiom%2Fgo-atlassian.svg?type=shield"/></a>
    <a href="https://codecov.io/gh/ctreminiom/go-atlassian"><img src="https://codecov.io/gh/ctreminiom/go-atlassian/branch/main/graph/badge.svg?token=G0KPNMTIRV"></a>
    <a href="https://www.codacy.com/gh/ctreminiom/go-atlassian/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ctreminiom/go-atlassian&amp;utm_campaign=Badge_Grade"><img src="https://app.codacy.com/project/badge/Grade/fe5c1b3c9fd64f84989ae51c42803456"/></a>
    <a href="https://github.com/ctreminiom/go-atlassian/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
    <a href="https://github.com/ctreminiom/go-atlassian/actions?query=workflow%3ATesting"><img src="https://img.shields.io/github/workflow/status/ctreminiom/go-atlassian/Testing?label=%F0%9F%A7%AA%20tests&style=flat&color=75C46B"></a>
    <a href="https://docs.go-atlassian.io/"><img src="https://img.shields.io/badge/%F0%9F%92%A1%20go-documentation-00ACD7.svg?style=flat"></a>
    <a href="https://bestpractices.coreinfrastructure.org/projects/4861"><img src="https://bestpractices.coreinfrastructure.org/projects/4861/badge"></a> 
</p>

go-atlassian is a Go module that enables the interaction with the Atlassian Cloud Services. It provides the Go implementation for operating the Atlassian Cloud platform.

## ✨ Features
- Supports Jira Software V2/V3 endpoints.
- Support the Jira Agile endpoints
- Interacts with the Jira Service Management entities.
- Manages the Atlassian Admin Cloud organizations.
- Manages the Atlassian Admin SCIM workflow.
- Checks Confluence Cloud content permissions.
- CRUD Confluence Cloud content (page, blogpost, comment, question).
- Add attachment into Confluence Cloud contents.
- Search contents and spaces.
- Support the Atlassian Document Format (ADF).  
- Every method has their corresponding example documented.

##  🔰 Installation
Make sure you have Go installed (download). Version `1.13` or higher is required.
```sh
$ go get -u -v github.com/ctreminiom/go-atlassian
```

## 📓 Documentation
Documentation is hosted live at https://docs.go-atlassian.io/

## 📝 Using the library
````go
package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira/v2"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := v2.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)

	issue, response, err := atlassian.Issue.Get(context.Background(), "KP-2", nil, []string{"transitions"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTTP Endpoint Used", response.Endpoint)

	log.Println(issue.Key)
	log.Println(issue.Fields.Reporter.AccountID)

	for _, transition := range issue.Transitions {
		log.Println(transition.Name, transition.ID, transition.To.ID, transition.HasScreen)
	}

	// Check if the issue contains sub-tasks
	if issue.Fields.Subtasks != nil {
		for _, subTask := range issue.Fields.Subtasks {
			log.Println("Sub-Task: ", subTask.Key, subTask.Fields.Summary)
		}
	}
}
````

## ⭐️ Project assistance
If you want to say **thank you** or/and support active development of `go-atlassian`:

- Add a [GitHub Star](https://github.com/ctreminiom/go-atlassian) to the project.
- Write interesting articles about project on [Dev.to](https://dev.to/), [Medium](https://medium.com/) or personal blog.
- Contributions, issues and feature requests are welcome!
-  Feel free to check [issues page](https://github.com/ctreminiom/go-atlassian/issues).

## 💡 Inspiration
The project was created with the purpose to provide a unique point to provide an interface for interacting with Atlassian products. This module is highly inspired by the Go library https://github.com/andygrunwald/go-jira
but focused on Cloud solutions.

## 🧪 Run Test Cases
```sh
go test -v ./...
```

## 📝 License
Copyright © 2021 [Carlos Treminio](https://github.com/ctreminiom).
This project is [MIT](https://opensource.org/licenses/MIT) licensed.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fctreminiom%2Fgo-atlassian.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fctreminiom%2Fgo-atlassian?ref=badge_large)

## 🌐 Credits
In addition to all the contributors we would like to thank to these companies:
-   [Atlassian](https://www.atlassian.com/)  for providing us Atlassian Admin/Jira/Confluence Standard licenses.
-   [JetBrains](https://www.jetbrains.com/)  for providing us with free licenses of  [GoLand](https://www.jetbrains.com/pycharm/)
-   [GitBook](https://www.gitbook.com/)  for providing us non-profit / open-source plan so hence I would like to express my thanks here.

<img src="./static/jetbrains-logo.svg">
<img src="./static/gitbook-logo.svg">