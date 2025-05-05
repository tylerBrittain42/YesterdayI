package add

import (
	"fmt"

	"github.com/tylerBrittain42/YesterdayI/internal/config"
)

func Add(c *config.Config) {
	fmt.Println()
	fmt.Printf("Starting add with jira=%v and task=%v\n", c.JiraTicket, c.Task)

}
