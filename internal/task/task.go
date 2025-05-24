package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/tylerBrittain42/YesterdayI/internal/config"
)

type Task struct {
	Content     string    `json:"content"`
	JiraTicket  string    `json:"jira_ticket"`
	DateCreated time.Time `json:"date_created"`
}

func (t Task) String() string {
	return fmt.Sprintf("\"%s\" %s %s", t.Content, t.JiraTicket, t.DateCreated.Format("01/02"))
}

func (t Task) Pretty() string {
	if t.JiraTicket != "" {
		return fmt.Sprintf("task: %s\nJIRA: %s", t.Content, t.JiraTicket)
	} else {
		return fmt.Sprintf("task: %s", t.Content)
	}
}

func newTask(content string, jira string) (Task, error) {
	if content == "" {
		return Task{}, errors.New("content has empty value")
	}
	return Task{Content: content, JiraTicket: jira, DateCreated: time.Now()}, nil

}

type TaskSlice []Task

func (s *TaskSlice) add(content string, jira string) error {
	task, err := newTask(content, jira)
	if err != nil {
		return err
	}
	*s = append(*s, task)

	return nil
}

func (s *TaskSlice) save(fName string) error {

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

func load(fName string) (TaskSlice, error) {
	_, err := os.Stat(fName)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("error checking file existence: %w", err)
		}
		return TaskSlice{}, nil
	}

	f, err := os.ReadFile(fName)
	if err != nil {
		return nil, fmt.Errorf("unable to load: %w", err)
	}

	tSlice := TaskSlice{}

	// guard clause bc json.Unmarshal will throw an error if it is given an empty slice
	if len(f) == 0 {
		return tSlice, nil
	}

	err = json.Unmarshal(f, &tSlice)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal: %w", err)
	}
	return tSlice, nil

}

func AddTask(fName string, c *config.Config) error {

	tSlice, err := load(fName)
	if err != nil {
		return err
	}

	tSlice.add(c.Content, c.JiraTicket)
	tSlice.save(fName)

	return nil

}

func View(fName string, c *config.Config) error {
	tSlice, err := load(fName)
	if err != nil {
		return err
	}
	if c.StartTime == "" && c.EndTime == "" && c.SpecificTime == "" {
		tSlice.viewAll()
	} else if c.StartTime != "" && c.EndTime == "" && c.SpecificTime == "" {
		tSlice.viewUntilNow(c)
	} else if c.StartTime != "" && c.EndTime != "" && c.SpecificTime == "" {
		tSlice.viewRange()
	} else if c.StartTime == "" && c.EndTime == "" && c.SpecificTime != "" {
		tSlice.viewSpecificTime()
	} else {
		return errors.New("unsupported combination of flags")

	}

	return nil
}

func (s *TaskSlice) viewAll() {
	lastDate := (*s)[0].DateCreated.Format("01/02")
	fmt.Printf("---%s---\n", lastDate)
	for _, v := range *s {
		thisDate := v.DateCreated.Format("01/02")
		if thisDate != lastDate {
			lastDate = thisDate
			fmt.Printf("\n---%s---\n", lastDate)
		}
		fmt.Println(v.Pretty())

	}
}
func (s *TaskSlice) viewUntilNow(c *config.Config) error {
	parsedDate, err := time.Parse("01/02", c.StartTime)
	if err != nil {
		return fmt.Errorf("unable to parse start date: %w", err)
	}
	minDate := time.Date(time.Now().Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, time.Local)

	lastDate := (*s)[0].DateCreated.Format("01/02")
	first := true
	for _, v := range *s {
		thisDate := v.DateCreated.Format("01/02")
		if v.DateCreated.After(minDate) {

			// weird funkyness to ensure that the first line is not empty
			if first {
				first = false
				fmt.Printf("---%s---\n", thisDate)
				lastDate = thisDate

			} else if thisDate != lastDate {
				lastDate = thisDate
				fmt.Printf("\n---%s---\n", thisDate)
			}
			fmt.Println(v.Pretty())
		}
	}
	return nil
}

func (s *TaskSlice) viewSpecificTime() {
	return
}

func (s *TaskSlice) viewRange() {
	return
}
