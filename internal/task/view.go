package task

import (
	"errors"
	"fmt"
	"time"

	"github.com/tylerBrittain42/YesterdayI/internal/config"
)

func View(fName string, c *config.Config) error {
	tSlice, err := load(fName)
	if err != nil {
		return err
	}
	if len(tSlice) == 0 {
		return fmt.Errorf("no entries found")
	}

	if c.StartTime == "" && c.EndTime == "" && c.SpecificTime == "" {
		tSlice.viewAll()
	} else if c.StartTime != "" && c.EndTime == "" && c.SpecificTime == "" {
		tSlice.viewUntilNow(c)
	} else if c.StartTime != "" && c.EndTime != "" && c.SpecificTime == "" {
		tSlice.viewRange(c)
	} else if c.StartTime == "" && c.EndTime == "" && c.SpecificTime != "" {
		tSlice.viewSpecificTime(c)
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

func (s *TaskSlice) viewRange(c *config.Config) error {
	parsedStartDate, err := time.Parse("01/02", c.StartTime)
	if err != nil {
		return fmt.Errorf("unable to parse start date: %w", err)
	}
	minDate := time.Date(time.Now().Year(), parsedStartDate.Month(), parsedStartDate.Day(), 0, 0, 0, 0, time.Local)

	parsedEndDate, err := time.Parse("01/02", c.EndTime)
	if err != nil {
		return fmt.Errorf("unable to parse end date: %w", err)
	}
	maxDate := time.Date(time.Now().Year(), parsedEndDate.Month(), parsedEndDate.Day()+1, 0, 0, 0, 0, time.Local)

	lastDate := (*s)[0].DateCreated.Format("01/02")
	first := true
	for _, v := range *s {
		thisDate := v.DateCreated.Format("01/02")
		if v.DateCreated.After(minDate) && v.DateCreated.Before(maxDate) {

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
	fmt.Println("done")
	return nil

}

func (s *TaskSlice) viewSpecificTime(c *config.Config) error {
	parsedDate, err := time.Parse("01/02", c.SpecificTime)
	if err != nil {
		return fmt.Errorf("unable to parse specific date: %w", err)
	}

	fmt.Printf("---%s---\n", parsedDate.Format("01/02"))
	for _, v := range *s {
		if v.DateCreated.Month() == parsedDate.Month() && v.DateCreated.Day() == parsedDate.Day() {
			fmt.Println(v.Pretty())
		}
	}
	return nil

}
