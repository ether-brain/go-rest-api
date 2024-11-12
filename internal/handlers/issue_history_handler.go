package handlers

import (
	"github.com/andygrunwald/go-jira"
	jiraClient "github.com/ether-brain/go-rest-api/internal/jira_client"
)

func Handle() (*string, *jira.Response, error) {
	issueStatus, response, err := jiraClient.GetCurrentIssueStatus()
	return issueStatus, response, err
}
