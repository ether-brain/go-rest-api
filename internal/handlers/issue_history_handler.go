package handlers

import (
	"fmt"

	jira_client "github.com/ether-brain/go-rest-api/internal/jira_client"
)

func Handle() string {
	issue_data := jira_client.Connect()
	fmt.Println(issue_data)
	return issue_data
}
