package task

import (
	"time"
)

type Task struct {
	Content     string
	JiraTicket  string
	dateCreated time.Time
}

func NewTask(content string, jira string) Task {
	return Task{Content: content, JiraTicket: jira, dateCreated: time.Now()}

}
