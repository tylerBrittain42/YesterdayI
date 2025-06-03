package task

import (
	"github.com/tylerBrittain42/YesterdayI/internal/config"
)

func AddTask(fName string, c *config.Config) error {

	tSlice, err := load(fName)
	if err != nil {
		return err
	}

	err = tSlice.add(c.Content, c.JiraTicket)
	if err != nil {
		return err
	}

	err = tSlice.save(fName)
	if err != nil {
		return err
	}

	return nil

}
