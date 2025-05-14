package add

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/tylerBrittain42/YesterdayI/internal/config"
	"github.com/tylerBrittain42/YesterdayI/internal/task"
)

func Add(f string, c *config.Config) error {
	if c.Content == "" {
		return errors.New("no content specified")
	}
	t := task.NewTask(c.Content, c.JiraTicket)

	tJson, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("error marshaling json: %w", err)
	}

	err = updateFile(f, tJson)
	err = os.WriteFile("taskLog.json", tJson, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println(string(tJson))
	return nil

}

func updateFile(f string, t []byte) error {
	// check if file exists and create it if it does not
	_, err := os.Stat(f)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("error checking file existence: %w", err)
		}
		file, err := os.Create(f)
		if err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}
		file.Close()
		fmt.Println("file successfully created")
	}

	// taskStream, err := os.ReadFile(f)
	//
	return nil
}
