package task

import (
	"errors"
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

func newTask(content string, jira string) (Task, error) {
	if content == "" {
		return Task{}, errors.New("content has empty value")
	}
	return Task{Content: content, JiraTicket: jira, dateCreated: time.Now()}, nil

}

type TaskSlice []Task

func (s *TaskSlice) Add(content string, jira string) error {
	task, err := newTask(content, jira)
	if err != nil {
		return err
	}
	*s = append(*s, task)

	return nil
}

func GetFromFile(fName string) []Task {

	return []Task{}

}
