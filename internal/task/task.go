package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Content     string    `json:"content"`
	JiraTicket  string    `json:"jira_ticket"`
	DateCreated time.Time `json:"date_created"`
}

func (t Task) String() string {
	return fmt.Sprintf("\"%s\" %s %s", t.Content, t.JiraTicket, t.DateCreated.Format("01/02"))
}

func newTask(content string, jira string) (Task, error) {
	if content == "" {
		return Task{}, errors.New("content has empty value")
	}
	return Task{Content: content, JiraTicket: jira, DateCreated: time.Now()}, nil

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

func (s *TaskSlice) Save(fName string) error {

	b, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("unable to save: %w", err)
	}

	err = os.WriteFile(fName, b, 0644)
	if err != nil {
		return fmt.Errorf("unable to save: %w", err)
	}

	return nil

}

func Load(fName string) (TaskSlice, error) {
	f, err := os.ReadFile(fName)
	if err != nil {
		return nil, fmt.Errorf("unable to load: %w", err)
	}

	tSlice := TaskSlice{}
	err = json.Unmarshal(f, &tSlice)
	if err != nil {
		return nil, fmt.Errorf("unable to load: %w", err)
	}
	return tSlice, nil

}
