# go-atlassian 

[![GoDoc][1]][2]
[![GoCard][3]][4]
[![Codecov][5]][6]
[![License][7]][8]
[![][9]][10]
[![a][11]][12]


[1]: https://pkg.go.dev/badge/github.com/ctreminiom/go-atlassian?utm_source=godoc
[2]: https://pkg.go.dev/github.com/ctreminiom/go-atlassian
[3]: https://goreportcard.com/badge/ctreminiom/go-atlassian
[4]: https://goreportcard.com/report/github.com/ctreminiom/go-atlassian
[5]: https://codecov.io/gh/ctreminiom/go-atlassian/branch/main/graph/badge.svg?token=G0KPNMTIRV
[6]: https://codecov.io/gh/ctreminiom/go-atlassian
[7]: https://img.shields.io/badge/license-MIT-blue.svg
[8]: https://github.com/ctreminiom/go-atlassian/blob/master/LICENSE
[9]:  https://img.shields.io/github/workflow/status/ctreminiom/go-atlassian/Testing?label=%F0%9F%A7%AA%20tests&style=flat&color=75C46B
[10]: https://github.com/ctreminiom/go-atlassian/actions?query=workflow%3ATesting
[11]: https://img.shields.io/badge/%F0%9F%92%A1%20go-documentation-00ACD7.svg?style=flat
[12]: https://docs.go-atlassian.io/

> `go-atlassian` is a [Atlassian Cloud](https://www.atlassian.com/cloud) client library written in Golang. It interacts with the following services:

### 📘 [Documentation](https://docs.go-atlassian.io/)

|Application|  Status|
|--|--|
|Jira Cloud | Available ✅|
|Jira Agile Cloud | In development 👷|
|Jira Service Management Cloud | In development 👷|
|Confluence Cloud | In development 👷|
|Atlassian Admin Cloud | In development 👷|

## Features
- Create issue issues with custom fields
- Manage the screens, screens schemes, issue type screen schemes and all endpoints that interacts with the customfields
- The 90% of the endpoints documented [here](https://developer.atlassian.com/cloud/jira/platform/rest/v3/intro/) were mapped and documented with examples.

## Installation 📖
Make sure you have a working Go 1.14+ workspace (_[instructions](https://golang.org/doc/install)_), then:
```sh
$ go get github.com/ctreminiom/go-atlassian/jira
```

## Usage ✒️
All interaction starts with a `jira.Client` struct. Create one with your Atlassian site host URL and a custom HTTP client if it's necessary.
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

    // You can set custom *http.Client here
	jiraCloud, err := jira.New(nil, host)
	if err != nil {
		return
	}

	jiraCloud.Auth.SetBasicAuth(mail, token)
	jiraCloud.Auth.SetUserAgent("curl/7.54.0")

	applicationRoles, response, err := jiraCloud.Role.Gets(context.Background())
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for _, applicationRole := range *applicationRoles {
		log.Printf("Application Role Name: %v", applicationRole.Name)
		log.Printf("Application Role Key: %v", applicationRole.Key)
		log.Printf("Application Role User Count: %v", applicationRole.UserCount)
	}

}
```

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
