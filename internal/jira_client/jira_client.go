package jira_client

import (
	"os"

	"github.com/andygrunwald/go-jira"
)

func GetCurrentIssueStatus() (*string, *jira.Response, error) {
	issue_key := "WTA-1896"
	server := "https://jira.petrovich.tech"
	tp := jira.BearerAuthTransport{
		Token:     os.Getenv("JIRA_TOKEN"),
		Transport: nil,
	}

	JiraClient, err := jira.NewClient(tp.Client(), server)
	if err != nil {
		return nil, nil, err
	}

	issue, response, err := JiraClient.Issue.Get(issue_key, nil)
	if err != nil {
		return nil, response, err
	}

	current_status := issue.Fields.Status.Name
	return &current_status, response, nil
}
