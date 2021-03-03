---
name: Bug report
about: Create a report to help us improve
title: ''
labels: bug
assignees: ctreminiom

---

**go-atlassian version**


**go-atlassian component**
- [ ] Jira Software Cloud
- [ ] Jira Agile Cloud
- [ ] Jira Service Management Cloud 
- [ ] Confluence Cloud
- [ ] Atlassian Admin Cloud

**Describe the bug** :bug:
A clear and concise description of what the bug is.

**To Reproduce** :construction: 
Steps to reproduce the behavior:

**Expected behavior** :white_check_mark: 
A clear and concise description of what you expected to happen.

**Screenshots** :page_facing_up: 
If applicable, add screenshots to help explain your problem.

**Additional context**
Add any other context about the problem here.

**Code snippet**
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

	jiraCloud, err := jira.New(nil, host)
	if err != nil {
		return
	}

	jiraCloud.Auth.SetBasicAuth(mail, token)
	jiraCloud.Auth.SetUserAgent("curl/7.54.0")

	// Steps to reproduce

}
```
