package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tylerBrittain42/YesterdayI/internal/config"
	"github.com/tylerBrittain42/YesterdayI/internal/task"
)

const fileName = "taskLog.json"

func main() {

	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	var conf config.Config
	addCmd.StringVar(&conf.Content, "t", "", "required task")
	addCmd.StringVar(&conf.JiraTicket, "j", "", "optional jira ticket")

	// viewCmd := flag.NewFlagSet("view", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("No sub command given")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		err := task.AddTask(fileName, &conf)
		if err != nil {
			fmt.Printf("Unable to add task: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("task added")

	case "view":
		fmt.Println("selected view")
	default:
		fmt.Println("unexpected sub command given")
		os.Exit(1)
	}

}
