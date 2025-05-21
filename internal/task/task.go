package task

import (
	"fmt"
	"time"
)

type Task struct {
	Content     string
	JiraTicket  string
	dateCreated time.Time
}

func (t Task) String() string {
	return fmt.Sprintf("\"%s\" %s %s", t.Content, t.JiraTicket, t.dateCreated.Format("01/02"))
}

func newTask(content string, jira string) Task {
	return Task{Content: content, JiraTicket: jira, dateCreated: time.Now()}

}

type TaskSlice []Task

func (s *TaskSlice) Add(content string, jira string) error {
	_ = newTask(content, jira)
	return nil
}

func GetFromFile(fName string) []Task {

	return []Task{}

}
