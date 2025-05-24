package task

import (
	"github.com/tylerBrittain42/YesterdayI/internal/config"
)

func AddTask(fName string, c *config.Config) error {

	tSlice, err := load(fName)
	if err != nil {
		return err
	}

	tSlice.add(c.Content, c.JiraTicket)
	tSlice.save(fName)

	return nil

}
