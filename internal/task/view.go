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
	} else if c.StartTime != "" && c.EndTime == "" && c.SpecificTime == "" {
		tSlice, err = tSlice.filterUntilNow(c)
		if err != nil {
			return err
		}
	} else if c.StartTime != "" && c.EndTime != "" && c.SpecificTime == "" {
		tSlice, err = tSlice.filterRange(c)
		if err != nil {
			return err
		}
	} else if c.StartTime == "" && c.EndTime == "" && c.SpecificTime != "" {
		tSlice, err = tSlice.filterSpecificTime(c)
		if err != nil {
			return err
		}
	} else {
		return errors.New("unsupported combination of flags")

	}
	err = printOutput(tSlice)
	if err != nil {
		return fmt.Errorf("unable to print output: %w", err)
	}
	fmt.Println("done")

	return nil
}

func (s TaskSlice) filterUntilNow(c *config.Config) (TaskSlice, error) {
	filteredSlice := TaskSlice{}
	parsedDate, err := time.Parse("01/02", c.StartTime)
	if err != nil {
		return TaskSlice{}, fmt.Errorf("unable to parse start date: %w", err)
	}
	minDate := time.Date(time.Now().Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, time.Local)
	for _, v := range s {
		if v.DateCreated.After(minDate) {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice, nil
}

func (s TaskSlice) filterRange(c *config.Config) (TaskSlice, error) {
	filteredSlice := TaskSlice{}
	parsedStartDate, err := time.Parse("01/02", c.StartTime)
	if err != nil {
		return TaskSlice{}, fmt.Errorf("unable to parse start date: %w", err)
	}
	minDate := time.Date(time.Now().Year(), parsedStartDate.Month(), parsedStartDate.Day(), 0, 0, 0, 0, time.Local)

	parsedEndDate, err := time.Parse("01/02", c.EndTime)
	if err != nil {
		return TaskSlice{}, fmt.Errorf("unable to parse end date: %w", err)
	}
	maxDate := time.Date(time.Now().Year(), parsedEndDate.Month(), parsedEndDate.Day()+1, 0, 0, 0, 0, time.Local)

	for _, v := range s {
		if v.DateCreated.After(minDate) && v.DateCreated.Before(maxDate) {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice, nil

}

func (s TaskSlice) filterSpecificTime(c *config.Config) (TaskSlice, error) {
	filteredSlice := TaskSlice{}
	parsedDate, err := time.Parse("01/02", c.SpecificTime)
	if err != nil {
		return TaskSlice{}, fmt.Errorf("unable to parse specific date: %w", err)
	}

	for _, v := range s {
		if v.DateCreated.Month() == parsedDate.Month() && v.DateCreated.Day() == parsedDate.Day() {
			filteredSlice = append(filteredSlice, v)
		}
	}
	return filteredSlice, nil

}

func printOutput(s TaskSlice) error {
	if len(s) == 0 {
		fmt.Printf("no entries found for the given criteria\n")
		return nil
	}
	curDate := s[0].DateCreated
	prettyCurDate := prettyDate(curDate)
	fmt.Printf("---%s---\n", prettyCurDate)
	for _, v := range s {
		if v.DateCreated.Month() != curDate.Month() || v.DateCreated.Day() != curDate.Day() {
			curDate = v.DateCreated
			prettyCurDate = prettyDate(curDate)
			fmt.Printf("---%s---\n", prettyCurDate)
		}
		fmt.Print(v.Pretty())
		fmt.Println()

	}

	return nil
}

func prettyDate(t time.Time) string {
	return fmt.Sprintf("%v/%v", int(t.Month()), t.Day())
}
