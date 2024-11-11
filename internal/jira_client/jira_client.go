package jira_client

import (
	"os"

	"github.com/andygrunwald/go-jira"
)

func Connect() string {
	issue_key := "WTA-1896"
	server := "jira.petrovich.tech"
	tp := jira.BearerAuthTransport{
		Token: os.Getenv("JIRA_TOKEN"),
	}

	JiraClient, _ := jira.NewClient(tp.Client(), server)
	// if err != nil {
	// 	panic(err)
	// }

	issue, _, _ := JiraClient.Issue.Get(issue_key, nil)

	current_status := issue.Fields.Status.Name
	return current_status
}
