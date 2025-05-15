package task

import (
	"time"
)

type Task struct {
	Content     string
	JiraTicket  string
	dateCreated time.Time
}

type TaskSlice []Task

func newTask(content string, jira string) Task {
	return Task{Content: content, JiraTicket: jira, dateCreated: time.Now()}

}

func (s TaskSlice) String() string {
	return "not implemented"
}
func (s *TaskSlice) Add(content string, jira string) error {
	_ = newTask(content, jira)
	return nil

}

func GetFromFile(fName string) []Task {

	return []Task{}

}
